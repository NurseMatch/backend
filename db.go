package main

import (
	"fmt"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"log"
	"os"
)

const (
	port     = 1433
	database = "nurse_match"
)

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

	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("Connected!")
	return db, err
}

func runMigration(db *gorm.DB) {
	err := db.AutoMigrate(
		&User{},
		&Consultant{},
		&Location{},
		&Assignment{},
		&WorkProfile{})
	if err != nil {
		return
	}
}
