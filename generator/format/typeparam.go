package format

import (
	"fmt"
	"go/types"
)

type TypeParamsFormatter struct {
	TypeFormatterBase
}

func (f *TypeParamsFormatter) under(expr types.Type) types.Type {
	if typeParam, ok := expr.(*types.TypeParam); ok {
		return typeParam
	}

	return nil
}

func (f *TypeParamsFormatter) CanFormat(expr types.Type) bool {
	return f.under(expr) != nil
}

func (f *TypeParamsFormatter) Signature(expr types.Type) string {
	typeParam := expr.(*types.TypeParam)
	return typeParam.String()
}

func (f *TypeParamsFormatter) Declaration(fieldName string, expr types.Type) string {
	return fmt.Sprintf("%s %s", f.Signature(expr), fieldName)
}

var _ TypeFormatter = (*TypeParamsFormatter)(nil)
