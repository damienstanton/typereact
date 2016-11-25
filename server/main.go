package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var port = os.Getenv("PORT")

func main() {
	r := mux.NewRouter()
	s := http.FileServer(http.Dir("../"))
	r.PathPrefix("/").Handler(s)
	http.Handle("/", r)
	if port == "" {
		port = "8000"
	}
	fmt.Printf("Listening on port %v...\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("problem :%v", err)
	}
}
