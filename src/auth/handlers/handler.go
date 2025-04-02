package handlers

import (
	"apiusersafe/src/auth/application"
	"encoding/json"
	"net/http"
)

type LoginRequest struct {
	Usuario    string `json:"usuario"`
	Contrasena string `json:"contrasena"`
}

type AuthHandler struct {
	loginService *application.LoginService
}

func NewAuthHandler(loginService *application.LoginService) *AuthHandler {
	return &AuthHandler{loginService: loginService}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response, err := h.loginService.Execute(req.Usuario, req.Contrasena)
	if err != nil {
		http.Error(w, "Credenciales inválidas", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
