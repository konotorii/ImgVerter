package main

import (
	"context"
	"fmt"
	"github.com/gabriel-vasile/mimetype"
	"github.com/redis/go-redis/v9"
	_ "io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type SavedImage struct {
	data     []byte `redis:"data"`
	mimeType string `redis:"mime_type"`
}

const userKey = "user"

var secret = []byte(os.Getenv("SECRET"))

var port, port_exists = strconv.Atoi(os.Getenv("PORT"))

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {

	fmt.Print(port)
	r := engine()
	r.Use(gin.Logger())
	if err := engine().Run(":8080"); err != nil {
		log.Fatal("unable to start:", err)
	}
}

func engine() *gin.Engine {
	ctx := context.Background()
	r := gin.New()
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADD"),
		Password: os.Getenv("REDIS_PAS"),
		DB:       1,
	})

	status, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalln("Redis connection was refused")
	}
	fmt.Println(status)

	r.Use(sessions.Sessions("session", cookie.NewStore(secret)))

	r.GET("/img", func(c *gin.Context) { getServeImage(c, rdb, ctx) })
	r.POST("/login", login)
	r.GET("/logout", logout)

	admin := r.Group("/admin")
	admin.Use(AuthRequired)
	{
		admin.POST("/upload", postImage)
	}

	return r
}

func getServeImage(c *gin.Context, rdb *redis.Client, ctx context.Context) {
	filePath := c.Query("path")

	replacer := strings.NewReplacer("'", "")

	filePath = replacer.Replace(filePath)

	filePath = "./public/" + filePath

	result, err := rdb.Get(ctx, filePath).Result()
	if err != nil {
		fmt.Println(err)
		fmt.Println("Couldn't find image.")
		buf, err := ioutil.ReadFile(filePath)

		if err != nil {
			fmt.Println("Reading file error", err)

			c.Status(500)
		}

		mtype, err := mimetype.DetectFile(filePath)

		if err != nil {
			fmt.Println("Getting mimetype error", err)

			c.Status(500)
		}

		rdb.Set(ctx, filePath, buf, 30*time.Minute)

		c.Data(200, mtype.String(), buf)
	} else {
		c.Data(200, "", []byte(result))
	}

}

func postImage(c *gin.Context) {
	// Single file
	file, _ := c.FormFile("file")
	log.Println(file.Filename)

	// Upload the file to specific dst.
	c.SaveUploadedFile(file, "./public")

	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
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
