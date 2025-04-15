package config

import (
	"fmt"
	"log"
	"medi_remind_backend/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "root:@tcp(127.0.0.1:3306)/medi_remind?charset=utf8mb4&parseTime=True&loc=Local"
	var err error

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Gagal konek ke database: %v", err)
	}

	fmt.Println("Database berhasil terkoneksi!")

	err = DB.AutoMigrate(&models.User{}, &models.Obat{})
	if err != nil {
		log.Fatalf("Gagal migrate database: %v", err)
	}
}
