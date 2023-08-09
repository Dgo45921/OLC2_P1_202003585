package instructions

import (
	"PY1/environment"
	"PY1/interfaces"
	"fmt"
)

type Print struct {
	Lin   int
	Col   int
	Value interface{}
}

func NewPrint(lin int, col int, val interface{}) Print {
	return Print{lin, col, val}
}

func (p Print) Execute(ast *environment.AST, env interface{}) interface{} {
	valueToPrint := p.Value.(interfaces.Expression).Execute(ast, env)
	consoleOut := fmt.Sprintf("%v", valueToPrint.Value)
	ast.SetPrint(consoleOut + "\n")
	return nil
}
