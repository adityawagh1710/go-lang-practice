package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Same custom middlewares as above
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
	fmt.Fprintln(w, "Hello from gorilla/mux!")
}

func main() {
	r := mux.NewRouter()

	// Global middleware (applied to everything registered on this router)
	r.Use(LoggingMiddleware)
	r.Use(mux.MiddlewareFunc(AuthMiddleware)) // mux also accepts MiddlewareFunc

	// Or apply per-subrouter (recommended for scoping)
	api := r.PathPrefix("/api").Subrouter()

	api.Use(AuthMiddleware) // only for /api/*

	api.HandleFunc("/hello", HelloHandler).Methods("GET")

	api.HandleFunc("/users/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		fmt.Fprintf(w, "User ID: %s", vars["id"])
	}).Methods("GET")

	log.Println("gorilla/mux server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
