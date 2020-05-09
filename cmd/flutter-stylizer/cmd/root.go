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

// Package cmd implements the methods for the flutter-stylizer
// command-line tool.
package cmd

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootHelp = `
`

var rootCmd = &cobra.Command{
	Use:   "flutter-stylizer [flags] [path ...]",
	Short: "flutter-stylizer is a stand-alone version of the VSCode plugin of the same name",
	Long:  rootHelp,
	RunE:  rootRunE,
}

func rootRunE(cmd *cobra.Command, args []string) error {
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
	rootCmd.Flags().BoolP("diff", "d", false, "display diffs instead of rewriting files")
	rootCmd.Flags().BoolP("list", "l", false, "list files whose formatting differs from flutter-stylizer's")
	rootCmd.Flags().BoolP("write", "w", false, "write result to (source) file instead of stdout")
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
