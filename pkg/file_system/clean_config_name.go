package file_system

import "strings"

func CleanConfigFilename(name string) string {
	name = strings.ReplaceAll(name, ":", "/")
	return strings.Replace(name, ".json", "", 1)
}
