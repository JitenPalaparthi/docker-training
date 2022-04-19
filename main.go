package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Pong")
	})
	port := os.Getenv("PORT")
	if port == "" {
		port = "50052"
	}

	http.ListenAndServe(":"+port, nil)
}
