package expr_evaler

import (
	"fmt"
	"sync"
	"unsafe"
)

import "bytes"

// VarTable variable table, all variables' type must be float64
type VarTable map[string]float64

var (
	varTableMap  = make(map[uintptr]VarTable)
	varTableLock sync.RWMutex

	resultMap  = make(map[uintptr]bool)
	resultLock sync.RWMutex
)

func remove(p *yyParserImpl) {
	ptr := (uintptr)(unsafe.Pointer(p))
	resultLock.Lock()
	defer resultLock.Unlock()
	varTableLock.Lock()
	defer varTableLock.Unlock()
	delete(resultMap, ptr)
	delete(varTableMap, ptr)
}

func setResult(p *yyParserImpl, result bool) {
	ptr := (uintptr)(unsafe.Pointer(p))
	resultLock.Lock()
	defer resultLock.Unlock()
	resultMap[ptr] = result
}

func getResult(p *yyParserImpl) bool {
	ptr := (uintptr)(unsafe.Pointer(p))
	resultLock.RLock()
	defer resultLock.RUnlock()
	result, ok := resultMap[ptr]
	if !ok { // assert false: ok should always be true
		panic("no result")
	}
	return result
}

func getVariable(p *yyParserImpl, vari string) (float64, bool) {
	ptr := (uintptr)(unsafe.Pointer(p))
	varTableLock.RLock()
	defer varTableLock.RUnlock()
	varTable, ok := varTableMap[ptr]
	if !ok { // assert false: ok should always be true
		panic("no variable table")
	}
	num, ok := varTable[vari]
	return num, ok
}

func setVarTable(p *yyParserImpl, varTable VarTable) {
	ptr := (uintptr)(unsafe.Pointer(p))
	varTableLock.Lock()
	defer varTableLock.Unlock()
	varTableMap[ptr] = varTable
}

// Eval evaluate this expr with this varTable;
func Eval(expr string, varTable VarTable) (result bool, err error) {
	lexer := NewLexer(bytes.NewReader([]byte(expr)))
	parser := &yyParserImpl{}
	setVarTable(parser, varTable)
	defer remove(parser)

	defer func() {
		// catch panic
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()
	parser.Parse(lexer)
	return getResult(parser), nil
}
