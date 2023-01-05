package http_server_mock

type Routes []Route

type Route struct {
	URL        string
	Method     string
	StatusCode int
	Body       []byte
	Filepath   string
	Errors     []error
}
