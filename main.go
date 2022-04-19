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
	http.ListenAndServe(":"+port, nil)
}
