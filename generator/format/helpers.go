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

// isDartKeyword checks if a given identifier is a Dart keyword
func isDartKeyword(name string) bool {
	switch name {
	// Keywords that can't be used as identifiers at all
	case "abstract", "as", "assert", "async", "await", "base", "break", "case",
		"catch", "class", "const", "continue", "covariant", "default", "deferred",
		"do", "dynamic", "else", "enum", "export", "extends", "extension", "external",
		"factory", "false", "final", "finally", "for", "Function", "get", "hide",
		"if", "implements", "import", "in", "interface", "is", "late", "library",
		"mixin", "new", "null", "of", "on", "operator", "part", "required",
		"rethrow", "return", "sealed", "set", "show", "static", "super", "switch",
		"sync", "this", "throw", "true", "try", "type", "typedef", "var", "void",
		"when", "with", "while", "yield":
		return true
	default:
		return false
	}
}

// escapeDartKeyword escapes a Dart keyword by appending a dollar sign
func escapeDartKeyword(name string) string {
	// Check both the original name and its lowercase version
	if isDartKeyword(name) || isDartKeyword(strings.ToLower(name)) {
		return name + "$"
	}
	return name
}

func GetFieldName(f *types.Var) string {
	if f.Anonymous() {
		panic(fmt.Sprintf("no name for field: %#v", f))
	}

	fieldName := strcase.ToLowerCamel(f.Name())
	return escapeDartKeyword(fieldName)
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
