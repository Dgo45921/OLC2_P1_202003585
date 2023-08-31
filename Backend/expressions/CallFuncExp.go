package expressions

import (
	"PY1/environment"
	"fmt"
)

type CallFuncExp struct {
	Lin        int
	Col        int
	Id         string
	Parameters []environment.FuncArg
}

func NewCallFuncExp(lin int, col int, val string, param []environment.FuncArg) CallFuncExp {
	exp := CallFuncExp{lin, col, val, param}
	return exp
}

func (p CallFuncExp) Execute(ast *environment.AST, env interface{}) environment.Symbol {
	foundFunc, exists := env.(environment.Environment).FindFunc(p.Id)

	if !exists {
		ast.SetPrint("Error, funcion no existente! \n")
		return environment.Symbol{
			Lin:   p.Lin,
			Col:   p.Col,
			Value: nil,
		}

	}

	fmt.Println(foundFunc)

	return environment.Symbol{
		Lin:   p.Lin,
		Col:   p.Col,
		Value: nil,
	}

}
