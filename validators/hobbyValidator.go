package validators

import (
	"github.com/Dasha-Kinsely/mega-narrator/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type HobbyValidator struct {
	Check struct {

	} `json:"check"`
	models.Hobby
}

func NewHobbyValidator() HobbyValidator{
	return HobbyValidator{}
}

func (m *HobbyValidator) BindContext(c *gin.Context) error {
	headerMethod := c.Request.Method
	headerContentType := c.ContentType()
	if err := c.ShouldBindWith(m, binding.Default(headerMethod, headerContentType)); err != nil {
		return err
	}
	return nil
}