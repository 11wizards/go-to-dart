package format

import (
	"fmt"
	"go/types"
)

// InterfaceFormatter handles the "any" type (interface{}) in Go
type InterfaceFormatter struct {
	TypeFormatterBase
}

func (f *InterfaceFormatter) CanFormat(expr types.Type) bool {
	if iface, ok := expr.(*types.Interface); ok {
		// Empty interface check (interface{} or any)
		return iface.Empty()
	}
	return false
}

func (f *InterfaceFormatter) Signature(_ types.Type) string {
	return "Object?"
}

func (f *InterfaceFormatter) Declaration(fieldName string, expr types.Type) string {
	return fmt.Sprintf("%s %s", f.Signature(expr), fieldName)
}

// Mark this as implementing TypeFormatter
var _ TypeFormatter = (*InterfaceFormatter)(nil)
