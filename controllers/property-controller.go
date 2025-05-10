package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/themrgeek/settleinn-backend/config"
	models "github.com/themrgeek/settleinn-backend/model"
)

func CreateProperty(c *gin.Context) {
	var property models.Property
	if err := c.ShouldBindJSON(&property); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid property data"})
		return
	}

	property.OwnerID = uint(c.GetInt("user_id"))
	if err := config.DB.Create(&property).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create property"})
		return
	}
	c.JSON(http.StatusOK, property)
}

func IncrementPropertyView(c *gin.Context) {
	var property models.Property
	id := c.Param("id")
	if err := config.DB.First(&property, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Property not found"})
		return
	}
	property.Views++
	config.DB.Save(&property)
	c.JSON(http.StatusOK, gin.H{"message": "View incremented"})
}
