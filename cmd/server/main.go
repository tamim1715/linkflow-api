package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/tamim447/internal/app"
	"github.com/tamim447/internal/config"
	"github.com/tamim447/internal/database"
)

func main() {

	config.LoadEnv()

	// Database Connection
	db, err := database.Connect(config.MongoURI, config.MongoDatabase)
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	// Create Echo instance
	e := echo.New()

	// Global Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Create Server with dependencies
	server := app.NewServer(e, db)

	// Register routes
	server.RegisterRoutes()

	// Start server
	e.Logger.Fatal(e.Start(":" + config.ServerPort))
}
