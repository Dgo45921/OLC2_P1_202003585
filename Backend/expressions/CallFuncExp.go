package expressions

import (
	"PY1/environment"
	"PY1/interfaces"
	"fmt"
	"reflect"
	"strings"
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
	// checking length of parameters and arguments
	if len(p.Parameters) != len(foundFunc.Args) {
		ast.SetPrint("Error: Cantidad de parametros no coincide! \n")
		return environment.Symbol{
			Lin:   p.Lin,
			Col:   p.Col,
			Value: nil,
		}
	}
	var newEnv = environment.NewEnvironment(env, environment.FUNC)
	// check array of values and types
	for index, _ := range p.Parameters {
		valParameter := p.Parameters[index].Value.(interfaces.Expression).Execute(ast, env)
		if foundFunc.Args[index].Id == "_" {
			if getTypeByString(foundFunc.Args[index].Type, ast, env, p.Parameters[index].Value.(interfaces.Expression)) == valParameter.Type {
				isByReference := foundFunc.Args[index].Reference
				if isByReference == p.Parameters[index].Reference {
					if isByReference {
						newEnv.SaveVariable(foundFunc.Args[index].SID, valParameter)

					} else {
						pivote := valParameter
						newEnv.SaveVariable(foundFunc.Args[index].SID, pivote)

					}

				} else {
					ast.SetPrint("Error: atributos definidos como valor por ref o por valor, no coinciden!\n")
					return environment.Symbol{
						Lin:   p.Lin,
						Col:   p.Col,
						Value: nil,
					}
				}

			} else {
				ast.SetPrint("Error: tipo de atributo no coincide con el argumento enviado!\n")
				return environment.Symbol{
					Lin:   p.Lin,
					Col:   p.Col,
					Value: nil,
				}

			}

		} else {
			exists, indexx := checkIfParameterExists(foundFunc.Args, p.Parameters[index].Id)
			if exists {

				if getTypeByString(foundFunc.Args[indexx].Type, ast, env, p.Parameters[index].Value.(interfaces.Expression)) == valParameter.Type {
					isByReference := foundFunc.Args[indexx].Reference
					if isByReference == p.Parameters[index].Reference {
						if isByReference {
							newEnv.SaveVariable(foundFunc.Args[index].SID, valParameter)

						} else {
							pivote := valParameter
							newEnv.SaveVariable(foundFunc.Args[index].SID, pivote)

						}

					} else {
						ast.SetPrint("Error: atributos definidos como valor por ref o por valor, no coinciden!\n")
						return environment.Symbol{
							Lin:   p.Lin,
							Col:   p.Col,
							Value: nil,
						}
					}

				} else {
					ast.SetPrint("Error: tipo de atributo no coincide con el argumento enviado!\n")
					return environment.Symbol{
						Lin:   p.Lin,
						Col:   p.Col,
						Value: nil,
					}

				}

			}

		}

	}

	// setting up the function

	for index, _ := range p.Parameters {
		if p.Parameters[index].Reference {
			fmt.Println(foundFunc)
			fmt.Println(newEnv)

		} else {

		}
	}

	return environment.Symbol{
		Lin:   p.Lin,
		Col:   p.Col,
		Value: nil,
	}

}

func getTypeByString(val string, ast *environment.AST, env interface{}, expression interfaces.Expression) environment.TipoExpresion {
	if val == "Int" {
		return environment.INTEGER
	} else if val == "Float" {
		return environment.FLOAT
	} else if val == "Bool" {
		return environment.BOOLEAN
	} else if val == "String" {
		return environment.STRING
	} else if val == "Character" {
		return environment.CHAR
	} else if strings.Contains(val, "[") {
		structExp := expression.Execute(ast, env)
		depth := GetDepth(structExp.Value.([]interface{}))

		if depth == countCharOccurrences(val, '[') {
			if depth == 1 {
				if strings.Contains(val, "Int") {
					return environment.VECTOR_INT
				} else if strings.Contains(val, "Float") {
					return environment.VECTOR_FLOAT
				} else if strings.Contains(val, "Bool") {
					return environment.VECTOR_BOOLEAN
				} else if strings.Contains(val, "String") {
					return environment.VECTOR_STRING
				} else if strings.Contains(val, "Character") {
					return environment.VECTOR_CHAR
				} else {
					StructType := strings.ReplaceAll(strings.ReplaceAll(val, "[", ""), "]", "")
					structcase := env.(environment.Environment).FindVar(StructType)
					if structcase.Type == environment.STRUCT_DEF && structExp.StructType == val {
						return environment.STRUCT_IMP
					} else {
						return environment.NULL
					}
				}

			} else {

				if strings.Contains(val, "Int") {
					return environment.MATRIX_INT
				} else if strings.Contains(val, "Float") {
					return environment.MATRIX_FLOAT
				} else if strings.Contains(val, "Bool") {
					return environment.MATRIX_BOOLEAN
				} else if strings.Contains(val, "String") {
					return environment.MATRIX_STRING
				} else if strings.Contains(val, "Character") {
					return environment.MATRIX_CHAR
				} else {
					return environment.NULL
				}
			}

		}
		return environment.NULL

	} else {
		structExp := expression.Execute(ast, env)
		structcase := env.(environment.Environment).FindVar(val)
		if structcase.Type == environment.STRUCT_DEF && structExp.StructType == val {
			return environment.STRUCT_IMP
		} else {
			return environment.NULL
		}
	}

}
func GetDepth(arr []interface{}) int {
	if len(arr) == 0 {
		return 1
	}

	maxDepth := 0
	for _, item := range arr {
		if nestedArr, ok := item.([]interface{}); ok {
			depth := GetDepth(nestedArr)
			if depth > maxDepth {
				maxDepth = depth
			}
		}
	}

	return maxDepth + 1
}

func deepCopy(src interface{}) interface{} {
	srcValue := reflect.ValueOf(src)
	if srcValue.Kind() != reflect.Struct {
		return nil
	}

	// Create a new instance of the same type
	dest := reflect.New(srcValue.Type()).Elem()

	for i := 0; i < srcValue.NumField(); i++ {
		srcField := srcValue.Field(i)
		destField := dest.Field(i)

		// If the field is a struct, recursively deep copy it
		if srcField.Kind() == reflect.Struct {
			destField.Set(deepCopy(srcField.Interface()).(reflect.Value))
		} else {
			// Copy other types directly
			destField.Set(srcField)
		}
	}

	return dest.Interface()
}

func checkIfParameterExists(arr []environment.FuncParam, str string) (bool, int) {
	for index, element := range arr {
		if element.Id == str {
			return true, index
		}

	}

	return false, 0
}
