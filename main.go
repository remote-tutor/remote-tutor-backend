package main

import (
	db "backend/database"
	"backend/deployment"
	"backend/routes"
	"backend/utils"
	"github.com/labstack/echo"
	"os"
)

func main() {
	if os.Getenv("APP_ENV") == "development" {
		development()
	} else {
		deployment.Deployment()
	}
}

func development() {
	e := echo.New()

	utils.InitializeEnvironmentVariables()
	db.MigrateTables()
	routes.InitializeRoutes(e)

	e.Logger.Fatal(e.Start(":3000"))
}