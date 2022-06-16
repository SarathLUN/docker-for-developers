package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func index(w http.ResponseWriter, _ *http.Request) {
	_, err := io.WriteString(w, "Hello from docker container!")
	if err != nil {
		log.Println(err)
	}
}
