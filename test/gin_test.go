package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestPingRoute(t *testing.T) {
	// Create a new Gin router
	router := gin.Default()

	// Define a GET handler for /ping
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Create a new HTTP request to the /ping endpoint
	req, err := http.NewRequest("GET", "/ping", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new HTTP recorder to capture the response
	rr := httptest.NewRecorder()

	// Send the request to the router
	router.ServeHTTP(rr, req)

	// Check that the status code is 200 OK
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check that the response body contains the expected message
	expected := `{"message":"pong"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
