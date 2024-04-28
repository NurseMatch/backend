package main

import (
	"backend/data"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"log"
	"os"
)

const (
	port     = 1433
	database = "nurse_match"
)

func connectToLocalDb() (*gorm.DB, error) {
	dsn := "host=localhost user=user1234 password=user1234 dbname=localdb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func connectToDb() (*gorm.DB, error) {
	// Get environment variables
	server := os.Getenv("DB_SERVER")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")

	// Build connection string
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		server, user, password, port, database)

	// Create connection pool
	db, err := gorm.Open(sqlserver.Open(connString), &gorm.Config{})
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}

	fmt.Printf("Connected!")
	return db, err
}

func runMigration(db *gorm.DB) {
	err := db.AutoMigrate(
		&data.User{},
		&data.Consultant{},
		&data.Location{},
		&data.Assignment{},
		&data.WorkProfile{},
		&data.Workplace{})
	if err != nil {
		return
	}
}
