## Project initialization

```bash
GIT_USERNAME=dennybiasiolli
GIT_REPO=sample-web-go
mkdir $GIT_REPO
cd $GIT_REPO
go mod init github.com/$GIT_USERNAME/$GIT_REPO
```

## Gin Web Framework

```bash
go get -u github.com/gin-gonic/gin
```

```go
// file: main.go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func helloHandler(c *gin.Context) {
	c.String(http.StatusOK, "Hello World!")
}

func main() {
	router := gin.Default()

	router.GET("/hello", helloHandler)

	router.Run()
}
```

```bash
go run .
```

## Add an API endpoint

```go
func helloApiHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World!",
	})
}

func main() {
    // ...
    api := router.Group("/api")
	api.GET("/hello", helloApiHandler)
}
```

## Serve static files

```go
router.Static("/static", "./static")
```

## Async requests

```go
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

// ...
router.GET("/async/delay/:seconds", delayHandler)
```
