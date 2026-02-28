package app

import (
	"github.com/tamim447/internal/constants"
	"github.com/tamim447/internal/handler"
	"github.com/tamim447/internal/middleware"
)

type RouterDependencies struct {
	AuthHandler     *handler.AuthHandler
	FeedbackHandler *handler.FeedbackHandler
	AuthMiddleware  *middleware.AuthMiddleware
}

func (s *Server) RegisterRoutes() {

	api := s.Echo.Group(constants.APIPrefix)

	auth := api.Group(constants.AuthPrefix)
	auth.POST(constants.RequestLinkPrefix,
		s.AuthHandler.RequestLink,
		middleware.AuthRateLimiter(),
	)
	auth.GET(constants.VerifyPrefix, s.AuthHandler.Verify)

	feedback := api.Group(constants.FeedbackPrefix)
	feedback.Use(s.AuthMiddleware.RequireJWT)
	feedback.POST("", s.FeedbackHandler.Submit)
}
