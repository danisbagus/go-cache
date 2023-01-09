package main

import (
	"os"

	"github.com/danisbagus/go-cache/routes"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	godotenv.Load()

	e := echo.New()
	routes.ApiRoutes(e)

	appPort := os.Getenv("PORT")
	e.Logger.Fatal(e.Start(":" + appPort))
}
