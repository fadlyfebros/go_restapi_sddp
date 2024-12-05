package controllers

import (
	"net/http"
	"strconv"

	"github.com/fadlyfebros/go_restapi_sddp/moduls"
	"github.com/gin-gonic/gin"
)

func GetAllParfum(c *gin.Context) {
	var parfums []moduls.Parfum
	moduls.DB.Find(&parfums)
	c.JSON(http.StatusOK, parfums)
}

func GetParfumByID(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	var parfum moduls.Parfum
	if err := moduls.DB.First(&parfum, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Parfum not found"})
		return
	}
	c.JSON(http.StatusOK, parfum)
}

func CreateParfum(c *gin.Context) {
	var parfum moduls.Parfum

	if err := c.ShouldBindJSON(&parfum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input", "error": err.Error()})
		return
	}

	if parfum.Name == "" || parfum.Price <= 0 || parfum.Stock <= 0 || parfum.Size == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "All fields (name, size, price, stock) are required and must be valid"})
		return
	}

	if err := moduls.DB.Create(&parfum).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create parfum", "error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Parfum successfully created",
		"parfum":  parfum,
	})
}

func UpdateParfum(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
		return
	}

	var parfum moduls.Parfum
	if err := moduls.DB.First(&parfum, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Parfum not found"})
		return
	}

	var updatedParfum moduls.Parfum
	if err := c.ShouldBindJSON(&updatedParfum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input", "error": err.Error()})
		return
	}

	// Validasi data yang diperbarui
	if updatedParfum.Name == "" || updatedParfum.Price <= 0 || updatedParfum.Stock <= 0 || updatedParfum.Size == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "All fields (name, size, price, stock) are required and must be valid"})
		return
	}

	// Update data parfum
	parfum.Name = updatedParfum.Name
	parfum.Price = updatedParfum.Price
	parfum.Stock = updatedParfum.Stock
	parfum.Size = updatedParfum.Size

	if err := moduls.DB.Save(&parfum).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update parfum", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Parfum successfully updated",
		"parfum":  parfum,
	})
}


func DeleteParfum(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	var parfum moduls.Parfum
	if err := moduls.DB.First(&parfum, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Parfum not found"})
		return
	}
	moduls.DB.Delete(&parfum)
	c.JSON(http.StatusOK, gin.H{"message": "Parfum deleted"})
}