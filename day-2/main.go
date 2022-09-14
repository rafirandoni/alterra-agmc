package main

import (
	"day-2/config"
	"day-2/lib/db"
	"day-2/models"
	"day-2/routes"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func main() {
	cfg := config.Load()

	db := db.NewDB()
	migrateDB(db)

	e := echo.New()

	routes.NewRouteV1(e)

	e.Start(":" + cfg.AppPort)
}

func migrateDB(*gorm.DB) {
	db.NewDB().AutoMigrate(&models.User{})
}
