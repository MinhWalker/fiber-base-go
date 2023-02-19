package main

import (
	"fiber-base-go/config"
	"fiber-base-go/internal/delivery/api/handlers"
	"fiber-base-go/internal/repository"
	"fiber-base-go/internal/services"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	Run(3000)
}

// Run start server
func Run(port int) {
	// Load the configuration
	cfg, err := config.LoadConfig("configs/dev.yaml")
	if err != nil {
		log.Fatalf("failed to load config: %s", err)
	}

	conn, err := config.ConnectDb(cfg)
	if err != nil {
		log.Fatalf("failed to connect database: %s", err)
	}

	// Automigrate the database schema
	config.DBMigrate(conn)

	// Create the repository
	userRepo := repository.NewStudentRepository(conn)

	// Create the service
	userService := services.NewStudentService(userRepo)

	// Create the Fiber app
	app := fiber.New()

	// Register the handler
	studentHandler := handlers.NewStudentHandler(userService)

	// Register the routes
	studentHandler.RegisterRoutes(app)

	// Start the server
	go func() {
		if err := app.Listen(":3000"); err != nil {
			log.Fatalf("failed to start server: %s", err)
		}
	}()

	// Wait for a signal to gracefully shut down the server
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig

	if err := app.Shutdown(); err != nil {
		log.Printf("error shutting down server: %s", err)
	}

}
