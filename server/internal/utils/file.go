// Package utils contain all utility functions
package utils

import "os"

func CheckAndCreateDirectoryIfNotExists(dir string) {
	_, err := os.Stat(dir)
	if err == nil {
		return
	}

	os.Mkdir(dir, os.ModePerm)
}
