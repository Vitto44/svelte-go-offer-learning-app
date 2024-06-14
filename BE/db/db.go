package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	var err error
	connStr := "user=postgres dbname=certainwager sslmode=disable password=0000 host=172.18.96.1"
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to database!")

	createTables()
}

func createTables() {
	tableCreationQueries := []string{
		`CREATE TABLE IF NOT EXISTS offers (
			id SERIAL PRIMARY KEY,
			title TEXT NOT NULL,
			description TEXT NOT NULL
		)`,
		`CREATE TABLE IF NOT EXISTS reviews (
			id SERIAL PRIMARY KEY,
			title TEXT NOT NULL,
			content TEXT NOT NULL,
			rating INTEGER NOT NULL
		)`,
		`CREATE TABLE IF NOT EXISTS blogs (
			id SERIAL PRIMARY KEY,
			title TEXT NOT NULL,
			content TEXT NOT NULL
		)`,
	}

	for _, query := range tableCreationQueries {
		_, err := DB.Exec(query)
		if err != nil {
			log.Fatalf("Error creating table: %v", err)
		}
	}

	fmt.Println("Tables created or already exist.")
}
