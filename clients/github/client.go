package github

import (
	"context"
	"net/http"
)

//go:generate mockgen -source=$GOFILE -destination=mock/$GOFILE
type Client interface {
	GetFavLanguage(ctx context.Context, userName string) (string, error)
}

type ClientImpl struct {
	HTTPClient *http.Client
	BaseURL    string
}

func NewClient(httpClient *http.Client, baseUrl string) *ClientImpl {
	return &ClientImpl{
		HTTPClient: httpClient,
		BaseURL:    baseUrl,
	}
}
