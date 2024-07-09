package api

import (
	_ "Practice/api"
	"Practice/internal/api/controllers"
	"Practice/internal/app"
	"Practice/internal/service"
	echoSwagger "github.com/swaggo/echo-swagger"
	"gorm.io/gorm"
)

func RouteController(a *app.App, db *gorm.DB) {
	DB := controllers.NewDatabase(db)
	a.EchoServer.POST("/register", DB.Registration)
	a.EchoServer.POST("/auth", DB.Authorization)
	a.EchoServer.POST("/sales", DB.NewSale, service.JWTMiddleware)
	a.EchoServer.DELETE("/sales/:id", DB.DelSale, service.JWTMiddleware)
	a.EchoServer.GET("/pdfReport", DB.PdfReport)
	a.EchoServer.GET("/JSONReport", DB.JSONReport)
	a.EchoServer.GET("/swagger/*", echoSwagger.WrapHandler)
	a.EchoServer.GET("/welcome", DB.Welcome, service.JWTMiddleware)
}
