package main

import (
	db "backend/database"
	"backend/routes"
	"backend/utils"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	utils.InitializeEnvironmentVariables()
	db.MigrateTables()
	routes.InitializeRoutes(e)

	e.Logger.Fatal(e.Start(":3000"))
}
