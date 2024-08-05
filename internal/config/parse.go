package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"wamock/internal/routing"
)

func ParseSingleFile(file *os.File) (map[string]*routing.Route, error) {
	routes := make(map[string]*routing.Route, 0)
	b, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(b, &routes); err != nil {
		return nil, err
	}
	for path, route := range routes {
		route.HydrateWithPath(path)
	}
	fmt.Println(routes)
	return routes, nil
}

func Parse(config RunConfig) (map[string]*routing.Route, error) {
	defer config.file.Close()
	if config.isDirectory {
		// @todo implements
		return nil, nil
	}
	return ParseSingleFile(config.file)
}
