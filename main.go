package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	router := gin.New()
	router.GET("/mark", func(c *gin.Context) {
		c.String(http.StatusOK, string(blackfriday.MarkdownBasic([]byte("**hi!**"))))
	})
	router.Run(":" + port)
}
