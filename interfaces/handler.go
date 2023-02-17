package interfaces

import (
	"fiber-base-go/application"
	"fiber-base-go/config"
	"fiber-base-go/domain"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"log"
	"strconv"
)

type StudentHandler struct {
	DB *gorm.DB
}

// Run start server
func Run(port int) {
	app := fiber.New()
	conn, _ := config.ConnectDb()

	SetupRoutes(app, conn)

	log.Printf("Server running at http://localhost:%d/", port)

	app.Listen(":3000")
}

// Routes returns the initialized router
func SetupRoutes(app *fiber.App, db *gorm.DB) {
	h := &StudentHandler{
		DB: db,
	}

	app.Get("/", h.ListStudents)

	app.Post("/fact", h.CreateStudent)

	app.Get("/migrate", h.Migrate)

	app.Get("/upload", h.Upload)
}

func (s *StudentHandler) ListStudents(c *fiber.Ctx) error {
	facts, err := application.GetAllStudents(s.DB)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(200).JSON(facts)
}

func (s *StudentHandler) CreateStudent(c *fiber.Ctx) error {
	var fact domain.Student
	if err := c.BodyParser(fact); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	err := application.AddStudent(s.DB, fact)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(200).JSON(fact)
}

func (s *StudentHandler) Upload(ctx *fiber.Ctx) error {
	id, err := strconv.ParseInt(ctx.Params("id"), 10, 64)
	if err != nil {
		return ctx.Status(422).JSON(fiber.Map{"errors": [1]string{"We were not able to process your expense"}})
	}
	file, err := ctx.FormFile("attachment")
	if err != nil {
		return ctx.Status(422).JSON(fiber.Map{"errors": [1]string{"We were not able upload your attachment"}})
	}
	ctx.SaveFile(file, fmt.Sprintf("./uploads/%s", file.Filename))
	var studentDB domain.Student
	s.DB.First(&studentDB, id)
	s.DB.Model(&studentDB).Update("attachment", file.Filename)
	return ctx.JSON(fiber.Map{"message": "Attachment uploaded successfully"})
}

// =============================
//    MIGRATE
// =============================

func (s *StudentHandler) Migrate(c *fiber.Ctx) error {
	err := config.DBMigrate(s.DB)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(200).Next()
}
