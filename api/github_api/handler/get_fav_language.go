package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type GetFavLanguageResponse struct {
	UserName         string `json:"user_name"`
	MostUsedLanguage string `json:"most_used_language"`
}

func (h *Handler) GetFavLanguage(c *fiber.Ctx) error {
	userName := c.Params("userName")

	language, err := h.GithubClient.GetFavLanguage(c.Context(), userName)
	if err != nil {
		fmt.Printf("Failed to get favourite language %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.JSON(GetFavLanguageResponse{
		MostUsedLanguage: language,
		UserName:         userName,
	})
}
