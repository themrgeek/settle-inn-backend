package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/themrgeek/settleinn-backend/config"
	models "github.com/themrgeek/settleinn-backend/model"
)

func ViewListings(c *gin.Context) {
	var properties []models.Property
	if err := config.DB.Find(&properties).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load listings"})
		return
	}
	c.JSON(http.StatusOK, properties)
}
