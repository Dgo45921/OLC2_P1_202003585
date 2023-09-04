package environment

type CustomSemanticError struct {
	line, column int
	msg          string
	ttype        string
}

type AST struct {
	Instructions []interface{}
	Print        string
	Errors       []CustomSemanticError
	Symbols      []Symbol
}

func NewAST(inst []interface{}, print string) AST {
	ast := AST{Instructions: inst, Print: print, Errors: make([]CustomSemanticError, 0), Symbols: make([]Symbol, 0)}
	return ast
}

func (a *AST) GetPrint() string {
	return a.Print
}

func (a *AST) SetPrint(ToPrint string) {
	a.Print = a.Print + ToPrint
}

func (a *AST) GetErrors() []CustomSemanticError {
	return a.Errors
}

func (a *AST) SetError(line int, col int, des string) {
	err := CustomSemanticError{
		line:   line,
		column: col,
		msg:    des,
		ttype:  "Semantico",
	}
	a.SetPrint("Error: " + des + "\n")
	a.Errors = append(a.Errors, err)
}
