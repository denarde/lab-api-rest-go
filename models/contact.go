package models

import (
	"database/sql"
	"log"
)

type Contact struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func CreateTable(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS contacts (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		email TEXT
	);`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

func InsertContact(db *sql.DB, c Contact) (int64, error) {
	query := `INSERT INTO contacts (name, email) VALUES (?, ?)`
	result, err := db.Exec(query, c.Name, c.Email)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func GetContacts(db *sql.DB) ([]Contact, error) {
	rows, err := db.Query("SELECT id, name, email FROM contacts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var contacts []Contact
	for rows.Next() {
		var c Contact
		if err := rows.Scan(&c.ID, &c.Name, &c.Email); err != nil {
			return nil, err
		}
		contacts = append(contacts, c)
	}
	return contacts, nil
}

func UpdateContact(db *sql.DB, c Contact) error {
	query := `UPDATE contacts SET name = ?, email = ? WHERE id = ?`
	_, err := db.Exec(query, c.Name, c.Email, c.ID)
	if err != nil {
		return err
	}
	return nil
}

func DeleteContact(db *sql.DB, id int) error {
	query := `DELETE FROM contacts WHERE id = ?`
	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
