package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	uri := "http://ajax.googleapis.com/ajax/services/search/web?v=1.0&rsz=8&q=golang"
	resp, err := http.Get(uri)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}
	fmt.Println(resp.Status)
}
