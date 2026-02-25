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

	cfg := config.Load()

	// 1️⃣ Database Connection
	db, err := database.Connect(cfg.MongoURI, cfg.MongoDatabase)
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	// 2️⃣ Create Echo instance
	e := echo.New()

	// 3️⃣ Global Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// 4️⃣ Create Server with dependencies
	server := app.NewServer(e, db)

	// 5️⃣ Register routes
	server.RegisterRoutes()

	// 6️⃣ Start server
	e.Logger.Fatal(e.Start(":8080"))
}
