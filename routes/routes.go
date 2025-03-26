package routes

import (
	"contact-api/handlers"
	"contact-api/middlewares"
	"database/sql"

	"github.com/go-chi/chi/v5"
)

func SetupRoutes(db *sql.DB) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middlewares.RequestID)

	r.Group(func(r chi.Router) {
		r.Use(middlewares.AuthMiddleware)
		r.Get("/contacts", handlers.GetContacts(db))
		r.Post("/contact", handlers.CreateContact(db))
		r.Put("/contact/{id}", handlers.UpdateContact(db))
		r.Delete("/contact/{id}", handlers.DeleteContact(db))
	})

	r.Post("/login", handlers.Login)

	return r
}
