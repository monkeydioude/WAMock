package main

import (
	"log"
	"os"
	"wamock/internal/config"
	"wamock/internal/handler"

	"github.com/monkeydioude/moon"
)

func main() {
	conf := config.RetrieveStartingConf(os.Args)
	// routes := make(map[string]any, 0)
	routes, err := config.Parse(conf)
	if err != nil {
		log.Fatal(err)
	}
	server := moon.Moon()
	server.MakeRouter(handler.GenerateRoutes(routes)...)
	moon.ServerRun(conf.GetServerAddr(), server)
}
