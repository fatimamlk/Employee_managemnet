package main

import (
	"EmployeeManagement/backend/config"
	"EmployeeManagement/backend/routes"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()

	// Initialisation du routeur Gin
	router := gin.Default()

	// Ajouter CORS middleware
	router.Use(cors.Default()) // Cela autorise tous les origines, vous pouvez personnaliser si nécessaire

	// Définir les routes des employés
	routes.EmployeeRoutes(router)

	// Démarrer le serveur sur le port 8080
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
