package instructions

import (
	"PY1/environment"
	"PY1/interfaces"
)

type StructDec struct {
	Lin int
	Col int
	Id  string
	Exp interfaces.Expression
}

func NewStructDec(lin int, col int, id string, expr interfaces.Expression) StructDec {
	return StructDec{lin, col, id, expr}
}

func (p StructDec) Execute(ast *environment.AST, env interface{}) interface{} {
	newStruct := p.Exp.Execute(ast, env)
	env.(environment.Environment).SaveStruct(p.Id, newStruct)
	return nil
}
