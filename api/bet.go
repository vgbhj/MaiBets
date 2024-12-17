package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vgbhj/MaiBets/db"
	"github.com/vgbhj/MaiBets/models"
	"github.com/vgbhj/MaiBets/service/bet_service"
)

type BetInput struct {
	EventName string    `json:"name" gorm:"not null"`
	BetAmount float64   `json:"bet_amount"`
	BetDate   time.Time `json:"bet_date"`
}

func AddBet(c *gin.Context) {
	// получение пользователя с мидлваре
	userId, exists := c.Get("currentUserId")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"user ID not found": exists})
		return
	}

	// Преобразование userId в int
	userIdInt, ok := userId.(int)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"user ID is not of type int": ok})
		return
	}

	var betInput BetInput
	if err := c.ShouldBindJSON(&betInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	eventName := betInput.EventName
	eventID, err := models.GetEventIDByName(eventName)
	if err != nil {
		// Обработка ошибки
		c.JSON(http.StatusBadRequest, gin.H{"GetEventIDByName error": err.Error()})
		fmt.Println("Error:", err)
		return
	}
	oddId, err := models.GetOddIDByEventId(eventID)
	if err != nil {
		// Обработка ошибки
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println("Error:", err)
		return
	}

	// NO TAK NELZUA DELAT

	// Проверка, хватает ли пользователю денег на балансе
	db := db.ConnectDB()
	defer db.Close()

	var userBalance float64
	err = db.QueryRow("SELECT balance FROM users WHERE id = $1", userIdInt).Scan(&userBalance)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not check user balance", "details": err.Error()})
		return
	}

	if userBalance < betInput.BetAmount {
		c.JSON(http.StatusBadRequest, gin.H{"error": "insufficient balance"})
		return
	}

	bet := models.Bet{
		ID:        1,
		ClientID:  userIdInt,
		EventID:   eventID,
		BetTypeID: 1,
		OddID:     oddId,
		BetAmount: betInput.BetAmount,
		Status:    "In Progress",
		BetDate:   betInput.BetDate,
	}

	// Вызов функции сервиса для добавления ставки
	if err := bet_service.AddBet(bet); err != nil {
		// Выводим текст ошибки в лог (опционально)
		fmt.Println("Error adding event:", err)

		// Возвращаем ошибку в ответе
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save event to database", "details": err.Error()})
		return
	}

	// Уменьшаем баланс пользователя на сумму ставки
	newBalance := userBalance - betInput.BetAmount
	_, err = db.Exec("UPDATE users SET balance = $1 WHERE id = $2", newBalance, userIdInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update user balance", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Bet added successfully"})
}

func GetBets(c *gin.Context) {
	// получение пользователя с мидлваре
	userId, exists := c.Get("currentUserId")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"user ID not found": exists})
		return
	}

	// Преобразование userId в int
	userIdInt, ok := userId.(int)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"user ID is not of type int": ok})
		return
	}

	db := db.ConnectDB()
	defer db.Close()

	// Получаем все ставки пользователя
	rows, err := db.Query("SELECT id, event_id, bet_amount, status, bet_date FROM bet WHERE client_id = $1", userIdInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve bets", "details": err.Error()})
		return
	}
	defer rows.Close()

	var bets []models.Bet
	for rows.Next() {
		var bet models.Bet
		if err := rows.Scan(&bet.ID, &bet.EventID, &bet.BetAmount, &bet.Status, &bet.BetDate); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not scan bet", "details": err.Error()})
			return
		}
		bets = append(bets, bet)
	}

	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred during rows iteration", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, bets)
}
