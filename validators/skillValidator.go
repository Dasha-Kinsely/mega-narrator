package validators

import (
	"github.com/Dasha-Kinsely/mega-narrator/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type SkillValidator struct {
	Check struct {

	} `json:"check"`
	ValidatedSkill models.Skill `json:"-"`
}

func NewSkillValidator() SkillValidator{
	return SkillValidator{}
}

func (m *SkillValidator) BindContext(c *gin.Context) error {
	headerMethod := c.Request.Method
	headerContentType := c.ContentType()
	if err := c.ShouldBindWith(m, binding.Default(headerMethod, headerContentType)); err != nil {
		return err
	}
	return nil
}