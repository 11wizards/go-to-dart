package format

import (
	"fmt"
	"go/ast"
)

type PrimitiveFormatter struct {
}

func (f *PrimitiveFormatter) toDartPrimitive(expr ast.Expr) string {
	if e, ok := expr.(*ast.Ident); ok && e.Obj == nil {
		switch e.Name {
		case "bool":
			return "bool"
		case "byte":
			return "int"
		case "float32":
			return "double"
		case "float64":
			return "double"
		case "int":
			return "int"
		case "int16":
			return "int"
		case "int32":
			return "int"
		case "int64":
			return "int"
		case "int8":
			return "int"
		case "rune":
			return "int"
		case "string":
			return "String"
		case "uint":
			return "int"
		case "uint16":
			return "int"
		case "uint32":
			return "int"
		case "uint64":
			return "int"
		case "uint8":
			return "int"
		case "uintptr":
			return "int"
		}
	}

	return ""
}

func (f *PrimitiveFormatter) CanFormat(expr ast.Expr) bool {
	return f.toDartPrimitive(expr) != ""
}

func (f *PrimitiveFormatter) Signature(expr ast.Expr) string {
	return f.toDartPrimitive(expr)
}

func (f *PrimitiveFormatter) Declaration(fieldName string, expr ast.Expr) string {
	return fmt.Sprintf("%s %s", f.Signature(expr), fieldName)
}

func (f *PrimitiveFormatter) Constructor(fieldName string, _ ast.Expr) string {
	return "required this." + fieldName
}
