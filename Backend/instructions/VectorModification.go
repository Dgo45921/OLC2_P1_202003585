package instructions

import (
	"PY1/environment"
	"PY1/expressions"
	"PY1/interfaces"
	"fmt"
	"reflect"
)

type VectorModification struct {
	Lin        int
	Col        int
	DestinyID  string
	Indexes    []interface{}
	Expression interfaces.Expression
}

func NewVectorModification(lin int, col int, id string, indexes []interface{}, val interfaces.Expression) VectorModification {
	asig := VectorModification{lin, col, id, indexes, val}
	return asig
}

func (p VectorModification) Execute(ast *environment.AST, env interface{}) interface{} {
	foundVar := env.(environment.Environment).FindVar(p.DestinyID)
	targetValue := p.Expression.Execute(ast, env)
	if foundVar.Type == environment.VECTOR {
		ast.SetPrint("Error: vector vacío!\n")
		return nil
	}

	var indexes = expressions.GetIndexes(p.Indexes, ast, env)

	if !expressions.AllNonNegativeIntegers(indexes) {
		ast.SetPrint("Error: el o los indices deben de ser un entero mayor o igual a 0!\n")
		return nil
	}

	if _, isArray := foundVar.Value.([]interface{}); isArray {
		if foundVar.Type == environment.VECTOR_STRUCT || foundVar.Type == environment.VECTOR_STRUCT || foundVar.Type == environment.VECTOR_STRING || foundVar.Type == environment.VECTOR_CHAR || foundVar.Type == environment.VECTOR_FLOAT || foundVar.Type == environment.VECTOR_BOOLEAN || foundVar.Type == environment.VECTOR_INT || foundVar.Type == environment.MATRIX_INT || foundVar.Type == environment.MATRIX_FLOAT || foundVar.Type == environment.MATRIX_STRING || foundVar.Type == environment.MATRIX_BOOLEAN || foundVar.Type == environment.MATRIX_CHAR {

			_, exists := expressions.GetIndexValue(foundVar.Value, indexes)

			if !exists {
				ast.SetPrint("Error: indice no existente!\n")
				return nil
			}

			fmt.Println(foundVar.Value)

			if setIndexValue(foundVar.Value, targetValue.Value, indexes) {
				ashkur := getCommonType(foundVar.Value)
				if ashkur != nil {
					env.(environment.Environment).UpdateVariable(p.DestinyID, foundVar)
				} else {
					ast.SetPrint("Error: la matriz debe de ser de un solo tipo!\n")
					return nil
				}

			}

		} else {
			ast.SetPrint("Error: el acceso [] solo funciona con vectores o matrices!\n")
			return nil

		}

	} else {
		ast.SetPrint("Error: este tipo de asignacion solo funciona en vectores o matrices! \n")
		return nil
	}

	return nil

}

func setIndexValue(arr interface{}, newValue interface{}, indexes []int) bool {
	if len(indexes) == 0 || reflect.ValueOf(arr).Kind() != reflect.Slice {
		return false
	}

	index := indexes[0]
	if index < 0 || index >= reflect.ValueOf(arr).Len() {
		return false
	}

	if len(indexes) == 1 {
		reflect.ValueOf(arr).Index(index).Set(reflect.ValueOf(newValue))
		return true
	}

	nextLevel := reflect.ValueOf(arr).Index(index).Interface()
	return setIndexValue(nextLevel, newValue, indexes[1:])
}
