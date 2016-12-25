package main

import (
	"log"
	"net/http"
	"testMicroservice/Data"
)

func main() {
	Data.InitUsers()
	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
