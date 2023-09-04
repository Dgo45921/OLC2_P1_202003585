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
		ast.SetError(p.Lin, p.Col, "Funcion ya declarada")
		return nil

	}
	if len(p.Args) > 1 {
		firstName := p.Args[0].SID
		for i := 1; i < len(p.Args); i++ {
			currentName := p.Args[i].SID
			if firstName == currentName {
				ast.SetError(p.Lin, p.Col, "nombre de parametro repetido")
				return nil
			}
		}
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
	} else {
		return environment.STRUCT_IMP
	}
}
