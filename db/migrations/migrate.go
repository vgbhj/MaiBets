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

	// Миграция для таблицы user
	_, err := dbConn.Exec(`
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		first_name VARCHAR(255),
		last_name VARCHAR(255),
		email VARCHAR(255) NOT NULL UNIQUE,
		phone VARCHAR(50),
		balance DECIMAL DEFAULT 0,
		access INTEGER REFERENCES access(id)
	);
	`)
	if err != nil {
		log.Fatal("Migration for users failed:", err)
	}

	// Миграция для таблицы event
	_, err = dbConn.Exec(`
	CREATE TABLE IF NOT EXISTS event (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		description TEXT,
		date TIMESTAMP NOT NULL,
		status VARCHAR(50) -- e.g., pending, live, finished
	);
	`)
	if err != nil {
		log.Fatal("Migration for event failed:", err)
	}

	// Миграция для таблицы access
	_, err = dbConn.Exec(`
	CREATE TABLE IF NOT EXISTS access (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		description TEXT
	);
	`)
	if err != nil {
		log.Fatal("Migration for access failed:", err)
	}

	// Добавление двух записей: casual и admin
	_, err = dbConn.Exec(`
	INSERT INTO access (name, description) VALUES
	('casual', 'Basic access level'),
	('admin', 'Administrative access level')
	`)
	if err != nil {
		log.Fatal("Inserting default access levels failed:", err)
	}

	// Миграция для таблицы bet_type
	_, err = dbConn.Exec(`
	CREATE TABLE IF NOT EXISTS bet_type (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255), -- e.g., win, draw, lose, total over/under
		description TEXT -- Description of the bet type
	);
	`)
	if err != nil {
		log.Fatal("Migration for bet_type failed:", err)
	}

	// Добавление пяти записей в bet_type
	_, err = dbConn.Exec(`
	INSERT INTO bet_type (name, description) VALUES
	('win', 'Bet on the team to win'),
	('draw', 'Bet on the match to end in a draw'),
	('lose', 'Bet on the team to lose'),
	('total over', 'Bet on the total score to be over a certain number'),
	('total under', 'Bet on the total score to be under a certain number')
	`)
	if err != nil {
		log.Fatal("Inserting default bet types failed:", err)
	}
	// Миграция для таблицы odd
	_, err = dbConn.Exec(`
	CREATE TABLE IF NOT EXISTS odd (
		id SERIAL PRIMARY KEY,
		odd_value DECIMAL, -- Value of the odd for the event
		event_id INTEGER REFERENCES event(id),
		updated_at TIMESTAMP -- When the odd were last updated
	);
	`)
	if err != nil {
		log.Fatal("Migration for odd failed:", err)
	}

	// Миграция для таблицы bet
	_, err = dbConn.Exec(`
	CREATE TABLE IF NOT EXISTS bet (
		id SERIAL PRIMARY KEY,
		client_id INTEGER REFERENCES users(id), -- Many bet to one client
		event_id INTEGER REFERENCES event(id), -- Many bet to one event
		bet_type_id INTEGER REFERENCES bet_type(id), -- Many bet to one bet type
		odd_id INTEGER REFERENCES odd(id), 
		bet_amount DECIMAL,
		status VARCHAR(50), -- e.g., pending, won, lost
		bet_date TIMESTAMP
	);
	`)
	if err != nil {
		log.Fatal("Migration for bet failed:", err)
	}

	// Миграция для таблицы payment
	_, err = dbConn.Exec(`
	CREATE TABLE IF NOT EXISTS payment (
		id SERIAL PRIMARY KEY,
		client_id INTEGER REFERENCES users(id),
		payment_date TIMESTAMP,
		amount DECIMAL,
		payment_type VARCHAR(50) -- e.g., deposit, withdrawal
	);
	`)
	if err != nil {
		log.Fatal("Migration for payment failed:", err)
	}

	log.Println("Migrations completed successfully.")
}
