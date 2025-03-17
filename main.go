package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/ItsLukV/Guild-Server/src/app"
	"github.com/ItsLukV/Guild-Server/src/controllers"
	"github.com/ItsLukV/Guild-Server/src/model"
	"github.com/ItsLukV/Guild-Server/src/utils"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var appData = model.App{
	Users: make([]model.User, 0),
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

	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())

	// router.LoadHTMLGlob("templates/*")
	router.LoadHTMLFiles("public/index.tmpl")

	router.GET("/", controller.GetDefault)

	api := router.Group("/api")

	api.GET("/users", controller.GetUsers)
	api.GET("/user", controller.GetUser)
	api.GET("/guildevent", controller.GetGuildEvent)
	api.GET("/guildevents", controller.GetGuildEvents)

	// Apply the TokenAuthMiddleware to all routes in this group
	api.Use(controller.TokenAuthMiddleware(&appData))

	// Define the routes in this group
	api.POST("/users", controller.PostUsers)

	api.POST("/guildevent", controller.CreateGuildEvent)

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	} else {
		log.Println("Server started successfully")
	}
}

func startDataFetcher() {

	utils.InsertPlayerData(appData.Engine, appData.Users)

	now := time.Now()

	var updateTime time.Duration
	number, err := strconv.Atoi(os.Getenv("UPDATE_TIME"))
	if os.Getenv("UPDATE_TIME") == "" || err != nil {
		updateTime = time.Hour
	} else {
		updateTime = time.Minute * time.Duration(number)
	}

	nextHour := now.Truncate(updateTime).Add(updateTime)
	timeUntilNextHour := time.Until(nextHour)

	// Sleep until the next hour
	time.Sleep(timeUntilNextHour)

	utils.InsertPlayerData(appData.Engine, appData.Users)

	// Start the ticker
	ticker := time.NewTicker(1 * updateTime)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			utils.InsertPlayerData(appData.Engine, appData.Users)
		}
	}
}
