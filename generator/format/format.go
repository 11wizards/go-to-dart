package format

import (
	"go/ast"
)

type TypeFormatter interface {
	CanFormat(expr ast.Expr) bool
	Signature(expr ast.Expr) string
	Declaration(fieldName string, expr ast.Expr) string
	Constructor(fieldName string, expr ast.Expr) string
}

var Formatters []TypeFormatter

func GetTypeFormatter(expr ast.Expr) TypeFormatter {
	for _, f := range Formatters {
		if f.CanFormat(expr) {
			return f
		}
	}

	panic("no formatter found for type")
}
