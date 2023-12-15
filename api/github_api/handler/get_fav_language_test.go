package handler

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	mock_github "dev/github-fav-language/clients/github/mock"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_GetFavLanguageClientFailure(t *testing.T) {
	// Setup
	ctrl := gomock.NewController(t)
	mockGithubClient := mock_github.NewMockClient(ctrl)
	mockGithubClient.EXPECT().GetFavLanguage(gomock.Any(), "antoniodaria").Return("", assert.AnError)

	app := fiber.New()
	_ = NewHandler(app, mockGithubClient)

	// Test
	url := "/github/fav-language/antoniodaria"
	req := httptest.NewRequest(http.MethodGet, url, nil)
	resp, _ := app.Test(req)
	defer resp.Body.Close()

	// Assertions
	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
}

func Test_GetFavLanguageSuccess(t *testing.T) {
	// Setup
	ctrl := gomock.NewController(t)
	mockGithubClient := mock_github.NewMockClient(ctrl)
	app := fiber.New()

	_ = NewHandler(app, mockGithubClient)

	// Expectations
	mockGithubClient.EXPECT().GetFavLanguage(gomock.Any(), "antoniodaria").Return("Go", nil)

	// Test
	url := "/github/fav-language/antoniodaria"
	req := httptest.NewRequest(http.MethodGet, url, nil)
	resp, _ := app.Test(req)
	defer resp.Body.Close()

	// convert response body to byte
	resBody, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)

	// Assertions
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, `{"user_name":"antoniodaria","most_used_language":"Go"}`, string(resBody))
}
