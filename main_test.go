package main

import (
	"os"
	"path"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/11wizards/go-to-dart/generator/options"
)

func runAndCompare(t *testing.T, input string, opts options.Options) {
	output := path.Join(os.TempDir(), "go-to-dart-test-output", t.Name())

	if err := os.RemoveAll(output); err != nil {
		t.Fatal(err)
		return
	}

	args := []string{"-i", input, "-o", output, "-m", string(opts.Mode), "-x", opts.Prefix}

	rootCmd.SetArgs(args)

	err := rootCmd.Execute()
	require.NoError(t, err, "command failed")

	files, err := filepath.Glob(output + "**/*.dart")
	require.NoError(t, err, "failed to list files")

	for _, file := range files {
		relativeFile, err := filepath.Rel(output, file)
		require.NoError(t, err, "failed to get relative path of %s to %s", file, output)

		expectedFile := filepath.Join(input, strings.Replace(relativeFile, ".dart.txt", ".dart", 1))

		_, err = os.Stat(expectedFile)
		require.NoError(t, err, "expected file %s to exist", expectedFile)

		actual, err := os.ReadFile(file)
		require.NoError(t, err, "failed to read file %s", file)

		expected, err := os.ReadFile(expectedFile)
		require.NoError(t, err, "failed to read file %s", expectedFile)

		require.Equal(t, string(expected), string(actual), "file %s does not match expected file %s", file, expectedFile)
	}
}

func TestEverything(t *testing.T) {
	runAndCompare(t, "./examples/everything", options.Options{
		Mode: options.JSON,
	})
}

func TestUser(t *testing.T) {
	runAndCompare(t, "./examples/user", options.Options{
		Mode: options.JSON,
	})
}

func TestFirestore(t *testing.T) {
	runAndCompare(t, "./examples/firestore", options.Options{
		Mode: options.Firestore,
	})
}

func TestMultipackage(t *testing.T) {
	runAndCompare(t, "./examples/multipackage", options.Options{
		Mode: options.Firestore,
	})
}

func TestGenerics(t *testing.T) {
	runAndCompare(t, "./examples/generics", options.Options{
		Mode: options.JSON,
	})
}

func TestEmbedded(t *testing.T) {
	runAndCompare(t, "./examples/embedded", options.Options{
		Mode: options.JSON,
	})
}

func TestPrefix(t *testing.T) {
	runAndCompare(t, "./examples/prefix", options.Options{
		Mode:   options.JSON,
		Prefix: "My",
	})
}
