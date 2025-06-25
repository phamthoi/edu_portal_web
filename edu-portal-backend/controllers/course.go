package controllers

import (
	"edu-portal-backend/config"
	"edu-portal-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// create new course
func CreateCourse(c *gin.Context) {
	var course models.Course
	if err := c.ShouldBindJSON(&course); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if result := config.DB.Create(&course); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusCreated, course)
}

// get all of course
func GetCourse(c *gin.Context) {
	var courses []models.Course
	config.DB.Find(&courses)
	c.JSON(http.StatusOK, courses)
}

// update the course
func UpdateCourse(c *gin.Context) {
	id := c.Param("id")
	var course models.Course
	if err := config.DB.First(&course, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
		return
	}
	var input models.Course
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	course.Code = input.Code
	course.Name = input.Name
	course.Description = input.Description
	course.Credit = input.Credit
	config.DB.Save(&course)
	c.JSON(http.StatusOK, course)
}

// delete the course
func DeleteCourse(c *gin.Context) {
	id := c.Param("id")
	var course models.Course
	if err := config.DB.First(&course, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
		return
	}
	config.DB.Delete(&course)
	c.JSON(http.StatusOK, gin.H{"message": "Deleted successfully"})
}
