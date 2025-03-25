package main

import (
	"contact-api/handlers"
	"contact-api/logger"
	"contact-api/middlewares"
	"contact-api/models"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	logger.InitLogger()

	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	models.CreateTable(db)

	r := chi.NewRouter()

	r.Use(middlewares.RequestID)

	r.Get("/contacts", handlers.GetContacts(db))
	r.Post("/contact", handlers.CreateContact(db))

	fmt.Println("Server running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
