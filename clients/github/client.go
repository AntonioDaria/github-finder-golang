package github

import (
	"context"
	"net/http"
)

type Client interface {
	GetFavLanguage(ctx context.Context, userName string) (string, error)
}

type ClientImpl struct {
	HTTPClient *http.Client
}

func NewClient(httpClient *http.Client) *ClientImpl {
	return &ClientImpl{
		HTTPClient: httpClient,
	}
}
