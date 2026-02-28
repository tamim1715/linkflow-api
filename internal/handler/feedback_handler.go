package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tamim447/internal/constants"
	"github.com/tamim447/internal/service"
)

type FeedbackHandler struct {
	Service *service.FeedbackService
}

func NewFeedbackHandler(s *service.FeedbackService) *FeedbackHandler {
	return &FeedbackHandler{
		Service: s,
	}
}

func (h *FeedbackHandler) Submit(c echo.Context) error {

	var req struct {
		Message string `json:"message"`
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest,
			map[string]string{constants.Error: constants.InvalidRequest})
	}

	userID := c.Get(constants.ContextUserID).(string)

	if err := h.Service.Submit(userID, req.Message); err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{constants.Error: err.Error()})
	}

	return c.JSON(http.StatusCreated,
		map[string]string{constants.Message: constants.FeedbackSubmitted})
}
