package instructions

import (
	"PY1/environment"
	"PY1/expressions"
	"PY1/interfaces"
)

type VarDec struct {
	Lin        int
	Col        int
	Id         string
	Type       interface{}
	Expression interface{}
}

func NewVarDec(lin int, col int, id string, tyype interface{}, val interface{}) VarDec {
	NewVarDeclaration := VarDec{lin, col, id, tyype, val}
	return NewVarDeclaration
}

func (p VarDec) Execute(ast *environment.AST, env interface{}) interface{} {

	if env.(environment.Environment).VariableExists(p.Id) {
		ast.SetPrint("Error, variable ya declarada!\n")
		return nil
	}

	if p.Type == nil {

		if _, ok := p.Expression.(interfaces.Expression); ok {
			expression := p.Expression.(interfaces.Expression)
			value := expression.Execute(ast, env)
			if value.Type == environment.VECTOR_INT || value.Type == environment.VECTOR_FLOAT || value.Type == environment.VECTOR_STRING || value.Type == environment.VECTOR_CHAR || value.Type == environment.VECTOR_BOOLEAN || value.Type == environment.MATRIX_INT || value.Type == environment.MATRIX_FLOAT || value.Type == environment.MATRIX_STRING || value.Type == environment.MATRIX_CHAR || value.Type == environment.MATRIX_BOOLEAN || value.Type == environment.VECTOR {
				val := expressions.DeepCopyArray(value.Value)
				value.Value = val
				env.(environment.Environment).SaveVariable(p.Id, value)
				return nil
			}

			env.(environment.Environment).SaveVariable(p.Id, value)
			return nil

		}
	}

	if _, ok := p.Expression.(interfaces.Expression); ok {

		expression := p.Expression.(interfaces.Expression)
		value := expression.Execute(ast, env)
		if value.Type == environment.NULL {
			if p.Type == "String" {
				value.Type = environment.STRING
			} else if p.Type == "Int" {
				value.Type = environment.INTEGER
			} else if p.Type == "Float" {
				value.Type = environment.FLOAT
			} else if p.Type == "Bool" {
				value.Type = environment.BOOLEAN
			} else if p.Type == "Character" {
				value.Type = environment.CHAR
			} else {
				typeStruct := env.(environment.Environment).FindVar(p.Type.(string))
				if typeStruct.Type == environment.STRUCT_DEF {
					value = environment.Symbol{Lin: 0, Col: 0, Type: environment.STRUCT_IMP, Value: nil, StructType: p.Type.(string)}
					env.(environment.Environment).SaveVariable(p.Id, value)

				} else {
					ast.SetPrint("Error: Tipo de variable no valida!\n")
				}

			}

			env.(environment.Environment).SaveVariable(p.Id, value)
			return nil
		}

		if p.Type == "String" && value.Type == environment.STRING {
			env.(environment.Environment).SaveVariable(p.Id, value)
			return nil
		} else if p.Type == "Int" && value.Type == environment.INTEGER {
			env.(environment.Environment).SaveVariable(p.Id, value)
			return nil
		} else if p.Type == "Character" && value.Type == environment.CHAR {
			env.(environment.Environment).SaveVariable(p.Id, value)
			return nil
		} else if p.Type == "Float" {
			if value.Type == environment.FLOAT {
				env.(environment.Environment).SaveVariable(p.Id, value)
				return nil

			} else if value.Type == environment.INTEGER {
				if _, ok := value.Value.(int); ok {
					value.Value = float64(value.Value.(int))
					value.Type = environment.FLOAT
					env.(environment.Environment).SaveVariable(p.Id, value)
					return nil

				}

			}

		} else if p.Type == "Bool" && value.Type == environment.BOOLEAN {
			env.(environment.Environment).SaveVariable(p.Id, value)
			return nil
		} else if value.Type == environment.VECTOR_INT || value.Type == environment.VECTOR_FLOAT || value.Type == environment.VECTOR_STRING || value.Type == environment.VECTOR_CHAR || value.Type == environment.VECTOR_BOOLEAN || value.Type == environment.MATRIX_INT || value.Type == environment.MATRIX_FLOAT || value.Type == environment.MATRIX_STRING || value.Type == environment.MATRIX_CHAR || value.Type == environment.MATRIX_BOOLEAN || value.Type == environment.VECTOR || value.Type == environment.VECTOR_STRUCT {
			val := expressions.DeepCopyArray(value.Value)
			value.Value = val
			env.(environment.Environment).SaveVariable(p.Id, value)
			return nil
		} else {
			typeStruct := env.(environment.Environment).FindVar(p.Type.(string))
			if typeStruct.Type == environment.STRUCT_DEF {
				if p.Type.(string) == value.StructType {
					value = environment.Symbol{Lin: p.Lin, Col: p.Col, Type: environment.STRUCT_IMP, Value: value.Value, StructType: p.Type.(string)}
					env.(environment.Environment).SaveVariable(p.Id, value)
					return nil
				} else {
					ast.SetPrint("Error: Tipo de struct distinto al definido!\n")
					return nil
				}

			} else {
				ast.SetPrint("Error: Tipo de variable no valida!\n")
			}

		}
	} else if p.Expression == nil {

		var value = environment.Symbol{Lin: 0, Col: 0, Type: environment.NULL, Value: nil}

		if p.Type == "String" {
			value = environment.Symbol{Lin: 0, Col: 0, Type: environment.STRING, Value: nil}
			env.(environment.Environment).SaveVariable(p.Id, value)
			return nil
		} else if p.Type == "Int" {
			value = environment.Symbol{Lin: 0, Col: 0, Type: environment.INTEGER, Value: nil}
			env.(environment.Environment).SaveVariable(p.Id, value)
			return nil
		} else if p.Type == "Character" {
			value = environment.Symbol{Lin: 0, Col: 0, Type: environment.CHAR, Value: nil}
			env.(environment.Environment).SaveVariable(p.Id, value)
			return nil
		} else if p.Type == "Float" {
			value = environment.Symbol{Lin: 0, Col: 0, Type: environment.FLOAT, Value: nil}
			env.(environment.Environment).SaveVariable(p.Id, value)
			return nil
		} else if p.Type == "Bool" {
			value = environment.Symbol{Lin: 0, Col: 0, Type: environment.BOOLEAN, Value: nil}
			env.(environment.Environment).SaveVariable(p.Id, value)
			return nil
		} else {
			typeStruct := env.(environment.Environment).FindVar(p.Type.(string))
			if typeStruct.Type == environment.STRUCT_DEF {
				value = environment.Symbol{Lin: 0, Col: 0, Type: environment.STRUCT_IMP, Value: nil, StructType: p.Type.(string)}
				env.(environment.Environment).SaveVariable(p.Id, value)

			} else {
				ast.SetPrint("Error: Tipo de variable no valida!\n")
			}

		}

	}

	return nil
}

func (p VarDec) GetVarDec(ast *environment.AST, env interface{}) interface{} {
	if env == nil {
		return nil
	}
	if p.Type == nil {

		if _, ok := p.Expression.(interfaces.Expression); ok {
			expression := p.Expression.(interfaces.Expression)
			value := expression.Execute(ast, env)
			return value
		}
	}

	if _, ok := p.Expression.(interfaces.Expression); ok {

		expression := p.Expression.(interfaces.Expression)
		value := expression.Execute(ast, env)
		if value.Type == environment.NULL {
			if p.Type == "String" {
				value.Type = environment.STRING
			} else if p.Type == "Int" {
				value.Type = environment.INTEGER
			} else if p.Type == "Float" {
				value.Type = environment.FLOAT
			} else if p.Type == "Bool" {
				value.Type = environment.BOOLEAN
			} else if p.Type == "Character" {
				value.Type = environment.CHAR
			}

			return value

		}

		if p.Type == "String" && value.Type == environment.STRING {
			return value

		} else if p.Type == "Int" && value.Type == environment.INTEGER {
			return value

		} else if p.Type == "Character" && value.Type == environment.CHAR {
			return value

		} else if p.Type == "Float" {
			if value.Type == environment.FLOAT {
				return value

			} else if value.Type == environment.INTEGER {
				if _, ok := value.Value.(int); ok {
					value.Value = float64(value.Value.(int))
					value.Type = environment.FLOAT
					return value

				}

			}

		} else if p.Type == "Bool" && value.Type == environment.BOOLEAN {
			return value

		} else {
			ast.SetPrint("Error: Type de variable no coincide con la asignación dada!")
			return nil
		}
	} else if p.Expression == nil {

		var value = environment.Symbol{Lin: 0, Col: 0, Type: environment.NULL, Value: nil}

		if p.Type == "String" {
			value = environment.Symbol{Lin: 0, Col: 0, Type: environment.STRING, Value: nil}
			return value

		} else if p.Type == "Int" {
			value = environment.Symbol{Lin: 0, Col: 0, Type: environment.INTEGER, Value: nil}
			return value

		} else if p.Type == "Character" {
			value = environment.Symbol{Lin: 0, Col: 0, Type: environment.CHAR, Value: nil}
			return value

		} else if p.Type == "Float" {
			value = environment.Symbol{Lin: 0, Col: 0, Type: environment.FLOAT, Value: nil}
			return value

		} else if p.Type == "Bool" {
			value = environment.Symbol{Lin: 0, Col: 0, Type: environment.BOOLEAN, Value: nil}
			return value

		}

	}

	return nil
}
