package instructions

import (
	"PY1/environment"
	"PY1/interfaces"
)

type If struct {
	Lin        int
	Col        int
	Condition  interfaces.Expression
	TrueBlock  []interface{}
	FalseBlock []interface{}
}

func NewIf(lin int, col int, condition interfaces.Expression, trueb []interface{}, falseb []interface{}) If {
	return If{lin, col, condition, trueb, falseb}
}

func (p If) Execute(ast *environment.AST, env interface{}) interface{} {

	var expResult = p.Condition.Execute(ast, env)
	if expResult.Type != environment.BOOLEAN {
		ast.SetPrint("Error: La expresion dentro del if debe ser booleana")
		return nil
	}

	if expResult.Value.(bool) == true {
		var newEnv = environment.NewEnvironment(env)
		for _, inst := range p.TrueBlock {
			inst.(interfaces.Instruction).Execute(ast, newEnv)
		}
	} else {
		if len(p.FalseBlock) > 0 {
			if nestedArray, isArray := p.FalseBlock[0].([]interface{}); isArray {
				var newEnv = environment.NewEnvironment(env)
				for _, inst := range nestedArray {
					if instruction, isInstruction := inst.(interfaces.Instruction); isInstruction {
						instruction.Execute(ast, newEnv)
					}
				}
			} else {
				var newEnv = environment.NewEnvironment(env)
				for _, inst := range p.FalseBlock {
					inst.(interfaces.Instruction).Execute(ast, newEnv)
				}
			}
		}
	}

	return nil
}
