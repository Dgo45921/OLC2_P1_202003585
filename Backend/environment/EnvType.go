package environment

type EnvType int

const (
	GLOBAL     EnvType = iota
	IF
	FOR
	WHILE
	METHOD
	FUNC
	SWITCH
	CASE
	ElSE
	ELSEIF
)
