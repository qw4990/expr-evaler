%{

package expr_evaler

import (
    "fmt"
)

%}

%union {
    number float64
    boolean bool
    variable string
}

%type <number> num_expr num_expr1 // num_expr1 主要是为了处理运算优先级, 使得 '*' 和 '/' 先被计算

%type <boolean> bool_expr bool_expr1 // bool_expr1 主要是为了处理运算优先级, 使得 '>' 和 '<' 先被计算

%token '+' '-' '*' '/' '(' ')' '>' '<' EQ GE LE AND OR

%token <number> NUM

%token <variable> VAR

%%

top:
    bool_expr
    {
        if $1 {
            setResult(yyrcvr, true)
        } else {
            setResult(yyrcvr, false)
        }
    }

bool_expr:
    '(' bool_expr ')'
    {
        $$ = $2
    }
|   bool_expr AND bool_expr
    {
        $$ = $1 && $3
    }
|   bool_expr OR bool_expr
    {
        $$ = $1 || $3
    }
|   bool_expr1

bool_expr1:
    num_expr '>' num_expr
    {
        $$ = $1 > $3
    }
|   num_expr '<' num_expr
    {
        $$ = $1 < $3
    }
|   num_expr EQ num_expr
    {
        $$ = $1 == $3
    }
|   num_expr GE num_expr
    {
        $$ = $1 >= $3
    }
|   num_expr LE num_expr
    {
        $$ = $1 <= $3
    }

num_expr:
    NUM
|   VAR
    {
        num, ok := getVariable(yyrcvr, $1)
        if !ok {
            panic(fmt.Sprintf("no variable: %v", $1))
        }
        $$ = num
    }
|   '+' num_expr
    {
        $$ = $2
    }
|   '-' num_expr
    {
        $$ = -$2
    }
|   '(' num_expr ')'
    {
        $$ = $2
    }
|   num_expr '+' num_expr
    {
        $$ = $1 + $3
    }
|   num_expr '-' num_expr
    {
        $$ = $1 - $3
    }
|   num_expr1

num_expr1:
    num_expr '*' num_expr
    {
        $$ = $1 * $3
    }
|   num_expr '/' num_expr
    {
        $$ = $1 / $3
    }

%%