package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// NewMySQLConnection establece una conexión con la base de datos MySQL.
func NewMySQLConnection() (*sql.DB, error) {
	// Leer variables de entorno
	dbUser := os.Getenv("DB_USER")
	//dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// Crear cadena de conexión
	dsn := fmt.Sprintf("%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbHost, dbPort, dbName)

	// Conectar a la base de datos
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error al conectar a MySQL: %v", err)
		return nil, err
	}

	// Verificar conexión
	if err = db.Ping(); err != nil {
		log.Fatalf("No se pudo conectar a MySQL: %v", err)
		return nil, err
	}

	log.Println("Conexión exitosa a MySQL")
	return db, nil
}
