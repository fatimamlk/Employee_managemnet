package controllers

import (
	"EmployeeManagement/backend/config"
	"EmployeeManagement/backend/models"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetEmployees(c *gin.Context) {
	var employees []models.Employee
	collection := config.DB.Collection("employees")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching employees"})
		return
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var employee models.Employee
		cursor.Decode(&employee)
		employees = append(employees, employee)
	}

	c.JSON(http.StatusOK, employees)
}

func CreateEmployee(c *gin.Context) {
	var employee models.Employee
	if err := c.BindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	employee.ID = primitive.NewObjectID()
	collection := config.DB.Collection("employees")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, employee)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating employee"})
		return
	}

	c.JSON(http.StatusCreated, employee)
}

// GetEmployeeByID - Récupérer un employé par ID
func GetEmployeeByID(c *gin.Context) {
	// Récupérer l'ID passé en paramètre dans l'URL
	idParam := c.Param("id")
	employeeID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employee ID"})
		return
	}

	// Rechercher l'employé dans la base de données
	var employee models.Employee
	collection := config.DB.Collection("employees")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Rechercher l'employé par son ID
	err = collection.FindOne(ctx, bson.M{"_id": employeeID}).Decode(&employee)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching employee"})
		}
		return
	}

	// Retourner l'employé trouvé en réponse
	c.JSON(http.StatusOK, employee)
}
func UpdateEmployee(c *gin.Context) {
	// Récupérer l'ID passé en paramètre dans l'URL
	idParam := c.Param("id")
	employeeID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employee ID"})
		return
	}

	// Définir une variable pour l'employé mis à jour
	var updatedEmployee models.Employee
	if err := c.ShouldBindJSON(&updatedEmployee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employee data"})
		return
	}

	// Mettre à jour les informations de l'employé
	collection := config.DB.Collection("employees")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Rechercher l'employé dans la base de données et mettre à jour
	updateResult, err := collection.UpdateOne(ctx, bson.M{"_id": employeeID}, bson.M{
		"$set": bson.M{
			"first_name":   updatedEmployee.FirstName,
			"last_name":    updatedEmployee.LastName,
			"email":        updatedEmployee.Email,
			"phone_number": updatedEmployee.PhoneNumber,
			"position":     updatedEmployee.Position,
			"department":   updatedEmployee.Department,
			"date_of_hire": updatedEmployee.DateOfHire,
		},
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update employee"})
		return
	}

	// Vérifier si un document a été mis à jour
	if updateResult.MatchedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		return
	}

	// Retourner un message de succès
	c.JSON(http.StatusOK, gin.H{"message": "Employee updated successfully"})
}
func DeleteEmployee(c *gin.Context) {
	// Récupérer l'ID passé en paramètre dans l'URL
	idParam := c.Param("id")
	employeeID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employee ID"})
		return
	}

	// Supprimer l'employé de la base de données
	collection := config.DB.Collection("employees")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	deleteResult, err := collection.DeleteOne(ctx, bson.M{"_id": employeeID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete employee"})
		return
	}

	// Vérifier si un document a été supprimé
	if deleteResult.DeletedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		return
	}

	// Retourner un message de succès
	c.JSON(http.StatusOK, gin.H{"message": "Employee deleted successfully"})
}
