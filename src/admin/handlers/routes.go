package handlers

import "github.com/gorilla/mux"

func RegisterAdminRoutes(r *mux.Router, handler *AdminHandler) {
	r.HandleFunc("/admin", handler.ListAdmins).Methods("GET")
	r.HandleFunc("/admin", handler.CreateAdmin).Methods("POST")
	r.HandleFunc("/admin/{id}", handler.UpdateAdmin).Methods("PUT")
	r.HandleFunc("/admin/{id}", handler.DeleteAdmin).Methods("DELETE")
}
