lexer grammar swiftLexer;

// --------------- Tokens

// keywords

RINT: 'Int';
RFLOAT: 'Float';
RBOOL: 'Bool';
RCHAR: 'Char';
RSTRING: 'String';



RTRUE: 'true';
RFALSE: 'false';
RPRINT: 'print';
RIF: 'if';
RELSE: 'else';
RWHILE: 'while';
RVAR: 'var';
RLET: 'let';
RCASE: 'case';
RSWITCH: 'switch';
RBREAK: 'break';
RDEFAULT: 'default';
RFOR: 'for';
RIN: 'in';
RGUARD : 'guard';
RCONTINUE: 'continue';
RRETURN: 'return';
RSTRUCTU: 'struct';
RFUNC: 'func';




// primitives
NUMBER : [0-9]+ ('.'[0-9]+)?;
STRING: '"'~["]*'"';
ID: [a-zA-Z][a-zA-Z0-9_]*;

// symbols

DIF:      '!=';
IG_IG:          '==';
NOT:            '!';
OR:             '||';
AND:            '&&';
IG:          '=';
MAY_IG:     '>=';
MEN_IG:     '<=';
MAYOR:          '>';
MENOR:          '<';
MUL:            '*';
DIV:            '/';
ADD:            '+';
SUB:            '-';
PARIZQ:         '(';
PARDER:         ')';
LLAVEIZQ:       '{';
LLAVEDER:       '}';

// skip
WHITESPACE: [\r\f\t\n]  -> skip;
COMMENT : '/*' .*? '*/' -> skip;
LINE_COMMENT : '//' ~[\r\n]* -> skip;

//fragment
//ESC_SEQ
//    :   '\\' ('\\'|'@'|'['|']'|'.'|'#'|'+'|'-'|'!'|':'|' ')
//    ;
