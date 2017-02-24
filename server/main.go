package main

// Main server of the cluster would run this and not worker.go.

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
)

func main() {
	log.Println("Running on port 8080")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Panicln("reading response body", err)
	}
	log.Println(string(data))
	b := bytes.NewReader(data)
  
    // inform worker 1
	go func() {
		_, err = http.Post("http://ec2-54-84-75-4.compute-1.amazonaws.com:8080", "text/octet-stream", b)
		if err != nil {
			log.Println("error connecting to other server: (2)", err)
		}
	}()
  
    // inform worker 2
	b2 := bytes.NewReader(data)
	go func() {
		_, err = http.Post("http://ec2-52-206-125-141.compute-1.amazonaws.com:8080", "text/octet-stream", b2)
		if err != nil {
			log.Println("error connecting to other server (3) :", err)
		}
	}()

    // check cache for existing file 
	if _, err := os.Stat(Hash(string(data)) + ".torrent"); !os.IsNotExist(err) {
		fmt.Fprintf(w, "http://ec2-54-83-190-222.compute-1.amazonaws.com:81/"+Hash(string(data))+".torrent")
		log.Println("Returning cached value")
		err = exec.Command("/bin/bash", "-c", "/home/ubuntu/c-t.sh "+Hash(string(data))).Start()
		if err != nil {
			log.Panicln(err)
		}

		return
	}
	c := make(chan string)
	defer close(c)
	go Download(string(data), c)
	for {
		select {
		case filename, _ := <-c:
			err := exec.Command("/bin/bash", "-c", "/home/ubuntu/c-t.sh "+filename).Start()
			if err != nil {
				log.Panicln(err)
			}

			fmt.Fprintf(w, "http://ec2-54-83-190-222.compute-1.amazonaws.com:81/"+filename+".torrent")
			log.Println("Returned link to torrent for", filename)
			return
		}
	}
}

func Hash(s string) string {
	h := sha256.New()
	io.WriteString(h, s)
	return hex.EncodeToString(h.Sum(nil))
}
