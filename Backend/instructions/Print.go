package instructions

import (
	"PY1/environment"
	"PY1/interfaces"
	"fmt"
	"strings"
)

type Print struct {
	Lin   int
	Col   int
	Value []interface{}
}

func NewPrint(lin int, col int, val []interface{}) Print {
	return Print{lin, col, val}
}

func (p Print) Execute(ast *environment.AST, env interface{}) interface{} {
	var printedValues []string

	for _, val := range p.Value {
		if expr, ok := val.(interfaces.Expression); ok {
			valueToPrint := expr.Execute(ast, env).Value
			printedValues = append(printedValues, fmt.Sprintf("%v", valueToPrint))
		} else {
			// Handle non-expression values as needed
		}
	}

	consoleOut := strings.Join(printedValues, " ")
	ast.SetPrint(consoleOut + "\n")

	return nil
}
