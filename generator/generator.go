package generator

import (
	"bytes"
	_ "embed"
	"fmt"
	"github.com/11wizards/go-to-dart/generator/format"
	"github.com/11wizards/go-to-dart/generator/format/mo"
	"github.com/11wizards/go-to-dart/generator/options"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"os"
	"path/filepath"
	"sort"
)

//go:embed dart/timestamp_converter.dart
var timestampConverterSrc string

func generateHeader(pkg *ast.Package, wr io.Writer, mode options.Mode) {
	if mode == options.Firestore {
		fmt.Fprint(wr, "import 'package:cloud_firestore/cloud_firestore.dart';\n")
	}

	fmt.Fprint(wr, "import 'package:json_annotation/json_annotation.dart';\n\n")
	fmt.Fprintf(wr, "part '%s.go.g.dart';\n\n", pkg.Name)

	if mode == options.Firestore {
		fmt.Fprint(wr, timestampConverterSrc)
		fmt.Fprint(wr, "\n\n")
	}
}

func createRegistry(mode options.Mode) *format.TypeFormatterRegistry {
	registry := format.NewTypeFormatterRegistry()

	base := format.TypeFormatterBase{Mode: mode}

	registry.RegisterTypeFormatter(&format.FallbackFormatter{TypeFormatterBase: base})
	registry.RegisterTypeFormatter(&format.StructFormatter{TypeFormatterBase: base})
	registry.RegisterTypeFormatter(&format.AliasFormatter{TypeFormatterBase: base})
	registry.RegisterTypeFormatter(&format.PrimitiveFormatter{TypeFormatterBase: base})
	registry.RegisterTypeFormatter(&format.TimeFormatter{TypeFormatterBase: base})
	registry.RegisterTypeFormatter(&format.PointerFormatter{TypeFormatterBase: base})
	registry.RegisterTypeFormatter(&format.ArrayFormatter{TypeFormatterBase: base})
	registry.RegisterTypeFormatter(&format.MapFormatter{TypeFormatterBase: base})
	registry.RegisterTypeFormatter(&mo.OptionFormatter{TypeFormatterBase: base})

	return registry
}

func generateClasses(pkg *ast.Package, wr io.Writer, mode options.Mode) {
	registry := createRegistry(mode)

	ast.Inspect(pkg, func(node ast.Node) bool {
		ts, ok := node.(*ast.TypeSpec)
		if ok {
			registry.KnownTypes[ts.Name.Name] = ts
			return false
		}

		return true
	})

	var list []struct {
		TypeSpec   *ast.TypeSpec
		StructType *ast.StructType
	}

	ast.Inspect(pkg, func(node ast.Node) bool {
		ts, ok := node.(*ast.TypeSpec)
		if !ok {
			return true
		}

		st, ok := ts.Type.(*ast.StructType)
		if !ok {
			return true
		}

		list = append(list, struct {
			TypeSpec   *ast.TypeSpec
			StructType *ast.StructType
		}{
			TypeSpec:   ts,
			StructType: st,
		})

		return false
	})

	sort.Slice(list, func(i, j int) bool {
		return list[i].TypeSpec.Name.Name < list[j].TypeSpec.Name.Name
	})

	for _, item := range list {
		generateDartClass(wr, item.TypeSpec, item.StructType, registry, mode)
	}
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

func Run(options options.Options) {
	fileSet := token.NewFileSet()
	f, err := parser.ParseDir(fileSet, options.Input, nil, 0)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, pkg := range f {
		var buf []byte
		wr := bytes.NewBuffer(buf)
		generateHeader(pkg, wr, options.Mode)
		generateClasses(pkg, wr, options.Mode)
		writeOut(options.Output, fmt.Sprintf("%s.go.dart", pkg.Name), wr)
	}
}
