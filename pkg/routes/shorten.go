package routes

import (
	"encoding/json"
	"log"
	"math/rand"
	"net"
	"net/http"
	"simple-url-shortener/pkg"
	"simple-url-shortener/pkg/database"
	"simple-url-shortener/urlverifier"
	"strconv"
)

type Data struct {
	URL string `json:"url"`
}

func ShortenHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("ShortenHandler is called")
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request"))
		return
	}
	defer r.Body.Close()

	ip, err := getIp(r.RemoteAddr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request"))
		return
	}
	c.Get(database.Ctx, ip)

	var url Data
	err = json.NewDecoder(r.Body).Decode(&url)
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
	id := rand.Uint64()
	shortenedURL := pkg.URLShortener(id, url.URL)
	log.Println(shortenedURL)

	err = c.Set(database.Ctx, strconv.FormatUint(id, 10), shortenedURL, 0).Err()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func getIp(remoteAddr string) (string, error) {
	ip, _, err := net.SplitHostPort(remoteAddr)
	if err != nil {
		return "", err
	}
	return ip, nil
}
