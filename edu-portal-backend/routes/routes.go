package routes

import (
	"edu-portal-backend/controllers"
	"edu-portal-backend/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoute() *gin.Engine {
	r := gin.Default()
	//Routes don't need auth. we can put it exception this group
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	//Route security(test)
	//Route security by middleware
	auth := r.Group("/api")
	auth.Use(middleware.AuthMiddlware())
	{
		auth.GET("/profile", func(c *gin.Context) {
			userID := c.GetUint("userID")
			role := c.GetString("role")
			c.JSON(200, gin.H{"user_id": userID, "role": role})
		})

		//course related API
		auth.GET("/courses", controllers.GetCourse)
		auth.POST("/course", controllers.CreateCourse)
		auth.PUT("/courses/:id", controllers.UpdateCourse)
		auth.DELETE("/courses/:id", controllers.DeleteCourse)

		// API related classes
		auth.GET("/classes", controllers.GetClasses)
		auth.POST("/classes", controllers.CreateClass)
		auth.PUT("/classes/:id", controllers.UpdateClass)
		auth.DELETE("/classes/:id", controllers.DeleteClass)

		//Enrollment API
		auth.POST("/enroll", controllers.EnrollClass)
		auth.GET("/classes/:id/students", controllers.GetStudentInClass)
		auth.PUT("/enrollments/:id/score", controllers.UpdateCourse)

		//manager user
		auth.GET("/users", controllers.GetUser)
		auth.POST("/users", controllers.CreateUser)       //admin create
		auth.PUT("/users/:id", controllers.UpdateUser)    // admin can change
		auth.DELETE("/users/:id", controllers.DeleteUser) //admin delete

		//student get score of their
		auth.GET("/my-grades", controllers.GetStudentGrades)
	}

	return r
}
