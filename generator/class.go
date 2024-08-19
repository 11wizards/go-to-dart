package generator

import (
	"fmt"
	"go/types"
	"io"

	"github.com/openconfig/goyang/pkg/indent"

	"github.com/11wizards/go-to-dart/generator/format"
	"github.com/11wizards/go-to-dart/generator/options"
)

type fieldProjection struct {
	field *types.Var
	tag   string
}

func extractFields(st *types.Struct) []fieldProjection {
	var fields []fieldProjection
	for i := 0; i < st.NumFields(); i++ {
		field := st.Field(i)
		if field.Embedded() {
			if embedStruct, ok := field.Type().Underlying().(*types.Struct); ok {
				embeddedFields := extractFields(embedStruct)
				fields = append(fields, embeddedFields...)
				continue
			}
		}
		tag := st.Tag(i)
		fields = append(fields, fieldProjection{field, tag})
	}
	return fields
}

func generateFields(wr io.Writer, st *types.Struct, registry *format.TypeFormatterRegistry, mode options.Mode) {
	fields := extractFields(st)
	for _, field := range fields {
		generateFieldDeclaration(wr, field.field, field.tag, registry, mode)
		fmt.Fprintln(wr, ";")
	}
	fmt.Fprintln(wr)
}

func generateConstructor(wr io.Writer, ts *types.TypeName, st *types.Struct, registry *format.TypeFormatterRegistry) {
	fmt.Fprintf(wr, "%s(", ts.Name())

	fields := extractFields(st)
	if len(fields) > 0 {
		fmt.Fprintln(wr, "{")
		for _, field := range fields {
			generateFieldConstrutor(indent.NewWriter(wr, "\t"), field.field, registry)
			fmt.Fprintln(wr, ",")
		}
		fmt.Fprint(wr, "}")
	}

	fmt.Fprintf(wr, ");")
	fmt.Fprintln(wr)
	fmt.Fprintln(wr)
}

func generateDartClass(outputFile io.Writer, ts *types.TypeName, st *types.Struct, registry *format.TypeFormatterRegistry, mode options.Mode) {
	formatter, ok := registry.GetTypeFormatter(ts.Type()).(format.StructFormatter)
	if !ok {
		panic(fmt.Sprintf("expected StructFormatter, got %T", registry.GetTypeFormatter(ts.Type())))
	}

	fmt.Fprintln(outputFile, formatter.Annotation(ts))
	if mode == options.Firestore {
		fmt.Fprintln(outputFile, "@_TimestampConverter()")
	}
	fmt.Fprintf(outputFile, "class %s {\n", formatter.Name(ts))

	wr := indent.NewWriter(outputFile, "\t")

	generateFields(wr, st, registry, mode)
	generateConstructor(wr, ts, st, registry)
	fmt.Fprint(wr, formatter.Serialization(ts))
	fmt.Fprint(wr, formatter.Deserialization(ts))

	fmt.Fprintln(outputFile, "}")
	fmt.Fprintln(outputFile, "")
}
