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
		"", "", "", "", "", "", "", "", "", "", "", "", "", "'\\u0000'",
	}
	staticData.SymbolicNames = []string{
		"", "COMMENT", "WS_BLANK", "WS", "KW_DESCRIPTION", "KW_TAGS", "KW_TYPE",
		"KW_NODE", "KW_COPY_SCHEDULE", "KW_DATASOURCE", "KW_TARGET_DATASOURCE",
		"KW_SQL_ARROW", "REST_OF_LINE", "SQL_LINE",
	}
	staticData.RuleNames = []string{
		"pipeFile", "statement", "directive", "nodeBlock", "sqlBlock", "sqlBody",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 13, 52, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 1, 0, 5, 0, 14, 8, 0, 10, 0, 12, 0, 17, 9, 0, 1, 0, 1, 0,
		1, 1, 1, 1, 3, 1, 23, 8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 37, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4,
		1, 4, 1, 4, 1, 5, 5, 5, 47, 8, 5, 10, 5, 12, 5, 50, 9, 5, 1, 5, 0, 0, 6,
		0, 2, 4, 6, 8, 10, 0, 0, 53, 0, 15, 1, 0, 0, 0, 2, 22, 1, 0, 0, 0, 4, 36,
		1, 0, 0, 0, 6, 38, 1, 0, 0, 0, 8, 42, 1, 0, 0, 0, 10, 48, 1, 0, 0, 0, 12,
		14, 3, 2, 1, 0, 13, 12, 1, 0, 0, 0, 14, 17, 1, 0, 0, 0, 15, 13, 1, 0, 0,
		0, 15, 16, 1, 0, 0, 0, 16, 18, 1, 0, 0, 0, 17, 15, 1, 0, 0, 0, 18, 19,
		5, 0, 0, 1, 19, 1, 1, 0, 0, 0, 20, 23, 3, 4, 2, 0, 21, 23, 3, 6, 3, 0,
		22, 20, 1, 0, 0, 0, 22, 21, 1, 0, 0, 0, 23, 3, 1, 0, 0, 0, 24, 25, 5, 4,
		0, 0, 25, 37, 5, 12, 0, 0, 26, 27, 5, 5, 0, 0, 27, 37, 5, 12, 0, 0, 28,
		29, 5, 6, 0, 0, 29, 37, 5, 12, 0, 0, 30, 31, 5, 9, 0, 0, 31, 37, 5, 12,
		0, 0, 32, 33, 5, 10, 0, 0, 33, 37, 5, 12, 0, 0, 34, 35, 5, 8, 0, 0, 35,
		37, 5, 12, 0, 0, 36, 24, 1, 0, 0, 0, 36, 26, 1, 0, 0, 0, 36, 28, 1, 0,
		0, 0, 36, 30, 1, 0, 0, 0, 36, 32, 1, 0, 0, 0, 36, 34, 1, 0, 0, 0, 37, 5,
		1, 0, 0, 0, 38, 39, 5, 7, 0, 0, 39, 40, 5, 12, 0, 0, 40, 41, 3, 8, 4, 0,
		41, 7, 1, 0, 0, 0, 42, 43, 5, 11, 0, 0, 43, 44, 3, 10, 5, 0, 44, 9, 1,
		0, 0, 0, 45, 47, 5, 13, 0, 0, 46, 45, 1, 0, 0, 0, 47, 50, 1, 0, 0, 0, 48,
		46, 1, 0, 0, 0, 48, 49, 1, 0, 0, 0, 49, 11, 1, 0, 0, 0, 50, 48, 1, 0, 0,
		0, 4, 15, 22, 36, 48,
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
	PipeLangEOF                  = antlr.TokenEOF
	PipeLangCOMMENT              = 1
	PipeLangWS_BLANK             = 2
	PipeLangWS                   = 3
	PipeLangKW_DESCRIPTION       = 4
	PipeLangKW_TAGS              = 5
	PipeLangKW_TYPE              = 6
	PipeLangKW_NODE              = 7
	PipeLangKW_COPY_SCHEDULE     = 8
	PipeLangKW_DATASOURCE        = 9
	PipeLangKW_TARGET_DATASOURCE = 10
	PipeLangKW_SQL_ARROW         = 11
	PipeLangREST_OF_LINE         = 12
	PipeLangSQL_LINE             = 13
)

// PipeLang rules.
const (
	PipeLangRULE_pipeFile  = 0
	PipeLangRULE_statement = 1
	PipeLangRULE_directive = 2
	PipeLangRULE_nodeBlock = 3
	PipeLangRULE_sqlBlock  = 4
	PipeLangRULE_sqlBody   = 5
)

// IPipeFileContext is an interface to support dynamic dispatch.
type IPipeFileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

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

func (s *PipeFileContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *PipeFileContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
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

	return t.(IStatementContext)
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
	p.SetState(15)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2032) != 0 {
		{
			p.SetState(12)
			p.Statement()
		}

		p.SetState(17)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(18)
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

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Directive() IDirectiveContext
	NodeBlock() INodeBlockContext

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PipeLangRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PipeLangRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PipeLangRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) Directive() IDirectiveContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDirectiveContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDirectiveContext)
}

func (s *StatementContext) NodeBlock() INodeBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INodeBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INodeBlockContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeLangListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeLangListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PipeLangVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PipeLang) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, PipeLangRULE_statement)
	p.SetState(22)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PipeLangKW_DESCRIPTION, PipeLangKW_TAGS, PipeLangKW_TYPE, PipeLangKW_COPY_SCHEDULE, PipeLangKW_DATASOURCE, PipeLangKW_TARGET_DATASOURCE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(20)
			p.Directive()
		}

	case PipeLangKW_NODE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(21)
			p.NodeBlock()
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

type TargetDatasourceDirContext struct {
	DirectiveContext
}

func NewTargetDatasourceDirContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TargetDatasourceDirContext {
	var p = new(TargetDatasourceDirContext)

	InitEmptyDirectiveContext(&p.DirectiveContext)
	p.parser = parser
	p.CopyAll(ctx.(*DirectiveContext))

	return p
}

func (s *TargetDatasourceDirContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TargetDatasourceDirContext) KW_TARGET_DATASOURCE() antlr.TerminalNode {
	return s.GetToken(PipeLangKW_TARGET_DATASOURCE, 0)
}

func (s *TargetDatasourceDirContext) REST_OF_LINE() antlr.TerminalNode {
	return s.GetToken(PipeLangREST_OF_LINE, 0)
}

func (s *TargetDatasourceDirContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeLangListener); ok {
		listenerT.EnterTargetDatasourceDir(s)
	}
}

func (s *TargetDatasourceDirContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeLangListener); ok {
		listenerT.ExitTargetDatasourceDir(s)
	}
}

func (s *TargetDatasourceDirContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PipeLangVisitor:
		return t.VisitTargetDatasourceDir(s)

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

func (s *TypeDirContext) KW_TYPE() antlr.TerminalNode {
	return s.GetToken(PipeLangKW_TYPE, 0)
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

type CopyScheduleDirContext struct {
	DirectiveContext
}

func NewCopyScheduleDirContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CopyScheduleDirContext {
	var p = new(CopyScheduleDirContext)

	InitEmptyDirectiveContext(&p.DirectiveContext)
	p.parser = parser
	p.CopyAll(ctx.(*DirectiveContext))

	return p
}

func (s *CopyScheduleDirContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CopyScheduleDirContext) KW_COPY_SCHEDULE() antlr.TerminalNode {
	return s.GetToken(PipeLangKW_COPY_SCHEDULE, 0)
}

func (s *CopyScheduleDirContext) REST_OF_LINE() antlr.TerminalNode {
	return s.GetToken(PipeLangREST_OF_LINE, 0)
}

func (s *CopyScheduleDirContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeLangListener); ok {
		listenerT.EnterCopyScheduleDir(s)
	}
}

func (s *CopyScheduleDirContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeLangListener); ok {
		listenerT.ExitCopyScheduleDir(s)
	}
}

func (s *CopyScheduleDirContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PipeLangVisitor:
		return t.VisitCopyScheduleDir(s)

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

func (s *DescriptionDirContext) KW_DESCRIPTION() antlr.TerminalNode {
	return s.GetToken(PipeLangKW_DESCRIPTION, 0)
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

func (s *TagsDirContext) KW_TAGS() antlr.TerminalNode {
	return s.GetToken(PipeLangKW_TAGS, 0)
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

type DatasourceDirContext struct {
	DirectiveContext
}

func NewDatasourceDirContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DatasourceDirContext {
	var p = new(DatasourceDirContext)

	InitEmptyDirectiveContext(&p.DirectiveContext)
	p.parser = parser
	p.CopyAll(ctx.(*DirectiveContext))

	return p
}

func (s *DatasourceDirContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DatasourceDirContext) KW_DATASOURCE() antlr.TerminalNode {
	return s.GetToken(PipeLangKW_DATASOURCE, 0)
}

func (s *DatasourceDirContext) REST_OF_LINE() antlr.TerminalNode {
	return s.GetToken(PipeLangREST_OF_LINE, 0)
}

func (s *DatasourceDirContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeLangListener); ok {
		listenerT.EnterDatasourceDir(s)
	}
}

func (s *DatasourceDirContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeLangListener); ok {
		listenerT.ExitDatasourceDir(s)
	}
}

func (s *DatasourceDirContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PipeLangVisitor:
		return t.VisitDatasourceDir(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PipeLang) Directive() (localctx IDirectiveContext) {
	localctx = NewDirectiveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, PipeLangRULE_directive)
	p.SetState(36)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PipeLangKW_DESCRIPTION:
		localctx = NewDescriptionDirContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(24)
			p.Match(PipeLangKW_DESCRIPTION)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(25)
			p.Match(PipeLangREST_OF_LINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PipeLangKW_TAGS:
		localctx = NewTagsDirContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(26)
			p.Match(PipeLangKW_TAGS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(27)
			p.Match(PipeLangREST_OF_LINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PipeLangKW_TYPE:
		localctx = NewTypeDirContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(28)
			p.Match(PipeLangKW_TYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(29)
			p.Match(PipeLangREST_OF_LINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PipeLangKW_DATASOURCE:
		localctx = NewDatasourceDirContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(30)
			p.Match(PipeLangKW_DATASOURCE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(31)
			p.Match(PipeLangREST_OF_LINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PipeLangKW_TARGET_DATASOURCE:
		localctx = NewTargetDatasourceDirContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(32)
			p.Match(PipeLangKW_TARGET_DATASOURCE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(33)
			p.Match(PipeLangREST_OF_LINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PipeLangKW_COPY_SCHEDULE:
		localctx = NewCopyScheduleDirContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(34)
			p.Match(PipeLangKW_COPY_SCHEDULE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(35)
			p.Match(PipeLangREST_OF_LINE)
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

// INodeBlockContext is an interface to support dynamic dispatch.
type INodeBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	KW_NODE() antlr.TerminalNode
	REST_OF_LINE() antlr.TerminalNode
	SqlBlock() ISqlBlockContext

	// IsNodeBlockContext differentiates from other interfaces.
	IsNodeBlockContext()
}

type NodeBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNodeBlockContext() *NodeBlockContext {
	var p = new(NodeBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PipeLangRULE_nodeBlock
	return p
}

func InitEmptyNodeBlockContext(p *NodeBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PipeLangRULE_nodeBlock
}

func (*NodeBlockContext) IsNodeBlockContext() {}

func NewNodeBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NodeBlockContext {
	var p = new(NodeBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PipeLangRULE_nodeBlock

	return p
}

func (s *NodeBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *NodeBlockContext) KW_NODE() antlr.TerminalNode {
	return s.GetToken(PipeLangKW_NODE, 0)
}

func (s *NodeBlockContext) REST_OF_LINE() antlr.TerminalNode {
	return s.GetToken(PipeLangREST_OF_LINE, 0)
}

func (s *NodeBlockContext) SqlBlock() ISqlBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISqlBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISqlBlockContext)
}

func (s *NodeBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NodeBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NodeBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeLangListener); ok {
		listenerT.EnterNodeBlock(s)
	}
}

func (s *NodeBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeLangListener); ok {
		listenerT.ExitNodeBlock(s)
	}
}

func (s *NodeBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PipeLangVisitor:
		return t.VisitNodeBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PipeLang) NodeBlock() (localctx INodeBlockContext) {
	localctx = NewNodeBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, PipeLangRULE_nodeBlock)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(38)
		p.Match(PipeLangKW_NODE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(39)
		p.Match(PipeLangREST_OF_LINE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(40)
		p.SqlBlock()
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

// ISqlBlockContext is an interface to support dynamic dispatch.
type ISqlBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	KW_SQL_ARROW() antlr.TerminalNode
	SqlBody() ISqlBodyContext

	// IsSqlBlockContext differentiates from other interfaces.
	IsSqlBlockContext()
}

type SqlBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySqlBlockContext() *SqlBlockContext {
	var p = new(SqlBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PipeLangRULE_sqlBlock
	return p
}

func InitEmptySqlBlockContext(p *SqlBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PipeLangRULE_sqlBlock
}

func (*SqlBlockContext) IsSqlBlockContext() {}

func NewSqlBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SqlBlockContext {
	var p = new(SqlBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PipeLangRULE_sqlBlock

	return p
}

func (s *SqlBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *SqlBlockContext) KW_SQL_ARROW() antlr.TerminalNode {
	return s.GetToken(PipeLangKW_SQL_ARROW, 0)
}

func (s *SqlBlockContext) SqlBody() ISqlBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISqlBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISqlBodyContext)
}

func (s *SqlBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SqlBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SqlBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeLangListener); ok {
		listenerT.EnterSqlBlock(s)
	}
}

func (s *SqlBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeLangListener); ok {
		listenerT.ExitSqlBlock(s)
	}
}

func (s *SqlBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PipeLangVisitor:
		return t.VisitSqlBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PipeLang) SqlBlock() (localctx ISqlBlockContext) {
	localctx = NewSqlBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, PipeLangRULE_sqlBlock)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(42)
		p.Match(PipeLangKW_SQL_ARROW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(43)
		p.SqlBody()
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

// ISqlBodyContext is an interface to support dynamic dispatch.
type ISqlBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllSQL_LINE() []antlr.TerminalNode
	SQL_LINE(i int) antlr.TerminalNode

	// IsSqlBodyContext differentiates from other interfaces.
	IsSqlBodyContext()
}

type SqlBodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySqlBodyContext() *SqlBodyContext {
	var p = new(SqlBodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PipeLangRULE_sqlBody
	return p
}

func InitEmptySqlBodyContext(p *SqlBodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PipeLangRULE_sqlBody
}

func (*SqlBodyContext) IsSqlBodyContext() {}

func NewSqlBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SqlBodyContext {
	var p = new(SqlBodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PipeLangRULE_sqlBody

	return p
}

func (s *SqlBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *SqlBodyContext) AllSQL_LINE() []antlr.TerminalNode {
	return s.GetTokens(PipeLangSQL_LINE)
}

func (s *SqlBodyContext) SQL_LINE(i int) antlr.TerminalNode {
	return s.GetToken(PipeLangSQL_LINE, i)
}

func (s *SqlBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SqlBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SqlBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeLangListener); ok {
		listenerT.EnterSqlBody(s)
	}
}

func (s *SqlBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PipeLangListener); ok {
		listenerT.ExitSqlBody(s)
	}
}

func (s *SqlBodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PipeLangVisitor:
		return t.VisitSqlBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PipeLang) SqlBody() (localctx ISqlBodyContext) {
	localctx = NewSqlBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, PipeLangRULE_sqlBody)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(48)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == PipeLangSQL_LINE {
		{
			p.SetState(45)
			p.Match(PipeLangSQL_LINE)
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
