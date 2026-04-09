package main

import (
	"net"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/gin-gonic/gin"
)

type HealthResponse struct {
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
	Version   string `json:"version,omitempty"`
	BuildTime string `json:"build_time,omitempty"`
}

func listUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"users": []string{}})
}

func getUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"id": c.Param("id")})
}

func createUser(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"status": "created"})
}

func updateUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "updated"})
}

func deleteUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"deleted": c.Param("id")})
}

func HealthCheck(c *gin.Context) {
	info, _ := debug.ReadBuildInfo()

	resp := HealthResponse{
		Status:    "healthy",
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Version:   info.GoVersion, // You can read from build info
		BuildTime: "2026-04-09",
	}

	c.JSON(http.StatusOK, resp)
}

// RegisterUserRoutes registers user-related routes to the provided router group.
func RegisterUserRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/users")
	{
		// Define user-related routes
		users.GET("", listUsers)
		users.GET("/:id", getUser)
		users.POST("", createUser)
		users.PUT("/:id", updateUser)
		users.DELETE("/:id", deleteUser)
	}
}

// Helper function
func getPortFromHost(host string, isHTTPS bool) (string, string) {
	host, port, err := net.SplitHostPort(host)

	if err == nil {
		return port, host
	}

	if isHTTPS {
		return "443", host
	}

	return "80", host
}

func main() {
	// Recommended for most projects (includes Logger + Recovery middleware)
	server := gin.Default()

	// For production, consider using gin.New() and adding only the necessary middleware for better performance and control.
	// Clean router - no built-in middleware (better for production control)
	// server := gin.New()

	// Handle 404 for undefined routes
	server.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Route not found"})
	})

	server.GET("/", func(c *gin.Context) {
		// Then use it inside handler:
		port, hostname := getPortFromHost(c.Request.Host, c.Request.TLS != nil)

		c.JSON(http.StatusOK, gin.H{
			"message":   hostname + ":" + port + " and is healthy",
			"status":    "success",
			"code":      http.StatusOK,
			"client_ip": c.ClientIP(), // ← This shows "::1" locally
		})
	})

	server.GET("/health", HealthCheck)

	server.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Grouping routes under /api/v1
	v1 := server.Group("/api/v1")
	RegisterUserRoutes(v1)

	server.Run(":8080")                          // Listen on 0.0.0.0:8080
	server.Run("127.0.0.1:8080")                 // Specific interface
	server.RunTLS(":443", "cert.pem", "key.pem") // HTTPS
	// Tell Gin which IPs are trusted proxies (important for production too)
	// For local development, trust loopback addresses
	server.SetTrustedProxies([]string{"127.0.0.1", "::1", "localhost"})
}
