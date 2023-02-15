package format

import (
	"fmt"
	"go/ast"
)

type TimeFormatter struct {
	TypeFormatterBase
}

func (f *TimeFormatter) CanFormat(expr ast.Expr) bool {
	if v, ok := expr.(*ast.SelectorExpr); ok {
		return v.X.(*ast.Ident).Name == "time" && v.Sel.Name == "Time"
	}
	return false
}

func (f *TimeFormatter) Signature(_ ast.Expr) string {
	return "DateTime"
}

func (f *TimeFormatter) Declaration(fieldName string, expr ast.Expr) string {
	return fmt.Sprintf("%s %s", f.Signature(expr), fieldName)
}

func (f *TimeFormatter) Constructor(fieldName string, _ ast.Expr) string {
	return fmt.Sprintf("required this.%s", fieldName)
}
