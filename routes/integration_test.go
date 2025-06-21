package routes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"example.com/rest-api/internal/testutils"
	"example.com/rest-api/middlewares"
	"github.com/gin-gonic/gin"
)

func setupRouter(t *testing.T) *gin.Engine {
	cleanup := testutils.SetupTestDB(t)
	t.Cleanup(cleanup)

	r := gin.Default()
	r.Use(middlewares.CORSMiddleware())
	RegisterRoutes(r)
	return r
}

func TestSignupLoginAndEventFlow(t *testing.T) {
	router := setupRouter(t)

	signupBody := `{"email":"demo@example.com","password":"secret"}`
	req := httptest.NewRequest(http.MethodPost, "/signup", bytes.NewBufferString(signupBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	if w.Code != http.StatusCreated {
		t.Fatalf("signup status %d", w.Code)
	}

	req = httptest.NewRequest(http.MethodPost, "/login", bytes.NewBufferString(signupBody))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("login status %d", w.Code)
	}
	var loginResp map[string]string
	if err := json.Unmarshal(w.Body.Bytes(), &loginResp); err != nil {
		t.Fatalf("decode login resp: %v", err)
	}
	token := loginResp["token"]
	if token == "" {
		t.Fatal("token empty")
	}

	eventPayload := fmt.Sprintf(`{"name":"Event","description":"Desc","location":"Loc","dateTime":"%s"}`,
		time.Now().UTC().Format(time.RFC3339))
	req = httptest.NewRequest(http.MethodPost, "/events", bytes.NewBufferString(eventPayload))
	req.Header.Set("Authorization", token)
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	if w.Code != http.StatusCreated {
		t.Fatalf("create event status %d", w.Code)
	}
	var createResp struct {
		Event struct {
			ID int64 `json:"id"`
		}
	}
	json.Unmarshal(w.Body.Bytes(), &createResp)
	id := createResp.Event.ID

	req = httptest.NewRequest(http.MethodGet, "/events", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("list events status %d", w.Code)
	}

	req = httptest.NewRequest(http.MethodGet, fmt.Sprintf("/events/%d", id), nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("get event status %d", w.Code)
	}

	req = httptest.NewRequest(http.MethodPost, fmt.Sprintf("/events/%d/register", id), nil)
	req.Header.Set("Authorization", token)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	if w.Code != http.StatusCreated {
		t.Fatalf("register status %d", w.Code)
	}

	req = httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/events/%d/register", id), nil)
	req.Header.Set("Authorization", token)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("cancel status %d", w.Code)
	}

	updatePayload := fmt.Sprintf(`{"name":"Changed","description":"Desc","location":"Loc","dateTime":"%s"}`,
		time.Now().UTC().Format(time.RFC3339))
	req = httptest.NewRequest(http.MethodPut, fmt.Sprintf("/events/%d", id), bytes.NewBufferString(updatePayload))
	req.Header.Set("Authorization", token)
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("update status %d", w.Code)
	}

	req = httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/events/%d", id), nil)
	req.Header.Set("Authorization", token)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("delete status %d", w.Code)
	}
}

func TestInvalidLoginAndUnauthorized(t *testing.T) {
	router := setupRouter(t)

	// signup user
	body := `{"email":"x@y.com","password":"pass"}`
	req := httptest.NewRequest(http.MethodPost, "/signup", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	if w.Code != http.StatusCreated {
		t.Fatalf("signup status %d", w.Code)
	}

	// wrong password login
	req = httptest.NewRequest(http.MethodPost, "/login", bytes.NewBufferString(`{"email":"x@y.com","password":"bad"}`))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	if w.Code != http.StatusUnauthorized {
		t.Fatalf("expected 401, got %d", w.Code)
	}

	// create event without token
	payload := fmt.Sprintf(`{"name":"E","description":"D","location":"L","dateTime":"%s"}`, time.Now().UTC().Format(time.RFC3339))
	req = httptest.NewRequest(http.MethodPost, "/events", bytes.NewBufferString(payload))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	if w.Code != http.StatusUnauthorized {
		t.Fatalf("expected 401 unauthorized, got %d", w.Code)
	}
}
