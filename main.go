package main

import (
	"log"
	"net/http"
)

func main() {
	const addr = ":8080"
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		// not working - error
		log.Fatalf("Server failed %v", err)
	}
}
