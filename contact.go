package handlers

import (
	"contact-api/logger"
	"contact-api/middlewares"
	"contact-api/models"
	"contact-api/utils"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func GetContacts(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		requestID := middlewares.GetRequestID(r.Context())
		log := logger.WithRequestID(requestID)

		page, _ := strconv.Atoi(r.URL.Query().Get("page"))
		limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
		name := r.URL.Query().Get("name")

		if page < 1 {
			page = 1
		}
		if limit < 1 {
			limit = 10
		}

		contacts, err := models.GetContacts(db, page, limit, name)
		if err != nil {
			log.Error("Error fetching contacts", err)
			utils.SendError(w, http.StatusInternalServerError, "Error fetching contacts")
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
			utils.SendError(w, http.StatusBadRequest, "Invalid request body")
			return
		}

		if err := contact.Validate(); err != nil {
			log.Warn("Validation failed", err)
			print("PRINT: Invalid contact data: " + err.Error())
			utils.SendError(w, http.StatusBadRequest, "Invalid contact data: "+err.Error())
			return
		}

		_, err := models.InsertContact(db, contact)
		if err != nil {
			log.Error("Failed to create contact", err)
			utils.SendError(w, http.StatusInternalServerError, "Failed to create contact")
			return
		}

		log.WithFields(map[string]interface{}{
			"name":  contact.Name,
			"email": contact.Email,
		}).Info("Contact created")

		w.WriteHeader(http.StatusCreated)
	}
}

func UpdateContact(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		requestID := middlewares.GetRequestID(r.Context())
		log := logger.WithRequestID(requestID)

		id := chi.URLParam(r, "id")
		contactID := 0
		_, err := fmt.Sscanf(id, "%d", &contactID)
		if err != nil || contactID <= 0 {
			log.WithField("id", id).Warn("Invalid contact ID")
			utils.SendError(w, http.StatusBadRequest, "Invalid contact ID")
			return
		}

		var contact models.Contact
		if err := json.NewDecoder(r.Body).Decode(&contact); err != nil {
			log.Warn("Invalid request body", err)
			utils.SendError(w, http.StatusBadRequest, "Invalid request body")
			return
		}

		contact.ID = contactID
		err = models.UpdateContact(db, contact)
		if err != nil {
			log.Error("Failed to update contact", err)
			utils.SendError(w, http.StatusInternalServerError, "Failed to update contact")
			return
		}

		log.WithFields(map[string]interface{}{
			"id":    contact.ID,
			"name":  contact.Name,
			"email": contact.Email,
		}).Info("Contact updated")

		w.WriteHeader(http.StatusOK)
	}
}

func DeleteContact(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		requestID := middlewares.GetRequestID(r.Context())
		log := logger.WithRequestID(requestID)

		id := chi.URLParam(r, "id")
		contactID := 0
		_, err := fmt.Sscanf(id, "%d", &contactID)
		if err != nil || contactID <= 0 {
			log.WithField("id", id).Warn("Invalid contact ID")
			utils.SendError(w, http.StatusBadRequest, "Invalid contact ID")
			return
		}

		err = models.DeleteContact(db, contactID)
		if err != nil {
			log.Error("Failed to delete contact", err)
			utils.SendError(w, http.StatusInternalServerError, "Failed to delete contact")
			return
		}

		log.WithField("id", contactID).Info("Contact deleted")
		w.WriteHeader(http.StatusNoContent)
	}
}
