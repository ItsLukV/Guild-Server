package controllers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"slices"

	"github.com/ItsLukV/Guild-Server/src/app"
	"github.com/gin-gonic/gin"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

func (con *Controller) CreateGuildEvent(c *gin.Context) {
	// TODO check if player is the db.
	// TODO check if guild evnet is already created.

	var request struct {
		Users     []string           `json:"users" binding:"required"`
		Duration  int                `json:"duration" binding:"required"`
		Type      app.GuildEventType `json:"type" binding:"required"`
		StartTime time.Time          `json:"start_time"`
	}

	// Bind JSON input to the GuildEvent struct
	if err := c.ShouldBindJSON(&request); err != nil {
		con.ErrorResponseWithUUID(c, http.StatusBadRequest, err, "Invalid JSON input")
		return
	}

	guildTypes := []app.GuildEventType{app.Dungeons, app.Diana}

	if !slices.Contains(guildTypes, request.Type) {
		con.ErrorResponseWithUUID(c, http.StatusBadRequest, nil, fmt.Sprintf("Invalid guild event type: %s", request.Type))
		return
	}

	if request.StartTime.IsZero() {
		request.StartTime = time.Now().Truncate(time.Hour)
	} else {
		request.StartTime = request.StartTime.Truncate(time.Hour)
	}

	alphabet := "abcdefghijklmnopqrstuvwxyz1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	id, err := gonanoid.Generate(alphabet, 21)
	if err != nil {
		con.ErrorResponseWithUUID(c, http.StatusInternalServerError, err, "Failed to generate guild event ID")
		return
	}

	guildEvent := app.GuildEvent{
		Id:        id,
		Users:     request.Users,
		Duration:  request.Duration,
		Type:      request.Type,
		StartTime: request.StartTime,
	}

	// Insert the new guild event into the database
	_, err = con.AppData.Engine.Insert(&guildEvent)
	if err != nil {
		con.ErrorResponseWithUUID(c, http.StatusInternalServerError, err, "Failed to insert guild event")
		return
	}

	// Return success response
	c.JSON(http.StatusCreated, gin.H{
		"message":  "Guild event created successfully",
		"guild id": guildEvent.Id,
	})
}

func (con *Controller) GetGuildEvent(c *gin.Context) {
	// TODO: Implement this function

	type GuildEventResponse struct {
		EventID   string             `json:"event_id"`
		StartTime time.Time          `json:"start_time"`
		Duration  int                `json:"duration"`
		Type      app.GuildEventType `json:"type"`
		UserIDs   []string           `json:"user_ids"`
		EventData []app.GuildData    `json:"event_data"`
	}

	id := c.Query("id")
	event := app.GuildEvent{Id: id}
	has, err := con.AppData.Engine.Get(&event)

	if err != nil {
		con.ErrorResponseWithUUID(c, http.StatusInternalServerError, err, "Failed to fetch guild event")
		return
	}

	if !has {
		con.ErrorResponseWithUUID(c, http.StatusNotFound, nil, fmt.Sprintf("Guild event with ID %s not found", id))
		return
	}

	var guildData []app.GuildData

	switch event.Type {
	case app.Dungeons:
		dungeonsData, err := fetchPlayerData[app.DungeonsData](con, event)
		if err != nil {
			con.ErrorResponseWithUUID(c, http.StatusInternalServerError, err, "Failed to fetch dungeons data")
			return
		}

		for _, data := range dungeonsData {
			guildData = append(guildData, data)
		}
	case app.Diana:
		dianaData, err := fetchPlayerData[app.DianaData](con, event)
		if err != nil {
			con.ErrorResponseWithUUID(c, http.StatusInternalServerError, err, "Failed to fetch Diana data")
			return
		}

		for _, data := range dianaData {
			guildData = append(guildData, data)
		}
	}

	// Return success response
	guildEventResponse := GuildEventResponse{
		EventID:   event.Id,
		StartTime: event.StartTime,
		Duration:  event.Duration,
		Type:      event.Type,
		UserIDs:   event.Users,
		EventData: guildData,
	}

	c.JSON(http.StatusOK, guildEventResponse)
}

func fetchPlayerData[T app.GuildData](con *Controller, event app.GuildEvent) ([]T, error) {
	records := make([]T, 0)

	con.AppData.Engine.ShowSQL(true)

	var err error
	// Query for all players at the specific FetchTime.
	err = con.AppData.Engine.
		Where("fetch_time = ?", event.StartTime.Add(time.Duration(event.Duration)*time.Hour)).
		Find(&records)
	if err != nil {
		return nil, fmt.Errorf("failed to query records for specific FetchTime: %v", err)
	}

	// If no records are found for the specific FetchTime, fetch the latest records for all players.
	if len(records) == 0 {
		var instance T
		tableName := instance.TableName()
		query := fmt.Sprintf(`
            SELECT *
            FROM %s AS d
            INNER JOIN (
                SELECT user_id, MAX(fetch_time) AS LatestFetchTime
                FROM %s
                GROUP BY user_id
            ) AS latest
            ON d.user_id = latest.user_id AND d.fetch_time = latest.LatestFetchTime
        `, tableName, tableName)

		err = con.AppData.Engine.SQL(query).Find(&records)
		if err != nil {
			return nil, fmt.Errorf("failed to query latest records: %v", err)
		}
	}

	startRecords := make([]T, 0)
	err = con.AppData.Engine.Where("DATE_TRUNC('hour', fetch_time) = ?", event.StartTime).Find(&startRecords)

	if err != nil {
		return nil, fmt.Errorf("failed to query records for specific FetchTime: %v", err)
	}

	if len(startRecords) != len(records) {
		return nil, fmt.Errorf("failed to fetch the same number of records for the start time (len: %v) and latest records (len: %v)", len(startRecords), len(records))
	}

	// Subtract the records at the start time from the latest records to get the event data
	for i, record := range records {

		log.Println("Record: ", record)
		log.Println("StartRecord: ", startRecords[i])
		startRecord := startRecords[i]
		subtractedRecord, err := record.Subtract(startRecord)
		if err != nil {
			return nil, fmt.Errorf("failed to subtract records: %v", err)
		}

		// Type assertion to ensure we assign a T type value
		if subtractedRecordT, ok := subtractedRecord.(T); ok {
			records[i] = subtractedRecordT
		} else {
			return nil, fmt.Errorf("failed to assert the subtracted record as type T")
		}
	}

	return records, nil
}
