package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"simple-url-shortener/pkg"
	"simple-url-shortener/urlverifier"
)

type Data struct {
	URL string `json:"url"`
}

func ShortenHandler(w http.ResponseWriter, r *http.Request) {
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
	// Check if url is valid
	verifier := urlverifier.NewVerifier(url.URL)
	err = verifier.Verify()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Input is valid
	shortenedURL := pkg.URLShortener(url.URL)
	log.Println(shortenedURL)
}
