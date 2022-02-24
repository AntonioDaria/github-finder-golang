package githubclient

import "net/http"

type httpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}
