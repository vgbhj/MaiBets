package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vgbhj/MaiBets/models"
	"github.com/vgbhj/MaiBets/service/event_service"
)

// AddMaterialHandler обрабатывает HTTP-запрос на добавление ивента
func AddEvent(c *gin.Context) {
	var event models.Event

	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Вызов функции сервиса для добавления материала
	if err := event_service.AddEvent(event); err != nil {
		// Выводим текст ошибки в лог (опционально)
		fmt.Println("Error adding event:", err)

		// Возвращаем ошибку в ответе
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save event to database", "details": err.Error()})
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

	material, err := event_service.GetEvent(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve event"})
		return
	}

	c.JSON(http.StatusOK, material)
}
