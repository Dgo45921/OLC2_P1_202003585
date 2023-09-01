package instructions

import (
	"PY1/environment"
)

type FuncDec struct {
	Lin        int
	Col        int
	Id         string
	Args       []environment.FuncParam
	ReturnType interface{}
	insBlock   []interface{}
}

func NewFuncDec(lin int, col int, id string, args []environment.FuncParam, ret interface{}, insb []interface{}) FuncDec {
	return FuncDec{lin, col, id, args, ret, insb}
}

func (p FuncDec) Execute(ast *environment.AST, env interface{}) interface{} {

	if env.(environment.Environment).FuncExists(p.Id) {
		ast.SetPrint("Error, funcion ya declarada!\n")
		return nil

	}

	// has a return type
	if p.ReturnType != nil {

		funcval := environment.FunctionSymbol{
			Lin:        p.Lin,
			Col:        p.Col,
			ReturnType: getReturnType(p.ReturnType.(string)),
			Args:       p.Args,
			InsBlock:   p.insBlock,
			StructType: "",
		}
		env.(environment.Environment).SaveFunc(p.Id, funcval)

	} else {

		funcval := environment.FunctionSymbol{
			Lin:        p.Lin,
			Col:        p.Col,
			ReturnType: environment.NULL,
			Args:       p.Args,
			InsBlock:   p.insBlock,
			StructType: "",
		}
		env.(environment.Environment).SaveFunc(p.Id, funcval)

	}

	return nil
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
