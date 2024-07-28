package handlers

import (
	"certainwager-be/db"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// SetupDB initializes the database for testing
func setupDB() {
	db.InitDB()
	// Insert test data if necessary
}

// CleanupDB cleans up the database after tests
func cleanupDB() {
	// Drop test data if necessary
}

// TestGetOffers tests the GetOffers handler
func TestGetOffers(t *testing.T) {
	setupDB()
	defer cleanupDB()

	req, err := http.NewRequest("GET", "/offers", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetOffers)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var offers []map[string]interface{}
	if err := json.NewDecoder(rr.Body).Decode(&offers); err != nil {
		t.Errorf("handler returned invalid JSON: %v", err)
	}

	// Check if the response contains expected data
	// This will depend on your test data in the database
}

// TestGetReviews tests the GetReviews handler
func TestGetReviews(t *testing.T) {
	setupDB()
	defer cleanupDB()

	req, err := http.NewRequest("GET", "/reviews", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetReviews)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var reviews []map[string]interface{}
	if err := json.NewDecoder(rr.Body).Decode(&reviews); err != nil {
		t.Errorf("handler returned invalid JSON: %v", err)
	}

	// Check if the response contains expected data
	// This will depend on your test data in the database
}

// TestGetBlogs tests the GetBlogs handler
func TestGetBlogs(t *testing.T) {
	setupDB()
	defer cleanupDB()

	req, err := http.NewRequest("GET", "/blogs", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetBlogs)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var blogs []map[string]interface{}
	if err := json.NewDecoder(rr.Body).Decode(&blogs); err != nil {
		t.Errorf("handler returned invalid JSON: %v", err)
	}

	// Check if the response contains expected data
	// This will depend on your test data in the database
}
