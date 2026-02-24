package handler

//type VerifyHandler struct {
//	Ver *service.VerifyResult
//}
//
//func (h *AuthHandler) Verify(w http.ResponseWriter, r *http.Request) {
//
//	token := r.URL.Query().Get("token")
//	if token == "" {
//		http.Error(w, "missing token", 400)
//		return
//	}
//
//	result, err := h.Auth.VerifyMagicLink(token)
//	if err != nil {
//
//		// Token errors are user-visible
//		if err == service.ErrTokenExpired {
//			http.Error(w, "token expired", 401)
//			return
//		}
//
//		if err == service.ErrTokenUsed {
//			http.Error(w, "token already used", 401)
//			return
//		}
//
//		http.Error(w, "invalid token", 401)
//		return
//	}
//
//	// Return JWT in JSON
//	json.NewEncoder(w).Encode(map[string]string{
//		"token": result.JWT,
//	})
//}
