package githubclient

type LangugeClient struct {
	Doer httpRequestDoer
}

type ErrorResponse struct {
	Message string
}
