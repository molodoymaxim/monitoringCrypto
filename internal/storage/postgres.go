package storage

import (
	"database/sql"
	"fmt"
	"log"
	"monitoringCrypto/internal/config"
)

type Storage struct {
	db *sql.DB
}

func ConnectToStorage(config *config.Config, isLocal bool) *Storage {
	var connStr string
	if isLocal {
		fmt.Println("Try local")
		connStr = fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
			config.PorstgresUserName, config.PorstgresPassword, config.PorstgresDatabase)
	} else {
		connStr = fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=require",
			config.PorstgresUserName, config.PorstgresPassword, config.PorstgresHost, config.PorstgresPort, config.PorstgresDatabase)
	}
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error during connection verification: %v", err)
	}

	log.Println("Connection to the database was successful")

	return &Storage{db: db}
}

func (s *Storage) Close() {
	s.db.Close()
}

func (s *Storage) CreateTables() {
	err := s.CreateCryptoPriceStorage()
	if err != nil {
		log.Fatal("Failed to create tender storage")
	}
}
