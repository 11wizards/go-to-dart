package format

import (
	"fmt"
	"go/ast"
)

type AliasFormatter struct {
	TypeFormatterBase
}

func (f *AliasFormatter) under(expr ast.Expr) *ast.Ident {
	if x, ok := expr.(*ast.Ident); ok && x.Obj != nil {
		if y, ok := x.Obj.Decl.(*ast.TypeSpec); ok {
			if z, ok := y.Type.(*ast.Ident); ok {
				return z
			}
		}
	}
	return nil
}

func (f *AliasFormatter) CanFormat(expr ast.Expr) bool {
	return f.under(expr) != nil
}

func (f *AliasFormatter) Signature(expr ast.Expr) string {
	u := f.under(expr)
	return f.Registry.GetTypeFormatter(u).Signature(u)
}

func (f *AliasFormatter) Declaration(fieldName string, expr ast.Expr) string {
	return fmt.Sprintf("%s %s", f.Signature(expr), fieldName)
}

func (f *AliasFormatter) Constructor(fieldName string, expr ast.Expr) string {
	u := f.under(expr)
	return f.Registry.GetTypeFormatter(u).Constructor(fieldName, u)
}
