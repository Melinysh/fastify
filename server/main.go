package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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

	filedata, err := ioutil.ReadFile(string(data))
	if err != nil {
		log.Println(err)
		fmt.Fprintln(w, err.Error())
		return
	}
	fmt.Fprintln(w, string(filedata))

	/*	c := make(chan string)
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
		}*/
}
