package api

import (
	"fmt"
	"math/rand/v2"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vgbhj/MaiBets/db"
	"github.com/vgbhj/MaiBets/models"
	"github.com/vgbhj/MaiBets/service/event_service"
	"github.com/vgbhj/MaiBets/service/odd_service"
)

// AddMaterialHandler обрабатывает HTTP-запрос на добавление ивента
func AddEvent(c *gin.Context) {
	var event models.Event

	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// NO TAK NELZUA DELAT

	db := db.ConnectDB()
	defer db.Close()
	var eventCount int
	err := db.QueryRow("SELECT COUNT(*) FROM users WHERE username = $1", event.Name).Scan(&eventCount)
	if eventCount > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "event name already used"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not check event name in database", "details": err.Error()})
		return
	}

	// Вызов функции сервиса для добавления ивента
	if err := event_service.AddEvent(event); err != nil {
		// Выводим текст ошибки в лог (опционально)
		fmt.Println("Error adding event:", err)

		// Возвращаем ошибку в ответе
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save event to database", "details": err.Error()})
		return
	}
	eventName := event.Name
	eventID, err := models.GetEventIDByName(eventName)
	if err != nil {
		// Обработка ошибки
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Event ID:", eventID)
	}

	odd := models.Odd{
		ID:        eventID,
		OddValue:  rand.Float64() * 2, // Генерируем случайное значение от 0 до 2
		EventID:   eventID,            // Ссылка на событие
		UpdatedAt: time.Now(),         // Устанавливаем текущее время
	}

	// Вызов функции сервиса для добавления кефа
	if err := odd_service.AddOdd(odd); err != nil {
		// Выводим текст ошибки в лог (опционально)
		fmt.Println("Error adding event:", err)

		// Возвращаем ошибку в ответе
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save odd to database", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Event added successfully"})
}

func GetEvent(c *gin.Context) {
	idStr := c.Param("id") // Получаем ID из параметров URL

	// Преобразуем строку в int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}

	event, err := event_service.GetEvent(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve event"})
		return
	}

	c.JSON(http.StatusOK, event)
}
