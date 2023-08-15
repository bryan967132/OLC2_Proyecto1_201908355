grammar Parser;

@header {
    import (
        expressions "TSwift/Classes/Expressions"
        instructions "TSwift/Classes/Instructions"
        interfaces "TSwift/Classes/Interfaces"
        utils "TSwift/Classes/Utils"
    )
}

options {
    language = Go;
    tokenVocab = Scanner;
}

init returns[[]interface{} result]:
    insts = instsglobal? EOF {$result = $insts.result} ;

instsglobal returns[[]interface{} result]:
    l = instsglobal i = instglobal {$result = $l.result;; $result = append($result, $i.result)} |
    i = instglobal                 {$result = []interface{}{$i.result}                        } ;

instglobal returns[interface{} result] :
    inst0 = instruction             {$result = $inst0.result} |
    inst1 = declfunc                |
    inst2 = defstruct TK_semicolon? ;

callfunc :
    TK_id TK_lpar listargs? TK_rpar ;

listargs :
    (TK_id TK_colon)? TK_amp? exp TK_comma listargs |
    (TK_id TK_colon)? TK_amp? exp                   ;

decvar :
    RW_var TK_id TK_colon type TK_equ exp  |
    RW_var TK_id TK_colon type TK_question |
    RW_var TK_id TK_equ exp                ;

deccst :
    RW_let TK_id TK_colon type TK_equ exp  |
    RW_let TK_id TK_equ exp                ;

declfunc :
    RW_func TK_id TK_lpar listparams? TK_rpar (TK_prompt type)? env ;

listparams :
    (TK_id | TK_under)? TK_id TK_colon RW_inout? type TK_comma listparams |
    (TK_id | TK_under)? TK_id TK_colon RW_inout? type                     ;

ifstruct :
    RW_if exp env RW_else ifstruct |
    RW_if exp env RW_else env      |
    RW_if exp env                  ;

switchstruct :
    RW_switch exp envs ;

envs :
    TK_lbrc casesdefault TK_rbrc ;

casesdefault :
    cases default |
    cases         |
    default       ;

cases :
    cases case |
    case       ;

case :
    RW_case exp TK_colon instructions ;

default :
    RW_default TK_colon instructions ;

loopfor :
    RW_for TK_id RW_in (exp | range) env ;

range :
    exp TK_dot TK_dot TK_dot exp ;

loopwhile :
    RW_while exp env ;

guard :
    RW_guard exp RW_else env ;

reasign :
    TK_id TK_equ exp ;

addsub :
    TK_id (TK_add | TK_sub) exp ;

decvector :
    RW_var TK_id TK_colon TK_lbrk type TK_rbrk TK_equ defvector ;

defvector :
    TK_lbrk listexp? TK_rbrk |
    simplevec                |
    TK_id                    ;

listexp returns[[]interfaces.Expression result] :
    l = listexp TK_comma e = exp {$result = $l.result;; $result = append($result, $e.result)} |
    e = exp                      {$result = []interfaces.Expression{$e.result}              } ;

simplevec :
    TK_lbrk type TK_rbrk TK_lpar RW_repeating TK_colon exp TK_comma RW_count TK_colon exp TK_rpar ;

funcvector :
    TK_id TK_dot RW_append TK_lpar exp TK_rpar                |
    TK_id TK_dot RW_removeLast TK_lpar TK_rpar                |
    TK_id TK_dot RW_remove TK_lpar RW_at TK_colon exp TK_rpar ;

decmatrix :
    RW_var TK_id (TK_colon typematrix)? TK_equ defmatrix ;

typematrix :
    TK_lbrk typematrix TK_rbrk |
    TK_lbrk type TK_rbrk       ;

defmatrix :
    listvector   |
    simplematrix ;

listvector :
    TK_lbrk listvector2 TK_rbrk ;

listvector2 :
    listvector2 TK_comma listvector |
    listvector                      |
    listexp                         ;

simplematrix :
    typematrix TK_lpar RW_repeating TK_colon simplematrix TK_comma RW_count TK_colon exp TK_rpar |
    typematrix TK_lpar RW_repeating TK_colon exp TK_comma RW_count TK_colon exp TK_rpar          ;


reasignvector :
    TK_id dims TK_equ exp ;

dims :
    TK_lbrk exp TK_rbrk dims |
    TK_lbrk exp TK_rbrk      ;

defstruct :
    RW_struct TK_id TK_lbrc listattribs TK_rbrc ;

listattribs :
    listattribs TK_semicolon? attrib |
    attrib TK_semicolon?             ;

attrib :
    (RW_let | RW_var) TK_id (TK_colon type)? (TK_equ exp)? |
    RW_mutating? declfunc ;

decstruct :
    (RW_let | RW_var) TK_id (TK_colon TK_id)? TK_equ TK_id TK_lpar listdupla? TK_rpar |
    (RW_let | RW_var) TK_id (TK_colon TK_id)? TK_equ TK_id TK_lpar TK_rpar       ;

listdupla :
    TK_id TK_colon exp TK_comma listdupla |
    TK_id TK_colon exp                    ;

useattribs :
    obj useattribs1     |
    obj TK_dot callfunc ;

obj :
    TK_id TK_lbrk exp TK_rbrk |
    TK_id                     ;

useattribs1 :
    TK_dot TK_id useattribs1 |
    TK_dot TK_id             ;

print returns[interfaces.Instruction result]:
    p = RW_print TK_lpar exps = listexp? TK_rpar {$result = instructions.NewPrint($p.line, $p.pos, $exps.result)} ;

env :
    TK_lbrc instructions? TK_rbrc ;

instructions returns[[]interface{} result] :
    l = instructions i = instruction {$result = $l.result;; $result = append($result, $i.result)} |
    i = instruction                  {$result = []interface{}{$i.result}                        } ;

instruction returns[interface{} result] :
    decvar                          TK_semicolon? |
    deccst                          TK_semicolon? |
    ifstruct                                      |
    switchstruct                                  |
    loopfor                                       |
    loopwhile                                     |
    guard                                         |
    (RW_self TK_dot)? reasign       TK_semicolon? |
    (RW_self TK_dot)? addsub        TK_semicolon? |
    decvector                       TK_semicolon? |
    funcvector                      TK_semicolon? |
    (RW_self TK_dot)? reasignvector TK_semicolon? |
    decmatrix                       TK_semicolon? |
    decstruct                       TK_semicolon? |
    (RW_self TK_dot)? useattribs    TK_semicolon? |
    (RW_self TK_dot)? callfunc      TK_semicolon? |
    inst17 = print                           TK_semicolon? {$result = $inst17.result} |
    RW_return exp                   TK_semicolon? |
    RW_return                       TK_semicolon? |
    RW_continue                     TK_semicolon? |
    RW_break                        TK_semicolon? ;

type :
    RW_String    |
    RW_Int       |
    RW_Bool      |
    RW_Character |
    RW_Float     |
    TK_id        ;

exp returns[interfaces.Expression result] :
    // ARITHMETICS
    s = TK_minus e2 = exp                             {$result = expressions.NewArithmetic($s.line, $s.pos, nil, $s.text, $e2.result)                                } |
    e1 = exp s = (TK_mult | TK_div | TK_mod) e2 = exp {$result = expressions.NewArithmetic($e1.result.LineN(), $e1.result.ColumnN(), $e1.result, $s.text, $e2.result)} |
    e1 = exp s = (TK_plus | TK_minus)        e2 = exp {$result = expressions.NewArithmetic($e1.result.LineN(), $e1.result.ColumnN(), $e1.result, $s.text, $e2.result)} |
    // RELATIONALS
    e1 = exp s = (TK_lessequ | TK_moreequ)   e2 = exp {$result = expressions.NewRelational($e1.result.LineN(), $e1.result.ColumnN(), $e1.result, $s.text, $e2.result)} |
    e1 = exp s = (TK_less    | TK_more)      e2 = exp {$result = expressions.NewRelational($e1.result.LineN(), $e1.result.ColumnN(), $e1.result, $s.text, $e2.result)} |
    e1 = exp s = (TK_equequ  | TK_notequ)    e2 = exp {$result = expressions.NewRelational($e1.result.LineN(), $e1.result.ColumnN(), $e1.result, $s.text, $e2.result)} |
    // LOGICS
    s = TK_not e2 = exp           {$result = expressions.NewLogic($s.line, $s.pos, nil, $s.text, $e2.result)                                } |
    e1 = exp s = TK_and e2 = exp  {$result = expressions.NewLogic($e1.result.LineN(), $e1.result.ColumnN(), $e1.result, $s.text, $e2.result)} |
    e1 = exp s = TK_or  e2 = exp  {$result = expressions.NewLogic($e1.result.LineN(), $e1.result.ColumnN(), $e1.result, $s.text, $e2.result)} |
    // CAST
    RW_Int    TK_lpar exp TK_rpar  |
    RW_Float  TK_lpar exp TK_rpar  |
    RW_String TK_lpar exp TK_rpar  |
    // Vector
    TK_id dims                     |
    TK_id TK_dot RW_isEmpty        |
    TK_id TK_dot RW_count          |
    // ATTRIBUTES STRUCT
    (RW_self TK_dot)? useattribs   |
    // CALL FUNCTION
    (RW_self TK_dot)? callfunc     |
    // ACCESS ID
    (RW_self TK_dot)? id = TK_id   |
    // NIL
    n = RW_nil                     {$result = expressions.NewPrimitive($p.line, $p.pos, $p.text, utils.NIL)    } |
    // PRIMITIVES
    p = TK_string                  {$result = expressions.NewPrimitive($p.line, $p.pos, $p.text[1 : len($p.text) - 1], utils.STRING) } |
    p = TK_char                    {$result = expressions.NewPrimitive($p.line, $p.pos, $p.text[1 : len($p.text) - 1], utils.CHAR)   } |
    p = TK_int                     {$result = expressions.NewPrimitive($p.line, $p.pos, $p.text, utils.INT)    } |
    p = TK_float                   {$result = expressions.NewPrimitive($p.line, $p.pos, $p.text, utils.FLOAT)  } |
    p = RW_true                    {$result = expressions.NewPrimitive($p.line, $p.pos, $p.text, utils.BOOLEAN)} |
    p = RW_false                   {$result = expressions.NewPrimitive($p.line, $p.pos, $p.text, utils.BOOLEAN)} |
    // GROUP
    TK_lpar e = exp TK_rpar        {$result = $e.result} ;