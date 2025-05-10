// server.go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/themrgeek/settleinn-backend/routes"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file. Make sure .env file exists in the root directory:", err)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true",
		os.Getenv("DB_USER"), os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	DB = db
	fmt.Println("Database connected")
}
func main() {
	ConnectDB()
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
