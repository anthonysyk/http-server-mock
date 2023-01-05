package routehelper

import (
	"fmt"
	"github.com/gorilla/mux"
	"math/rand"
	"path"
	"reflect"
	"regexp"
	"runtime"
	"strings"
	"time"
)

func PrintRoutes(router *mux.Router) {
	router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		pathTemplate, err := route.GetPathTemplate()
		if err == nil {
			methods, _ := route.GetMethods()
			handlerName := runtime.FuncForPC(reflect.ValueOf(route.GetHandler()).Pointer()).Name()
			for _, m := range methods {
				fmt.Printf("[%s] %s - %s\n", m, pathTemplate, path.Base(handlerName))
			}
		}
		return err
	})
}

func ReplaceQueryParamWithID(s string) string {
	rand.Seed(time.Now().UnixNano())
	re := regexp.MustCompile(`\{[^}]+\}`)
	return re.ReplaceAllStringFunc(s, func(match string) string {
		return fmt.Sprintf("%d", rand.Intn(1000000)+1)
	})
}

func ReplaceQueryParamsPlaceholdersWithValues(s string, values map[string]string) string {
	re := regexp.MustCompile(`\{[^}]+\}`)
	return re.ReplaceAllStringFunc(s, func(match string) string {
		key := strings.Replace(match, "{", "", 1)
		key = strings.Replace(key, "}", "", 1)
		return fmt.Sprintf(values[key])
	})
}
