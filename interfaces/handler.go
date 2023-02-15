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
	app.Get("/", ListFacts)

	app.Post("/fact", CreateFact)

	app.Get("/migrate", Migrate)
}

func ListFacts(c *fiber.Ctx) error {
	facts, err := application.GetAllFacts()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(200).JSON(facts)
}

func CreateFact(c *fiber.Ctx) error {
	var fact domain.Fact
	if err := c.BodyParser(fact); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	err := application.AddFact(fact)
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