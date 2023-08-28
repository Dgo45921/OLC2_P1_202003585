package instructions

import (
	"PY1/environment"
	"PY1/interfaces"
)

type AppendVector struct {
	Lin        int
	Col        int
	Id         string
	Expression interfaces.Expression
}

func NewAppendVector(lin int, col int, id string, val interfaces.Expression) AppendVector {
	asig := AppendVector{lin, col, id, val}
	return asig
}

func (p AppendVector) Execute(ast *environment.AST, env interface{}) interface{} {
	foundVar := env.(environment.Environment).FindVar(p.Id)
	if foundVar.Type == environment.VECTOR_STRUCT || foundVar.Type == environment.VECTOR_INT || foundVar.Type == environment.VECTOR_FLOAT || foundVar.Type == environment.VECTOR_CHAR || foundVar.Type == environment.VECTOR_STRING || foundVar.Type == environment.VECTOR_BOOLEAN || foundVar.Type == environment.VECTOR {
		value := p.Expression.Execute(ast, env)

		if foundVar.Type == environment.VECTOR {
			if value.Type == environment.INTEGER {
				foundVar.Type = environment.VECTOR_INT
				if _, isArray := foundVar.Value.([]interface{}); isArray {
					foundVar.Value = append(foundVar.Value.([]interface{}), value.Value)
				}
				env.(environment.Environment).UpdateVariable(p.Id, foundVar)
			} else if value.Type == environment.FLOAT {
				foundVar.Type = environment.VECTOR_FLOAT
				if _, isArray := foundVar.Value.([]interface{}); isArray {
					foundVar.Value = append(foundVar.Value.([]interface{}), value.Value)
				}
				env.(environment.Environment).UpdateVariable(p.Id, foundVar)
			} else if value.Type == environment.BOOLEAN {
				foundVar.Type = environment.VECTOR_BOOLEAN
				if _, isArray := foundVar.Value.([]interface{}); isArray {
					foundVar.Value = append(foundVar.Value.([]interface{}), value.Value)
				}
				env.(environment.Environment).UpdateVariable(p.Id, foundVar)
			} else if value.Type == environment.CHAR {
				foundVar.Type = environment.VECTOR_CHAR
				if _, isArray := foundVar.Value.([]interface{}); isArray {
					foundVar.Value = append(foundVar.Value.([]interface{}), value.Value)
				}
				env.(environment.Environment).UpdateVariable(p.Id, foundVar)
			} else if value.Type == environment.STRING {
				foundVar.Type = environment.VECTOR_STRING
				if _, isArray := foundVar.Value.([]interface{}); isArray {
					foundVar.Value = append(foundVar.Value.([]interface{}), value.Value)
				}
				env.(environment.Environment).UpdateVariable(p.Id, foundVar)
			} else if value.Type == environment.STRUCT_IMP {
				foundVar.Type = environment.VECTOR_STRUCT
				if _, isArray := foundVar.Value.([]interface{}); isArray {
					foundVar.Value = append(foundVar.Value.([]interface{}), value.Value)
				}
				env.(environment.Environment).UpdateVariable(p.Id, foundVar)
			} else {
				ast.SetPrint("Error: tipo de concatenacion incompatible! \n")
			}
		} else {
			if value.Type == environment.INTEGER && foundVar.Type == environment.VECTOR_INT {
				if _, isArray := foundVar.Value.([]interface{}); isArray {
					foundVar.Value = append(foundVar.Value.([]interface{}), value.Value)
				}
				env.(environment.Environment).UpdateVariable(p.Id, foundVar)
			} else if value.Type == environment.FLOAT && foundVar.Type == environment.VECTOR_FLOAT {
				if _, isArray := foundVar.Value.([]interface{}); isArray {
					foundVar.Value = append(foundVar.Value.([]interface{}), value.Value)
				}
				env.(environment.Environment).UpdateVariable(p.Id, foundVar)
			} else if value.Type == environment.BOOLEAN && foundVar.Type == environment.VECTOR_BOOLEAN {
				if _, isArray := foundVar.Value.([]interface{}); isArray {
					foundVar.Value = append(foundVar.Value.([]interface{}), value.Value)
				}
				env.(environment.Environment).UpdateVariable(p.Id, foundVar)
			} else if value.Type == environment.STRING && foundVar.Type == environment.VECTOR_STRING {
				if _, isArray := foundVar.Value.([]interface{}); isArray {
					foundVar.Value = append(foundVar.Value.([]interface{}), value.Value)
				}
				env.(environment.Environment).UpdateVariable(p.Id, foundVar)
			} else if value.Type == environment.CHAR && foundVar.Type == environment.VECTOR_CHAR {
				if _, isArray := foundVar.Value.([]interface{}); isArray {
					foundVar.Value = append(foundVar.Value.([]interface{}), value.Value)
				}
				env.(environment.Environment).UpdateVariable(p.Id, foundVar)
			} else if value.Type == environment.STRUCT_IMP && foundVar.Type == environment.VECTOR_STRUCT {
				if _, isArray := foundVar.Value.([]interface{}); isArray {
					foundVar.Value = append(foundVar.Value.([]interface{}), value.Value)
				}
				env.(environment.Environment).UpdateVariable(p.Id, foundVar)
			} else {
				ast.SetPrint("Error: tipo de concatenacion incompatible! \n")
			}
		}

	} else {
		ast.SetPrint("Error: La función append solo funciona con vectores! \n")
	}

	return nil

}
