package main

import (
	"fmt"
	"imgverter/connectors"
	"imgverter/routes/asset"
	"imgverter/routes/upload"
	"imgverter/util"
	_ "io"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/konotorii/go-consola"
)

func init() {
	if err := godotenv.Load(); err != nil {
		consola.Error("No .env was found!")
	}

	util.ConfigInit()

	_, _ = connectors.RedisDatabaseInit()
}

func main() {
	r := engine()
	r.Use(gin.Logger())
	if err := engine().Run(fmt.Sprintf(":%d", util.Config.Port)); err != nil {
		consola.Error("unable to start:", err)
	}
}

func engine() *gin.Engine {
	r := gin.New()

	r.Use(sessions.Sessions("session", cookie.NewStore([]byte(util.Config.CookieSecret))))

	r.GET("/i/:id", asset.FetchImage)

	r.POST("/upload/img", upload.PostImage)
	r.POST("/upload/file", upload.PostFile)

	return r
}
