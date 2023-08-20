package instructions

import (
	"PY1/environment"
	"PY1/interfaces"
)

type While struct {
	Lin       int
	Col       int
	Condition interfaces.Expression
	insBlock  []interface{}
}

func NewWhile(lin int, col int, condition interfaces.Expression, insBlock []interface{}) While {
	return While{lin, col, condition, insBlock}
}

func (p While) Execute(ast *environment.AST, env interface{}) interface{} {
	var newEnv = environment.NewEnvironment(env)

	var conditionResult = p.Condition.Execute(ast, newEnv)
	if conditionResult.Type != environment.BOOLEAN {
		ast.SetPrint("La condicion debe de ser booleana!")
	} else {
		contador := 0
		for conditionResult.Value == true && contador < 5000 {
			for _, inst := range p.insBlock {
				inst.(interfaces.Instruction).Execute(ast, newEnv)
			}
			conditionResult = p.Condition.Execute(ast, newEnv)
			contador++
		}

	}

	return nil
}
