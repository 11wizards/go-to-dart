package format

import (
	"fmt"
	"go/ast"
)

type FallbackFormatter struct {
	TypeFormatterBase
}

func (f *FallbackFormatter) under(expr ast.Expr) *ast.TypeSpec {
	if e, ok := expr.(*ast.Ident); ok && e.Obj == nil {
		if val, ok := f.Registry.KnownTypes[e.Name]; ok {
			return val
		}
	}
	return nil
}

func (f *FallbackFormatter) CanFormat(expr ast.Expr) bool {
	return f.under(expr) != nil
}

func (f *FallbackFormatter) Signature(expr ast.Expr) string {
	t := f.under(expr)
	switch t.Type.(type) {
	case *ast.Ident:
		return f.Registry.GetTypeFormatter(t.Type).Signature(t.Type)
	default:
		return t.Name.Name
	}
}

func (f *FallbackFormatter) Attribute(_ ast.Expr) string {
	return ""
}

func (f *FallbackFormatter) Declaration(fieldName string, expr ast.Expr) string {
	return fmt.Sprintf("%s %s", f.Signature(expr), fieldName)
}

func (f *FallbackFormatter) Constructor(fieldName string, _ ast.Expr) string {
	return "required this." + fieldName
}
