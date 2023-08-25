package instructions

import (
	"PY1/environment"
	"PY1/interfaces"
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
	env.(environment.Environment).SaveVariable(p.Id, value)
	if _, isString := p.Type.(string); isString {
		if countCharOccurrences(p.Type.(string), ']') == getDepth(value.Value.([]interface{})) {
			var matrixType = getMatrixType(value.Type)
			value.Type = matrixType
			env.(environment.Environment).SaveVariable(p.Id, value)

		} else {
			ast.SetPrint("Error: El tamaño con el que se inicializa no consiste con el tamaño definido\n")
		}
	}

	return nil
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
