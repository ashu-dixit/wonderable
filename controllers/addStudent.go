package controllers

import (
	"net/http"
	"wonderable/config"
	"wonderable/models"

	"github.com/gin-gonic/gin"
)

type AssignTeacherRequest struct {
	TeacherID uint `json:"teacher_id" binding:"required"`
	StudentID uint `json:"student_id" binding:"required"`
}

// AssignStudentToTeacher assigns a student to a teacher
func AssignStudentToTeacher(c *gin.Context) {
	var input AssignTeacherRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var teacher models.Teacher
	var student models.Student

	if err := config.DB.First(&teacher, input.TeacherID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Teacher not found"})
		return
	}

	if err := config.DB.First(&student, input.StudentID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}

	// Assign the student to the teacher
	config.DB.Model(&teacher).Association("Students").Append(&student)

	c.JSON(http.StatusOK, gin.H{"message": "Student assigned to teacher successfully"})
}


type AssignParentRequest struct {
	ParentID  uint `json:"parent_id" binding:"required"`
	StudentID uint `json:"student_id" binding:"required"`
}

// AssignStudentToParent assigns a student to a parent
func AssignStudentToParent(c *gin.Context) {
	var input AssignParentRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var parent models.Parent
	var student models.Student

	if err := config.DB.First(&parent, input.ParentID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Parent not found"})
		return
	}

	if err := config.DB.First(&student, input.StudentID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}

	// Assign the student to the parent
	config.DB.Model(&parent).Association("Students").Append(&student)

	c.JSON(http.StatusOK, gin.H{"message": "Student assigned to parent successfully"})
}

func AddStudent(c *gin.Context) {
	var studentInput struct {
		Name  string `json:"name" binding:"required"`
		Email string `json:"email" binding:"required,email"`
		SubjectIDs []uint `json:"subject_ids"` // List of subject IDs to assign
	}

	// Bind input data
	if err := c.ShouldBindJSON(&studentInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create User (for student)
	newUser := models.User{
		Email: studentInput.Email,
		Role:  "Student",
	}

	if err := config.DB.Create(&newUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}


	var subjects []*models.Subject
	if len(studentInput.SubjectIDs) > 0 {
		var dbSubjects []models.Subject
		if err := config.DB.Where("id IN ?", studentInput.SubjectIDs).Find(&dbSubjects).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch subjects"})
			return
		}

		for i := range dbSubjects {
			subjects = append(subjects, &dbSubjects[i])
		}
	}

	// Create Student linked to new User
	newStudent := models.Student{
		UserID: newUser.ID,
		User:   newUser,
		Name:   studentInput.Name,
		Email:  studentInput.Email,
		Subjects: subjects, // Assign subjects to the student

	}

	if err := config.DB.Create(&newStudent).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err, "Message": "Failed to add student"})
		return
	}

	// Return success response
	c.JSON(http.StatusCreated, gin.H{
		"message": "Student added successfully",
		"student": newStudent,
	})
}
