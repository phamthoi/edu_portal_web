package main

import (
	"edu-portal-backend/config"
	"edu-portal-backend/models"
)

func main() {
	config.ConnectDB()

	//migrate models(di chuyển models)
	config.DB.AutoMigrate(&models.User{})

}
