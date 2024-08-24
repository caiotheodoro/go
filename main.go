package main

import (
	"fmt"
	"net/http"
)

func main() {
	server := &http.Server{
		Addr:    ":3000",
		Handler: http.HandlerFunc(handler),
	}

	if err := server.ListenAndServe(); err != nil {
		fmt.Println("Server error: ", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		if r.URL.Path == "/hello" {
			w.Write([]byte("Hello, World!"))
			return
		}
	}

	if r.Method == http.MethodPost {
		w.Write([]byte("Hello, World!"))
		return
	}

	w.Write([]byte("Hello, World!"))
}
