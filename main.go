package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/vgbhj/MaiBets/api"
	"github.com/vgbhj/MaiBets/config"
	"github.com/vgbhj/MaiBets/db"
	_ "github.com/vgbhj/MaiBets/docs"
	"github.com/vgbhj/MaiBets/middleware"
	"github.com/vgbhj/MaiBets/service/ticker_service"
)

func init() {
	config.LoadEnvs()
	db.ConnectDB()
}

func main() {
	router := gin.Default()
	ticker_service.StartTicker()

	// Serve static files (CSS, JS, images)
	router.Static("/static", "./frontend/static")

	// Load HTML templates from the frontend folder
	router.LoadHTMLGlob("frontend/*.html")

	// Frontend routes
	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	router.GET("/login", func(c *gin.Context) {
		c.HTML(200, "login.html", nil)
	})

	router.GET("/user_bets", func(c *gin.Context) {
		c.HTML(200, "user_bets.html", nil)
	})

	router.GET("/event/:id", func(c *gin.Context) {
		eventID := c.Param("id")
		c.HTML(200, "event.html", gin.H{
			"EventID": eventID,
		})
	})

	// API routes
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.POST("/api/signup", api.CreateUser)
	router.POST("/api/login", api.Login)

	router.POST("/api/event", api.AddEvent)
	router.GET("/api/event/:id", api.GetEvent)
	router.GET("/api/events", api.GetEvents)

	router.POST("/api/bet", middleware.CheckAuth, api.AddBet)
	router.GET("/api/bets", middleware.CheckAuth, api.GetBets)

	router.GET("/api/user", middleware.CheckAuth, api.GetUser)

	// Run the server
	router.Run()
}
