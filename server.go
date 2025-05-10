// server.go
package main

import (
	"github.com/themrgeek/settleinn-backend/config"
	"github.com/themrgeek/settleinn-backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()
	r := gin.Default()
	routes.SetupRoutes(r)
	r.Run(":8080")
}
