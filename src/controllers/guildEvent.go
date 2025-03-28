package controllers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"slices"

	"github.com/ItsLukV/Guild-Server/src/model"
	"github.com/gin-gonic/gin"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"xorm.io/xorm"
)

func (con *Controller) CreateGuildEvent(c *gin.Context) {
	// Define the request struct
	var request struct {
		Users     []string             `json:"users" binding:"required"`
		Duration  int                  `json:"duration" binding:"required"`
		Type      model.GuildEventType `json:"type" binding:"required"`
		StartTime time.Time            `json:"start_time"`
		IsHidden  bool                 `json:"is_hidden"`
	}

	// Bind JSON input to the GuildEvent struct
	if err := c.ShouldBindJSON(&request); err != nil {
		con.ErrorResponseWithUUID(c, http.StatusBadRequest, err, "Invalid JSON input")
		return
	}

	// Check if guild has a known event type
	guildTypes := []model.GuildEventType{model.Dungeons, model.Diana, model.Mining}

	if !slices.Contains(guildTypes, request.Type) {
		con.ErrorResponseWithUUID(c, http.StatusBadRequest, nil, fmt.Sprintf("Invalid guild event type: %s", request.Type))
		return
	}

	// Check if the user is in db
	if request.StartTime.IsZero() {
		request.StartTime = time.Now().Truncate(time.Hour)
	} else {
		request.StartTime = request.StartTime.Truncate(time.Hour)
	}

	// Create a id for guild event
	alphabet := "abcdefghijklmnopqrstuvwxyz1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	id, err := gonanoid.Generate(alphabet, 21)
	if err != nil {
		con.ErrorResponseWithUUID(c, http.StatusInternalServerError, err, "Failed to generate guild event ID")
		return
	}

	// Create event
	guildEvent := model.GuildEvent{
		Id:        id,
		Users:     request.Users,
		Duration:  request.Duration,
		Type:      request.Type,
		StartTime: request.StartTime,
		IsHidden:  request.IsHidden,
	}

	// Insert the new guild event into the database
	session := con.AppData.Engine.NewSession()
	defer session.Close()

	// Begin session
	if err := session.Begin(); err != nil {
		log.Printf("Failed to start transaction: %v", err)
		return
	}

	missingUsers, err := checkForMissingUsers(session, guildEvent.Users)
	if err != nil {
		con.ErrorResponseWithUUID(c, http.StatusInternalServerError, err, "Failed to check for missing users")
		session.Rollback()
		return
	}

	if len(missingUsers) > 0 {
		con.ErrorResponseWithUUID(c, http.StatusNotFound, nil, fmt.Sprintf("Missing users: %v", missingUsers))
		session.Rollback()
		return
	}

	_, err = session.Insert(&guildEvent)
	if err != nil {
		con.ErrorResponseWithUUID(c, http.StatusInternalServerError, err, "Failed to insert guild event")
		session.Rollback()
		return
	}

	session.Commit()

	// Return success response
	c.JSON(http.StatusCreated, gin.H{
		"message":  "Guild event created successfully",
		"guild id": guildEvent.Id,
	})
}

func checkForMissingUsers(session *xorm.Session, users []string) ([]string, error) {
	// Query the database to check which user IDs exist
	var existingUsers []model.User
	err := session.In("id", users).Find(&existingUsers)
	if err != nil {
		log.Fatalf("Failed to fetch users from the database: %v", err)
	}

	// Check for missing IDs
	existingIDs := make(map[string]bool)
	for _, user := range existingUsers {
		existingIDs[user.Id] = true
	}

	var missingIDs []string
	for _, id := range users {
		if !existingIDs[id] {
			missingIDs = append(missingIDs, id)
		}
	}

	return missingIDs, nil
}

func (con *Controller) GetGuildEvent(c *gin.Context) {
	type GuildEventResponse struct {
		EventID   string                 `json:"id"`
		StartTime time.Time              `json:"start_time"`
		Duration  int                    `json:"duration"`
		Type      model.GuildEventType   `json:"type"`
		UserIDs   []string               `json:"users"`
		EventData []model.GuildEventData `json:"event_data"`
	}

	session := con.AppData.Engine.NewSession()
	defer session.Close()

	// fetch the id of guild event
	id := c.Query("id")
	event := model.GuildEvent{Id: id}

	if id == "" {
		con.ErrorResponseWithUUID(c, http.StatusBadRequest, nil, "Missing guild event ID")
		return
	}

	if err := session.Begin(); err != nil {
		log.Printf("Failed to start transaction: %v", err)
		return
	}

	// Checking if guild exits
	has, err := session.Get(&event)
	if err != nil {
		con.ErrorResponseWithUUID(c, http.StatusInternalServerError, err, "Failed to fetch guild event")
		session.Rollback()
		return
	}

	if !has {
		con.ErrorResponseWithUUID(c, http.StatusNotFound, nil, fmt.Sprintf("Guild event with ID %s not found", id))
		session.Rollback()
		return
	}

	var guildData []model.GuildEventData

	if event.IsHidden {
		// Return success response
		guildEventResponse := GuildEventResponse{
			EventID:   event.Id,
			StartTime: event.StartTime,
			Duration:  event.Duration,
			Type:      event.Type,
			UserIDs:   event.Users,
			EventData: []model.GuildEventData{},
		}

		c.JSON(http.StatusOK, guildEventResponse)
		return
	}

	println(event.Type)

	// Getting guild data
	switch event.Type {
	case model.Dungeons:
		dungeonsData, err := fetchPlayerData[model.DungeonsData](session, event)
		if err != nil {
			con.ErrorResponseWithUUID(c, http.StatusInternalServerError, err, "Failed to fetch dungeons data")
			session.Rollback()
			return
		}

		for _, data := range dungeonsData {
			guildData = append(guildData, data)
		}
	case model.Diana:
		dianaData, err := fetchPlayerData[model.DianaData](session, event)
		if err != nil {
			con.ErrorResponseWithUUID(c, http.StatusInternalServerError, err, "Failed to fetch miana data")
			session.Rollback()
			return
		}

		for _, data := range dianaData {
			guildData = append(guildData, data)
		}
	case model.Mining:
		Data, err := fetchPlayerData[model.MiningData](session, event)
		if err != nil {
			con.ErrorResponseWithUUID(c, http.StatusInternalServerError, err, "Failed to fetch mining data")
			session.Rollback()
			return
		}

		for _, data := range Data {
			guildData = append(guildData, data)
		}
	}

	session.Commit()

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

func fetchPlayerData[T model.GuildEventData](session *xorm.Session, event model.GuildEvent) ([]T, error) {
	var records []T
	// records := make([]T, 0)
	endTime := event.StartTime.Add(time.Duration(event.Duration) * time.Hour)

	var err error

	// Query for all players at the specific FetchTime.
	err = session.Where("fetch_time = ?", endTime).In("user_id", event.Users).Find(&records)
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
            ON d.user_id = latest.user_id AND d.fetch_time = latest.LatestFetchTime;
        `, tableName, tableName)

		var recordsCopy []T

		err = session.SQL(query).Find(&recordsCopy)
		if err != nil {
			return nil, fmt.Errorf("failed to query latest records: %v", err)
		}

		for _, eventData := range recordsCopy {
			if slices.Contains(event.Users, eventData.GetUserID()) {
				records = append(records, eventData)
			}
		}
	}

	// Get the records guild event start
	startRecords := make([]T, 0)
	err = session.Where("fetch_time = ?", event.StartTime).In("user_id", event.Users).Find(&startRecords)

	if err != nil {
		return nil, fmt.Errorf("failed to query records for specific FetchTime: %v", err)
	}

	if len(startRecords) != len(records) {
		return nil, fmt.Errorf("failed to fetch the same number of records for the start time (len: %v) and latest records (len: %v)", len(startRecords), len(records))
	}

	// Subtract the records at the start time from the latest records to get the event data
	for i, record := range records {
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

func (con *Controller) GetGuildEvents(c *gin.Context) {
	session := con.AppData.Engine.NewSession()
	defer session.Close()

	// Begin session
	if err := session.Begin(); err != nil {
		con.ErrorResponseWithUUID(c, http.StatusInternalServerError, err, "Failed to start transaction")
		return
	}

	var guildEvents []model.GuildEvent
	err := session.Find(&guildEvents)
	if err != nil {
		con.ErrorResponseWithUUID(c, http.StatusInternalServerError, err, "Failed to fetch guild events")
		session.Rollback()
		return
	}
	session.Commit()

	if len(guildEvents) == 0 {
		c.JSON(http.StatusOK, make([]model.GuildEvent, 0))
		return
	}

	// Return success response
	c.JSON(http.StatusOK, guildEvents)
}
