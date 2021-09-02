package setups

import (
	"github.com/Dasha-Kinsely/mega-narrator/helpers"
	"github.com/gin-gonic/gin"
)

func (server *Server) InitializeRoutes() {

	// Upload Single File
	server.Router.POST("/upload/single", func(c *gin.Context){
		if err := helpers.Upload(c, "single"); err != nil {
			return
		}
	})
	// Upload Multiple File
	server.Router.POST("/upload/multiple", func(c *gin.Context){
		if err := helpers.Upload(c, "multiple"); err != nil {
			return
		}
	})
	//server.Router.POST("/", )
}