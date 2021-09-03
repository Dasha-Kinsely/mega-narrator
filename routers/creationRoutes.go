package routers

import (
	"net/http"

	"github.com/Dasha-Kinsely/mega-narrator/controllers"
	"github.com/gin-gonic/gin"
)

func HandleForm(c *gin.Context){
	controllers.FormCheck(c)
}

func HandleReview(c *gin.Context){
	c.JSON(http.StatusOK, "passed this route")
}

func CreationRoutes(r *gin.RouterGroup) {
	r.POST("/form", HandleForm)
	r.GET("/review", HandleReview)
}

/*
r.PUT("/user", )
	r.PUT("/education", )
	r.PUT("/skill")
	r.PUT("/showcase")
	r.PUT("/experience")
	r.PUT("/album")
	r.PUT("/blog")
	r.PUT("/hobby")
	r.PUT("/contact")
	*/