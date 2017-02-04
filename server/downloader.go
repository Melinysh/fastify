package main

import (
	"hash/fnv"
	"io/ioutil"
	"log"
	"net/http"
)

// Download gets the file and saves it
func Download(link string, ch chan string) {
	resp, err := http.Get(link)
	if err != nil {
		log.Println("ERROR: fetching link", link, err)
		ch <- "ERROR"
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	filename = hash(link)
	err = iotutil.WriteFile(filename, body, 777)
	if err != nil {
		log.Println("ERROR: enable to save file", string(body), err)
		ch <- "ERROR"
	}
	ch <- filename
}

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}
