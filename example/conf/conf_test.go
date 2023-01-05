package conf

import (
	_ "embed"
	"fmt"
	"github.com/anthonysyk/http-server-mock"
	"github.com/anthonysyk/http-server-mock/pkg/routehelper"
	"github.com/anthonysyk/http-server-mock/pkg/stdout"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestRoutes(t *testing.T) {
	routes, err := http_server_mock.GetRoutes(RoutesYAML)
	assert.NoError(t, err)
	assert.Len(t, routes, 8)
}

func TestRouter(t *testing.T) {
	router, err := http_server_mock.GenerateRouter(RoutesYAML)
	assert.NoError(t, err)
	output := stdout.Record(func() { routehelper.PrintRoutes(router) })
	expectedOutput := `[GET] /movies/{id} - http-server-mock.handler.func1
[GET] /movies/{moviesId}/actors/{actorsId} - http-server-mock.handler.func1
[GET] /top-actors - http-server-mock.handler.func1
[DELETE] /genres - http-server-mock.handler.func1
[GET] /genres - http-server-mock.handler.func1
[POST] /genres - http-server-mock.handler.func1
[GET] /health - http-server-mock.handler.func1
[GET] /top-movies - http-server-mock.handler.func1
`
	assert.ElementsMatch(t, strings.Split(expectedOutput, "\n"), strings.Split(output, "\n"))
}

//go:embed routes.yml
var RoutesYAML string

func TestServer(t *testing.T) {
	router, err := http_server_mock.GenerateRouter(RoutesYAML)
	assert.NoError(t, err)
	go http.ListenAndServe(":8080", router)

	client := &http.Client{}

	tests := []struct {
		url                string
		method             string
		expectedStatusCode int
		expectedResponse   string
	}{
		{
			url:                "/health",
			method:             http.MethodGet,
			expectedStatusCode: 200,
			expectedResponse:   `{"status":"OK"}`,
		},
		{
			url:                "/genres",
			method:             http.MethodGet,
			expectedStatusCode: 200,
			expectedResponse:   `["Action","Adventure","Animation","Comedy","Crime","Drama","Family","Fantasy","Music","Science Fiction","Thriller"]`,
		},
		{
			url:                "/genres",
			method:             http.MethodPost,
			expectedStatusCode: 200,
			expectedResponse:   `{"message":"new genre added"}`,
		},
		{
			url:                "/genres",
			method:             http.MethodDelete,
			expectedStatusCode: 200,
			expectedResponse:   `{"message":"genre deleted"}`,
		},
		{
			url:                "/top-movies",
			method:             http.MethodGet,
			expectedStatusCode: 200,
			expectedResponse:   TopMoviesPayload,
		},
		{
			url:                "/top-actors",
			method:             http.MethodGet,
			expectedStatusCode: 200,
			expectedResponse:   TopActorsPayload,
		},
		{
			url:                "/movies/399566",
			method:             http.MethodGet,
			expectedStatusCode: 200,
			expectedResponse:   Movie399566Payload,
		},
		{
			url:                "/movies/399566/actors/15556",
			method:             http.MethodGet,
			expectedStatusCode: 200,
			expectedResponse:   Movie399566Actor15556Payload,
		},
		{
			url:                "/movies/suggest",
			method:             http.MethodGet,
			expectedStatusCode: 200,
			expectedResponse:   MovieSuggestPayload,
		},
		{
			url:                routehelper.ReplaceQueryParamWithID("/movies/{id}"),
			method:             http.MethodGet,
			expectedStatusCode: 404,
		},
		{
			url:                routehelper.ReplaceQueryParamWithID("/movies/{id}/actors/{id}"),
			method:             http.MethodGet,
			expectedStatusCode: 404,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%s%s", test.method, test.url), func(t *testing.T) {
			url := fmt.Sprintf("%s%s", "http://localhost:8080", test.url)
			req, reqErr := http.NewRequest(test.method, url, nil)
			if reqErr != nil {
				t.Fatal(reqErr)
			}
			res, resErr := client.Do(req)
			defer res.Body.Close()
			if resErr != nil {
				t.Fatal(resErr)
			}
			body, bodyErr := io.ReadAll(res.Body)
			if bodyErr != nil {
				t.Fatal(bodyErr)
			}

			assert.Equal(t, test.expectedStatusCode, res.StatusCode)
			assert.Equal(t, test.expectedResponse, string(body))
		})
	}

}
