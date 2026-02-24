package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/tamim447/internal/service"
	"net/http"
	"strings"
)

type contextKey string

const UserIDKey contextKey = "userId"

//type AuthMiddleware struct {
//	Secret string
//}
//
//func (m *AuthMiddleware) RequireAuth(next http.HandlerFunc) http.HandlerFunc {
//
//	return func(w http.ResponseWriter, r *http.Request) {
//
//		authHeader := r.Header.Get("Authorization")
//		if authHeader == "" {
//			http.Error(w, "missing authorization header", http.StatusUnauthorized)
//			return
//		}
//
//		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
//
//		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
//			return []byte(m.Secret), nil
//		})
//
//		if err != nil || !token.Valid {
//			http.Error(w, "invalid token", http.StatusUnauthorized)
//			return
//		}
//
//		claims := token.Claims.(jwt.MapClaims)
//
//		userID := claims["userId"].(string)
//
//		ctx := context.WithValue(r.Context(), UserIDKey, userID)
//
//		next(w, r.WithContext(ctx))
//	}
//}

type AuthMiddleware struct {
	JWT *service.JWTService
}

func NewAuthMiddleware(jwt *service.JWTService) *AuthMiddleware {
	return &AuthMiddleware{JWT: jwt}
}

func (m *AuthMiddleware) RequireJWT(next echo.HandlerFunc) echo.HandlerFunc {

	return func(c echo.Context) error {

		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return c.JSON(http.StatusUnauthorized,
				map[string]string{"error": "missing token"})
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		userID, err := m.JWT.Validate(tokenString)
		if err != nil {
			return c.JSON(http.StatusUnauthorized,
				map[string]string{"error": "invalid token"})
		}

		c.Set("userID", userID)

		return next(c)
	}
}
