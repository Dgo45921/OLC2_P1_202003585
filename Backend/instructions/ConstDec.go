package instructions

import (
	"PY1/environment"
	"PY1/interfaces"
)

type ConstDec struct {
	Lin        int
	Col        int
	Id         string
	Type       interface{}
	Expression interface{}
}

func NewConstDec(lin int, col int, id string, tyype interface{}, val interface{}) ConstDec {
	NewConstDeclaration := ConstDec{lin, col, id, tyype, val}
	return NewConstDeclaration
}

func (p ConstDec) Execute(ast *environment.AST, env interface{}) interface{} {
	if env == nil {
		return nil
	}
	if p.Type == nil {

		if _, ok := p.Expression.(interfaces.Expression); ok {
			expression := p.Expression.(interfaces.Expression)
			value := expression.Execute(ast, env)
			value.Const = true
			env.(environment.Environment).SaveVariable(p.Id, value)
			return nil

		}
	}

	if _, ok := p.Expression.(interfaces.Expression); ok {

		expression := p.Expression.(interfaces.Expression)
		value := expression.Execute(ast, env)
		value.Const = true

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
	}

	return nil
}
