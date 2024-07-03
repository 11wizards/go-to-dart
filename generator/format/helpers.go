package format

import (
	"bytes"
	"fmt"
	"go/types"
	"reflect"
	"strings"

	"github.com/11wizards/go-to-dart/generator/options"
	"github.com/iancoleman/strcase"
)

func GetFieldName(f *types.Var) string {
	if f.Anonymous() {
		panic(fmt.Sprintf("no name for field: %#v", f))
	}

	return strcase.ToLowerCamel(f.Name())
}

func GetJSONFieldName(tag string, mode options.Mode) string {
	var tagName string
	if mode == options.Firestore {
		tagName = "firestore"
	} else {
		tagName = "json"
	}

	if tag != "" {
		val := reflect.StructTag(strings.Trim(tag, "`"))
		value, ok := val.Lookup(tagName)
		if ok {
			return strings.Split(value, ",")[0]
		}
	}

	return ""
}

func GenerateTypeParams(registry *TypeFormatterRegistry, named *types.Named) string {
	buf := new(bytes.Buffer)
	typeArgs := named.TypeArgs()
	typeParams := named.TypeParams()

	if typeArgs.Len() > 0 {
		fmt.Fprint(buf, "<")
		for i := 0; i < typeArgs.Len(); i++ {
			arg := typeArgs.At(i)
			name := registry.GetTypeFormatter(arg).Signature(arg)
			if i > 0 {
				fmt.Fprint(buf, ", ")
			}
			fmt.Fprint(buf, name)
		}
		fmt.Fprint(buf, ">")
	} else if typeParams.Len() > 0 {
		fmt.Fprint(buf, "<")
		for i := 0; i < typeParams.Len(); i++ {
			param := typeParams.At(i).String()
			if i > 0 {
				fmt.Fprint(buf, ", ")
			}
			fmt.Fprint(buf, param)
		}
		fmt.Fprint(buf, ">")
	}

	return buf.String()
}
