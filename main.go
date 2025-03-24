package main

import (
	"contact-api/handlers"
	"contact-api/models"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	models.CreateTable(db)

	http.HandleFunc("/contacts", handlers.GetContacts(db))
	http.HandleFunc("/contact", handlers.CreateContact(db))

	fmt.Println("Server running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
