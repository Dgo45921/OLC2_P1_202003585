grammar SwiftGrammar; 
// import SwiftLexer; 

options {
  tokenVocab = SwiftLexer;
}

@header {
    import "PY1/interfaces"
    import "PY1/environment"
    import "PY1/expressions"
    import "PY1/instructions"
    import "strings"
}


s returns [[]interface{} code]
: block EOF
    {   
        $code = $block.blk
    }
;

block returns [[]interface{} blk]
@init{
    $blk = []interface{}{}
    var listInt []IInstructionContext
  }
: ins+=instruction+
    {
        listInt = localctx.(*BlockContext).GetIns()
        for _, e := range listInt {
            $blk = append($blk, e.GetInst())
        }
    }
;


arguments returns [[]interface{} args]
@init {
    $args = []interface{}{}

}
    :argument COMA arguments  { $args = append($args, $argument.e)
                                       for _, arg := range $arguments.args {
                                           $args = append($args, arg)
                                       }
                                 }
    |argument { $args = append( $args , $argument.e) }
    | {}
    ;

argument returns [interface{} e]
    : expr { $e = $expr.e; }
    ;

// INSTRUCTIONS
instruction returns [interfaces.Instruction inst]
: printstmt PTOCOMA?  { $inst = $printstmt.prnt}
| ifstmt { }
;

// STATEMENTS----------------------------------------------------------------------------------------------
printstmt returns [interfaces.Instruction prnt]
: RPRINT PARIZQ arguments PARDER { $prnt = instructions.NewPrint($RPRINT.line,$RPRINT.pos,$arguments.args)}
;

ifstmt  
: RIF PARIZQ expr PARDER LLAVEIZQ block LLAVEDER
;

// EXPRESSIONS -----------------------------------------------------------------
expr returns [interfaces.Expression e]
:
 PARIZQ expr PARDER { $e = $expr.e }
| op=(SUB | NOT) left=expr { $e = expressions.NewUnaryOperation($left.start.GetLine(), $left.start.GetColumn(), $left.e, $op.text) }
| left=expr op=(MUL|DIV|MODULE) right=expr { $e = expressions.NewArithmeticOperation($left.start.GetLine(), $left.start.GetColumn(), $left.e, $op.text, $right.e) }//TODO ACA DEBERIA IR EL MODULO
| left=expr op=(ADD|SUB) right=expr { $e = expressions.NewArithmeticOperation($left.start.GetLine(), $left.start.GetColumn(), $left.e, $op.text, $right.e) }
| left=expr op=(MAY_IG|MAYOR|MEN_IG|MENOR) right=expr { $e = expressions.NewRelationalOperation($left.start.GetLine(), $left.start.GetColumn(), $left.e, $op.text, $right.e) }
| left=expr op=(IG_IG|DIF) right=expr { $e = expressions.NewRelationalOperation($left.start.GetLine(), $left.start.GetColumn(), $left.e, $op.text, $right.e) }
| left=expr op=AND right=expr { $e = expressions.NewBooleanOperation($left.start.GetLine(), $left.start.GetColumn(), $left.e, $op.text, $right.e) }
| left=expr op=OR right=expr { $e = expressions.NewBooleanOperation($left.start.GetLine(), $left.start.GetColumn(), $left.e, $op.text, $right.e) }

| ID
| NUMBER
    {
        if (strings.Contains($NUMBER.text,".")){
            num,err := strconv.ParseFloat($NUMBER.text, 64);
            if err!= nil{
                fmt.Println(err)
            }
            $e = expressions.NewPrimitive($NUMBER.line,$NUMBER.pos,num,environment.FLOAT)
        }else{
            num,err := strconv.Atoi($NUMBER.text)
            if err!= nil{
                fmt.Println(err)
            }
            $e = expressions.NewPrimitive($NUMBER.line,$NUMBER.pos,num,environment.INTEGER)
        }
    }
| STRING
    {
        str := $STRING.text
        if (len(str) == 3){
            $e = expressions.NewPrimitive($STRING.line, $STRING.pos, str[1:len(str)-1],environment.CHAR)
        } else{
            $e = expressions.NewPrimitive($STRING.line, $STRING.pos, str[1:len(str)-1],environment.STRING)
        }

    }                        
| RTRUE { $e = expressions.NewPrimitive($RTRUE.line, $RTRUE.pos, true, environment.BOOLEAN) }
| RFALSE { $e = expressions.NewPrimitive($RFALSE.line, $RFALSE.pos, false, environment.BOOLEAN) }
;
