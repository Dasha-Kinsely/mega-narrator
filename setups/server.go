package models

import (
	"github.com/Dasha-Kinsely/mega-narrator/middlewares"
	fileadapter "github.com/casbin/casbin/v2/persist/file-adapter"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type Server struct {
	Router *gin.Engine
	RedisClient *redis.Client
	FileAdapter *fileadapter.Adapter
	TokenInterface *middlewares.TokenInterface
}

var ActiveServer Server

func (server *Server) InitializeServer(redis_host, redis_port, redis_pass string) {
	server.Router = gin.Default()
	server.RedisClient = NewRedisDB()
}