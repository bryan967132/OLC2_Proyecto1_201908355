grammar Parser;

tokens {
    RW_Int, RW_Float, RW_String, RW_Bool, RW_Character, RW_var, RW_let,
    RW_if, RW_else, RW_for, RW_while, RW_guard, RW_switch, RW_case, RW_default,
    RW_break, RW_continue, RW_return,
    RW_true, RW_false, RW_nil, RW_func, RW_inout, RW_in,
    RW_append, RW_removeLast, RW_remove, RW_isEmpty, RW_count, RW_repeating,
    RW_print, TK_prompt, TK_under,
    TK_char, TK_string, TK_int, TK_float, TK_id, TK_add, TK_sub,
    TK_plus, TK_minus, TK_mult, TK_div, TK_mod,
    TK_equequ, TK_notequ, TK_lessequ, TK_moreequ, TK_equ, TK_less, TK_more,
    TK_and, TK_or, TK_not,
    TK_lpar, TK_rpar, TK_lbrc, TK_rbrc, TK_lbrk, TK_rbrk,
    TK_dot, TK_comma, TK_colon, TK_semicolon, TK_question, TK_amp
}

init :
    instsglobal? EOF ;

instsglobal :
    instglobal instsglobal |
    instglobal             ;

instglobal :
    instruction |
    declfunc    ;

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
    case cases |
    case       ;

case :
    RW_case exp TK_colon instructions ;

default :
    RW_default TK_colon instructions ;

loopfor :
    RW_for TK_id RW_in (exp | range) env ;

range :
    TK_int TK_dot TK_dot TK_dot TK_int ;

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

listexp :
    exp TK_comma listexp |
    exp                  ;

simplevec :
    TK_lbrk type TK_rbrk TK_lpar RW_repeating TK_colon exp TK_comma RW_count TK_colon exp TK_rpar ;

funcvector :
    TK_id TK_dot RW_append TK_lpar exp TK_rpar |
    TK_id TK_dot RW_removeLast TK_lpar TK_rpar |
    TK_id TK_dot RW_remove TK_lpar exp TK_rpar ;

print :
    RW_print TK_lpar exp? TK_rpar ;

env :
    TK_lbrc instructions? TK_rbrc ;

instructions :
    instruction instructions |
    instruction              ;

instruction :
    decvar        |
    deccst        |
    ifstruct      |
    switchstruct  |
    loopfor       |
    loopwhile     |
    guard         |
    reasign       |
    addsub        |
    decvector     |
    funcvector    |
    callfunc      |
    print         |
    RW_return exp |
    RW_return     |
    RW_continue   |
    RW_break      ;

type :
    RW_String    |
    RW_Int       |
    RW_Bool      |
    RW_Character |
    RW_Float     ;

exp :
    // ARITHMETICS
    s = TK_minus e2 = exp                             |
    e1 = exp s = (TK_mult | TK_div | TK_mod) e2 = exp |
    e1 = exp s = (TK_plus | TK_minus)        e2 = exp |
    // RELATIONALS
    e1 = exp s = (TK_lessequ | TK_moreequ)   e2 = exp |
    e1 = exp s = (TK_less    | TK_more)      e2 = exp |
    e1 = exp s = (TK_equequ  | TK_notequ)    e2 = exp |
    // LOGICS
    s = TK_not e2 = exp          |
    e1 = exp s = TK_and e2 = exp |
    e1 = exp s = TK_or  e2 = exp |
    // CAST
    RW_Int    TK_lpar exp TK_rpar |
    RW_Float  TK_lpar exp TK_rpar |
    RW_String TK_lpar exp TK_rpar |
    // Vector
    TK_id TK_lbrk exp TK_rbrk     |
    TK_id TK_dot RW_isEmpty       |
    TK_id TK_dot RW_count         |
    // CALL FUNCTION
    callfunc                |
    // NIL
    n = RW_nil              |
    // PRIMITIVES
    id = TK_id              |
    p = TK_string           |
    p = TK_char             |
    p = TK_int              |
    p = TK_float            |
    p = RW_true             |
    p = RW_false            |
    // GROUP
    TK_lpar e = exp TK_rpar ;