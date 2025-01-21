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
	dianaData, dungeonsData := FetchData(users)

	session := engine.NewSession()
	defer session.Close()

	// Begin transaction
	if err := session.Begin(); err != nil {
		log.Printf("Failed to start transaction: %v", err)
		return
	}

	// InserUsert data for Diana
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

func FetchData(users []app.User) ([]app.DianaData, []app.DungeonsData) {
	outDiana := make([]app.DianaData, 0)
	outDungeons := make([]app.DungeonsData, 0)

	for _, user := range users {
		if !user.FetchData {
			continue
		}

		profile := user.ActiveProfileUUID

		data, err := FetchPlayerData(user.Id, profile)
		if err != nil {
			log.Println("Failed to fetch api: ", err)
		}

		dianaData := IntoDianaData(*data, user.Id, user.Id)
		dianaData.FetchTime = time.Now().Truncate(time.Hour)
		outDiana = append(outDiana, dianaData)

		dungeonsData := IntoDungeonsData(*data, user.Id, user.Id)
		dungeonsData.FetchTime = time.Now().Truncate(time.Hour)
		outDungeons = append(outDungeons, dungeonsData)

	}
	return outDiana, outDungeons
}
