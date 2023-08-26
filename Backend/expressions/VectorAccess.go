package expressions

import (
	"PY1/environment"
	"PY1/interfaces"
	"reflect"
)

type VectorAccess struct {
	Lin   int
	Col   int
	Id    string
	Index []interface{}
}

func NewVectorAccess(lin int, col int, id string, index []interface{}) VectorAccess {
	exp := VectorAccess{lin, col, id, index}
	return exp
}

func (p VectorAccess) Execute(ast *environment.AST, env interface{}) environment.Symbol {
	foundVar := env.(environment.Environment).FindVar(p.Id)
	if foundVar.Type == environment.VECTOR {
		ast.SetPrint("Error: vector vacío!\n")
		return environment.Symbol{
			Lin:   p.Lin,
			Col:   p.Col,
			Value: nil,
		}
	}

	var indexes = getIndexes(p.Index, ast, env)

	if !allNonNegativeIntegers(indexes) {
		ast.SetPrint("Error: el o los indices deben de ser un entero mayor o igual a 0!\n")
		return environment.Symbol{
			Lin:   p.Lin,
			Col:   p.Col,
			Value: nil,
		}
	}

	if _, isArray := foundVar.Value.([]interface{}); isArray {
		if foundVar.Type == environment.VECTOR_STRING || foundVar.Type == environment.VECTOR_CHAR || foundVar.Type == environment.VECTOR_FLOAT || foundVar.Type == environment.VECTOR_BOOLEAN || foundVar.Type == environment.VECTOR_INT || foundVar.Type == environment.MATRIX_INT || foundVar.Type == environment.MATRIX_FLOAT || foundVar.Type == environment.MATRIX_STRING || foundVar.Type == environment.MATRIX_BOOLEAN || foundVar.Type == environment.MATRIX_CHAR {

			val, exists := getIndexValue(foundVar.Value, indexes)

			if !exists {
				ast.SetPrint("Error: indice no existente!\n")
				return environment.Symbol{
					Lin:   p.Lin,
					Col:   p.Col,
					Value: nil,
				}
			}

			var accesstype = environment.INTEGER
			if foundVar.Value == environment.VECTOR_INT || foundVar.Value == environment.MATRIX_INT {
				accesstype = environment.INTEGER
			} else if foundVar.Value == environment.VECTOR_FLOAT || foundVar.Value == environment.MATRIX_FLOAT {
				accesstype = environment.FLOAT
			} else if foundVar.Value == environment.VECTOR_STRING || foundVar.Value == environment.MATRIX_STRING {
				accesstype = environment.STRING
			} else if foundVar.Value == environment.VECTOR_CHAR || foundVar.Value == environment.MATRIX_CHAR {
				accesstype = environment.CHAR
			} else if foundVar.Value == environment.VECTOR_BOOLEAN || foundVar.Value == environment.MATRIX_BOOLEAN {
				accesstype = environment.BOOLEAN
			}

			return environment.Symbol{
				Lin:   p.Lin,
				Col:   p.Col,
				Value: val,
				Type:  accesstype,
				Const: false,
			}

		} else {
			ast.SetPrint("Error: el acceso [] solo funciona con vectores o matrices!\n")
			return environment.Symbol{
				Lin:   p.Lin,
				Col:   p.Col,
				Value: nil,
			}

		}

	}
	ast.SetPrint("Error: el acceso [] solo funciona con vectores o matrices!\n")
	return environment.Symbol{
		Lin:   p.Lin,
		Col:   p.Col,
		Value: nil,
	}

}

func getIndexValue(arr interface{}, indexes []int) (interface{}, bool) {
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
	return getIndexValue(nextLevel, indexes[1:])
}

func getIndexes(val []interface{}, ast *environment.AST, env interface{}) []int {
	indexes := make([]int, len(val))
	for i, index := range val {
		var retrievedVal = index.(interfaces.Expression).Execute(ast, env).Value
		indexes[i] = retrievedVal.(int)
	}

	return indexes
}

func allNonNegativeIntegers(arr []int) bool {
	for _, num := range arr {
		if num < 0 {
			return false
		}
	}
	return true
}
