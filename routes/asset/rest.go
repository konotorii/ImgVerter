package asset

import (
	"fmt"
	"imgverter/util"
	"net/url"
	"os"
	"path"
	"strings"

	"github.com/gabriel-vasile/mimetype"
	"github.com/gin-gonic/gin"
	"github.com/konotorii/go-consola"
)

func FetchRest(c *gin.Context) {
	fileId, _ := url.PathUnescape(c.Request.URL.String())

	enableWebP := c.Query("webp")

	u, _ := url.Parse(fileId)
	result := u.Path

	filePath := util.Config.PublicFolder

	filePath = path.Clean(filePath + result)

	mtype, err := mimetype.DetectFile(filePath)

	if err != nil {
		consola.Error("Getting mimetype error", err)

		c.Status(500)
		return
	}

	if strings.Contains(mtype.String(), "image") {
		existsWebP := util.WebpExists(filePath)

		if util.Config.UploadSettings.EnableWebpConversion || enableWebP == "true" {
			if existsWebP {
				filePath = strings.Replace(filePath, path.Ext(filePath), ".webp", -1)

				mtype, err = mimetype.DetectFile(filePath)
				if err != nil {
					consola.Error("Getting mimetype error", err)

					c.Status(500)
					return
				}
			} else {
				link, err := util.EncodeWebP(filePath)
				if err != nil {
					consola.Error(err)
				}
				if link != nil {
					filePath = strings.Replace(filePath, path.Ext(filePath), ".webp", -1)

					mtype, err = mimetype.DetectFile(filePath)
					if err != nil {
						consola.Error("Getting mimetype error", err)

						c.Status(500)
						return
					}
				}
			}
		}
	}

	consola.Info(fmt.Sprintf("Fetching %s | %s | %s", result, c.Request.UserAgent(), c.RemoteIP()))

	buf, err := os.ReadFile(filePath)
	if err != nil {
		consola.Error(err)
		c.Status(500)
		return
	}

	c.Data(200, mtype.String(), buf)
}
