package interfaces

import "PY1/environment"

type Expression interface {
	Ejecutar(ast *environment.AST, env interface{}) environment.Symbol
}
