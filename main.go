package main

import (
	"log"
	"net/http"
	"os"
)

type Server struct {
	Address string
}

func NewServer(addr string) *Server {
	return &Server{
		Address: addr,
	}
}

func (s *Server) Start() {
	log.Fatal(http.ListenAndServe(s.Address, nil))
}

func main() {
	addr := os.Getenv("ADDRESS")
	if addr == "" {
		addr = ":8080"
	}
	http.HandleFunc("/", mainHandler)
	s := NewServer(addr)
	s.Start()

}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Main is called")
}
