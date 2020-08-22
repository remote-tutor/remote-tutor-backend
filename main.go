package main

import (
	db "backend/database"
	"backend/routes"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	db.MigrateTables()
	routes.InitializeRoutes(e)

	e.Logger.Fatal(e.Start(":3000"))
}
