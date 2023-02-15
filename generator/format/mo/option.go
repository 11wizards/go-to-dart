package mo

import (
	"fmt"
	"github.com/11wizards/go-to-dart/generator/format"
	"go/ast"
)

type OptionFormatter struct {
	format.TypeFormatterBase
}

func (f *OptionFormatter) under(expr ast.Expr) (format.TypeFormatter, ast.Expr) {
	e := expr.(*ast.IndexExpr).Index
	formatter := f.Registry.GetTypeFormatter(e)
	return formatter, e
}

func (f *OptionFormatter) CanFormat(expr ast.Expr) bool {
	if x, ok := expr.(*ast.IndexExpr); ok {
		if y, ok := x.X.(*ast.SelectorExpr); ok && y.X.(*ast.Ident).Name == "mo" && y.Sel.Name == "Option" {
			return true
		}
	}
	return false
}

func (f *OptionFormatter) Signature(expr ast.Expr) string {
	formatter, expr := f.under(expr)
	return fmt.Sprintf("%s?", formatter.Signature(expr))
}

func (f *OptionFormatter) Declaration(fieldName string, expr ast.Expr) string {
	return fmt.Sprintf("%s %s", f.Signature(expr), fieldName)
}

func (f *OptionFormatter) Constructor(fieldName string, _ ast.Expr) string {
	return "this." + fieldName
}
