package main

import (
	"log"
	"net/http"
	"testing"
)

type server int

func (*server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
	write, err := w.Write([]byte("Hello World!"))
	log.Println(write)
	log.Println(err)
}

func TestServeHTTP(t *testing.T) {
	var s server
	http.ListenAndServe("localhost:9999", &s)
}
