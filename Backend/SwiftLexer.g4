lexer grammar SwiftLexer;

// --------------- Tokens
// types
RINT: 'Int';
RFLOAT: 'Float';
RBOOL: 'Bool';
RSTRING: 'String';
RCHARACTER: 'Character';

// reserved words
RTRUE: 'true';
RFALSE: 'false';
RPRINT: 'print';
RIF: 'if';
RELSE: 'else';
RWHILE: 'while';
RVAR: 'var';
RLET: 'let';
RNIL: 'nil';
RBREAK: 'break';
RCONTINUE:'continue';

// primitives
NUMBER : [0-9]+ ('.'[0-9]+)?;
STRING: '"' ('\\' ('\\' | '"' | 'r'| 't'| 'n') | ~('\\' | '"'))* '"';

ID: ([a-zA-Z_])[a-zA-Z0-9_]*;

// symbols

UNARYPLUS:      '+=';
UNARYMINUS:     '-=';
DIF:            '!=';
IG_IG:          '==';
NOT:            '!';
OR:             '||';
AND:            '&&';
IG:             '=';
MAY_IG:         '>=';
MEN_IG:         '<=';
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
PTOCOMA:        ';';
DOSPTOS:        ':';
MODULE:         '%' ;
COMA:           ',' ;
QM:             '?' ;

// skip
WHITESPACE: [ \\\r\n\t]+ -> skip;
COMMENT : '/*' .*? '*/' -> skip;
LINE_COMMENT : '//' ~[\r\n]* -> skip;

//fragment
//ESC_SEQ
//    :   '\\' ('\\'|'@'|'['|']'|'.'|'#'|'+'|'-'|'!'|':'|' ')
//    ;
