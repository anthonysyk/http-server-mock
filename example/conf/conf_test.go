package conf

import (
	"github.com/anthonysyk/http-server-mock"
	"github.com/stretchr/testify/assert"
	"testing"

	_ "embed"
)

//go:embed routes.yml
var RoutesYAML string

func TestModel(t *testing.T) {
	routes, err := http_server_mock.GetRoutes(RoutesYAML)
	assert.NoError(t, err)
	assert.Len(t, routes, 8)
}

// run server
// do some calls
// catch output stdout
// verify output is ok
