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

	var conditionResult = p.Condition.Execute(ast, env)
	if conditionResult.Type != environment.BOOLEAN {
		ast.SetPrint("La condicion debe de ser booleana!")
	} else {
		contador := 0
	outerloop:
		for conditionResult.Value == true && contador < 5000 {
			var newEnv = environment.NewEnvironment(env, environment.WHILE)
			for _, inst := range p.insBlock {
				var response = inst.(interfaces.Instruction).Execute(ast, newEnv)
				if response != nil {
					if _, isBreak := response.(Break); isBreak {
						return nil
					} else if _, isContinue := response.(Continue); isContinue {
						continue outerloop
					}
				}

			}
			conditionResult = p.Condition.Execute(ast, newEnv)
			contador++
		}

	}

	return nil
}
