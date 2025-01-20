package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/ItsLukV/Guild-Server/src/app"
	"github.com/gin-gonic/gin"
)

func (con *Controller) GetDiana(c *gin.Context) {
	con.ErrorResponseWithUUID(c, http.StatusNotImplemented, fmt.Errorf("not implemented"), "This endpoint is not implemented yet")
	return

	user := c.Param("user")

	var topEntries []app.DianaData

	err := con.AppData.Engine.Where("user_id = ?", user).Limit(10, 0).Find(&topEntries)
	if err != nil {
		log.Println("Failed to fetch top 10 entries: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to fetch data",
		})
		return
	}

	// After fetching DianaData, load the associated User
	for i, entry := range topEntries {
		var user app.User
		has, err := con.AppData.Engine.ID(entry.UserId).Get(&user)
		if err != nil {
			log.Println("Failed to load user data: ", err)
			continue
		}
		if has {
			topEntries[i].UserId = user.Id
		}
	}
	i, _ := strconv.Atoi(user)
	name := (con.AppData.Users)[i].Id

	c.JSON(http.StatusOK, gin.H{
		"user":    name,
		"entries": topEntries,
	})
}
