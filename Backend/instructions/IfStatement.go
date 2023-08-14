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
	var newEnv = environment.NewEnvironment(env)
	var expResult = p.Condition.Execute(ast, newEnv)
	if expResult.Type != environment.BOOLEAN {
		ast.SetPrint("Error: La expresion dentro del if debe ser booleana")
		return nil
	}

	if expResult.Value.(bool) == true {
		for _, inst := range p.TrueBlock {
			inst.(interfaces.Instruction).Execute(ast, newEnv)
		}
	} else {
		// Check if FalseBlock contains instructions
		if len(p.FalseBlock) > 0 {
			// Assuming the first element of FalseBlock is an array
			if nestedArray, isArray := p.FalseBlock[0].([]interface{}); isArray {
				for _, inst := range nestedArray {
					if instruction, isInstruction := inst.(interfaces.Instruction); isInstruction {
						instruction.Execute(ast, newEnv)
					}
				}
			} else {
				for _, inst := range p.FalseBlock {
					inst.(interfaces.Instruction).Execute(ast, newEnv)
				}
			}
		}
	}

	return nil
}
