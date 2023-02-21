package format

import (
	"fmt"
	"go/ast"
)

type ArrayFormatter struct {
	TypeFormatterBase
}

func (f *ArrayFormatter) under(expr ast.Expr) (TypeFormatter, ast.Expr) {
	arrayExpr := expr.(*ast.ArrayType)
	formatter := f.Registry.GetTypeFormatter(arrayExpr.Elt)
	return formatter, arrayExpr.Elt
}

func (f *ArrayFormatter) CanFormat(expr ast.Expr) bool {
	_, ok := expr.(*ast.ArrayType)
	return ok
}

func (f *ArrayFormatter) Signature(expr ast.Expr) string {
	formatter, expr := f.under(expr)
	return fmt.Sprintf("List<%s>", formatter.Signature(expr))
}

func (f *ArrayFormatter) DefaultValue(expr ast.Expr) string {
	return fmt.Sprintf("<%s>[]", f.Signature(expr))
}

func (f *ArrayFormatter) Declaration(fieldName string, expr ast.Expr) string {
	return fmt.Sprintf("%s %s", f.Signature(expr), fieldName)
}

func (f *ArrayFormatter) Constructor(fieldName string, _ ast.Expr) string {
	return "required this." + fieldName
}
