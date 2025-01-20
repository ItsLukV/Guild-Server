package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ItsLukV/Guild-Server/src/app"
	"github.com/gin-gonic/gin"
)

func (con *Controller) GetDungeonsData(c *gin.Context) {
	con.ErrorResponseWithUUID(c, http.StatusNotImplemented, fmt.Errorf("not implemented"), "This endpoint is not implemented yet")
	return

	name := c.Param("user")

	var dungeonData app.DungeonsData

	// Construct the query
	has, err := con.AppData.Engine.
		Join("INNER", "users", "dungeons_data.user_id = users.id").
		Where("LOWER(users.name) = LOWER(?)", name).
		OrderBy("dungeons_data.fetch_time DESC").
		Get(&dungeonData)

	// Handle errors
	if err != nil {
		log.Print(err)
	}
	if !has {
		log.Print(fmt.Errorf("no dungeon data found for name: %s", name))
	}

	c.JSON(http.StatusOK, gin.H{
		"user": name,
		"data": dungeonData,
	})
}
