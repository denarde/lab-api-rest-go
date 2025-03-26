package main

import (
	"contact-api/handlers"
	"contact-api/logger"
	"contact-api/models"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func setupTestDB() *sql.DB {
	logger.InitLogger()

	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}
	models.CreateTable(db)
	return db
}

func setupRouter(db *sql.DB) *chi.Mux {
	r := chi.NewRouter()
	r.Get("/contacts", handlers.GetContacts(db))
	r.Post("/contact", handlers.CreateContact(db))
	r.Put("/contact/{id}", handlers.UpdateContact(db))
	r.Delete("/contact/{id}", handlers.DeleteContact(db))
	return r
}

func TestCreateContact(t *testing.T) {
	db := setupTestDB()
	defer db.Close()
	router := setupRouter(db)

	reqBody := `{"name": "John Doe", "email": "john.doe@example.com"}`
	req := httptest.NewRequest("POST", "/contact", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusCreated, rec.Code)
}

func TestCreateContact_ErrorResponse(t *testing.T) {
	db := setupTestDB()
	defer db.Close()
	router := setupRouter(db)

	reqBody := `{"name": "josh"}`
	req := httptest.NewRequest("POST", "/contact", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)

	var errorResponse struct {
		Error struct {
			Message string `json:"message"`
			Code    int    `json:"code"`
		} `json:"error"`
	}

	err := json.NewDecoder(rec.Body).Decode(&errorResponse)
	assert.NoError(t, err)

	assert.Equal(t, "Invalid contact data: Key: 'Contact.Email' Error:Field validation for 'Email' failed on the 'required' tag", errorResponse.Error.Message)
	assert.Equal(t, 400, errorResponse.Error.Code)
}

func TestGetContacts(t *testing.T) {
	db := setupTestDB()
	defer db.Close()
	router := setupRouter(db)

	models.InsertContact(db, models.Contact{Name: "Jane Doe", Email: "jane@example.com"})

	req := httptest.NewRequest("GET", "/contacts", nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)

	var contacts []models.Contact
	err := json.Unmarshal(rec.Body.Bytes(), &contacts)
	assert.NoError(t, err)
	assert.Len(t, contacts, 1)
}

func TestUpdateContact(t *testing.T) {
	db := setupTestDB()
	defer db.Close()
	router := setupRouter(db)

	id, _ := models.InsertContact(db, models.Contact{Name: "Old Name", Email: "old@example.com"})

	reqBody := `{"name": "New Name", "email": "new@example.com"}`
	req := httptest.NewRequest("PUT", "/contact/"+strconv.FormatInt(id, 10), strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestDeleteContact(t *testing.T) {
	db := setupTestDB()
	defer db.Close()
	router := setupRouter(db)

	id, _ := models.InsertContact(db, models.Contact{Name: "To Delete", Email: "delete@example.com"})

	req := httptest.NewRequest("DELETE", "/contact/"+strconv.FormatInt(id, 10), nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusNoContent, rec.Code)
}
