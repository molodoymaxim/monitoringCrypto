package storage

import (
	"fmt"
	"log"
)

func (s *Storage) CreateCryptoPriceStorage() error {
	createTableSql := `
  	CREATE TABLE IF NOT EXISTS tenders (
		id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
		name VARCHAR(100) NOT NULL,
  	    price REAL NOT NULL,
		currency VARCHAR(100) NOT NULL,
		timeCheck TIMESTAMP NOT NULL
	 );`

	_, err := s.db.Exec(createTableSql)
	if err != nil {
		msgErr := fmt.Errorf("Error creating table:", err)
		log.Println(msgErr)
		return msgErr
	}

	return nil
}

func (s *Storage) InsertCryptoPriceStorage(price float32, currency string) error {
	const op = "storage.CreateTender"

	insertQuery := `
		INSERT INTO tenders (name, description, serviceType, status, organizationId, creatorUsername,version, createdAt)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id;
	`

	stmt, err := s.db.Prepare(insertQuery)
	defer stmt.Close()
	if err != nil {
		return "", fmt.Errorf("%s. Error preparing statement: %v", op, err)
	}

	var ID string
	err = stmt.QueryRow(tender.Name, tender.Description, tender.ServiceType, tender.Status, tender.OrganizationId, tender.CreatorUsername, tender.Version, tender.CreatedAt).Scan(&ID)
	if err != nil {
		return "", fmt.Errorf("%s. Error executing query: %v", op, err)
	}

	return ID, nil
}
