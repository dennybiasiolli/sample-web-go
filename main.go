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

func helloAdvancedHandler(c *gin.Context) {
	name := c.Param("name")
	age := c.Param("age")
	ageInt, err := strconv.Atoi(age)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid input")
		return
	}
	cool := c.Param("cool")
	coolBool, err := strconv.ParseBool(cool)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid input")
		return
	}
	if coolBool {
		c.String(http.StatusOK, "You're a cool %d year old, %s!", ageInt, name)
	} else {
		c.String(http.StatusOK, "%s, we need to talk about your coolness.", name)
	}
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
	router.GET("/hello_advanced/:name/:age/:cool", helloAdvancedHandler)

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
