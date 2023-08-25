package environment

type TipoExpresion int

const (
	INTEGER TipoExpresion = iota //0
	FLOAT                        //1
	STRING                       //2
	CHAR                         //2
	BOOLEAN                      //3
	VECTOR
	VECTOR_INT //4
	VECTOR_FLOAT
	VECTOR_CHAR
	VECTOR_STRING
	VECTOR_BOOLEAN
	MATRIX
	NULL //5
)
