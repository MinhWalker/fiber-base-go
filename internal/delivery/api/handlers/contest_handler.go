package handlers

import (
	"fiber-base-go/internal/delivery/api/request"
	"fiber-base-go/internal/model"
	"fiber-base-go/internal/services"

	"github.com/gofiber/fiber/v2"
)

type contestHandler struct {
	svc services.ContestService
}

func NewContestHandler(svc services.ContestService) *contestHandler {
	return &contestHandler{
		svc: svc,
	}
}

func (h *contestHandler) RegisterRoutes(api fiber.Router) {
	api.Post("/contests", h.CreateContest)
}

func (h *contestHandler) CreateContest(c *fiber.Ctx) error {
	// Parse request body
	req := new(request.CreateContestRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Bad Request",
		})
	}

	var contest *model.Contest
	var err error
	if contest, err = h.svc.CreateContest(req.Name, req.Grades, req.Classes, req.AllSchool); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(contest)
}
