package instructions

import (
	"PY1/environment"
)

type Break struct {
	Lin int
	Col int
}

func NewBreak(lin int, col int) Break {
	breakInstr := Break{lin, col}
	return breakInstr
}

func (p Break) Execute(ast *environment.AST, env interface{}) interface{} {
	if !env.(environment.Environment).InsideLoop() {
		ast.SetPrint("Error: sentencia break debe de estar dentro de un ciclo!\n")
	}
	return p
}
