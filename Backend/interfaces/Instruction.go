package interfaces

import "PY1/environment"

type Instruction interface {
	Ejecutar(ast *environment.AST, env interface{}) interface{}
}
