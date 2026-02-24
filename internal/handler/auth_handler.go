package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/tamim447/internal/constants"
	"github.com/tamim447/internal/service"
)

type AuthHandler struct {
	Auth *service.AuthService
}

//func (h *AuthHandler) RequestLink(w http.ResponseWriter, r *http.Request) {
//
//	var body struct {
//		Email string `json:"email"`
//	}
//
//	json.NewDecoder(r.Body).Decode(&body)
//
//	err := h.Auth.RequestMagicLink(body.Email)
//	if err != nil {
//		http.Error(w, "failed to send link", 500)
//		return
//	}
//
//	w.WriteHeader(http.StatusOK)
//}

func (h *AuthHandler) RequestLink(c echo.Context) error {

	var req struct {
		Email string `json:"email"`
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(400, map[string]string{"error": "invalid request"})
	}

	if err := h.Auth.RequestMagicLink(req.Email); err != nil {
		return c.JSON(500, map[string]string{"error": err.Error()})
	}

	return c.JSON(200, map[string]string{"message": "magic link sent"})
}

func (h *AuthHandler) Verify(c echo.Context) error {

	token := c.QueryParam(constants.Token)

	jwtToken, err := h.Auth.Verify(token)
	if err != nil {
		return c.JSON(401, map[string]string{"error": err.Error()})
	}

	return c.JSON(200, map[string]string{constants.Token: jwtToken})
}
