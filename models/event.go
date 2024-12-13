package models

import (
	"fmt"
	"time"

	"github.com/vgbhj/MaiBets/db"
)

type Event struct {
	ID     int       `json:"id" gorm:"primary_key"`
	Name   string    `json:"name" gorm:"not null"`
	Desc   string    `json:"description"`
	Date   time.Time `json:"date" gorm:"not null"`
	Status string    `json:"status"`
}

// AddEvent добавляет новое событие в базу данных
func AddEvent(data map[string]interface{}) error {
	dbConn := db.ConnectDB() // Получаем соединение с базой данных
	defer dbConn.Close()

	event := Event{
		ID:     data["id"].(int),
		Name:   data["name"].(string),
		Desc:   data["desc"].(string),
		Date:   data["date"].(time.Time),
		Status: data["status"].(string),
	}

	_, err := dbConn.Exec("INSERT INTO event (id, name, description, date, status) VALUES ($1, $2, $3, $4, $5)",
		event.ID, event.Name, event.Desc, event.Date, event.Status)
	if err != nil {
		return err
	}

	return nil
}

// GetEvent получает событие по его ID
func GetEvent(id int) (*Event, error) {
	dbConn := db.ConnectDB() // Получаем соединение с базой данных
	defer dbConn.Close()

	var event Event

	err := dbConn.QueryRow("SELECT id, name, description, date, status FROM event WHERE id = $1", id).Scan(
		&event.ID, &event.Name, &event.Desc, &event.Date, &event.Status)

	if err != nil {
		// Выводим ошибку в лог
		fmt.Println("Error retrieving event:", err)

		// Возвращаем ошибку
		return nil, err
	}

	return &event, nil
}
