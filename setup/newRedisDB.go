package setup

import (
	"log"

	"github.com/go-redis/redis/v8"
)

func NewRedisDB(host, port, pass string) *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr: host + ":" + port,
		Password: pass,
		DB: 0,
	})
	return redisClient
}

func (server *Server) Close() {
	if server.RedisClient != nil {
		if err := server.RedisClient.Close(); err != nil {
			log.Fatal(err)
		}
		server.RedisClient = nil
	}
}