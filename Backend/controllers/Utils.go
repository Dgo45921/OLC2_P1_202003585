package controllers

import (
	"fmt"
)

func getVizCode(lexerErrors *CustomLexicalErrorListener, parserErrors *CustomSyntaxErrorListener) string {
	vizcode := "digraph G {\n  node [shape=plaintext];\n  labelloc=\"t\";\n  label=\"Reporte de Errores\";\n  \n  a [label=<\n    <table border=\"0\" cellborder=\"1\" cellspacing=\"0\">\n      <tr><td bgcolor=\"lightgrey\"><b>Tipo</b></td><td bgcolor=\"lightgrey\"><b>Descripción</b></td><td bgcolor=\"lightgrey\"><b>Línea</b></td><td bgcolor=\"lightgrey\"><b>Columna</b></td></tr>"
	fmt.Println(lexerErrors.Errors)
	fmt.Println(parserErrors.Errors)

	vizcode += "</table>\n  >];\n}\n"

	return vizcode
}
