package format

import (
	"bytes"
	"fmt"
	"go/types"
)

type GenericStructFormatter struct {
	TypeFormatterBase
}

func (f *GenericStructFormatter) under(expr types.Type) types.Type {
	if namedType, ok := expr.(*types.Named); ok && namedType.TypeParams().Len() > 0 {
		if structType, ok := namedType.Underlying().(*types.Struct); ok {
			return structType
		}
	}

	return nil
}

func (f *GenericStructFormatter) nameWithoutTypeParams(expr types.Type) string {
	namedType := expr.(*types.Named)

	if f.Options.Prefix != "" && f.Registry.IsKnownNamedType(namedType) {
		return fmt.Sprintf("%s%s", f.Options.Prefix, namedType.Obj().Name())
	}

	return namedType.Obj().Name()
}

func (f *GenericStructFormatter) CanFormat(expr types.Type) bool {
	return f.under(expr) != nil
}

func (f *GenericStructFormatter) Signature(expr types.Type) string {
	namedType := expr.(*types.Named)
	return fmt.Sprintf("%s%s", f.nameWithoutTypeParams(expr), GenerateTypeParams(f.Registry, namedType))
}

func (f *GenericStructFormatter) Declaration(fieldName string, expr types.Type) string {
	return fmt.Sprintf("%s %s", f.Signature(expr), fieldName)
}

func (f *GenericStructFormatter) Name(expr *types.TypeName, includeTypeParams bool) string {
	if includeTypeParams {
		return f.Signature(expr.Type())
	}

	return f.nameWithoutTypeParams(expr.Type())
}

func (f *GenericStructFormatter) Serialization(expr *types.TypeName) string {
	typeParams := expr.Type().(*types.Named).TypeParams()
	buf := new(bytes.Buffer)
	fmt.Fprint(buf, "Map<String, dynamic> toJson(")

	for i := 0; i < typeParams.Len(); i++ {
		if i > 0 {
			fmt.Fprint(buf, ", ")
		}

		tp := typeParams.At(i)

		fmt.Fprintf(buf, "Object Function(%s) toJson%s", tp, tp)
	}

	fmt.Fprintf(buf, ") => _$%sToJson(this", f.nameWithoutTypeParams(expr.Type()))

	for i := 0; i < typeParams.Len(); i++ {
		tp := typeParams.At(i)

		fmt.Fprintf(buf, ", toJson%s", tp)
	}

	fmt.Fprint(buf, ");\n\n")

	return buf.String()
}

func (f *GenericStructFormatter) Deserialization(expr *types.TypeName) string {
	typeParams := expr.Type().(*types.Named).TypeParams()
	buf := new(bytes.Buffer)
	fmt.Fprintf(buf, "factory %s.fromJson(Map<String, dynamic> json", f.nameWithoutTypeParams(expr.Type()))

	for i := 0; i < typeParams.Len(); i++ {
		tp := typeParams.At(i)
		fmt.Fprintf(buf, ", %s Function(Object? json) fromJson%s", tp, tp)
	}

	fmt.Fprintf(buf, ") => _$%sFromJson(json", f.nameWithoutTypeParams(expr.Type()))

	for i := 0; i < typeParams.Len(); i++ {
		tp := typeParams.At(i)

		fmt.Fprintf(buf, ", fromJson%s", tp)
	}

	fmt.Fprint(buf, ");\n")

	return buf.String()
}

func (f *GenericStructFormatter) Annotation(expr *types.TypeName) string {
	return "@JsonSerializable(explicitToJson: true, genericArgumentFactories: true)"
}

var _ StructFormatter = (*GenericStructFormatter)(nil)
