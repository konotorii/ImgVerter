package main

//import "net/http"
//import "io"
//import "log"

//import "github.com/sunshineplan/imgconv"

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Run("localhost:8000")
}

func convertImage(c *gin.Context) {
	url := c.Param("url")
}
