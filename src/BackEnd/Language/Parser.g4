grammar Parser;

tokens {
    RW_Int, RW_Float, RW_String, RW_Bool, RW_Character, RW_var, RW_let,
    RW_if, RW_else, RW_for, RW_while, RW_switch, RW_case, RW_default,
    RW_break, RW_continue, RW_return,
    RW_true, RW_false, RW_nil, RW_func, RW_inout, RW_print, TK_prompt, TK_under,
    TK_string, TK_char, TK_int, TK_float, TK_id, TK_add, TK_sub,
    TK_plus, TK_minus, TK_mult, TK_div, TK_pow, TK_mod,
    TK_equequ, TK_notequ, TK_lessequ, TK_moreequ, TK_equ, TK_less, TK_more,
    TK_and, TK_or, TK_not,
    TK_lpar, TK_rpar, TK_lbrc, TK_rbrc, TK_lbrk, TK_rbrk, TK_comma, TK_colon, TK_semicolon, TK_question
}

init :
    instructions EOF |
    EOF              ;

instructions :
    instruction instructions |
    instruction              ;

instruction :
    decvar   |
    deccst   |
    declfunc |
    ifstruct ;

decvar :
    RW_var TK_id TK_colon type TK_equ exp  |
    RW_var TK_id TK_equ exp                |
    RW_var TK_id TK_colon type TK_question ;

deccst :
    RW_let TK_id TK_colon type TK_equ exp  |
    RW_let TK_id TK_colon type TK_question ;

declfunc :
    RW_func TK_id TK_lpar params? TK_rpar (TK_prompt type)? env ;

params :
    TK_comma (TK_id | TK_under) ? TK_id TK_colon RW_inout? type params |
    (TK_id | TK_under) ? TK_id TK_colon RW_inout? type                 ;

ifstruct :
    RW_if exp env RW_else ifstruct |
    RW_if exp env RW_else env      |
    RW_if exp env                  ;

env :
    TK_lbrc instructions? TK_rbrc ;

type :
    RW_String    |
    RW_Int       |
    RW_Bool      |
    RW_Character |
    RW_Float     ;

exp :
    // ARITHMETICS
    <assoc = right> s = TK_minus e2 = exp            |
    <assoc = left>  e1 = exp s = TK_mult    e2 = exp |
    <assoc = left>  e1 = exp s = TK_div     e2 = exp |
    <assoc = left>  e1 = exp s = TK_plus    e2 = exp |
    <assoc = left>  e1 = exp s = TK_minus   e2 = exp |
    <assoc = left>  e1 = exp s = TK_pow     e2 = exp |
    <assoc = left>  e1 = exp s = TK_mod     e2 = exp |
    // RELATIONALS
    <assoc = left>  e1 = exp s = TK_equequ  e2 = exp |
    <assoc = left>  e1 = exp s = TK_notequ  e2 = exp |
    <assoc = left>  e1 = exp s = TK_lessequ e2 = exp |
    <assoc = left>  e1 = exp s = TK_moreequ e2 = exp |
    <assoc = left>  e1 = exp s = TK_less    e2 = exp |
    <assoc = left>  e1 = exp s = TK_more    e2 = exp |
    // LOGICS
    <assoc = right> s = TK_not e2 = exp              |
    <assoc = left>  e1 = exp s = TK_and     e2 = exp |
    <assoc = left>  e1 = exp s = TK_or      e2 = exp |
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
    TK_lpar e = exp TK_rpar ;