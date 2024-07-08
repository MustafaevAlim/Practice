package api

import (
	_ "Practice/api"
	"Practice/internal/api/controllers"
	"Practice/internal/app"
	"Practice/internal/service"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func RouteController(a *app.App) {
	a.EchoServer.POST("/register", controllers.Registration)
	a.EchoServer.POST("/auth", controllers.Authorization)
	a.EchoServer.POST("/sales", controllers.NewSale, service.JWTMiddleware)
	a.EchoServer.GET("/pdfReport", controllers.GetPdfReport)
	a.EchoServer.GET("/swagger/*", echoSwagger.WrapHandler)
	a.EchoServer.GET("/welcome", controllers.Welcome, service.JWTMiddleware)
}
