package expr_evaler_test

import (
	"fmt"
	"testing"

	"github.com/qw4990/expr-evaler"
)

type testCase struct {
	expr   string
	result bool
	vtable expr_evaler.VarTable
	ifErr  bool
}

func runCases(t *testing.T, cases []*testCase) {
	for _, cas := range cases {
		result, err := expr_evaler.Eval(cas.expr, cas.vtable)

		if cas.ifErr {
			fmt.Println(cas.expr, ": ", err)
			if err == nil {
				t.Fatal(cas.expr)
			}
		} else {
			if err != nil {
				t.Fatal(cas.expr)
			}
			if result != cas.result {
				t.Fatal(cas.expr)
			}
		}
	}
}

// TestSimpleExprs .
func TestSimpleExprs(t *testing.T) {
	cases := []*testCase{
		{"2 > 1", true, nil, false},
		{"2 > 3", false, nil, false},
		{"2 + 3 > 4", true, nil, false},
		{" 2 * 2*2*2 > 17", false, nil, false},
		{"3 / 2 > 1", true, nil, false},
		{"3 / 2 > 1.4444", true, nil, false},
		{"3 / 2 > 1.55555", false, nil, false},
		{"1.5 == 3/ 2", true, nil, false},
		{"1.5 >= 3/2", true, nil, false},
		{"1.5 <= 3 /2", true, nil, false},
		{"3 > 2 or 2 > 3", true, nil, false},
		{"3 > 2 and 2 > 3", false, nil, false},
		{"2+3*5 < (2+3)*5", true, nil, false},
		{"3*3*3 == 27", true, nil, false},
		{"333*3 > 0 and 123 * 0 == 0 and 2 >= 1", true, nil, false},
		{"-23 * -10 == 230", true, nil, false},
		{"-2*-2*-2*-2*-2 == -32", true, nil, false},
	}

	runCases(t, cases)
}

func TestWithVariables(t *testing.T) {
	cases := []*testCase{
		{"a > b", true, expr_evaler.VarTable{
			"a": 2.2222,
			"b": -23,
		}, false},
		{" a > b", false, expr_evaler.VarTable{
			"a": 2.2222,
			"b": 034,
		}, false},
		{" a + 34 > b", true, expr_evaler.VarTable{
			"a": 0.2222,
			"b": 034,
		}, false},
	}

	runCases(t, cases)
}

func TestError(t *testing.T) {
	cases := []*testCase{
		{"a >< b", true, expr_evaler.VarTable{
			"a": 2.2222,
			"b": -23,
		}, true}, // syntax error
		{"a < b", true, expr_evaler.VarTable{
			"a": 2.2222,
		}, true}, // no variable
		{"2+2 ", true, nil, true}, // syntax error
	}

	runCases(t, cases)
}
