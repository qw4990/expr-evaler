# Expr Evaler(EE)
EE is used to evaluate a bool expr, and return its result(true/false);

In EE's exprs, you can use:

+ <b>predefined</b> variables, but all variables' type must be float64;
+ numbers(float64), '+', '-', '*', '/', '(', ')' to do calculations;
+ '>', '<', '==', '>=', '<=', 'and', 'or' to do logical operations;

# API & Examples
## API
```
// Eval evaluate this expr with this varTable;
Eval(expr string, varTable VarTable) (result bool, err error)

// VarTable variable table, all variables' type must be float64
type VarTable map[string]float64
```

The syntax of EE is very simple, so you can see some examples below to catch it;

And completed syntax of EE is defined in ee.y;

## Simple Expr
```
Eval("2 * 2*2*2 > 17", nil)

Eval("2+3*5 < (2+3)*5", nil)

Eval("3*3*3 == 27", nil)

Eval("333*3 > 0 and 123 * 0 == 0 and 2 >= 1", nil)

Eval("-23 * -10 == 230", nil)

Eval("-2*-2*-2*-2*-2 == -32", nil)
```

## With Variables
```
Eval("a > b", VarTable{
	"a": 2.2222,
	"b": 034,
})

Eval("a + 34 > b", VarTable{
    "a": 0.2222,
	"b": 034,
})
```

## More Examples
Please see ee_test.go :)

# Dependencies
EE based on goyacc and nex(https://github.com/blynn/nex);