package main

import (
	"github.com/dasha-kinsely/mega-narrator/helpers"
	"github.com/gin-gonic/gin"
)

func Init() {

}

func main() {
	r := gin.Default()
	r.MaxMultipartMemory = 512 << 20
	r.POST("/upload", func(c *gin.Context){
		if err := helpers.Upload(c, "single"); err != nil {
			return
		}
	})
	r.Run(":8080")
}