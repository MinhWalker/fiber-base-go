package handlers

import (
	"fmt"
	"gorm.io/gorm"
	"net/http"
	"strconv"

	"fiber-base-go/internal/delivery/api/request"
	"fiber-base-go/internal/model"
	"fiber-base-go/internal/services"
	"fiber-base-go/internal/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)

type studentHandler struct {
	svc services.StudentService
}

func NewStudentHandler(svc services.StudentService) *studentHandler {
	return &studentHandler{
		svc: svc,
	}
}

func (h *studentHandler) RegisterRoutes(app *fiber.App) {
	api := app.Group("/api/v1") // Prefix API version
	api.Post("/students", h.createStudent)
	api.Get("/students/:id", h.getStudentByID)
	api.Put("/students/:id", h.updateStudent)
	api.Post("/students/import", h.importStudents)
}

func (h *studentHandler) importStudents(ctx *fiber.Ctx) error {
	file, err := ctx.FormFile("file")
	if err != nil {
		return errors.Wrap(err, "studentHandler.ImportStudents")
	}

	// Open the uploaded file
	src, err := file.Open()
	if err != nil {
		return errors.Wrap(err, "studentHandler.ImportStudents")
	}
	defer src.Close()

	// Parse the file
	batches, err := utils.ParseCSV(src, 100)
	if err != nil {
		return errors.Wrap(err, "studentHandler.ImportStudents")
	}

	// Save the students to the database
	totalCount := 0
	for _, batch := range batches {
		if err := h.svc.ImportStudent(batch); err != nil {
			return errors.Wrap(err, "studentHandler.ImportStudents")
		}
		totalCount += len(batch)
	}

	return ctx.JSON(fiber.Map{
		"message": fmt.Sprintf("Imported %d students", totalCount),
	})
}

func (h *studentHandler) getStudentByID(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "invalid student ID")
	}

	student, err := h.svc.GetStudent(uint(id))
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, "failed to get student")
	}

	return ctx.JSON(student)
}

func (h *studentHandler) createStudent(ctx *fiber.Ctx) error {
	var student model.Student
	if err := ctx.BodyParser(&student); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	err := h.svc.CreateStudent(&student)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(200).JSON(student)
}

func (h *studentHandler) updateStudent(ctx *fiber.Ctx) error {
	idStr := ctx.Params("id")

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid student ID",
		})
	}

	var req request.UpdateStudentRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse request body",
		})
	}

	student, err := h.svc.GetStudent(uint(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "student not found",
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "cannot get student",
		})
	}

	if req.Name != "" {
		student.Name = req.Name
	}
	if req.Class != "" {
		student.Class = req.Class
	}
	if !req.Birthday.IsZero() {
		student.Birthday = req.Birthday
	}

	if err := h.svc.UpdateStudent(student); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "cannot update student",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(student)
}
