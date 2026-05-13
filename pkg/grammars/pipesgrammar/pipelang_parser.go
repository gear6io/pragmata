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
		"SCHEDULE", "DESCRIPTION_ML", "DESCRIPTION", "SOURCES", "PARAMS", "PIPELINE",
		"REST_OF_LINE", "BLOCK_LINE", "BLOCK_BLANK", "BLOCK_COMMENT", "NODE_HEADER",
		"PRQL_LINE", "PIPE_BLANK", "PIPE_COMMENT",
	}
	staticData.RuleNames = []string{
		"pipeFile", "directive", "pipelineBlock", "pipelineNode",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 22, 69, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 1, 0, 5,
		0, 10, 8, 0, 10, 0, 12, 0, 13, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 25, 8, 1, 10, 1, 12, 1, 28, 9, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 40, 8, 1, 10,
		1, 12, 1, 43, 9, 1, 1, 1, 1, 1, 5, 1, 47, 8, 1, 10, 1, 12, 1, 50, 9, 1,
		1, 1, 1, 1, 3, 1, 54, 8, 1, 1, 2, 5, 2, 57, 8, 2, 10, 2, 12, 2, 60, 9,
		2, 1, 3, 1, 3, 5, 3, 64, 8, 3, 10, 3, 12, 3, 67, 9, 3, 1, 3, 0, 0, 4, 0,
		2, 4, 6, 0, 0, 80, 0, 11, 1, 0, 0, 0, 2, 53, 1, 0, 0, 0, 4, 58, 1, 0, 0,
		0, 6, 61, 1, 0, 0, 0, 8, 10, 3, 2, 1, 0, 9, 8, 1, 0, 0, 0, 10, 13, 1, 0,
		0, 0, 11, 9, 1, 0, 0, 0, 11, 12, 1, 0, 0, 0, 12, 14, 1, 0, 0, 0, 13, 11,
		1, 0, 0, 0, 14, 15, 5, 0, 0, 1, 15, 1, 1, 0, 0, 0, 16, 17, 5, 4, 0, 0,
		17, 54, 5, 15, 0, 0, 18, 19, 5, 5, 0, 0, 19, 54, 5, 15, 0, 0, 20, 21, 5,
		11, 0, 0, 21, 54, 5, 15, 0, 0, 22, 26, 5, 10, 0, 0, 23, 25, 5, 16, 0, 0,
		24, 23, 1, 0, 0, 0, 25, 28, 1, 0, 0, 0, 26, 24, 1, 0, 0, 0, 26, 27, 1,
		0, 0, 0, 27, 54, 1, 0, 0, 0, 28, 26, 1, 0, 0, 0, 29, 30, 5, 6, 0, 0, 30,
		54, 5, 15, 0, 0, 31, 32, 5, 7, 0, 0, 32, 54, 5, 15, 0, 0, 33, 34, 5, 8,
		0, 0, 34, 54, 5, 15, 0, 0, 35, 36, 5, 9, 0, 0, 36, 54, 5, 15, 0, 0, 37,
		41, 5, 12, 0, 0, 38, 40, 5, 16, 0, 0, 39, 38, 1, 0, 0, 0, 40, 43, 1, 0,
		0, 0, 41, 39, 1, 0, 0, 0, 41, 42, 1, 0, 0, 0, 42, 54, 1, 0, 0, 0, 43, 41,
		1, 0, 0, 0, 44, 48, 5, 13, 0, 0, 45, 47, 5, 16, 0, 0, 46, 45, 1, 0, 0,
		0, 47, 50, 1, 0, 0, 0, 48, 46, 1, 0, 0, 0, 48, 49, 1, 0, 0, 0, 49, 54,
		1, 0, 0, 0, 50, 48, 1, 0, 0, 0, 51, 52, 5, 14, 0, 0, 52, 54, 3, 4, 2, 0,
		53, 16, 1, 0, 0, 0, 53, 18, 1, 0, 0, 0, 53, 20, 1, 0, 0, 0, 53, 22, 1,
		0, 0, 0, 53, 29, 1, 0, 0, 0, 53, 31, 1, 0, 0, 0, 53, 33, 1, 0, 0, 0, 53,
		35, 1, 0, 0, 0, 53, 37, 1, 0, 0, 0, 53, 44, 1, 0, 0, 0, 53, 51, 1, 0, 0,
		0, 54, 3, 1, 0, 0, 0, 55, 57, 3, 6, 3, 0, 56, 55, 1, 0, 0, 0, 57, 60, 1,
		0, 0, 0, 58, 56, 1, 0, 0, 0, 58, 59, 1, 0, 0, 0, 59, 5, 1, 0, 0, 0, 60,
		58, 1, 0, 0, 0, 61, 65, 5, 19, 0, 0, 62, 64, 5, 20, 0, 0, 63, 62, 1, 0,
		0, 0, 64, 67, 1, 0, 0, 0, 65, 63, 1, 0, 0, 0, 65, 66, 1, 0, 0, 0, 66, 7,
		1, 0, 0, 0, 67, 65, 1, 0, 0, 0, 7, 11, 26, 41, 48, 53, 58, 65,
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
	PipeLangDESCRIPTION_ML = 10
	PipeLangDESCRIPTION    = 11
	PipeLangSOURCES        = 12
	PipeLangPARAMS         = 13
	PipeLangPIPELINE       = 14
	PipeLangREST_OF_LINE   = 15
	PipeLangBLOCK_LINE     = 16
	PipeLangBLOCK_BLANK    = 17
	PipeLangBLOCK_COMMENT  = 18
	PipeLangNODE_HEADER    = 19
	PipeLangPRQL_LINE      = 20
	PipeLangPIPE_BLANK     = 21
	PipeLangPIPE_COMMENT   = 22
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

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&32752) != 0 {
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

type NameDirContext struct {
	DirectiveContext
}

func NewNameDirContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NameDirContext {
	var p = new(NameDirContext)

	InitEmptyDirectiveContext(&p.DirectiveContext)
	p.parser = parser
	p.CopyAll(ctx.(*DirectiveContext))

	return p
}

func (s *NameDirContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NameDirContext) NAME() antlr.TerminalNode {
	return s.GetToken(PipeLangNAME, 0)
}

func (s *NameDirContext) REST_OF_LINE() antlr.TerminalNode {
	return s.GetToken(PipeLangREST_OF_LINE, 0)
}

func (s *NameDirContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeLangListener); ok {
		listenerT.EnterNameDir(s)
	}
}

func (s *NameDirContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeLangListener); ok {
		listenerT.ExitNameDir(s)
	}
}

func (s *NameDirContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PipeLangVisitor:
		return t.VisitNameDir(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParamsDirContext struct {
	DirectiveContext
}

func NewParamsDirContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParamsDirContext {
	var p = new(ParamsDirContext)

	InitEmptyDirectiveContext(&p.DirectiveContext)
	p.parser = parser
	p.CopyAll(ctx.(*DirectiveContext))

	return p
}

func (s *ParamsDirContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamsDirContext) PARAMS() antlr.TerminalNode {
	return s.GetToken(PipeLangPARAMS, 0)
}

func (s *ParamsDirContext) AllBLOCK_LINE() []antlr.TerminalNode {
	return s.GetTokens(PipeLangBLOCK_LINE)
}

func (s *ParamsDirContext) BLOCK_LINE(i int) antlr.TerminalNode {
	return s.GetToken(PipeLangBLOCK_LINE, i)
}

func (s *ParamsDirContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeLangListener); ok {
		listenerT.EnterParamsDir(s)
	}
}

func (s *ParamsDirContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeLangListener); ok {
		listenerT.ExitParamsDir(s)
	}
}

func (s *ParamsDirContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PipeLangVisitor:
		return t.VisitParamsDir(s)

	default:
		return t.VisitChildren(s)
	}
}

type PipelineDirContext struct {
	DirectiveContext
}

func NewPipelineDirContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PipelineDirContext {
	var p = new(PipelineDirContext)

	InitEmptyDirectiveContext(&p.DirectiveContext)
	p.parser = parser
	p.CopyAll(ctx.(*DirectiveContext))

	return p
}

func (s *PipelineDirContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PipelineDirContext) PIPELINE() antlr.TerminalNode {
	return s.GetToken(PipeLangPIPELINE, 0)
}

func (s *PipelineDirContext) PipelineBlock() IPipelineBlockContext {
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

func (s *PipelineDirContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeLangListener); ok {
		listenerT.EnterPipelineDir(s)
	}
}

func (s *PipelineDirContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeLangListener); ok {
		listenerT.ExitPipelineDir(s)
	}
}

func (s *PipelineDirContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PipeLangVisitor:
		return t.VisitPipelineDir(s)

	default:
		return t.VisitChildren(s)
	}
}

type TypeDirContext struct {
	DirectiveContext
}

func NewTypeDirContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeDirContext {
	var p = new(TypeDirContext)

	InitEmptyDirectiveContext(&p.DirectiveContext)
	p.parser = parser
	p.CopyAll(ctx.(*DirectiveContext))

	return p
}

func (s *TypeDirContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeDirContext) TYPE() antlr.TerminalNode {
	return s.GetToken(PipeLangTYPE, 0)
}

func (s *TypeDirContext) REST_OF_LINE() antlr.TerminalNode {
	return s.GetToken(PipeLangREST_OF_LINE, 0)
}

func (s *TypeDirContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeLangListener); ok {
		listenerT.EnterTypeDir(s)
	}
}

func (s *TypeDirContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeLangListener); ok {
		listenerT.ExitTypeDir(s)
	}
}

func (s *TypeDirContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PipeLangVisitor:
		return t.VisitTypeDir(s)

	default:
		return t.VisitChildren(s)
	}
}

type DestinationDirContext struct {
	DirectiveContext
}

func NewDestinationDirContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DestinationDirContext {
	var p = new(DestinationDirContext)

	InitEmptyDirectiveContext(&p.DirectiveContext)
	p.parser = parser
	p.CopyAll(ctx.(*DirectiveContext))

	return p
}

func (s *DestinationDirContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DestinationDirContext) DESTINATION() antlr.TerminalNode {
	return s.GetToken(PipeLangDESTINATION, 0)
}

func (s *DestinationDirContext) REST_OF_LINE() antlr.TerminalNode {
	return s.GetToken(PipeLangREST_OF_LINE, 0)
}

func (s *DestinationDirContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeLangListener); ok {
		listenerT.EnterDestinationDir(s)
	}
}

func (s *DestinationDirContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeLangListener); ok {
		listenerT.ExitDestinationDir(s)
	}
}

func (s *DestinationDirContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PipeLangVisitor:
		return t.VisitDestinationDir(s)

	default:
		return t.VisitChildren(s)
	}
}

type DescriptionMLDirContext struct {
	DirectiveContext
}

func NewDescriptionMLDirContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DescriptionMLDirContext {
	var p = new(DescriptionMLDirContext)

	InitEmptyDirectiveContext(&p.DirectiveContext)
	p.parser = parser
	p.CopyAll(ctx.(*DirectiveContext))

	return p
}

func (s *DescriptionMLDirContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DescriptionMLDirContext) DESCRIPTION_ML() antlr.TerminalNode {
	return s.GetToken(PipeLangDESCRIPTION_ML, 0)
}

func (s *DescriptionMLDirContext) AllBLOCK_LINE() []antlr.TerminalNode {
	return s.GetTokens(PipeLangBLOCK_LINE)
}

func (s *DescriptionMLDirContext) BLOCK_LINE(i int) antlr.TerminalNode {
	return s.GetToken(PipeLangBLOCK_LINE, i)
}

func (s *DescriptionMLDirContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeLangListener); ok {
		listenerT.EnterDescriptionMLDir(s)
	}
}

func (s *DescriptionMLDirContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeLangListener); ok {
		listenerT.ExitDescriptionMLDir(s)
	}
}

func (s *DescriptionMLDirContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PipeLangVisitor:
		return t.VisitDescriptionMLDir(s)

	default:
		return t.VisitChildren(s)
	}
}

type DescriptionDirContext struct {
	DirectiveContext
}

func NewDescriptionDirContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DescriptionDirContext {
	var p = new(DescriptionDirContext)

	InitEmptyDirectiveContext(&p.DirectiveContext)
	p.parser = parser
	p.CopyAll(ctx.(*DirectiveContext))

	return p
}

func (s *DescriptionDirContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DescriptionDirContext) DESCRIPTION() antlr.TerminalNode {
	return s.GetToken(PipeLangDESCRIPTION, 0)
}

func (s *DescriptionDirContext) REST_OF_LINE() antlr.TerminalNode {
	return s.GetToken(PipeLangREST_OF_LINE, 0)
}

func (s *DescriptionDirContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeLangListener); ok {
		listenerT.EnterDescriptionDir(s)
	}
}

func (s *DescriptionDirContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeLangListener); ok {
		listenerT.ExitDescriptionDir(s)
	}
}

func (s *DescriptionDirContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PipeLangVisitor:
		return t.VisitDescriptionDir(s)

	default:
		return t.VisitChildren(s)
	}
}

type OwnerDirContext struct {
	DirectiveContext
}

func NewOwnerDirContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OwnerDirContext {
	var p = new(OwnerDirContext)

	InitEmptyDirectiveContext(&p.DirectiveContext)
	p.parser = parser
	p.CopyAll(ctx.(*DirectiveContext))

	return p
}

func (s *OwnerDirContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OwnerDirContext) OWNER() antlr.TerminalNode {
	return s.GetToken(PipeLangOWNER, 0)
}

func (s *OwnerDirContext) REST_OF_LINE() antlr.TerminalNode {
	return s.GetToken(PipeLangREST_OF_LINE, 0)
}

func (s *OwnerDirContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeLangListener); ok {
		listenerT.EnterOwnerDir(s)
	}
}

func (s *OwnerDirContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeLangListener); ok {
		listenerT.ExitOwnerDir(s)
	}
}

func (s *OwnerDirContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PipeLangVisitor:
		return t.VisitOwnerDir(s)

	default:
		return t.VisitChildren(s)
	}
}

type ScheduleDirContext struct {
	DirectiveContext
}

func NewScheduleDirContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ScheduleDirContext {
	var p = new(ScheduleDirContext)

	InitEmptyDirectiveContext(&p.DirectiveContext)
	p.parser = parser
	p.CopyAll(ctx.(*DirectiveContext))

	return p
}

func (s *ScheduleDirContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ScheduleDirContext) SCHEDULE() antlr.TerminalNode {
	return s.GetToken(PipeLangSCHEDULE, 0)
}

func (s *ScheduleDirContext) REST_OF_LINE() antlr.TerminalNode {
	return s.GetToken(PipeLangREST_OF_LINE, 0)
}

func (s *ScheduleDirContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeLangListener); ok {
		listenerT.EnterScheduleDir(s)
	}
}

func (s *ScheduleDirContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeLangListener); ok {
		listenerT.ExitScheduleDir(s)
	}
}

func (s *ScheduleDirContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PipeLangVisitor:
		return t.VisitScheduleDir(s)

	default:
		return t.VisitChildren(s)
	}
}

type SourcesDirContext struct {
	DirectiveContext
}

func NewSourcesDirContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SourcesDirContext {
	var p = new(SourcesDirContext)

	InitEmptyDirectiveContext(&p.DirectiveContext)
	p.parser = parser
	p.CopyAll(ctx.(*DirectiveContext))

	return p
}

func (s *SourcesDirContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SourcesDirContext) SOURCES() antlr.TerminalNode {
	return s.GetToken(PipeLangSOURCES, 0)
}

func (s *SourcesDirContext) AllBLOCK_LINE() []antlr.TerminalNode {
	return s.GetTokens(PipeLangBLOCK_LINE)
}

func (s *SourcesDirContext) BLOCK_LINE(i int) antlr.TerminalNode {
	return s.GetToken(PipeLangBLOCK_LINE, i)
}

func (s *SourcesDirContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeLangListener); ok {
		listenerT.EnterSourcesDir(s)
	}
}

func (s *SourcesDirContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeLangListener); ok {
		listenerT.ExitSourcesDir(s)
	}
}

func (s *SourcesDirContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PipeLangVisitor:
		return t.VisitSourcesDir(s)

	default:
		return t.VisitChildren(s)
	}
}

type TagsDirContext struct {
	DirectiveContext
}

func NewTagsDirContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TagsDirContext {
	var p = new(TagsDirContext)

	InitEmptyDirectiveContext(&p.DirectiveContext)
	p.parser = parser
	p.CopyAll(ctx.(*DirectiveContext))

	return p
}

func (s *TagsDirContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TagsDirContext) TAGS() antlr.TerminalNode {
	return s.GetToken(PipeLangTAGS, 0)
}

func (s *TagsDirContext) REST_OF_LINE() antlr.TerminalNode {
	return s.GetToken(PipeLangREST_OF_LINE, 0)
}

func (s *TagsDirContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeLangListener); ok {
		listenerT.EnterTagsDir(s)
	}
}

func (s *TagsDirContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeLangListener); ok {
		listenerT.ExitTagsDir(s)
	}
}

func (s *TagsDirContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PipeLangVisitor:
		return t.VisitTagsDir(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PipeLang) Directive() (localctx IDirectiveContext) {
	localctx = NewDirectiveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, PipeLangRULE_directive)
	var _la int

	p.SetState(53)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PipeLangTYPE:
		localctx = NewTypeDirContext(p, localctx)
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
		localctx = NewNameDirContext(p, localctx)
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
		localctx = NewDescriptionDirContext(p, localctx)
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
		localctx = NewDescriptionMLDirContext(p, localctx)
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
		localctx = NewTagsDirContext(p, localctx)
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
		localctx = NewOwnerDirContext(p, localctx)
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
		localctx = NewDestinationDirContext(p, localctx)
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
		localctx = NewScheduleDirContext(p, localctx)
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

	case PipeLangSOURCES:
		localctx = NewSourcesDirContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(37)
			p.Match(PipeLangSOURCES)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(41)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == PipeLangBLOCK_LINE {
			{
				p.SetState(38)
				p.Match(PipeLangBLOCK_LINE)
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
		}

	case PipeLangPARAMS:
		localctx = NewParamsDirContext(p, localctx)
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(44)
			p.Match(PipeLangPARAMS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(48)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == PipeLangBLOCK_LINE {
			{
				p.SetState(45)
				p.Match(PipeLangBLOCK_LINE)
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
		}

	case PipeLangPIPELINE:
		localctx = NewPipelineDirContext(p, localctx)
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(51)
			p.Match(PipeLangPIPELINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(52)
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
	p.SetState(58)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == PipeLangNODE_HEADER {
		{
			p.SetState(55)
			p.PipelineNode()
		}

		p.SetState(60)
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
		p.SetState(61)
		p.Match(PipeLangNODE_HEADER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(65)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == PipeLangPRQL_LINE {
		{
			p.SetState(62)
			p.Match(PipeLangPRQL_LINE)
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
