//line ee.y:2
package expr_evaler

import __yyfmt__ "fmt"

//line ee.y:3
import (
	"fmt"
)

//line ee.y:11
type yySymType struct {
	yys      int
	number   float64
	boolean  bool
	variable string
}

const EQ = 57346
const GE = 57347
const LE = 57348
const AND = 57349
const OR = 57350
const NUM = 57351
const VAR = 57352

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"'('",
	"')'",
	"'>'",
	"'<'",
	"EQ",
	"GE",
	"LE",
	"AND",
	"OR",
	"NUM",
	"VAR",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line ee.y:118

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 21
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 79

var yyAct = [...]int{

	5, 8, 9, 1, 14, 25, 11, 12, 4, 24,
	26, 10, 0, 0, 6, 7, 31, 32, 33, 34,
	35, 36, 37, 38, 39, 0, 40, 20, 21, 22,
	23, 0, 30, 15, 16, 17, 18, 19, 8, 9,
	0, 0, 3, 20, 21, 22, 23, 20, 21, 22,
	23, 6, 7, 15, 16, 17, 18, 19, 29, 2,
	0, 0, 0, 13, 11, 12, 0, 0, 0, 0,
	0, 27, 28, 20, 21, 22, 23, 0, 30,
}
var yyPact = [...]int{

	34, -1000, -9, 34, -1000, 43, -1000, -1000, -3, -3,
	-1000, 34, 34, 49, 23, -3, -3, -3, -3, -3,
	-3, -3, -3, -3, 39, -3, 39, -9, -9, -1000,
	-1000, 39, 39, 39, 39, 39, 39, 39, 39, 39,
	69,
}
var yyPgo = [...]int{

	0, 0, 11, 59, 8, 3,
}
var yyR1 = [...]int{

	0, 5, 3, 3, 3, 3, 4, 4, 4, 4,
	4, 1, 1, 1, 1, 1, 1, 1, 1, 2,
	2,
}
var yyR2 = [...]int{

	0, 1, 3, 3, 3, 1, 3, 3, 3, 3,
	3, 1, 1, 2, 2, 3, 3, 3, 1, 3,
	3,
}
var yyChk = [...]int{

	-1000, -5, -3, 8, -4, -1, 17, 18, 4, 5,
	-2, 15, 16, -3, -1, 10, 11, 12, 13, 14,
	4, 5, 6, 7, -1, 8, -1, -3, -3, 9,
	9, -1, -1, -1, -1, -1, -1, -1, -1, -1,
	-1,
}
var yyDef = [...]int{

	0, -2, 1, 0, 5, 0, 11, 12, 0, 0,
	18, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 13, 0, 14, 3, 4, 2,
	15, 6, 7, 8, 9, 10, 16, 17, 19, 20,
	0,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	8, 9, 6, 4, 3, 5, 3, 7, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	11, 3, 10,
}
var yyTok2 = [...]int{

	2, 3, 12, 13, 14, 15, 16, 17, 18,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ee.y:31
		{
			if yyDollar[1].boolean {
				setResult(yyrcvr, true)
			} else {
				setResult(yyrcvr, false)
			}
		}
	case 2:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ee.y:41
		{
			yyVAL.boolean = yyDollar[2].boolean
		}
	case 3:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ee.y:45
		{
			yyVAL.boolean = yyDollar[1].boolean && yyDollar[3].boolean
		}
	case 4:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ee.y:49
		{
			yyVAL.boolean = yyDollar[1].boolean || yyDollar[3].boolean
		}
	case 6:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ee.y:56
		{
			yyVAL.boolean = yyDollar[1].number > yyDollar[3].number
		}
	case 7:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ee.y:60
		{
			yyVAL.boolean = yyDollar[1].number < yyDollar[3].number
		}
	case 8:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ee.y:64
		{
			yyVAL.boolean = yyDollar[1].number == yyDollar[3].number
		}
	case 9:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ee.y:68
		{
			yyVAL.boolean = yyDollar[1].number >= yyDollar[3].number
		}
	case 10:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ee.y:72
		{
			yyVAL.boolean = yyDollar[1].number <= yyDollar[3].number
		}
	case 12:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ee.y:79
		{
			num, ok := getVariable(yyrcvr, yyDollar[1].variable)
			if !ok {
				panic(fmt.Sprintf("no variable: %v", yyDollar[1].variable))
			}
			yyVAL.number = num
		}
	case 13:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ee.y:87
		{
			yyVAL.number = yyDollar[2].number
		}
	case 14:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ee.y:91
		{
			yyVAL.number = -yyDollar[2].number
		}
	case 15:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ee.y:95
		{
			yyVAL.number = yyDollar[2].number
		}
	case 16:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ee.y:99
		{
			yyVAL.number = yyDollar[1].number + yyDollar[3].number
		}
	case 17:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ee.y:103
		{
			yyVAL.number = yyDollar[1].number - yyDollar[3].number
		}
	case 19:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ee.y:110
		{
			yyVAL.number = yyDollar[1].number * yyDollar[3].number
		}
	case 20:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ee.y:114
		{
			yyVAL.number = yyDollar[1].number / yyDollar[3].number
		}
	}
	goto yystack /* stack new state and value */
}
