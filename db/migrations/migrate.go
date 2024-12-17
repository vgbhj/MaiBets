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

	// Миграция для таблицы event
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

	// Функция триггера для обновления статуса события
	_, err = dbConn.Exec(`
	CREATE OR REPLACE FUNCTION update_event_status()
	RETURNS TRIGGER AS $update_event_status$
	BEGIN
		IF NEW.date < NOW() THEN
			UPDATE event SET status = 'end' WHERE id = NEW.id;
		END IF;
		RETURN NEW;
	END;
	$update_event_status$ LANGUAGE plpgsql;
	`)
	if err != nil {
		log.Fatal("Creating trigger function failed:", err)
	}

	// Создание триггера для обновления статуса события
	// _, err = dbConn.Exec(`
	// CREATE TRIGGER update_event_status
	// AFTER UPDATE ON event
	// FOR EACH ROW
	// EXECUTE FUNCTION update_event_status();
	// `)
	// if err != nil {
	// 	log.Fatal("Creating trigger failed:", err)
	// }

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

	log.Println("Migrations completed successfully.")
}
