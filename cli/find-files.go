/*
Copyright © 2021 Glenn M. Lewis

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cli

import (
	"io/fs"
	"log"
	"path/filepath"
	"strings"
)

func findDartFiles(path string, excludeFiles []string, quiet bool) []string {
	if !quiet && len(excludeFiles) > 0 {
		log.Printf("excluding files: %+v", excludeFiles)
	}
	var result []string

	excluded := func(path string) bool {
		for _, pattern := range excludeFiles {
			if strings.HasPrefix(pattern, "**/") {
				base := filepath.Base(path)
				match, err := filepath.Match(pattern[3:], base)
				if err != nil {
					log.Fatalf("bad glob pattern %q: %v", pattern[3:], err)
				}
				if match {
					if !quiet {
						log.Printf("Skipping file %q due to exclude pattern %q ...", path, pattern)
					}
					return true
				}
				continue
			}

			match, err := filepath.Match(pattern, path)
			if err != nil {
				log.Fatalf("bad glob pattern %q: %v", pattern, err)
			}
			if match {
				if !quiet {
					log.Printf("Skipping file %q due to exclude pattern %q ...", path, pattern)
				}
				return true
			}
		}

		return false
	}

	f := func(path string, d fs.DirEntry, err error) error {
		if strings.HasSuffix(path, ".dart") && !excluded(path) {
			result = append(result, path)
		}
		return nil
	}

	if err := filepath.WalkDir(path, f); err != nil {
		log.Fatalf("WalkDir: %v", err)
	}

	return result
}
