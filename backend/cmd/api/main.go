package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/sebhv/sys-his-clin/backend/internal/domain"
	"github.com/sebhv/sys-his-clin/backend/internal/platform/db"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("Advertencia: No se pudo cargar el archivo .env")
	}

	database := db.DBConnection()

	if err := database.AutoMigrate(
		&domain.Paciente{},
	); err != nil {
		log.Fatal("Error en migración:", err)
	}

	// Levantar servidor
	log.Println("Servidor corriendo en :8080")
	http.ListenAndServe(":8080", nil)

}
