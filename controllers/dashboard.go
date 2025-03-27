package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Parent Dashboard (Only accessible by Parents)
func ParentDashboard(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Welcome to the Parent Dashboard"})
}

// Teacher Dashboard (Only accessible by Teachers)
func TeacherDashboard(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Welcome to the Teacher Dashboard"})
}
