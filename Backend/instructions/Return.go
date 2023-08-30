package instructions

import (
	"PY1/environment"
	"PY1/interfaces"
)

type Return struct {
	Lin int
	Col int
	Exp interfaces.Expression
}

func NewReturn(lin int, col int, exp interfaces.Expression) Return {
	breakInstr := Return{lin, col, exp}
	return breakInstr
}

func (p Return) Execute(ast *environment.AST, env interface{}) interface{} {
	return p
}
