package main

import (
	"github.com/guillego/cajita/internal/handler"
	"github.com/guillego/cajita/internal/store"
	"log"
	"net/http"
)

func main() {
	s := store.NewStore()
	h := handler.NewHandler(s)

	http.HandleFunc("/get", h.GetHandler)
	http.HandleFunc("/set", h.SetHandler)
	http.HandleFunc("/delete", h.DeleteHandler)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}

