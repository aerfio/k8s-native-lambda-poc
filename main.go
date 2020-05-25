package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("handling request from %s....", r.Host)
		_, _ = fmt.Fprintf(w, "Hi there, I love %s!", "nature")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("starting server...")
	log.Println("REDIS_PORT:", os.Getenv("REDIS_PORT"))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
