package main

import (
	_ "Practice/api"
	"Practice/internal/api"
	"Practice/internal/app"
)

func main() {
	a, err := app.NewApp()
	if err != nil {
		panic(err)
	}
	api.RouteController(a)

	err = a.Run()
	if err != nil {
		panic(err)
	}

}
