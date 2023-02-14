package format

import (
	"fmt"
	"go/ast"
)

type MapFormatter struct {
}

func (f *MapFormatter) under(expr ast.Expr) (TypeFormatter, TypeFormatter, ast.Expr, ast.Expr) {
	mapExpr := expr.(*ast.MapType)
	keyFormatter := GetTypeFormatter(mapExpr.Key)
	valueFormatter := GetTypeFormatter(mapExpr.Value)
	return keyFormatter, valueFormatter, mapExpr.Key, mapExpr.Value
}

func (f *MapFormatter) CanFormat(expr ast.Expr) bool {
	_, ok := expr.(*ast.MapType)
	return ok
}

func (f *MapFormatter) Signature(expr ast.Expr) string {
	keyFormatter, valueFormatter, keyExpr, valueExpr := f.under(expr)
	return fmt.Sprintf("Map<%s, %s>", keyFormatter.Signature(keyExpr), valueFormatter.Signature(valueExpr))
}

func (f *MapFormatter) Declaration(fieldName string, expr ast.Expr) string {
	return fmt.Sprintf("%s %s", f.Signature(expr), fieldName)
}

func (f *MapFormatter) Constructor(fieldName string, _ ast.Expr) string {
	return "required this." + fieldName
}
