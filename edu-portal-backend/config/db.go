package config

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB //khai báo biến toàn cục

func ConnectDB() {
	dsn := "root:1234@tcp(127.0.0.1:3306)/edu_portal?charset=utf8mb4&parseTime=True&loc=Local"
	/* name_database: edu_portal
	charset=utf8mb4: bộ mã UTF-8 mở rộng(hỗ trợ emoji)
	parseTime=true: cho phép GORM hiểu định dạng DATETIME.
	loc=local: đặt timezone là hệ thống local.
	*/
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//&gorm.Config{}: config GORM Default

	if err != nil {
		log.Fatal("Cannot connect to database: ", err)
	}

	DB = db
}
