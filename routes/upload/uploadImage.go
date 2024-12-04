package upload

import (
	"fmt"
	"imgverter/util"
	"path"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/konotorii/go-consola"
)

func PostImage(c *gin.Context) {
	key := c.Query("key")

	if key == util.Config.UploadKey {
		file, _ := c.FormFile("file")
		consola.Log("Uploaded file: " + file.Filename)

		id, err := uuid.NewV4()
		if err != nil {
			consola.Error(err)
			c.Status(500)
			return
		}

		ext := path.Ext(file.Filename)
		fileName := fmt.Sprintf("%s%s", id, ext)
		webFileName := fmt.Sprintf("%s%s", id, ext)
		filePath := path.Clean(fmt.Sprintf("%s/i/%s%s", util.Config.PublicFolder, id, ext))

		err = c.SaveUploadedFile(file, filePath)
		if err != nil {
			consola.Error(err)
			c.Status(500)
			return
		}

		link, err := util.EncodeWebP(filePath)
		if err != nil {
			consola.Error(err)
		}
		if link != nil {
			webFileName = strings.Replace(webFileName, ext, ".webp", -1)
		}

		c.JSON(200, gin.H{"url": fmt.Sprintf("https://%s/i/%s", util.Config.Domain, fileName), "id": id, "fileName": fileName, "ext": ext, "webUrl": webFileName})
	} else {
		c.Status(400)
	}
}
