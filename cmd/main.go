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

	// Create the student repository
	studentRepo := repository.NewStudentRepository(conn)

	// Create the student service
	studentService := services.NewStudentService(studentRepo)

	// Create the contest repository
	contestRepo := repository.NewContestRepository(conn)

	// Create the contest service
	contestService := services.NewContestService(contestRepo, studentRepo)

	roomRepo := repository.NewRoomRepository(conn)
	roomService := services.NewRoomService(roomRepo)

	go func() {
		populateService := services.NewPopulateService(roomService, studentService)
		if err := populateService.Populate(); err != nil {
			log.Fatalf("failed to populate: %s", err)
		}
	}()

	// Create the Fiber app
	app := fiber.New()

	// Register the student handler
	studentHandler := handlers.NewStudentHandler(studentService)

	// Register the student handler
	contestHandler := handlers.NewContestHandler(contestService)

	// Register the student routes
	studentHandler.RegisterRoutes(app)
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
