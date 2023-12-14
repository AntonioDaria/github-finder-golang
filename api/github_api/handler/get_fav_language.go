package handler

import (
	"dev/github-fav-language/packages/api"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type GetFavLanguageResponse struct {
	UserName         string `json:"user_name"`
	MostUsedLanguage string `json:"most_used_language"`
}

// GetFavLanguage godoc
// @Summary Get favorite language
// @Description Get the User Github favorite language
// @Tags Github
// @Accept json
// @Produce json
// @Param userName path string true "User Name"
// @Success 200 {object} GetFavLanguageResponse
// @Failure 404 {object} api.JSONError
// @Failure 500 {object} api.JSONError
// @Router /v1/github/fav-language/{userName} [get]
func (h *Handler) GetFavLanguage(c *fiber.Ctx) error {
	userName := c.Params("userName")

	language, err := h.GithubClient.GetFavLanguage(c.Context(), userName)
	if err != nil {
		fmt.Printf("Failed to get favorite language %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(api.JSONError{
			Error: "Failed to get favorite language",
		})
	}

	return c.JSON(GetFavLanguageResponse{
		MostUsedLanguage: language,
		UserName:         userName,
	})
}
