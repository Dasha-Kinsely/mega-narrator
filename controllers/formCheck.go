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
	case "experience":
	case "showcase":
	case "album":
	case "blog":
	case "hobby":
	case "contact":
	default:
		errorserializers.ContextJSON(c, "form not found", "finding form")
		return
	}
}