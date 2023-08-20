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
| vardec PTOCOMA?  { $inst = $vardec.newdec}
| constdec PTOCOMA? {$inst = $constdec.newconst}
| asignation PTOCOMA? {$inst = $asignation.newasignation}
| unarysum PTOCOMA?  {$inst = $unarysum.newunarysum}
| unarysub PTOCOMA?  {$inst = $unarysub.newunarysub}
| ifstmt {$inst = $ifstmt.newif}
;
//--------------------------
ifstmt returns [interfaces.Instruction newif]
: RIF  ex=expr  LLAVEIZQ b=block LLAVEDER elseifstatement
    {$newif = instructions.NewIf($RIF.line, $RIF.pos, $ex.e, $b.blk, $elseifstatement.newelse)}
;

elseifstatement returns [[] interface{} newelse]
: RELSE ifstmt                {$newelse = append($newelse, $ifstmt.newif)}
| RELSE LLAVEIZQ b=block  LLAVEDER             {$newelse = append($newelse, $b.blk)}
|                             {}
;

// STATEMENTS----------------------------------------------------------------------------------------------
printstmt returns [interfaces.Instruction prnt]
: RPRINT PARIZQ arguments PARDER { $prnt = instructions.NewPrint($RPRINT.line,$RPRINT.pos,$arguments.args)}
;

vardec returns [interfaces.Instruction newdec]
: RVAR ID DOSPTOS typpe=(RINT|RFLOAT|RBOOL|RSTRING|RCHARACTER) IG ex=expr { $newdec = instructions.NewVarDec($RVAR.line,$RVAR.pos,$ID.text,$typpe.text, $ex.e)}
| RVAR ID IG ex=expr                                                      { $newdec = instructions.NewVarDec($RVAR.line,$RVAR.pos,$ID.text, nil, $ex.e)}
| RVAR ID DOSPTOS typpe=(RINT|RFLOAT|RBOOL|RSTRING|RCHARACTER) QM         { $newdec = instructions.NewVarDec($RVAR.line,$RVAR.pos,$ID.text, $typpe.text, nil)}
;

constdec returns [interfaces.Instruction newconst]
: RLET ID DOSPTOS typpe=(RINT|RFLOAT|RBOOL|RSTRING|RCHARACTER) IG ex=expr { $newconst = instructions.NewConstDec($RLET.line,$RLET.pos,$ID.text,$typpe.text, $ex.e)}
| RLET ID IG ex=expr                                                      { $newconst = instructions.NewConstDec($RLET.line,$RLET.pos,$ID.text, nil, $ex.e)}
;

asignation returns [interfaces.Instruction newasignation]
: ID IG ex=expr { $newasignation = instructions.NewAsignation($ID.line,$ID.pos,$ID.text, $ex.e)}
;




unarysum returns [interfaces.Instruction newunarysum]
:ID UNARYPLUS ex=expr { $newunarysum = instructions.NewUnarySum($ID.line,$ID.pos,$ID.text, "+=", $ex.e)}
;

unarysub returns [interfaces.Instruction newunarysub]
:ID UNARYMINUS ex=expr { $newunarysub = instructions.NewUnarySum($ID.line,$ID.pos,$ID.text, "-=", $ex.e)}
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
| ID                        { $e = expressions.NewVariableAccess($ID.text) }
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
        var chari = str[1:len(str)-1]
        chari = strings.ReplaceAll(chari, "\\n","\n")
        chari = strings.ReplaceAll(chari, "\\r","\r")
        chari = strings.ReplaceAll(chari, "\\t","\t")
        chari = strings.ReplaceAll(chari, "\\\\","\\")
        chari = strings.ReplaceAll(chari, "\\\"","\"")
        if (len(str) == 3){

            $e = expressions.NewPrimitive($STRING.line, $STRING.pos, chari,environment.CHAR)
        } else{

            $e = expressions.NewPrimitive($STRING.line, $STRING.pos, chari,environment.STRING)
        }

    }                        
| RTRUE { $e = expressions.NewPrimitive($RTRUE.line, $RTRUE.pos, true, environment.BOOLEAN) }
| RFALSE { $e = expressions.NewPrimitive($RFALSE.line, $RFALSE.pos, false, environment.BOOLEAN) }
| RNIL  { $e = expressions.NewPrimitive($RFALSE.line, $RFALSE.pos, nil, environment.NULL) }
;
