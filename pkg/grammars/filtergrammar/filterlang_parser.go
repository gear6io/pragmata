// Code generated from FilterLang.g4 by ANTLR 4.13.2. DO NOT EDIT.

package filtergrammar // FilterLang
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type FilterLang struct {
	*antlr.BaseParser
}

var FilterLangParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func filterlangParserInit() {
	staticData := &FilterLangParserStaticData
	staticData.LiteralNames = []string{
		"", "", "'<='", "'>='", "'!='", "'<'", "'>'", "'='", "'::'", "'.'",
		"','", "'('", "')'",
	}
	staticData.SymbolicNames = []string{
		"", "WS", "OP_LTE", "OP_GTE", "OP_NEQ", "OP_LT", "OP_GT", "OP_EQ", "CAST_OP",
		"DOT", "COMMA", "LPAREN", "RPAREN", "HAS_ANY", "HAS_ALL", "AND", "NOT",
		"OR", "IN", "HAS", "CONTAINS", "ILIKE", "LIKE", "MATCHES", "EXISTS",
		"IS", "NULL", "TRUE", "FALSE", "FLOAT", "INTEGER", "STRING", "IDENT",
	}
	staticData.RuleNames = []string{
		"filterExpr", "filter", "condition", "op", "path", "typeName", "scalarList",
		"scalar",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 32, 157, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 28, 8, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 5, 1, 36, 8, 1, 10, 1, 12, 1, 39, 9, 1, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 122, 8, 2, 1, 3, 1, 3, 1, 4, 1,
		4, 1, 4, 5, 4, 129, 8, 4, 10, 4, 12, 4, 132, 9, 4, 1, 5, 1, 5, 1, 5, 1,
		5, 3, 5, 138, 8, 5, 1, 6, 1, 6, 1, 6, 5, 6, 143, 8, 6, 10, 6, 12, 6, 146,
		9, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 155, 8, 7, 1, 7,
		0, 1, 2, 8, 0, 2, 4, 6, 8, 10, 12, 14, 0, 1, 1, 0, 2, 7, 177, 0, 16, 1,
		0, 0, 0, 2, 27, 1, 0, 0, 0, 4, 121, 1, 0, 0, 0, 6, 123, 1, 0, 0, 0, 8,
		125, 1, 0, 0, 0, 10, 133, 1, 0, 0, 0, 12, 139, 1, 0, 0, 0, 14, 154, 1,
		0, 0, 0, 16, 17, 3, 2, 1, 0, 17, 18, 5, 0, 0, 1, 18, 1, 1, 0, 0, 0, 19,
		20, 6, 1, -1, 0, 20, 21, 5, 16, 0, 0, 21, 28, 3, 2, 1, 3, 22, 23, 5, 11,
		0, 0, 23, 24, 3, 2, 1, 0, 24, 25, 5, 12, 0, 0, 25, 28, 1, 0, 0, 0, 26,
		28, 3, 4, 2, 0, 27, 19, 1, 0, 0, 0, 27, 22, 1, 0, 0, 0, 27, 26, 1, 0, 0,
		0, 28, 37, 1, 0, 0, 0, 29, 30, 10, 5, 0, 0, 30, 31, 5, 17, 0, 0, 31, 36,
		3, 2, 1, 6, 32, 33, 10, 4, 0, 0, 33, 34, 5, 15, 0, 0, 34, 36, 3, 2, 1,
		5, 35, 29, 1, 0, 0, 0, 35, 32, 1, 0, 0, 0, 36, 39, 1, 0, 0, 0, 37, 35,
		1, 0, 0, 0, 37, 38, 1, 0, 0, 0, 38, 3, 1, 0, 0, 0, 39, 37, 1, 0, 0, 0,
		40, 41, 3, 8, 4, 0, 41, 42, 5, 8, 0, 0, 42, 43, 3, 10, 5, 0, 43, 44, 3,
		6, 3, 0, 44, 45, 3, 14, 7, 0, 45, 122, 1, 0, 0, 0, 46, 47, 3, 8, 4, 0,
		47, 48, 3, 6, 3, 0, 48, 49, 3, 14, 7, 0, 49, 122, 1, 0, 0, 0, 50, 51, 3,
		8, 4, 0, 51, 52, 5, 18, 0, 0, 52, 53, 5, 11, 0, 0, 53, 54, 3, 12, 6, 0,
		54, 55, 5, 12, 0, 0, 55, 122, 1, 0, 0, 0, 56, 57, 3, 8, 4, 0, 57, 58, 5,
		16, 0, 0, 58, 59, 5, 18, 0, 0, 59, 60, 5, 11, 0, 0, 60, 61, 3, 12, 6, 0,
		61, 62, 5, 12, 0, 0, 62, 122, 1, 0, 0, 0, 63, 64, 3, 8, 4, 0, 64, 65, 5,
		19, 0, 0, 65, 66, 3, 14, 7, 0, 66, 122, 1, 0, 0, 0, 67, 68, 3, 8, 4, 0,
		68, 69, 5, 13, 0, 0, 69, 70, 5, 11, 0, 0, 70, 71, 3, 12, 6, 0, 71, 72,
		5, 12, 0, 0, 72, 122, 1, 0, 0, 0, 73, 74, 3, 8, 4, 0, 74, 75, 5, 14, 0,
		0, 75, 76, 5, 11, 0, 0, 76, 77, 3, 12, 6, 0, 77, 78, 5, 12, 0, 0, 78, 122,
		1, 0, 0, 0, 79, 80, 3, 8, 4, 0, 80, 81, 5, 20, 0, 0, 81, 82, 3, 14, 7,
		0, 82, 122, 1, 0, 0, 0, 83, 84, 3, 8, 4, 0, 84, 85, 5, 16, 0, 0, 85, 86,
		5, 20, 0, 0, 86, 87, 3, 14, 7, 0, 87, 122, 1, 0, 0, 0, 88, 89, 3, 8, 4,
		0, 89, 90, 5, 22, 0, 0, 90, 91, 3, 14, 7, 0, 91, 122, 1, 0, 0, 0, 92, 93,
		3, 8, 4, 0, 93, 94, 5, 21, 0, 0, 94, 95, 3, 14, 7, 0, 95, 122, 1, 0, 0,
		0, 96, 97, 3, 8, 4, 0, 97, 98, 5, 16, 0, 0, 98, 99, 5, 22, 0, 0, 99, 100,
		3, 14, 7, 0, 100, 122, 1, 0, 0, 0, 101, 102, 3, 8, 4, 0, 102, 103, 5, 23,
		0, 0, 103, 104, 3, 14, 7, 0, 104, 122, 1, 0, 0, 0, 105, 106, 3, 8, 4, 0,
		106, 107, 5, 24, 0, 0, 107, 122, 1, 0, 0, 0, 108, 109, 3, 8, 4, 0, 109,
		110, 5, 16, 0, 0, 110, 111, 5, 24, 0, 0, 111, 122, 1, 0, 0, 0, 112, 113,
		3, 8, 4, 0, 113, 114, 5, 25, 0, 0, 114, 115, 5, 26, 0, 0, 115, 122, 1,
		0, 0, 0, 116, 117, 3, 8, 4, 0, 117, 118, 5, 25, 0, 0, 118, 119, 5, 16,
		0, 0, 119, 120, 5, 26, 0, 0, 120, 122, 1, 0, 0, 0, 121, 40, 1, 0, 0, 0,
		121, 46, 1, 0, 0, 0, 121, 50, 1, 0, 0, 0, 121, 56, 1, 0, 0, 0, 121, 63,
		1, 0, 0, 0, 121, 67, 1, 0, 0, 0, 121, 73, 1, 0, 0, 0, 121, 79, 1, 0, 0,
		0, 121, 83, 1, 0, 0, 0, 121, 88, 1, 0, 0, 0, 121, 92, 1, 0, 0, 0, 121,
		96, 1, 0, 0, 0, 121, 101, 1, 0, 0, 0, 121, 105, 1, 0, 0, 0, 121, 108, 1,
		0, 0, 0, 121, 112, 1, 0, 0, 0, 121, 116, 1, 0, 0, 0, 122, 5, 1, 0, 0, 0,
		123, 124, 7, 0, 0, 0, 124, 7, 1, 0, 0, 0, 125, 130, 5, 32, 0, 0, 126, 127,
		5, 9, 0, 0, 127, 129, 5, 32, 0, 0, 128, 126, 1, 0, 0, 0, 129, 132, 1, 0,
		0, 0, 130, 128, 1, 0, 0, 0, 130, 131, 1, 0, 0, 0, 131, 9, 1, 0, 0, 0, 132,
		130, 1, 0, 0, 0, 133, 137, 5, 32, 0, 0, 134, 135, 5, 11, 0, 0, 135, 136,
		5, 32, 0, 0, 136, 138, 5, 12, 0, 0, 137, 134, 1, 0, 0, 0, 137, 138, 1,
		0, 0, 0, 138, 11, 1, 0, 0, 0, 139, 144, 3, 14, 7, 0, 140, 141, 5, 10, 0,
		0, 141, 143, 3, 14, 7, 0, 142, 140, 1, 0, 0, 0, 143, 146, 1, 0, 0, 0, 144,
		142, 1, 0, 0, 0, 144, 145, 1, 0, 0, 0, 145, 13, 1, 0, 0, 0, 146, 144, 1,
		0, 0, 0, 147, 155, 5, 31, 0, 0, 148, 155, 5, 29, 0, 0, 149, 155, 5, 30,
		0, 0, 150, 155, 5, 27, 0, 0, 151, 155, 5, 28, 0, 0, 152, 155, 5, 26, 0,
		0, 153, 155, 5, 32, 0, 0, 154, 147, 1, 0, 0, 0, 154, 148, 1, 0, 0, 0, 154,
		149, 1, 0, 0, 0, 154, 150, 1, 0, 0, 0, 154, 151, 1, 0, 0, 0, 154, 152,
		1, 0, 0, 0, 154, 153, 1, 0, 0, 0, 155, 15, 1, 0, 0, 0, 8, 27, 35, 37, 121,
		130, 137, 144, 154,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// FilterLangInit initializes any static state used to implement FilterLang. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewFilterLang(). You can call this function if you wish to initialize the static state ahead
// of time.
func FilterLangInit() {
	staticData := &FilterLangParserStaticData
	staticData.once.Do(filterlangParserInit)
}

// NewFilterLang produces a new parser instance for the optional input antlr.TokenStream.
func NewFilterLang(input antlr.TokenStream) *FilterLang {
	FilterLangInit()
	this := new(FilterLang)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &FilterLangParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "FilterLang.g4"

	return this
}

// FilterLang tokens.
const (
	FilterLangEOF      = antlr.TokenEOF
	FilterLangWS       = 1
	FilterLangOP_LTE   = 2
	FilterLangOP_GTE   = 3
	FilterLangOP_NEQ   = 4
	FilterLangOP_LT    = 5
	FilterLangOP_GT    = 6
	FilterLangOP_EQ    = 7
	FilterLangCAST_OP  = 8
	FilterLangDOT      = 9
	FilterLangCOMMA    = 10
	FilterLangLPAREN   = 11
	FilterLangRPAREN   = 12
	FilterLangHAS_ANY  = 13
	FilterLangHAS_ALL  = 14
	FilterLangAND      = 15
	FilterLangNOT      = 16
	FilterLangOR       = 17
	FilterLangIN       = 18
	FilterLangHAS      = 19
	FilterLangCONTAINS = 20
	FilterLangILIKE    = 21
	FilterLangLIKE     = 22
	FilterLangMATCHES  = 23
	FilterLangEXISTS   = 24
	FilterLangIS       = 25
	FilterLangNULL     = 26
	FilterLangTRUE     = 27
	FilterLangFALSE    = 28
	FilterLangFLOAT    = 29
	FilterLangINTEGER  = 30
	FilterLangSTRING   = 31
	FilterLangIDENT    = 32
)

// FilterLang rules.
const (
	FilterLangRULE_filterExpr = 0
	FilterLangRULE_filter     = 1
	FilterLangRULE_condition  = 2
	FilterLangRULE_op         = 3
	FilterLangRULE_path       = 4
	FilterLangRULE_typeName   = 5
	FilterLangRULE_scalarList = 6
	FilterLangRULE_scalar     = 7
)

// IFilterExprContext is an interface to support dynamic dispatch.
type IFilterExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Filter() IFilterContext
	EOF() antlr.TerminalNode

	// IsFilterExprContext differentiates from other interfaces.
	IsFilterExprContext()
}

type FilterExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFilterExprContext() *FilterExprContext {
	var p = new(FilterExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FilterLangRULE_filterExpr
	return p
}

func InitEmptyFilterExprContext(p *FilterExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FilterLangRULE_filterExpr
}

func (*FilterExprContext) IsFilterExprContext() {}

func NewFilterExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FilterExprContext {
	var p = new(FilterExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FilterLangRULE_filterExpr

	return p
}

func (s *FilterExprContext) GetParser() antlr.Parser { return s.parser }

func (s *FilterExprContext) Filter() IFilterContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFilterContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFilterContext)
}

func (s *FilterExprContext) EOF() antlr.TerminalNode {
	return s.GetToken(FilterLangEOF, 0)
}

func (s *FilterExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FilterExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FilterExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterLangListener); ok {
		listenerT.EnterFilterExpr(s)
	}
}

func (s *FilterExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterLangListener); ok {
		listenerT.ExitFilterExpr(s)
	}
}

func (s *FilterExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FilterLangVisitor:
		return t.VisitFilterExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FilterLang) FilterExpr() (localctx IFilterExprContext) {
	localctx = NewFilterExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, FilterLangRULE_filterExpr)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(16)
		p.filter(0)
	}
	{
		p.SetState(17)
		p.Match(FilterLangEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFilterContext is an interface to support dynamic dispatch.
type IFilterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsFilterContext differentiates from other interfaces.
	IsFilterContext()
}

type FilterContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFilterContext() *FilterContext {
	var p = new(FilterContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FilterLangRULE_filter
	return p
}

func InitEmptyFilterContext(p *FilterContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FilterLangRULE_filter
}

func (*FilterContext) IsFilterContext() {}

func NewFilterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FilterContext {
	var p = new(FilterContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FilterLangRULE_filter

	return p
}

func (s *FilterContext) GetParser() antlr.Parser { return s.parser }

func (s *FilterContext) CopyAll(ctx *FilterContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *FilterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FilterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type AndExprContext struct {
	FilterContext
}

func NewAndExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AndExprContext {
	var p = new(AndExprContext)

	InitEmptyFilterContext(&p.FilterContext)
	p.parser = parser
	p.CopyAll(ctx.(*FilterContext))

	return p
}

func (s *AndExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndExprContext) AllFilter() []IFilterContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFilterContext); ok {
			len++
		}
	}

	tst := make([]IFilterContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFilterContext); ok {
			tst[i] = t.(IFilterContext)
			i++
		}
	}

	return tst
}

func (s *AndExprContext) Filter(i int) IFilterContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFilterContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFilterContext)
}

func (s *AndExprContext) AND() antlr.TerminalNode {
	return s.GetToken(FilterLangAND, 0)
}

func (s *AndExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterLangListener); ok {
		listenerT.EnterAndExpr(s)
	}
}

func (s *AndExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterLangListener); ok {
		listenerT.ExitAndExpr(s)
	}
}

func (s *AndExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FilterLangVisitor:
		return t.VisitAndExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type GroupExprContext struct {
	FilterContext
}

func NewGroupExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GroupExprContext {
	var p = new(GroupExprContext)

	InitEmptyFilterContext(&p.FilterContext)
	p.parser = parser
	p.CopyAll(ctx.(*FilterContext))

	return p
}

func (s *GroupExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GroupExprContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(FilterLangLPAREN, 0)
}

func (s *GroupExprContext) Filter() IFilterContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFilterContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFilterContext)
}

func (s *GroupExprContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(FilterLangRPAREN, 0)
}

func (s *GroupExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterLangListener); ok {
		listenerT.EnterGroupExpr(s)
	}
}

func (s *GroupExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterLangListener); ok {
		listenerT.ExitGroupExpr(s)
	}
}

func (s *GroupExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FilterLangVisitor:
		return t.VisitGroupExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type NotExprContext struct {
	FilterContext
}

func NewNotExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotExprContext {
	var p = new(NotExprContext)

	InitEmptyFilterContext(&p.FilterContext)
	p.parser = parser
	p.CopyAll(ctx.(*FilterContext))

	return p
}

func (s *NotExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotExprContext) NOT() antlr.TerminalNode {
	return s.GetToken(FilterLangNOT, 0)
}

func (s *NotExprContext) Filter() IFilterContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFilterContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFilterContext)
}

func (s *NotExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterLangListener); ok {
		listenerT.EnterNotExpr(s)
	}
}

func (s *NotExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterLangListener); ok {
		listenerT.ExitNotExpr(s)
	}
}

func (s *NotExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FilterLangVisitor:
		return t.VisitNotExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type CondExprContext struct {
	FilterContext
}

func NewCondExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CondExprContext {
	var p = new(CondExprContext)

	InitEmptyFilterContext(&p.FilterContext)
	p.parser = parser
	p.CopyAll(ctx.(*FilterContext))

	return p
}

func (s *CondExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CondExprContext) Condition() IConditionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionContext)
}

func (s *CondExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterLangListener); ok {
		listenerT.EnterCondExpr(s)
	}
}

func (s *CondExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterLangListener); ok {
		listenerT.ExitCondExpr(s)
	}
}

func (s *CondExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FilterLangVisitor:
		return t.VisitCondExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type OrExprContext struct {
	FilterContext
}

func NewOrExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OrExprContext {
	var p = new(OrExprContext)

	InitEmptyFilterContext(&p.FilterContext)
	p.parser = parser
	p.CopyAll(ctx.(*FilterContext))

	return p
}

func (s *OrExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrExprContext) AllFilter() []IFilterContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFilterContext); ok {
			len++
		}
	}

	tst := make([]IFilterContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFilterContext); ok {
			tst[i] = t.(IFilterContext)
			i++
		}
	}

	return tst
}

func (s *OrExprContext) Filter(i int) IFilterContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFilterContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFilterContext)
}

func (s *OrExprContext) OR() antlr.TerminalNode {
	return s.GetToken(FilterLangOR, 0)
}

func (s *OrExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterLangListener); ok {
		listenerT.EnterOrExpr(s)
	}
}

func (s *OrExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterLangListener); ok {
		listenerT.ExitOrExpr(s)
	}
}

func (s *OrExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FilterLangVisitor:
		return t.VisitOrExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FilterLang) Filter() (localctx IFilterContext) {
	return p.filter(0)
}

func (p *FilterLang) filter(_p int) (localctx IFilterContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewFilterContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IFilterContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, FilterLangRULE_filter, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(27)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case FilterLangNOT:
		localctx = NewNotExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(20)
			p.Match(FilterLangNOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(21)
			p.filter(3)
		}

	case FilterLangLPAREN:
		localctx = NewGroupExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(22)
			p.Match(FilterLangLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(23)
			p.filter(0)
		}
		{
			p.SetState(24)
			p.Match(FilterLangRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case FilterLangIDENT:
		localctx = NewCondExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(26)
			p.Condition()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(37)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(35)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
			case 1:
				localctx = NewOrExprContext(p, NewFilterContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FilterLangRULE_filter)
				p.SetState(29)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(30)
					p.Match(FilterLangOR)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(31)
					p.filter(6)
				}

			case 2:
				localctx = NewAndExprContext(p, NewFilterContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FilterLangRULE_filter)
				p.SetState(32)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(33)
					p.Match(FilterLangAND)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(34)
					p.filter(5)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(39)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConditionContext is an interface to support dynamic dispatch.
type IConditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsConditionContext differentiates from other interfaces.
	IsConditionContext()
}

type ConditionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionContext() *ConditionContext {
	var p = new(ConditionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FilterLangRULE_condition
	return p
}

func InitEmptyConditionContext(p *ConditionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FilterLangRULE_condition
}

func (*ConditionContext) IsConditionContext() {}

func NewConditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionContext {
	var p = new(ConditionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FilterLangRULE_condition

	return p
}

func (s *ConditionContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionContext) CopyAll(ctx *ConditionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type NotExistsCheckContext struct {
	ConditionContext
}

func NewNotExistsCheckContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotExistsCheckContext {
	var p = new(NotExistsCheckContext)

	InitEmptyConditionContext(&p.ConditionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ConditionContext))

	return p
}

func (s *NotExistsCheckContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotExistsCheckContext) Path() IPathContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPathContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPathContext)
}

func (s *NotExistsCheckContext) NOT() antlr.TerminalNode {
	return s.GetToken(FilterLangNOT, 0)
}

func (s *NotExistsCheckContext) EXISTS() antlr.TerminalNode {
	return s.GetToken(FilterLangEXISTS, 0)
}

func (s *NotExistsCheckContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterLangListener); ok {
		listenerT.EnterNotExistsCheck(s)
	}
}

func (s *NotExistsCheckContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterLangListener); ok {
		listenerT.ExitNotExistsCheck(s)
	}
}

func (s *NotExistsCheckContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FilterLangVisitor:
		return t.VisitNotExistsCheck(s)

	default:
		return t.VisitChildren(s)
	}
}

type NotInListContext struct {
	ConditionContext
}

func NewNotInListContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotInListContext {
	var p = new(NotInListContext)

	InitEmptyConditionContext(&p.ConditionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ConditionContext))

	return p
}

func (s *NotInListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotInListContext) Path() IPathContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPathContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPathContext)
}

func (s *NotInListContext) NOT() antlr.TerminalNode {
	return s.GetToken(FilterLangNOT, 0)
}

func (s *NotInListContext) IN() antlr.TerminalNode {
	return s.GetToken(FilterLangIN, 0)
}

func (s *NotInListContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(FilterLangLPAREN, 0)
}

func (s *NotInListContext) ScalarList() IScalarListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IScalarListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IScalarListContext)
}

func (s *NotInListContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(FilterLangRPAREN, 0)
}

func (s *NotInListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterLangListener); ok {
		listenerT.EnterNotInList(s)
	}
}

func (s *NotInListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterLangListener); ok {
		listenerT.ExitNotInList(s)
	}
}

func (s *NotInListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FilterLangVisitor:
		return t.VisitNotInList(s)

	default:
		return t.VisitChildren(s)
	}
}

type InListContext struct {
	ConditionContext
}

func NewInListContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InListContext {
	var p = new(InListContext)

	InitEmptyConditionContext(&p.ConditionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ConditionContext))

	return p
}

func (s *InListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InListContext) Path() IPathContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPathContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPathContext)
}

func (s *InListContext) IN() antlr.TerminalNode {
	return s.GetToken(FilterLangIN, 0)
}

func (s *InListContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(FilterLangLPAREN, 0)
}

func (s *InListContext) ScalarList() IScalarListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IScalarListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IScalarListContext)
}

func (s *InListContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(FilterLangRPAREN, 0)
}

func (s *InListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterLangListener); ok {
		listenerT.EnterInList(s)
	}
}

func (s *InListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterLangListener); ok {
		listenerT.ExitInList(s)
	}
}

func (s *InListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FilterLangVisitor:
		return t.VisitInList(s)

	default:
		return t.VisitChildren(s)
	}
}

type MatchesOpContext struct {
	ConditionContext
}

func NewMatchesOpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MatchesOpContext {
	var p = new(MatchesOpContext)

	InitEmptyConditionContext(&p.ConditionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ConditionContext))

	return p
}

func (s *MatchesOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MatchesOpContext) Path() IPathContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPathContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPathContext)
}

func (s *MatchesOpContext) MATCHES() antlr.TerminalNode {
	return s.GetToken(FilterLangMATCHES, 0)
}

func (s *MatchesOpContext) Scalar() IScalarContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IScalarContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IScalarContext)
}

func (s *MatchesOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterLangListener); ok {
		listenerT.EnterMatchesOp(s)
	}
}

func (s *MatchesOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterLangListener); ok {
		listenerT.ExitMatchesOp(s)
	}
}

func (s *MatchesOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FilterLangVisitor:
		return t.VisitMatchesOp(s)

	default:
		return t.VisitChildren(s)
	}
}

type HasAllContext struct {
	ConditionContext
}

func NewHasAllContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *HasAllContext {
	var p = new(HasAllContext)

	InitEmptyConditionContext(&p.ConditionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ConditionContext))

	return p
}

func (s *HasAllContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HasAllContext) Path() IPathContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPathContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPathContext)
}

func (s *HasAllContext) HAS_ALL() antlr.TerminalNode {
	return s.GetToken(FilterLangHAS_ALL, 0)
}

func (s *HasAllContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(FilterLangLPAREN, 0)
}

func (s *HasAllContext) ScalarList() IScalarListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IScalarListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IScalarListContext)
}

func (s *HasAllContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(FilterLangRPAREN, 0)
}

func (s *HasAllContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterLangListener); ok {
		listenerT.EnterHasAll(s)
	}
}

func (s *HasAllContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterLangListener); ok {
		listenerT.ExitHasAll(s)
	}
}

func (s *HasAllContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FilterLangVisitor:
		return t.VisitHasAll(s)

	default:
		return t.VisitChildren(s)
	}
}

type ILikeOpContext struct {
	ConditionContext
}

func NewILikeOpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ILikeOpContext {
	var p = new(ILikeOpContext)

	InitEmptyConditionContext(&p.ConditionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ConditionContext))

	return p
}

func (s *ILikeOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ILikeOpContext) Path() IPathContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPathContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPathContext)
}

func (s *ILikeOpContext) ILIKE() antlr.TerminalNode {
	return s.GetToken(FilterLangILIKE, 0)
}

func (s *ILikeOpContext) Scalar() IScalarContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IScalarContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IScalarContext)
}

func (s *ILikeOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterLangListener); ok {
		listenerT.EnterILikeOp(s)
	}
}

func (s *ILikeOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterLangListener); ok {
		listenerT.ExitILikeOp(s)
	}
}

func (s *ILikeOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FilterLangVisitor:
		return t.VisitILikeOp(s)

	default:
		return t.VisitChildren(s)
	}
}

type CastCompareContext struct {
	ConditionContext
}

func NewCastCompareContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CastCompareContext {
	var p = new(CastCompareContext)

	InitEmptyConditionContext(&p.ConditionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ConditionContext))

	return p
}

func (s *CastCompareContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CastCompareContext) Path() IPathContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPathContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPathContext)
}

func (s *CastCompareContext) CAST_OP() antlr.TerminalNode {
	return s.GetToken(FilterLangCAST_OP, 0)
}

func (s *CastCompareContext) TypeName() ITypeNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeNameContext)
}

func (s *CastCompareContext) Op() IOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOpContext)
}

func (s *CastCompareContext) Scalar() IScalarContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IScalarContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IScalarContext)
}

func (s *CastCompareContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterLangListener); ok {
		listenerT.EnterCastCompare(s)
	}
}

func (s *CastCompareContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterLangListener); ok {
		listenerT.ExitCastCompare(s)
	}
}

func (s *CastCompareContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FilterLangVisitor:
		return t.VisitCastCompare(s)

	default:
		return t.VisitChildren(s)
	}
}

type LikeOpContext struct {
	ConditionContext
}

func NewLikeOpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LikeOpContext {
	var p = new(LikeOpContext)

	InitEmptyConditionContext(&p.ConditionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ConditionContext))

	return p
}

func (s *LikeOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LikeOpContext) Path() IPathContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPathContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPathContext)
}

func (s *LikeOpContext) LIKE() antlr.TerminalNode {
	return s.GetToken(FilterLangLIKE, 0)
}

func (s *LikeOpContext) Scalar() IScalarContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IScalarContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IScalarContext)
}

func (s *LikeOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterLangListener); ok {
		listenerT.EnterLikeOp(s)
	}
}

func (s *LikeOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterLangListener); ok {
		listenerT.ExitLikeOp(s)
	}
}

func (s *LikeOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FilterLangVisitor:
		return t.VisitLikeOp(s)

	default:
		return t.VisitChildren(s)
	}
}

type NotLikeOpContext struct {
	ConditionContext
}

func NewNotLikeOpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotLikeOpContext {
	var p = new(NotLikeOpContext)

	InitEmptyConditionContext(&p.ConditionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ConditionContext))

	return p
}

func (s *NotLikeOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotLikeOpContext) Path() IPathContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPathContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPathContext)
}

func (s *NotLikeOpContext) NOT() antlr.TerminalNode {
	return s.GetToken(FilterLangNOT, 0)
}

func (s *NotLikeOpContext) LIKE() antlr.TerminalNode {
	return s.GetToken(FilterLangLIKE, 0)
}

func (s *NotLikeOpContext) Scalar() IScalarContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IScalarContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IScalarContext)
}

func (s *NotLikeOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterLangListener); ok {
		listenerT.EnterNotLikeOp(s)
	}
}

func (s *NotLikeOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterLangListener); ok {
		listenerT.ExitNotLikeOp(s)
	}
}

func (s *NotLikeOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FilterLangVisitor:
		return t.VisitNotLikeOp(s)

	default:
		return t.VisitChildren(s)
	}
}

type IsNotNullContext struct {
	ConditionContext
}

func NewIsNotNullContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IsNotNullContext {
	var p = new(IsNotNullContext)

	InitEmptyConditionContext(&p.ConditionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ConditionContext))

	return p
}

func (s *IsNotNullContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IsNotNullContext) Path() IPathContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPathContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPathContext)
}

func (s *IsNotNullContext) IS() antlr.TerminalNode {
	return s.GetToken(FilterLangIS, 0)
}

func (s *IsNotNullContext) NOT() antlr.TerminalNode {
	return s.GetToken(FilterLangNOT, 0)
}

func (s *IsNotNullContext) NULL() antlr.TerminalNode {
	return s.GetToken(FilterLangNULL, 0)
}

func (s *IsNotNullContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterLangListener); ok {
		listenerT.EnterIsNotNull(s)
	}
}

func (s *IsNotNullContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterLangListener); ok {
		listenerT.ExitIsNotNull(s)
	}
}

func (s *IsNotNullContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FilterLangVisitor:
		return t.VisitIsNotNull(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExistsCheckContext struct {
	ConditionContext
}

func NewExistsCheckContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExistsCheckContext {
	var p = new(ExistsCheckContext)

	InitEmptyConditionContext(&p.ConditionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ConditionContext))

	return p
}

func (s *ExistsCheckContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExistsCheckContext) Path() IPathContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPathContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPathContext)
}

func (s *ExistsCheckContext) EXISTS() antlr.TerminalNode {
	return s.GetToken(FilterLangEXISTS, 0)
}

func (s *ExistsCheckContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterLangListener); ok {
		listenerT.EnterExistsCheck(s)
	}
}

func (s *ExistsCheckContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterLangListener); ok {
		listenerT.ExitExistsCheck(s)
	}
}

func (s *ExistsCheckContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FilterLangVisitor:
		return t.VisitExistsCheck(s)

	default:
		return t.VisitChildren(s)
	}
}

type CompareContext struct {
	ConditionContext
}

func NewCompareContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CompareContext {
	var p = new(CompareContext)

	InitEmptyConditionContext(&p.ConditionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ConditionContext))

	return p
}

func (s *CompareContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompareContext) Path() IPathContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPathContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPathContext)
}

func (s *CompareContext) Op() IOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOpContext)
}

func (s *CompareContext) Scalar() IScalarContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IScalarContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IScalarContext)
}

func (s *CompareContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterLangListener); ok {
		listenerT.EnterCompare(s)
	}
}

func (s *CompareContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterLangListener); ok {
		listenerT.ExitCompare(s)
	}
}

func (s *CompareContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FilterLangVisitor:
		return t.VisitCompare(s)

	default:
		return t.VisitChildren(s)
	}
}

type HasElementContext struct {
	ConditionContext
}

func NewHasElementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *HasElementContext {
	var p = new(HasElementContext)

	InitEmptyConditionContext(&p.ConditionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ConditionContext))

	return p
}

func (s *HasElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HasElementContext) Path() IPathContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPathContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPathContext)
}

func (s *HasElementContext) HAS() antlr.TerminalNode {
	return s.GetToken(FilterLangHAS, 0)
}

func (s *HasElementContext) Scalar() IScalarContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IScalarContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IScalarContext)
}

func (s *HasElementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterLangListener); ok {
		listenerT.EnterHasElement(s)
	}
}

func (s *HasElementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterLangListener); ok {
		listenerT.ExitHasElement(s)
	}
}

func (s *HasElementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FilterLangVisitor:
		return t.VisitHasElement(s)

	default:
		return t.VisitChildren(s)
	}
}

type ContainsOpContext struct {
	ConditionContext
}

func NewContainsOpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ContainsOpContext {
	var p = new(ContainsOpContext)

	InitEmptyConditionContext(&p.ConditionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ConditionContext))

	return p
}

func (s *ContainsOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContainsOpContext) Path() IPathContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPathContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPathContext)
}

func (s *ContainsOpContext) CONTAINS() antlr.TerminalNode {
	return s.GetToken(FilterLangCONTAINS, 0)
}

func (s *ContainsOpContext) Scalar() IScalarContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IScalarContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IScalarContext)
}

func (s *ContainsOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterLangListener); ok {
		listenerT.EnterContainsOp(s)
	}
}

func (s *ContainsOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterLangListener); ok {
		listenerT.ExitContainsOp(s)
	}
}

func (s *ContainsOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FilterLangVisitor:
		return t.VisitContainsOp(s)

	default:
		return t.VisitChildren(s)
	}
}

type IsNullContext struct {
	ConditionContext
}

func NewIsNullContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IsNullContext {
	var p = new(IsNullContext)

	InitEmptyConditionContext(&p.ConditionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ConditionContext))

	return p
}

func (s *IsNullContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IsNullContext) Path() IPathContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPathContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPathContext)
}

func (s *IsNullContext) IS() antlr.TerminalNode {
	return s.GetToken(FilterLangIS, 0)
}

func (s *IsNullContext) NULL() antlr.TerminalNode {
	return s.GetToken(FilterLangNULL, 0)
}

func (s *IsNullContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterLangListener); ok {
		listenerT.EnterIsNull(s)
	}
}

func (s *IsNullContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterLangListener); ok {
		listenerT.ExitIsNull(s)
	}
}

func (s *IsNullContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FilterLangVisitor:
		return t.VisitIsNull(s)

	default:
		return t.VisitChildren(s)
	}
}

type HasAnyContext struct {
	ConditionContext
}

func NewHasAnyContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *HasAnyContext {
	var p = new(HasAnyContext)

	InitEmptyConditionContext(&p.ConditionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ConditionContext))

	return p
}

func (s *HasAnyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HasAnyContext) Path() IPathContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPathContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPathContext)
}

func (s *HasAnyContext) HAS_ANY() antlr.TerminalNode {
	return s.GetToken(FilterLangHAS_ANY, 0)
}

func (s *HasAnyContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(FilterLangLPAREN, 0)
}

func (s *HasAnyContext) ScalarList() IScalarListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IScalarListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IScalarListContext)
}

func (s *HasAnyContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(FilterLangRPAREN, 0)
}

func (s *HasAnyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterLangListener); ok {
		listenerT.EnterHasAny(s)
	}
}

func (s *HasAnyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterLangListener); ok {
		listenerT.ExitHasAny(s)
	}
}

func (s *HasAnyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FilterLangVisitor:
		return t.VisitHasAny(s)

	default:
		return t.VisitChildren(s)
	}
}

type NotContainsOpContext struct {
	ConditionContext
}

func NewNotContainsOpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotContainsOpContext {
	var p = new(NotContainsOpContext)

	InitEmptyConditionContext(&p.ConditionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ConditionContext))

	return p
}

func (s *NotContainsOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotContainsOpContext) Path() IPathContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPathContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPathContext)
}

func (s *NotContainsOpContext) NOT() antlr.TerminalNode {
	return s.GetToken(FilterLangNOT, 0)
}

func (s *NotContainsOpContext) CONTAINS() antlr.TerminalNode {
	return s.GetToken(FilterLangCONTAINS, 0)
}

func (s *NotContainsOpContext) Scalar() IScalarContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IScalarContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IScalarContext)
}

func (s *NotContainsOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterLangListener); ok {
		listenerT.EnterNotContainsOp(s)
	}
}

func (s *NotContainsOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterLangListener); ok {
		listenerT.ExitNotContainsOp(s)
	}
}

func (s *NotContainsOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FilterLangVisitor:
		return t.VisitNotContainsOp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FilterLang) Condition() (localctx IConditionContext) {
	localctx = NewConditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, FilterLangRULE_condition)
	p.SetState(121)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		localctx = NewCastCompareContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(40)
			p.Path()
		}
		{
			p.SetState(41)
			p.Match(FilterLangCAST_OP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(42)
			p.TypeName()
		}
		{
			p.SetState(43)
			p.Op()
		}
		{
			p.SetState(44)
			p.Scalar()
		}

	case 2:
		localctx = NewCompareContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(46)
			p.Path()
		}
		{
			p.SetState(47)
			p.Op()
		}
		{
			p.SetState(48)
			p.Scalar()
		}

	case 3:
		localctx = NewInListContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(50)
			p.Path()
		}
		{
			p.SetState(51)
			p.Match(FilterLangIN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(52)
			p.Match(FilterLangLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(53)
			p.ScalarList()
		}
		{
			p.SetState(54)
			p.Match(FilterLangRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewNotInListContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(56)
			p.Path()
		}
		{
			p.SetState(57)
			p.Match(FilterLangNOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(58)
			p.Match(FilterLangIN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(59)
			p.Match(FilterLangLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(60)
			p.ScalarList()
		}
		{
			p.SetState(61)
			p.Match(FilterLangRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		localctx = NewHasElementContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(63)
			p.Path()
		}
		{
			p.SetState(64)
			p.Match(FilterLangHAS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(65)
			p.Scalar()
		}

	case 6:
		localctx = NewHasAnyContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(67)
			p.Path()
		}
		{
			p.SetState(68)
			p.Match(FilterLangHAS_ANY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(69)
			p.Match(FilterLangLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(70)
			p.ScalarList()
		}
		{
			p.SetState(71)
			p.Match(FilterLangRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		localctx = NewHasAllContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(73)
			p.Path()
		}
		{
			p.SetState(74)
			p.Match(FilterLangHAS_ALL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(75)
			p.Match(FilterLangLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(76)
			p.ScalarList()
		}
		{
			p.SetState(77)
			p.Match(FilterLangRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 8:
		localctx = NewContainsOpContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(79)
			p.Path()
		}
		{
			p.SetState(80)
			p.Match(FilterLangCONTAINS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(81)
			p.Scalar()
		}

	case 9:
		localctx = NewNotContainsOpContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(83)
			p.Path()
		}
		{
			p.SetState(84)
			p.Match(FilterLangNOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(85)
			p.Match(FilterLangCONTAINS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(86)
			p.Scalar()
		}

	case 10:
		localctx = NewLikeOpContext(p, localctx)
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(88)
			p.Path()
		}
		{
			p.SetState(89)
			p.Match(FilterLangLIKE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(90)
			p.Scalar()
		}

	case 11:
		localctx = NewILikeOpContext(p, localctx)
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(92)
			p.Path()
		}
		{
			p.SetState(93)
			p.Match(FilterLangILIKE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(94)
			p.Scalar()
		}

	case 12:
		localctx = NewNotLikeOpContext(p, localctx)
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(96)
			p.Path()
		}
		{
			p.SetState(97)
			p.Match(FilterLangNOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(98)
			p.Match(FilterLangLIKE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(99)
			p.Scalar()
		}

	case 13:
		localctx = NewMatchesOpContext(p, localctx)
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(101)
			p.Path()
		}
		{
			p.SetState(102)
			p.Match(FilterLangMATCHES)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(103)
			p.Scalar()
		}

	case 14:
		localctx = NewExistsCheckContext(p, localctx)
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(105)
			p.Path()
		}
		{
			p.SetState(106)
			p.Match(FilterLangEXISTS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 15:
		localctx = NewNotExistsCheckContext(p, localctx)
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(108)
			p.Path()
		}
		{
			p.SetState(109)
			p.Match(FilterLangNOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(110)
			p.Match(FilterLangEXISTS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 16:
		localctx = NewIsNullContext(p, localctx)
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(112)
			p.Path()
		}
		{
			p.SetState(113)
			p.Match(FilterLangIS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(114)
			p.Match(FilterLangNULL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 17:
		localctx = NewIsNotNullContext(p, localctx)
		p.EnterOuterAlt(localctx, 17)
		{
			p.SetState(116)
			p.Path()
		}
		{
			p.SetState(117)
			p.Match(FilterLangIS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(118)
			p.Match(FilterLangNOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(119)
			p.Match(FilterLangNULL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOpContext is an interface to support dynamic dispatch.
type IOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OP_EQ() antlr.TerminalNode
	OP_NEQ() antlr.TerminalNode
	OP_LT() antlr.TerminalNode
	OP_GT() antlr.TerminalNode
	OP_LTE() antlr.TerminalNode
	OP_GTE() antlr.TerminalNode

	// IsOpContext differentiates from other interfaces.
	IsOpContext()
}

type OpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOpContext() *OpContext {
	var p = new(OpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FilterLangRULE_op
	return p
}

func InitEmptyOpContext(p *OpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FilterLangRULE_op
}

func (*OpContext) IsOpContext() {}

func NewOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OpContext {
	var p = new(OpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FilterLangRULE_op

	return p
}

func (s *OpContext) GetParser() antlr.Parser { return s.parser }

func (s *OpContext) OP_EQ() antlr.TerminalNode {
	return s.GetToken(FilterLangOP_EQ, 0)
}

func (s *OpContext) OP_NEQ() antlr.TerminalNode {
	return s.GetToken(FilterLangOP_NEQ, 0)
}

func (s *OpContext) OP_LT() antlr.TerminalNode {
	return s.GetToken(FilterLangOP_LT, 0)
}

func (s *OpContext) OP_GT() antlr.TerminalNode {
	return s.GetToken(FilterLangOP_GT, 0)
}

func (s *OpContext) OP_LTE() antlr.TerminalNode {
	return s.GetToken(FilterLangOP_LTE, 0)
}

func (s *OpContext) OP_GTE() antlr.TerminalNode {
	return s.GetToken(FilterLangOP_GTE, 0)
}

func (s *OpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterLangListener); ok {
		listenerT.EnterOp(s)
	}
}

func (s *OpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterLangListener); ok {
		listenerT.ExitOp(s)
	}
}

func (s *OpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FilterLangVisitor:
		return t.VisitOp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FilterLang) Op() (localctx IOpContext) {
	localctx = NewOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, FilterLangRULE_op)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(123)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&252) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPathContext is an interface to support dynamic dispatch.
type IPathContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIDENT() []antlr.TerminalNode
	IDENT(i int) antlr.TerminalNode
	AllDOT() []antlr.TerminalNode
	DOT(i int) antlr.TerminalNode

	// IsPathContext differentiates from other interfaces.
	IsPathContext()
}

type PathContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPathContext() *PathContext {
	var p = new(PathContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FilterLangRULE_path
	return p
}

func InitEmptyPathContext(p *PathContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FilterLangRULE_path
}

func (*PathContext) IsPathContext() {}

func NewPathContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PathContext {
	var p = new(PathContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FilterLangRULE_path

	return p
}

func (s *PathContext) GetParser() antlr.Parser { return s.parser }

func (s *PathContext) AllIDENT() []antlr.TerminalNode {
	return s.GetTokens(FilterLangIDENT)
}

func (s *PathContext) IDENT(i int) antlr.TerminalNode {
	return s.GetToken(FilterLangIDENT, i)
}

func (s *PathContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(FilterLangDOT)
}

func (s *PathContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(FilterLangDOT, i)
}

func (s *PathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PathContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PathContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterLangListener); ok {
		listenerT.EnterPath(s)
	}
}

func (s *PathContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterLangListener); ok {
		listenerT.ExitPath(s)
	}
}

func (s *PathContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FilterLangVisitor:
		return t.VisitPath(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FilterLang) Path() (localctx IPathContext) {
	localctx = NewPathContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, FilterLangRULE_path)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(125)
		p.Match(FilterLangIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(130)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FilterLangDOT {
		{
			p.SetState(126)
			p.Match(FilterLangDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(127)
			p.Match(FilterLangIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(132)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypeNameContext is an interface to support dynamic dispatch.
type ITypeNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIDENT() []antlr.TerminalNode
	IDENT(i int) antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode

	// IsTypeNameContext differentiates from other interfaces.
	IsTypeNameContext()
}

type TypeNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeNameContext() *TypeNameContext {
	var p = new(TypeNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FilterLangRULE_typeName
	return p
}

func InitEmptyTypeNameContext(p *TypeNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FilterLangRULE_typeName
}

func (*TypeNameContext) IsTypeNameContext() {}

func NewTypeNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeNameContext {
	var p = new(TypeNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FilterLangRULE_typeName

	return p
}

func (s *TypeNameContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeNameContext) AllIDENT() []antlr.TerminalNode {
	return s.GetTokens(FilterLangIDENT)
}

func (s *TypeNameContext) IDENT(i int) antlr.TerminalNode {
	return s.GetToken(FilterLangIDENT, i)
}

func (s *TypeNameContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(FilterLangLPAREN, 0)
}

func (s *TypeNameContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(FilterLangRPAREN, 0)
}

func (s *TypeNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterLangListener); ok {
		listenerT.EnterTypeName(s)
	}
}

func (s *TypeNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterLangListener); ok {
		listenerT.ExitTypeName(s)
	}
}

func (s *TypeNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FilterLangVisitor:
		return t.VisitTypeName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FilterLang) TypeName() (localctx ITypeNameContext) {
	localctx = NewTypeNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, FilterLangRULE_typeName)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(133)
		p.Match(FilterLangIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(137)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FilterLangLPAREN {
		{
			p.SetState(134)
			p.Match(FilterLangLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(135)
			p.Match(FilterLangIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(136)
			p.Match(FilterLangRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IScalarListContext is an interface to support dynamic dispatch.
type IScalarListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllScalar() []IScalarContext
	Scalar(i int) IScalarContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsScalarListContext differentiates from other interfaces.
	IsScalarListContext()
}

type ScalarListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyScalarListContext() *ScalarListContext {
	var p = new(ScalarListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FilterLangRULE_scalarList
	return p
}

func InitEmptyScalarListContext(p *ScalarListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FilterLangRULE_scalarList
}

func (*ScalarListContext) IsScalarListContext() {}

func NewScalarListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ScalarListContext {
	var p = new(ScalarListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FilterLangRULE_scalarList

	return p
}

func (s *ScalarListContext) GetParser() antlr.Parser { return s.parser }

func (s *ScalarListContext) AllScalar() []IScalarContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IScalarContext); ok {
			len++
		}
	}

	tst := make([]IScalarContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IScalarContext); ok {
			tst[i] = t.(IScalarContext)
			i++
		}
	}

	return tst
}

func (s *ScalarListContext) Scalar(i int) IScalarContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IScalarContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IScalarContext)
}

func (s *ScalarListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(FilterLangCOMMA)
}

func (s *ScalarListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(FilterLangCOMMA, i)
}

func (s *ScalarListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ScalarListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ScalarListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterLangListener); ok {
		listenerT.EnterScalarList(s)
	}
}

func (s *ScalarListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterLangListener); ok {
		listenerT.ExitScalarList(s)
	}
}

func (s *ScalarListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FilterLangVisitor:
		return t.VisitScalarList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FilterLang) ScalarList() (localctx IScalarListContext) {
	localctx = NewScalarListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, FilterLangRULE_scalarList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(139)
		p.Scalar()
	}
	p.SetState(144)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FilterLangCOMMA {
		{
			p.SetState(140)
			p.Match(FilterLangCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(141)
			p.Scalar()
		}

		p.SetState(146)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IScalarContext is an interface to support dynamic dispatch.
type IScalarContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsScalarContext differentiates from other interfaces.
	IsScalarContext()
}

type ScalarContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyScalarContext() *ScalarContext {
	var p = new(ScalarContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FilterLangRULE_scalar
	return p
}

func InitEmptyScalarContext(p *ScalarContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FilterLangRULE_scalar
}

func (*ScalarContext) IsScalarContext() {}

func NewScalarContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ScalarContext {
	var p = new(ScalarContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FilterLangRULE_scalar

	return p
}

func (s *ScalarContext) GetParser() antlr.Parser { return s.parser }

func (s *ScalarContext) CopyAll(ctx *ScalarContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ScalarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ScalarContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ScalarFalseContext struct {
	ScalarContext
}

func NewScalarFalseContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ScalarFalseContext {
	var p = new(ScalarFalseContext)

	InitEmptyScalarContext(&p.ScalarContext)
	p.parser = parser
	p.CopyAll(ctx.(*ScalarContext))

	return p
}

func (s *ScalarFalseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ScalarFalseContext) FALSE() antlr.TerminalNode {
	return s.GetToken(FilterLangFALSE, 0)
}

func (s *ScalarFalseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterLangListener); ok {
		listenerT.EnterScalarFalse(s)
	}
}

func (s *ScalarFalseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterLangListener); ok {
		listenerT.ExitScalarFalse(s)
	}
}

func (s *ScalarFalseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FilterLangVisitor:
		return t.VisitScalarFalse(s)

	default:
		return t.VisitChildren(s)
	}
}

type ScalarStringContext struct {
	ScalarContext
}

func NewScalarStringContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ScalarStringContext {
	var p = new(ScalarStringContext)

	InitEmptyScalarContext(&p.ScalarContext)
	p.parser = parser
	p.CopyAll(ctx.(*ScalarContext))

	return p
}

func (s *ScalarStringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ScalarStringContext) STRING() antlr.TerminalNode {
	return s.GetToken(FilterLangSTRING, 0)
}

func (s *ScalarStringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterLangListener); ok {
		listenerT.EnterScalarString(s)
	}
}

func (s *ScalarStringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterLangListener); ok {
		listenerT.ExitScalarString(s)
	}
}

func (s *ScalarStringContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FilterLangVisitor:
		return t.VisitScalarString(s)

	default:
		return t.VisitChildren(s)
	}
}

type ScalarTrueContext struct {
	ScalarContext
}

func NewScalarTrueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ScalarTrueContext {
	var p = new(ScalarTrueContext)

	InitEmptyScalarContext(&p.ScalarContext)
	p.parser = parser
	p.CopyAll(ctx.(*ScalarContext))

	return p
}

func (s *ScalarTrueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ScalarTrueContext) TRUE() antlr.TerminalNode {
	return s.GetToken(FilterLangTRUE, 0)
}

func (s *ScalarTrueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterLangListener); ok {
		listenerT.EnterScalarTrue(s)
	}
}

func (s *ScalarTrueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterLangListener); ok {
		listenerT.ExitScalarTrue(s)
	}
}

func (s *ScalarTrueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FilterLangVisitor:
		return t.VisitScalarTrue(s)

	default:
		return t.VisitChildren(s)
	}
}

type ScalarNullContext struct {
	ScalarContext
}

func NewScalarNullContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ScalarNullContext {
	var p = new(ScalarNullContext)

	InitEmptyScalarContext(&p.ScalarContext)
	p.parser = parser
	p.CopyAll(ctx.(*ScalarContext))

	return p
}

func (s *ScalarNullContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ScalarNullContext) NULL() antlr.TerminalNode {
	return s.GetToken(FilterLangNULL, 0)
}

func (s *ScalarNullContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterLangListener); ok {
		listenerT.EnterScalarNull(s)
	}
}

func (s *ScalarNullContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterLangListener); ok {
		listenerT.ExitScalarNull(s)
	}
}

func (s *ScalarNullContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FilterLangVisitor:
		return t.VisitScalarNull(s)

	default:
		return t.VisitChildren(s)
	}
}

type ScalarFloatContext struct {
	ScalarContext
}

func NewScalarFloatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ScalarFloatContext {
	var p = new(ScalarFloatContext)

	InitEmptyScalarContext(&p.ScalarContext)
	p.parser = parser
	p.CopyAll(ctx.(*ScalarContext))

	return p
}

func (s *ScalarFloatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ScalarFloatContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(FilterLangFLOAT, 0)
}

func (s *ScalarFloatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterLangListener); ok {
		listenerT.EnterScalarFloat(s)
	}
}

func (s *ScalarFloatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterLangListener); ok {
		listenerT.ExitScalarFloat(s)
	}
}

func (s *ScalarFloatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FilterLangVisitor:
		return t.VisitScalarFloat(s)

	default:
		return t.VisitChildren(s)
	}
}

type ScalarIntContext struct {
	ScalarContext
}

func NewScalarIntContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ScalarIntContext {
	var p = new(ScalarIntContext)

	InitEmptyScalarContext(&p.ScalarContext)
	p.parser = parser
	p.CopyAll(ctx.(*ScalarContext))

	return p
}

func (s *ScalarIntContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ScalarIntContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(FilterLangINTEGER, 0)
}

func (s *ScalarIntContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterLangListener); ok {
		listenerT.EnterScalarInt(s)
	}
}

func (s *ScalarIntContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterLangListener); ok {
		listenerT.ExitScalarInt(s)
	}
}

func (s *ScalarIntContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FilterLangVisitor:
		return t.VisitScalarInt(s)

	default:
		return t.VisitChildren(s)
	}
}

type ScalarBareWordContext struct {
	ScalarContext
}

func NewScalarBareWordContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ScalarBareWordContext {
	var p = new(ScalarBareWordContext)

	InitEmptyScalarContext(&p.ScalarContext)
	p.parser = parser
	p.CopyAll(ctx.(*ScalarContext))

	return p
}

func (s *ScalarBareWordContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ScalarBareWordContext) IDENT() antlr.TerminalNode {
	return s.GetToken(FilterLangIDENT, 0)
}

func (s *ScalarBareWordContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterLangListener); ok {
		listenerT.EnterScalarBareWord(s)
	}
}

func (s *ScalarBareWordContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterLangListener); ok {
		listenerT.ExitScalarBareWord(s)
	}
}

func (s *ScalarBareWordContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FilterLangVisitor:
		return t.VisitScalarBareWord(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FilterLang) Scalar() (localctx IScalarContext) {
	localctx = NewScalarContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, FilterLangRULE_scalar)
	p.SetState(154)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case FilterLangSTRING:
		localctx = NewScalarStringContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(147)
			p.Match(FilterLangSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case FilterLangFLOAT:
		localctx = NewScalarFloatContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(148)
			p.Match(FilterLangFLOAT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case FilterLangINTEGER:
		localctx = NewScalarIntContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(149)
			p.Match(FilterLangINTEGER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case FilterLangTRUE:
		localctx = NewScalarTrueContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(150)
			p.Match(FilterLangTRUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case FilterLangFALSE:
		localctx = NewScalarFalseContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(151)
			p.Match(FilterLangFALSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case FilterLangNULL:
		localctx = NewScalarNullContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(152)
			p.Match(FilterLangNULL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case FilterLangIDENT:
		localctx = NewScalarBareWordContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(153)
			p.Match(FilterLangIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *FilterLang) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *FilterContext = nil
		if localctx != nil {
			t = localctx.(*FilterContext)
		}
		return p.Filter_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *FilterLang) Filter_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 4)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
