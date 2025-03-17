package utils

import (
	"fmt"
	"log"
	"time"

	"github.com/ItsLukV/Guild-Server/src/model"
	"xorm.io/xorm"
)

func InsertPlayerData(engine *xorm.Engine, users []model.User) {
	log.Println("Fetching data")

	// Begin transaction
	session := engine.NewSession()
	defer session.Close()
	if err := session.Begin(); err != nil {
		log.Printf("Failed to start transaction: %v", err)
		return
	}

	data := make(map[string][]model.GuildEventData)

	for _, user := range users {
		if !user.FetchData {
			continue
		}
		dianaPlayerData, dungeonsPlayerData, miningPlayerData := FetchData(user)
		if !hasPlayerData[model.DianaData](engine.NewSession(), user) {
			data["diana"] = append(data["diana"], dianaPlayerData)
		}
		if !hasPlayerData[model.DungeonsData](engine.NewSession(), user) {
			data["dungeons"] = append(data["dungeons"], dungeonsPlayerData)
		}
		if !hasPlayerData[model.MiningData](engine.NewSession(), user) {
			data["mining"] = append(data["mining"], miningPlayerData)
			log.Println(data["mining"])
		}
	}

	for name, playerData := range data {
		// Insert data for Diana
		if len(playerData) > 0 {
			if err := insertData(session, playerData); err != nil {
				log.Printf("Error inserting %s data: %v", name, err)
				_ = session.Rollback()
				return
			}
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
func insertData(session *xorm.Session, data interface{}) error {
	_, err := session.Insert(data)
	if err != nil {
		return fmt.Errorf("failed to insert data: %v", err)
	}
	log.Println("Data inserted successfully!")
	return nil
}

func FetchData(user model.User) (model.DianaData, model.DungeonsData, model.MiningData) {
	profile := user.ActiveProfileUUID

	data, err := FetchPlayerData(user.Id, profile)
	if err != nil {
		log.Println("Failed to fetch api: ", err)
	}

	dianaData := IntoDianaData(*data, user.Id)
	dianaData.FetchTime = time.Now().Truncate(time.Hour)

	dungeonsData := IntoDungeonsData(*data, user.Id)
	dungeonsData.FetchTime = time.Now().Truncate(time.Hour)

	miningData := IntoMiningData(*data, user.Id)
	miningData.FetchTime = time.Now().Truncate(time.Hour)

	return dianaData, dungeonsData, miningData
}

func hasPlayerData[T model.GuildEventData](session *xorm.Session, user model.User) bool {
	playerData := new(T)
	currentTime := time.Now().Truncate(time.Hour)

	found, err := session.Where("user_id = ? AND fetch_time = ?", user.Id, currentTime).Get(playerData)
	if err != nil {
		log.Printf("Error fetching player data: %v", err)
		return false
	}
	session.Close()
	return found
}
