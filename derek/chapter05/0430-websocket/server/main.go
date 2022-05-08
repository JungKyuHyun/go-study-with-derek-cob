package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Listening on port :8000")
	log.Panic(http.ListenAndServe("localhost:8000", http.HandlerFunc(wsHandler)))
}