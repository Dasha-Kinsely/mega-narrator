package errorserializers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ContextJSON(c *gin.Context, issue string, step string) {
	switch issue{
		case "load env":
			log.Println(step)
			c.JSON(http.StatusExpectationFailed, "Error occurred during env file loading")
			return
		case "upload":
			log.Println(step)
			c.JSON(http.StatusBadRequest, "Error occurred during file uploading")
			return
		default:
			c.JSON(http.StatusInternalServerError, "Error < Unknown >")
			return
	}
}