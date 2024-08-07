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
		if r.Method.Match(routing.ALL) {
			routing.ForEachMethod(func(m routing.Method) {
				sliceRoutes = append(sliceRoutes, moon.NewRoute(r.Path, m.String(), handle(r)))
			})
		} else if r.Method.Match(routing.NONE) {
			continue
		}
		sliceRoutes = append(sliceRoutes, moon.NewRoute(r.Path, r.Method.String(), handle(r)))
	}
	return sliceRoutes
}
