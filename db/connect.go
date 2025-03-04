package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type User struct {
	ID        uint32
	Email     string
	Password  string
	FirstName string
	LastName  string
	Phone     string
}

type Company struct {
	ID           uint32
	ContactEmail string
	Password     string
	Name         string
	Address      string
	Postcode     string
	Country      string
}

func ConnectDB() *sql.DB {
	connStr := os.Getenv("DB_URL")

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully connected to database")

	return db
}
