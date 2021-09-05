package validators

import (
	"mime/multipart"

	"github.com/Dasha-Kinsely/mega-narrator/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type AlbumValidator struct {
	Check struct {
		Title string `form:"title" json:"title" binding:"required"`
		Photo1 *multipart.FileHeader `form:"photo1" json:"photo1" binding:"required"`
		Photo2 *multipart.FileHeader `form:"photo2" json:"photo2"`
		Photo3 *multipart.FileHeader `form:"photo3" json:"photo3"`
		Photo4 *multipart.FileHeader `form:"photo4" json:"photo4"`
	} `json:"check"`
	ValidatedAlbum models.Album `json:"-"`
}

func NewAlbumValidator() AlbumValidator{
	return AlbumValidator{}
}

func (m *AlbumValidator) BindContext(c *gin.Context) error {
	headerMethod := c.Request.Method
	headerContentType := c.ContentType()
	if err := c.ShouldBindWith(m, binding.Default(headerMethod, headerContentType)); err != nil {
		return err
	}
	m.ValidatedAlbum.Title = m.Check.Title
	m.ValidatedAlbum.Photo1 = &m.Check.Photo1.Filename
	m.ValidatedAlbum.Photo2 = &m.Check.Photo2.Filename
	m.ValidatedAlbum.Photo3 = &m.Check.Photo3.Filename
	m.ValidatedAlbum.Photo4 = &m.Check.Photo4.Filename
	return nil
}