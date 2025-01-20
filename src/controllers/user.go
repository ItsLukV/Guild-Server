package controllers

import (
	"net/http"

	"github.com/ItsLukV/Guild-Server/src/app"
	"github.com/ItsLukV/Guild-Server/src/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/rand"
)

func (con *Controller) GetDefault(c *gin.Context) {
	words := []string{"It's me mario", "Meow", "LukV", "LukV", "LukV", "LukV", "LukV"}

	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": words[rand.Intn(len(words))],
	})
}

func (con *Controller) GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"users": con.AppData.Users,
	})
}

func (con *Controller) PostUsers(c *gin.Context) {
	// Define a temporary struct for binding only the Uuid field
	var input struct {
		Uuid string `json:"uuid" binding:"required"`
	}

	// Bind JSON input to the input struct
	if err := c.ShouldBindJSON(&input); err != nil {
		con.ErrorResponseWithUUID(c, http.StatusBadRequest, err, "Invalid JSON input")
		return
	}

	// Create a new User instance with only the Uuid field
	newUser := app.User{
		Id:        input.Uuid,
		FetchData: true,
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

	// Add the new user to the Users slice
	con.AppData.Users = append(con.AppData.Users, newUser)

	// Return success response
	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"user id": newUser.Id,
	})
}
