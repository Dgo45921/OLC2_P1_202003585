package instructions

import (
	"PY1/environment"
	"PY1/interfaces"
	"reflect"
	"strings"
)

type MatrixDec struct {
	Lin  int
	Col  int
	Id   string
	Type interface{}
	Def  interfaces.Expression
}

func NewMatrixDec(lin int, col int, id string, tyype interface{}, def interfaces.Expression) MatrixDec {
	NewMatrixDeclaration := MatrixDec{lin, col, id, tyype, def}
	return NewMatrixDeclaration
}

func (p MatrixDec) Execute(ast *environment.AST, env interface{}) interface{} {

	value := p.Def.Execute(ast, env)
	deepness := getDepth(value.Value.([]interface{}))
	if _, isString := p.Type.(string); isString {
		if countCharOccurrences(p.Type.(string), ']') == deepness {
			if deepness == 1 {
				var matrixType = getMatrixType(value.Type)
				value.Type = matrixType
				env.(environment.Environment).SaveVariable(p.Id, value)
				return nil
			}

			arrayType := getArrayType(value.Value)
			if arrayType != nil {
				if arrayType == reflect.TypeOf(1) && strings.Contains(p.Type.(string), "Int") {
					value.Type = environment.MATRIX_INT
				} else if arrayType == reflect.TypeOf("x") && strings.Contains(p.Type.(string), "String") {
					value.Type = environment.MATRIX_STRING
				} else if arrayType == reflect.TypeOf('c') && strings.Contains(p.Type.(string), "Character") {
					value.Type = environment.MATRIX_CHAR
				} else if arrayType == reflect.TypeOf(5.121) && strings.Contains(p.Type.(string), "Float") {
					value.Type = environment.MATRIX_FLOAT
				} else if arrayType == reflect.TypeOf(false) && strings.Contains(p.Type.(string), "Bool") {
					value.Type = environment.MATRIX_BOOLEAN
				} else {
					ast.SetPrint("Error: Matriz no coincide con tipo de dato definido!\n")
					return nil
				}
				env.(environment.Environment).SaveVariable(p.Id, value)
				return nil
			} else {
				ast.SetPrint("Error: Matriz con varios tipos de dato!\n")
				return nil
			}

		} else {
			ast.SetPrint("Error: El tamaño con el que se inicializa la matriz no consiste con el tamaño definido\n")
			return nil
		}
	} else {
		arrayType := getArrayType(value.Value)
		if arrayType != nil {
			if arrayType == reflect.TypeOf(1) {
				value.Type = environment.MATRIX_INT
			} else if arrayType == reflect.TypeOf("x") {
				value.Type = environment.MATRIX_STRING
			} else if arrayType == reflect.TypeOf('c') {
				value.Type = environment.MATRIX_CHAR
			} else if arrayType == reflect.TypeOf(5.121) {
				value.Type = environment.MATRIX_FLOAT
			} else if arrayType == reflect.TypeOf(false) {
				value.Type = environment.MATRIX_BOOLEAN
			} else {
				ast.SetPrint("Error: Matriz no coincide con tipo de dato definido!\n")
				return nil
			}
			env.(environment.Environment).SaveVariable(p.Id, value)
			return nil
		} else {
			ast.SetPrint("Error: Matriz con varios tipos de dato!\n")
			return nil
		}

	}

}

func getDepth(arr []interface{}) int {
	if len(arr) == 0 {
		return 1
	}

	maxDepth := 0
	for _, item := range arr {
		if nestedArr, ok := item.([]interface{}); ok {
			depth := getDepth(nestedArr)
			if depth > maxDepth {
				maxDepth = depth
			}
		}
	}

	return maxDepth + 1
}
func countCharOccurrences(input string, char rune) int {
	count := 0
	for _, c := range input {
		if c == char {
			count++
		}
	}
	return count
}

func getMatrixType(typee environment.TipoExpresion) environment.TipoExpresion {
	if typee == environment.VECTOR_INT {
		return environment.MATRIX_INT
	} else if typee == environment.VECTOR_FLOAT {
		return environment.MATRIX_FLOAT
	} else if typee == environment.VECTOR_BOOLEAN {
		return environment.MATRIX_BOOLEAN
	} else if typee == environment.VECTOR_CHAR {
		return environment.MATRIX_CHAR
	} else if typee == environment.VECTOR_STRING {
		return environment.MATRIX_STRING
	}
	return environment.NULL
}

func getArrayType(arr interface{}) reflect.Type {
	var elementType reflect.Type

	switch arr.(type) {
	case []interface{}:
		for _, item := range arr.([]interface{}) {
			itemType := getArrayType(item)
			if elementType == nil {
				elementType = itemType
			} else if elementType != itemType {
				return nil
			}
		}
	default:
		elementType = reflect.TypeOf(arr)

		// Check for string type with length 1
		if elementType == reflect.TypeOf("") {
			str := arr.(string)
			if len(str) == 1 {
				elementType = reflect.TypeOf('c') // Replace 'c' with the actual character
			}
		}
	}

	return elementType
}
