package expressions

import (
	"PY1/environment"
)

type CountVector struct {
	Lin int
	Col int
	Id  string
}

func NewCountVector(lin int, col int, val string) CountVector {
	exp := CountVector{lin, col, val}
	return exp
}

func (p CountVector) Execute(ast *environment.AST, env interface{}) environment.Symbol {
	foundVar := env.(environment.Environment).FindVar(p.Id)
	if _, isArray := foundVar.Value.([]interface{}); isArray {
		if foundVar.Type == environment.VECTOR || foundVar.Type == environment.VECTOR_STRING || foundVar.Type == environment.VECTOR_CHAR || foundVar.Type == environment.VECTOR_FLOAT || foundVar.Type == environment.VECTOR_BOOLEAN || foundVar.Type == environment.VECTOR_INT {
			long := len(foundVar.Value.([]interface{}))
			return environment.Symbol{
				Lin:   p.Lin,
				Col:   p.Col,
				Value: long,
				Type:  environment.BOOLEAN,
				Const: false,
			}

		}

	}
	ast.SetPrint("Error: la funcion count solo funciona con vectores!\n")
	return environment.Symbol{
		Lin:   p.Lin,
		Col:   p.Col,
		Value: nil,
	}

}
