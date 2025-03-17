package controllers

import (
	"log"
	"net/http"
	"runtime"

	"github.com/ItsLukV/Guild-Server/src/model"
	"github.com/gin-gonic/gin"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"golang.org/x/exp/rand"
)

type Controller struct {
	AppData *model.App
}

func NewController(appData *model.App) *Controller {
	return &Controller{AppData: appData}
}

const (
	Red   = "\033[31m"
	Reset = "\033[0m"
)

func (c *Controller) ErrorResponseWithUUID(ctx *gin.Context, errorCode int, err error, message string) {
	alphabet := "abcdefghijklmnopqrstuvwxyz1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	errorID, errWithID := gonanoid.Generate(alphabet, 21)
	if errWithID != nil {
		log.Println("Failed to generate error ID: ", errWithID)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"errorID": errorID,
		})
		return
	}

	// Capture the current call stack
	stackBuf := make([]byte, 1024)
	stackSize := runtime.Stack(stackBuf, false)
	stackTrace := string(stackBuf[:stackSize])

	if err != nil {
		log.Printf("%sError ID: %v \nMessage: %s\nError: %v\nStack Trace:\n%s%s",
			Red, errorID, message, err, stackTrace, Reset)
	} else {
		log.Printf("%sError ID: %v \nMessage: %s\nStack Trace:\n%s%s",
			Red, errorID, message, stackTrace, Reset)
	}

	ctx.JSON(errorCode, gin.H{
		"errorID": errorID,
	})
}

func (con *Controller) GetDefault(c *gin.Context) {
	words := []string{"It's me mario", "Meow", "LukV", "LukV", "LukV", "LukV", "LukV"}

	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": words[rand.Intn(len(words))],
	})
}
