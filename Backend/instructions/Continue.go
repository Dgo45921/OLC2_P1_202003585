package instructions

import (
	"PY1/environment"
	"PY1/interfaces"
)

type Continue struct {
	Lin        int
	Col        int
	Expression interfaces.Expression
}

func NewContinue(lin int, col int, exp interfaces.Expression) Continue {
	breakInstr := Continue{lin, col, exp}
	return breakInstr
}

func (p Continue) Execute(ast *environment.AST, env interface{}) interface{} {
	if !env.(environment.Environment).InsideLoop() {
		ast.SetPrint("Error: sentencia continue debe de estar dentro de un ciclo!")
	}
	return p
}
