package main

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

// Download gets the file and saves it
func Download(link string, ch chan string) {
	log.Println("Got request to download", link)
	resp, err := http.Get(link)
	if err != nil {
		log.Println("ERROR: fetching link", link, err)
		ch <- "ERROR"
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	log.Println("Got body for link", link)
	filename := hash(link)
	err = ioutil.WriteFile(filename, body, 0777)
	if err != nil {
		log.Println("ERROR: enable to save file", link, len(body), err)
		ch <- "ERROR"
	}
	log.Println("Got file saved", filename)
	ch <- filename
}

func hash(s string) string {
	h := sha256.New()
	io.WriteString(h, s)
	return hex.EncodeToString(h.Sum(nil))
}
