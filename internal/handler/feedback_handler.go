package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/tamim447/internal/service"
	"net/http"
)

//type FeedbackHandler struct {
//	Service *service.FeedbackService
//}
//
//func (h *FeedbackHandler) Submit(w http.ResponseWriter, r *http.Request) {
//
//	userID := r.Context().Value(middleware.UserIDKey).(string)
//
//	var body struct {
//		Message string `json:"message"`
//	}
//
//	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
//		http.Error(w, "invalid body", 400)
//		return
//	}
//
//	if body.Message == "" {
//		http.Error(w, "message required", 400)
//		return
//	}
//
//	if err := h.Service.Submit(userID, body.Message); err != nil {
//		http.Error(w, "failed to save feedback", 500)
//		return
//	}
//
//	w.WriteHeader(http.StatusCreated)
//}

type FeedbackHandler struct {
	Service *service.FeedbackService
}

func NewFeedbackHandler(s *service.FeedbackService) *FeedbackHandler {
	return &FeedbackHandler{Service: s}
}

func (h *FeedbackHandler) Submit(c echo.Context) error {

	var req struct {
		Message string `json:"message"`
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest,
			map[string]string{"error": "invalid request"})
	}

	userID := c.Get("userID").(string)

	if err := h.Service.Submit(userID, req.Message); err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated,
		map[string]string{"message": "feedback submitted"})
}
