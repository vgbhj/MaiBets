package main

import (
	"github.com/gin-gonic/gin"
	"github.com/vgbhj/MaiBets/api"
	"github.com/vgbhj/MaiBets/config"
	"github.com/vgbhj/MaiBets/db"
)

func init() {
	config.LoadEnvs()
	db.ConnectDB()

}

func main() {
	router := gin.Default()

	router.POST("/auth/signup", api.CreateUser)
	router.POST("/auth/login", api.Login)
	router.POST("/event", api.AddEvent)
	router.GET("/event/:id", api.GetEvent)
	// router.GET("/user/profile", middlewares.CheckAuth, controllers.GetUserProfile)
	router.Run()
}
