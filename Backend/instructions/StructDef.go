package instructions

import (
	"PY1/environment"
)

type StructDef struct {
	Lin      int
	Col      int
	Id       string
	insBlock []interface{}
	Type     environment.TipoExpresion
}

func NewStructDef(lin int, col int, id string, insBlock []interface{}) StructDef {
	return StructDef{lin, col, id, insBlock, environment.STRUCT_DEF}
}

func (p StructDef) Execute(ast *environment.AST, env interface{}) interface{} {
	if env.(environment.Environment).Prev != nil {
		ast.SetPrint("Error: Los structs solo pueden ser declarados en el ambito global!\n")
		return nil
	}

	structMap := []environment.KeyValue{}

	for _, inst := range p.insBlock {

		if _, isVarDec := inst.(VarDec); isVarDec {
			response := inst.(VarDec).GetVarDec(ast, env)
			if response != nil {
				newKeyValuePair := environment.KeyValue{inst.(VarDec).Id, response}
				structMap = append(structMap, newKeyValuePair)
			} else {
				ast.SetPrint("Error: El tipo de asignacion a un atributo var no fue válida!\n")
				return nil

			}

		} else if _, isConstDec := inst.(ConstDec); isConstDec {

			response := inst.(ConstDec).GetConstDec(ast, env)
			if response != nil {
				newKeyValuePair := environment.KeyValue{inst.(ConstDec).Id, response}
				structMap = append(structMap, newKeyValuePair)
			} else {
				ast.SetPrint("Error: El tipo de asignacion a un atributo const no fue válida!\n")
				return nil

			}

		} else if _, isVecDec := inst.(VecDec); isVecDec {
			response := inst.(VecDec).GetVecDec(ast, env)
			if response != nil {
				newKeyValuePair := environment.KeyValue{inst.(VecDec).Id, response}
				structMap = append(structMap, newKeyValuePair)
			} else {
				ast.SetPrint("Error: El tipo de asignacion a un atributo vector no fue válida!\n")
				return nil

			}

		} else if _, isMatrixDec := inst.(MatrixDec); isMatrixDec {

			response := inst.(MatrixDec).GetMatrixDec(ast, env)
			if response != nil {
				newKeyValuePair := environment.KeyValue{inst.(MatrixDec).Id, response}
				structMap = append(structMap, newKeyValuePair)
			} else {
				ast.SetPrint("Error: El tipo de asignacion a un atributo matriz no fue válida!\n")
				return nil

			}

		}

	}

	newstruct := environment.Symbol{
		Lin:   p.Lin,
		Col:   p.Col,
		Value: structMap,
		Type:  p.Type,
	}

	env.(environment.Environment).SaveStruct(p.Id, newstruct)

	return nil
}
