package generator

import (
	"fmt"
	"github.com/openconfig/goyang/pkg/indent"
	"go/ast"
	"io"
)

func generateFields(wr io.Writer, st *ast.StructType) {
	for _, f := range st.Fields.List {
		generateFieldDeclaration(wr, f)
		fmt.Fprintln(wr, ";")
	}
	fmt.Fprintln(wr)
}

func generateConstructor(wr io.Writer, ts *ast.TypeSpec, st *ast.StructType) {
	fmt.Fprintf(wr, "%s({\n", ts.Name)

	for _, f := range st.Fields.List {
		generateFieldConstrutor(indent.NewWriter(wr, "\t"), f)
		fmt.Fprintln(wr, ",")
	}

	fmt.Fprintf(wr, "});")
	fmt.Fprintln(wr)
	fmt.Fprintln(wr)
}

func generateSerialization(wr io.Writer, ts *ast.TypeSpec) {
	fmt.Fprintf(wr, "Map<String, dynamic> toJson() => _$%sToJson(this);\n\n", ts.Name)
}

func generateDeserialization(wr io.Writer, ts *ast.TypeSpec) {
	fmt.Fprintf(wr, "factory %s.fromJson(Map<String, dynamic> json) => _$%sFromJson(json);\n", ts.Name, ts.Name)
}

func generateDartClass(outputFile io.Writer, ts *ast.TypeSpec, st *ast.StructType) bool {
	fmt.Fprintln(outputFile, "@JsonSerializable()")
	fmt.Fprintf(outputFile, "class %s {\n", ts.Name)

	wr := indent.NewWriter(outputFile, "\t")

	generateFields(wr, st)
	generateConstructor(wr, ts, st)
	generateSerialization(wr, ts)
	generateDeserialization(wr, ts)

	fmt.Fprintln(outputFile, "}")
	fmt.Fprintln(outputFile, "")

	return false
}
