package instructions

import (
	"PY1/environment"
)

type RemoveLastVector struct {
	Lin int
	Col int
	Id  string
}

func NewRemoveLastVector(lin int, col int, id string) RemoveLastVector {
	asig := RemoveLastVector{lin, col, id}
	return asig
}

func (p RemoveLastVector) Execute(ast *environment.AST, env interface{}) interface{} {
	if env.(environment.Environment).VariableExists(p.Id) {
		foundVar := env.(environment.Environment).FindVar(p.Id)
		if foundVar.Const {
			ast.SetPrint("Error: No se puede modificar un vector constante!\n")
			return nil

		}
		if _, isArray := foundVar.Value.([]interface{}); isArray {
			if len(foundVar.Value.([]interface{})) == 0 {
				ast.SetPrint("Error: Vector vacio!\n")
				return nil
			}

			oldSlice := foundVar.Value.([]interface{})
			newSlice := oldSlice[:len(oldSlice)-1]
			foundVar.Value = newSlice
			env.(environment.Environment).UpdateVariable(p.Id, foundVar)
		} else {
			ast.SetPrint("Error: la funcion removelast solo funciona en vectores!\n")
		}
	}else if env.(environment.Environment).ReferenceExists(p.Id) {
		foundVar := env.(environment.Environment).FindReference(p.Id)
		if foundVar.Const {
			ast.SetPrint("Error: No se puede modificar un vector constante!\n")
			return nil

		}
		if _, isArray := foundVar.Value.([]interface{}); isArray {
			if len(foundVar.Value.([]interface{})) == 0 {
				ast.SetPrint("Error: Vector vacio!\n")
				return nil
			}

			oldSlice := foundVar.Value.([]interface{})
			newSlice := oldSlice[:len(oldSlice)-1]
			foundVar.Value = newSlice
			env.(environment.Environment).UpdateReference(p.Id, foundVar)
		} else {
			ast.SetPrint("Error: la funcion removelast solo funciona en vectores!\n")
		}
	}

	return nil

}
