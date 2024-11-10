package main

import (
    "github.com/fadlyfebros/go_restapi_sddp/controllers"
    "github.com/fadlyfebros/go_restapi_sddp/moduls"

    "github.com/gin-gonic/gin"
)

func main() {
    // Koneksi database
    moduls.ConnectDatabase()

    // Inisialisasi router
    router := gin.Default()

    // Routing
    router.GET("/api/product", controllers.Index)
    router.GET("/api/product/:id", controllers.Show)
    router.POST("/api/product", controllers.Create)
    router.PUT("/api/product/:id", controllers.Update)
    router.DELETE("/api/product/:id", controllers.Delete)

    // Jalankan server di port 8080
    router.Run(":8080")
}