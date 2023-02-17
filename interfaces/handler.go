package interfaces

import (
	"bufio"
	"encoding/csv"
	"fiber-base-go/config"
	"fiber-base-go/domain"
	"fiber-base-go/infrastructure/persistence"
	"fiber-base-go/services"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"io"
	"log"
	"net/http"
	"time"
)

type StudentHandler struct {
	services services.StudentService
}

// Run start server
func Run(port int) {
	app := fiber.New()
	conn, _ := config.ConnectDb()

	Migrate(conn)

	SetupRoutes(app, conn)

	log.Printf("Server running at http://localhost:%d/", port)

	app.Listen(":3000")
}

// Routes returns the initialized router
func SetupRoutes(app *fiber.App, db *gorm.DB) {
	repo := &persistence.StudentRepository{Conn: db}
	service := &services.StudentService{Repo: *repo}
	h := &StudentHandler{services: *service}

	app.Get("/", h.ListStudents)

	app.Post("/fact", h.CreateStudent)

	app.Get("/upload", h.Upload)
}

func (s *StudentHandler) ListStudents(c *fiber.Ctx) error {
	facts, err := s.services.GetAllStudents()
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

	err := s.services.AddStudent(fact)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(200).JSON(fact)
}

func (s *StudentHandler) Upload(c *fiber.Ctx) error {
	// Read the file from the request body
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Error reading file from request body",
		})
	}
	// Open the file and create a new reader
	csvFile, err := file.Open()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error opening CSV file",
		})
	}
	defer csvFile.Close()
	reader := csv.NewReader(bufio.NewReader(csvFile))
	// Read the CSV records one by one and create new student records
	var students []domain.Student
	for {
		// Read the next record from the CSV file
		record, err := reader.Read()
		if err == io.EOF {
			// Reached end of file
			break
		}
		if err != nil {
			// Error reading record
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"message": "Error reading CSV record",
			})
		}
		// Parse the student data from the record
		birthday, err := time.Parse("2006-01-02", record[1])
		if err != nil {
			// Error parsing date
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid date format in CSV record",
			})
		}
		student := domain.Student{
			Name:     record[0],
			Birthday: &birthday,
			Class:    record[2],
		}
		// Add the new student to the list
		students = append(students, student)
	}
	// Save the new student records to the database
	for _, student := range students {
		result := s.services.Repo.Create(&student)
		if result.Error != nil {
			// Error saving record to database
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"message": "Error saving student record",
			})
		}
	}
	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"message": fmt.Sprintf("Created %d new student records", len(students)),
	})
}

// =============================
//    MIGRATE
// =============================

func Migrate(conn *gorm.DB) error {
	err := config.DBMigrate(conn)
	if err != nil {
		return err
	}

	return nil
}
