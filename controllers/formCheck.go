package controllers

import (
	"github.com/Dasha-Kinsely/mega-narrator/serializers/errorserializers"
	"github.com/Dasha-Kinsely/mega-narrator/validators"
	"github.com/gin-gonic/gin"
)

func FormCheck(c *gin.Context) {
	step := c.Param("step")
	// Workflow: 1) Check form binding -> 2) Handle form-data & db-dao & response-serialization 
	switch step{ 
	case "user":
		newUser := validators.NewUserValidator()
		if err := newUser.BindContext(c); err != nil {
			errorserializers.ContextJSON(c, "form binding", "failed to bind user form")
			return
		}
	case "education":
		newEducation := validators.NewEducationValidator()
		if err := newEducation.BindContext(c); err != nil {
			errorserializers.ContextJSON(c, "form binding", "failed to bind education form")
			return
		}
	case "skill":
		newSkill := validators.NewSkillValidator()
		if err := newSkill.BindContext(c); err != nil {
			errorserializers.ContextJSON(c, "form binding", "failed to bind skill form")
			return
		}
	case "experience":
		newExperience := validators.NewExperienceValidator()
		if err := newExperience.BindContext(c); err != nil {
			errorserializers.ContextJSON(c, "form binding", "failed to bind experience form")
			return
		}
	case "showcase":
		newShowcase := validators.NewShowcaseValidator()
		if err := newShowcase.BindContext(c); err != nil {
			errorserializers.ContextJSON(c, "form binding", "failed to bind showcase form")
			return
		}
	case "album":
		newAlbum := validators.NewAlbumValidator()
		if err := newAlbum.BindContext(c); err != nil {
			errorserializers.ContextJSON(c, "form binding", "failed to bind album form")
			return
		}
	case "blog":
		newBlog := validators.NewBlogValidator()
		if err := newBlog.BindContext(c); err != nil {
			errorserializers.ContextJSON(c, "form binding", "failed to bind blog form")
			return
		}
	case "hobby":
		newHobby := validators.NewHobbyValidator()
		if err := newHobby.BindContext(c); err != nil {
			errorserializers.ContextJSON(c, "form binding", "failed to bind hobby form")
			return
		}
	case "contact":
		newContact := validators.NewContactValidator()
		if err := newContact.BindContext(c); err != nil {
			errorserializers.ContextJSON(c, "form binding", "failed to bind contact form")
			return
		}
	default:
		errorserializers.ContextJSON(c, "form not found", "finding form")
		return
	}
}