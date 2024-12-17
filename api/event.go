package api

import (
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

// @Summary AddEvent
// @Description Add a new event
// @Tags events
// @Accept json
// @Produce json
// @Param event body models.Event true "Event"
// @Success 200 {object} models.SuccessResponse "Event added successfully"
// @Failure 400 {object} models.ErrorResponse "Invalid input"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /api/event [post]
// AddEvent обрабатывает HTTP-запрос на добавление события
func AddEvent(c *gin.Context) {
	var event models.Event

	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid input", Details: err.Error()})
		return
	}

	db := db.ConnectDB()
	defer db.Close()
	var eventCount int
	err := db.QueryRow("SELECT COUNT(*) FROM event WHERE name = $1", event.Name).Scan(&eventCount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Could not check event name in database", Details: err.Error()})
		return
	}

	if eventCount > 0 {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Event name already used"})
		return
	}

	// Вызов функции сервиса для добавления события
	if err := event_service.AddEvent(event); err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Could not save event to database", Details: err.Error()})
		return
	}

	eventName := event.Name
	eventID, err := models.GetEventIDByName(eventName)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "GetEventIDByName error", Details: err.Error()})
		return
	}

	odd := models.Odd{
		ID:        eventID,
		OddValue:  rand.Float64() * 2, // Генерация случайного коэффициента от 0 до 2
		EventID:   eventID,            // Привязка к событию
		UpdatedAt: time.Now(),
	}

	// Добавление коэффициента
	if err := odd_service.AddOdd(odd); err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Could not save odd to database", Details: err.Error()})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{Message: "Event added successfully"})
}

// @Summary GetEvent
// @Description Retrieve a single event by its ID
// @Tags events
// @Accept json
// @Produce json
// @Param id path int true "Event ID"
// @Success 200 {object} models.Event "Event details"
// @Failure 400 {object} models.ErrorResponse "Invalid event ID"
// @Failure 500 {object} models.ErrorResponse "Could not retrieve event"
// @Router /api/event/{id} [get]
func GetEvent(c *gin.Context) {
	idStr := c.Param("id") // Получаем ID из параметров URL

	// Преобразование строки в int
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

// @Summary GetEvents
// @Description Retrieve all events
// @Tags events
// @Accept json
// @Produce json
// @Success 200 {array} models.Event "List of events"
// @Failure 500 {object} models.ErrorResponse "Could not retrieve events"
// @Router /api/events [get]
func GetEvents(c *gin.Context) {
	db := db.ConnectDB()
	defer db.Close()

	rows, err := db.Query("SELECT id, name, description, date, status FROM event")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve events", "details": err.Error()})
		return
	}
	defer rows.Close()

	var events []models.Event
	for rows.Next() {
		var event models.Event
		if err := rows.Scan(&event.ID, &event.Name, &event.Desc, &event.Date, &event.Status); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not scan event", "details": err.Error()})
			return
		}
		events = append(events, event)
	}

	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred during rows iteration", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, events)
}
