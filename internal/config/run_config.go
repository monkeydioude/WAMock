package config

import (
	"flag"
	"log"
	"os"
	"wamock/pkg/file_system"
)

type RunConfig struct {
	file                  *os.File
	confPath              string
	isDirectory           bool
	coroutineRefreshTimer uint8
}

func New(confPath string, coroutineRefreshTimer uint8) RunConfig {
	isDirectory, file, err := file_system.IsDirectory(confPath)
	if err != nil {
		log.Fatalf("RunConfig.New: %s\n", err.Error())
	}
	return RunConfig{
		file:                  file,
		confPath:              confPath,
		isDirectory:           isDirectory,
		coroutineRefreshTimer: coroutineRefreshTimer,
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

func (r RunConfig) CoroutineRefreshTimer() uint8 {
	return r.coroutineRefreshTimer
}

func RetrieveStartingConf(args []string) RunConfig {
	if len(args) < 2 {
		log.Fatal("Missing path to config file or config file directory")
	}
	confPath := args[1]
	flagSet := flag.NewFlagSet("custom", flag.ExitOnError)
	crRefrTimer := flagSet.Int("x", 0, "-x <coroutine refresh timer>")
	flagSet.Parse(args[2:])
	return New(confPath, uint8(*crRefrTimer))
}
