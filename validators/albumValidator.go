package validators

import (
	"github.com/Dasha-Kinsely/mega-narrator/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type AlbumValidator struct {
	Check struct {

	} `json:"check"`
	models.Album
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
	return nil
}