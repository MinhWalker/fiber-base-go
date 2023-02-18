package interfaces

import (
	"encoding/csv"
	"fiber-base-go/config"
	"fiber-base-go/domain"
	"fiber-base-go/services"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"io"
	"net/http"
	"time"
)

type StudentHandler struct {
	Services *services.StudentService
}

func (s *StudentHandler) ListStudents(c *fiber.Ctx) error {
	facts, err := s.Services.GetAllStudents()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(200).JSON(facts)
}

func (s *StudentHandler) CreateStudent(c *fiber.Ctx) error {
	var student domain.Student
	if err := c.BodyParser(&student); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	err := s.Services.AddStudent(student)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(200).JSON(student)
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

	r := csv.NewReader(csvFile)
	headers, err := r.Read()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	var students []*domain.Student
	for {
		row, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		student := domain.Student{}
		for i, header := range headers {
			switch header {
			case "Name":
				student.Name = row[i]
				fmt.Println("DEBUGGGG name", row[i])
			case "Class":
				student.Class = row[i]
				fmt.Println("DEBUGGGG Class", row[i])
			case "Birthday":
				birthday, err := time.Parse("2006-01-02", row[i])
				if err != nil {
					return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
				}
				student.Birthday = birthday
				fmt.Println("DEBUGGGG Class", birthday)
			}
		}
		students = append(students, &student)
	}

	if err := s.Services.UpsertStudent(students); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": ""})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Students imported successfully"})
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
