package models

import "gorm.io/gorm"

type Course struct {
	gorm.Model
	Code        string `gorm:"unique"` //course code
	Name        string
	Description string
	Credit      int // credit number of course(số tín chỉ của học phần)
}
