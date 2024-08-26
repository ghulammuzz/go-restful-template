package main

import (
	"log"
	"net/http"
	"time"

	"github.com/ghulammuzz/go-restful-template/config"
	"github.com/ghulammuzz/go-restful-template/internal/middleware/compression"
	"github.com/ghulammuzz/go-restful-template/internal/middleware/cors"
	"github.com/ghulammuzz/go-restful-template/internal/middleware/rate"
	"github.com/ghulammuzz/go-restful-template/internal/routes"
	"github.com/ghulammuzz/go-restful-template/pkg/env"
	"github.com/ghulammuzz/go-restful-template/pkg/logger"
)

func main() {
	env.LoadEnv()

	profile := env.GetEnv("APP_PROFILE")
	logger.InitLogger(profile)

	db, err := config.Initialize()
	if err != nil {
		logger.Error("Failed to connect to database", "error", err)
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	mux := http.NewServeMux()

	routes.SetupRoutes(mux, db)
	corsConfig := cors.CORSConfig{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	}
	corsMux := cors.CORS(corsConfig)(mux)

	compressedMux := compression.GzipCompressionMiddleware(corsMux)
	rateLimiter := rate.NewRateLimiter(100, time.Minute, time.Minute*5)
	rateLimitedMux := rateLimiter.Middleware(compressedMux)

	port := ":8080"
	logger.Info("Server is starting on port", "port", port)
	if err := http.ListenAndServe(port, rateLimitedMux); err != nil {
		logger.Error("Failed to start server", "error", err)
		log.Fatalf("Failed to start server: %v", err)
	}
}
