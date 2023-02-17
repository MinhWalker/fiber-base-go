package interfaces

import (
	"fiber-base-go/application"
	"fiber-base-go/config"
	"fiber-base-go/domain"
	"github.com/gofiber/fiber/v2"
	"log"
)

// Run start server
func Run(port int) {
	app := fiber.New()
	SetupRoutes(app)

	log.Printf("Server running at http://localhost:%d/", port)

	app.Listen(":3000")
}

// Routes returns the initialized router
func SetupRoutes(app *fiber.App) {
	app.Get("/", ListStudents)

	app.Post("/fact", CreateStudent)

	app.Get("/migrate", Migrate)
}

func ListStudents(c *fiber.Ctx) error {
	facts, err := application.GetAllStudents()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(200).JSON(facts)
}

func CreateStudent(c *fiber.Ctx) error {
	var fact domain.Student
	if err := c.BodyParser(fact); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	err := application.AddStudent(fact)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(200).JSON(fact)
}

// =============================
//    MIGRATE
// =============================

func Migrate(c *fiber.Ctx) error {
	_, err := config.DBMigrate()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(200).Next()
}
