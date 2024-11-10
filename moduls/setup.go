package moduls

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB


// Fungsi untuk menghubungkan ke database
func ConnectDatabase() {
	dsn := "root:@tcp(http://localhost:8080)/go_restapi_sddp"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	// Migrasi model Product ke dalam database
	err = DB.AutoMigrate(&Product{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
}
