package app

import (
	"github.com/labstack/echo/v4"
	"github.com/tamim447/internal/email"
	"github.com/tamim447/internal/handler"
	"github.com/tamim447/internal/middleware"
	"github.com/tamim447/internal/repository/mongodb"
	"github.com/tamim447/internal/service"
	"github.com/tamim447/internal/slack"
	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	Echo            *echo.Echo
	AuthHandler     *handler.AuthHandler
	FeedbackHandler *handler.FeedbackHandler
	AuthMiddleware  *middleware.AuthMiddleware
}

func NewServer(e *echo.Echo, db *mongo.Database) *Server {

	// Repositories
	tokenRepo := mongodb.NewMongoTokenRepository(db)
	userRepo := mongodb.NewMongoUserRepository(db)
	feedbackRepo := mongodb.NewMongoFeedbackRepository(db)

	// Services
	jwtService := service.NewJWTService("super-secret-key")

	authService := &service.AuthService{
		Users:    userRepo,
		Tokens:   tokenRepo,
		TokenGen: &service.TokenService{},
		JWT:      jwtService,
		Email:    &email.MockSender{},
	}

	feedbackService := &service.FeedbackService{
		Repo:  feedbackRepo,
		Slack: &slack.MockClient{},
	}

	//authService := service.AuthService(userRepo, tokenRepo, jwtService)
	//feedbackService := service.NewFeedbackService(feedbackRepo)

	// Handlers
	authHandler := &handler.AuthHandler{
		Auth: authService,
	}

	feedbackHandler := &handler.FeedbackHandler{
		Service: feedbackService,
	}
	//authHandler := handler.NewAuthHandler(authService)
	//feedbackHandler := handler.NewFeedbackHandler(feedbackService)

	// Middleware
	//authMiddleware := &middleware.AuthMiddleware{
	//	Secret: "super-secret-key",
	//}

	authMiddleware := middleware.NewAuthMiddleware(jwtService)

	return &Server{
		Echo:            e,
		AuthHandler:     authHandler,
		FeedbackHandler: feedbackHandler,
		AuthMiddleware:  authMiddleware,
	}
}
