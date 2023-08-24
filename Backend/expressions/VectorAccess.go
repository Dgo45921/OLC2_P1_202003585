package expressions

import (
	"PY1/environment"
	"PY1/interfaces"
)

type VectorAccess struct {
	Lin   int
	Col   int
	Id    string
	Index interfaces.Expression
}

func NewVectorAccess(lin int, col int, id string, index interfaces.Expression) VectorAccess {
	exp := VectorAccess{lin, col, id, index}
	return exp
}

func (p VectorAccess) Execute(ast *environment.AST, env interface{}) environment.Symbol {
	foundVar := env.(environment.Environment).FindVar(p.Id)
	if foundVar.Type == environment.VECTOR {
		ast.SetPrint("Error: vector vacío!\n")
		return environment.Symbol{
			Lin:   p.Lin,
			Col:   p.Col,
			Value: nil,
		}
	}

	var index = p.Index.Execute(ast, env)

	if index.Type != environment.INTEGER {
		ast.SetPrint("Error: el índice debe de ser un entero!\n")
		return environment.Symbol{
			Lin:   p.Lin,
			Col:   p.Col,
			Value: nil,
		}
	}

	if _, isArray := foundVar.Value.([]interface{}); isArray {
		if foundVar.Type == environment.VECTOR_STRING || foundVar.Type == environment.VECTOR_CHAR || foundVar.Type == environment.VECTOR_FLOAT || foundVar.Type == environment.VECTOR_BOOLEAN || foundVar.Type == environment.VECTOR_INT {
			val := foundVar.Value.([]interface{})[index.Value.(int)]

			var accesstype = environment.INTEGER
			if foundVar.Value == environment.VECTOR_INT {
				accesstype = environment.INTEGER
			} else if foundVar.Value == environment.VECTOR_FLOAT {
				accesstype = environment.FLOAT
			} else if foundVar.Value == environment.VECTOR_STRING {
				accesstype = environment.STRING
			} else if foundVar.Value == environment.VECTOR_CHAR {
				accesstype = environment.CHAR
			} else if foundVar.Value == environment.VECTOR_BOOLEAN {
				accesstype = environment.BOOLEAN
			}

			return environment.Symbol{
				Lin:   p.Lin,
				Col:   p.Col,
				Value: val,
				Type:  accesstype,
				Const: false,
			}

		} else {
			ast.SetPrint("Error: el acceso [] solo funciona con vectores!\n")
			return environment.Symbol{
				Lin:   p.Lin,
				Col:   p.Col,
				Value: nil,
			}

		}

	}
	ast.SetPrint("Error: el acceso [] solo funciona con vectores!\n")
	return environment.Symbol{
		Lin:   p.Lin,
		Col:   p.Col,
		Value: nil,
	}

}
