package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/tamim447/internal/constants"
	"github.com/tamim447/internal/service"
	"net/http"
	"strings"
)

type contextKey string

const UserIDKey contextKey = "userId"

type AuthMiddleware struct {
	JWT *service.JWTService
}

func NewAuthMiddleware(jwt *service.JWTService) *AuthMiddleware {
	return &AuthMiddleware{JWT: jwt}
}

func (m *AuthMiddleware) RequireJWT(next echo.HandlerFunc) echo.HandlerFunc {

	return func(c echo.Context) error {

		authHeader := c.Request().Header.Get(constants.AuthorizationHeader)
		if authHeader == "" {
			return c.JSON(http.StatusUnauthorized,
				map[string]string{constants.Error: constants.ErrMissingToken})
		}

		tokenString := strings.TrimPrefix(authHeader, constants.BearerPrefix)

		userID, err := m.JWT.Validate(tokenString)
		if err != nil {
			return c.JSON(http.StatusUnauthorized,
				map[string]string{constants.Error: constants.ErrInvalidToken})
		}

		c.Set("userID", userID)

		return next(c)
	}
}
