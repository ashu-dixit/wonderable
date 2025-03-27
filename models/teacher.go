package models

import "gorm.io/gorm"

type Teacher struct {
	gorm.Model
	UserID   uint       `gorm:"unique;not null"` // Links to User table
	User     User       `gorm:"foreignKey:UserID"`
	Students []*Student `gorm:"many2many:teacher_students;" json:"students"`
}
