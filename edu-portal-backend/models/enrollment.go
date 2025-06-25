package models

import "gorm.io/gorm"

type Enrollment struct {
	gorm.Model
	StudentID uint     //user_id of student
	ClassID   uint     //registerd class
	Score     *float64 // score, it can null if teacher has not graded them yet.

	//relationship
	Student User `gorm:"foreignKey:StudentID"`
	Class   Class
}

/* StudentID: Exception Key pointer to User with role is student.
ClassID: exception key pointer to class table which the student learned.
*/
