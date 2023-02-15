package generator

import (
	"bytes"
	"fmt"
	"github.com/11wizards/go-to-dart/generator/format"
	"github.com/11wizards/go-to-dart/generator/format/mo"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"os"
	"path/filepath"
)

func createRegistry() *format.TypeFormatterRegistry {
	registry := format.NewTypeFormatterRegistry()

	registry.RegisterTypeFormatter(&format.FallbackFormatter{})
	registry.RegisterTypeFormatter(&format.StructFormatter{})
	registry.RegisterTypeFormatter(&format.AliasFormatter{})
	registry.RegisterTypeFormatter(&format.PrimitiveFormatter{})
	registry.RegisterTypeFormatter(&format.TimeFormatter{})
	registry.RegisterTypeFormatter(&format.PointerFormatter{})
	registry.RegisterTypeFormatter(&format.ArrayFormatter{})
	registry.RegisterTypeFormatter(&format.MapFormatter{})
	registry.RegisterTypeFormatter(&mo.OptionFormatter{})

	return registry

}

func traversePackage(f *ast.Package, outputFile io.Writer) {
	fmt.Fprint(outputFile, "import 'package:json_annotation/json_annotation.dart';\n\n")
	fmt.Fprintf(outputFile, "part '%s.g.dart';\n\n", f.Name)

	registry := createRegistry()

	ast.Inspect(f, func(node ast.Node) bool {
		ts, ok := node.(*ast.TypeSpec)
		if ok {
			registry.KnownTypes[ts.Name.Name] = ts
			return false
		}

		return true
	})

	ast.Inspect(f, func(node ast.Node) bool {
		ts, ok := node.(*ast.TypeSpec)
		if !ok {
			return true
		}

		st, ok := ts.Type.(*ast.StructType)
		if !ok {
			return true
		}

		return generateDartClass(outputFile, ts, st, registry)
	})
}

func writeOut(output, outputDartFile string, wr *bytes.Buffer) {
	if _, err := os.Stat(output); os.IsNotExist(err) {
		err = os.MkdirAll(output, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}

	outputFilePath := filepath.Join(output, outputDartFile)
	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		panic(err)
	}

	defer func() { _ = outputFile.Close() }()

	_, err = outputFile.Write(wr.Bytes())

	if err != nil {
		panic(err)
	}

	fmt.Printf("Processed: %s -> %s\n", outputDartFile, outputFilePath)
}

func Run(input string, output string) {
	fileSet := token.NewFileSet()
	f, err := parser.ParseDir(fileSet, input, nil, 0)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, pkg := range f {
		var buf []byte
		wr := bytes.NewBuffer(buf)
		traversePackage(pkg, wr)
		writeOut(output, fmt.Sprintf("%s.dart", pkg.Name), wr)
	}
}
