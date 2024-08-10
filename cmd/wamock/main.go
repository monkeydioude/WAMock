package main

import (
	"log"
	"os"
	"time"
	"wamock/internal/config"
	"wamock/internal/handler"

	"github.com/monkeydioude/moon"
)

func loadConfig(mh *moon.Handler, args []string) config.RunConfig {
	conf := config.RetrieveStartingConf(args)
	routes, err := config.Parse(conf)
	if err != nil {
		log.Fatal(err)
	}
	mh.MakeRouter(handler.GenerateRoutes(routes)...)
	return conf
}

func triggerCoroutine(conf config.RunConfig, mh *moon.Handler, args []string) {
	if conf.CoroutineRefreshTimer() == 0 {
		return
	}
	for {
		done := make(chan struct{})
		go func() {
			defer func() {
				if r := recover(); r != nil {
					log.Printf("[ERR ] Config auto reload panicked: %s. Relaunching it\n", r)
					close(done)
				}
			}()
			log.Printf("[INFO] Starting config auto reload every %s\n", time.Duration(conf.CoroutineRefreshTimer())*time.Second)
			for {
				time.Sleep(time.Duration(conf.CoroutineRefreshTimer()) * time.Second)
				loadConfig(mh, args)
			}
		}()
		<-done
	}
}

func main() {
	server := moon.Moon()
	conf := loadConfig(server, os.Args)
	go triggerCoroutine(conf, server, os.Args)
	moon.ServerRun(conf.GetServerAddr(), server)
}
