package instructions

import (
	"PY1/environment"
)

type FuncDec struct {
	Lin        int
	Col        int
	Id         string
	Args       []interface{}
	ReturnType interface{}
	insBlock   []interface{}
}

func NewFuncDec(lin int, col int, id string, args []interface{}, ret interface{}, insb []interface{}) FuncDec {
	return FuncDec{lin, col, id, args, ret, insb}
}

func (p FuncDec) Execute(ast *environment.AST, env interface{}) interface{} {

	if !env.(environment.Environment).FuncExists(p.Id) {
		ast.SetPrint("Error, funcion ya declarada!\n")
		return nil

	}

	// has a return type
	if p.ReturnType != nil {
		containsreturn := checkIfReturn(p.insBlock)

		if !containsreturn {
			ast.SetPrint("Error: La funcion no tiene un return! \n")
			return nil
		}

		funcval := environment.FunctionSymbol{
			Lin:        p.Lin,
			Col:        p.Col,
			ReturnType: getReturnType(p.ReturnType.(string)),
			InsBlock:   p.insBlock,
			StructType: "",
		}
		env.(environment.Environment).SaveFunc(p.Id, funcval)

	} else {

		funcval := environment.FunctionSymbol{
			Lin:        p.Lin,
			Col:        p.Col,
			ReturnType: environment.NULL,
			InsBlock:   p.insBlock,
			StructType: "",
		}
		env.(environment.Environment).SaveFunc(p.Id, funcval)

	}

	return nil
}

func checkIfReturn(block []interface{}) bool {
	for _, inst := range block {
		if _, isReturn := inst.(Return); isReturn {
			return true

		}
	}

	return false
}

func getReturnType(str string) environment.TipoExpresion {
	if str == "String" {
		return environment.STRING
	} else if str == "Int" {
		return environment.INTEGER
	} else if str == "Float" {
		return environment.FLOAT
	} else if str == "Bool" {
		return environment.BOOLEAN
	} else if str == "Character" {
		return environment.CHAR
	}

	return environment.NULL
}
