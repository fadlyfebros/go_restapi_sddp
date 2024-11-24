package moduls

import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "log"
)

var DB *gorm.DB

// Fungsi untuk menghubungkan ke database
func ConnectDatabase() {
<<<<<<< HEAD
    dsn := "root:@tcp(localhost:3306)/go_restapi_sddp"
    var err error
    DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to the database:", err)
    }
=======
	dsn := "root:@tcp(127.0.0.1:3306)/go_restapi_sddp"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
>>>>>>> 399dcdc9ca4e492a45cecbd0e5d560b70594a515

    // Migrasi model Parfum ke dalam database
    err = DB.AutoMigrate(&Parfum{}, &Cart{}, &CartItem{}, &Order{})
    if err != nil {
        log.Fatal("Failed to migrate database:", err)
    }

    // Tambahkan data dummy
    DB.Create(&[]Parfum{
        {ID: 1, Name: "Parfum Chanel No. 5", Price: 150.00, Stock: 10, Category: "Parfum"},
        {ID: 2, Name: "Parfum Dior Sauvage", Price: 120.00, Stock: 5, Category: "Parfum"},
        {ID: 3, Name: "Parfum Gucci Bloom", Price: 100.00, Stock: 8, Category: "Parfum"},
        {ID: 4, Name: "Parfum Vanilla Clouds", Price: 128.00, Stock: 5, Category: "Parfum"},
    })
}
