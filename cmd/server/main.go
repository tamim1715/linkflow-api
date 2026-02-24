package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/tamim447/internal/app"
	"github.com/tamim447/internal/config"
	"github.com/tamim447/internal/database"
)

//func main() {
//
//	// 1️⃣ Database
//	db, err := database.Connect("mongodb://localhost:27017", "linkflow")
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// 2️⃣ Repositories
//	userRepo := &mongoRepo.UserRepo{
//		Collection: db.Collection("users"),
//	}
//
//	tokenRepo := &mongoRepo.TokenRepo{
//		Collection: db.Collection("magic_tokens"),
//	}
//
//	feedbackRepo := &mongoRepo.FeedbackRepo{
//		Collection: db.Collection("feedbacks"),
//	}
//
//	authService := &service.AuthService{
//		Users:    userRepo,
//		Tokens:   tokenRepo,
//		TokenGen: &service.TokenService{},
//		JWT:      &service.JWTService{Secret: "supersecret"},
//		Email:    &email.MockSender{},
//	}
//
//	feedbackService := &service.FeedbackService{
//		Repo:  feedbackRepo,
//		Slack: &slack.MockClient{},
//	}
//
//	authHandler := &handler.AuthHandler{
//		Auth: authService,
//	}
//
//	feedbackHandler := &handler.FeedbackHandler{
//		Service: feedbackService,
//	}
//
//	// Middleware
//	authMiddleware := &middleware.AuthMiddleware{
//		Secret: "supersecret",
//	}
//
//	router := app.NewRouter(app.RouterDependencies{
//		AuthHandler:     authHandler,
//		FeedbackHandler: feedbackHandler,
//		AuthMiddleware:  authMiddleware,
//	})
//
//	log.Println("Server running on :8080")
//	http.ListenAndServe(":8080", router)
//}

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
