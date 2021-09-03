package errorserializers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ContextJSON(c *gin.Context, issue string, step string) {
	log.Println(step)
	switch issue{
		case "load env":
			c.JSON(http.StatusExpectationFailed, "Error occurred during env file loading")
			return
		case "upload":
			c.JSON(http.StatusBadRequest, "Error occurred during file uploading")
			return
		case "form not found":
			c.JSON(http.StatusNotFound, "Controllers associated with this route was not found")
			return
		case "form binding":
			c.JSON(http.StatusBadRequest, "Failed during form checking process")
			return
		default:
			c.JSON(http.StatusInternalServerError, "Error < Unknown >")
			return
	}
}