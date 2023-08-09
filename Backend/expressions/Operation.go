package expressions

import (
	"PY1/environment"
	"PY1/interfaces"
	"fmt"
	"strconv"
)

type Operation struct {
	Lin      int
	Col      int
	LeftExp  interfaces.Expression
	Operator string
	RightExp interfaces.Expression
}

func NewOperation(lin int, col int, Op1 interfaces.Expression, Operador string, Op2 interfaces.Expression) Operation {
	exp := Operation{Lin: lin, Col: col, LeftExp: Op1, Operator: Operador, RightExp: Op2}
	return exp
}

func (o Operation) Execute(ast *environment.AST, env interface{}) environment.Symbol {
	var dominante environment.TipoExpresion

	tablaDominante := [5][5]environment.TipoExpresion{
		//		INTEGER			FLOAT				STRING				BOOLEAN				NULL
		{environment.INTEGER, environment.FLOAT, environment.STRING, environment.BOOLEAN, environment.NULL},
		//FLOAT
		{environment.FLOAT, environment.FLOAT, environment.STRING, environment.NULL, environment.NULL},
		//STRING
		{environment.STRING, environment.STRING, environment.STRING, environment.STRING, environment.NULL},
		//BOOLEAN
		{environment.BOOLEAN, environment.NULL, environment.STRING, environment.BOOLEAN, environment.NULL},
		//NULL
		{environment.NULL, environment.NULL, environment.NULL, environment.NULL, environment.NULL},
	}

	var op1, op2 environment.Symbol
	op1 = o.LeftExp.Execute(ast, env)
	op2 = o.RightExp.Execute(ast, env)

	switch o.Operator {
	case "+":
		{
			//validar tipo dominante
			dominante = tablaDominante[op1.Type][op2.Type]
			//valida el tipo
			if dominante == environment.INTEGER {
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Type: dominante, Value: op1.Value.(int) + op2.Value.(int)}
			} else if dominante == environment.FLOAT {
				val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", op1.Value), 64)
				val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", op2.Value), 64)
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Type: dominante, Value: val1 + val2}
			} else if dominante == environment.STRING {
				r1 := fmt.Sprintf("%v", op1.Value)
				r2 := fmt.Sprintf("%v", op2.Value)
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Type: dominante, Value: r1 + r2}
			} else {
				ast.SetError("ERROR: No es posible sumar")
			}
		}
	case "-":
		{
			dominante = tablaDominante[op1.Type][op2.Type]

			if dominante == environment.INTEGER {
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Type: dominante, Value: op1.Value.(int) - op2.Value.(int)}
			} else if dominante == environment.FLOAT {
				val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", op1.Value), 64)
				val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", op2.Value), 64)
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Type: dominante, Value: val1 - val2}
			} else {
				ast.SetError("ERROR: No es posible restar")
			}
		}
	case "*":
		{
			dominante = tablaDominante[op1.Type][op2.Type]
			if dominante == environment.INTEGER {
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Type: dominante, Value: op1.Value.(int) * op2.Value.(int)}
			} else if dominante == environment.FLOAT {
				val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", op1.Value), 64)
				val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", op2.Value), 64)
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Type: dominante, Value: val1 * val2}
			} else {
				ast.SetError("ERROR: No es posible Multiplicar")
			}
		}
	case "/":
		{
			dominante = tablaDominante[op1.Type][op2.Type]
			if dominante == environment.INTEGER {
				if op2.Value.(int) != 0 {
					return environment.Symbol{Lin: o.Lin, Col: o.Col, Type: dominante, Value: op1.Value.(int) / op2.Value.(int)}
				} else {
					ast.SetError("ERROR: No es posible dividir en cero")
				}

			} else if dominante == environment.FLOAT {
				val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", op1.Value), 64)
				val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", op2.Value), 64)
				if val2 != 0 {
					return environment.Symbol{Lin: o.Lin, Col: o.Col, Type: dominante, Value: val1 / val2}
				} else {
					ast.SetError("ERROR: No es posible dividir en cero")
				}
			} else {
				ast.SetError("ERROR: No es posible Dividir")
			}

		}
	case "<":
		{
			dominante = tablaDominante[op1.Type][op2.Type]
			if dominante == environment.INTEGER {
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Type: environment.BOOLEAN, Value: op1.Value.(int) < op2.Value.(int)}
			} else if dominante == environment.FLOAT {
				val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", op1.Value), 64)
				val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", op2.Value), 64)
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Type: environment.BOOLEAN, Value: val1 < val2}
			} else {
				ast.SetError("ERROR: No es posible comparar <")
			}
		}
	case ">":
		{
			dominante = tablaDominante[op1.Type][op2.Type]
			if dominante == environment.INTEGER {
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Type: environment.BOOLEAN, Value: op1.Value.(int) > op2.Value.(int)}
			} else if dominante == environment.FLOAT {
				val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", op1.Value), 64)
				val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", op2.Value), 64)
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Type: environment.BOOLEAN, Value: val1 > val2}
			} else {
				ast.SetError("ERROR: No es posible comparar >")
			}
		}
	case "<=":
		{
			dominante = tablaDominante[op1.Type][op2.Type]
			if dominante == environment.INTEGER {
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Type: environment.BOOLEAN, Value: op1.Value.(int) <= op2.Value.(int)}
			} else if dominante == environment.FLOAT {
				val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", op1.Value), 64)
				val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", op2.Value), 64)
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Type: environment.BOOLEAN, Value: val1 <= val2}
			} else {
				ast.SetError("ERROR: No es posible comparar <=")
			}
		}
	case ">=":
		{
			dominante = tablaDominante[op1.Type][op2.Type]
			if dominante == environment.INTEGER {
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Type: environment.BOOLEAN, Value: op1.Value.(int) >= op2.Value.(int)}
			} else if dominante == environment.FLOAT {
				val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", op1.Value), 64)
				val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", op2.Value), 64)
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Type: environment.BOOLEAN, Value: val1 >= val2}
			} else {
				ast.SetError("ERROR: No es posible comparar >=")
			}
		}
	case "==":
		{
			if op1.Type == op2.Type {
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Type: environment.BOOLEAN, Value: op1.Value == op2.Value}
			} else {
				ast.SetError("ERROR: No es posible comparar == ")
			}
		}
	case "!=":
		{
			if op1.Type == op2.Type {
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Type: environment.BOOLEAN, Value: op1.Value != op2.Value}
			} else {
				ast.SetError("ERROR: No es posible comparar !=")
			}
		}
	case "&&":
		{
			if (op1.Type == environment.BOOLEAN) && (op2.Type == environment.BOOLEAN) {
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Type: environment.BOOLEAN, Value: op1.Value.(bool) && op2.Value.(bool)}
			} else {
				ast.SetError("ERROR: tipo no compatible &&")
			}
		}
	case "||":
		{
			if (op1.Type == environment.BOOLEAN) && (op2.Type == environment.BOOLEAN) {
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Type: environment.BOOLEAN, Value: op1.Value.(bool) || op2.Value.(bool)}
			} else {
				ast.SetError("ERROR: tipo no compatible ||")
			}
		}
	}

	var result interface{}
	return environment.Symbol{Lin: o.Lin, Col: o.Col, Type: environment.NULL, Value: result}
}
