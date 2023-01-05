package routehelper

import (
	"github.com/anthonysyk/http-server-mock/pkg/stdout"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestPrintRoutes(t *testing.T) {
	router := mux.NewRouter()
	router.HandleFunc("/route1", func(w http.ResponseWriter, r *http.Request) {}).Methods("GET")
	router.HandleFunc("/route2", func(w http.ResponseWriter, r *http.Request) {}).Methods("POST")

	output := stdout.Record(func() { PrintRoutes(router) })

	expected := `[GET] /route1 - routehelper.TestPrintRoutes.func1
[POST] /route2 - routehelper.TestPrintRoutes.func2
`
	assert.Equal(t, expected, output)
}

func TestReplaceQueryParamWithID(t *testing.T) {
	url := "/movies/{id}/actors/{id}"
	result := ReplaceQueryParamWithID(url)
	assert.NotEqual(t, url, result)
}

func TestReplaceQueryParamsPlaceholdersWithValues(t *testing.T) {
	url := "/movies/{movieId}/actors/{actorId}"
	path := "payload/movies/{movieId}/actors/{actorId}.json"

	values := make(map[string]string)
	values["movieId"] = "123"
	values["actorId"] = "456"

	resultURL := ReplaceQueryParamsPlaceholdersWithValues(url, values)
	assert.Equal(t, "/movies/123/actors/456", resultURL)

	resultPath := ReplaceQueryParamsPlaceholdersWithValues(path, values)
	assert.Equal(t, "payload/movies/123/actors/456.json", resultPath)
}
