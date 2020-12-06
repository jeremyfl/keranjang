package internal

import (
	"github.com/go-redis/redis/v8"
)

// NewDatabaseClient new database of redis client
func NewDatabaseClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
