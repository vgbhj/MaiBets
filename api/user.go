package api

import (
	"database/sql"
	"net/http"
	"strconv"

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
	// Получение userId из параметров запроса
	userId := c.Param("userId")
	if userId == "" {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "User ID not provided"})
		return
	}

	// Преобразование userId в int
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid user ID"})
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
