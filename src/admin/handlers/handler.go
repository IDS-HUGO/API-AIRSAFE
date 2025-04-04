package handlers

import (
	"apiusersafe/src/admin/application"
	"apiusersafe/src/admin/domain"
	authdomain "apiusersafe/src/auth/domain" // Add with alias to avoid conflict
	"apiusersafe/src/auth/infrastructure"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type AdminHandler struct {
	CreateService *application.CreateAdminService
	UpdateService *application.UpdateAdminService
	DeleteService *application.DeleteAdminService
	ListService   *application.ListAdminsService
	userRepo      *infrastructure.MySQLUserRepository
}

func NewAdminHandler(
	createService *application.CreateAdminService,
	updateService *application.UpdateAdminService,
	deleteService *application.DeleteAdminService,
	listService *application.ListAdminsService,
	userRepo *infrastructure.MySQLUserRepository,
) *AdminHandler {
	return &AdminHandler{
		CreateService: createService,
		UpdateService: updateService,
		DeleteService: deleteService,
		ListService:   listService,
		userRepo:      userRepo,
	}
}

func (h *AdminHandler) CreateAdmin(w http.ResponseWriter, r *http.Request) {
	var admin domain.Admin
	if err := json.NewDecoder(r.Body).Decode(&admin); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.CreateService.Execute(&admin); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := h.userRepo.CreateUser(&authdomain.User{
		Usuario:    admin.Usuario,
		Contrasena: admin.Contrasena,
		Role:       "admin",
	}); err != nil {
		http.Error(w, "Error creating user entry", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *AdminHandler) ListAdmins(w http.ResponseWriter, r *http.Request) {
	admins, err := h.ListService.Execute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(admins)
}

func (h *AdminHandler) UpdateAdmin(w http.ResponseWriter, r *http.Request) {

    idStr := mux.Vars(r)["id"]
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "ID inválido", http.StatusBadRequest)
        return
    }

    var admin domain.Admin
    if err := json.NewDecoder(r.Body).Decode(&admin); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    admin.ID = id

    if err := h.UpdateService.Execute(&admin); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
}


func (h *AdminHandler) DeleteAdmin(w http.ResponseWriter, r *http.Request) {

    idStr := mux.Vars(r)["id"]
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