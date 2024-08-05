package config

import (
	"flag"
	"fmt"
	"log"
	"os"
	"wamock/pkg/file_system"
)

type RunConfig struct {
	file                  *os.File
	confPath              string
	isDirectory           bool
	coroutineRefreshTimer uint16
	port                  uint16
}

func New(confPath string, coroutineRefreshTimer, port uint16) RunConfig {
	isDirectory, file, err := file_system.IsDirectory(confPath)
	if err != nil {
		log.Fatalf("RunConfig.New: %s\n", err.Error())
	}
	return RunConfig{
		file:                  file,
		confPath:              confPath,
		isDirectory:           isDirectory,
		coroutineRefreshTimer: coroutineRefreshTimer,
		port:                  port,
	}
}

func (r RunConfig) ConfPath() string {
	return r.confPath
}

func (r RunConfig) IsConfigPathDirectory() bool {
	return r.isDirectory
}

func (r RunConfig) ShouldUseCoroutine() bool {
	return r.coroutineRefreshTimer > 0
}

func (r RunConfig) CoroutineRefreshTimer() uint16 {
	return r.coroutineRefreshTimer
}

func (r RunConfig) GetServerAddr() string {
	return fmt.Sprintf("0.0.0.0:%d", r.port)
}

func RetrieveStartingConf(args []string) RunConfig {
	if len(args) < 2 {
		log.Fatal("Missing path to config file or config file directory")
	}
	confPath := args[1]
	flagSet := flag.NewFlagSet("custom", flag.ExitOnError)
	crRefrTimer := flagSet.Int("x", 0, "-x <coroutine refresh timer>")
	port := flagSet.Int("p", 8088, "-p <port>")
	flagSet.Parse(args[2:])
	return New(confPath, uint16(*crRefrTimer), uint16(*port))
}
