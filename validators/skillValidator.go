package validators

import (
	"mime/multipart"

	"github.com/Dasha-Kinsely/mega-narrator/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type SkillValidator struct {
	Check struct {
		Title string `form:"title" json:"title" binding:"required,max=255"`
		Icon *multipart.FileHeader `form:"icon" json:"icon" binding:"max=255"`
		AcquiredYears string `form:"acquiredyears" json:"acquiredyears" binding:"max=255"`
		Fluency int8 `form:"fluency" json:"fluency" binding:"required,max=255"`
		Description string `form:"description" json:"description" binding:"max=255"`
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
	m.ValidatedSkill.Title = m.Check.Title
	m.ValidatedSkill.AcquiredYears = m.Check.AcquiredYears
	m.ValidatedSkill.Icon = &m.Check.Icon.Filename
	m.ValidatedSkill.Fluency = m.Check.Fluency
	m.ValidatedSkill.Description = m.Check.Description
	return nil
}