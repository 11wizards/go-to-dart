package generator

import (
	"fmt"
	"go/types"
	"io"

	"github.com/11wizards/go-to-dart/generator/format"
	"github.com/11wizards/go-to-dart/generator/options"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

func generateFieldJSONKey(writer io.Writer, f *types.Var, tag string, registry *format.TypeFormatterRegistry, opts options.Options) format.TypeFormatter {
	formatter := registry.GetTypeFormatter(f.Type())
	fieldName := format.GetFieldName(f)
	jsonFieldName := format.GetJSONFieldName(tag, opts.Mode)

	keyProperties := map[string]string{}

	if jsonFieldName != "" && jsonFieldName != fieldName {
		keyProperties["name"] = fmt.Sprintf("\"%s\"", jsonFieldName)
	} else if jsonFieldName == "" {
		keyProperties["name"] = fmt.Sprintf("\"%s\"", f.Name())
	}

	if formatter.DefaultValue(f.Type()) != "" {
		keyProperties["defaultValue"] = formatter.DefaultValue(f.Type())
	}

	if len(keyProperties) > 0 {
		fmt.Fprint(writer, "@JsonKey(")
		first := true

		keys := maps.Keys(keyProperties)
		slices.Sort(keys)

		for _, key := range keys {
			value := keyProperties[key]

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

func generateFieldDeclaration(writer io.Writer, f *types.Var, tag string, registry *format.TypeFormatterRegistry, opts options.Options) {
	formatter := generateFieldJSONKey(writer, f, tag, registry, opts)
	fmt.Fprintf(writer, "final %s", formatter.Declaration(format.GetFieldName(f), f.Type()))
}

func generateFieldConstrutor(writer io.Writer, f *types.Var, registry *format.TypeFormatterRegistry) {
	formatter := registry.GetTypeFormatter(f.Type())
	fmt.Fprint(writer, formatter.Constructor(format.GetFieldName(f), f.Type()))
}
