package instructions

import (
	"PY1/environment"
	"PY1/interfaces"
	"strings"
)

type CallFuncInst struct {
	Lin        int
	Col        int
	Id         string
	Parameters []environment.FuncArg
}

func NewCallFuncInst(lin int, col int, val string, param []environment.FuncArg) CallFuncInst {
	exp := CallFuncInst{lin, col, val, param}
	return exp
}

func (p CallFuncInst) Execute(ast *environment.AST, env interface{}) interface{} {
	foundFunc, exists := env.(environment.Environment).FindFunc(p.Id)

	if !exists {
		ast.SetPrint("Error, funcion no existente! \n")
		return nil

	}
	// checking length of parameters and arguments
	if len(p.Parameters) != len(foundFunc.Args) {
		ast.SetPrint("Error: Cantidad de parametros no coincide! \n")
		return nil
	}
	newEnv := environment.NewEnvironment(env, environment.FUNC)
	// check array of values and types
	for index, _ := range p.Parameters {
		valParameter := p.Parameters[index].Value.(interfaces.Expression).Execute(ast, env)
		if valParameter.Type == environment.VECTOR_STRING || valParameter.Type == environment.VECTOR_STRUCT || valParameter.Type == environment.VECTOR_CHAR || valParameter.Type == environment.VECTOR_FLOAT || valParameter.Type == environment.VECTOR_BOOLEAN || valParameter.Type == environment.VECTOR_INT {
			valParameter.Value = DeepCopyArray(valParameter.Value)
		}
		if foundFunc.Args[index].Id == "_" {
			if getTypeByString(foundFunc.Args[index].Type, ast, env, p.Parameters[index].Value.(interfaces.Expression)) == valParameter.Type {
				isByReference := foundFunc.Args[index].Reference
				if isByReference == p.Parameters[index].Reference {
					if isByReference {
						if env.(environment.Environment).VariableExists(p.Parameters[index].RealId) {
							newEnv.SaveReference(foundFunc.Args[index].SID, valParameter)
						} else {
							ast.SetPrint("Error:La referencia solo sirve con variables!\n")
						}

					} else {
						pivote := valParameter
						newEnv.SaveVariable(foundFunc.Args[index].SID, pivote)

					}

				} else {
					ast.SetPrint("Error: atributos definidos como valor por ref o por valor, no coinciden!\n")
					return nil
				}

			} else {
				ast.SetPrint("Error: tipo de atributo no coincide con el argumento enviado!\n")
				return nil
			}

		} else {
			exists, indexx := checkIfParameterExists(foundFunc.Args, p.Parameters[index].Id)
			if exists {

				if getTypeByString(foundFunc.Args[indexx].Type, ast, env, p.Parameters[index].Value.(interfaces.Expression)) == valParameter.Type {
					isByReference := foundFunc.Args[indexx].Reference
					if isByReference == p.Parameters[index].Reference {
						if isByReference {
							if env.(environment.Environment).VariableExists(p.Parameters[index].RealId) {
								newEnv.SaveReference(foundFunc.Args[index].SID, valParameter)
							} else {
								ast.SetPrint("Error:La referencia solo sirve con variables!\n")
							}

						} else {
							pivote := valParameter
							newEnv.SaveVariable(foundFunc.Args[index].SID, pivote)

						}

					} else {
						ast.SetPrint("Error: atributos definidos como valor por ref o por valor, no coinciden!\n")
						return nil
					}

				} else {
					ast.SetPrint("Error: tipo de atributo no coincide con el argumento enviado!\n")
					return nil

				}

			}

		}

	}

	// setting up the function
	for _, inst := range foundFunc.InsBlock {
		// is not any of that cases
		var response = inst.(interfaces.Instruction).Execute(ast, newEnv)
		if response != nil {
			if _, isReturn := response.(environment.Symbol); isReturn {
				valretorno := response.(environment.Symbol)
				if valretorno.Type == foundFunc.ReturnType {
					if valretorno.Type == environment.STRUCT_IMP {
						founstructdef := newEnv.FindVar(valretorno.StructType)
						if founstructdef.Type == environment.STRUCT_DEF && founstructdef.StructType == valretorno.StructType {
							for index, parameter := range p.Parameters {
								if parameter.Reference {
									_, indexx := checkIfParameterExists(foundFunc.Args, p.Parameters[index].Id)
									newEnv.SetReferenceValues(parameter.RealId, foundFunc.Args[indexx].SID)
								}
							}
						} else {
							ast.SetPrint("Error, el tipo de retorno no coincide con el definido con la funcion!\n")
							return nil
						}
					}

					for index, parameter := range p.Parameters {
						if parameter.Reference {
							_, indexx := checkIfParameterExists(foundFunc.Args, p.Parameters[index].Id)
							newEnv.SetReferenceValues(parameter.RealId, foundFunc.Args[indexx].SID)
						}
					}

					return nil

				} else {
					ast.SetPrint("Error: El tipo de retorno definido en la funcion no coincide con el valor del return\n")

					return nil

				}

			}
		} else {
			continue
		}

	}

	for index, parameter := range p.Parameters {
		_, indexx := checkIfParameterExists(foundFunc.Args, p.Parameters[index].Id)
		newEnv.SetReferenceValues(parameter.RealId, foundFunc.Args[indexx].SID)
	}

	return nil

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

func checkIfParameterExists(arr []environment.FuncParam, str string) (bool, int) {
	for index, element := range arr {
		if element.Id == str {
			return true, index
		}

	}

	return false, 0
}
