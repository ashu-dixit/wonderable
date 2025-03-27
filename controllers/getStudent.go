package controllers

import (
	"net/http"
	"wonderable/models"
	"wonderable/config"

	"github.com/gin-gonic/gin"
)

func GetStudentsByTeacher(c *gin.Context) {
	userID, exists := c.Get("userID") // This is UserID, not TeacherID

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "UserID not found"})
		return
	}

	var teacher models.Teacher
	if err := config.DB.Where("user_id = ?", userID).Preload("Students").First(&teacher).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Teacher not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"students": teacher.Students})
}

func GetStudentsByParent(c *gin.Context) {
	userID, exists := c.Get("userID") // This is UserID

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "UserID not found"})
		return
	}

	var parent models.Parent
	if err := config.DB.Where("user_id = ?", userID).Preload("Students").First(&parent).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Parent not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"students": parent.Students})
}


