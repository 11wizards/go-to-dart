package generator

import (
	"fmt"
	"github.com/11wizards/go-to-dart/generator/format"
	"github.com/11wizards/go-to-dart/generator/options"
	"go/ast"
	"io"
)

func generateFieldDeclaration(writer io.Writer, f *ast.Field, registry *format.TypeFormatterRegistry, mode options.Mode) {
	formatter := registry.GetTypeFormatter(f.Type)
	fieldName := format.GetFieldName(f)
	jsonFieldName := format.GetJSONFieldName(f, mode)

	if jsonFieldName != "" && jsonFieldName != fieldName {
		fmt.Fprintf(writer, "@JsonKey(name: '%s')\n", jsonFieldName)
	} else if jsonFieldName == "" {
		fmt.Fprintf(writer, "@JsonKey(name: '%s')\n", f.Names[0].Name)
	}

	fmt.Fprintf(writer, "final %s", formatter.Declaration(format.GetFieldName(f), f.Type))
}

func generateFieldConstrutor(writer io.Writer, f *ast.Field, registry *format.TypeFormatterRegistry) {
	formatter := registry.GetTypeFormatter(f.Type)
	fmt.Fprint(writer, formatter.Constructor(format.GetFieldName(f), f.Type))
}
