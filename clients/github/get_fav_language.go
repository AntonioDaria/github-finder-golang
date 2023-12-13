package github

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Repo struct {
	Language string
}

func (c *ClientImpl) GetFavLanguage(ctx context.Context, userName string) (string, error) {

	url := fmt.Sprintf(`https://api.github.com/users/%s/repos`, userName)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		fmt.Printf("Failed to generate a request %v", err)
		return "", err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		fmt.Printf("Failed to retrieve language %v", err)
		return "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Failed to read response body %v", err)
		return "", err
	}

	if res.StatusCode != 200 {
		fmt.Printf("Failed to read response body %v", err)
		return "", err
	}

	var response []Repo
	jsonErr := json.Unmarshal(body, &response)
	if jsonErr != nil {
		fmt.Printf("Failed to unmarshal response body %v", err)
		return "", err
	}

	return extractFavoriteLanguage(response), nil
}

func extractFavoriteLanguage(repos []Repo) string {
	var value string
	var maxCount int
	languageMap := make(map[string]int)

	for _, language := range repos {
		if languageMap[language.Language] == 0 {
			languageMap[language.Language] = 1
		} else {
			languageMap[language.Language]++
		}
		if languageMap[language.Language] > maxCount {
			value = language.Language
			maxCount = languageMap[language.Language]
		}
	}

	return value
}
