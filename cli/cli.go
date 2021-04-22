/*
Copyright © 2020 Glenn M. Lewis

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

// Package cli implements the methods for the flutter-stylizer
// command-line tool.
package cli

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gmlewis/go-flutter-stylizer/dart"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// This version string should stay in sync with the VSCode plugin
// to reduce confusion.
const VERSION = "0.0.20"

var cfgFile string

var rootHelp = `This is flutter-stylizer version %v.

To stylize all Dart files in directory tree, run:
  $ flutter-stylizer -w ./...
To show stylizer differences for all Dart files in directory tree, run:
  $ flutter-stylizer -d ./...
To list all Dart files with stylizer differences in directory tree, run:
  $ flutter-stylizer -l ./...
To show stylized output for a source file, run:
  $ flutter-stylizer lib/path/to/file.dart

For more information, see the repo: https://github.com/gmlewis/go-flutter-stylizer
`

var rootCmd = &cobra.Command{
	Use:   "flutter-stylizer [flags] [path ... | ./...]",
	Short: "flutter-stylizer is a stand-alone version of the VSCode plugin of the same name",
	Long:  fmt.Sprintf(rootHelp, VERSION),
	RunE:  rootRunE,
}

func rootRunE(cmd *cobra.Command, args []string) error {
	diff, err := cmd.Flags().GetBool("diff")
	if err != nil {
		log.Fatalf("diff: %v", err)
	}
	list, err := cmd.Flags().GetBool("list")
	if err != nil {
		log.Fatalf("list: %v", err)
	}
	write, err := cmd.Flags().GetBool("write")
	if err != nil {
		log.Fatalf("write: %v", err)
	}
	verbose, err := cmd.Flags().GetBool("verbose")
	if err != nil {
		log.Fatalf("verbose: %v", err)
	}
	quiet, err := cmd.Flags().GetBool("quiet")
	if err != nil {
		log.Fatalf("quiet: %v", err)
	}

	var flagCount int
	f := func(b bool) {
		if b {
			flagCount++
		}
	}

	f(diff)
	f(list)
	f(write)

	if flagCount > 1 {
		log.Fatalf("Must supply only one of --diff (-d), --list (-l), and --write (-w).")
	}

	opts := dart.Options{
		Diff:    diff,
		List:    list,
		Quiet:   quiet,
		Verbose: verbose,
		Write:   write,
	}
	c := dart.New(opts)

	if len(args) == 0 {
		log.Fatalf("Must supply at least one filename or './...' to process all files in directory tree.")
	}

	var newArgs []string
	for i := 0; i < len(args); i++ {
		if strings.HasSuffix(args[i], "/...") || strings.HasSuffix(args[i], `\...`) {
			dartFiles := findDartFiles(args[i][0 : len(args[i])-4])
			newArgs = append(newArgs, dartFiles...)
			continue
		}
		newArgs = append(newArgs, args[i])
	}
	if len(newArgs) == 0 {
		log.Fatalf("No *.dart files found.")
	}

	if !quiet {
		log.Printf("Stylizing %v files...", len(newArgs))
	}

	var anyDiffs bool
	for i := 0; i < len(newArgs); i++ {
		arg := newArgs[i]
		if verbose && !quiet {
			log.Printf("Stylizing file: %v ...", arg)
		}
		diffs, err := c.StylizeFile(arg)
		if err != nil {
			log.Fatalf("StylizeFile(%q): %v", arg, err)
		}
		anyDiffs = anyDiffs || diffs
	}

	if anyDiffs && (opts.Diff || opts.List) {
		log.Fatalf("Differences were found. Exit code 1.")
	}

	if !quiet {
		log.Printf("Done.")
	}

	return nil
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.flutter-stylizer.yaml)")
	rootCmd.Flags().BoolP("diff", "d", false, "display diffs (cannot be used with -l or -w); exit code 1 on diffs")
	rootCmd.Flags().BoolP("list", "l", false, "list files whose formatting differs from flutter-stylizer's (cannot be used with -d or -w); exit code 1 on diffs")
	rootCmd.Flags().BoolP("write", "w", false, "write result to (source) file instead of stdout (cannot be used with -d or -l); exit code 0 on diffs")
	rootCmd.Flags().BoolP("verbose", "v", false, "write progress details to stderr")
	rootCmd.Flags().BoolP("quiet", "q", false, "don't print unless there are errors")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".flutter-stylizer" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".flutter-stylizer")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
