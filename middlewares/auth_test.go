package middlewares

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"example.com/rest-api/utils"
	"github.com/gin-gonic/gin"
)

func TestAuthenticate(t *testing.T) {
	r := gin.New()
	r.Use(Authenticate)
	r.GET("/p", func(c *gin.Context) { c.Status(http.StatusOK) })

	// missing token
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/p", nil)
	r.ServeHTTP(w, req)
	if w.Code != http.StatusUnauthorized {
		t.Errorf("expected 401, got %d", w.Code)
	}

	// valid token without Bearer prefix
	token, _ := utils.GenerateToken("a", 1)
	req = httptest.NewRequest(http.MethodGet, "/p", nil)
	req.Header.Set("Authorization", token)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}

	// valid token with Bearer prefix
	req = httptest.NewRequest(http.MethodGet, "/p", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("expected 200 with Bearer, got %d", w.Code)
	}
}