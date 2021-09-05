package validators

import (
	"log"
	"mime/multipart"
	"time"

	"github.com/Dasha-Kinsely/mega-narrator/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type EducationValidator struct {
	Check struct {
		Degree string `form:"degree" json:"" binding:"required,max=255"`
		InstitutePhoto *multipart.FileHeader `form:"institutephoto" json:"institutephoto" binding:"omitempty,url"`
		Institute string `form:"institute" json:"institute" binding:"required,max=255"`
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
	if err := c.ShouldBindWith(m, binding.Default(headerMethod, headerContentType)); err != nil {
		log.Println(err.Error())
		return err
	}
	m.ValidatedEducation.Degree = m.Check.Degree
	m.ValidatedEducation.Institute = m.Check.Institute
	m.ValidatedEducation.StartTime = m.Check.StartTime
	m.ValidatedEducation.EndTime = m.Check.EndTime
	m.ValidatedEducation.InstitutePhoto = &m.Check.InstitutePhoto.Filename
	return nil
}