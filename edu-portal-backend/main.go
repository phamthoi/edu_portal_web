package main

import (
	"edu-portal-backend/config"
	"edu-portal-backend/models"
	"edu-portal-backend/routes"
)

func main() {
	config.ConnectDB()

	//migrate models(di chuyển models)
	config.DB.AutoMigrate(
		&models.User{},
		&models.Course{},
		&models.Class{},
		&models.Enrollment{},
	)
	r := routes.SetupRoute()
	r.Run(":8080")

}
