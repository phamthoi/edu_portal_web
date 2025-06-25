package controllers

import (
	"edu-portal-backend/config"
	"edu-portal-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// this is just for admin
func isAdmin(c *gin.Context) bool {
	return c.GetString("role") == "admin"
}

// get list of used by role(lấy danh sách sử dụng theo vai trò)
func GetUser(c *gin.Context) {
	if !isAdmin(c) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	role := c.Query("role") //this can student or teacher role.
	var users []models.User
	config.DB.Where("role = ?", role).Find(&users)
	c.JSON(http.StatusOK, users)
}

// create user(teacher or student)
func CreateUser(c *gin.Context) {
	if !isAdmin(c) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Hashing error"})
		return
	}

	//hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "hashing error"})
		return
	}
	user.Password = string(hashedPassword)
	result := config.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusCreated, user)
}

// update user
func UpdateUser(c *gin.Context) {
	if !isAdmin(c) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	id := c.Param("id")
	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	var input struct {
		FullName string
		Email    string
		Role     string
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.FullName = input.FullName
	user.Email = input.Email
	user.Role = input.Role

	config.DB.Save(&user)
	c.JSON(http.StatusOK, user)
}

// Delete user
func DeleteUser(c *gin.Context) {
	if !isAdmin(c) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	id := c.Param("id")
	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	config.DB.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}
