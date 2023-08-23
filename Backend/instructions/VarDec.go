package instructions

import (
	"PY1/environment"
	"PY1/interfaces"
)

// VarDec TODO CHECK NOT REDECLARED VARIABLE
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
	if env == nil {
		return nil
	}
	if p.Type == nil {

		if _, ok := p.Expression.(interfaces.Expression); ok {
			expression := p.Expression.(interfaces.Expression)
			value := expression.Execute(ast, env)
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
		}

	}

	return nil
}
