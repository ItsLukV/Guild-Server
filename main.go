package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/exp/rand"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")

	words := []string{"It's me mario", "Meow", "LukV", "LukV", "LukV", "LukV", "LukV"}

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": words[rand.Intn(len(words))],
		})
	})
	router.Run(":8080")
}
