package handlers

import (
	"contact-api/logger"
	"contact-api/middlewares"
	"contact-api/models"
	"database/sql"
	"encoding/json"
	"net/http"
)

func GetContacts(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		requestID := middlewares.GetRequestID(r.Context())
		log := logger.WithRequestID(requestID)

		contacts, err := models.GetContacts(db)
		if err != nil {
			log.Error("Error fetching contacts", err)
			http.Error(w, "Error fetching contacts", http.StatusInternalServerError)
			return
		}

		log.WithField("count", len(contacts)).Info("Retrieved contacts")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(contacts)
	}
}

func CreateContact(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		requestID := middlewares.GetRequestID(r.Context())
		log := logger.WithRequestID(requestID)

		var contact models.Contact
		if err := json.NewDecoder(r.Body).Decode(&contact); err != nil {
			log.Warn("Invalid request body", err)
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		_, err := models.InsertContact(db, contact)
		if err != nil {
			log.Error("Failed to create contact", err)
			http.Error(w, "Failed to create contact", http.StatusInternalServerError)
			return
		}

		log.WithFields(map[string]interface{}{
			"name":  contact.Name,
			"email": contact.Email,
		}).Info("Contact created")

		w.WriteHeader(http.StatusCreated)
	}
}
