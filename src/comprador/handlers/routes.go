package handler

import "github.com/gorilla/mux"

func RegisterCompradorRoutes(r *mux.Router, handler *CompradorHandler) {
	r.HandleFunc("/comprador", handler.ListCompradores).Methods("GET")
	r.HandleFunc("/comprador", handler.CreateComprador).Methods("POST")
	r.HandleFunc("/comprador/{id}", handler.UpdateComprador).Methods("PUT")
	r.HandleFunc("/comprador/{id}", handler.DeleteComprador).Methods("DELETE")
}
