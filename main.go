package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"simple-url-shortener/urlverifier"
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
	log.Printf("Starting server on %s", s.Address)
	log.Fatal(http.ListenAndServe(s.Address, nil))
}

type Data struct {
	URL string `json:"url"`
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
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request"))
		return
	}
	defer r.Body.Close()
	var url Data
	err := json.NewDecoder(r.Body).Decode(&url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	verifier := urlverifier.NewVerifier(url.URL)
	err = verifier.Verify()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Input is valid
	
}
