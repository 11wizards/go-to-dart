package generator

import (
	"fmt"
	"github.com/11wizards/go-to-dart/generator/format"
	"github.com/11wizards/go-to-dart/generator/options"
	"go/ast"
	"io"
)

func generateFieldJSONKey(writer io.Writer, f *ast.Field, registry *format.TypeFormatterRegistry, mode options.Mode) format.TypeFormatter {
	formatter := registry.GetTypeFormatter(f.Type)
	fieldName := format.GetFieldName(f)
	jsonFieldName := format.GetJSONFieldName(f, mode)

	keyProperties := map[string]string{}

	if jsonFieldName != "" && jsonFieldName != fieldName {
		keyProperties["name"] = fmt.Sprintf("\"%s\"", jsonFieldName)
	} else if jsonFieldName == "" {
		keyProperties["name"] = fmt.Sprintf("\"%s\"", f.Names[0].Name)
	}

	if formatter.DefaultValue(f.Type) != "" {
		keyProperties["defaultValue"] = formatter.DefaultValue(f.Type)
	}

	if len(keyProperties) > 0 {
		fmt.Fprint(writer, "@JsonKey(")
		first := true
		for key, value := range keyProperties {
			if !first {
				fmt.Fprint(writer, ", ")
			} else {
				first = false
			}

			fmt.Fprintf(writer, "%s: %s", key, value)
		}
		fmt.Fprintf(writer, ")")

	}
	return formatter
}

func generateFieldDeclaration(writer io.Writer, f *ast.Field, registry *format.TypeFormatterRegistry, mode options.Mode) {
	formatter := generateFieldJSONKey(writer, f, registry, mode)

	fmt.Fprintf(writer, "final %s", formatter.Declaration(format.GetFieldName(f), f.Type))
}

func generateFieldConstrutor(writer io.Writer, f *ast.Field, registry *format.TypeFormatterRegistry) {
	formatter := registry.GetTypeFormatter(f.Type)
	fmt.Fprint(writer, formatter.Constructor(format.GetFieldName(f), f.Type))
}
