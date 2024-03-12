package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func helloHandler(c *gin.Context) {
	c.String(http.StatusOK, "Hello World!")
}

func helloApiHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World!",
	})
}

func setupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/hello", helloHandler)

	api := router.Group("/api")
	api.GET("/hello", helloApiHandler)

	router.Static("/static", "./static")

	return router
}

func main() {
	router := setupRouter()
	router.Run()
}
