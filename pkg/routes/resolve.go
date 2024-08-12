package routes

import (
	"fmt"
	"log"
	"net/http"
	"simple-url-shortener/pkg/database"
	"strings"

	"github.com/redis/go-redis/v9"
)

func ResolveHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	ip, _ := getIp(r.RemoteAddr)
	log.Println("IP", ip)
	log.Println("Resolved called with the:", r.URL.Path)
	requestURI := strings.TrimPrefix(r.RequestURI, "/")
	uri := strings.Split(requestURI, "/")
	log.Println("uri:",len(uri))
	if len(uri) != 1  || requestURI == "" {
		w.Write([]byte(fmt.Sprintf("%s is not valid", r.RequestURI)))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id := uri[0]
	url, err := c.Get(database.Ctx, id).Result()
	if err == redis.Nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(url))
	log.Println("requestURI", url)
}
