package validators

import (
	"github.com/Dasha-Kinsely/mega-narrator/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type BlogValidator struct {
	Check struct {

	} `json:"check"`
	models.Blog
}

func NewBlogValidator() BlogValidator{
	return BlogValidator{}
}

func (m *BlogValidator) BindContext(c *gin.Context) error {
	headerMethod := c.Request.Method
	headerContentType := c.ContentType()
	if err := c.ShouldBindWith(m, binding.Default(headerMethod, headerContentType)); err != nil {
		return err
	}
	return nil
}