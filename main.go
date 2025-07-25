package main

import (
	"log"
	"net/http"
)

func main() {

	http.Handle("/", http.FileServer(http.Dir("public"))) // by default - relative path is used here

	const addr = ":8080"
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		// not working - error
		log.Fatalf("Server failed %v", err)
	}
}
