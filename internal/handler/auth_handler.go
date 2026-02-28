package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/tamim447/internal/constants"
	"github.com/tamim447/internal/service"
)

type AuthHandler struct {
	Auth *service.AuthService
}

func NewAuthHandler(auth *service.AuthService) *AuthHandler {
	return &AuthHandler{
		Auth: auth,
	}
}

func (h *AuthHandler) RequestLink(c echo.Context) error {

	var req struct {
		Email string `json:"email"`
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(400, map[string]string{constants.Error: constants.InvalidRequest})
	}

	if err := h.Auth.RequestMagicLink(req.Email); err != nil {
		return c.JSON(500, map[string]string{constants.Error: err.Error()})
	}

	return c.JSON(200, map[string]string{constants.Message: constants.MagicLinkSent})
}

func (h *AuthHandler) Verify(c echo.Context) error {

	token := c.QueryParam(constants.Token)

	jwtToken, err := h.Auth.Verify(token)
	if err != nil {
		return c.JSON(401, map[string]string{constants.Error: err.Error()})
	}

	return c.JSON(200, map[string]string{constants.Token: jwtToken})
}
