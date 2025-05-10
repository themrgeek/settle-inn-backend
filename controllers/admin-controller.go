package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/themrgeek/settleinn-backend/config"
	models "github.com/themrgeek/settleinn-backend/model"
)

func GetAllUsers(c *gin.Context) {
	var users []models.User
	if err := config.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get users"})
		return
	}
	c.JSON(http.StatusOK, users)
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&models.User{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Delete failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}
func AdminDashboard(c *gin.Context) {
	var stats struct {
		TotalUsers    int64
		ActiveUsers   int64
		InactiveUsers int64
		LatestUsers   []models.User
	}

	// Get total users count
	config.DB.Model(&models.User{}).Count(&stats.TotalUsers)

	// Get active users count
	config.DB.Model(&models.User{}).Where("status = ?", "active").Count(&stats.ActiveUsers)

	// Get inactive users count
	config.DB.Model(&models.User{}).Where("status = ?", "inactive").Count(&stats.InactiveUsers)

	// Get latest 5 users
	config.DB.Order("created_at desc").Limit(5).Find(&stats.LatestUsers)

	c.JSON(http.StatusOK, gin.H{
		"statistics": stats,
	})
}
