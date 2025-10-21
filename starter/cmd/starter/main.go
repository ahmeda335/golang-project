package main

import (
	"log"
	"net/http"
)

func RegisterControllers() {
	// minimal controller registration so the server runs without the missing package
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	})
}

func main() {
	RegisterControllers()
	log.Println("starting server on :3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}
