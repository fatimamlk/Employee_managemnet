package main

import (
	"EmployeeManagement/backend/config"
	"EmployeeManagement/backend/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()
	router := gin.Default()
	routes.EmployeeRoutes(router)
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
