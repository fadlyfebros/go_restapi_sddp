package moduls

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "root:@tcp(127.0.0.1:3306)/go_restapi_sddp"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	// Tambahkan semua model ke AutoMigrate
	err = DB.AutoMigrate(
		&Parfum{},
		&Cart{},
		&CartItem{},
		&Order{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// Insert dummy data untuk parfum
	dummyData := []Parfum{
		{ID: 1, Name: "Dior Sauvage", Size: "50 ml", Price: 150000.00, Stock: 10},
		{ID: 2, Name: "Bleu de Chanel", Size: "100 ml", Price: 120000.00, Stock: 5},
		{ID: 3, Name: "Parfum Gucci Bloom", Size: "30 ml", Price: 100000.00, Stock: 8},
		{ID: 4, Name: "Creed Aventus", Size: "40 ml", Price: 128000.00, Stock: 5},
		{ID: 5, Name: "Paco Rabanne 1 Million ", Size: "100 ml", Price: 128000000.00, Stock: 5},
	}
	if err := DB.Create(&dummyData).Error; err != nil {
		log.Println("Dummy data already exists or failed to insert:", err)
	}
}