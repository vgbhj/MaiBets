package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vgbhj/MaiBets/models"
)

// @Summary GetUser
// @Description Retrieve user information by ID
// @Tags users
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} models.User "User information"
// @Failure 400 {object} models.ErrorResponse "Invalid user ID"
// @Failure 404 {object} models.ErrorResponse "User not found"
// @Failure 500 {object} models.ErrorResponse "Could not retrieve user"
// @Router /api/user/ [get]
// GetUser обрабатывает HTTP-запрос на получение информации о пользователе
func GetUser(c *gin.Context) {
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

	// Получение пользователя из базы данных
	user, err := models.GetUser(userIdInt)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "User not found"})
		} else {
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Could not retrieve user", Details: err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, user)
}
