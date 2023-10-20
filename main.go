package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, Go HTTP Server!")
	})

	port := ":8080"
	fmt.Printf("Server is running on port %s...\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
