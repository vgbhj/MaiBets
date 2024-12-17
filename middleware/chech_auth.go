package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/vgbhj/MaiBets/db"
	"github.com/vgbhj/MaiBets/models"
)

func CheckAuth(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{Error: "A"})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	authToken := strings.Split(authHeader, " ")
	if len(authToken) != 2 || authToken[0] != "Bearer" {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{Error: "Invalid token format"})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	tokenString := authToken[1]
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET")), nil
	})

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{Error: "Invalid or expired token"})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{Error: "Invalid token"})
		c.Abort()
		return
	}

	if float64(time.Now().Unix()) > claims["exp"].(float64) {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{Error: "Token expired"})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	var user models.User

	db := db.ConnectDB()
	defer db.Close()
	row := db.QueryRow("SELECT id, username FROM users WHERE id = $1", claims["id"])
	err2 := row.Scan(&user.ID, &user.Username)
	if err2 != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	c.Set("currentUser", user)
	c.Set("currentUserId", user.ID)

	c.Next()
}
