package main

import (
	"github.com/fadlyfebros/go_restapi_sddp/controllers"
	"github.com/fadlyfebros/go_restapi_sddp/moduls"
	"github.com/gin-gonic/gin"
)

func main() {
	// Koneksi ke database
	moduls.ConnectDatabase()

	// Inisialisasi router
	router := gin.Default()

	// Routing untuk produk/parfum
	router.GET("/api/parfum", controllers.GetAllParfum)
	router.GET("/api/parfum/:id", controllers.GetParfumByID)
	router.POST("/api/parfum", controllers.CreateParfum)
	router.PUT("/api/parfum/:id", controllers.UpdateParfum)
	router.DELETE("/api/parfum/:id", controllers.DeleteParfum)

	// Routing untuk keranjang (cart)
	router.POST("/api/cart/add", controllers.AddToCart)
	router.GET("/api/cart/view", controllers.ViewCart)

	// Routing untuk checkout
	router.POST("/api/order/checkout", controllers.Checkout)

	// Jalankan server
	router.Run(":8080")
}