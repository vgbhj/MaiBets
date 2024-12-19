package models

import (
	"fmt"
	"time"

	"github.com/vgbhj/MaiBets/db"
)

type Event struct {
	// юзаем только в GET запросах
	ID     int       `json:"id" gorm:"primary_key"`
	Name   string    `json:"name" gorm:"not null" example:"Champions League Final"`
	Desc   string    `json:"description" example:"Final match of the 2024 Champions League"`
	Date   time.Time `json:"date" gorm:"not null" example:"2024-06-01T20:00:00Z"`
	Status string    `json:"status" example:"live"`
}

// AddEvent добавляет новое событие в базу данных
func AddEvent(data map[string]interface{}) error {
	dbConn := db.ConnectDB() // Получаем соединение с базой данных
	defer dbConn.Close()

	event := Event{
		Name:   data["name"].(string),
		Desc:   data["desc"].(string),
		Date:   data["date"].(time.Time),
		Status: data["status"].(string),
	}

	_, err := dbConn.Exec("INSERT INTO event (name, description, date, status) VALUES ($1, $2, $3, $4)",
		event.Name, event.Desc, event.Date, event.Status)
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

// GetEventIDByName получает ID события по его имени
func GetEventIDByName(name string) (int, error) {
	dbConn := db.ConnectDB() // Получаем соединение с базой данных
	defer dbConn.Close()

	var eventID int

	err := dbConn.QueryRow("SELECT id FROM event WHERE name = $1", name).Scan(&eventID)
	if err != nil {
		// Выводим ошибку в лог
		fmt.Println("Error retrieving event ID by name:", err)

		// Возвращаем ошибку
		return 0, err
	}

	return eventID, nil
}

// UpdateEventStatus обновляет статус событий, у которых дата меньше или равна текущему времени
func UpdateEventStatus() error {
	dbConn := db.ConnectDB() // Получаем соединение с базой данных
	defer dbConn.Close()
	// Обновляем статус событий
	_, err := dbConn.Exec(`
        UPDATE event
        SET status = 'finished'
        WHERE date <= $1 AND status != 'finished'
    `, time.Now())

	return err
}

// GetEventStatusByID получает статус события по его ID
func GetEventStatusByID(id int) (string, error) {
	dbConn := db.ConnectDB() // Получаем соединение с базой данных
	defer dbConn.Close()

	var status string

	err := dbConn.QueryRow("SELECT status FROM event WHERE id = $1", id).Scan(&status)
	if err != nil {
		// Выводим ошибку в лог
		fmt.Println("Error retrieving event status by ID:", err)

		// Возвращаем ошибку
		return "", err
	}

	return status, nil
}
