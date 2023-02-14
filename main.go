package main

import (
	"fmt"
	"github.com/11wizards/go-to-dart/generator"
	"github.com/spf13/cobra"
	"os"
)

var Input string
var Output string

var rootCmd = &cobra.Command{
	Use:   "go-to-dart",
	Short: "Go-to-Dart is a tool to generate Dart classes from Go structs",
	Run: func(cmd *cobra.Command, args []string) {
		generator.Run(Input, Output)
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&Input, "input", "i", "", "Input directory to read from")
	rootCmd.PersistentFlags().StringVarP(&Output, "output", "o", "", "Output directory to write to")

	if err := rootCmd.MarkPersistentFlagRequired("input"); err != nil {
		panic(err)
	}

	if err := rootCmd.MarkPersistentFlagRequired("output"); err != nil {
		panic(err)
	}
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
