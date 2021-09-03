package controllers

import (
	"github.com/Dasha-Kinsely/mega-narrator/serializers/errorserializers"
	"github.com/Dasha-Kinsely/mega-narrator/validators"
	"github.com/gin-gonic/gin"
)

func FormCheck(c *gin.Context) {
	step := c.PostForm("step")
	// Workflow: 1) Check form binding -> 2) Handle form-data & db-dao & response-serialization 
	switch step{ 
	case "user":
		newUser := validators.NewUserValidator()
		if err := newUser.BindContext(c); err != nil {
			errorserializers.ContextJSON(c, "form binding", "failed to bind form")
			return
		}
	case "education":
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