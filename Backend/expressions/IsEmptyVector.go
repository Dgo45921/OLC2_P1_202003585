package expressions

import (
	"PY1/environment"
)

type IsEmptyVector struct {
	Lin int
	Col int
	Id  string
}

func NewIsEmptyVector(lin int, col int, val string) IsEmptyVector {
	exp := IsEmptyVector{lin, col, val}
	return exp
}

func (p IsEmptyVector) Execute(ast *environment.AST, env interface{}) environment.Symbol {
	foundVar := env.(environment.Environment).FindVar(p.Id)
	if _, isArray := foundVar.Value.([]interface{}); isArray {
		if foundVar.Type == environment.VECTOR || foundVar.Type == environment.VECTOR_STRING || foundVar.Type == environment.VECTOR_CHAR || foundVar.Type == environment.VECTOR_FLOAT || foundVar.Type == environment.VECTOR_BOOLEAN || foundVar.Type == environment.VECTOR_INT {
			if len(foundVar.Value.([]interface{})) == 0 {
				return environment.Symbol{
					Lin:   p.Lin,
					Col:   p.Col,
					Value: true,
					Type:  environment.BOOLEAN,
					Const: false,
				}
			} else {

				return environment.Symbol{
					Lin:   p.Lin,
					Col:   p.Col,
					Value: false,
					Type:  environment.BOOLEAN,
					Const: false,
				}

			}

		}

	}
	ast.SetPrint("Error: la funcion isEmpty solo funciona con vectores!\n")
	return environment.Symbol{
		Lin:   p.Lin,
		Col:   p.Col,
		Value: nil,
	}

}
