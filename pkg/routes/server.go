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
	status := c.Ping(database.Ctx)
	log.Printf("Connecting to the db Number %d\n", dbNo)
	if status.Err() != nil {
		log.Fatalf("Could not connect to the DB: %s",status.Err().Error())
	}
	log.Println("Connected to the DB")
}

func (s *Server) Start() {
	log.Printf("Starting server on %s", s.Address)
	log.Fatal(http.ListenAndServe(s.Address, nil))
}
