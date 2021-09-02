package helpers

import (
	"os"

	"github.com/dasha-kinsely/mega-narrator/serializers/errorserializers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

/* This will return a string value based on input key, you may user strconv.ParseInt or strconv.ParseBool
to make further modification */
func LoadEnvVarAsString(c *gin.Context, step string, key string) (value string, err error) {
	if err := godotenv.Load(); err != nil {
		errorserializers.ContextJSON(c, "load env", step)
		return "", err
	}
	return os.Getenv(key), nil
}