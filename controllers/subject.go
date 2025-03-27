package controllers

import (
	"net/http"
	"wonderable/config"
	"wonderable/models"

	"github.com/gin-gonic/gin"
)

// AddSubject - API to add a new subject
func AddSubject(c *gin.Context) {
	var input struct {
		Name string `json:"name" binding:"required"`
	}

	// Bind input data
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if subject already exists
	var existingSubject models.Subject
	if err := config.DB.Where("name = ?", input.Name).First(&existingSubject).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Subject already exists"})
		return
	}

	// Create new subject
	newSubject := models.Subject{Name: input.Name}
	if err := config.DB.Create(&newSubject).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add subject"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Subject added successfully", "subject": newSubject})
}


func GetSubjectsByStudentID(c *gin.Context) {
	studentID := c.Param("studentID") // Get Student ID from URL parameter

	var student models.Student

	// Fetch student and preload subjects
	if err := config.DB.Preload("Subjects").First(&student, studentID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}

	// Return subjects
	c.JSON(http.StatusOK, gin.H{"student_id": student.ID, "subjects": student.Subjects})
}