package instructions

import (
	"PY1/environment"
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
	if env.(environment.Environment).VariableExists(p.DestinyID) {
		foundVar := env.(environment.Environment).FindVar(p.DestinyID)
		if foundVar.Const {
			ast.SetPrint("Error: No se puede modificar un vector constante!\n")
			return nil

		}
		targetValue := p.Expression.Execute(ast, env)
		if foundVar.Type == environment.VECTOR {
			ast.SetPrint("Error: vector vacío!\n")
			return nil
		}

		var indexes = GetIndexes(p.Indexes, ast, env)

		if !AllNonNegativeIntegers(indexes) {
			ast.SetPrint("Error: el o los indices deben de ser un entero mayor o igual a 0!\n")
			return nil
		}

		if _, isArray := foundVar.Value.([]interface{}); isArray {
			if foundVar.Type == environment.VECTOR_STRUCT || foundVar.Type == environment.VECTOR_STRING || foundVar.Type == environment.VECTOR_CHAR || foundVar.Type == environment.VECTOR_FLOAT || foundVar.Type == environment.VECTOR_BOOLEAN || foundVar.Type == environment.VECTOR_INT || foundVar.Type == environment.MATRIX_INT || foundVar.Type == environment.MATRIX_FLOAT || foundVar.Type == environment.MATRIX_STRING || foundVar.Type == environment.MATRIX_BOOLEAN || foundVar.Type == environment.MATRIX_CHAR {

				_, exists := GetIndexValue(foundVar.Value, indexes)

				if !exists {
					ast.SetPrint("Error: indice no existente!\n")
					return nil
				}

				fmt.Println(foundVar.Value)

				newValue := DeepCopyArray(foundVar.Value)
				if setIndexValue(newValue, targetValue.Value, indexes) {
					ashkur := getCommonType(newValue)
					if ashkur != nil {
						foundVar.Value = newValue
						env.(environment.Environment).UpdateVariable(p.DestinyID, foundVar)
					} else {
						ast.SetPrint("Error: la matriz o vector debe de ser de un solo tipo!\n")
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
	} else if env.(environment.Environment).ReferenceExists(p.DestinyID) {
		foundVar := env.(environment.Environment).FindReference(p.DestinyID)
		if foundVar.Const {
			ast.SetPrint("Error: No se puede modificar un vector constante!\n")
			return nil

		}
		targetValue := p.Expression.Execute(ast, env)
		if foundVar.Type == environment.VECTOR {
			ast.SetPrint("Error: vector vacío!\n")
			return nil
		}

		var indexes = GetIndexes(p.Indexes, ast, env)

		if !AllNonNegativeIntegers(indexes) {
			ast.SetPrint("Error: el o los indices deben de ser un entero mayor o igual a 0!\n")
			return nil
		}

		if _, isArray := foundVar.Value.([]interface{}); isArray {
			if foundVar.Type == environment.VECTOR_STRUCT || foundVar.Type == environment.VECTOR_STRING || foundVar.Type == environment.VECTOR_CHAR || foundVar.Type == environment.VECTOR_FLOAT || foundVar.Type == environment.VECTOR_BOOLEAN || foundVar.Type == environment.VECTOR_INT || foundVar.Type == environment.MATRIX_INT || foundVar.Type == environment.MATRIX_FLOAT || foundVar.Type == environment.MATRIX_STRING || foundVar.Type == environment.MATRIX_BOOLEAN || foundVar.Type == environment.MATRIX_CHAR {

				_, exists := GetIndexValue(foundVar.Value, indexes)

				if !exists {
					ast.SetPrint("Error: indice no existente!\n")
					return nil
				}

				fmt.Println(foundVar.Value)

				newValue := DeepCopyArray(foundVar.Value)
				if setIndexValue(newValue, targetValue.Value, indexes) {
					ashkur := getCommonType(newValue)
					if ashkur != nil {
						foundVar.Value = newValue
						env.(environment.Environment).UpdateReference(p.DestinyID, foundVar)
					} else {
						ast.SetPrint("Error: la matriz o vector debe de ser de un solo tipo!\n")
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

func GetIndexes(val []interface{}, ast *environment.AST, env interface{}) []int {
	indexes := make([]int, len(val))
	for i, index := range val {
		var retrievedVal = index.(interfaces.Expression).Execute(ast, env).Value
		indexes[i] = retrievedVal.(int)
	}

	return indexes
}

func AllNonNegativeIntegers(arr []int) bool {
	for _, num := range arr {
		if num < 0 {
			return false
		}
	}
	return true
}

func GetIndexValue(arr interface{}, indexes []int) (interface{}, bool) {
	if len(indexes) == 0 || reflect.ValueOf(arr).Kind() != reflect.Slice {
		return nil, false
	}

	index := indexes[0]
	if index < 0 || index >= reflect.ValueOf(arr).Len() {
		return nil, false
	}

	if len(indexes) == 1 {
		return reflect.ValueOf(arr).Index(index).Interface(), true
	}

	nextLevel := reflect.ValueOf(arr).Index(index).Interface()
	return GetIndexValue(nextLevel, indexes[1:])
}
