package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/BarryDaBee/go_restaurant_management/internal/config"
	"github.com/BarryDaBee/go_restaurant_management/internal/repository/storage/postgresql"
	"github.com/BarryDaBee/go_restaurant_management/models"
	"github.com/BarryDaBee/go_restaurant_management/pkg/router"
	"github.com/BarryDaBee/go_restaurant_management/utility"
	"github.com/go-playground/validator/v10"
)

func main() {
	// Initialize logger
	logger := utility.NewLogger()

	// Load configuration
	configuration := config.Setup(".env")

	// Initialize database storage
	storage := postgresql.NewPostgresStore(configuration)
	storage.Connect(logger, configuration.Databases)

	// Run migrations
	if err := storage.Migrate(&models.User{}, &models.Food{}, &models.Order{}); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}
	utility.LogAndPrint(logger, "Database migrations completed")

	// Initialize validator
	validatorRef := validator.New()

	// Setup router with all routes
	r := router.Setup(logger, validatorRef, storage, configuration)

	// Determine port
	port := configuration.Server.Port
	if port == "" {
		port = os.Getenv("PORT")
		if port == "" {
			port = "8080"
		}
	}

	// Create HTTP server
	httpServer := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	// Start server in goroutine
	go func() {
		utility.LogAndPrint(logger, fmt.Sprintf("Server is running on port %s", port))
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	utility.LogAndPrint(logger, "Shutting down server...")

	// Give outstanding requests time to complete
	_ = time.Second * 5

	utility.LogAndPrint(logger, "Server exited properly")
}
