package upload

import (
	"fmt"
	"imgverter/util"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/konotorii/go-consola"
)

func PostFile(c *gin.Context) {
	key := c.Query("key")

	if len(util.Config.UploadKey) < 6 {
		c.Status(403)
		return
	}

	if key == util.Config.UploadKey {
		file, _ := c.FormFile("file")
		consola.Log(file.Filename)

		id, err := uuid.NewV4()
		if err != nil {
			consola.Error(err)
			c.Status(500)
			return
		}

		ext := path.Ext(file.Filename)
		fileName := fmt.Sprintf("%s%s", id, ext)
		filePath := path.Clean(fmt.Sprintf("%s/files/%s%s", util.Config.PublicFolder, id, ext))

		err = c.SaveUploadedFile(file, filePath)
		if err != nil {
			consola.Error(err)
			c.Status(500)
		}

		c.JSON(200, gin.H{"url": fmt.Sprintf("https://%s/f/%s", util.Config.Domain, fileName), "id": id, "filename": fileName, "ext": ext})
	} else {
		c.Status(400)
	}
}
