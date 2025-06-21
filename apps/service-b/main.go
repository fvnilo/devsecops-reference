package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fvnilo/devsecops-reference/libs/reverse_go"
)

func main() {
	http.HandleFunc("/b", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, reverse_go.String("Hello from Service B"))
	})

	log.Println("Service B listening on :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
