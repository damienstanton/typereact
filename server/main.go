package main

import (
	"net/http"

	"log"
)

func main() {
	fs := fileServer()
	log.Print("Listening on port 8000")
	log.Fatal(http.ListenAndServe(":8000", fs))
}

func fileServer() http.Handler {
	return http.FileServer(http.Dir("../"))
}
