package main

import (
	"log"
	"net/http"
)

func main() {
	log.Print("Run TestServer 127.0.0.1...")
	http.ListenAndServe(":80", http.FileServer(http.Dir(".")))

}
