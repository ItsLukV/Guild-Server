package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ItsLukV/Guild-Server/src/app"
	"github.com/ItsLukV/Guild-Server/src/utils"
	"github.com/gin-gonic/gin"
)

func (con *Controller) GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"users": con.AppData.Users,
	})
}

func (con *Controller) PostUsers(c *gin.Context) {
	// Define a temporary struct for binding only the Uuid field
	var input struct {
		Uuid             string `json:"uuid" binding:"required"`
		DiscordSnowflake string `json:"discord_username" binding:"required"`
	}

	// Bind JSON input to the input struct
	if err := c.ShouldBindJSON(&input); err != nil {
		con.ErrorResponseWithUUID(c, http.StatusBadRequest, err, "Invalid JSON input")
		return
	}

	// Create a new User instance with only the Uuid field
	newUser := app.User{
		Id:               input.Uuid,
		DiscordSnowflake: input.DiscordSnowflake,
		FetchData:        true,
	}

	// Fetch the active profile UUID from the hypixel api
	activeProfileUUID, err := utils.FetchActivePlayerProfile(newUser.Id)
	if err != nil {
		con.ErrorResponseWithUUID(c, http.StatusInternalServerError, err, "Failed to fetch active profile UUID")
		return
	}
	newUser.ActiveProfileUUID = activeProfileUUID

	// Insert the new user into the database
	_, err = con.AppData.Engine.Insert(&newUser)
	if err != nil {
		con.ErrorResponseWithUUID(c, http.StatusInternalServerError, err, "Failed to insert user")
		return
	}

	// Insert player data for the new user
	utils.InsertPlayerData(con.AppData.Engine, []app.User{newUser})

	// Add the new user to the Users slice
	con.AppData.Users = append(con.AppData.Users, newUser)

	// Return success response
	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"user id": newUser.Id,
	})
}

func (con *Controller) GetUserData(c *gin.Context) {
	session := con.AppData.Engine.NewSession()
	defer session.Close()

	id := c.Query("id")

	if id == "" {
		con.ErrorResponseWithUUID(c, http.StatusBadRequest, nil, "Missing user ID")
		return
	}

	if err := session.Begin(); err != nil {
		log.Printf("Failed to start transaction: %v", err)
		return
	}

	user := app.User{Id: id}

	// Checking if guild exits
	has, err := session.Get(&user)
	if err != nil {
		con.ErrorResponseWithUUID(c, http.StatusInternalServerError, err, "Failed to fetch user")
		session.Rollback()
		return
	}

	if !has {
		con.ErrorResponseWithUUID(c, http.StatusNotFound, nil, fmt.Sprintf("User with ID %s not found", id))
		session.Rollback()
		return
	}

	playerDianaData := app.DianaData{UserId: user.Id}
	playerDungeonsData := app.DungeonsData{UserId: user.Id}

	_, err = con.AppData.Engine.OrderBy("fetch_time").Get(&playerDianaData)
	if err != nil {
		con.ErrorResponseWithUUID(c, http.StatusInternalServerError, err, "Failed to query player data")
		return
	}

	_, err = con.AppData.Engine.OrderBy("fetch_time").Get(&playerDungeonsData)
	if err != nil {
		con.ErrorResponseWithUUID(c, http.StatusInternalServerError, err, "Failed to query player data")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user":         user,
		"dianaData":    playerDianaData,
		"dungeonsData": playerDungeonsData,
	})
}
