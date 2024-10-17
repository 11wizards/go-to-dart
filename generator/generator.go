package generator

import (
	"bytes"
	_ "embed"
	"fmt"
	"go/types"
	"io"
	"os"
	"path/filepath"
	"sort"

	"github.com/11wizards/go-to-dart/generator/format"
	"github.com/11wizards/go-to-dart/generator/format/mo"
	"github.com/11wizards/go-to-dart/generator/options"

	"golang.org/x/tools/go/packages"
)

//go:embed dart/timestamp_converter.dart
var timestampConverterSrc string

func generateHeader(pkg *packages.Package, wr io.Writer, opts options.Options) {
	fmt.Fprintln(wr, "// ignore_for_file: always_use_package_imports")

	if opts.Mode == options.Firestore {
		fmt.Fprint(wr, "import 'package:cloud_firestore/cloud_firestore.dart';\n")
	}

	fmt.Fprint(wr, "import 'package:copy_with_extension/copy_with_extension.dart';\n")
	fmt.Fprint(wr, "import 'package:equatable/equatable.dart';\n")
	fmt.Fprint(wr, "import 'package:json_annotation/json_annotation.dart';\n")
	for _, imp := range opts.Imports {
		fmt.Fprintf(wr, "import '%s';\n", imp)
	}

	fmt.Fprintf(wr, "\npart '%s.go.g.dart';\n\n", pkg.Name)

	if opts.Mode == options.Firestore {
		fmt.Fprint(wr, timestampConverterSrc)
		fmt.Fprint(wr, "\n\n")
	}
}

func createRegistry(options options.Options) *format.TypeFormatterRegistry {
	registry := format.NewTypeFormatterRegistry()

	typeFormatterBase := format.TypeFormatterBase{Options: options}

	registry.RegisterTypeFormatter(&format.AliasFormatter{TypeFormatterBase: typeFormatterBase})
	registry.RegisterTypeFormatter(&format.ConcreteStructFormatter{TypeFormatterBase: typeFormatterBase})
	registry.RegisterTypeFormatter(&format.GenericStructFormatter{TypeFormatterBase: typeFormatterBase})
	registry.RegisterTypeFormatter(&format.PrimitiveFormatter{TypeFormatterBase: typeFormatterBase})
	registry.RegisterTypeFormatter(&format.TimeFormatter{TypeFormatterBase: typeFormatterBase})
	registry.RegisterTypeFormatter(&format.PointerFormatter{TypeFormatterBase: typeFormatterBase})
	registry.RegisterTypeFormatter(&format.ArrayFormatter{TypeFormatterBase: typeFormatterBase})
	registry.RegisterTypeFormatter(&format.MapFormatter{TypeFormatterBase: typeFormatterBase})
	registry.RegisterTypeFormatter(&format.TypeParamsFormatter{TypeFormatterBase: typeFormatterBase})
	registry.RegisterTypeFormatter(&mo.OptionFormatter{TypeFormatterBase: typeFormatterBase})

	return registry
}

func generateClasses(pkg *packages.Package, wr io.Writer, options options.Options) {
	registry := createRegistry(options)

	for _, value := range pkg.TypesInfo.Defs {
		if value == nil {
			continue
		}

		if typeName, ok := value.(*types.TypeName); ok && typeName.Exported() {
			registry.AddKnownType(typeName.Type())
		}
	}

	var list []struct {
		TypeName   *types.TypeName
		StructType *types.Struct
	}

	for _, value := range pkg.TypesInfo.Defs {
		if value == nil {
			continue
		}

		if typeName, ok := value.(*types.TypeName); ok && typeName.Exported() {
			if structType, ok := typeName.Type().Underlying().(*types.Struct); ok {
				list = append(list, struct {
					TypeName   *types.TypeName
					StructType *types.Struct
				}{
					TypeName:   typeName,
					StructType: structType,
				})
			}
		}
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i].TypeName.Name() < list[j].TypeName.Name()
	})

	for _, item := range list {
		generateDartClass(wr, item.TypeName, item.StructType, registry, options)
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

func Run(opts options.Options) {
	if abs, err := filepath.Abs(opts.Input); err == nil {
		opts.Input = abs
	} else {
		panic(err)
	}

	pkgs, err := packages.Load(&packages.Config{
		Dir:  opts.Input,
		Mode: packages.NeedName | packages.NeedFiles | packages.NeedImports | packages.NeedDeps | packages.NeedTypes | packages.NeedSyntax | packages.NeedTypesInfo,
	}, opts.Input)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, pkg := range pkgs {
		if len(pkg.Errors) > 0 {
			for _, err := range pkg.Errors {
				fmt.Println(err)
			}

			os.Exit(1)
		}
	}

	for _, pkg := range pkgs {
		var buf []byte
		wr := bytes.NewBuffer(buf)
		generateHeader(pkg, wr, opts)
		generateClasses(pkg, wr, opts)
		writeOut(opts.Output, fmt.Sprintf("%s.go.dart", pkg.Name), wr)
	}
}
