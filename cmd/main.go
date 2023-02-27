package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"fiber-base-go/config"
	"fiber-base-go/internal/delivery/api/handlers"
	"fiber-base-go/internal/repository"
	"fiber-base-go/internal/services"

	"github.com/gofiber/fiber/v2"
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

	oauthConfig := config.InitOAuthConfig(cfg)

	// Automigrate the database schema
	config.DBMigrate(conn)

	// Create the repository
	studentRepo := repository.NewStudentRepository(conn)
	userRepo := repository.NewUserRepository(conn)

	// Create the service
	studentService := services.NewStudentService(studentRepo)
	userService := services.NewUserService(userRepo)
	// Create the contest repository
	contestRepo := repository.NewContestRepository(conn)

	// Create the contest service
	contestService := services.NewContestService(contestRepo, studentRepo)

	// Create the Fiber app
	app := fiber.New()

	// Register the student handler
	studentHandler := handlers.NewStudentHandler(studentService)
	userHandler := handlers.NewUserHandler(oauthConfig, userService)

	// Register the student handler
	contestHandler := handlers.NewContestHandler(contestService)

	// Register the student routes
	studentHandler.RegisterRoutes(app)
	userHandler.RegisterRoutes(app)
	contestHandler.RegisterRoutes(app)

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
