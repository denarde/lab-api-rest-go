package models

import (
	"database/sql"
	"log"

	"github.com/go-playground/validator/v10"
)

type Contact struct {
	ID    int    `json:"id"`
	Name  string `json:"name" validate:"required,min=3"`
	Email string `json:"email" validate:"required,email"`
}

var validate = validator.New()

func (c *Contact) Validate() error {
	return validate.Struct(c)
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

func GetContacts(db *sql.DB, page, limit int, name string) ([]Contact, error) {
	offset := (page - 1) * limit

	query := "SELECT id, name, email FROM contacts"
	var args []interface{}

	if name != "" {
		query += " WHERE name LIKE ?"
		args = append(args, "%"+name+"%")
	}

	query += " LIMIT ? OFFSET ?"
	args = append(args, limit, offset)
	rows, err := db.Query(query, args...)
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
