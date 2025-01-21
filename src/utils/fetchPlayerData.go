package utils

import (
	"fmt"
	"log"
	"time"

	"github.com/ItsLukV/Guild-Server/src/app"
	"xorm.io/xorm"
)

func InsertPlayerData(engine *xorm.Engine, users []app.User) {
	log.Println("Fetching data")

	// Begin transaction
	session := engine.NewSession()
	defer session.Close()
	if err := session.Begin(); err != nil {
		log.Printf("Failed to start transaction: %v", err)
		return
	}

	var dianaData []app.DianaData
	var dungeonsData []app.DungeonsData
	for _, user := range users {
		if !user.FetchData {
			continue
		}
		dianaPlayerData, dungeonsPlayerData := FetchData(user)
		if !hasPlayerData[app.DianaData](engine.NewSession(), user) {
			dianaData = append(dianaData, dianaPlayerData)
		}
		if !hasPlayerData[app.DungeonsData](engine.NewSession(), user) {
			dungeonsData = append(dungeonsData, dungeonsPlayerData)
		}
	}

	// Insert data for Diana
	if len(dianaData) > 0 {
		if err := insertData(session, dianaData, "Diana"); err != nil {
			log.Printf("Error inserting Diana data: %v", err)
			_ = session.Rollback()
			return
		}
	}

	// Insert data for Dungeons
	if len(dungeonsData) > 0 {
		if err := insertData(session, dungeonsData, "Dungeons"); err != nil {
			log.Printf("Error inserting Dungeons data: %v", err)
			_ = session.Rollback()
			return
		}
	}

	// Commit transaction
	if err := session.Commit(); err != nil {
		log.Printf("Failed to commit transaction: %v", err)
	} else {
		log.Printf("Data inserted successfully for all users (%v)!", len(users))
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

func FetchData(user app.User) (app.DianaData, app.DungeonsData) {
	profile := user.ActiveProfileUUID

	data, err := FetchPlayerData(user.Id, profile)
	if err != nil {
		log.Println("Failed to fetch api: ", err)
	}

	dianaData := IntoDianaData(*data, user.Id)
	dianaData.FetchTime = time.Now().Truncate(time.Hour)

	dungeonsData := IntoDungeonsData(*data, user.Id)
	dungeonsData.FetchTime = time.Now().Truncate(time.Hour)

	return dianaData, dungeonsData
}

func hasPlayerData[T app.GuildEventData](session *xorm.Session, user app.User) bool {
	playerData := new(T)
	currentTime := time.Now().Truncate(time.Hour)

	found, err := session.Where("user_id = ? AND fetch_time = ?", user.Id, currentTime).Get(playerData)
	if err != nil {
		log.Printf("Error fetching player data: %v", err)
		return false
	}

	return found
}
