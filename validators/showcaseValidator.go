package validators

import (
	"github.com/Dasha-Kinsely/mega-narrator/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type ShowcaseValidator struct {
	Check struct {

	} `json:"check"`
	ValidatedShowcase models.Showcase `json:"-"`
}

func NewShowcaseValidator() ShowcaseValidator{
	return ShowcaseValidator{}
}

func (m *ShowcaseValidator) BindContext(c *gin.Context) error {
	headerMethod := c.Request.Method
	headerContentType := c.ContentType()
	if err := c.ShouldBindWith(m, binding.Default(headerMethod, headerContentType)); err != nil {
		return err
	}
	return nil
}