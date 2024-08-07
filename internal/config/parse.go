package config

import (
	"encoding/json"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"wamock/internal/routing"
	"wamock/pkg/file_system"
)

const FILE_SLICE_SIZE int = 1

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
	return routes, nil
}

// parseDirChunk reads json config files from a dirPath directory,
// and try to convert their path to a apiMock endpoint route.
// Then, the content of the json config files and unmarshall
// their content to match a route.
func parseDirChunk(
	dirPath string,
	files []fs.FileInfo,
	routes map[string]*routing.Route,
) error {
	for _, f := range files {
		route := routing.Route{}
		file, err := os.Open(dirPath + "/" + f.Name())
		if err != nil {
			return err
		}
		b, err := io.ReadAll(file)
		if err != nil {
			return err
		}
		if err = json.Unmarshal(b, &route); err != nil {
			return err
		}
		apiMockPath := file_system.CleanConfigFilename(f.Name())
		route.HydrateWithPath(apiMockPath)
		routes[apiMockPath] = &route
	}
	return nil
}

// ParseDirectory reads a directory and try to convert each config file
// into a mock api endpoint.
func ParseDirectory(file *os.File) (map[string]*routing.Route, error) {
	routes := make(map[string]*routing.Route, 0)
	files, err := file.Readdir(FILE_SLICE_SIZE)
	if err != nil {
		return nil, err
	}
	lenFiles := len(files)
	dirPath, err := filepath.Abs(file.Name())
	if err != nil {
		return nil, err
	}
	for lenFiles > 0 {
		if err = parseDirChunk(dirPath, files, routes); err != nil {
			log.Println(err)
			continue
		}
		if lenFiles < FILE_SLICE_SIZE {
			break
		}
		files, _ = file.Readdir(FILE_SLICE_SIZE)
		lenFiles = len(files)
	}
	return routes, nil
}

func Parse(config RunConfig) (map[string]*routing.Route, error) {
	defer config.file.Close()
	if config.isDirectory {
		return ParseDirectory(config.file)
	}
	return ParseSingleFile(config.file)
}
