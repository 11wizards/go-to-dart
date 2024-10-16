package format

import (
	"fmt"
	"go/types"
)

type ConcreteStructFormatter struct {
	TypeFormatterBase
}

func (f *ConcreteStructFormatter) under(expr types.Type) *types.Struct {
	if namedType, ok := expr.(*types.Named); ok {
		if structType, ok := namedType.Underlying().(*types.Struct); ok {
			return structType
		}
	}
	return nil
}

func (f *ConcreteStructFormatter) CanFormat(expr types.Type) bool {
	return f.under(expr) != nil
}

func (f *ConcreteStructFormatter) Signature(expr types.Type) string {
	if f.Options.Prefix != "" && f.Registry.IsKnownNamedType(expr.(*types.Named)) {
		return fmt.Sprintf("%s%s", f.Options.Prefix, expr.(*types.Named).Obj().Name())
	}

	return expr.(*types.Named).Obj().Name()
}

func (f *ConcreteStructFormatter) Declaration(fieldName string, expr types.Type) string {
	return fmt.Sprintf("%s %s", f.Signature(expr), fieldName)
}

func (f *ConcreteStructFormatter) Name(expr *types.TypeName) string {
	return f.Signature(expr.Type())
}

func (f *ConcreteStructFormatter) Serialization(expr *types.TypeName) string {
	return fmt.Sprintf("Map<String, dynamic> toJson() => _$%sToJson(this);\n\n", f.Signature(expr.Type()))
}

func (f *ConcreteStructFormatter) Deserialization(expr *types.TypeName) string {
	return fmt.Sprintf("factory %s.fromJson(Map<String, dynamic> json) => _$%sFromJson(json);\n", f.Signature(expr.Type()), f.Signature(expr.Type()))
}

func (t *ConcreteStructFormatter) Annotation(expr *types.TypeName) string {
	return "@JsonSerializable(explicitToJson: true)"
}

var _ StructFormatter = (*ConcreteStructFormatter)(nil)
