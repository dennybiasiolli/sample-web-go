package main

import (
	"net/http"
	"strconv"
	"time"

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

func delayHandler(c *gin.Context) {
	seconds := c.Param("seconds")
	secondsInt, err := strconv.Atoi(seconds)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid input")
		return
	}
	time.Sleep(time.Duration(secondsInt) * time.Second)
	c.String(http.StatusOK, "Waited for %d seconds", secondsInt)
}

func setupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/hello", helloHandler)

	api := router.Group("/api")
	api.GET("/hello", helloApiHandler)

	router.Static("/static", "./static")
	router.GET("/async/delay/:seconds", delayHandler)

	return router
}

func main() {
	router := setupRouter()
	router.Run()
}
