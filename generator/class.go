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

func generateFields(wr io.Writer, registry *format.TypeFormatterRegistry, opts options.Options, fields []fieldProjection) {
	for _, field := range fields {
		generateFieldDeclaration(wr, field.field, field.tag, registry, opts)
		fmt.Fprintln(wr, ";")
	}
	fmt.Fprintln(wr)
}

func generateConstructor(wr io.Writer, ts *types.TypeName, registry *format.TypeFormatterRegistry, fields []fieldProjection) {
	fmt.Fprintf(wr, "const %s(", ts.Name())

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

func generateEquatable(wr io.Writer, fields []fieldProjection) {
	fmt.Fprintln(wr, "@override")
	fmt.Fprint(wr, "List<Object?> get props => [")
	if len(fields) > 0 {
		fmt.Fprintln(wr)
	}

	iwr := indent.NewWriter(wr, "\t")
	for _, field := range fields {
		fmt.Fprintf(iwr, "%s,\n", format.GetFieldName(field.field))
	}

	fmt.Fprintln(wr, "];")
}

func generateDartClass(outputFile io.Writer, ts *types.TypeName, st *types.Struct, registry *format.TypeFormatterRegistry, opts options.Options) {
	formatter, ok := registry.GetTypeFormatter(ts.Type()).(format.StructFormatter)
	if !ok {
		panic(fmt.Sprintf("expected StructFormatter, got %T", registry.GetTypeFormatter(ts.Type())))
	}

	fields := extractFields(st)

	if len(fields) > 0 {
		fmt.Fprintln(outputFile, "@CopyWith()")
	}
	fmt.Fprintln(outputFile, formatter.Annotation(ts))
	if opts.Mode == options.Firestore {
		fmt.Fprintln(outputFile, "@_TimestampConverter()")
	}
	fmt.Fprintf(outputFile, "class %s extends Equatable {\n", formatter.Name(ts))

	wr := indent.NewWriter(outputFile, "\t")

	generateFields(wr, registry, opts, fields)
	generateConstructor(wr, ts, registry, fields)
	fmt.Fprint(wr, formatter.Serialization(ts))
	fmt.Fprintln(wr, formatter.Deserialization(ts))
	generateEquatable(wr, fields)

	fmt.Fprintln(outputFile, "}")
	fmt.Fprintln(outputFile, "")
}
