package helpers

import (
	"path/filepath"

	"github.com/dasha-kinsely/meganarrator/serializers/errorserializers"
	"github.com/gin-gonic/gin"
)

func Upload(c *gin.Context, count string) error {
	// reading file, make sure it exists. Handle singlular and multiple files differently
	if count == "single" {
		file, err := c.FormFile("attachment")
		if err != nil {
			errorserializers.ContextJSON(c, "upload", "reading file upload")
			return err
		}
		// saving file to path
		filename := filepath.Base(file.Filename)
		if err := c.SaveUploadedFile(file, filename); err != nil {
			errorserializers.ContextJSON(c, "upload", "saving file upload")
			return err
		}
	} else if count == "multiple" {
		form, err := c.MultipartForm()
		if err != nil {
			errorserializers.ContextJSON(c, "upload", "reading file upload")
			return err
		}
		files := form.File["files"]
		// saving file to path, one by one
		for _, file := range files {
			filename := filepath.Base(file.Filename)
			if err := c.SaveUploadedFile(file, filename); err != nil {
				errorserializers.ContextJSON(c, "upload", "saving file upload")
				return err
			}
		}
	}
	return nil
}