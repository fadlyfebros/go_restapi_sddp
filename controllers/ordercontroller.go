package controllers

import (
	"net/http"
	"strconv"

	"github.com/fadlyfebros/go_restapi_sddp/moduls"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Checkout(c *gin.Context) {
	// Ambil UserID dari Header
	userID := c.GetHeader("UserID")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "UserID is required"})
		return
	}

	parsedUserID, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid UserID format"})
		return
	}

	// Cari Keranjang Berdasarkan UserID
	var cart moduls.Cart
	if err := moduls.DB.Preload("Items.Parfum").Where("user_id = ?", parsedUserID).First(&cart).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Cart not found"})
		return
	}

	// Hitung Total Harga
	total := 0.0
	for _, item := range cart.Items {
		total += float64(item.Quantity) * item.Parfum.Price
	}

	// Buat Order Baru
	order := moduls.Order{
		UserID:      parsedUserID,
		CartID:      cart.ID,
		TotalAmount: total,
		Status:      "Pending",
	}
	moduls.DB.Create(&order)

	// Update Stok Barang
	for _, item := range cart.Items {
		moduls.DB.Model(&moduls.Parfum{}).Where("id = ?", item.ParfumID).Update("stock", gorm.Expr("stock - ?", item.Quantity))
	}

	// Hapus Item di Keranjang Setelah Checkout
	moduls.DB.Where("cart_id = ?", cart.ID).Delete(&moduls.CartItem{})

	c.JSON(http.StatusCreated, gin.H{
		"message": "Checkout successful",
		"order":   order,
	})
}