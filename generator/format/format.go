package format

import (
	"fmt"
	"github.com/11wizards/go-to-dart/generator/options"
	"go/ast"
)

type TypeFormatter interface {
	SetRegistry(registry *TypeFormatterRegistry)
	CanFormat(expr ast.Expr) bool
	Signature(expr ast.Expr) string
	DefaultValue(expr ast.Expr) string
	Declaration(fieldName string, expr ast.Expr) string
	Constructor(fieldName string, expr ast.Expr) string
}

type TypeFormatterBase struct {
	Registry *TypeFormatterRegistry
	Mode     options.Mode
}

func (t *TypeFormatterBase) SetRegistry(registry *TypeFormatterRegistry) {
	t.Registry = registry
}

type TypeFormatterRegistry struct {
	KnownTypes map[string]*ast.TypeSpec
	Formatters []TypeFormatter
}

func NewTypeFormatterRegistry() *TypeFormatterRegistry {
	return &TypeFormatterRegistry{
		KnownTypes: make(map[string]*ast.TypeSpec),
		Formatters: make([]TypeFormatter, 0),
	}
}
func (t *TypeFormatterRegistry) RegisterTypeFormatter(formatter TypeFormatter) {
	t.Formatters = append(t.Formatters, formatter)
	formatter.SetRegistry(t)
}

func (t *TypeFormatterRegistry) GetTypeFormatter(expr ast.Expr) TypeFormatter {
	// walks the t.Formatters in reverse order
	// so that the last registered formatter is the first to be checked
	for i := len(t.Formatters) - 1; i >= 0; i-- {
		f := t.Formatters[i]
		if f.CanFormat(expr) {
			return f
		}
	}

	panic(fmt.Sprintf("no formatter found for %v", expr))
}
