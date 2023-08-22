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
    insts = instsglobal EOF {$result = $insts.result  } |
    EOF                     {$result = []interface{}{}} ;

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

decvar returns[interfaces.Instruction result] :
    d = RW_var id = TK_id TK_colon t = type TK_equ e = exp {$result = instructions.NewInitID($d.line, $d.pos, true, $id.text, $t.result.Value.(utils.Type), $e.result)} |
    d = RW_var id = TK_id TK_colon t = type TK_question    {$result = instructions.NewInitID($d.line, $d.pos, true, $id.text, $t.result.Value.(utils.Type), nil)      } |
    d = RW_var id = TK_id TK_equ e = exp                   {$result = instructions.NewInitID($d.line, $d.pos, true, $id.text, utils.NIL, $e.result)                   } ;

deccst returns[interfaces.Instruction result] :
    d = RW_let id = TK_id TK_colon t = type TK_equ e = exp {$result = instructions.NewInitID($d.line, $d.pos, false, $id.text, $t.result.Value.(utils.Type), $e.result)} |
    d = RW_let id = TK_id TK_equ e = exp                   {$result = instructions.NewInitID($d.line, $d.pos, false, $id.text, utils.NIL, $e.result)                   } ;

declfunc :
    RW_func TK_id TK_lpar listparams TK_rpar TK_prompt type env |
    RW_func TK_id TK_lpar listparams TK_rpar env                |
    RW_func TK_id TK_lpar TK_rpar TK_prompt type env            |
    RW_func TK_id TK_lpar TK_rpar env                           ;

listparams :
    (TK_id | TK_under)? TK_id TK_colon RW_inout? type TK_comma listparams |
    (TK_id | TK_under)? TK_id TK_colon RW_inout? type                     ;

ifstruct returns[interfaces.Instruction result] :
    r = RW_if cn = exp b1 = env RW_else b2 = ifstruct  {$result = instructions.NewIf($r.line, $r.pos, $cn.result, $b1.result, $b2.result)                                        } |
    r = RW_if cn = exp b1 = env RW_else b3 = env       {$result = instructions.NewIf($r.line, $r.pos, $cn.result, $b1.result, (interface{}($b3.result)).(interfaces.Instruction))} |
    r = RW_if cn = exp b1 = env                        {$result = instructions.NewIf($r.line, $r.pos, $cn.result, $b1.result, nil)                                               } ;

switchstruct returns[interfaces.Instruction result] :
    s = RW_switch e = exp b = envs {$result = instructions.NewSwitch($s.line, $s.pos, $e.result, $b.result[0], $b.result[1].(interfaces.Instruction))} ;

envs returns[[]interface{} result] :
    TK_lbrc cd = casesdefault TK_rbrc {$result = $cd.result             } |
    TK_lbrc TK_rbrc                   {$result = []interface{}{nil, nil}} ;

casesdefault returns[[]interface{} result] :
    c = cases d = default {$result = []interface{}{$c.result, $d.result}} |
    c = cases             {$result = []interface{}{$c.result, nil}      } |
    d = default           {$result = []interface{}{nil,       $d.result}} ;

cases returns[[]interfaces.Instruction result] :
    l = cases c = case {$result = $l.result;; $result = append($result, $c.result)} |
    c = case           {$result = []interfaces.Instruction{$c.result}             } ;

case returns[interfaces.Instruction result] :
    c = RW_case e = exp TK_colon b = instructions {$result = instructions.NewCase($c.line, $c.pos, $e.result, instructions.NewBlock($c.line, $c.pos, $b.result))      } |
    c = RW_case e = exp TK_colon                  {$result = instructions.NewCase($c.line, $c.pos, $e.result, instructions.NewBlock($c.line, $c.pos, []interface{}{}))} ;

default returns[interfaces.Instruction result] :
    d = RW_default TK_colon b = instructions {$result = instructions.NewBlock($d.line, $d.pos, $b.result)      } |
    d = RW_default TK_colon                  {$result = instructions.NewBlock($d.line, $d.pos, []interface{}{})} ;

loopfor returns[interfaces.Instruction result] :
    f = RW_for id = TK_id RW_in r = range b = env {$result = instructions.NewFor($f.line, $f.pos, $id.text, nil, $r.result[0], $r.result[1], $b.result)} |
    f = RW_for id = TK_id RW_in e = exp b = env   {$result = instructions.NewFor($f.line, $f.pos, $id.text, $e.result, nil, nil, $b.result)            } ;

range returns[[]interfaces.Expression result] :
    e1 = exp TK_dot TK_dot TK_dot e2 = exp {$result = []interfaces.Expression{$e1.result, $e2.result}} ;

loopwhile returns[interfaces.Instruction result] :
    w = RW_while e = exp b = env {$result = instructions.NewWhile($w.line, $w.pos, $e.result, $b.result)} ;

guard returns[interfaces.Instruction result] :
    g = RW_guard e = exp RW_else b = env {$result = instructions.NewGuard($g.line, $g.pos, $e.result, $b.result)} ;

reasign returns[interfaces.Instruction result] :
    id = TK_id TK_equ e = exp {$result = instructions.NewAsignID($id.line, $id.pos, $id.text, $e.result)} ;

addsub returns[interfaces.Instruction result] :
    id = TK_id s = (TK_add | TK_sub) e = exp {$result = instructions.NewAddSub($id.line, $id.pos, $id.text, $s.text, $e.result)} ;

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
    p = RW_print TK_lpar exps = listexp TK_rpar {$result = instructions.NewPrint($p.line, $p.pos, $exps.result)} |
    p = RW_print TK_lpar TK_rpar                {$result = instructions.NewPrint($p.line, $p.pos, nil)         } ;

env returns[interfaces.Instruction result] :
    l = TK_lbrc ins = instructions TK_rbrc {$result = instructions.NewBlock($l.line, $l.pos, $ins.result)    } |
    l = TK_lbrc TK_rbrc                    {$result = instructions.NewBlock($l.line, $l.pos, []interface{}{})} ;

instructions returns[[]interface{} result] :
    l = instructions i = instruction {$result = $l.result;; $result = append($result, $i.result)} |
    i = instruction                  {$result = []interface{}{$i.result}                        } ;

instruction returns[interface{} result] :
    inst1 =  decvar                          TK_semicolon? {$result = $inst1.result} |
    inst2 =  deccst                          TK_semicolon? {$result = $inst2.result} |
    inst3 =  ifstruct                                      {$result = $inst3.result} |
    inst4 =  switchstruct                                  {$result = $inst4.result} |
    inst5 =  loopfor                                       {$result = $inst5.result} |
    inst6 =  loopwhile                                     {$result = $inst6.result} |
    inst7 =  guard                                         {$result = $inst7.result} |
    (RW_self TK_dot)? inst8 = reasign        TK_semicolon? {$result = $inst8.result} |
    (RW_self TK_dot)? inst9 = addsub         TK_semicolon? {$result = $inst9.result} |
    decvector                       TK_semicolon? |
    funcvector                      TK_semicolon? |
    (RW_self TK_dot)? reasignvector TK_semicolon? |
    decmatrix                       TK_semicolon? |
    decstruct                       TK_semicolon? |
    (RW_self TK_dot)? useattribs    TK_semicolon? |
    (RW_self TK_dot)? callfunc      TK_semicolon? |
    inst17 = print                           TK_semicolon? {$result = $inst17.result} |
    inst18 = RW_return e = exp               TK_semicolon? {$result = expressions.NewReturn($inst18.line, $inst18.line, $e.result)}|
    inst19 = RW_return                       TK_semicolon? {$result = expressions.NewReturn($inst19.line, $inst19.line, nil)      }|
    inst20 = RW_continue                     TK_semicolon? {$result = instructions.NewContinue($inst20.line, $inst20.line)        }|
    inst21 = RW_break                        TK_semicolon? {$result = instructions.NewBreak($inst21.line, $inst21.line)           };

type returns[utils.AttribsType result] :
    t = RW_String    {$result = *utils.NewAttribsType($t.line, $t.pos, utils.STRING) } |
    t = RW_Int       {$result = *utils.NewAttribsType($t.line, $t.pos, utils.INT)    } |
    t = RW_Bool      {$result = *utils.NewAttribsType($t.line, $t.pos, utils.BOOLEAN)} |
    t = RW_Character {$result = *utils.NewAttribsType($t.line, $t.pos, utils.CHAR)   } |
    t = RW_Float     {$result = *utils.NewAttribsType($t.line, $t.pos, utils.FLOAT)  } |
    t = TK_id        {$result = *utils.NewAttribsType($t.line, $t.pos, $t.text)      } ;

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
    t = type TK_lpar e = exp TK_rpar {$result = expressions.NewCast($t.result.Line, $t.result.Column, $t.result.Value.(utils.Type), $e.result)} |
    // Vector
    TK_id dims                     |
    TK_id TK_dot RW_isEmpty        |
    TK_id TK_dot RW_count          |
    // ATTRIBUTES STRUCT
    (RW_self TK_dot)? useattribs   |
    // CALL FUNCTION
    (RW_self TK_dot)? callfunc     |
    // ACCESS ID
    (RW_self TK_dot)? id = TK_id   {$result = expressions.NewAccessID($id.line, $id.pos, $id.text)             }|
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