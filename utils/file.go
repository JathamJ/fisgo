package utils

import (
	"os"
	"path/filepath"
)

// FileExist check if file exist
func FileExist(filePath string) bool {
	_, err := os.Stat(filePath)
	return err == nil || os.IsExist(err)
}

// CheckAndCreateFileDir will create dir if dir not exist
func CheckAndCreateFileDir(path string) error {
	dir := filepath.Dir(path)
	if FileExist(dir) {
		return nil
	}
	return os.MkdirAll(dir, os.ModePerm)
}
