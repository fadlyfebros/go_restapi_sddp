package controllers

import (
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

    // Validasi Parfum berdasarkan ID
    var parfum moduls.Parfum
    if err := moduls.DB.First(&parfum, cartItem.ParfumID).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"message": "Parfum not found"})
        return
    }

    // Cek stok yang cukup
    if parfum.Stock < cartItem.Quantity {
        c.JSON(http.StatusBadRequest, gin.H{"message": "Insufficient stock", "available_stock": parfum.Stock})
        return
    }

    // Ambil UserID dari header
    userID := c.GetHeader("UserID")
    if userID == "" {
        c.JSON(http.StatusUnauthorized, gin.H{"message": "UserID is required"})
        return
    }
    parsedUserID, _ := strconv.ParseInt(userID, 10, 64)

    // Cek atau buat keranjang untuk UserID
    var cart moduls.Cart
    if err := moduls.DB.Where("user_id = ?", parsedUserID).First(&cart).Error; err != nil {
        cart = moduls.Cart{UserID: parsedUserID}
        if err := moduls.DB.Create(&cart).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create cart"})
            return
        }
    }

    // Periksa apakah item sudah ada di keranjang
    var existingCartItem moduls.CartItem
    if err := moduls.DB.Where("cart_id = ? AND parfum_id = ?", cart.ID, cartItem.ParfumID).First(&existingCartItem).Error; err == nil {
        // Tambahkan jumlah jika sudah ada
        existingCartItem.Quantity += cartItem.Quantity
        if existingCartItem.Quantity > parfum.Stock {
            c.JSON(http.StatusBadRequest, gin.H{"message": "Insufficient stock for the updated quantity", "available_stock": parfum.Stock})
            return
        }
        moduls.DB.Save(&existingCartItem)
    } else {
        // Tambahkan item baru ke keranjang
        cartItem.CartID = cart.ID
        cartItem.UserID = parsedUserID
        if err := moduls.DB.Create(&cartItem).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to add item to cart"})
            return
        }
    }

    // Kurangi stok parfum
    parfum.Stock -= cartItem.Quantity
    moduls.DB.Save(&parfum)

    // Preload data parfum untuk respons JSON
    moduls.DB.Preload("Parfum").First(&cartItem, cartItem.ID)

    c.JSON(http.StatusCreated, gin.H{
        "message":   "Item added to cart",
        "cart_item": cartItem,
    })
}

func ViewCart(c *gin.Context) {
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

	var cart moduls.Cart
	if err := moduls.DB.Preload("Items.Parfum").Where("user_id = ?", parsedUserID).First(&cart).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Cart not found"})
		return
	}

	items := []gin.H{}
	total := 0.0
	for _, item := range cart.Items {
		itemTotal := float64(item.Quantity) * item.Parfum.Price
		total += itemTotal

		items = append(items, gin.H{
			"parfum_name": item.Parfum.Name,
			"quantity":    item.Quantity,
			"price":       item.Parfum.Price,
			"item_total":  itemTotal,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"cart_id": cart.ID,
		"user_id": cart.UserID,
		"items":   items,
		"total":   total,
	})
}