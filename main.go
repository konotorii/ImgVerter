package main

import (
	"context"
	"fmt"
	"github.com/TwiN/go-color"
	"github.com/gabriel-vasile/mimetype"
	"github.com/redis/go-redis/v9"
	_ "io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/konotorii/go-consola"
)

const userKey = "user"

var secret = []byte(os.Getenv("SECRET"))

var port, port_exists = strconv.Atoi(os.Getenv("PORT"))

var rdb *redis.Client
var ctx = context.Background()

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Println(color.Ize(color.White, "No .env was found!"))
	}
}

func main() {

	fmt.Print(port)
	r := engine()
	r.Use(gin.Logger())
	if err := engine().Run(":8080"); err != nil {
		consola.Error("unable to start:", err)
	}
}

func engine() *gin.Engine {
	r := gin.New()
	rdb = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADD"),
		Password: os.Getenv("REDIS_PAS"),
		DB:       1,
	})

	status, err := rdb.Ping(ctx).Result()
	if err != nil {
		consola.Error("Redis connection was refused")
	}
	consola.Log(status)

	r.Use(sessions.Sessions("session", cookie.NewStore(secret)))

	// Handle long routes
	r.GET("/:fileId1", getServeFile)
	r.GET("/:fileId1/:fileId2", getServeFile)
	r.GET("/:fileId1/:fileId2/:fileId3", getServeFile)
	r.GET("/:fileId1/:fileId2/:fileId3/:fileId4", getServeFile)
	r.GET("/:fileId1/:fileId2/:fileId3/:fileId4/:fileId5", getServeFile)
	r.GET("/:fileId1/:fileId2/:fileId3/:fileId4/:fileId5/:fileId6", getServeFile)
	r.GET("/:fileId1/:fileId2/:fileId3/:fileId4/:fileId5/:fileId6/:fileId7", getServeFile)
	r.GET("/:fileId1/:fileId2/:fileId3/:fileId4/:fileId5/:fileId6/:fileId7/:fileId8", getServeFile)

	r.GET("/img", getServeImage)
	r.POST("/login", login)
	r.GET("/logout", logout)

	r.POST("/upload", postScreenshot)

	admin := r.Group("/admin")
	admin.Use(AuthRequired)
	{
		admin.POST("/upload", postScreenshot)
	}

	return r
}

func getServeFile(c *gin.Context) {
	fileId1 := c.Param("fileId1")
	fileId2 := c.Param("fileId2")
	fileId3 := c.Param("fileId3")
	fileId4 := c.Param("fileId4")
	fileId5 := c.Param("fileId5")
	fileId6 := c.Param("fileId6")
	fileId7 := c.Param("fileId7")
	fileId8 := c.Param("fileId8")

	filePath := "/public/"

	if len(fileId1) > 0 {
		filePath = filePath + fileId1
	}
	if len(fileId2) > 0 {
		filePath = filePath + "/" + fileId2
	}
	if len(fileId3) > 0 {
		filePath = filePath + "/" + fileId3
	}
	if len(fileId4) > 0 {
		filePath = filePath + "/" + fileId4
	}
	if len(fileId5) > 0 {
		filePath = filePath + "/" + fileId5
	}
	if len(fileId6) > 0 {
		filePath = filePath + "/" + fileId6
	}
	if len(fileId7) > 0 {
		filePath = filePath + "/" + fileId7
	}
	if len(fileId8) > 0 {
		filePath = filePath + "/" + fileId8
	}

	mtype, err := mimetype.DetectFile(filePath)

	if err != nil {
		consola.Error("Getting mimetype error", err)

		c.Status(500)
	}

	buf, err := ioutil.ReadFile(filePath)

	c.Data(200, mtype.String(), buf)
}

func getServeImage(c *gin.Context) {
	filePath := c.Query("path")

	replacer := strings.NewReplacer("'", "")

	filePath = replacer.Replace(filePath)

	filePath = "/public/" + filePath

	result, err := rdb.Get(ctx, filePath).Result()
	if err != nil {
		consola.Error(err)
		consola.Warning("Couldn't find image.")
		buf, err := ioutil.ReadFile(filePath)

		if err != nil {
			consola.Error("Reading file error", err)

			c.Status(500)
		}

		mtype, err := mimetype.DetectFile(filePath)

		if err != nil {
			consola.Error("Getting mimetype error", err)

			c.Status(500)
		}

		rdb.Set(ctx, filePath, buf, 30*time.Minute)

		c.Data(200, mtype.String(), buf)
	} else {
		c.Data(200, "", []byte(result))
	}

}

func postImage(c *gin.Context) {

}

func postVideo(c *gin.Context) {

}

func postFile(c *gin.Context) {
	key := c.Query("key")

	if key == os.Getenv("KEY") {
		// Single file
		file, _ := c.FormFile("file")
		consola.Log(file.Filename)

		// Upload the file to specific dst.
		err := c.SaveUploadedFile(file, "/public/download/"+file.Filename)
		if err != nil {
			consola.Error(err)
			c.Status(500)
		}

		c.JSON(200, gin.H{"url": fmt.Sprintf("https://img.kono.services/download/" + file.Filename)})
	} else {
		c.Status(400)
	}
}

func postScreenshot(c *gin.Context) {
	key := c.Query("key")

	if key == os.Getenv("KEY") {
		// Single file
		file, _ := c.FormFile("file")
		consola.Log("Uploaded file: " + file.Filename)

		// Upload the file to specific dst.
		err := c.SaveUploadedFile(file, "/public/ss/"+file.Filename)
		if err != nil {
			consola.Error(err)
			c.Status(500)
		}

		c.JSON(200, gin.H{"url": fmt.Sprintf("https://img.kono.services/ss/" + file.Filename)})
	} else {
		c.Status(400)
	}
}

func AuthRequired(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(userKey)
	if user == nil {
		// Abort the request with the appropriate error code
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	// Continue down the chain to handler etc
	c.Next()
}

func login(c *gin.Context) {
	session := sessions.Default(c)
	username := c.PostForm("username")
	password := c.PostForm("password")

	// Validate form input
	if strings.Trim(username, " ") == "" || strings.Trim(password, " ") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameters can't be empty"})
		return
	}

	// Check for username and password match, usually from a database
	if username != "hello" || password != "itsme" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
		return
	}

	// Save the username in the session
	session.Set(userKey, username) // In real world usage you'd set this to the users ID
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully authenticated user"})
}

func logout(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(userKey)
	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session token"})
		return
	}
	session.Delete(userKey)
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}
