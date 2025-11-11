package main

import (
	"fmt"
	"os"

	"github.com/Aditya7880900936/auth_golang.git/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env only if it exists (safe for Render)
	if err := godotenv.Load(); err != nil {
		fmt.Println("⚠️  No .env file found — using Render environment variables")
	}

	// Get the PORT from environment
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // default for local
	}

	router := gin.New()
	router.Use(gin.Logger())

	// Register routes
	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	// Test routes
	router.GET("/api-1", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for api-1"})
	})

	router.GET("/api-2", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for api-2"})
	})

	fmt.Println("🚀 Server is running on port", port)
	router.Run(":" + port)
}
