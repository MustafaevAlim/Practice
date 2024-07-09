package main

import (
	_ "Practice/api"
	"Practice/internal/api"
	"Practice/internal/app"
)

// @title Generate Geport
// @version 1.0
// @description Генерация отчета по продажам за месяц.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Введите JWT токен следующим образом: "Bearer {токен}"

// @host localhost:8080
// @BasePath /
func main() {
	a, err := app.NewApp()
	if err != nil {
		panic(err)
	}
	api.RouteController(a, a.DB)
	err = a.Run()
	if err != nil {
		panic(err)
	}

}
