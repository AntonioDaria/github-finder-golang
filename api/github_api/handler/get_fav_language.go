package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) GetFavLanguage(c *fiber.Ctx) error {
	language, err := h.GithubClient.GetFavLanguage(c.Context(), c.Params("userName"))
	if err != nil {
		fmt.Printf("Failed to get favourite language %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.JSON(language)
}
