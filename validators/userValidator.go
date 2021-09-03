package validators

import (
	"github.com/Dasha-Kinsely/mega-narrator/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type UserValidator struct {
	Check struct {
		Photo string `form:"photo" json:"photo" binding:"required,max=255"`
		FirstName string `form:"firstname" json:"firstname" binding:"required,alphanum,max=255"`
		LastName string `form:"lastname" json:"lastname" binding:"required,alphanum,max=255"`
		TitlePhrase string `form:"titlephrase" json:"titlephrase" binding:"required,alphanum,max=255"`
		Citizenship string `form:"citizenship" json:"citizenship" binding:"required,alphanum,max=255"`
		Gender string `form:"gender" json:"gender" binding:"required,alphanum,max=255"`
		Languages string `form:"languages" json:"languages" binding:"required"`
		Birthdate string `form:"birthdate" json:"birthdate" binding:"required,max=255"`
		ResumeFile string `form:"resumefile" json:"resumefile" binding:"required,max=255"`
	} `json:"check"`
	ValidatedNewUser models.User `json:"-"`
}

func NewUserValidator() UserValidator {
	return UserValidator{}
} 

func (m *UserValidator) BindContext(c *gin.Context) error {
	headerMethod := c.Request.Method
	headerContentType := c.ContentType()
	if err := c.ShouldBindWith(m, binding.Default(headerMethod, headerContentType)); err != nil {
		return err
	}
	return nil
}