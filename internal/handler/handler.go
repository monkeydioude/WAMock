package handler

import (
	"encoding/json"
	"wamock/internal/routing"

	"github.com/monkeydioude/moon"
)

func handle(route *routing.Route) func(*moon.Request) ([]byte, int, error) {
	return func(req *moon.Request) ([]byte, int, error) {
		b, err := json.Marshal(route.Response)
		if err != nil {
			return nil, 500, err
		}
		return b, 200, nil
	}
}

func GenerateRoutes(routes map[string]*routing.Route) []*moon.Route {
	sliceRoutes := make([]*moon.Route, 0)
	for _, r := range routes {
		moonRoute := moon.NewRoute(r.Path, r.Method.String(), handle(r))
		sliceRoutes = append(sliceRoutes, moonRoute)
	}
	return sliceRoutes
}
