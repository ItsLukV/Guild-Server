package controllers

import (
	"log"
	"net/http"

	"github.com/ItsLukV/Guild-Server/src/app"
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
	var newUser app.User

	// Bind JSON input to the User struct
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON input"})
		return
	}

	// Insert the new user into the database
	affected, err := con.AppData.Engine.Insert(&newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert user"})
		log.Println("Failed to insert user: ", err)
		return
	}

	// Add the new user to the Users slice
	con.AppData.Users = append(con.AppData.Users, newUser)

	// Return success response
	c.JSON(http.StatusCreated, gin.H{
		"message":  "User created successfully",
		"affected": affected,
		"id":       newUser.Id,
	})
}
