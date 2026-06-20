package handler_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestHealthEndpoint(t *testing.T) {

	gin.SetMode(gin.TestMode)

	router := gin.New()

	router.GET(
		"/health",
		func(c *gin.Context) {
			c.JSON(
				http.StatusOK,
				gin.H{
					"status": "healthy",
				},
			)
		},
	)

	req := httptest.NewRequest(
		http.MethodGet,
		"/health",
		nil,
	)

	resp := httptest.NewRecorder()

	router.ServeHTTP(
		resp,
		req,
	)

	if resp.Code != http.StatusOK {
		t.Fatalf(
			"expected %d got %d",
			http.StatusOK,
			resp.Code,
		)
	}

	var body map[string]interface{}

	err := json.Unmarshal(
		resp.Body.Bytes(),
		&body,
	)

	if err != nil {
		t.Fatal(err)
	}

	if body["status"] != "healthy" {
		t.Fatal("unexpected response")
	}
}
