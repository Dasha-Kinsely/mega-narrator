package routers

import (
	"github.com/Dasha-Kinsely/mega-narrator/helpers"
	"github.com/gin-gonic/gin"
)

// Mostly consisted of helper function routes
func MiscRoutes(r *gin.RouterGroup) {
	// Upload Single File
	r.POST("/upload/single", func(c *gin.Context){
		if err := helpers.Upload(c, "single"); err != nil {
			return
		}
	})
	// Upload Multiple File
	r.POST("/upload/multiple", func(c *gin.Context){
		if err := helpers.Upload(c, "multiple"); err != nil {
			return
		}
	})
}