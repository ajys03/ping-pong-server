package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("ping pong server running...")

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		// Write "pong" back to the client
		fmt.Fprintf(w, "pong\n")
	})

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
