package http_server_mock

import (
	"errors"
	"gopkg.in/yaml.v3"
	"os"
)

func GetRoutes(config string) (Routes, error) {
	m := make(map[interface{}]interface{})
	err := yaml.Unmarshal([]byte(config), &m)
	if err != nil {
		return nil, err
	}
	paths := getPaths(m)
	var routes Routes
	for _, ut := range paths {
		methodTails := getMethod(ut.Tail)
		for _, mt := range methodTails {
			statusTails := getStatus(mt.Tail)
			for _, st := range statusTails {
				filepath, body, bodyErr := getBody(st.Tail)
				routes = append(routes, Route{
					URL:        ut.URL,
					Method:     mt.Method,
					StatusCode: st.Status,
					Body:       body,
					Filepath:   filepath,
					Errors:     []error{bodyErr},
				})
			}
		}
	}
	return routes, nil
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

func getMethod(m map[string]interface{}) []MethodTail {
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

func getBody(m map[string]interface{}) (string, []byte, error) {
	if len(m) > 1 {
		return "", nil, errors.New("cannot set both body and filepath for a specific status code")
	}
	if body, ok := m["body"]; ok && body != "" {
		return "", []byte(body.(string)), nil
	}
	if filepath, ok := m["filepath"]; ok && filepath != "" {
		fp := filepath.(string)
		content, err := os.ReadFile(fp)
		if err != nil {
			return fp, nil, err
		}
		return fp, content, nil
	}

	return "", nil, errors.New("no body found")
}
