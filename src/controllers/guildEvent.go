package controllers

import (
	"log"
	"net/http"

	"github.com/ItsLukV/Guild-Server/src/app"
	"github.com/gin-gonic/gin"
)

func (con *Controller) CreateGuildEvent(c *gin.Context) {
	var guildEvent app.GuildEvent

	// Bind JSON input to the GuildEvent struct
	if err := c.ShouldBindJSON(&guildEvent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON input"})
		return
	}

	// Insert the new guild event into the database
	affected, err := con.AppData.Engine.Insert(&guildEvent)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert guild event"})
		log.Println("Failed to insert guild event: ", err)
		return
	}

	// Return success response
	c.JSON(http.StatusCreated, gin.H{
		"message":  "Guild event created successfully",
		"affected": affected,
		"id":       guildEvent.Id,
	})
}
