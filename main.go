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

	r := mux.NewRouter()

	// Middleware de CORS
	r.Use(corsMiddleware)
	r.Use(loggingMiddleware)
	
	// Tus rutas
	

	r.HandleFunc("/auth/login", authHandlerInstance.Login).Methods("POST", "OPTIONS")

	r.HandleFunc("/admin", adminHandlerInstance.ListAdmins).Methods("GET")
	r.HandleFunc("/admin/create", adminHandlerInstance.CreateAdmin).Methods("POST")
	r.HandleFunc("/admin/{id}", adminHandlerInstance.UpdateAdmin).Methods("PUT")
	r.HandleFunc("/admin/{id}", adminHandlerInstance.DeleteAdmin).Methods("DELETE")

	r.HandleFunc("/comprador", compradorHandlerInstance.ListCompradores).Methods("GET")
	r.HandleFunc("/comprador/create", compradorHandlerInstance.CreateComprador).Methods("POST", "OPTIONS")
	r.HandleFunc("/comprador/{id}", compradorHandlerInstance.UpdateComprador).Methods("PUT")
	r.HandleFunc("/comprador/{id}", compradorHandlerInstance.DeleteComprador).Methods("DELETE")

	serverAddr := fmt.Sprintf(":%s", cfg.ServerPort)
	log.Printf("Servidor corriendo en http://localhost%s\n", serverAddr)

	server := &http.Server{
		Addr:    serverAddr,
		Handler: r,
	}

	log.Fatal(server.ListenAndServe())
}

func corsMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Verifica si la respuesta tiene las cabeceras necesarias
        w.Header().Set("Access-Control-Allow-Origin", "*")  // Permitir todos los orígenes
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
        w.Header().Set("Access-Control-Allow-Credentials", "true")

        if r.Method == "OPTIONS" {
            w.WriteHeader(http.StatusOK)
            return
        }

        next.ServeHTTP(w, r)
    })
}


func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("📡 %s %s %s", r.RemoteAddr, r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}
