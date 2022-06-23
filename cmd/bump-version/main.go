// -*- compile-command: "go test -v ./..."; -*-

// bump-version reads a file and looks for this pattern:
// const VERSION = "a.b.c"
// (on a line by itself) where "a", "b", and "c" are numbers.
// (The file can use single or double quotes.)
//
// It then increments either the major (a), minor (b), or patch (c)
// numbers based on the supplied argument.
// It rewrites the file with the new version, commits the change
// with "git commit" (with -commit) and tags it with "git tag" (with -tag).
//
// Usage:
//   go run cmd/bump-version/main.go [-n] patch src/main.js
//
// The `-n` flag parses the file but makes no changes or git actions.
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strconv"
)

var (
	commit = flag.Bool("commit", false, "Commit the change.")
	dryRun = flag.Bool("n", false, "Dry run - report version but do not make changes")
	tag    = flag.Bool("tag", false, "Tag the change.")

	versionRE = regexp.MustCompile(`\nconst VERSION = (['"])(\d+)\.(\d+)\.(\d+)['"]\n`)
)

func main() {
	flag.Parse()

	if flag.NArg() != 2 {
		log.Fatal("Usage: bump-version patch filename")
	}

	filename := flag.Arg(1)
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("ioutil.ReadFile: %v", err)
	}

	m := versionRE.FindStringSubmatch(string(buf))
	if len(m) != 5 {
		log.Fatalf("Unable to find 'const VERSION' string in file %q", filename)
	}

	cvt := func(s string) int {
		v, err := strconv.Atoi(s)
		if err != nil {
			log.Fatalf("strconv.Atoi: %v", err)
		}
		return v
	}

	quote := m[1]
	major := cvt(m[2])
	minor := cvt(m[3])
	patch := cvt(m[4])

	switch flag.Arg(0) {
	case "major":
		major++
		minor = 0
		patch = 0
	case "minor":
		minor++
		patch = 0
	case "patch":
		patch++
	default:
		log.Fatalf("Second argument must be one of: 'major', 'minor', or 'patch'.")
	}

	newVersion := fmt.Sprintf("%v.%v.%v", major, minor, patch)
	buf = versionRE.ReplaceAll(buf, []byte(fmt.Sprintf("\nconst VERSION = %v%v%v\n", quote, newVersion, quote)))
	if *dryRun {
		log.Printf("New version: %v\nDry-run mode; no changes made. Done.", newVersion)
		os.Exit(0)
	}

	if err := ioutil.WriteFile(filename, buf, 0644); err != nil {
		log.Fatalf("ioutil.WriteFile: %v", err)
	}

	log.Printf("Bumped version to %v.", newVersion)

	log.Printf("Running: git add %v", filename)
	buf, err = exec.Command("git", "add", filename).CombinedOutput()
	if err != nil {
		log.Fatalf("git add %v: %v\n%s", filename, err, buf)
	}

	if *commit {
		log.Printf("Running: git commit -m %v", newVersion)
		buf, err = exec.Command("git", "commit", "-m", newVersion).CombinedOutput()
		if err != nil {
			log.Fatalf("git commit -m %v: %v\n%s", newVersion, err, buf)
		}
	}

	if *tag {
		log.Printf("Running: git tag %v", newVersion)
		buf, err = exec.Command("git", "tag", newVersion).CombinedOutput()
		if err != nil {
			log.Fatalf("git tag %v: %v\n%s", newVersion, err, buf)
		}
	}

	log.Printf("Done.")
}
