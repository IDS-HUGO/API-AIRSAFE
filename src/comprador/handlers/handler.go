package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	authdomain "apiusersafe/src/auth/domain"
	"apiusersafe/src/auth/infrastructure"
	"apiusersafe/src/comprador/application"
	"apiusersafe/src/comprador/domain"
)

type CompradorHandler struct {
	CreateService *application.CreateCompradorService
	UpdateService *application.UpdateCompradorService
	DeleteService *application.DeleteCompradorService
	ListService   *application.ListCompradoresService
	userRepo      *infrastructure.MySQLUserRepository
}

func NewCompradorHandler(
	createService *application.CreateCompradorService,
	updateService *application.UpdateCompradorService,
	deleteService *application.DeleteCompradorService,
	listService *application.ListCompradoresService,
	userRepo *infrastructure.MySQLUserRepository,
) *CompradorHandler {
	return &CompradorHandler{
		CreateService: createService,
		UpdateService: updateService,
		DeleteService: deleteService,
		ListService:   listService,
		userRepo:      userRepo,
	}
}

func (h *CompradorHandler) CreateComprador(w http.ResponseWriter, r *http.Request) {
	var comprador domain.Comprador
	if err := json.NewDecoder(r.Body).Decode(&comprador); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.CreateService.Execute(&comprador); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := h.userRepo.CreateUser(&authdomain.User{
		Usuario:    comprador.Usuario,
		Contrasena: comprador.Contrasena,
		Role:       "comprador",
	}); err != nil {
		http.Error(w, "Error creating user entry", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *CompradorHandler) ListCompradores(w http.ResponseWriter, r *http.Request) {
	compradores, err := h.ListService.Execute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(compradores)
}

func (h *CompradorHandler) UpdateComprador(w http.ResponseWriter, r *http.Request) {
	var comprador domain.Comprador
	if err := json.NewDecoder(r.Body).Decode(&comprador); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.UpdateService.Execute(&comprador); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *CompradorHandler) DeleteComprador(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	if err := h.DeleteService.Execute(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
