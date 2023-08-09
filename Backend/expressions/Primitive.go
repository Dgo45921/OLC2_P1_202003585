package expressions

import (
	"PY1/environment"
)

type Primitive struct {
	Lin   int
	Col   int
	Valor interface{}
	Tipo  environment.TipoExpresion
}

func NewPrimitive(lin int, col int, valor interface{}, tipo environment.TipoExpresion) Primitive {
	exp := Primitive{lin, col, valor, tipo}
	return exp
}

func (p Primitive) Execute(ast *environment.AST, env interface{}) environment.Symbol {
	return environment.Symbol{
		Lin:   p.Lin,
		Col:   p.Col,
		Type:  p.Tipo,
		Value: p.Valor,
	}
}
