package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/fadlyfebros/go_restapi_sddp/moduls"
	"github.com/gin-gonic/gin"
)

func AddToCart(c *gin.Context) {
    var cartItem moduls.CartItem
    if err := c.ShouldBindJSON(&cartItem); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input"})
        return
    }

    // Validasi parfum
    var parfum moduls.Parfum
    if err := moduls.DB.First(&parfum, cartItem.ParfumID).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"message": "Parfum not found"})
        return
    }

    if parfum.Stock < cartItem.Quantity {
        c.JSON(http.StatusBadRequest, gin.H{"message": "Insufficient stock"})
        return
    }

    // Validasi UserID
    userID := c.GetHeader("UserID")
    if userID == "" {
        c.JSON(http.StatusUnauthorized, gin.H{"message": "UserID is required"})
        return
    }
    parsedUserID, _ := strconv.ParseInt(userID, 10, 64)

    // Cari atau buat keranjang
    var cart moduls.Cart
    if err := moduls.DB.Where("user_id = ?", parsedUserID).First(&cart).Error; err != nil {
        cart = moduls.Cart{UserID: parsedUserID}
        moduls.DB.Create(&cart)
    }

    // Tambahkan item ke keranjang
    cartItem.CartID = cart.ID
    cartItem.UserID = parsedUserID
    moduls.DB.Create(&cartItem)

    // Preload data parfum untuk respons JSON
    moduls.DB.Preload("Parfum").First(&cartItem, cartItem.ID)

    c.JSON(http.StatusCreated, gin.H{
        "message":   "Item added to cart",
        "cart_item": cartItem,
    })
}

func ViewCart(c *gin.Context) {
    // Ambil UserID dari Header
    userID := c.GetHeader("UserID")
    if userID == "" {
        c.JSON(http.StatusUnauthorized, gin.H{"message": "UserID is required"})
        return
    }

    // Parsing UserID dari string ke int64
    parsedUserID, err := strconv.ParseInt(userID, 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid UserID format"})
        return
    }

    // Log untuk debug
    fmt.Println("UserID diterima:", parsedUserID)

    // Cari keranjang berdasarkan UserID
    var cart moduls.Cart
    if err := moduls.DB.Preload("Items.Parfum").Where("user_id = ?", parsedUserID).First(&cart).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"message": "Cart not found"})
        return
    }

    // Kirim respons berisi keranjang
    c.JSON(http.StatusOK, gin.H{
        "cart": cart,
    })
}