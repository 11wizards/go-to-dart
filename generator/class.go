package generator

import (
	"fmt"
	"go/types"
	"io"

	"github.com/11wizards/go-to-dart/generator/format"
	"github.com/11wizards/go-to-dart/generator/options"
	"github.com/openconfig/goyang/pkg/indent"
)

func generateFields(wr io.Writer, st *types.Struct, registry *format.TypeFormatterRegistry, mode options.Mode) {
	for i := 0; i < st.NumFields(); i++ {
		field := st.Field(i)
		tag := st.Tag(i)
		generateFieldDeclaration(wr, field, tag, registry, mode)
		fmt.Fprintln(wr, ";")
	}
	fmt.Fprintln(wr)
}

func generateConstructor(wr io.Writer, ts *types.TypeName, st *types.Struct, registry *format.TypeFormatterRegistry) {
	fmt.Fprintf(wr, "%s(", ts.Name())

	if st.NumFields() > 0 {
		fmt.Fprintln(wr, "{")
		for i := 0; i < st.NumFields(); i++ {
			f := st.Field(i)
			generateFieldConstrutor(indent.NewWriter(wr, "\t"), f, registry)
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
