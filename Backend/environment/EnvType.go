package environment

type EnvType int

const (
	GLOBAL EnvType = iota
	IF
	FOR
	WHILE
	METHOD
	FUNC
	CASE
	DEFAULT
	ElSE
	ELSEIF
)
