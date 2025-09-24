package utils

import (
	"path/filepath"
	"runtime"
)

func GetErrorLocation() (string, int) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "unknown"
		line = 0
	}
	file = filepath.Base(file)

	return file, line
}
