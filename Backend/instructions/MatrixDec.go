package instructions

import (
	"PY1/environment"
	"PY1/interfaces"
)

type MatrixDec struct {
	Lin  int
	Col  int
	Id   string
	Type interface{}
	Def  interfaces.Expression
}

func NewMatrixDec(lin int, col int, id string, tyype interface{}, def interfaces.Expression) MatrixDec {
	NewMatrixDeclaration := MatrixDec{lin, col, id, tyype, def}
	return NewMatrixDeclaration
}

func (p MatrixDec) Execute(ast *environment.AST, env interface{}) interface{} {

	value := p.Def.Execute(ast, env)
	env.(environment.Environment).SaveVariable(p.Id, value)

	return nil
}
