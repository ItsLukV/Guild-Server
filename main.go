package main

import (
	"log"
	"time"

	"github.com/ItsLukV/Guild-Server/src/app"
	"github.com/ItsLukV/Guild-Server/src/controllers"
	"github.com/ItsLukV/Guild-Server/src/middleware"
	"github.com/ItsLukV/Guild-Server/src/utils"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var appData = app.App{
	Users: make([]app.User, 0),
}

var controller controllers.Controller

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	// Load environment variables
	if err := godotenv.Load(".env"); err != nil {
		log.Println("Error loading .env file")
	}
}

func main() {
	var err error
	appData.Engine, err = app.SyncDatabase(&appData.Users) // app.SyncDatabase returns *xorm.Engine
	if err != nil {
		log.Fatalf("Failed to sync database: %v", err)
	}
	defer appData.Engine.Close()

	controller = *controllers.NewController(&appData)

	// Start a background Goroutine to fetch data every hour
	go startDataFetcher()

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")

	router.GET("/", controller.GetDefault)

	needToken := router.Group("/api")

	needToken.GET("/users", controller.GetUsers)
	needToken.GET("/guildevent", controller.GetGuildEvent)

	// Apply the TokenAuthMiddleware to all routes in this group
	needToken.Use(middleware.TokenAuthMiddleware(&appData))

	// Define the routes in this group
	needToken.POST("/users", controller.PostUsers)
	needToken.GET("/diana/:user", controller.GetDiana)
	needToken.GET("/dungeons/:user", controller.GetDungeonsData)

	needToken.POST("/guildevent", controller.CreateGuildEvent)

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func startDataFetcher() {
	now := time.Now()
	nextHour := now.Truncate(time.Hour).Add(time.Hour)
	timeUntilNextHour := time.Until(nextHour)

	// Sleep until the next hour
	time.Sleep(timeUntilNextHour)

	// Start the ticker
	ticker := time.NewTicker(1 * time.Hour)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			utils.InsertPlayerData(appData.Engine, appData.Users)
		}
	}
}
