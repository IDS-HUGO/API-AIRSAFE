package main

import (
	"fmt"
	"log"
	"net/http"

	"apiusersafe/config"
	"apiusersafe/src/adapters/db"
	adminApp "apiusersafe/src/admin/application"
	adminHandler "apiusersafe/src/admin/handlers"
	adminInfra "apiusersafe/src/admin/infrastructure"
	authApp "apiusersafe/src/auth/application"
	authHandler "apiusersafe/src/auth/handlers"
	authInfra "apiusersafe/src/auth/infrastructure"
	compApp "apiusersafe/src/comprador/application"
	compHandler "apiusersafe/src/comprador/handlers"
	compInfra "apiusersafe/src/comprador/infrastructure"

	"github.com/gorilla/mux"
)

func main() {
	cfg := config.LoadConfig()

	database, err := db.NewMySQLConnection()
	if err != nil {
		log.Fatalf("Error al conectar a la base de datos: %v", err)
	}
	defer database.Close()

	// Repositorios
	adminRepo := adminInfra.NewAdminRepositoryMysql(database)
	compradorRepo := compInfra.NewMySQLCompradorRepository(database)
	userRepo := authInfra.NewMySQLUserRepository(database)

	// JWT Service
	jwtService := authInfra.NewJWTService(cfg.JWTSecret)

	// Auth Service
	loginService := authApp.NewLoginService(userRepo, jwtService)

	// Servicios de Administrador
	createAdminService := &adminApp.CreateAdminService{Repo: adminRepo}
	updateAdminService := &adminApp.UpdateAdminService{Repo: adminRepo}
	deleteAdminService := &adminApp.DeleteAdminService{Repo: adminRepo}
	listAdminsService := &adminApp.ListAdminsService{Repo: adminRepo}

	// Servicios de Comprador
	createCompradorService := compApp.NewCreateCompradorService(compradorRepo)
	updateCompradorService := compApp.NewUpdateCompradorService(compradorRepo)
	deleteCompradorService := compApp.NewDeleteCompradorService(compradorRepo)
	listCompradoresService := compApp.NewListCompradoresService(compradorRepo)

	// Handlers
	authHandlerInstance := authHandler.NewAuthHandler(loginService)

	adminHandlerInstance := adminHandler.NewAdminHandler(
		createAdminService,
		updateAdminService,
		deleteAdminService,
		listAdminsService,
		userRepo,
	)

	compradorHandlerInstance := compHandler.NewCompradorHandler(
		createCompradorService,
		updateCompradorService,
		deleteCompradorService,
		listCompradoresService,
		userRepo,
	)

	// Crear el enrutador HTTP
	r := mux.NewRouter()

	// Middleware para CORS y logging
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}

			next.ServeHTTP(w, r)
		})
	})
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
			next.ServeHTTP(w, r)
		})
	})

	// Auth routes
	r.HandleFunc("/auth/login", authHandlerInstance.Login).Methods("POST", "OPTIONS")

	// Admin routes
	r.HandleFunc("/admin", adminHandlerInstance.ListAdmins).Methods("GET")
	r.HandleFunc("/admin/create", adminHandlerInstance.CreateAdmin).Methods("POST")
	r.HandleFunc("/admin/update", adminHandlerInstance.UpdateAdmin).Methods("PUT")
	r.HandleFunc("/admin/delete", adminHandlerInstance.DeleteAdmin).Methods("DELETE")

	// Comprador routes
	r.HandleFunc("/comprador", compradorHandlerInstance.ListCompradores).Methods("GET")
	r.HandleFunc("/comprador/create", compradorHandlerInstance.CreateComprador).Methods("POST")
	r.HandleFunc("/comprador/update", compradorHandlerInstance.UpdateComprador).Methods("PUT")
	r.HandleFunc("/comprador/delete", compradorHandlerInstance.DeleteComprador).Methods("DELETE")

	serverAddr := fmt.Sprintf(":%s", cfg.ServerPort)
	log.Printf("Servidor corriendo en http://localhost%s\n", serverAddr)

	server := &http.Server{
		Addr:    serverAddr,
		Handler: r,
	}

	log.Fatal(server.ListenAndServe())
}
