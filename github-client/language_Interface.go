package githubclient

import "context"

type LanguageInterface interface {
	GetFavLanguage(ctx context.Context, userName string) (string, error)
}
