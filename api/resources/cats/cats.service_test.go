package cats

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	Router(r)
	return r
}

func resetCats() {
	cats = []cat{
		{ID: "1", Name: "Whiskers", Breed: "Siamese", Age: 3},
		{ID: "2", Name: "Tiger", Breed: "Bengal", Age: 7},
		{ID: "3", Name: "Mittens", Breed: "Tabby", Age: 5},
	}
}

func TestCRUD(t *testing.T) {
	resetCats()
	router := setupRouter()

	// create
	newCat := cat{Name: "NewCat", Breed: "Mixed", Age: 1}
	body, _ := json.Marshal(newCat)
	req, _ := http.NewRequest(http.MethodPost, "/cats", bytes.NewBuffer(body))
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	if resp.Code != http.StatusCreated {
		t.Fatalf("expected status %d, got %d", http.StatusCreated, resp.Code)
	}
	var created cat
	if err := json.Unmarshal(resp.Body.Bytes(), &created); err != nil {
		t.Fatalf("failed to parse response: %v", err)
	}

	// read
	req, _ = http.NewRequest(http.MethodGet, "/cats/"+created.ID, nil)
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	if resp.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, resp.Code)
	}

	// update
	updated := cat{Name: "Updated", Breed: "Mixed", Age: 2}
	body, _ = json.Marshal(updated)
	req, _ = http.NewRequest(http.MethodPut, "/cats/"+created.ID, bytes.NewBuffer(body))
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	if resp.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, resp.Code)
	}

	// delete
	req, _ = http.NewRequest(http.MethodDelete, "/cats/"+created.ID, nil)
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	if resp.Code != http.StatusNoContent {
		t.Fatalf("expected status %d, got %d", http.StatusNoContent, resp.Code)
	}
}
