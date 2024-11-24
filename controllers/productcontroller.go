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
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input"})
		return
	}
	moduls.DB.Create(&parfum)
	c.JSON(http.StatusCreated, parfum)
}

func UpdateParfum(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	var parfum moduls.Parfum
	if err := moduls.DB.First(&parfum, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Parfum not found"})
		return
	}
	if err := c.ShouldBindJSON(&parfum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input"})
		return
	}
	moduls.DB.Save(&parfum)
	c.JSON(http.StatusOK, parfum)
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