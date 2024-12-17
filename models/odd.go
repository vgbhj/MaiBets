package models

import (
	"fmt"
	"time"

	"github.com/vgbhj/MaiBets/db"
)

type Odd struct {
	ID        int       `json:"id" gorm:"primary_key"`      // юзаем только для GET
	OddValue  float64   `json:"odd_value" gorm:"not null"`  // Значение ставки
	EventID   int       `json:"event_id" gorm:"not null"`   // Ссылка на событие
	UpdatedAt time.Time `json:"updated_at" gorm:"not null"` // Время последнего обновления
}

// AddOdd добавляет новое значение ставки в базу данных
func AddOdd(data map[string]interface{}) error {
	dbConn := db.ConnectDB() // Получаем соединение с базой данных
	defer dbConn.Close()

	odd := Odd{
		OddValue:  data["odd_value"].(float64),
		EventID:   data["event_id"].(int),
		UpdatedAt: time.Now(), // Устанавливаем текущее время
	}

	_, err := dbConn.Exec("INSERT INTO odd (odd_value, event_id, updated_at) VALUES ($1, $2, $3)",
		odd.OddValue, odd.EventID, odd.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

// GetOdd получает значение ставки по его ID
func GetOdd(id int) (*Odd, error) {
	dbConn := db.ConnectDB() // Получаем соединение с базой данных
	defer dbConn.Close()

	var odd Odd

	err := dbConn.QueryRow("SELECT id, odd_value, event_id, updated_at FROM odd WHERE id = $1", id).Scan(
		&odd.ID, &odd.OddValue, &odd.EventID, &odd.UpdatedAt)

	if err != nil {
		// Выводим ошибку в лог
		fmt.Println("Error retrieving odd:", err)

		// Возвращаем ошибку
		return nil, err
	}

	return &odd, nil
}

func GetOddIDByEventId(id int) (int, error) {
	dbConn := db.ConnectDB() // Получаем соединение с базой данных
	defer dbConn.Close()

	var oddID int

	err := dbConn.QueryRow("SELECT id FROM odd WHERE event_id = $1", id).Scan(&oddID)
	if err != nil {
		// Выводим ошибку в лог
		fmt.Println("Error retrieving event ID by name:", err)

		// Возвращаем ошибку
		return 0, err
	}

	return oddID, nil
}
