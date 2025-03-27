package models

import "gorm.io/gorm"

type Subject struct {
	gorm.Model
	Name string `gorm:"unique;not null" json:"name"`
	Students []*Student `gorm:"many2many:student_subjects;" json:"students"`
}
