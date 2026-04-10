package main

import (
	"gin-gorm-api/config"
	"gin-gorm-api/routes"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

func main() {
	godotenv.Load()

	config.ConnectDB()

	r := gin.Default()

	r.SetTrustedProxies([]string{"127.0.0.1", "::1"})

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"error":      "route not found",
			"request_id": c.GetString("request_id"),
		})
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message":    "Welcome to the Gin-GORM API",
			"request_id": c.GetString("request_id"),
		})
	})

	routes.SetupRoutes(r)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	if err := r.Run(":" + port); err != nil {
		log.Error().Err(err).Msg("server failed")
		os.Exit(1)
	}
}
