package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

// ConnectDB устанавливает соединение с базой данных
func ConnectDB() *sql.DB {
	// Open the connection
	db, err := sql.Open("postgres", os.Getenv("DB_URL"))
	// connStr := "host=0.0.0.0 user=user password=password dbname=mydatabase port=5432 sslmode=disable"
	// db, err := sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}

	// check the connection
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	// return the connection
	return db
}
