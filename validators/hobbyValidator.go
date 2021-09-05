package validators

import (
	"github.com/Dasha-Kinsely/mega-narrator/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type HobbyValidator struct {
	Check struct {
		Title string `form:"title" json:"title" binding:"required"`
		AltDescription string `form:"altdescription" json:"altdescription" binding:"required"`
	} `json:"check"`
	ValidatedHobby models.Hobby `json:"-"`
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
	m.ValidatedHobby.Title = m.Check.Title
	m.ValidatedHobby.AltDescription = m.Check.AltDescription
	return nil
}