package instructions

import (
	"PY1/environment"
	"PY1/expressions"
	"PY1/interfaces"
)

type VecDec struct {
	Lin     int
	Col     int
	Id      string
	Type    string
	DefType interface{}
	Exp     interface{}
}

func NewVecDec(lin int, col int, id string, tyype string, deftype interface{}, exp interface{}) VecDec {
	NewVecDeclaration := VecDec{lin, col, id, tyype, deftype, exp}
	return NewVecDeclaration
}

func (p VecDec) Execute(ast *environment.AST, env interface{}) interface{} {

	if _, isArray := p.Exp.(expressions.VariableAccess); isArray {
		var response = p.Exp.(interfaces.Expression).Execute(ast, env)
		types := environment.VECTOR_INT
		switch typo := p.Type; typo {
		case "Int":
			types = environment.VECTOR_INT
		case "Float":
			types = environment.VECTOR_FLOAT
		case "Character":
			types = environment.VECTOR_CHAR
		case "Bool":
			types = environment.VECTOR_BOOLEAN
		case "String":
			types = environment.VECTOR_STRING

		}

		if types == response.Type {

			var symbol = environment.Symbol{
				Lin:   p.Lin,
				Col:   p.Col,
				Type:  types,
				Value: response.Value,
				Const: false,
			}
			env.(environment.Environment).SaveVariable(p.Id, symbol)
			return nil

		} else {
			ast.SetPrint("Error: variable no es del mismo tipo de vector definido!\n")
		}

		return nil
	}

	if _, isArray := p.Exp.(expressions.Vector); isArray {
		types := environment.VECTOR_INT
		switch typo := p.Type; typo {
		case "Int":
			types = environment.VECTOR_INT
		case "Float":
			types = environment.VECTOR_FLOAT
		case "Character":
			types = environment.VECTOR_CHAR
		case "Bool":
			types = environment.VECTOR_BOOLEAN
		case "String":
			types = environment.VECTOR_STRING
		default:
			res := env.(environment.Environment).FindVar(p.Type)
			if res.Type == environment.STRUCT_DEF {
				types = environment.VECTOR_STRUCT
			}

		}

		var response = p.Exp.(interfaces.Expression).Execute(ast, env)

		if types == response.Type || response.Type == environment.VECTOR {
			var symbol = environment.Symbol{
				Lin:   p.Lin,
				Col:   p.Col,
				Type:  types,
				Value: response.Value,
				Const: false,
			}
			env.(environment.Environment).SaveVariable(p.Id, symbol)
			return nil

		} else {
			ast.SetPrint("Error: elemento del arreglo, no coincide con el tipo!\n")
		}
		return nil
	}

	if p.DefType != nil && p.Exp == nil {
		typesVector := environment.VECTOR_INT
		switch typo := p.DefType; typo {
		case "Int":
			typesVector = environment.VECTOR_INT
		case "Float":
			typesVector = environment.VECTOR_FLOAT
		case "Character":
			typesVector = environment.VECTOR_CHAR
		case "Bool":
			typesVector = environment.VECTOR_BOOLEAN
		case "String":
			typesVector = environment.VECTOR_STRING

		}
		var emptyArray []interface{} = []interface{}{}
		var symbol = environment.Symbol{
			Lin:   p.Lin,
			Col:   p.Col,
			Type:  typesVector,
			Value: emptyArray,
			Const: false,
		}
		if p.DefType.(string) == p.Type {
			env.(environment.Environment).SaveVariable(p.Id, symbol)
		} else {
			ast.SetPrint("Error: El tipo definido al inicio no es igual al definido por el ultimo\n")
		}

		return nil

	}

	return nil
}

func (p VecDec) GetVecDec(ast *environment.AST, env interface{}) interface{} {

	if _, isArray := p.Exp.(expressions.VariableAccess); isArray {
		var response = p.Exp.(interfaces.Expression).Execute(ast, env)
		types := environment.VECTOR_INT
		switch typo := p.Type; typo {
		case "Int":
			types = environment.VECTOR_INT
		case "Float":
			types = environment.VECTOR_FLOAT
		case "Character":
			types = environment.VECTOR_CHAR
		case "Bool":
			types = environment.VECTOR_BOOLEAN
		case "String":
			types = environment.VECTOR_STRING

		}

		if types == response.Type {

			var symbol = environment.Symbol{
				Lin:   p.Lin,
				Col:   p.Col,
				Type:  types,
				Value: response.Value,
				Const: false,
			}
			return symbol

		} else {
			ast.SetPrint("Error: variable no es del mismo tipo de vector definido!\n")
		}

		return nil
	}

	if _, isArray := p.Exp.(expressions.Vector); isArray {
		types := environment.VECTOR_INT
		switch typo := p.Type; typo {
		case "Int":
			types = environment.VECTOR_INT
		case "Float":
			types = environment.VECTOR_FLOAT
		case "Character":
			types = environment.VECTOR_CHAR
		case "Bool":
			types = environment.VECTOR_BOOLEAN
		case "String":
			types = environment.VECTOR_STRING

		}

		var response = p.Exp.(interfaces.Expression).Execute(ast, env)

		if types == response.Type || response.Type == environment.VECTOR {
			var symbol = environment.Symbol{
				Lin:   p.Lin,
				Col:   p.Col,
				Type:  types,
				Value: response.Value,
				Const: false,
			}
			return symbol

		} else {
			ast.SetPrint("Error: elemento del arreglo, no coincide con el tipo!\n")
			return nil
		}
	}

	if p.DefType != nil && p.Exp == nil {
		typesVector := environment.VECTOR_INT
		switch typo := p.DefType; typo {
		case "Int":
			typesVector = environment.VECTOR_INT
		case "Float":
			typesVector = environment.VECTOR_FLOAT
		case "Character":
			typesVector = environment.VECTOR_CHAR
		case "Bool":
			typesVector = environment.VECTOR_BOOLEAN
		case "String":
			typesVector = environment.VECTOR_STRING

		}
		var emptyArray []interface{} = []interface{}{}
		var symbol = environment.Symbol{
			Lin:   p.Lin,
			Col:   p.Col,
			Type:  typesVector,
			Value: emptyArray,
			Const: false,
		}

		if p.DefType.(string) == p.Type {
			return symbol
		} else {
			ast.SetPrint("Error: El tipo definido al inicio no es igual al definido por el ultimo\n")
			return nil
		}

	}

	return nil
}
