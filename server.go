// server.go
package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/themrgeek/settleinn-backend/config"
	"github.com/themrgeek/settleinn-backend/routes"
)

func main() {
	config.ConnectDB()
	r := gin.Default()
	routes.SetupRoutes(r)

	// Try running with SSL/HTTPS first
	err := r.RunTLS(":443", "./certs/server.crt", "./certs/server.key")
	if err != nil {
		log.Printf("Failed to start HTTPS server: %v\nFalling back to HTTP on port 8080", err)
		// Fallback to HTTP on port 8080
		err = r.Run(":8080")
		if err != nil {
			log.Fatal("Failed to start HTTP server: ", err)
		}
	}
}
