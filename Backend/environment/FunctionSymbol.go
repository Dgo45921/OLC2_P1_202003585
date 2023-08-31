package environment

type FunctionSymbol struct {
	Lin        int
	Col        int
	ReturnType TipoExpresion
	Args       []FuncParam
	InsBlock   []interface{}
	StructType string
}

func (s FunctionSymbol) GetType() string {
	if s.ReturnType == STRING {
		return "STRING"
	}
	if s.ReturnType == INTEGER {
		return "INTEGER"
	}
	if s.ReturnType == FLOAT {
		return "FLOAT"
	}
	if s.ReturnType == CHAR {
		return "CHARACTER"
	}
	if s.ReturnType == BOOLEAN {
		return "BOOLEAN"
	}

	return "NULL"
}
