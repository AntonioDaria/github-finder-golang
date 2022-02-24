package githubclient

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pulumi/pulumi/sdk/v3/go/common/util/logging"
)

func (c *LangugeClient) GetFavLanguage(ctx context.Context, userName string) (string, error) {

	url := fmt.Sprintf(`https://api.github.com/users/%s/repos`, userName)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		logging.Errorf("Failed to generate a request %v", err)
		return "", err
	}

	req = req.WithContext(ctx)
	res, err := c.Doer.Do(req)
	fmt.Println("xxx", res)

	if err != nil {
		logging.Errorf("Failed to retrieve language %v", err)
		return "", err
	}

	if res.StatusCode != 200 {
		err := fmt.Errorf("error returned from github api")
		logging.Errorf("Github api error: %v", err)
		return "", err
	}
	var response []Repo
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		logging.Errorf("Can not decode JSON %v", err)
		return "", err
	}
	//fmt.Printf("%+v\n", response)
	return extractFavouriteLanguage(response), nil
}

type Repo struct {
	Language string
}

func extractFavouriteLanguage(repos []Repo) string {
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
