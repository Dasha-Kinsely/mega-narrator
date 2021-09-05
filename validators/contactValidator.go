package validators

import (
	"github.com/Dasha-Kinsely/mega-narrator/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type ContactValidator struct {
	Check struct {
		Title string `form:"title" json:"title" binding:"required,max=255"`
		LinkedIn string `form:"linkedin" json:"linkedin" binding:"max=255"`
		Facebook string `form:"facebook" json:"facebook" binding:"max=255"`
		Twitter string `form:"twitter" json:"twitter" binding:"max=255"`
		Github string `form:"github" json:"github" binding:"max=255"`
		Phone string `form:"phone" json:"phone" binding:"required,max=255"`
	} `json:"check"`
	ValidatedContact models.Contact `json:"-"`
}

func NewContactValidator() ContactValidator{
	return ContactValidator{}
}

func (m *ContactValidator) BindContext(c *gin.Context) error {
	headerMethod := c.Request.Method
	headerContentType := c.ContentType()
	if err := c.ShouldBindWith(m, binding.Default(headerMethod, headerContentType)); err != nil {
		return err
	}
	m.ValidatedContact.Title = m.Check.Title
	m.ValidatedContact.LinkedIn = m.Check.LinkedIn
	m.ValidatedContact.Facebook = m.Check.Facebook
	m.ValidatedContact.Twitter = m.Check.Twitter
	m.ValidatedContact.Github = m.Check.Github
	m.ValidatedContact.Phone = m.Check.Phone
	return nil
}