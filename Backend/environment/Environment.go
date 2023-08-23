package environment

type Environment struct {
	Prev        interface{}
	SymbolTable map[string]Symbol
	Scope       EnvType
}

func NewEnvironment(prev interface{}, scope EnvType) Environment {
	return Environment{
		Prev:        prev,
		SymbolTable: make(map[string]Symbol),
		Scope:       scope,
	}
}

func (env Environment) SaveVariable(id string, value Symbol) {
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
