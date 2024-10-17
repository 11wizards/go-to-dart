package format

import (
	"fmt"
	"go/types"

	"github.com/11wizards/go-to-dart/generator/options"
)

type TypeFormatter interface {
	SetRegistry(registry *TypeFormatterRegistry)
	CanFormat(expr types.Type) bool
	Signature(expr types.Type) string
	DefaultValue(expr types.Type) string
	Declaration(fieldName string, expr types.Type) string
	Constructor(fieldName string, expr types.Type) string
}

type TypeFormatterBase struct {
	Registry *TypeFormatterRegistry
	Options  options.Options
}

func (t *TypeFormatterBase) SetRegistry(registry *TypeFormatterRegistry) {
	t.Registry = registry
}

func (t *TypeFormatterBase) DefaultValue(expr types.Type) string {
	return ""
}

func (f *TypeFormatterBase) Constructor(fieldName string, _ types.Type) string {
	return "required this." + fieldName
}

type StructFormatter interface {
	TypeFormatter
	Name(expr *types.TypeName) string
	Serialization(expr *types.TypeName) string
	Deserialization(expr *types.TypeName) string
	Annotation(expr *types.TypeName) string
}

type TypeFormatterRegistry struct {
	Formatters []TypeFormatter

	knownTypes      map[types.Type]struct{}
	knownNamedTypes map[string]struct{}
}

func (t *TypeFormatterRegistry) AddKnownType(typ types.Type) {
	t.knownTypes[typ] = struct{}{}
	if namedType, ok := typ.(*types.Named); ok {
		t.knownNamedTypes[namedType.Obj().Type().String()] = struct{}{}
	}
}

func (t *TypeFormatterRegistry) IsKnownNamedType(namedType *types.Named) bool {
	name := namedType.Obj().Type().String()
	_, ok := t.knownNamedTypes[name]
	return ok
}

func NewTypeFormatterRegistry() *TypeFormatterRegistry {
	return &TypeFormatterRegistry{
		Formatters:      make([]TypeFormatter, 0),
		knownTypes:      make(map[types.Type]struct{}),
		knownNamedTypes: make(map[string]struct{}),
	}
}

func (t *TypeFormatterRegistry) RegisterTypeFormatter(formatter TypeFormatter) {
	t.Formatters = append(t.Formatters, formatter)
	formatter.SetRegistry(t)
}

func (t *TypeFormatterRegistry) GetTypeFormatter(expr types.Type) TypeFormatter {
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
