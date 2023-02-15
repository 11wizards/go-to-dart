package format

import (
	"fmt"
	"go/ast"
)

type PointerFormatter struct {
	TypeFormatterBase
}

func (f *PointerFormatter) under(expr ast.Expr) (TypeFormatter, ast.Expr) {
	starExpr := expr.(*ast.StarExpr)
	formatter := f.Registry.GetTypeFormatter(starExpr.X)
	return formatter, starExpr.X
}

func (f *PointerFormatter) CanFormat(expr ast.Expr) bool {
	_, ok := expr.(*ast.StarExpr)
	return ok
}

func (f *PointerFormatter) Signature(expr ast.Expr) string {
	formatter, expr := f.under(expr)
	return fmt.Sprintf("%s?", formatter.Signature(expr))
}

func (f *PointerFormatter) Attribute(_ ast.Expr) string {
	return ""
}

func (f *PointerFormatter) Declaration(fieldName string, expr ast.Expr) string {
	return fmt.Sprintf("%s %s", f.Signature(expr), fieldName)
}

func (f *PointerFormatter) Constructor(fieldName string, _ ast.Expr) string {
	return "this." + fieldName
}
