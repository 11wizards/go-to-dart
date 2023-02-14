package generator

import (
	"fmt"
	"github.com/11wizards/go-to-dart/generator/format"
	"go/ast"
	"io"
)

func generateFieldDeclaration(writer io.Writer, f *ast.Field) {
	formatter := format.GetTypeFormatter(f.Type)
	fieldName := format.GetFieldName(f)
	jsonFieldName := format.GetJSONFieldName(f)

	if jsonFieldName != "" && jsonFieldName != fieldName {
		fmt.Fprintf(writer, "@JsonKey(name: '%s') ", jsonFieldName)
	} else if jsonFieldName == "" {
		fmt.Fprintf(writer, "@JsonKey(name: '%s') ", f.Names[0].Name)
	}

	fmt.Fprintf(writer, "final %s", formatter.Declaration(format.GetFieldName(f), f.Type))
}

func generateFieldConstrutor(writer io.Writer, f *ast.Field) {
	formatter := format.GetTypeFormatter(f.Type)
	fmt.Fprint(writer, formatter.Constructor(format.GetFieldName(f), f.Type))
}
