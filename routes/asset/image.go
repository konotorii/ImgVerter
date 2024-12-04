package asset

import (
	"imgverter/util"
	_ "io"
	"os"
	"path"
	"strings"

	"github.com/gabriel-vasile/mimetype"
	"github.com/gin-gonic/gin"
	"github.com/konotorii/go-consola"
)

func FetchImage(c *gin.Context) {
	fileId := c.Param("id")

	filePath := util.Config.PublicFolder

	filePath = path.Clean(filePath + "/i/" + fileId)

	existsWebP := util.WebpExists(filePath)

	if existsWebP {
		filePath = strings.Replace(filePath, path.Ext(filePath), ".webp", -1)
	} else {
		link, err := util.EncodeWebP(filePath)
		if err != nil {
			consola.Error(err)
		}
		if link != nil {
			filePath = strings.Replace(filePath, path.Ext(filePath), ".webp", -1)
		}
	}

	mtype, err := mimetype.DetectFile(filePath)

	if err != nil {
		consola.Error("Getting mimetype error", err)

		c.Status(500)
	}

	buf, err := os.ReadFile(filePath)
	if err != nil {
		consola.Error(err)
		c.Status(500)
		return
	}

	c.Data(200, mtype.String(), buf)
}
