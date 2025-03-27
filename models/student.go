package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	UserID    uint        `gorm:"unique;not null"` // Links to User table
	User      User        `gorm:"foreignKey:UserID"`
	Name      string      `json:"name"`
	Email     string      `gorm:"unique;not null" json:"email"`
	Teachers  []*Teacher  `gorm:"many2many:teacher_students;" json:"teachers"`
	Parents   []*Parent   `gorm:"many2many:parent_students;" json:"parents"`
	Subjects  []*Subject  `gorm:"many2many:student_subjects;" json:"subjects"` // ✅ Relationship with subjects

}
