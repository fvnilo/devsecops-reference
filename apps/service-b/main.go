package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/a", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello from Service B")
	})

	log.Println("Service B listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
