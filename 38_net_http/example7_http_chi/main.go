package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// Custom middleware (same as before)
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("→ %s %s", r.Method, r.URL.Path)

		next.ServeHTTP(w, r)

		log.Printf("← %s %s (%v)", r.Method, r.URL.Path, time.Since(start))
	})
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from chi!")
}

func main() {
	r := chi.NewRouter()

	// Global middleware stack (applied to all routes)
	r.Use(middleware.RequestID) // built-in
	r.Use(middleware.RealIP)
	r.Use(LoggingMiddleware)    // custom
	r.Use(middleware.Recoverer) // panic recovery

	// Route-specific / grouped middleware
	r.Route("/api", func(r chi.Router) {
		//r.Use(AuthMiddleware) // only applies to /api/*

		r.Get("/hello", HelloHandler)
		r.Get("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
			id := chi.URLParam(r, "id")
			fmt.Fprintf(w, "User ID: %s", id)
		})
	})

	log.Println("chi server starting on :8090")
	log.Fatal(http.ListenAndServe(":8090", r))
}
