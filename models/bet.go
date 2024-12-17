package models

import (
	"time"

	"github.com/vgbhj/MaiBets/db"
)

// Bet представляет модель ставки
type Bet struct {
	ID        int       `json:"id" gorm:"primary_key"` // Идентификатор ставки
	ClientID  int       `json:"client_id"`             // Идентификатор клиента
	EventID   int       `json:"event_id"`              // Идентификатор события
	BetTypeID int       `json:"bet_type_id"`           // Идентификатор типа ставки
	OddID     int       `json:"odd_id"`                // Идентификатор коэффициента (ссылка на Odd)
	BetAmount float64   `json:"bet_amount"`            // Сумма ставки
	Status    string    `json:"status"`                // Статус ставки (например, pending, won, lost)
	BetDate   time.Time `json:"bet_date"`              // Дата ставки
}

// AddBet добавляет новую ставку в базу данных
func AddBet(data map[string]interface{}) error {
	dbConn := db.ConnectDB() // Получаем соединение с базой данных
	defer dbConn.Close()

	bet := Bet{
		ClientID:  data["client_id"].(int),
		EventID:   data["event_id"].(int),
		BetTypeID: data["bet_type_id"].(int),
		OddID:     data["odd_id"].(int),
		BetAmount: data["bet_amount"].(float64),
		Status:    data["status"].(string),
		BetDate:   data["bet_date"].(time.Time),
	}

	_, err := dbConn.Exec("INSERT INTO bet (client_id, event_id, bet_type_id, odd_id, bet_amount, status, bet_date) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		bet.ClientID, bet.EventID, bet.BetTypeID, bet.OddID, bet.BetAmount, bet.Status, bet.BetDate)
	if err != nil {
		return err
	}

	return nil
}
