package controllers

import (
	"edu-portal-backend/config"
	"edu-portal-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Student registed class.
func EnrollClass(c *gin.Context) {
	var input struct {
		ClassID uint `json:"class_id"`
	}
	userID := c.GetUint("userID")
	role := c.GetString("role")

	if role != "student" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only students can enroll"})
		return
	}
	enrollment := models.Enrollment{
		StudentID: userID,
		ClassID:   input.ClassID,
	}
	result := config.DB.Create(&enrollment)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Enrolled successfully"})
}

// watch student list in the class(this is for admin/teacher)
func GetStudentInClass(c *gin.Context) {
	classID := c.Param("id")
	var enrollments []models.Enrollment

	config.DB.Preload("Student").Where("class_id = ?", classID).Find(&enrollments)
	c.JSON(http.StatusOK, enrollments)
}

// the teacher grading(giảng viên chấm điểm)
func UpdateScore(c *gin.Context) {
	enrollmentID := c.Param("id")
	var input struct {
		Score float64 `json:"score"`
	}
	var enrollment models.Enrollment

	if err := config.DB.First(&enrollment, enrollmentID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "enrollment not found"})
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	enrollment.Score = &input.Score
	config.DB.Save(&enrollment)
	c.JSON(http.StatusOK, enrollment)
}

func GetStudentGrades(c *gin.Context) {
	role := c.GetString("role")
	userID := c.GetUint("userID")

	if role != "student" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only students can view their grades"})
		return
	}

	var enrollments []models.Enrollment
	config.DB.Preload("Class.Course").Where("student_id = ?", userID).Find(&enrollments)

	//create result list simplely
	var results []gin.H
	for _, e := range enrollments {
		results = append(results, gin.H{
			"course":   e.Class.Course.Name,
			"semester": e.Class.Semester,
			"year":     e.Class.Year,
			"score":    e.Score,
		})
	}

	c.JSON(http.StatusOK, results)
}
