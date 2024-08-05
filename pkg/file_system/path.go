package file_system

import (
	"io/fs"
	"os"
)

func IsDirectory(path string) (bool, *os.File, error) {
	var file *os.File = nil
	var fileStat fs.FileInfo = nil
	var err error = nil

	if file, err = os.Open(path); err != nil {
		return false, nil, err
	}
	if fileStat, err = file.Stat(); err != nil {
		file.Close()
		return false, nil, err
	}
	return fileStat.IsDir(), file, nil
}
