package http_server_mock

import (
	"fmt"
	"github.com/anthonysyk/http-server-mock/pkg/routehelper"
	"github.com/gorilla/mux"
	"net/http"
)

func GenerateRouter(config string) (*mux.Router, error) {
	router := mux.NewRouter()

	routes, err := GetRoutes(config)
	if err != nil {
		return nil, err
	}

	for _, r := range routes {
		router.HandleFunc(r.URL, handler(r)).Methods(r.Method)
	}

	routehelper.PrintRoutes(router)
	return router, nil
}

func handler(route Route) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		content, err := route.GetBody(mux.Vars(r))
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(route.StatusCode)
		fmt.Fprintf(w, string(content))
	}
}
