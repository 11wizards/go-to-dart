package format

import (
	"fmt"
	"github.com/iancoleman/strcase"
	"go/ast"
	"reflect"
	"strings"
)

func GetFieldName(f *ast.Field) string {
	if f.Names == nil {
		panic(fmt.Sprintf("no name for field: %#v", f))
	}

	return strcase.ToLowerCamel(f.Names[0].Name)
}

func GetJSONFieldName(f *ast.Field) string {
	// Check for json struct field tag
	if f.Tag != nil {
		val := reflect.StructTag(strings.Trim(f.Tag.Value, "`"))
		tag, ok := val.Lookup("json")
		if ok {
			return tag
		}
	}

	return ""
}
