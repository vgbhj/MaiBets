package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vgbhj/MaiBets/db"
	"github.com/vgbhj/MaiBets/models"
	"github.com/vgbhj/MaiBets/service/bet_service"
)

// BetInput - структура для ввода ставки
type BetInput struct {
	EventName string    `json:"name" gorm:"not null" example:"Champions League Final"`
	BetAmount float64   `json:"bet_amount" example:"10.10"`
	BetDate   time.Time `json:"bet_date" example:"2024-06-01T20:00:00Z"`
}

// @Summary AddBet
// @Description Add a new bet for the current user
// @Tags bets
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param bet body BetInput true "Bet Details"
// @Success 200 {object} models.SuccessResponse "Bet added successfully"
// @Failure 400 {object} models.ErrorResponse "Invalid input or insufficient balance"
// @Failure 400 {object} models.ErrorResponse "Event is finished"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /api/bet [post]
// AddBet обрабатывает HTTP-запрос на добавление новой ставки
func AddBet(c *gin.Context) {
	// получение пользователя с мидлваре
	userId, exists := c.Get("currentUserId")
	if !exists {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "User ID not found"})
		return
	}

	// Преобразование userId в int
	userIdInt, ok := userId.(int)
	if !ok {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "User ID is not of type int"})
		return
	}

	// Валидация входных данных
	var betInput BetInput
	if err := c.ShouldBindJSON(&betInput); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid input", Details: err.Error()})
		return
	}

	// Получаем ID события по его имени
	eventID, err := models.GetEventIDByName(betInput.EventName)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Event not found", Details: err.Error()})
		return
	}

	// Получаем статус события
	eventStatus, err := models.GetEventStatusByID(eventID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Could not retrieve event status", Details: err.Error()})
		return
	}

	// Проверка, завершено ли событие
	if eventStatus == "finished" {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Event is finished"})
		return
	}

	// Получаем ID коэффициента для события
	oddId, err := models.GetOddIDByEventId(eventID)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Odd not found", Details: err.Error()})
		return
	}

	// Подключение к БД и проверка баланса пользователя
	db := db.ConnectDB()
	defer db.Close()

	var userBalance float64
	err = db.QueryRow("SELECT balance FROM users WHERE id = $1", userIdInt).Scan(&userBalance)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Could not check user balance", Details: err.Error()})
		return
	}

	// Проверка достаточности баланса
	if userBalance < betInput.BetAmount {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Insufficient balance"})
		return
	}

	// Создание модели ставки
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

	// Добавление ставки через сервис
	if err := bet_service.AddBet(bet); err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Could not save bet to database", Details: err.Error()})
		return
	}

	// Обновление баланса пользователя
	newBalance := userBalance - betInput.BetAmount
	_, err = db.Exec("UPDATE users SET balance = $1 WHERE id = $2", newBalance, userIdInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Could not update user balance", Details: err.Error()})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{Message: "Bet added successfully"})
}

// @Summary GetBets
// @Description Retrieve all bets for the current user
// @Tags bets
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Success 200 {array} models.Bet "List of user bets"
// @Failure 400 {object} models.ErrorResponse "Invalid user ID"
// @Failure 500 {object} models.ErrorResponse "Could not retrieve bets"
// @Router /api/bets [get]
// GetBets обрабатывает HTTP-запрос на получение всех ставок пользователя
func GetBets(c *gin.Context) {
	// получение пользователя с мидлваре
	userId, exists := c.Get("currentUserId")
	if !exists {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "User ID not found"})
		return
	}

	// Преобразование userId в int
	userIdInt, ok := userId.(int)
	if !ok {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "User ID is not of type int"})
		return
	}

	// Подключение к БД и получение списка ставок
	db := db.ConnectDB()
	defer db.Close()

	rows, err := db.Query("SELECT id, event_id, bet_amount, status, bet_date FROM bet WHERE client_id = $1", userIdInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Could not retrieve bets", Details: err.Error()})
		return
	}
	defer rows.Close()

	var bets []models.Bet
	for rows.Next() {
		var bet models.Bet
		if err := rows.Scan(&bet.ID, &bet.EventID, &bet.BetAmount, &bet.Status, &bet.BetDate); err != nil {
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Could not scan bet", Details: err.Error()})
			return
		}
		bets = append(bets, bet)
	}

	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Error occurred during rows iteration", Details: err.Error()})
		return
	}

	c.JSON(http.StatusOK, bets)
}
