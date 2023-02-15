package format

import (
	"fmt"
	"go/ast"
)

type StructFormatter struct {
	TypeFormatterBase
}

func (f *StructFormatter) under(expr ast.Expr) *ast.TypeSpec {
	if x, ok := expr.(*ast.Ident); ok && x.Obj != nil {
		if y, ok := x.Obj.Decl.(*ast.TypeSpec); ok {
			if _, ok := y.Type.(*ast.StructType); ok {
				return y
			}
		}
	}
	return nil
}

func (f *StructFormatter) CanFormat(expr ast.Expr) bool {
	return f.under(expr) != nil
}

func (f *StructFormatter) Signature(expr ast.Expr) string {
	u := f.under(expr)
	return u.Name.Name
}

func (f *StructFormatter) Declaration(fieldName string, expr ast.Expr) string {
	return fmt.Sprintf("%s %s", f.Signature(expr), fieldName)
}

func (f *StructFormatter) Constructor(fieldName string, _ ast.Expr) string {
	return "required this." + fieldName
}
