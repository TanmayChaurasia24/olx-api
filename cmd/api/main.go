package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("all ok"))
	})

	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		log.Fatal(fmt.Sprintf("Error starting server: %v", err))
	}
}
