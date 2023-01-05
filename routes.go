package http_server_mock

import (
	"fmt"
	"github.com/anthonysyk/http-server-mock/pkg/routehelper"
	"os"
)

type Routes []Route

type Route struct {
	URL        string
	Method     string
	StatusCode int
	InlineBody []byte
	Filepath   string
	Errors     []error
}

func (r Route) GetStatusCode() int {
	return r.StatusCode
}

func (r Route) GetBody(pathParams map[string]string) ([]byte, error) {
	if r.InlineBody != nil {
		return r.InlineBody, nil
	}

	body, err := r.filepathBody(pathParams)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (r Route) filepathBody(pathParams map[string]string) ([]byte, error) {
	filepath := routehelper.ReplaceQueryParamsPlaceholdersWithValues(r.Filepath, pathParams)
	fmt.Println(filepath)
	content, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	return content, nil
}

// todo : path param non unique, method http exists and uppercase
func (rs Routes) Validate() []error {

	return nil
}
