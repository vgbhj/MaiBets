package models

import (
	"fmt"
	"time"

	"github.com/vgbhj/MaiBets/db"
)

type User struct {
	ID          int     `json:"id" gorm:"primary_key"`
	Username    string  `json:"username" gorm:"unique"`
	Password    string  `json:"password"`
	FirstName   string  `json:"first_name"`
	LastName    string  `json:"last_name"`
	Phone       string  `json:"phone"`
	Balance     float64 `json:"balance" gorm:"default:100"`
	AccessLevel int     `json:"access_level"`
	CreatedAt   time.Time
}

// GetUser получает пользователя по его ID
func GetUser(id int) (*User, error) {
	dbConn := db.ConnectDB() // Получаем соединение с базой данных
	defer dbConn.Close()

	var user User

	err := dbConn.QueryRow("SELECT id, username, password, first_name, last_name, phone, balance, access_level FROM users WHERE id = $1", id).Scan(
		&user.ID, &user.Username, &user.Password, &user.FirstName, &user.LastName, &user.Phone, &user.Balance, &user.AccessLevel)

	if err != nil {
		// Выводим ошибку в лог
		fmt.Println("Error retrieving user:", err)

		// Возвращаем ошибку
		return nil, err
	}

	return &user, nil
}
