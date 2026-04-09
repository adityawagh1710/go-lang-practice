package main

import (
	"errors"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// ── Sentinel errors ───────────────────────────────────────────────────────────

var (
	ErrNotFound     = errors.New("resource not found")
	ErrUnauthorized = errors.New("unauthorized")
	ErrBadRequest   = errors.New("bad request")
)

// ── Middleware ─────────────────────────────────────────────────────────────────

func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.GetHeader("X-Request-ID")
		if requestID == "" {
			requestID = uuid.NewString()
		}
		c.Set("request_id", requestID)
		c.Header("X-Request-ID", requestID)
		c.Next()
	}
}

func SimpleLogger(logger *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next() // handlers run here

		status := c.Writer.Status()

		// pick log level based on status code
		attrs := []any{
			slog.String("method", c.Request.Method),
			slog.String("path", c.Request.URL.Path),
			slog.String("query", c.Request.URL.RawQuery),
			slog.Int("status", status),
			slog.Duration("latency", time.Since(start)),
			slog.String("ip", c.ClientIP()),
			slog.String("request_id", c.GetString("request_id")), // from RequestIDMiddleware
		}

		switch {
		case status >= 500:
			logger.Error("request", attrs...)
		case status >= 400:
			logger.Warn("request", attrs...)
		default:
			logger.Info("request", attrs...)
		}
	}
}

func ErrorHandlerMiddleware(logger *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) == 0 || c.Writer.Written() {
			return
		}

		err := c.Errors.Last().Err

		// log every accumulated error
		for _, e := range c.Errors {
			logger.Error("handler error",
				slog.String("path", c.Request.URL.Path),
				slog.String("method", c.Request.Method),
				slog.String("request_id", c.GetString("request_id")),
				slog.String("error", e.Error()),
			)
		}

		// map sentinel errors → status codes
		var statusCode int
		switch {
		case errors.Is(err, ErrNotFound):
			statusCode = http.StatusNotFound
		case errors.Is(err, ErrUnauthorized):
			statusCode = http.StatusUnauthorized
		case errors.Is(err, ErrBadRequest):
			statusCode = http.StatusBadRequest
		default:
			statusCode = http.StatusInternalServerError
		}

		c.JSON(statusCode, gin.H{
			"error":      err.Error(),
			"status":     statusCode,
			"request_id": c.GetString("request_id"),
		})
	}
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Authorization header is required",
				"message": "Please add Bearer token",
			})
			return
		}

		const prefix = "Bearer "
		if len(authHeader) > len(prefix) && authHeader[:len(prefix)] == prefix {
			token := authHeader[len(prefix):]
			if token == "super-secret-token" {
				c.Set("user_id", "user-12345")
				c.Set("role", "admin")
				c.Set("is_authenticated", true)
				c.Next()
				return
			}
		}

		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error":   "Invalid token",
			"message": "Please use 'super-secret-token'",
		})
	}
}

// ── Handlers ───────────────────────────────────────────────────────────────────

func getProfile(c *gin.Context) {
	userID := c.MustGet("user_id").(string)
	c.JSON(http.StatusOK, gin.H{
		"message":    "Profile fetched",
		"user_id":    userID,
		"role":       c.GetString("role"),
		"request_id": c.GetString("request_id"),
	})
}

type UserRequest struct {
	Name     string `json:"name"     binding:"required,min=3,max=50"`
	Email    string `json:"email"    binding:"required,email"`
	Age      int    `json:"age"      binding:"required,gte=18,lte=100"`
	Password string `json:"password" binding:"required,min=8"`
	Role     string `json:"role"     binding:"oneof=user admin moderator"`
}

func createUser(c *gin.Context) { // renamed from getUsers
	var req UserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":      err.Error(),
			"status":     "failed",
			"message":    "Invalid request payload",
			"request_id": c.GetString("request_id"),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message":    "User created successfully",
		"user":       req,
		"request_id": c.GetString("request_id"),
	})
}

// ── Main ───────────────────────────────────────────────────────────────────────
// Recovery → RequestID → ErrorHandler → SimpleLogger → AuthMiddleware → Handlers
func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	server := gin.New() // bare engine — we add middleware manually below

	// Proxy trust — must be set BEFORE Run
	server.SetTrustedProxies([]string{"127.0.0.1", "::1"})

	// Global middleware — order matters
	server.Use(gin.Recovery())                 // 1. catch panics first
	server.Use(RequestIDMiddleware())          // 2. stamp every request with an ID
	server.Use(ErrorHandlerMiddleware(logger)) // 3. catches c.Error() calls after chain
	server.Use(SimpleLogger(logger))           // 4. logs last — captures real status + latency

	server.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"error":      "route not found",
			"request_id": c.GetString("request_id"),
		})
	})

	// Protected routes
	v2 := server.Group("/api/v2")
	v2.Use(AuthMiddleware())
	{
		v2.GET("/profile", getProfile)
		v2.POST("/users", createUser)

		v2.GET("/test-notfound", func(c *gin.Context) {
			c.Error(ErrNotFound)
		})

		v2.GET("/test-badrequest", func(c *gin.Context) {
			c.Error(ErrBadRequest)
		})

		v2.GET("/test-unauth", func(c *gin.Context) {
			c.Error(ErrUnauthorized)
		})
	}

	// Single Run call — this blocks
	if err := server.Run(":8080"); err != nil {
		logger.Error("server failed", slog.String("error", err.Error()))
		os.Exit(1)
	}
}
