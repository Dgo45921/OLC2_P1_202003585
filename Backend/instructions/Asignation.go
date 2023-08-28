package instructions

import (
	"PY1/environment"
	"PY1/expressions"
	"PY1/interfaces"
)

type Asignation struct {
	Lin        int
	Col        int
	Id         string
	Expression interfaces.Expression
}

func NewAsignation(lin int, col int, id string, val interfaces.Expression) Asignation {
	asig := Asignation{lin, col, id, val}
	return asig
}

func (p Asignation) Execute(ast *environment.AST, env interface{}) interface{} {
	if env == nil {
		return nil
	}
	foundVar := env.(environment.Environment).FindVar(p.Id)
	if !foundVar.Const {
		value := p.Expression.Execute(ast, env)

		if value.Type == environment.NULL {
			foundVar.Value = nil
			env.(environment.Environment).UpdateVariable(p.Id, foundVar)
			return nil

		}

		if foundVar.Type == environment.NULL {
			foundVar.Type = value.Type
			env.(environment.Environment).UpdateVariable(p.Id, foundVar)
			return nil
		}

		if value.Type == foundVar.Type {
			if foundVar.Type == environment.VECTOR_INT || foundVar.Type == environment.VECTOR_FLOAT || foundVar.Type == environment.VECTOR_STRING || foundVar.Type == environment.VECTOR_CHAR || foundVar.Type == environment.VECTOR_BOOLEAN || foundVar.Type == environment.MATRIX_INT || foundVar.Type == environment.MATRIX_FLOAT || foundVar.Type == environment.MATRIX_STRING || foundVar.Type == environment.MATRIX_CHAR || foundVar.Type == environment.MATRIX_BOOLEAN || foundVar.Type == environment.VECTOR {
				foundVar.Value = expressions.DeepCopyArray(value.Value)
				env.(environment.Environment).UpdateVariable(p.Id, foundVar)
				return nil
			}
			foundVar.Value = value.Value
			env.(environment.Environment).UpdateVariable(p.Id, foundVar)

		} else {
			ast.SetPrint("Error:No se puede asignar valor: " + value.GetType() + " a variable " + foundVar.GetType() + "\n")
			foundVar.Value = nil
			env.(environment.Environment).UpdateVariable(p.Id, foundVar)

		}

	} else {
		ast.SetPrint("Error: el valor de una constante no puede ser cambiado\n")
	}

	return nil

}
