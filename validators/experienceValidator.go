package validators

import (
	"time"

	"github.com/Dasha-Kinsely/mega-narrator/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type ExperienceValidator struct {
	Check struct {
		Title string `form:"title" json:"title" binding:"required,max=255"`
		Organization string `form:"organization" json:"organization" binding:"required,max=255"`
		StartTime time.Time `form:"starttime" json:"starttime" binding:"required" time_utc:"1" time_format:"2006-01-02"`
		EndTime time.Time `form:"endtime" json:"endtime" binding:"required" time_utc:"1" time_format:"2006-01-02"`
		Responsibility string `form:"responsibility" json:"responsibility" binding:"required,max=255"`
		ReferenceContact string `form:"referencecontact" json:"referencecontact" binding:"max=255"`
	} `json:"check"`
	ValidatedExperience models.Experience `json:"-"`
}

func NewExperienceValidator() ExperienceValidator{
	return ExperienceValidator{}
}

func (m *ExperienceValidator) BindContext(c *gin.Context) error {
	headerMethod := c.Request.Method
	headerContentType := c.ContentType()
	if err := c.ShouldBindWith(m, binding.Default(headerMethod, headerContentType)); err != nil {
		return err
	}
	m.ValidatedExperience.Title = m.Check.Title
	m.ValidatedExperience.Organization = m.Check.Organization
	m.ValidatedExperience.StartTime = m.Check.StartTime
	m.ValidatedExperience.EndTime = m.Check.EndTime
	m.ValidatedExperience.Responsibility = m.Check.Responsibility
	m.ValidatedExperience.ReferenceContact = m.Check.ReferenceContact
	return nil
}