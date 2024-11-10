package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/fadlyfebros/go_restapi_sddp/moduls"
	"github.com/gin-gonic/gin"
)

// GET /api/product
func Index(c *gin.Context) {
    var products []moduls.Product
    moduls.DB.Find(&products)
    c.JSON(http.StatusOK, products)
}

// GET /api/product/:id
func Show(c *gin.Context) {
    id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
    var product moduls.Product
    if err := moduls.DB.First(&product, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"message": "Product not found"})
        return
    }
    c.JSON(http.StatusOK, product)
}

// POST /api/product
func Create(c *gin.Context) {
    fmt.Println("Create function called")
    var product moduls.Product
    if err := c.ShouldBindJSON(&product); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input"})
        return
    }
    moduls.DB.Create(&product)
    c.JSON(http.StatusCreated, product)
}
// PUT /api/product/:id
func Update(c *gin.Context) {
    id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
    var product moduls.Product
    if err := moduls.DB.First(&product, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"message": "Product not found"})
        return
    }
    if err := c.ShouldBindJSON(&product); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input"})
        return
    }
    moduls.DB.Save(&product)
    c.JSON(http.StatusOK, product)
}

// DELETE /api/product/:id
func Delete(c *gin.Context) {
    id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
    var product moduls.Product
    if err := moduls.DB.First(&product, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"message": "Product not found"})
        return
    }
    moduls.DB.Delete(&product)
    c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
}