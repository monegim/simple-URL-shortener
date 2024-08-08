package routes

import (
	"log"
	"net/http"
	"simple-url-shortener/pkg/database"

	"github.com/redis/go-redis/v9"
)

type Server struct {
	Address string
}

func NewServer(addr string) *Server {
	return &Server{
		Address: addr,
	}
}

var c *redis.Client

func init() {
	dbNo := 0
	c = database.Client(dbNo)
	
}

func (s *Server) Start() {
	log.Printf("Starting server on %s", s.Address)
	log.Fatal(http.ListenAndServe(s.Address, nil))
}
