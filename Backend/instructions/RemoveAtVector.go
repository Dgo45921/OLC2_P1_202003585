package instructions

import (
	"PY1/environment"
	"PY1/interfaces"
)

type RemoveAtVector struct {
	Lin   int
	Col   int
	Id    string
	Index interfaces.Expression
}

func NewRemoveAtVector(lin int, col int, id string, index interfaces.Expression) RemoveAtVector {
	asig := RemoveAtVector{lin, col, id, index}
	return asig
}

func (p RemoveAtVector) Execute(ast *environment.AST, env interface{}) interface{} {
	foundVar := env.(environment.Environment).FindVar(p.Id)
	indexval := p.Index.Execute(ast, env)
	if slice, isArray := foundVar.Value.([]interface{}); isArray {
		if indexval.Type == environment.INTEGER {
			if indexval.Value.(int) >= 0 && indexval.Value.(int) < len(slice) && !foundVar.Const {
				foundVar.Value = append(slice[:indexval.Value.(int)], slice[indexval.Value.(int)+1:]...)
				env.(environment.Environment).UpdateVariable(p.Id, foundVar)
			} else {
				ast.SetPrint("Error: Indice no disponible en el vector!\n")
			}
		} else {
			ast.SetPrint("Error: el indice luego de 'at:' debe ser un entero!\n")
		}
	} else {
		ast.SetPrint("Error: la funcion remove(at) solo funciona en vectores!\n")
	}
	return nil

}
