package validators

import (
	"github.com/Dasha-Kinsely/mega-narrator/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type ExperienceValidator struct {
	Check struct {

	} `json:"check"`
	ValidatedExperience models.Experience `json:"-"`
}

func NewExperienceValidator() ExperienceValidator{
	return ExperienceValidator{}
}

func (m *ExperienceValidator) BindContext(c *gin.Context) error {
	headerMethod := c.Request.Method
	headerContentType := c.ContentType()
	if err := c.ShouldBindWith(m, binding.Default(headerMethod, headerContentType)); err != nil {
		return err
	}
	return nil
}