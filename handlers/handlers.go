package handlers

import githubclient "dev/github-fav-language/github-client"

type Handler struct {
	GithubClient githubclient.LanguageInterface
}
