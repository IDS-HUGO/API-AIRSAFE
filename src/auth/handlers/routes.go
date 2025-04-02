package handlers

import "github.com/gorilla/mux"

func RegisterAuthRoutes(r *mux.Router, handler *AuthHandler) {
	r.HandleFunc("/auth/login", handler.Login).Methods("POST", "OPTIONS")
}
