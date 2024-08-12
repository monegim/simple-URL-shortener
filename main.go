package main

import (
	"net/http"
	"os"
	"simple-url-shortener/pkg/routes"
)

func main() {
	addr := os.Getenv("ADDRESS")
	if addr == "" {
		addr = ":8081"
	}
	http.HandleFunc("/", routes.ResolveHandler)
	http.HandleFunc("/api/v1", routes.ShortenHandler)
	s := routes.NewServer(addr)
	s.Start()
}
