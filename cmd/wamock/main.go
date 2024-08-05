package main

import (
	"encoding/json"
	"log"
	"os"
	"wamock/internal/config"
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

func generate_routes(routes map[string]*routing.Route) []*moon.Route {
	sliceRoutes := make([]*moon.Route, 0)
	for _, r := range routes {
		moonRoute := moon.NewRoute(r.Path, r.Method.String(), handle(r))
		sliceRoutes = append(sliceRoutes, moonRoute)
	}
	return sliceRoutes
}

func main() {
	conf := config.RetrieveStartingConf(os.Args)
	// routes := make(map[string]any, 0)
	routes, err := config.Parse(conf)
	if err != nil {
		log.Fatal(err)
	}
	server := moon.Moon()
	server.MakeRouter(generate_routes(routes)...)
	moon.ServerRun("0.0.0.0:8088", server)
}
