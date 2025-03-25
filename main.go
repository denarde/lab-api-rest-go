package main

import (
	"contact-api/logger"
	"contact-api/models"
	"contact-api/routes"
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

	r = routes.SetupRoutes(db)

	fmt.Println("Server running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
