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
: ins+=instruction*
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
| vecdec PTOCOMA? {$inst = $vecdec.newvecdec}
| appendvec PTOCOMA? {$inst = $appendvec.newappendvec}
| removelastvec PTOCOMA? {$inst = $removelastvec.newremovelastvec}
| removeatvec PTOCOMA?  {$inst = $removeatvec.newremoveat}
| asignation PTOCOMA? {$inst = $asignation.newasignation}
| unarysum PTOCOMA?  {$inst = $unarysum.newunarysum}
| unarysub PTOCOMA?  {$inst = $unarysub.newunarysub}
| breakstatement PTOCOMA? {$inst = $breakstatement.newbreak}
| continuestatement PTOCOMA? {$inst = $continuestatement.newcontinue}
| ifstmt {$inst = $ifstmt.newif}
| while_statement {$inst = $while_statement.newwhile}
| switchstatement {$inst = $switchstatement.newswitch}
;
//--------------------------
removeatvec returns [interfaces.Instruction newremoveat]
: ID PTO RREMOVEAT PARIZQ RRAT DOSPTOS expr PARDER {$newremoveat = instructions.NewRemoveAtVector($ID.line, $ID.pos,$ID.text ,$expr.e)}
  ;

appendvec returns [interfaces.Instruction newappendvec]
: ID PTO RAPPEND PARIZQ expr PARDER {$newappendvec = instructions.NewAppendVector($ID.line, $ID.pos,$ID.text ,$expr.e)}

;

removelastvec returns [interfaces.Instruction newremovelastvec]
: ID PTO RREMOVELAST PARIZQ PARDER {$newremovelastvec = instructions.NewRemoveLastVector($ID.line, $ID.pos,$ID.text)}

;


vecdec returns [interfaces.Instruction newvecdec]
: RVAR ID DOSPTOS OBRA typpe=(RINT|RFLOAT|RBOOL|RSTRING|RCHARACTER) CBRA IG OBRA typpe2=(RINT|RFLOAT|RBOOL|RSTRING|RCHARACTER) CBRA OBRA CBRA {$newvecdec = instructions.NewVecDec($RVAR.line, $RVAR.pos, $ID.text, $typpe.text, $typpe2.text, nil )}
| RVAR firstid=ID DOSPTOS  OBRA typpe=(RINT|RFLOAT|RBOOL|RSTRING|RCHARACTER) CBRA IG expr {$newvecdec = instructions.NewVecDec($RVAR.line, $RVAR.pos, $firstid.text, $typpe.text, nil,$expr.e )}
;

breakstatement returns [interfaces.Instruction newbreak]
: RBREAK    {$newbreak = instructions.NewBreak($RBREAK.line, $RBREAK.pos)}

;

continuestatement returns [interfaces.Instruction newcontinue]
: RCONTINUE    {$newcontinue = instructions.NewContinue($RCONTINUE.line, $RCONTINUE.pos)}

;

switchstatement returns [interfaces.Instruction newswitch]
: RSWITCH  ex=expr  LLAVEIZQ caselist LLAVEDER   {$newswitch = instructions.NewSwitch($RSWITCH.line, $RSWITCH.pos, $ex.e, $caselist.newcaselist, nil)}
 |RSWITCH  ex=expr   LLAVEIZQ caselist defaultstatement LLAVEDER   {$newswitch = instructions.NewSwitch($RSWITCH.line, $RSWITCH.pos, $ex.e, $caselist.newcaselist, $defaultstatement.newdefault)}
;


caselist returns [[]interface{} newcaselist]
@init {
    $newcaselist = []interface{}{}
}
    :case  caselist  { $newcaselist = append($newcaselist, $case.newcase)
                                       for _, arg := range $caselist.newcaselist {
                                           $newcaselist = append($newcaselist, arg)
                                       }
                                 }
    |case { $newcaselist = append( $newcaselist , $case.newcase) }
    | {}
    ;

case returns [interfaces.Instruction newcase]
: RCASE ex=expr DOSPTOS  b=block   {$newcase = instructions.NewCase($RCASE.line, $RCASE.pos, $ex.e, $b.blk)}
;

defaultstatement returns [[] interface{} newdefault]
: RDEFAULT DOSPTOS b=block  LLAVEDER             {$newdefault =  $b.blk}
;

//-----------------------------

ifstmt returns [interfaces.Instruction newif]
: RIF  ex=expr  LLAVEIZQ b=block LLAVEDER eliflist  {$newif = instructions.NewIf($RIF.line, $RIF.pos, $ex.e, $b.blk, $eliflist.neweliflist, nil)}
 |RIF  ex=expr  LLAVEIZQ b=block LLAVEDER eliflist elsestament {$newif = instructions.NewIf($RIF.line, $RIF.pos, $ex.e, $b.blk, $eliflist.neweliflist, $elsestament.newelse)}
;


eliflist returns [[]interface{} neweliflist]
@init {
    $neweliflist = []interface{}{}
}
    :elif  eliflist  { $neweliflist = append($neweliflist, $elif.newelif)
                                       for _, arg := range $eliflist.neweliflist {
                                           $neweliflist = append($neweliflist, arg)
                                       }
                                 }
    |elif { $neweliflist = append( $neweliflist , $elif.newelif) }
    | {}
    ;

elif returns [interfaces.Instruction newelif]
: RELSE RIF ex=expr LLAVEIZQ  b=block  LLAVEDER {$newelif = instructions.NewIf($RELSE.line, $RELSE.pos, $ex.e, $b.blk, nil, nil)}
;

elsestament returns [[] interface{} newelse]
: RELSE LLAVEIZQ b=block  LLAVEDER             {$newelse =  $b.blk}
;
// STATEMENTS----------------------------------------------------------------------------------------------
printstmt returns [interfaces.Instruction prnt]
: RPRINT PARIZQ arguments PARDER { $prnt = instructions.NewPrint($RPRINT.line,$RPRINT.pos,$arguments.args)}
;

while_statement returns [interfaces.Instruction newwhile]
: RWHILE expr LLAVEIZQ b=block LLAVEDER { $newwhile = instructions.NewWhile($RWHILE.line,$RWHILE.pos,$expr.e, $b.blk)}
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

isemptyvec returns [interfaces.Expression newisemptyvec]
: ID PTO RISEMPTY   {$newisemptyvec = expressions.NewIsEmptyVector($ID.line, $ID.pos,$ID.text)}
  ;

countvec returns [interfaces.Expression newcountvec]
: ID PTO RCOUNT   {$newcountvec = expressions.NewCountVector($ID.line, $ID.pos,$ID.text)}
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
| isemptyvec {$e = $isemptyvec.newisemptyvec}
| countvec {$e = $countvec.newcountvec}
| OBRA arguments CBRA {$e = expressions.NewVector($OBRA.line, $OBRA.pos, $arguments.args) }
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
