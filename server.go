package main

import (
	"log"
	"settleinn-backend/config"
	"settleinn-backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()

	r := gin.Default()
	routes.SetupRoutes(r)

	certFile := "cert/cert.pem"
	keyFile := "cert/key.pem"

	log.Println("Starting secure server on https://localhost:443")
	if err := r.RunTLS(":443", certFile, keyFile); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
