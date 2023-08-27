package expressions

import (
	"PY1/environment"
)

type StructAccess struct {
	Lin      int
	Col      int
	ID       string
	Accesses []string
}

func NewStructAccess(lin int, col int, id string, accesses []string) StructAccess {
	structaccess := StructAccess{lin, col, id, accesses}
	return structaccess
}

// TODO GET THE VARIABLE STRUCT_IMP TO ADD LOGIC

func (p StructAccess) Execute(ast *environment.AST, env interface{}) environment.Symbol {

	return environment.Symbol{Lin: p.Lin, Col: p.Col, Type: environment.NULL, Value: nil}
}
