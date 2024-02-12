package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("ping pong server running...")

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		// Write "pong" back to the client
		_, err := fmt.Fprintf(w, "pong\n")
		if err != nil {
			return
		}
	})

	fmt.Println("Server is running on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
