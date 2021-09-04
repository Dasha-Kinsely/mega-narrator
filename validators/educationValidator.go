package validators

import (
	"log"
	"time"

	"github.com/Dasha-Kinsely/mega-narrator/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type EducationValidator struct {
	Check struct {
		Degree string `form:"degree" json:"" binding:"required,max=255"`
		InstitutePhoto *string `form:"instituephoto" json:"instituephoto"`
		Institute string `form:"institue" json:"institue" binding:"required,max=255"`
		StartTime time.Time `form:"starttime" json:"starttime" time_utc:"1" time_format:"2006-01-02"`
		EndTime time.Time `form:"endtime" json:"endtime" time_utc:"1" time_format:"2006-01-02"`
	} `json:"check"`
	ValidatedEducation models.Education `json:"-"`
}

func NewEducationValidator() EducationValidator{
	return EducationValidator{}
}

func (m *EducationValidator) BindContext(c *gin.Context) error {
	headerMethod := c.Request.Method
	headerContentType := c.ContentType()
	log.Println(c.PostForm("starttime"))
	if err := c.ShouldBindWith(m, binding.Default(headerMethod, headerContentType)); err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}