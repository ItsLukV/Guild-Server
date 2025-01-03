package main

import (
	"fmt"
	"log"
	"time"

	"github.com/ItsLukV/Guild-Server/src/app"
	"github.com/ItsLukV/Guild-Server/src/controllers"
	"github.com/ItsLukV/Guild-Server/src/utils"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"xorm.io/xorm"
)

var appData = app.App{
	Users: make([]app.User, 0),
}

var controller controllers.Controller

func init() {

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

	api := router.Group("/api")

	api.GET("/users", controller.GetUsers)
	api.POST("/users", controller.PostUsers)

	api.GET("/diana/:user", controller.GetDiana)
	api.GET("/dungeons/:user", controller.GetDungeonsData)

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func startDataFetcher() {
	/*
		now := time.Now()
		nextHour := now.Truncate(time.Hour).Add(time.Hour)
		timeUntilNextHour := time.Until(nextHour)

		// Sleep until the next hour
		time.Sleep(timeUntilNextHour)
	*/

	// Start the ticker
	ticker := time.NewTicker(1 * time.Hour)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			log.Println("Fetching data")
			dianaData, dungeonsData := FetchData(appData.Users)

			session := appData.Engine.NewSession()
			defer session.Close()

			// Begin transaction
			if err := session.Begin(); err != nil {
				log.Printf("Failed to start transaction: %v", err)
				return
			}

			// Insert data for Diana
			if len(dianaData) > 0 {
				if err := insertData(session, dianaData, "Diana"); err != nil {
					log.Printf("Error inserting Diana data: %v", err)
					_ = session.Rollback()
					continue
				}
			}

			// Insert data for Dungeons
			if len(dungeonsData) > 0 {
				if err := insertData(session, dungeonsData, "Dungeons"); err != nil {
					log.Printf("Error inserting Dungeons data: %v", err)
					_ = session.Rollback()
					continue
				}
			}

			// Commit transaction
			if err := session.Commit(); err != nil {
				log.Printf("Failed to commit transaction: %v", err)
			} else {
				log.Println("Data inserted successfully!")
			}
		}
	}
}

// insertData is a helper function to insert data into the database and log any errors
func insertData(session *xorm.Session, data interface{}, dataType string) error {
	_, err := session.Insert(data)
	if err != nil {
		return fmt.Errorf("failed to insert %s data: %v", dataType, err)
	}
	log.Printf("%s data inserted successfully!", dataType)
	return nil
}

func FetchData(users []app.User) ([]app.DianaData, []app.DungeonsData) {
	outDiana := make([]app.DianaData, 0)
	outDungeons := make([]app.DungeonsData, 0)

	for _, user := range users {
		uuid, err := utils.GetMCUUID(user.Name)
		if err != nil {
			log.Println("Failed to fetch api: ", err)
		}

		profile := user.ActiveProfileUUID

		data, err := utils.FetchPlayerData(uuid.Id, profile)
		if err != nil {
			log.Println("Failed to fetch api: ", err)
		}

		dianaData := utils.IntoDianaData(*data, user.Id, uuid)
		outDiana = append(outDiana, dianaData)

		dungeonsData := utils.IntoDungeonsData(*data, user.Id, uuid)
		outDungeons = append(outDungeons, dungeonsData)

	}
	return outDiana, outDungeons
}
