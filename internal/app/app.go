package app

import (
	"github.com/FelixMorenoT/go-rampup/internal/http/controller/ping"
	"github.com/gin-gonic/gin"
)

type App struct {
	Router *gin.Engine
}

func NewApp() *App {
	app := &App{
		Router: gin.Default(),
	}

	app.Routes()

	return app
}

func (app *App) Start(port string) {

	if err := app.Router.Run(":" + port); err != nil {
		panic("Error Running Server")
	}
}

func (app *App) Routes() {
	pingController := ping.NewController()

	app.Router.GET("/ping", pingController.Ping)
}
