package app

import (
	_ "Practice/api"
	"Practice/internal/repository"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type App struct {
	EchoServer *echo.Echo
	DB         *gorm.DB
}

func NewApp() (*App, error) {
	app := &App{}
	app.EchoServer = echo.New()
	app.DB = repository.InitDB()
	return app, nil
}

func (a *App) Run() error {
	a.EchoServer.Use(middleware.Logger())
	a.EchoServer.Use(middleware.Recover())
	if err := a.EchoServer.Start(":8080"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
	return nil
}
