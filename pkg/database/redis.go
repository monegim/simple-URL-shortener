package database

import (
	"os"

	"github.com/redis/go-redis/v9"
)

func Connect(dbNo int) *redis.Client {
	DB_ADDR := os.Getenv("DB_ADDR")
	if DB_ADDR == "" {
		DB_ADDR = "localhost:6379"
	}
	opts := redis.Options{
		Addr: DB_ADDR,
		DB:   dbNo,
	}
	return redis.NewClient(&opts)
}
