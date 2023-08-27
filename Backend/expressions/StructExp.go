package expressions

import (
	"PY1/environment"
	"fmt"
	"html/template"
	"os"
)

type StructExp struct {
	Lin    int
	Col    int
	ID     string
	Fields map[string]interface{}
}

func NewStructExp(lin int, col int, id string, accesses map[string]interface{}) StructExp {
	structaccess := StructExp{lin, col, id, accesses}
	return structaccess
}

func (p StructExp) Execute(ast *environment.AST, env interface{}) environment.Symbol {
	typeStruct := env.(environment.Environment).FindVar(p.ID)
	if typeStruct.Type == environment.STRUCT_DEF {
		t := template.Must(template.New("").Funcs(template.FuncMap{
			"isEven": isEven,
		}).Parse(templ))

		m := map[string]string{
			"a": "A", "b": "B", "c": "C", "d": "D",
		}
		if err := t.Execute(os.Stdout, m); err != nil {
			panic(err)
		}

		for key, value := range typeStruct.Value.(map[string]interface{}) {
			fmt.Printf("Key: %s, Value: %v\n", key, value)

		}
	} else {
		ast.SetPrint("Error: el struct: " + p.ID + "no existe!\n")
	}

	return environment.Symbol{Lin: p.Lin, Col: p.Col, Type: environment.NULL, Value: nil}
}

func isEven() func() bool {
	e := false
	return func() bool {
		e = !e
		return e
	}
}

const templ = `{{$e := isEven}}
{{- range $k, $v := . -}}
    [even:{{call $e}}] key={{$k}}; value={{$v}}
{{end}}`
