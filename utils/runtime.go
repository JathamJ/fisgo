package utils

import (
	"os"
	"path/filepath"
)

//Get workspace dir path
func GetWorkSpacePath() (string, error) {
	path, err := os.Executable()
	if err != nil {
		return "", err
	}
	return filepath.Dir(path), nil
}
