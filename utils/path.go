package utils

import (
	"os"
	"path/filepath"
	"strings"
)

// return a path is directory
func IsDir(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}
	return fileInfo.IsDir()
}

// get file type by extension
func FileTypeByExt(path string) string {
	return strings.TrimLeft(filepath.Ext(path), ".")
}
