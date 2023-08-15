package main

import (
	"fmt"
	"os"

	"github.com/11wizards/go-to-dart/generator"
	"github.com/11wizards/go-to-dart/generator/options"
	"github.com/spf13/cobra"
)

var input, output, mode string
var imports []string

var rootCmd = &cobra.Command{
	Use:   "go-to-dart",
	Short: "Go-to-Dart is a tool to generate Dart classes from Go structs",
	Run: func(cmd *cobra.Command, args []string) {
		o := options.Options{
			Input:   input,
			Output:  output,
			Imports: imports,
			Mode:    options.Mode(mode),
		}

		if o.Mode != options.JSON && o.Mode != options.Firestore {
			fmt.Println("Mode must be either json or firestore")
			os.Exit(1)
		}
		generator.Run(o)
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&input, "input", "i", "", "Input directory to read from")
	rootCmd.PersistentFlags().StringVarP(&output, "output", "o", "", "Output directory to write to")
	rootCmd.PersistentFlags().StringVarP(&mode, "mode", "m", "json", "Mode to run in: json or firestore")
	rootCmd.PersistentFlags().StringSliceVarP(&imports, "imports", "p", []string{}, "Additional imports to add to the generated file")

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
