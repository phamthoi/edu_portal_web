package controllers

import (
	"edu-portal-backend/config"
	"edu-portal-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// create new class
func CreateClass(c *gin.Context) {
	var class models.Class
	if err := c.ShouldBindJSON(&class); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if result := config.DB.Create(&class); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusCreated, class)
}

// get all class
func GetClasses(c *gin.Context) {
	var classes []models.Class
	config.DB.Preload("Course").Preload("teacher").Find(&classes)
	c.JSON(http.StatusOK, classes)
}

// update class information
func UpdateClass(c *gin.Context) {
	id := c.Param("id")
	var class models.Class
	if err := config.DB.First(&class, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	var input models.Class
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	class.CourseID = input.CourseID
	class.TeacherID = input.TeacherID
	class.Semester = input.Semester
	class.Year = input.Year

	config.DB.Save(&class)
	c.JSON(http.StatusOK, class)
}

// Delete class
func DeleteClass(c *gin.Context) {
	id := c.Param("id")
	var class models.Class
	if err := config.DB.First(&class, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Class not found"})
		return
	}
	config.DB.Delete(&class)
	c.JSON(http.StatusOK, gin.H{"message": "Deleted successfully"})
}

/* "createclass": input JSON data send from client, connect into class
- check error of binding JSON.
- If it doesn't have error, it will create a new record in "classes" tables and return JSON data of the class which we just create.

"GetClasses:"
- get classes list from database
- use "Preload" to get course information and teacher which regarding.

"UpdateClass:"
- Find class base on ID from URL.
- Binding JSON data to update
- save changes and return updated data

"DeleteClass:"
- find class base on ID, if it dosen't see. it will error.
*/
