package validators

import (
	"mime/multipart"

	"github.com/Dasha-Kinsely/mega-narrator/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type BlogValidator struct {
	Check struct {
		Title string `form:"title" json:"title" binding:"required,max=255"`
		Image *multipart.FileHeader `form:"image" json:"image" binding:"required"`
		Text string `form:"text" json:"text" binding:"required"`
		Link string `form:"link" json:"link" binding:"max=255"`
	} `json:"check"`
	ValidatedBlog models.Blog `json:"-"`
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
	m.ValidatedBlog.Title = m.Check.Title
	m.ValidatedBlog.Image = &m.Check.Image.Filename
	m.ValidatedBlog.Text = m.Check.Text
	m.ValidatedBlog.Link = m.Check.Link
	return nil
}