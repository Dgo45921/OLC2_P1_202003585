package environment

type FuncArg struct {
	Id        string
	Value     interface{}
	Reference bool
}

type FuncParam struct {
	Id        string
	SID 	  string
	Type      string
	Reference bool
}
