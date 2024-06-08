package routes

import (
	"fmt"
	"io"
	"net/http"
)

func Root(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "this is my website")
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got Hello/ request\n")
	io.WriteString(w, "Hello there")
}