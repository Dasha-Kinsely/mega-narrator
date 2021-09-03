package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Signin(c *gin.Context){
	c.JSON(http.StatusOK, "passed")
}

func AuthRoutes(r *gin.RouterGroup) {
	r.POST("/signin", Signin)
}