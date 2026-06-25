// Code generated from PipeLang.g4 by ANTLR 4.13.2. DO NOT EDIT.

package grammar // PipeLang
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

type PipeLang struct {
	*antlr.BaseParser
}

var PipeLangParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func pipelangParserInit() {
	staticData := &PipeLangParserStaticData
	staticData.SymbolicNames = []string{
		"", "COMMENT", "WS_BLANK", "WS", "TYPE", "NAME", "TAGS", "OWNER", "DESTINATION",
		"SCHEDULE", "UNIQUE_KEY", "DESCRIPTION_ML", "DESCRIPTION", "SOURCES",
		"PARAMS", "PIPELINE", "REST_OF_LINE", "BLOCK_LINE", "BLOCK_BLANK", "BLOCK_COMMENT",
		"SOURCE_LINE", "SRCS_BLANK", "SRCS_COMMENT", "PARAM_LINE", "PRMS_BLANK",
		"PRMS_COMMENT", "NODE_HEADER", "PRQL_LINE", "PIPE_BLANK", "PIPE_COMMENT",
	}
	staticData.RuleNames = []string{
		"pipeFile", "directive", "pipelineBlock", "pipelineNode",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 29, 71, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 1, 0, 5,
		0, 10, 8, 0, 10, 0, 12, 0, 13, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 25, 8, 1, 10, 1, 12, 1, 28, 9, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 42,
		8, 1, 10, 1, 12, 1, 45, 9, 1, 1, 1, 1, 1, 5, 1, 49, 8, 1, 10, 1, 12, 1,
		52, 9, 1, 1, 1, 1, 1, 3, 1, 56, 8, 1, 1, 2, 5, 2, 59, 8, 2, 10, 2, 12,
		2, 62, 9, 2, 1, 3, 1, 3, 5, 3, 66, 8, 3, 10, 3, 12, 3, 69, 9, 3, 1, 3,
		0, 0, 4, 0, 2, 4, 6, 0, 0, 83, 0, 11, 1, 0, 0, 0, 2, 55, 1, 0, 0, 0, 4,
		60, 1, 0, 0, 0, 6, 63, 1, 0, 0, 0, 8, 10, 3, 2, 1, 0, 9, 8, 1, 0, 0, 0,
		10, 13, 1, 0, 0, 0, 11, 9, 1, 0, 0, 0, 11, 12, 1, 0, 0, 0, 12, 14, 1, 0,
		0, 0, 13, 11, 1, 0, 0, 0, 14, 15, 5, 0, 0, 1, 15, 1, 1, 0, 0, 0, 16, 17,
		5, 4, 0, 0, 17, 56, 5, 16, 0, 0, 18, 19, 5, 5, 0, 0, 19, 56, 5, 16, 0,
		0, 20, 21, 5, 12, 0, 0, 21, 56, 5, 16, 0, 0, 22, 26, 5, 11, 0, 0, 23, 25,
		5, 17, 0, 0, 24, 23, 1, 0, 0, 0, 25, 28, 1, 0, 0, 0, 26, 24, 1, 0, 0, 0,
		26, 27, 1, 0, 0, 0, 27, 56, 1, 0, 0, 0, 28, 26, 1, 0, 0, 0, 29, 30, 5,
		6, 0, 0, 30, 56, 5, 16, 0, 0, 31, 32, 5, 7, 0, 0, 32, 56, 5, 16, 0, 0,
		33, 34, 5, 8, 0, 0, 34, 56, 5, 16, 0, 0, 35, 36, 5, 9, 0, 0, 36, 56, 5,
		16, 0, 0, 37, 38, 5, 10, 0, 0, 38, 56, 5, 16, 0, 0, 39, 43, 5, 13, 0, 0,
		40, 42, 5, 20, 0, 0, 41, 40, 1, 0, 0, 0, 42, 45, 1, 0, 0, 0, 43, 41, 1,
		0, 0, 0, 43, 44, 1, 0, 0, 0, 44, 56, 1, 0, 0, 0, 45, 43, 1, 0, 0, 0, 46,
		50, 5, 14, 0, 0, 47, 49, 5, 23, 0, 0, 48, 47, 1, 0, 0, 0, 49, 52, 1, 0,
		0, 0, 50, 48, 1, 0, 0, 0, 50, 51, 1, 0, 0, 0, 51, 56, 1, 0, 0, 0, 52, 50,
		1, 0, 0, 0, 53, 54, 5, 15, 0, 0, 54, 56, 3, 4, 2, 0, 55, 16, 1, 0, 0, 0,
		55, 18, 1, 0, 0, 0, 55, 20, 1, 0, 0, 0, 55, 22, 1, 0, 0, 0, 55, 29, 1,
		0, 0, 0, 55, 31, 1, 0, 0, 0, 55, 33, 1, 0, 0, 0, 55, 35, 1, 0, 0, 0, 55,
		37, 1, 0, 0, 0, 55, 39, 1, 0, 0, 0, 55, 46, 1, 0, 0, 0, 55, 53, 1, 0, 0,
		0, 56, 3, 1, 0, 0, 0, 57, 59, 3, 6, 3, 0, 58, 57, 1, 0, 0, 0, 59, 62, 1,
		0, 0, 0, 60, 58, 1, 0, 0, 0, 60, 61, 1, 0, 0, 0, 61, 5, 1, 0, 0, 0, 62,
		60, 1, 0, 0, 0, 63, 67, 5, 26, 0, 0, 64, 66, 5, 27, 0, 0, 65, 64, 1, 0,
		0, 0, 66, 69, 1, 0, 0, 0, 67, 65, 1, 0, 0, 0, 67, 68, 1, 0, 0, 0, 68, 7,
		1, 0, 0, 0, 69, 67, 1, 0, 0, 0, 7, 11, 26, 43, 50, 55, 60, 67,
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

// PipeLangInit initializes any static state used to implement PipeLang. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewPipeLang(). You can call this function if you wish to initialize the static state ahead
// of time.
func PipeLangInit() {
	staticData := &PipeLangParserStaticData
	staticData.once.Do(pipelangParserInit)
}

// NewPipeLang produces a new parser instance for the optional input antlr.TokenStream.
func NewPipeLang(input antlr.TokenStream) *PipeLang {
	PipeLangInit()
	this := new(PipeLang)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &PipeLangParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "PipeLang.g4"

	return this
}

// PipeLang tokens.
const (
	PipeLangEOF            = antlr.TokenEOF
	PipeLangCOMMENT        = 1
	PipeLangWS_BLANK       = 2
	PipeLangWS             = 3
	PipeLangTYPE           = 4
	PipeLangNAME           = 5
	PipeLangTAGS           = 6
	PipeLangOWNER          = 7
	PipeLangDESTINATION    = 8
	PipeLangSCHEDULE       = 9
	PipeLangUNIQUE_KEY     = 10
	PipeLangDESCRIPTION_ML = 11
	PipeLangDESCRIPTION    = 12
	PipeLangSOURCES        = 13
	PipeLangPARAMS         = 14
	PipeLangPIPELINE       = 15
	PipeLangREST_OF_LINE   = 16
	PipeLangBLOCK_LINE     = 17
	PipeLangBLOCK_BLANK    = 18
	PipeLangBLOCK_COMMENT  = 19
	PipeLangSOURCE_LINE    = 20
	PipeLangSRCS_BLANK     = 21
	PipeLangSRCS_COMMENT   = 22
	PipeLangPARAM_LINE     = 23
	PipeLangPRMS_BLANK     = 24
	PipeLangPRMS_COMMENT   = 25
	PipeLangNODE_HEADER    = 26
	PipeLangPRQL_LINE      = 27
	PipeLangPIPE_BLANK     = 28
	PipeLangPIPE_COMMENT   = 29
)

// PipeLang rules.
const (
	PipeLangRULE_pipeFile      = 0
	PipeLangRULE_directive     = 1
	PipeLangRULE_pipelineBlock = 2
	PipeLangRULE_pipelineNode  = 3
)

// IPipeFileContext is an interface to support dynamic dispatch.
type IPipeFileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllDirective() []IDirectiveContext
	Directive(i int) IDirectiveContext

	// IsPipeFileContext differentiates from other interfaces.
	IsPipeFileContext()
}

type PipeFileContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPipeFileContext() *PipeFileContext {
	var p = new(PipeFileContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PipeLangRULE_pipeFile
	return p
}

func InitEmptyPipeFileContext(p *PipeFileContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PipeLangRULE_pipeFile
}

func (*PipeFileContext) IsPipeFileContext() {}

func NewPipeFileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PipeFileContext {
	var p = new(PipeFileContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PipeLangRULE_pipeFile

	return p
}

func (s *PipeFileContext) GetParser() antlr.Parser { return s.parser }

func (s *PipeFileContext) EOF() antlr.TerminalNode {
	return s.GetToken(PipeLangEOF, 0)
}

func (s *PipeFileContext) AllDirective() []IDirectiveContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDirectiveContext); ok {
			len++
		}
	}

	tst := make([]IDirectiveContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDirectiveContext); ok {
			tst[i] = t.(IDirectiveContext)
			i++
		}
	}

	return tst
}

func (s *PipeFileContext) Directive(i int) IDirectiveContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDirectiveContext); ok {
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

	return t.(IDirectiveContext)
}

func (s *PipeFileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PipeFileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PipeFileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeLangListener); ok {
		listenerT.EnterPipeFile(s)
	}
}

func (s *PipeFileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeLangListener); ok {
		listenerT.ExitPipeFile(s)
	}
}

func (s *PipeFileContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PipeLangVisitor:
		return t.VisitPipeFile(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PipeLang) PipeFile() (localctx IPipeFileContext) {
	localctx = NewPipeFileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, PipeLangRULE_pipeFile)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(11)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&65520) != 0 {
		{
			p.SetState(8)
			p.Directive()
		}

		p.SetState(13)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(14)
		p.Match(PipeLangEOF)
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

// IDirectiveContext is an interface to support dynamic dispatch.
type IDirectiveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsDirectiveContext differentiates from other interfaces.
	IsDirectiveContext()
}

type DirectiveContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDirectiveContext() *DirectiveContext {
	var p = new(DirectiveContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PipeLangRULE_directive
	return p
}

func InitEmptyDirectiveContext(p *DirectiveContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PipeLangRULE_directive
}

func (*DirectiveContext) IsDirectiveContext() {}

func NewDirectiveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DirectiveContext {
	var p = new(DirectiveContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PipeLangRULE_directive

	return p
}

func (s *DirectiveContext) GetParser() antlr.Parser { return s.parser }

func (s *DirectiveContext) CopyAll(ctx *DirectiveContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *DirectiveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DirectiveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type OwnerContext struct {
	DirectiveContext
}

func NewOwnerContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OwnerContext {
	var p = new(OwnerContext)

	InitEmptyDirectiveContext(&p.DirectiveContext)
	p.parser = parser
	p.CopyAll(ctx.(*DirectiveContext))

	return p
}

func (s *OwnerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OwnerContext) OWNER() antlr.TerminalNode {
	return s.GetToken(PipeLangOWNER, 0)
}

func (s *OwnerContext) REST_OF_LINE() antlr.TerminalNode {
	return s.GetToken(PipeLangREST_OF_LINE, 0)
}

func (s *OwnerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeLangListener); ok {
		listenerT.EnterOwner(s)
	}
}

func (s *OwnerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeLangListener); ok {
		listenerT.ExitOwner(s)
	}
}

func (s *OwnerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PipeLangVisitor:
		return t.VisitOwner(s)

	default:
		return t.VisitChildren(s)
	}
}

type PipelineClauseContext struct {
	DirectiveContext
}

func NewPipelineClauseContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PipelineClauseContext {
	var p = new(PipelineClauseContext)

	InitEmptyDirectiveContext(&p.DirectiveContext)
	p.parser = parser
	p.CopyAll(ctx.(*DirectiveContext))

	return p
}

func (s *PipelineClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PipelineClauseContext) PIPELINE() antlr.TerminalNode {
	return s.GetToken(PipeLangPIPELINE, 0)
}

func (s *PipelineClauseContext) PipelineBlock() IPipelineBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPipelineBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPipelineBlockContext)
}

func (s *PipelineClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeLangListener); ok {
		listenerT.EnterPipelineClause(s)
	}
}

func (s *PipelineClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeLangListener); ok {
		listenerT.ExitPipelineClause(s)
	}
}

func (s *PipelineClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PipeLangVisitor:
		return t.VisitPipelineClause(s)

	default:
		return t.VisitChildren(s)
	}
}

type ScheduleContext struct {
	DirectiveContext
}

func NewScheduleContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ScheduleContext {
	var p = new(ScheduleContext)

	InitEmptyDirectiveContext(&p.DirectiveContext)
	p.parser = parser
	p.CopyAll(ctx.(*DirectiveContext))

	return p
}

func (s *ScheduleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ScheduleContext) SCHEDULE() antlr.TerminalNode {
	return s.GetToken(PipeLangSCHEDULE, 0)
}

func (s *ScheduleContext) REST_OF_LINE() antlr.TerminalNode {
	return s.GetToken(PipeLangREST_OF_LINE, 0)
}

func (s *ScheduleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeLangListener); ok {
		listenerT.EnterSchedule(s)
	}
}

func (s *ScheduleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeLangListener); ok {
		listenerT.ExitSchedule(s)
	}
}

func (s *ScheduleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PipeLangVisitor:
		return t.VisitSchedule(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParamsClauseContext struct {
	DirectiveContext
}

func NewParamsClauseContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParamsClauseContext {
	var p = new(ParamsClauseContext)

	InitEmptyDirectiveContext(&p.DirectiveContext)
	p.parser = parser
	p.CopyAll(ctx.(*DirectiveContext))

	return p
}

func (s *ParamsClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamsClauseContext) PARAMS() antlr.TerminalNode {
	return s.GetToken(PipeLangPARAMS, 0)
}

func (s *ParamsClauseContext) AllPARAM_LINE() []antlr.TerminalNode {
	return s.GetTokens(PipeLangPARAM_LINE)
}

func (s *ParamsClauseContext) PARAM_LINE(i int) antlr.TerminalNode {
	return s.GetToken(PipeLangPARAM_LINE, i)
}

func (s *ParamsClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeLangListener); ok {
		listenerT.EnterParamsClause(s)
	}
}

func (s *ParamsClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeLangListener); ok {
		listenerT.ExitParamsClause(s)
	}
}

func (s *ParamsClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PipeLangVisitor:
		return t.VisitParamsClause(s)

	default:
		return t.VisitChildren(s)
	}
}

type UniqueKeyContext struct {
	DirectiveContext
}

func NewUniqueKeyContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UniqueKeyContext {
	var p = new(UniqueKeyContext)

	InitEmptyDirectiveContext(&p.DirectiveContext)
	p.parser = parser
	p.CopyAll(ctx.(*DirectiveContext))

	return p
}

func (s *UniqueKeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UniqueKeyContext) UNIQUE_KEY() antlr.TerminalNode {
	return s.GetToken(PipeLangUNIQUE_KEY, 0)
}

func (s *UniqueKeyContext) REST_OF_LINE() antlr.TerminalNode {
	return s.GetToken(PipeLangREST_OF_LINE, 0)
}

func (s *UniqueKeyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeLangListener); ok {
		listenerT.EnterUniqueKey(s)
	}
}

func (s *UniqueKeyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeLangListener); ok {
		listenerT.ExitUniqueKey(s)
	}
}

func (s *UniqueKeyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PipeLangVisitor:
		return t.VisitUniqueKey(s)

	default:
		return t.VisitChildren(s)
	}
}

type NameContext struct {
	DirectiveContext
}

func NewNameContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NameContext {
	var p = new(NameContext)

	InitEmptyDirectiveContext(&p.DirectiveContext)
	p.parser = parser
	p.CopyAll(ctx.(*DirectiveContext))

	return p
}

func (s *NameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NameContext) NAME() antlr.TerminalNode {
	return s.GetToken(PipeLangNAME, 0)
}

func (s *NameContext) REST_OF_LINE() antlr.TerminalNode {
	return s.GetToken(PipeLangREST_OF_LINE, 0)
}

func (s *NameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeLangListener); ok {
		listenerT.EnterName(s)
	}
}

func (s *NameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeLangListener); ok {
		listenerT.ExitName(s)
	}
}

func (s *NameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PipeLangVisitor:
		return t.VisitName(s)

	default:
		return t.VisitChildren(s)
	}
}

type DestinationContext struct {
	DirectiveContext
}

func NewDestinationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DestinationContext {
	var p = new(DestinationContext)

	InitEmptyDirectiveContext(&p.DirectiveContext)
	p.parser = parser
	p.CopyAll(ctx.(*DirectiveContext))

	return p
}

func (s *DestinationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DestinationContext) DESTINATION() antlr.TerminalNode {
	return s.GetToken(PipeLangDESTINATION, 0)
}

func (s *DestinationContext) REST_OF_LINE() antlr.TerminalNode {
	return s.GetToken(PipeLangREST_OF_LINE, 0)
}

func (s *DestinationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeLangListener); ok {
		listenerT.EnterDestination(s)
	}
}

func (s *DestinationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeLangListener); ok {
		listenerT.ExitDestination(s)
	}
}

func (s *DestinationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PipeLangVisitor:
		return t.VisitDestination(s)

	default:
		return t.VisitChildren(s)
	}
}

type DescriptionContext struct {
	DirectiveContext
}

func NewDescriptionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DescriptionContext {
	var p = new(DescriptionContext)

	InitEmptyDirectiveContext(&p.DirectiveContext)
	p.parser = parser
	p.CopyAll(ctx.(*DirectiveContext))

	return p
}

func (s *DescriptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DescriptionContext) DESCRIPTION() antlr.TerminalNode {
	return s.GetToken(PipeLangDESCRIPTION, 0)
}

func (s *DescriptionContext) REST_OF_LINE() antlr.TerminalNode {
	return s.GetToken(PipeLangREST_OF_LINE, 0)
}

func (s *DescriptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeLangListener); ok {
		listenerT.EnterDescription(s)
	}
}

func (s *DescriptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeLangListener); ok {
		listenerT.ExitDescription(s)
	}
}

func (s *DescriptionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PipeLangVisitor:
		return t.VisitDescription(s)

	default:
		return t.VisitChildren(s)
	}
}

type TypeContext struct {
	DirectiveContext
}

func NewTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeContext {
	var p = new(TypeContext)

	InitEmptyDirectiveContext(&p.DirectiveContext)
	p.parser = parser
	p.CopyAll(ctx.(*DirectiveContext))

	return p
}

func (s *TypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeContext) TYPE() antlr.TerminalNode {
	return s.GetToken(PipeLangTYPE, 0)
}

func (s *TypeContext) REST_OF_LINE() antlr.TerminalNode {
	return s.GetToken(PipeLangREST_OF_LINE, 0)
}

func (s *TypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeLangListener); ok {
		listenerT.EnterType(s)
	}
}

func (s *TypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeLangListener); ok {
		listenerT.ExitType(s)
	}
}

func (s *TypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PipeLangVisitor:
		return t.VisitType(s)

	default:
		return t.VisitChildren(s)
	}
}

type SourcesClauseContext struct {
	DirectiveContext
}

func NewSourcesClauseContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SourcesClauseContext {
	var p = new(SourcesClauseContext)

	InitEmptyDirectiveContext(&p.DirectiveContext)
	p.parser = parser
	p.CopyAll(ctx.(*DirectiveContext))

	return p
}

func (s *SourcesClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SourcesClauseContext) SOURCES() antlr.TerminalNode {
	return s.GetToken(PipeLangSOURCES, 0)
}

func (s *SourcesClauseContext) AllSOURCE_LINE() []antlr.TerminalNode {
	return s.GetTokens(PipeLangSOURCE_LINE)
}

func (s *SourcesClauseContext) SOURCE_LINE(i int) antlr.TerminalNode {
	return s.GetToken(PipeLangSOURCE_LINE, i)
}

func (s *SourcesClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeLangListener); ok {
		listenerT.EnterSourcesClause(s)
	}
}

func (s *SourcesClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeLangListener); ok {
		listenerT.ExitSourcesClause(s)
	}
}

func (s *SourcesClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PipeLangVisitor:
		return t.VisitSourcesClause(s)

	default:
		return t.VisitChildren(s)
	}
}

type DescriptionMLContext struct {
	DirectiveContext
}

func NewDescriptionMLContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DescriptionMLContext {
	var p = new(DescriptionMLContext)

	InitEmptyDirectiveContext(&p.DirectiveContext)
	p.parser = parser
	p.CopyAll(ctx.(*DirectiveContext))

	return p
}

func (s *DescriptionMLContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DescriptionMLContext) DESCRIPTION_ML() antlr.TerminalNode {
	return s.GetToken(PipeLangDESCRIPTION_ML, 0)
}

func (s *DescriptionMLContext) AllBLOCK_LINE() []antlr.TerminalNode {
	return s.GetTokens(PipeLangBLOCK_LINE)
}

func (s *DescriptionMLContext) BLOCK_LINE(i int) antlr.TerminalNode {
	return s.GetToken(PipeLangBLOCK_LINE, i)
}

func (s *DescriptionMLContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeLangListener); ok {
		listenerT.EnterDescriptionML(s)
	}
}

func (s *DescriptionMLContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeLangListener); ok {
		listenerT.ExitDescriptionML(s)
	}
}

func (s *DescriptionMLContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PipeLangVisitor:
		return t.VisitDescriptionML(s)

	default:
		return t.VisitChildren(s)
	}
}

type TagsContext struct {
	DirectiveContext
}

func NewTagsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TagsContext {
	var p = new(TagsContext)

	InitEmptyDirectiveContext(&p.DirectiveContext)
	p.parser = parser
	p.CopyAll(ctx.(*DirectiveContext))

	return p
}

func (s *TagsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TagsContext) TAGS() antlr.TerminalNode {
	return s.GetToken(PipeLangTAGS, 0)
}

func (s *TagsContext) REST_OF_LINE() antlr.TerminalNode {
	return s.GetToken(PipeLangREST_OF_LINE, 0)
}

func (s *TagsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeLangListener); ok {
		listenerT.EnterTags(s)
	}
}

func (s *TagsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeLangListener); ok {
		listenerT.ExitTags(s)
	}
}

func (s *TagsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PipeLangVisitor:
		return t.VisitTags(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PipeLang) Directive() (localctx IDirectiveContext) {
	localctx = NewDirectiveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, PipeLangRULE_directive)
	var _la int

	p.SetState(55)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PipeLangTYPE:
		localctx = NewTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(16)
			p.Match(PipeLangTYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(17)
			p.Match(PipeLangREST_OF_LINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PipeLangNAME:
		localctx = NewNameContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(18)
			p.Match(PipeLangNAME)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(19)
			p.Match(PipeLangREST_OF_LINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PipeLangDESCRIPTION:
		localctx = NewDescriptionContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(20)
			p.Match(PipeLangDESCRIPTION)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(21)
			p.Match(PipeLangREST_OF_LINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PipeLangDESCRIPTION_ML:
		localctx = NewDescriptionMLContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(22)
			p.Match(PipeLangDESCRIPTION_ML)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(26)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == PipeLangBLOCK_LINE {
			{
				p.SetState(23)
				p.Match(PipeLangBLOCK_LINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(28)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case PipeLangTAGS:
		localctx = NewTagsContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(29)
			p.Match(PipeLangTAGS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(30)
			p.Match(PipeLangREST_OF_LINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PipeLangOWNER:
		localctx = NewOwnerContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(31)
			p.Match(PipeLangOWNER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(32)
			p.Match(PipeLangREST_OF_LINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PipeLangDESTINATION:
		localctx = NewDestinationContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(33)
			p.Match(PipeLangDESTINATION)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(34)
			p.Match(PipeLangREST_OF_LINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PipeLangSCHEDULE:
		localctx = NewScheduleContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(35)
			p.Match(PipeLangSCHEDULE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(36)
			p.Match(PipeLangREST_OF_LINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PipeLangUNIQUE_KEY:
		localctx = NewUniqueKeyContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(37)
			p.Match(PipeLangUNIQUE_KEY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(38)
			p.Match(PipeLangREST_OF_LINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PipeLangSOURCES:
		localctx = NewSourcesClauseContext(p, localctx)
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(39)
			p.Match(PipeLangSOURCES)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(43)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == PipeLangSOURCE_LINE {
			{
				p.SetState(40)
				p.Match(PipeLangSOURCE_LINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(45)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case PipeLangPARAMS:
		localctx = NewParamsClauseContext(p, localctx)
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(46)
			p.Match(PipeLangPARAMS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(50)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == PipeLangPARAM_LINE {
			{
				p.SetState(47)
				p.Match(PipeLangPARAM_LINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(52)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case PipeLangPIPELINE:
		localctx = NewPipelineClauseContext(p, localctx)
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(53)
			p.Match(PipeLangPIPELINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(54)
			p.PipelineBlock()
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

// IPipelineBlockContext is an interface to support dynamic dispatch.
type IPipelineBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllPipelineNode() []IPipelineNodeContext
	PipelineNode(i int) IPipelineNodeContext

	// IsPipelineBlockContext differentiates from other interfaces.
	IsPipelineBlockContext()
}

type PipelineBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPipelineBlockContext() *PipelineBlockContext {
	var p = new(PipelineBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PipeLangRULE_pipelineBlock
	return p
}

func InitEmptyPipelineBlockContext(p *PipelineBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PipeLangRULE_pipelineBlock
}

func (*PipelineBlockContext) IsPipelineBlockContext() {}

func NewPipelineBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PipelineBlockContext {
	var p = new(PipelineBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PipeLangRULE_pipelineBlock

	return p
}

func (s *PipelineBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *PipelineBlockContext) AllPipelineNode() []IPipelineNodeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPipelineNodeContext); ok {
			len++
		}
	}

	tst := make([]IPipelineNodeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPipelineNodeContext); ok {
			tst[i] = t.(IPipelineNodeContext)
			i++
		}
	}

	return tst
}

func (s *PipelineBlockContext) PipelineNode(i int) IPipelineNodeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPipelineNodeContext); ok {
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

	return t.(IPipelineNodeContext)
}

func (s *PipelineBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PipelineBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PipelineBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeLangListener); ok {
		listenerT.EnterPipelineBlock(s)
	}
}

func (s *PipelineBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeLangListener); ok {
		listenerT.ExitPipelineBlock(s)
	}
}

func (s *PipelineBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PipeLangVisitor:
		return t.VisitPipelineBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PipeLang) PipelineBlock() (localctx IPipelineBlockContext) {
	localctx = NewPipelineBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, PipeLangRULE_pipelineBlock)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(60)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == PipeLangNODE_HEADER {
		{
			p.SetState(57)
			p.PipelineNode()
		}

		p.SetState(62)
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

// IPipelineNodeContext is an interface to support dynamic dispatch.
type IPipelineNodeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NODE_HEADER() antlr.TerminalNode
	AllPRQL_LINE() []antlr.TerminalNode
	PRQL_LINE(i int) antlr.TerminalNode

	// IsPipelineNodeContext differentiates from other interfaces.
	IsPipelineNodeContext()
}

type PipelineNodeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPipelineNodeContext() *PipelineNodeContext {
	var p = new(PipelineNodeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PipeLangRULE_pipelineNode
	return p
}

func InitEmptyPipelineNodeContext(p *PipelineNodeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PipeLangRULE_pipelineNode
}

func (*PipelineNodeContext) IsPipelineNodeContext() {}

func NewPipelineNodeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PipelineNodeContext {
	var p = new(PipelineNodeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PipeLangRULE_pipelineNode

	return p
}

func (s *PipelineNodeContext) GetParser() antlr.Parser { return s.parser }

func (s *PipelineNodeContext) NODE_HEADER() antlr.TerminalNode {
	return s.GetToken(PipeLangNODE_HEADER, 0)
}

func (s *PipelineNodeContext) AllPRQL_LINE() []antlr.TerminalNode {
	return s.GetTokens(PipeLangPRQL_LINE)
}

func (s *PipelineNodeContext) PRQL_LINE(i int) antlr.TerminalNode {
	return s.GetToken(PipeLangPRQL_LINE, i)
}

func (s *PipelineNodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PipelineNodeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PipelineNodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeLangListener); ok {
		listenerT.EnterPipelineNode(s)
	}
}

func (s *PipelineNodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeLangListener); ok {
		listenerT.ExitPipelineNode(s)
	}
}

func (s *PipelineNodeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PipeLangVisitor:
		return t.VisitPipelineNode(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PipeLang) PipelineNode() (localctx IPipelineNodeContext) {
	localctx = NewPipelineNodeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, PipeLangRULE_pipelineNode)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(63)
		p.Match(PipeLangNODE_HEADER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(67)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == PipeLangPRQL_LINE {
		{
			p.SetState(64)
			p.Match(PipeLangPRQL_LINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(69)
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
