package instructions

import (
	"PY1/environment"
	"PY1/interfaces"
)

type VarDec struct {
	Lin        int
	Col        int
	Id         string
	Type       interface{}
	Expression interfaces.Expression
}

func NewVarDec(lin int, col int, id string, tyype interface{}, val interfaces.Expression) VarDec {
	NewVarDeclaration := VarDec{lin, col, id, tyype, val}
	return NewVarDeclaration
}

func (p VarDec) Execute(ast *environment.AST, env interface{}) interface{} {
	if env == nil {
		return nil
	}

	var value = p.Expression.Execute(ast, env)
	env.(environment.Environment).SaveVariable(p.Id, value)

	return nil
}
