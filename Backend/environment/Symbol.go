package environment



type Symbol struct {
	Lin   int
	Col   int
	Type  TipoExpresion
	Value interface{}
	Const bool
}

func (s Symbol) GetType() string {
	if s.Type == STRING {
		return "STRING"
	}
	if s.Type == INTEGER {
		return "INTEGER"
	}
	if s.Type == FLOAT {
		return "FLOAT"
	}
	if s.Type == CHAR {
		return "CHARACTER"
	}
	if s.Type == BOOLEAN {
		return "BOOLEAN"
	}

	return "NULL"
}