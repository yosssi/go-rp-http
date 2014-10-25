package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Raspberry Pi HTTP server")
}

func main() {
	http.HandleFunc("/", handler)
	if err := http.ListenAndServe(":80", nil); err != nil {
		panic(err)
	}
}
