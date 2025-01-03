package controllers

import (
	"github.com/ItsLukV/Guild-Server/src/app"
)

type Controller struct {
	AppData *app.App
}

func NewController(appData *app.App) *Controller {
	return &Controller{AppData: appData}
}
