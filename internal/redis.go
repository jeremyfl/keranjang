package internal

import (
	"os"

	"github.com/go-redis/redis/v8"
)

// NewDatabaseClient new database of redis client
func NewDatabaseClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     os.Getenv("DB_HOST_URL"),
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
