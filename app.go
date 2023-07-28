package main

import (
	"context"
	"myauth/application/model"
	"myauth/application/service"
	"myauth/application/util"
)

// App struct
type App struct {
	ctx        context.Context
	appService *service.ApplicationService
}

func (a *App) ListAlgoritm() string {
	return model.NewMessage(true, util.LIST_ALGORITM).ToJSON()
}

// NewApp creates a new App application struct
func Build(service *service.ApplicationService) *App {
	return &App{
		appService: service,
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}
