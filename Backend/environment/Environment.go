package environment

type Environment struct {
	Prev          interface{}
	SymbolTable   map[string]Symbol
	FunctionTable map[string]FunctionSymbol
	Scope         EnvType
}

func NewEnvironment(prev interface{}, scope EnvType) Environment {
	return Environment{
		Prev:          prev,
		SymbolTable:   make(map[string]Symbol),
		FunctionTable: make(map[string]FunctionSymbol),
		Scope:         scope,
	}
}

func (env Environment) VariableExists(id string) bool {

	if _, ok := env.SymbolTable[id]; ok {
		return true
	}

	return false
}
func (env Environment) SaveVariable(id string, value Symbol) {
	env.SymbolTable[id] = value
}

func (env Environment) FuncExists(id string) bool {

	if _, ok := env.FunctionTable[id]; ok {
		return true
	}

	return false
}
func (env Environment) SaveFunc(id string, value FunctionSymbol) {
	env.FunctionTable[id] = value
}

func (env Environment) SaveStruct(id string, value Symbol) {
	env.SymbolTable[id] = value
}

func (env Environment) UpdateVariable(id string, value Symbol) {
	var envTemporal = env
	for {
		if _, ok := envTemporal.SymbolTable[id]; ok {
			envTemporal.SymbolTable[id] = value
			break
		}
		if envTemporal.Prev != nil {
			envTemporal = envTemporal.Prev.(Environment)
			continue
		}
		break

	}

}

func (env Environment) FindVar(id string) Symbol {
	var envTemporal = env
	for {
		if foundVar, ok := envTemporal.SymbolTable[id]; ok {
			return foundVar
		}
		if envTemporal.Prev != nil {
			envTemporal = envTemporal.Prev.(Environment)
			continue
		}
		return Symbol{Lin: 0, Col: 0, Type: NULL, Value: nil}

	}

}

func (env Environment) FindFunc(id string) (FunctionSymbol, bool) {
	var envTemporal = env
	for {
		if foundVar, ok := envTemporal.FunctionTable[id]; ok {
			return foundVar, true
		}
		if envTemporal.Prev != nil {
			envTemporal = envTemporal.Prev.(Environment)
			continue
		}
		return FunctionSymbol{Lin: 0, Col: 0}, false

	}

}

func (env Environment) InsideLoop() bool {
	var tmpEnv = env
	for {
		if tmpEnv.Scope == WHILE || tmpEnv.Scope == FOR || tmpEnv.Scope == CASE || tmpEnv.Scope == DEFAULT {
			return true
		}
		if tmpEnv.Prev == nil {
			break
		} else {
			tmpEnv = tmpEnv.Prev.(Environment)
		}
	}
	return false
}
