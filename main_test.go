package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHelloRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/hello", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, w.Code, http.StatusOK)
	assert.Equal(t, w.Body.String(), "Hello World!")
}

func TestApiHelloRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/hello", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, w.Code, http.StatusOK)
	result := gin.H{}
	if err := json.NewDecoder(w.Body).Decode(&result); err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, result, gin.H{
		"message": "Hello World!",
	})
}

func TestStaticRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/static/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, w.Code, http.StatusOK)
	body := w.Body.String()
	assert.Contains(t, body, "<title>Static HTML page</title>")
	assert.Contains(t, body, "<h1>Static HTML page</h1>")
}

func TestDelayRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/async/delay/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, w.Code, http.StatusOK)
	assert.Equal(t, w.Body.String(), "Waited for 1 seconds")
}
