package main

import (
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
}
