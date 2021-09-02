package setups

import (
	"github.com/Dasha-Kinsely/mega-narrator/middlewares"
	fileadapter "github.com/casbin/casbin/v2/persist/file-adapter"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type Server struct {
	Router *gin.Engine
	Context *gin.Context
	RedisClient *redis.Client
	FileAdapter *fileadapter.Adapter
	TokenInterface *middlewares.TokenInterface
}

func (server *Server) InitializeServer(redis_host, redis_port, redis_pass string) {
	// Init gin
	server.Router = gin.Default()
	server.Router.MaxMultipartMemory = 512 << 20
	// Init redis
	server.RedisClient = NewRedisDB(redis_host, redis_port, redis_pass)
	// Init casbin
	//server.FileAdapter = fileadapter.NewAdapter("config/")
	// Init routes
	server.InitializeRoutes()
}