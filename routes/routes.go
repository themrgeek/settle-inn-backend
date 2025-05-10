package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/settleinn-backend/controllers"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/login", controllers.Login)

	tenant := r.Group("/tenant")
	tenant.Use(middlewares.AuthMiddleware("tenant"))
	{
		tenant.GET("/properties", controllers.ViewProperties)
	}

	owner := r.Group("/owner")
	owner.Use(middlewares.AuthMiddleware("owner"))
	{
		owner.GET("/bookings", controllers.ViewBookings)
		owner.GET("/views", controllers.PropertyViews)
	}

	admin := r.Group("/admin")
	admin.Use(middlewares.AuthMiddleware("admin"))
	{
		admin.GET("/dashboard", controllers.AdminDashboard)
	}
}
