package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"unique;not null"`
	Password string `json:"password"`
	Role     string `gorm:"not null"` // "Teacher", "Parent", "Student"
}
