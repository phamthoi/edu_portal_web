package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Password string
	Role     string //"admin", "teacher", "student"
	FullName string
	Email    string
}

// Encrypt password when creating user
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

/*
BeforeCreate(): hook of GORM: it will auto before GORM INSERT INTO to create a new record.
bcrypt.GenerateFromPassword(): use brypt libraries to encrypt password.

Bcrypt: is a algorithm one way encryption (thuật toán mã hóa một chiều có chủ ý làm chậm: giúp chống lại tấn công brute-force)
khác với SHA-256 hay MD5 rất nhanh, hacker có thể băm hàng tỷ mật khẩu mỗi giây.
*/
