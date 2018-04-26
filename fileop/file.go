package fileop

import (
	"os"
)

func IsFile(path string) bool {
	fi, err := os.Stat(path)

	if err != nil {
		return false
	}

	switch mode := fi.Mode(); {
	case mode.IsDir():
		return false
	case mode.IsRegular():
		return true
	}

	return false
}

func OnlyFiles(paths map[string]bool) (files []string) {
	for path, _ := range paths {
		if IsFile(path) {
			files = append(files, path)
		}
	}

	return files
}
