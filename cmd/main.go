package main

import (
	"fiber-base-go/config"
	"fiber-base-go/infrastructure/persistence"
	"fiber-base-go/interfaces"
	"fiber-base-go/services"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"log"
)

func main() {
	Run(3000)
}

// Run start server
func Run(port int) {
	app := fiber.New()
	conn, _ := config.ConnectDb()

	interfaces.Migrate(conn)

	SetupRoutes(app, conn)

	log.Printf("Server running at http://localhost:%d/", port)

	app.Listen(":3000")
}

// Routes returns the initialized router
func SetupRoutes(app *fiber.App, db *gorm.DB) {
	repo := &persistence.StudentRepository{Conn: db}
	service := &services.StudentService{Repo: repo}
	h := interfaces.StudentHandler{Services: service}

	app.Get("/", h.ListStudents)

	app.Post("/student", h.CreateStudent)

	app.Post("/upload", h.Upload)
}
