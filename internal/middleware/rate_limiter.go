package middleware

import (
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/tamim447/internal/constants"
)

type client struct {
	count     int
	lastReset time.Time
}

var (
	clients = make(map[string]*client)
	mu      sync.Mutex
)

func AuthRateLimiter() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			ip, _, _ := net.SplitHostPort(c.Request().RemoteAddr)
			if ip == "" {
				ip = c.RealIP()
			}

			mu.Lock()
			defer mu.Unlock()

			cl, exists := clients[ip]

			if !exists {
				clients[ip] = &client{
					count:     1,
					lastReset: time.Now(),
				}
				return next(c)
			}

			if time.Since(cl.lastReset) > constants.AuthRateLimitWindow {
				cl.count = 1
				cl.lastReset = time.Now()
				return next(c)
			}

			if cl.count >= constants.AuthRateLimitRequests {
				return c.JSON(http.StatusTooManyRequests, map[string]string{
					constants.Error: constants.TooManyRequests,
				})
			}

			cl.count++
			return next(c)
		}
	}
}
