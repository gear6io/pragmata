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
	staticData.LiteralNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"':'", "'{'", "'}'", "','", "'-'",
	}
	staticData.SymbolicNames = []string{
		"", "COMMENT", "WS_BLANK", "WS", "TYPE", "NAME", "TAGS", "OWNER", "DESTINATION",
		"SCHEDULE", "UNIQUE_KEY", "DESCRIPTION_ML", "DESCRIPTION", "SOURCES",
		"PARAMS", "PIPELINE", "IDENTIFIER", "COLON_TOK", "LBRACE", "RBRACE",
		"COMMA", "DASH", "STRING_LIT", "NUMBER", "INLINE_CMT", "VALUE", "VALUE_NL",
		"SECTION_LINE", "SECTION_BLANK", "SECTION_COMMENT",
	}
	staticData.RuleNames = []string{
		"pipeFile", "directive", "source", "param", "paramValue",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 29, 89, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 1, 0, 5, 0, 12, 8, 0, 10, 0, 12, 0, 15, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 27, 8, 1, 10, 1, 12, 1, 30, 9,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 5, 1, 44, 8, 1, 10, 1, 12, 1, 47, 9, 1, 1, 1, 1, 1, 5, 1, 51, 8, 1,
		10, 1, 12, 1, 54, 9, 1, 1, 1, 1, 1, 5, 1, 58, 8, 1, 10, 1, 12, 1, 61, 9,
		1, 3, 1, 63, 8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 71, 8, 2,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 83, 8,
		3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 0, 0, 5, 0, 2, 4, 6, 8, 0, 1, 2, 0, 16,
		16, 22, 23, 101, 0, 13, 1, 0, 0, 0, 2, 62, 1, 0, 0, 0, 4, 70, 1, 0, 0,
		0, 6, 72, 1, 0, 0, 0, 8, 86, 1, 0, 0, 0, 10, 12, 3, 2, 1, 0, 11, 10, 1,
		0, 0, 0, 12, 15, 1, 0, 0, 0, 13, 11, 1, 0, 0, 0, 13, 14, 1, 0, 0, 0, 14,
		16, 1, 0, 0, 0, 15, 13, 1, 0, 0, 0, 16, 17, 5, 0, 0, 1, 17, 1, 1, 0, 0,
		0, 18, 19, 5, 4, 0, 0, 19, 63, 5, 25, 0, 0, 20, 21, 5, 5, 0, 0, 21, 63,
		5, 25, 0, 0, 22, 23, 5, 12, 0, 0, 23, 63, 5, 25, 0, 0, 24, 28, 5, 11, 0,
		0, 25, 27, 5, 27, 0, 0, 26, 25, 1, 0, 0, 0, 27, 30, 1, 0, 0, 0, 28, 26,
		1, 0, 0, 0, 28, 29, 1, 0, 0, 0, 29, 63, 1, 0, 0, 0, 30, 28, 1, 0, 0, 0,
		31, 32, 5, 6, 0, 0, 32, 63, 5, 25, 0, 0, 33, 34, 5, 7, 0, 0, 34, 63, 5,
		25, 0, 0, 35, 36, 5, 8, 0, 0, 36, 63, 5, 25, 0, 0, 37, 38, 5, 9, 0, 0,
		38, 63, 5, 25, 0, 0, 39, 40, 5, 10, 0, 0, 40, 63, 5, 25, 0, 0, 41, 45,
		5, 13, 0, 0, 42, 44, 3, 4, 2, 0, 43, 42, 1, 0, 0, 0, 44, 47, 1, 0, 0, 0,
		45, 43, 1, 0, 0, 0, 45, 46, 1, 0, 0, 0, 46, 63, 1, 0, 0, 0, 47, 45, 1,
		0, 0, 0, 48, 52, 5, 14, 0, 0, 49, 51, 3, 6, 3, 0, 50, 49, 1, 0, 0, 0, 51,
		54, 1, 0, 0, 0, 52, 50, 1, 0, 0, 0, 52, 53, 1, 0, 0, 0, 53, 63, 1, 0, 0,
		0, 54, 52, 1, 0, 0, 0, 55, 59, 5, 15, 0, 0, 56, 58, 5, 27, 0, 0, 57, 56,
		1, 0, 0, 0, 58, 61, 1, 0, 0, 0, 59, 57, 1, 0, 0, 0, 59, 60, 1, 0, 0, 0,
		60, 63, 1, 0, 0, 0, 61, 59, 1, 0, 0, 0, 62, 18, 1, 0, 0, 0, 62, 20, 1,
		0, 0, 0, 62, 22, 1, 0, 0, 0, 62, 24, 1, 0, 0, 0, 62, 31, 1, 0, 0, 0, 62,
		33, 1, 0, 0, 0, 62, 35, 1, 0, 0, 0, 62, 37, 1, 0, 0, 0, 62, 39, 1, 0, 0,
		0, 62, 41, 1, 0, 0, 0, 62, 48, 1, 0, 0, 0, 62, 55, 1, 0, 0, 0, 63, 3, 1,
		0, 0, 0, 64, 65, 5, 21, 0, 0, 65, 66, 5, 16, 0, 0, 66, 67, 5, 17, 0, 0,
		67, 71, 5, 16, 0, 0, 68, 69, 5, 21, 0, 0, 69, 71, 5, 16, 0, 0, 70, 64,
		1, 0, 0, 0, 70, 68, 1, 0, 0, 0, 71, 5, 1, 0, 0, 0, 72, 73, 5, 16, 0, 0,
		73, 74, 5, 17, 0, 0, 74, 75, 5, 18, 0, 0, 75, 76, 5, 16, 0, 0, 76, 77,
		5, 17, 0, 0, 77, 82, 5, 16, 0, 0, 78, 79, 5, 20, 0, 0, 79, 80, 5, 16, 0,
		0, 80, 81, 5, 17, 0, 0, 81, 83, 3, 8, 4, 0, 82, 78, 1, 0, 0, 0, 82, 83,
		1, 0, 0, 0, 83, 84, 1, 0, 0, 0, 84, 85, 5, 19, 0, 0, 85, 7, 1, 0, 0, 0,
		86, 87, 7, 0, 0, 0, 87, 9, 1, 0, 0, 0, 8, 13, 28, 45, 52, 59, 62, 70, 82,
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
	PipeLangEOF             = antlr.TokenEOF
	PipeLangCOMMENT         = 1
	PipeLangWS_BLANK        = 2
	PipeLangWS              = 3
	PipeLangTYPE            = 4
	PipeLangNAME            = 5
	PipeLangTAGS            = 6
	PipeLangOWNER           = 7
	PipeLangDESTINATION     = 8
	PipeLangSCHEDULE        = 9
	PipeLangUNIQUE_KEY      = 10
	PipeLangDESCRIPTION_ML  = 11
	PipeLangDESCRIPTION     = 12
	PipeLangSOURCES         = 13
	PipeLangPARAMS          = 14
	PipeLangPIPELINE        = 15
	PipeLangIDENTIFIER      = 16
	PipeLangCOLON_TOK       = 17
	PipeLangLBRACE          = 18
	PipeLangRBRACE          = 19
	PipeLangCOMMA           = 20
	PipeLangDASH            = 21
	PipeLangSTRING_LIT      = 22
	PipeLangNUMBER          = 23
	PipeLangINLINE_CMT      = 24
	PipeLangVALUE           = 25
	PipeLangVALUE_NL        = 26
	PipeLangSECTION_LINE    = 27
	PipeLangSECTION_BLANK   = 28
	PipeLangSECTION_COMMENT = 29
)

// PipeLang rules.
const (
	PipeLangRULE_pipeFile   = 0
	PipeLangRULE_directive  = 1
	PipeLangRULE_source     = 2
	PipeLangRULE_param      = 3
	PipeLangRULE_paramValue = 4
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
	p.SetState(13)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&65520) != 0 {
		{
			p.SetState(10)
			p.Directive()
		}

		p.SetState(15)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(16)
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

func (s *NameDirContext) VALUE() antlr.TerminalNode {
	return s.GetToken(PipeLangVALUE, 0)
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

func (s *PipelineClauseContext) AllSECTION_LINE() []antlr.TerminalNode {
	return s.GetTokens(PipeLangSECTION_LINE)
}

func (s *PipelineClauseContext) SECTION_LINE(i int) antlr.TerminalNode {
	return s.GetToken(PipeLangSECTION_LINE, i)
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

type UniqueKeyDirContext struct {
	DirectiveContext
}

func NewUniqueKeyDirContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UniqueKeyDirContext {
	var p = new(UniqueKeyDirContext)

	InitEmptyDirectiveContext(&p.DirectiveContext)
	p.parser = parser
	p.CopyAll(ctx.(*DirectiveContext))

	return p
}

func (s *UniqueKeyDirContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UniqueKeyDirContext) UNIQUE_KEY() antlr.TerminalNode {
	return s.GetToken(PipeLangUNIQUE_KEY, 0)
}

func (s *UniqueKeyDirContext) VALUE() antlr.TerminalNode {
	return s.GetToken(PipeLangVALUE, 0)
}

func (s *UniqueKeyDirContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeLangListener); ok {
		listenerT.EnterUniqueKeyDir(s)
	}
}

func (s *UniqueKeyDirContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeLangListener); ok {
		listenerT.ExitUniqueKeyDir(s)
	}
}

func (s *UniqueKeyDirContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PipeLangVisitor:
		return t.VisitUniqueKeyDir(s)

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

func (s *TypeDirContext) VALUE() antlr.TerminalNode {
	return s.GetToken(PipeLangVALUE, 0)
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

func (s *DestinationDirContext) VALUE() antlr.TerminalNode {
	return s.GetToken(PipeLangVALUE, 0)
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

func (s *ParamsClauseContext) AllParam() []IParamContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParamContext); ok {
			len++
		}
	}

	tst := make([]IParamContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParamContext); ok {
			tst[i] = t.(IParamContext)
			i++
		}
	}

	return tst
}

func (s *ParamsClauseContext) Param(i int) IParamContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamContext); ok {
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

	return t.(IParamContext)
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

func (s *DescriptionMLDirContext) AllSECTION_LINE() []antlr.TerminalNode {
	return s.GetTokens(PipeLangSECTION_LINE)
}

func (s *DescriptionMLDirContext) SECTION_LINE(i int) antlr.TerminalNode {
	return s.GetToken(PipeLangSECTION_LINE, i)
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

func (s *DescriptionDirContext) VALUE() antlr.TerminalNode {
	return s.GetToken(PipeLangVALUE, 0)
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

func (s *OwnerDirContext) VALUE() antlr.TerminalNode {
	return s.GetToken(PipeLangVALUE, 0)
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

func (s *ScheduleDirContext) VALUE() antlr.TerminalNode {
	return s.GetToken(PipeLangVALUE, 0)
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

func (s *SourcesClauseContext) AllSource() []ISourceContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISourceContext); ok {
			len++
		}
	}

	tst := make([]ISourceContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISourceContext); ok {
			tst[i] = t.(ISourceContext)
			i++
		}
	}

	return tst
}

func (s *SourcesClauseContext) Source(i int) ISourceContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISourceContext); ok {
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

	return t.(ISourceContext)
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

func (s *TagsDirContext) VALUE() antlr.TerminalNode {
	return s.GetToken(PipeLangVALUE, 0)
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

	p.SetState(62)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PipeLangTYPE:
		localctx = NewTypeDirContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(18)
			p.Match(PipeLangTYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(19)
			p.Match(PipeLangVALUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PipeLangNAME:
		localctx = NewNameDirContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(20)
			p.Match(PipeLangNAME)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(21)
			p.Match(PipeLangVALUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PipeLangDESCRIPTION:
		localctx = NewDescriptionDirContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(22)
			p.Match(PipeLangDESCRIPTION)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(23)
			p.Match(PipeLangVALUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PipeLangDESCRIPTION_ML:
		localctx = NewDescriptionMLDirContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(24)
			p.Match(PipeLangDESCRIPTION_ML)
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

		for _la == PipeLangSECTION_LINE {
			{
				p.SetState(25)
				p.Match(PipeLangSECTION_LINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(30)
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
			p.SetState(31)
			p.Match(PipeLangTAGS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(32)
			p.Match(PipeLangVALUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PipeLangOWNER:
		localctx = NewOwnerDirContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(33)
			p.Match(PipeLangOWNER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(34)
			p.Match(PipeLangVALUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PipeLangDESTINATION:
		localctx = NewDestinationDirContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(35)
			p.Match(PipeLangDESTINATION)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(36)
			p.Match(PipeLangVALUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PipeLangSCHEDULE:
		localctx = NewScheduleDirContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(37)
			p.Match(PipeLangSCHEDULE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(38)
			p.Match(PipeLangVALUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PipeLangUNIQUE_KEY:
		localctx = NewUniqueKeyDirContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(39)
			p.Match(PipeLangUNIQUE_KEY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(40)
			p.Match(PipeLangVALUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PipeLangSOURCES:
		localctx = NewSourcesClauseContext(p, localctx)
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(41)
			p.Match(PipeLangSOURCES)
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

		for _la == PipeLangDASH {
			{
				p.SetState(42)
				p.Source()
			}

			p.SetState(47)
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
			p.SetState(48)
			p.Match(PipeLangPARAMS)
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

		for _la == PipeLangIDENTIFIER {
			{
				p.SetState(49)
				p.Param()
			}

			p.SetState(54)
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
			p.SetState(55)
			p.Match(PipeLangPIPELINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(59)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == PipeLangSECTION_LINE {
			{
				p.SetState(56)
				p.Match(PipeLangSECTION_LINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(61)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
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

// ISourceContext is an interface to support dynamic dispatch.
type ISourceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsSourceContext differentiates from other interfaces.
	IsSourceContext()
}

type SourceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySourceContext() *SourceContext {
	var p = new(SourceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PipeLangRULE_source
	return p
}

func InitEmptySourceContext(p *SourceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PipeLangRULE_source
}

func (*SourceContext) IsSourceContext() {}

func NewSourceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SourceContext {
	var p = new(SourceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PipeLangRULE_source

	return p
}

func (s *SourceContext) GetParser() antlr.Parser { return s.parser }

func (s *SourceContext) CopyAll(ctx *SourceContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *SourceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SourceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SimpleSourceContext struct {
	SourceContext
	name antlr.Token
}

func NewSimpleSourceContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SimpleSourceContext {
	var p = new(SimpleSourceContext)

	InitEmptySourceContext(&p.SourceContext)
	p.parser = parser
	p.CopyAll(ctx.(*SourceContext))

	return p
}

func (s *SimpleSourceContext) GetName() antlr.Token { return s.name }

func (s *SimpleSourceContext) SetName(v antlr.Token) { s.name = v }

func (s *SimpleSourceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SimpleSourceContext) DASH() antlr.TerminalNode {
	return s.GetToken(PipeLangDASH, 0)
}

func (s *SimpleSourceContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(PipeLangIDENTIFIER, 0)
}

func (s *SimpleSourceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeLangListener); ok {
		listenerT.EnterSimpleSource(s)
	}
}

func (s *SimpleSourceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeLangListener); ok {
		listenerT.ExitSimpleSource(s)
	}
}

func (s *SimpleSourceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PipeLangVisitor:
		return t.VisitSimpleSource(s)

	default:
		return t.VisitChildren(s)
	}
}

type AliasedSourceContext struct {
	SourceContext
	alias antlr.Token
	table antlr.Token
}

func NewAliasedSourceContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AliasedSourceContext {
	var p = new(AliasedSourceContext)

	InitEmptySourceContext(&p.SourceContext)
	p.parser = parser
	p.CopyAll(ctx.(*SourceContext))

	return p
}

func (s *AliasedSourceContext) GetAlias() antlr.Token { return s.alias }

func (s *AliasedSourceContext) GetTable() antlr.Token { return s.table }

func (s *AliasedSourceContext) SetAlias(v antlr.Token) { s.alias = v }

func (s *AliasedSourceContext) SetTable(v antlr.Token) { s.table = v }

func (s *AliasedSourceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AliasedSourceContext) DASH() antlr.TerminalNode {
	return s.GetToken(PipeLangDASH, 0)
}

func (s *AliasedSourceContext) COLON_TOK() antlr.TerminalNode {
	return s.GetToken(PipeLangCOLON_TOK, 0)
}

func (s *AliasedSourceContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(PipeLangIDENTIFIER)
}

func (s *AliasedSourceContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(PipeLangIDENTIFIER, i)
}

func (s *AliasedSourceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeLangListener); ok {
		listenerT.EnterAliasedSource(s)
	}
}

func (s *AliasedSourceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeLangListener); ok {
		listenerT.ExitAliasedSource(s)
	}
}

func (s *AliasedSourceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PipeLangVisitor:
		return t.VisitAliasedSource(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PipeLang) Source() (localctx ISourceContext) {
	localctx = NewSourceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, PipeLangRULE_source)
	p.SetState(70)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		localctx = NewAliasedSourceContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(64)
			p.Match(PipeLangDASH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(65)

			var _m = p.Match(PipeLangIDENTIFIER)

			localctx.(*AliasedSourceContext).alias = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(66)
			p.Match(PipeLangCOLON_TOK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(67)

			var _m = p.Match(PipeLangIDENTIFIER)

			localctx.(*AliasedSourceContext).table = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewSimpleSourceContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(68)
			p.Match(PipeLangDASH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(69)

			var _m = p.Match(PipeLangIDENTIFIER)

			localctx.(*SimpleSourceContext).name = _m
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

// IParamContext is an interface to support dynamic dispatch.
type IParamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetPname returns the pname token.
	GetPname() antlr.Token

	// GetDtype returns the dtype token.
	GetDtype() antlr.Token

	// SetPname sets the pname token.
	SetPname(antlr.Token)

	// SetDtype sets the dtype token.
	SetDtype(antlr.Token)

	// Getter signatures
	AllCOLON_TOK() []antlr.TerminalNode
	COLON_TOK(i int) antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	AllIDENTIFIER() []antlr.TerminalNode
	IDENTIFIER(i int) antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	COMMA() antlr.TerminalNode
	ParamValue() IParamValueContext

	// IsParamContext differentiates from other interfaces.
	IsParamContext()
}

type ParamContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	pname  antlr.Token
	dtype  antlr.Token
}

func NewEmptyParamContext() *ParamContext {
	var p = new(ParamContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PipeLangRULE_param
	return p
}

func InitEmptyParamContext(p *ParamContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PipeLangRULE_param
}

func (*ParamContext) IsParamContext() {}

func NewParamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamContext {
	var p = new(ParamContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PipeLangRULE_param

	return p
}

func (s *ParamContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamContext) GetPname() antlr.Token { return s.pname }

func (s *ParamContext) GetDtype() antlr.Token { return s.dtype }

func (s *ParamContext) SetPname(v antlr.Token) { s.pname = v }

func (s *ParamContext) SetDtype(v antlr.Token) { s.dtype = v }

func (s *ParamContext) AllCOLON_TOK() []antlr.TerminalNode {
	return s.GetTokens(PipeLangCOLON_TOK)
}

func (s *ParamContext) COLON_TOK(i int) antlr.TerminalNode {
	return s.GetToken(PipeLangCOLON_TOK, i)
}

func (s *ParamContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(PipeLangLBRACE, 0)
}

func (s *ParamContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(PipeLangIDENTIFIER)
}

func (s *ParamContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(PipeLangIDENTIFIER, i)
}

func (s *ParamContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(PipeLangRBRACE, 0)
}

func (s *ParamContext) COMMA() antlr.TerminalNode {
	return s.GetToken(PipeLangCOMMA, 0)
}

func (s *ParamContext) ParamValue() IParamValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamValueContext)
}

func (s *ParamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeLangListener); ok {
		listenerT.EnterParam(s)
	}
}

func (s *ParamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeLangListener); ok {
		listenerT.ExitParam(s)
	}
}

func (s *ParamContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PipeLangVisitor:
		return t.VisitParam(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PipeLang) Param() (localctx IParamContext) {
	localctx = NewParamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, PipeLangRULE_param)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(72)

		var _m = p.Match(PipeLangIDENTIFIER)

		localctx.(*ParamContext).pname = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(73)
		p.Match(PipeLangCOLON_TOK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(74)
		p.Match(PipeLangLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(75)
		p.Match(PipeLangIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(76)
		p.Match(PipeLangCOLON_TOK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(77)

		var _m = p.Match(PipeLangIDENTIFIER)

		localctx.(*ParamContext).dtype = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(82)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PipeLangCOMMA {
		{
			p.SetState(78)
			p.Match(PipeLangCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(79)
			p.Match(PipeLangIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(80)
			p.Match(PipeLangCOLON_TOK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(81)
			p.ParamValue()
		}

	}
	{
		p.SetState(84)
		p.Match(PipeLangRBRACE)
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

// IParamValueContext is an interface to support dynamic dispatch.
type IParamValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING_LIT() antlr.TerminalNode
	NUMBER() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode

	// IsParamValueContext differentiates from other interfaces.
	IsParamValueContext()
}

type ParamValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamValueContext() *ParamValueContext {
	var p = new(ParamValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PipeLangRULE_paramValue
	return p
}

func InitEmptyParamValueContext(p *ParamValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PipeLangRULE_paramValue
}

func (*ParamValueContext) IsParamValueContext() {}

func NewParamValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamValueContext {
	var p = new(ParamValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PipeLangRULE_paramValue

	return p
}

func (s *ParamValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamValueContext) STRING_LIT() antlr.TerminalNode {
	return s.GetToken(PipeLangSTRING_LIT, 0)
}

func (s *ParamValueContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(PipeLangNUMBER, 0)
}

func (s *ParamValueContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(PipeLangIDENTIFIER, 0)
}

func (s *ParamValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeLangListener); ok {
		listenerT.EnterParamValue(s)
	}
}

func (s *ParamValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeLangListener); ok {
		listenerT.ExitParamValue(s)
	}
}

func (s *ParamValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PipeLangVisitor:
		return t.VisitParamValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PipeLang) ParamValue() (localctx IParamValueContext) {
	localctx = NewParamValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, PipeLangRULE_paramValue)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(86)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&12648448) != 0) {
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
