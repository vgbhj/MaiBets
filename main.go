package main

import (
	"github.com/gin-gonic/gin"
	"github.com/vgbhj/MaiBets/api"
	"github.com/vgbhj/MaiBets/config"
	"github.com/vgbhj/MaiBets/db"
	"github.com/vgbhj/MaiBets/middleware"
)

func init() {
	config.LoadEnvs()
	db.ConnectDB()

}

func main() {
	router := gin.Default()

	router.POST("/api/auth/signup", api.CreateUser)
	router.POST("/api/auth/login", api.Login)
	router.POST("/api/event", api.AddEvent)
	router.GET("/api/event/:id", api.GetEvent)
	router.GET("/api/events", api.GetEvents)
	// router.GET("/user/profile", middlewares.CheckAuth, controllers.GetUserProfile)
	router.POST("/api/bet", middleware.CheckAuth, api.AddBet)
	router.GET("/api/bets", middleware.CheckAuth, api.GetBets)
	router.Run()
}
