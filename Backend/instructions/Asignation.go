package instructions

import (
	"PY1/environment"
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

		if value.Type == foundVar.Type {
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
