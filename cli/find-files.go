package cli

import (
	"io/fs"
	"log"
	"path/filepath"
	"strings"
)

func findDartFiles(path string) []string {
	var result []string

	f := func(path string, d fs.DirEntry, err error) error {
		if strings.HasSuffix(path, ".dart") {
			result = append(result, path)
		}
		return nil
	}

	if err := filepath.WalkDir(path, f); err != nil {
		log.Fatalf("WalkDir: %v", err)
	}

	return result
}
