package expressions

import (
	"PY1/environment"
	"PY1/interfaces"
	"fmt"
)

type StructAccess struct {
	Lin      int
	Col      int
	ID       string
	Accesses []string
}

func NewStructAccess(lin int, col int, id string, accesses []string) StructAccess {
	structaccess := StructAccess{lin, col, id, accesses}
	return structaccess
}


func (p StructAccess) Execute(ast *environment.AST, env interface{}) environment.Symbol {
	foundVar := env.(environment.Environment).FindVar(p.ID)
	if foundVar.Type == environment.STRUCT_IMP {

		for _, acc := range p.Accesses {
			if _, isArrayKeyValue := foundVar.Value.([]environment.KeyValue); isArrayKeyValue {
				for _, kv := range foundVar.Value.([]environment.KeyValue) {
					if kv.Key == acc {
						result, err := searchNestedValue(kv, p.Accesses)
						if err != nil {
							return environment.Symbol{Lin: p.Lin, Col: p.Col, Type: environment.NULL, Value: nil}
						} else {
							result = result.(interfaces.Expression).Execute(ast, env)
							return environment.Symbol{Lin: p.Lin, Col: p.Col, Type: result.(environment.Symbol).Type, Value: result.(environment.Symbol).Value}
						}

					}

				}
			}

		}

	} else {
		return environment.Symbol{Lin: p.Lin, Col: p.Col, Type: environment.NULL, Value: nil}

	}

	return environment.Symbol{Lin: p.Lin, Col: p.Col, Type: environment.NULL, Value: nil}
}

func searchNestedValue(data environment.KeyValue, keys []string) (interface{}, error) {
	stack := []environment.KeyValue{{Key: "", Value: data}}

	for _, key := range keys {
		var newStack []environment.KeyValue

		for _, kv := range stack {
			if nestedData, ok := kv.Value.(environment.KeyValue); ok && nestedData.Key == key {
				newStack = append(newStack, nestedData)
			}
		}

		if len(newStack) == 0 {
			return nil, fmt.Errorf("key '%s' not found", key)
		}

		stack = newStack
	}

	if len(stack) > 0 {
		return stack[len(stack)-1].Value, nil
	}

	return nil, fmt.Errorf("value not found")
}
