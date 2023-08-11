package expressions

import (
	"PY1/environment"
)

type VariableAccess struct {
	ID string
}

func NewVariableAccess(id string) VariableAccess {
	exp := VariableAccess{id}
	return exp
}

func (p VariableAccess) Execute(ast *environment.AST, env interface{}) environment.Symbol {
	result := env.(environment.Environment).FindVar(p.ID)
	return result
}
