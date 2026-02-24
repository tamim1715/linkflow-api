package app

import (
	"github.com/tamim447/internal/constants"
	"github.com/tamim447/internal/handler"
	"github.com/tamim447/internal/middleware"
)

//func RegisterRoutes(auth *handler.AuthHandler) {
//
//	http.HandleFunc("/auth/request-link", auth.RequestLink)
//	http.HandleFunc("/auth/verify", auth.Verify)
//}

type RouterDependencies struct {
	AuthHandler     *handler.AuthHandler
	FeedbackHandler *handler.FeedbackHandler
	AuthMiddleware  *middleware.AuthMiddleware
}

//func NewRouter(dep RouterDependencies) *http.ServeMux {
//
//	mux := http.NewServeMux()
//
//	// Public routes
//	mux.HandleFunc("/auth/request-link", Method(http.MethodPost, dep.AuthHandler.RequestLink))
//	mux.HandleFunc(
//		"/auth/request-link",
//		Method(http.MethodPost, dep.AuthHandler.RequestLink),
//	)
//	mux.HandleFunc("/auth/verify", dep.AuthHandler.Verify)
//
//	// Protected routes
//	mux.HandleFunc(
//		"/feedback",
//		dep.AuthMiddleware.RequireAuth(dep.FeedbackHandler.Submit),
//	)
//
//	return mux
//}

func (s *Server) RegisterRoutes() {

	api := s.Echo.Group(constants.APIPrefix)

	auth := api.Group(constants.AuthPrefix)
	auth.POST(constants.RequestLinkPrefix, s.AuthHandler.RequestLink)
	auth.GET(constants.VerifyPrefix, s.AuthHandler.Verify)

	feedback := api.Group(constants.FeedbackPrefix)
	feedback.Use(s.AuthMiddleware.RequireJWT)
	feedback.POST("", s.FeedbackHandler.Submit)
}
