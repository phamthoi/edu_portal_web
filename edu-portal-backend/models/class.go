package models

import "gorm.io/gorm"

type Class struct {
	gorm.Model
	CourseID  uint   // exception key connect with study section.
	TeacherID uint   //exception key connect with teacher
	Semester  string // semester information,
	Year      int

	//relationship
	Course  Course `gorm:"constraint:OnUpdaet:CASCADE, OnDelete:SET NULL;"`
	Teacher User   `gorm:"constraint:OnUpdate:CASCADE, OnDelete: SET NULL;"`
}

//use gorm.Model to have some default ID, CreatedAt, UpdatedAt,...
// CourseID and Teacher ID use to make a exception key to connect with courses and users table.
// Course and Teacher will help you access to get information relate course and teacher through GORM easily.
