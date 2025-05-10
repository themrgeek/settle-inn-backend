package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/themrgeek/settleinn-backend/config"
	models "github.com/themrgeek/settleinn-backend/model"
)

func ListOwnerBookings(c *gin.Context) {
	ownerID := c.GetInt("user_id")
	var bookings []models.Booking
	if err := config.DB.Joins("JOIN properties ON bookings.property_id = properties.id").Where("properties.owner_id = ?", ownerID).Find(&bookings).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load bookings"})
		return
	}
	c.JSON(http.StatusOK, bookings)
}
func ViewPropertyStats(c *gin.Context) {
	ownerID := c.GetInt("user_id")
	var properties []models.Property
	if err := config.DB.Where("owner_id = ?", ownerID).Find(&properties).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load properties"})
		return
	}
	c.JSON(http.StatusOK, properties)
}
