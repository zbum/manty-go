package handler

import (
	"encoding/json"
	"manty-go/service"
	"manty-go/utils"
	"net/http"
)

type LoginHandler struct {
	UserService *service.UserService
}

func NewLoginHandler(userService *service.UserService) *LoginHandler {
	return &LoginHandler{UserService: userService}
}

func (h *LoginHandler) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	type loginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	var req loginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid JSON body"))
		return
	}

	user := h.UserService.Authenticate(req.Username, req.Password)
	if user == nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Login failed: invalid credentials"))
		return
	}

	// 로그인 성공: 세션 생성 및 쿠키 설정
	utils.SetSession(w, user.ID, user.Username)

	w.Write([]byte("Login successful! Welcome, " + user.Username))
}
