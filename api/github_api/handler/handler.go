package handler

import (
	"dev/github-fav-language/clients/github"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	GithubClient github.Client
}

func NewHandler(app fiber.Router, githubClient github.Client) *Handler {
	handler := &Handler{
		GithubClient: githubClient,
	}
	handler.connectRoutes(app)
	return handler
}

func (h *Handler) connectRoutes(app fiber.Router) {
	githubGroup := app.Group("/github")

	githubGroup.Get("/fav-language/:userName", h.GetFavLanguage)
}
