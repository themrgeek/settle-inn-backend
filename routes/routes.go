package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/themrgeek/settleinn-backend/controllers"
	middlewares "github.com/themrgeek/settleinn-backend/middleware"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Welcome to Settle Inn")
		c.JSON(200, gin.H{
			"message": "Welcome to the SettleInn API",
		})
	})
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	tenant := r.Group("/tenant")
	tenant.Use(middlewares.AuthMiddleware("tenant"))
	{
		tenant.GET("/properties", controllers.ViewListings)
	}

	owner := r.Group("/owner")
	owner.Use(middlewares.AuthMiddleware("owner"))
	{
		owner.GET("/bookings", controllers.ListOwnerBookings)
		owner.GET("/views", controllers.ViewPropertyStats)
	}

	admin := r.Group("/admin")
	admin.Use(middlewares.AuthMiddleware("admin"))
	{
		admin.GET("/dashboard", controllers.AdminDashboard)
	}
}
