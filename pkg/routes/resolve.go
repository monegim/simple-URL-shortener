package routes

import (
	"log"
	"net/http"
)

func ResolveHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Resolved called with the:", r.URL.Path)
}
