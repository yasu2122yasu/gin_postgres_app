package controller

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetRouter(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := GetRouter()

	testCases := []struct {
		method         string
		endpoint       string
		expectedStatus int
	}{
		{http.MethodGet, "/", http.StatusOK},
		{http.MethodGet, "/show/1", http.StatusOK},
		{http.MethodGet, "/create", http.StatusOK},
		{http.MethodPost, "/create", http.StatusOK},
		{http.MethodGet, "/edit/1", http.StatusOK},
		{http.MethodPost, "/edit", http.StatusOK},
		{http.MethodGet, "/delete/1", http.StatusOK},
		{http.MethodPost, "/delete", http.StatusOK},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("%s %s", testCase.method, testCase.endpoint), func(t *testing.T) {
			req, _ := http.NewRequest(testCase.method, testCase.endpoint, nil)
			resp := httptest.NewRecorder()

			router.ServeHTTP(resp, req)

			if resp.Code != testCase.expectedStatus {
				t.Errorf("Expected status code %v, got %v", testCase.expectedStatus, resp.Code)
			}
		})
	}
}
