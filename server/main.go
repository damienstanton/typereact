package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// freePort asks the kernel for an open port and returns it as a string
// based on github.com/phayes/freeport
func freePort() string {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		panic(err)
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		panic(err)
	}
	defer l.Close()
	return strconv.Itoa(l.Addr().(*net.TCPAddr).Port)
}

func main() {
	port := freePort()
	router := mux.NewRouter()
	dir := http.FileServer(http.Dir("../"))

	router.PathPrefix("/").Handler(dir)
	http.Handle("/", router)

	fmt.Printf("Listening for info on port %v...\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("problem :%v", err)
	}
}
