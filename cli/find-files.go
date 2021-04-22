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
