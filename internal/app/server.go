package app

import (
	"github.com/labstack/echo/v4"
	"github.com/tamim447/internal/constants"
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

	tokenGen := service.NewTokenService()
	emailSender := email.NewMockSender()
	slackMockClient := slack.NewMockClient()

	// Services
	jwtService := service.NewJWTService(constants.SecretKey)

	authService := service.NewAuthService(
		userRepo,
		tokenRepo,
		tokenGen,
		jwtService,
		emailSender,
	)

	feedbackService := service.NewFeedbackService(
		feedbackRepo,
		slackMockClient,
	)

	// Handlers
	authHandler := handler.NewAuthHandler(authService)
	feedbackHandler := handler.NewFeedbackHandler(feedbackService)

	// Middleware
	authMiddleware := middleware.NewAuthMiddleware(jwtService)

	return &Server{
		Echo:            e,
		AuthHandler:     authHandler,
		FeedbackHandler: feedbackHandler,
		AuthMiddleware:  authMiddleware,
	}
}
