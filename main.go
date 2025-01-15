package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const (
	outDir = "./out"
	srcDir = "./src"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "gostatic",
		Short: "a simple static site generator.",
	}

	buildCmd := &cobra.Command{
		Use:   "build",
		Short: "Run the build function",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Running build function...")
			Build()
		},
	}

	initCmd := &cobra.Command{
		Use:   "init",
		Short: "Initialize the file structure",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Initializing file structure...")
			Init()
		},
	}

	rootCmd.AddCommand(buildCmd)
	rootCmd.AddCommand(initCmd)
    rootCmd.CompletionOptions.DisableDefaultCmd = true

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

