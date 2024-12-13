package main

import (
	"log"

	"github.com/vgbhj/MaiBets/config"
	"github.com/vgbhj/MaiBets/db"
)

func init() {
	config.LoadEnvs()
	db.ConnectDB() // from database.go
}

func main() {
	dbConn := db.ConnectDB() // Получаем соединение с базой данных
	defer dbConn.Close()     // Закрываем соединение после завершения работы

	// Миграция для таблицы users
	_, err := dbConn.Exec(`
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		username VARCHAR(255) NOT NULL UNIQUE,
		password VARCHAR(255) NOT NULL,
		created_at TIMESTAMP DEFAULT NOW()
	);
	`)
	if err != nil {
		log.Fatal("Migration for users failed:", err)
	}

	// Миграция для таблицы events
	_, err = dbConn.Exec(`
	CREATE TABLE IF NOT EXISTS event (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		description TEXT,
		date TIMESTAMP NOT NULL,
		status VARCHAR(50)
	);
	`)
	if err != nil {
		log.Fatal("Migration for event failed:", err)
	}

	log.Println("Migrations completed successfully.")
}
