package handlers

import (
	"context"
	"encoding/json"
	"fiber-base-go/internal/model"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/oauth2"
)

var oauthStateString = "pseudo-random"

type userHandler struct {
	conf *oauth2.Config
}

func NewUserHandler(conf *oauth2.Config) *userHandler {
	return &userHandler{
		conf: conf,
	}
}

func (h *userHandler) RegisterRoutes(app *fiber.App) {
	api := app.Group("/api/v1") // Prefix API version
	api.Get("/oauth", h.googleLogin)
	api.Get("/oauth/callback", h.callback)
}

func (h *userHandler) googleLogin(ctx *fiber.Ctx) error {
	url := h.conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
	return ctx.Redirect(url)
}

func (h *userHandler) callback(ctx *fiber.Ctx) error {
	//content, err := h.getUserInfo(ctx.Query("state"), ctx.Query("code"))
	//if err != nil {
	//	return ctx.SendStatus(http.StatusInternalServerError)
	//}
	//
	//var user model.User
	//h.conf.Clauses(clause.OnConflict{
	//	Columns:   []clause.Column{{Name: "email"}},
	//	DoUpdates: clause.AssignmentColumns([]string{"updated_at"}),
	//}).FirstOrCreate(&user, User{Email: content.Email})
	//
	//return ctx.JSON(user)
	return nil
}

func (h *userHandler) getUserInfo(state string, code string) (*model.User, error) {
	if state != oauthStateString {
		return nil, fmt.Errorf("invalid oauth state")
	}

	token, err := h.conf.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("failed to exchange token: %s", err.Error())
	}

	client := h.conf.Client(context.Background(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo?alt=json")
	if err != nil {
		return nil, fmt.Errorf("failed to get user info: %s", err.Error())
	}
	defer resp.Body.Close()

	var userInfo model.User
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return nil, fmt.Errorf("failed to decode user info: %s", err.Error())
	}

	return &userInfo, nil
}