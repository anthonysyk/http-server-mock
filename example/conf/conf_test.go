package conf

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
	"testing"

	_ "embed"
)

//go:embed routes.yml
var RoutesYAML string

func TestModel(t *testing.T) {
	m := make(map[interface{}]interface{})
	err := yaml.Unmarshal([]byte(RoutesYAML), &m)
	assert.NoError(t, err)
	paths := getPaths(m)
	for _, ut := range paths {
		//fmt.Printf("%+v\n", ut)
		methodTails := getMethod(ut.URL, ut.Tail)
		for _, mt := range methodTails {
			//fmt.Printf("%+v\n", mt)
			statusTails := getStatus(mt.Tail)
			for _, st := range statusTails {
				fmt.Printf("%s, %s, %v, tail : %+v\n", ut.URL, mt.Method, st.Status, st.Tail)
			}
		}
	}
}

type URLTail struct {
	URL  string
	Tail map[string]interface{}
}

func getPaths(m map[interface{}]interface{}) []URLTail {
	var URLTails []URLTail
	for url, path := range m["paths"].(map[string]interface{}) {
		URLTails = append(URLTails, URLTail{URL: url, Tail: path.(map[string]interface{})})
	}
	return URLTails
}

type MethodTail struct {
	Method string
	Tail   map[interface{}]interface{}
}

func getMethod(url string, m map[string]interface{}) []MethodTail {
	var methodTails []MethodTail
	for k, v := range m {
		methodTail := MethodTail{
			Method: k,
			Tail:   v.(map[interface{}]interface{}),
		}
		methodTails = append(methodTails, methodTail)
	}

	return methodTails
}

type StatusTail struct {
	Status int
	Tail   map[string]interface{}
}

func getStatus(m map[interface{}]interface{}) []StatusTail {
	var statusTails []StatusTail
	for k, v := range m {
		if _, ok := k.(int); !ok {
			panic("status code must be an integer")
		}
		statusTails = append(statusTails, StatusTail{
			Status: k.(int),
			Tail:   v.(map[string]interface{}),
		})
	}

	return statusTails
}

func getBody(url, method string, statusCode int, m map[string]interface{}) ([]byte, error) {
	if len(m) > 1 {
		return nil, errors.New("cannot set multiple bodies for a specific status code")
	}
	for _, t := range m {
		return t.([]byte), nil
	}
	return nil, errors.New("no body found")
}

func getFilepath(m map[string]interface{}) (string, error) {
	if len(m) > 1 {
		return "nil", errors.New("cannot set multiple filenames for a specific status code")
	}
	for _, t := range m {
		return t.(string), nil
	}
	return "", errors.New("no filename found")
}

type Routes []Route
type Route struct {
	URL        string
	Method     string
	StatusCode int
	Body       []byte
	Filepath   string
}
