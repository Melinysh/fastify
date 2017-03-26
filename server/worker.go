package main

// A go program to run instead of main.go on the other worker servers in the cluster

import (
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
			return
		}
	}
}

func Hash(s string) string {
	h := sha256.New()
	io.WriteString(h, s)
	return hex.EncodeToString(h.Sum(nil))
}
