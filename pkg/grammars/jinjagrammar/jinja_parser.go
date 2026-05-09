// Code generated from Jinja.g4 by ANTLR 4.13.2. DO NOT EDIT.

package jinjagrammar // Jinja
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

type Jinja struct {
	*antlr.BaseParser
}

var JinjaParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func jinjaParserInit() {
	staticData := &JinjaParserStaticData
	staticData.LiteralNames = []string{
		"", "'{{'", "'{%'", "'{#'", "", "'}}'", "", "", "", "", "", "", "'.'",
		"'|'", "'=='", "'!='", "", "", "'%}'", "'if'", "'elif'", "'else'", "'endif'",
		"'for'", "'in'", "'endfor'", "'macro'", "'endmacro'", "'set'", "", "",
		"", "", "", "", "", "'#}'",
	}
	staticData.SymbolicNames = []string{
		"", "EXPR_START", "BLOCK_START", "COMMENT_START", "TEXT", "EXPR_END",
		"E_NAME", "E_STRING", "E_NUMBER", "E_LPAREN", "E_RPAREN", "E_COMMA",
		"E_DOT", "E_PIPE", "E_EQ", "E_NEQ", "E_ASSIGN", "E_WS", "BLOCK_END",
		"KW_IF", "KW_ELIF", "KW_ELSE", "KW_ENDIF", "KW_FOR", "KW_IN", "KW_ENDFOR",
		"KW_MACRO", "KW_ENDMACRO", "KW_SET", "B_NAME", "B_STRING", "B_ASSIGN",
		"B_LPAREN", "B_RPAREN", "B_COMMA", "B_WS", "COMMENT_END", "COMMENT_BODY",
	}
	staticData.RuleNames = []string{
		"template", "element", "expr", "primary", "argList", "arg", "ifBlock",
		"elifClause", "elseClause", "blockExpr", "forBlock", "macroBlock", "paramList",
		"param", "setStmt",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 37, 214, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 1, 0, 5, 0,
		32, 8, 0, 10, 0, 12, 0, 35, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 50, 8, 1, 10, 1, 12, 1, 53,
		9, 1, 1, 1, 3, 1, 56, 8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 3, 2, 67, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 75,
		8, 2, 1, 2, 3, 2, 78, 8, 2, 5, 2, 80, 8, 2, 10, 2, 12, 2, 83, 9, 2, 1,
		3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 92, 8, 3, 1, 4, 1, 4, 1, 4,
		5, 4, 97, 8, 4, 10, 4, 12, 4, 100, 9, 4, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5,
		106, 8, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 5, 6, 113, 8, 6, 10, 6, 12, 6,
		116, 9, 6, 1, 6, 5, 6, 119, 8, 6, 10, 6, 12, 6, 122, 9, 6, 1, 6, 3, 6,
		125, 8, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 5, 7,
		136, 8, 7, 10, 7, 12, 7, 139, 9, 7, 1, 8, 1, 8, 1, 8, 1, 8, 5, 8, 145,
		8, 8, 10, 8, 12, 8, 148, 9, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		3, 9, 157, 8, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 5, 10,
		166, 8, 10, 10, 10, 12, 10, 169, 9, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1,
		11, 1, 11, 1, 11, 1, 11, 1, 11, 3, 11, 180, 8, 11, 1, 11, 1, 11, 1, 11,
		5, 11, 185, 8, 11, 10, 11, 12, 11, 188, 9, 11, 1, 11, 1, 11, 1, 11, 1,
		11, 1, 12, 1, 12, 1, 12, 5, 12, 197, 8, 12, 10, 12, 12, 12, 200, 9, 12,
		1, 13, 1, 13, 1, 13, 3, 13, 205, 8, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1,
		14, 1, 14, 1, 14, 1, 14, 0, 1, 4, 15, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18,
		20, 22, 24, 26, 28, 0, 1, 1, 0, 29, 30, 229, 0, 33, 1, 0, 0, 0, 2, 55,
		1, 0, 0, 0, 4, 57, 1, 0, 0, 0, 6, 91, 1, 0, 0, 0, 8, 93, 1, 0, 0, 0, 10,
		105, 1, 0, 0, 0, 12, 107, 1, 0, 0, 0, 14, 130, 1, 0, 0, 0, 16, 140, 1,
		0, 0, 0, 18, 156, 1, 0, 0, 0, 20, 158, 1, 0, 0, 0, 22, 174, 1, 0, 0, 0,
		24, 193, 1, 0, 0, 0, 26, 201, 1, 0, 0, 0, 28, 206, 1, 0, 0, 0, 30, 32,
		3, 2, 1, 0, 31, 30, 1, 0, 0, 0, 32, 35, 1, 0, 0, 0, 33, 31, 1, 0, 0, 0,
		33, 34, 1, 0, 0, 0, 34, 36, 1, 0, 0, 0, 35, 33, 1, 0, 0, 0, 36, 37, 5,
		0, 0, 1, 37, 1, 1, 0, 0, 0, 38, 56, 5, 4, 0, 0, 39, 40, 5, 1, 0, 0, 40,
		41, 3, 4, 2, 0, 41, 42, 5, 5, 0, 0, 42, 56, 1, 0, 0, 0, 43, 56, 3, 12,
		6, 0, 44, 56, 3, 20, 10, 0, 45, 56, 3, 22, 11, 0, 46, 56, 3, 28, 14, 0,
		47, 51, 5, 3, 0, 0, 48, 50, 5, 37, 0, 0, 49, 48, 1, 0, 0, 0, 50, 53, 1,
		0, 0, 0, 51, 49, 1, 0, 0, 0, 51, 52, 1, 0, 0, 0, 52, 54, 1, 0, 0, 0, 53,
		51, 1, 0, 0, 0, 54, 56, 5, 36, 0, 0, 55, 38, 1, 0, 0, 0, 55, 39, 1, 0,
		0, 0, 55, 43, 1, 0, 0, 0, 55, 44, 1, 0, 0, 0, 55, 45, 1, 0, 0, 0, 55, 46,
		1, 0, 0, 0, 55, 47, 1, 0, 0, 0, 56, 3, 1, 0, 0, 0, 57, 58, 6, 2, -1, 0,
		58, 59, 3, 6, 3, 0, 59, 81, 1, 0, 0, 0, 60, 61, 10, 3, 0, 0, 61, 62, 5,
		12, 0, 0, 62, 80, 5, 6, 0, 0, 63, 64, 10, 2, 0, 0, 64, 66, 5, 9, 0, 0,
		65, 67, 3, 8, 4, 0, 66, 65, 1, 0, 0, 0, 66, 67, 1, 0, 0, 0, 67, 68, 1,
		0, 0, 0, 68, 80, 5, 10, 0, 0, 69, 70, 10, 1, 0, 0, 70, 71, 5, 13, 0, 0,
		71, 77, 5, 6, 0, 0, 72, 74, 5, 9, 0, 0, 73, 75, 3, 8, 4, 0, 74, 73, 1,
		0, 0, 0, 74, 75, 1, 0, 0, 0, 75, 76, 1, 0, 0, 0, 76, 78, 5, 10, 0, 0, 77,
		72, 1, 0, 0, 0, 77, 78, 1, 0, 0, 0, 78, 80, 1, 0, 0, 0, 79, 60, 1, 0, 0,
		0, 79, 63, 1, 0, 0, 0, 79, 69, 1, 0, 0, 0, 80, 83, 1, 0, 0, 0, 81, 79,
		1, 0, 0, 0, 81, 82, 1, 0, 0, 0, 82, 5, 1, 0, 0, 0, 83, 81, 1, 0, 0, 0,
		84, 92, 5, 6, 0, 0, 85, 92, 5, 7, 0, 0, 86, 92, 5, 8, 0, 0, 87, 88, 5,
		9, 0, 0, 88, 89, 3, 4, 2, 0, 89, 90, 5, 10, 0, 0, 90, 92, 1, 0, 0, 0, 91,
		84, 1, 0, 0, 0, 91, 85, 1, 0, 0, 0, 91, 86, 1, 0, 0, 0, 91, 87, 1, 0, 0,
		0, 92, 7, 1, 0, 0, 0, 93, 98, 3, 10, 5, 0, 94, 95, 5, 11, 0, 0, 95, 97,
		3, 10, 5, 0, 96, 94, 1, 0, 0, 0, 97, 100, 1, 0, 0, 0, 98, 96, 1, 0, 0,
		0, 98, 99, 1, 0, 0, 0, 99, 9, 1, 0, 0, 0, 100, 98, 1, 0, 0, 0, 101, 102,
		5, 6, 0, 0, 102, 103, 5, 16, 0, 0, 103, 106, 3, 4, 2, 0, 104, 106, 3, 4,
		2, 0, 105, 101, 1, 0, 0, 0, 105, 104, 1, 0, 0, 0, 106, 11, 1, 0, 0, 0,
		107, 108, 5, 2, 0, 0, 108, 109, 5, 19, 0, 0, 109, 110, 3, 18, 9, 0, 110,
		114, 5, 18, 0, 0, 111, 113, 3, 2, 1, 0, 112, 111, 1, 0, 0, 0, 113, 116,
		1, 0, 0, 0, 114, 112, 1, 0, 0, 0, 114, 115, 1, 0, 0, 0, 115, 120, 1, 0,
		0, 0, 116, 114, 1, 0, 0, 0, 117, 119, 3, 14, 7, 0, 118, 117, 1, 0, 0, 0,
		119, 122, 1, 0, 0, 0, 120, 118, 1, 0, 0, 0, 120, 121, 1, 0, 0, 0, 121,
		124, 1, 0, 0, 0, 122, 120, 1, 0, 0, 0, 123, 125, 3, 16, 8, 0, 124, 123,
		1, 0, 0, 0, 124, 125, 1, 0, 0, 0, 125, 126, 1, 0, 0, 0, 126, 127, 5, 2,
		0, 0, 127, 128, 5, 22, 0, 0, 128, 129, 5, 18, 0, 0, 129, 13, 1, 0, 0, 0,
		130, 131, 5, 2, 0, 0, 131, 132, 5, 20, 0, 0, 132, 133, 3, 18, 9, 0, 133,
		137, 5, 18, 0, 0, 134, 136, 3, 2, 1, 0, 135, 134, 1, 0, 0, 0, 136, 139,
		1, 0, 0, 0, 137, 135, 1, 0, 0, 0, 137, 138, 1, 0, 0, 0, 138, 15, 1, 0,
		0, 0, 139, 137, 1, 0, 0, 0, 140, 141, 5, 2, 0, 0, 141, 142, 5, 21, 0, 0,
		142, 146, 5, 18, 0, 0, 143, 145, 3, 2, 1, 0, 144, 143, 1, 0, 0, 0, 145,
		148, 1, 0, 0, 0, 146, 144, 1, 0, 0, 0, 146, 147, 1, 0, 0, 0, 147, 17, 1,
		0, 0, 0, 148, 146, 1, 0, 0, 0, 149, 157, 5, 29, 0, 0, 150, 151, 5, 29,
		0, 0, 151, 152, 5, 31, 0, 0, 152, 157, 5, 30, 0, 0, 153, 154, 5, 29, 0,
		0, 154, 155, 5, 31, 0, 0, 155, 157, 5, 29, 0, 0, 156, 149, 1, 0, 0, 0,
		156, 150, 1, 0, 0, 0, 156, 153, 1, 0, 0, 0, 157, 19, 1, 0, 0, 0, 158, 159,
		5, 2, 0, 0, 159, 160, 5, 23, 0, 0, 160, 161, 5, 29, 0, 0, 161, 162, 5,
		24, 0, 0, 162, 163, 5, 29, 0, 0, 163, 167, 5, 18, 0, 0, 164, 166, 3, 2,
		1, 0, 165, 164, 1, 0, 0, 0, 166, 169, 1, 0, 0, 0, 167, 165, 1, 0, 0, 0,
		167, 168, 1, 0, 0, 0, 168, 170, 1, 0, 0, 0, 169, 167, 1, 0, 0, 0, 170,
		171, 5, 2, 0, 0, 171, 172, 5, 25, 0, 0, 172, 173, 5, 18, 0, 0, 173, 21,
		1, 0, 0, 0, 174, 175, 5, 2, 0, 0, 175, 176, 5, 26, 0, 0, 176, 177, 5, 29,
		0, 0, 177, 179, 5, 32, 0, 0, 178, 180, 3, 24, 12, 0, 179, 178, 1, 0, 0,
		0, 179, 180, 1, 0, 0, 0, 180, 181, 1, 0, 0, 0, 181, 182, 5, 33, 0, 0, 182,
		186, 5, 18, 0, 0, 183, 185, 3, 2, 1, 0, 184, 183, 1, 0, 0, 0, 185, 188,
		1, 0, 0, 0, 186, 184, 1, 0, 0, 0, 186, 187, 1, 0, 0, 0, 187, 189, 1, 0,
		0, 0, 188, 186, 1, 0, 0, 0, 189, 190, 5, 2, 0, 0, 190, 191, 5, 27, 0, 0,
		191, 192, 5, 18, 0, 0, 192, 23, 1, 0, 0, 0, 193, 198, 3, 26, 13, 0, 194,
		195, 5, 34, 0, 0, 195, 197, 3, 26, 13, 0, 196, 194, 1, 0, 0, 0, 197, 200,
		1, 0, 0, 0, 198, 196, 1, 0, 0, 0, 198, 199, 1, 0, 0, 0, 199, 25, 1, 0,
		0, 0, 200, 198, 1, 0, 0, 0, 201, 204, 5, 29, 0, 0, 202, 203, 5, 31, 0,
		0, 203, 205, 5, 30, 0, 0, 204, 202, 1, 0, 0, 0, 204, 205, 1, 0, 0, 0, 205,
		27, 1, 0, 0, 0, 206, 207, 5, 2, 0, 0, 207, 208, 5, 28, 0, 0, 208, 209,
		5, 29, 0, 0, 209, 210, 5, 31, 0, 0, 210, 211, 7, 0, 0, 0, 211, 212, 5,
		18, 0, 0, 212, 29, 1, 0, 0, 0, 22, 33, 51, 55, 66, 74, 77, 79, 81, 91,
		98, 105, 114, 120, 124, 137, 146, 156, 167, 179, 186, 198, 204,
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

// JinjaInit initializes any static state used to implement Jinja. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewJinja(). You can call this function if you wish to initialize the static state ahead
// of time.
func JinjaInit() {
	staticData := &JinjaParserStaticData
	staticData.once.Do(jinjaParserInit)
}

// NewJinja produces a new parser instance for the optional input antlr.TokenStream.
func NewJinja(input antlr.TokenStream) *Jinja {
	JinjaInit()
	this := new(Jinja)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &JinjaParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Jinja.g4"

	return this
}

// Jinja tokens.
const (
	JinjaEOF           = antlr.TokenEOF
	JinjaEXPR_START    = 1
	JinjaBLOCK_START   = 2
	JinjaCOMMENT_START = 3
	JinjaTEXT          = 4
	JinjaEXPR_END      = 5
	JinjaE_NAME        = 6
	JinjaE_STRING      = 7
	JinjaE_NUMBER      = 8
	JinjaE_LPAREN      = 9
	JinjaE_RPAREN      = 10
	JinjaE_COMMA       = 11
	JinjaE_DOT         = 12
	JinjaE_PIPE        = 13
	JinjaE_EQ          = 14
	JinjaE_NEQ         = 15
	JinjaE_ASSIGN      = 16
	JinjaE_WS          = 17
	JinjaBLOCK_END     = 18
	JinjaKW_IF         = 19
	JinjaKW_ELIF       = 20
	JinjaKW_ELSE       = 21
	JinjaKW_ENDIF      = 22
	JinjaKW_FOR        = 23
	JinjaKW_IN         = 24
	JinjaKW_ENDFOR     = 25
	JinjaKW_MACRO      = 26
	JinjaKW_ENDMACRO   = 27
	JinjaKW_SET        = 28
	JinjaB_NAME        = 29
	JinjaB_STRING      = 30
	JinjaB_ASSIGN      = 31
	JinjaB_LPAREN      = 32
	JinjaB_RPAREN      = 33
	JinjaB_COMMA       = 34
	JinjaB_WS          = 35
	JinjaCOMMENT_END   = 36
	JinjaCOMMENT_BODY  = 37
)

// Jinja rules.
const (
	JinjaRULE_template   = 0
	JinjaRULE_element    = 1
	JinjaRULE_expr       = 2
	JinjaRULE_primary    = 3
	JinjaRULE_argList    = 4
	JinjaRULE_arg        = 5
	JinjaRULE_ifBlock    = 6
	JinjaRULE_elifClause = 7
	JinjaRULE_elseClause = 8
	JinjaRULE_blockExpr  = 9
	JinjaRULE_forBlock   = 10
	JinjaRULE_macroBlock = 11
	JinjaRULE_paramList  = 12
	JinjaRULE_param      = 13
	JinjaRULE_setStmt    = 14
)

// ITemplateContext is an interface to support dynamic dispatch.
type ITemplateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllElement() []IElementContext
	Element(i int) IElementContext

	// IsTemplateContext differentiates from other interfaces.
	IsTemplateContext()
}

type TemplateContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTemplateContext() *TemplateContext {
	var p = new(TemplateContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JinjaRULE_template
	return p
}

func InitEmptyTemplateContext(p *TemplateContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JinjaRULE_template
}

func (*TemplateContext) IsTemplateContext() {}

func NewTemplateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TemplateContext {
	var p = new(TemplateContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JinjaRULE_template

	return p
}

func (s *TemplateContext) GetParser() antlr.Parser { return s.parser }

func (s *TemplateContext) EOF() antlr.TerminalNode {
	return s.GetToken(JinjaEOF, 0)
}

func (s *TemplateContext) AllElement() []IElementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IElementContext); ok {
			len++
		}
	}

	tst := make([]IElementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IElementContext); ok {
			tst[i] = t.(IElementContext)
			i++
		}
	}

	return tst
}

func (s *TemplateContext) Element(i int) IElementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElementContext); ok {
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

	return t.(IElementContext)
}

func (s *TemplateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TemplateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TemplateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JinjaListener); ok {
		listenerT.EnterTemplate(s)
	}
}

func (s *TemplateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JinjaListener); ok {
		listenerT.ExitTemplate(s)
	}
}

func (s *TemplateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JinjaVisitor:
		return t.VisitTemplate(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Jinja) Template() (localctx ITemplateContext) {
	localctx = NewTemplateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, JinjaRULE_template)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(33)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&30) != 0 {
		{
			p.SetState(30)
			p.Element()
		}

		p.SetState(35)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(36)
		p.Match(JinjaEOF)
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

// IElementContext is an interface to support dynamic dispatch.
type IElementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsElementContext differentiates from other interfaces.
	IsElementContext()
}

type ElementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElementContext() *ElementContext {
	var p = new(ElementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JinjaRULE_element
	return p
}

func InitEmptyElementContext(p *ElementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JinjaRULE_element
}

func (*ElementContext) IsElementContext() {}

func NewElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElementContext {
	var p = new(ElementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JinjaRULE_element

	return p
}

func (s *ElementContext) GetParser() antlr.Parser { return s.parser }

func (s *ElementContext) CopyAll(ctx *ElementContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type RawTextContext struct {
	ElementContext
}

func NewRawTextContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RawTextContext {
	var p = new(RawTextContext)

	InitEmptyElementContext(&p.ElementContext)
	p.parser = parser
	p.CopyAll(ctx.(*ElementContext))

	return p
}

func (s *RawTextContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RawTextContext) TEXT() antlr.TerminalNode {
	return s.GetToken(JinjaTEXT, 0)
}

func (s *RawTextContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JinjaListener); ok {
		listenerT.EnterRawText(s)
	}
}

func (s *RawTextContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JinjaListener); ok {
		listenerT.ExitRawText(s)
	}
}

func (s *RawTextContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JinjaVisitor:
		return t.VisitRawText(s)

	default:
		return t.VisitChildren(s)
	}
}

type MacroTagContext struct {
	ElementContext
}

func NewMacroTagContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MacroTagContext {
	var p = new(MacroTagContext)

	InitEmptyElementContext(&p.ElementContext)
	p.parser = parser
	p.CopyAll(ctx.(*ElementContext))

	return p
}

func (s *MacroTagContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MacroTagContext) MacroBlock() IMacroBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMacroBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMacroBlockContext)
}

func (s *MacroTagContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JinjaListener); ok {
		listenerT.EnterMacroTag(s)
	}
}

func (s *MacroTagContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JinjaListener); ok {
		listenerT.ExitMacroTag(s)
	}
}

func (s *MacroTagContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JinjaVisitor:
		return t.VisitMacroTag(s)

	default:
		return t.VisitChildren(s)
	}
}

type IfTagContext struct {
	ElementContext
}

func NewIfTagContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfTagContext {
	var p = new(IfTagContext)

	InitEmptyElementContext(&p.ElementContext)
	p.parser = parser
	p.CopyAll(ctx.(*ElementContext))

	return p
}

func (s *IfTagContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfTagContext) IfBlock() IIfBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfBlockContext)
}

func (s *IfTagContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JinjaListener); ok {
		listenerT.EnterIfTag(s)
	}
}

func (s *IfTagContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JinjaListener); ok {
		listenerT.ExitIfTag(s)
	}
}

func (s *IfTagContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JinjaVisitor:
		return t.VisitIfTag(s)

	default:
		return t.VisitChildren(s)
	}
}

type SetTagContext struct {
	ElementContext
}

func NewSetTagContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SetTagContext {
	var p = new(SetTagContext)

	InitEmptyElementContext(&p.ElementContext)
	p.parser = parser
	p.CopyAll(ctx.(*ElementContext))

	return p
}

func (s *SetTagContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SetTagContext) SetStmt() ISetStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISetStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISetStmtContext)
}

func (s *SetTagContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JinjaListener); ok {
		listenerT.EnterSetTag(s)
	}
}

func (s *SetTagContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JinjaListener); ok {
		listenerT.ExitSetTag(s)
	}
}

func (s *SetTagContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JinjaVisitor:
		return t.VisitSetTag(s)

	default:
		return t.VisitChildren(s)
	}
}

type CommentTagContext struct {
	ElementContext
}

func NewCommentTagContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CommentTagContext {
	var p = new(CommentTagContext)

	InitEmptyElementContext(&p.ElementContext)
	p.parser = parser
	p.CopyAll(ctx.(*ElementContext))

	return p
}

func (s *CommentTagContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommentTagContext) COMMENT_START() antlr.TerminalNode {
	return s.GetToken(JinjaCOMMENT_START, 0)
}

func (s *CommentTagContext) COMMENT_END() antlr.TerminalNode {
	return s.GetToken(JinjaCOMMENT_END, 0)
}

func (s *CommentTagContext) AllCOMMENT_BODY() []antlr.TerminalNode {
	return s.GetTokens(JinjaCOMMENT_BODY)
}

func (s *CommentTagContext) COMMENT_BODY(i int) antlr.TerminalNode {
	return s.GetToken(JinjaCOMMENT_BODY, i)
}

func (s *CommentTagContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JinjaListener); ok {
		listenerT.EnterCommentTag(s)
	}
}

func (s *CommentTagContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JinjaListener); ok {
		listenerT.ExitCommentTag(s)
	}
}

func (s *CommentTagContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JinjaVisitor:
		return t.VisitCommentTag(s)

	default:
		return t.VisitChildren(s)
	}
}

type ForTagContext struct {
	ElementContext
}

func NewForTagContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForTagContext {
	var p = new(ForTagContext)

	InitEmptyElementContext(&p.ElementContext)
	p.parser = parser
	p.CopyAll(ctx.(*ElementContext))

	return p
}

func (s *ForTagContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForTagContext) ForBlock() IForBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForBlockContext)
}

func (s *ForTagContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JinjaListener); ok {
		listenerT.EnterForTag(s)
	}
}

func (s *ForTagContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JinjaListener); ok {
		listenerT.ExitForTag(s)
	}
}

func (s *ForTagContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JinjaVisitor:
		return t.VisitForTag(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExprTagContext struct {
	ElementContext
}

func NewExprTagContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprTagContext {
	var p = new(ExprTagContext)

	InitEmptyElementContext(&p.ElementContext)
	p.parser = parser
	p.CopyAll(ctx.(*ElementContext))

	return p
}

func (s *ExprTagContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprTagContext) EXPR_START() antlr.TerminalNode {
	return s.GetToken(JinjaEXPR_START, 0)
}

func (s *ExprTagContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprTagContext) EXPR_END() antlr.TerminalNode {
	return s.GetToken(JinjaEXPR_END, 0)
}

func (s *ExprTagContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JinjaListener); ok {
		listenerT.EnterExprTag(s)
	}
}

func (s *ExprTagContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JinjaListener); ok {
		listenerT.ExitExprTag(s)
	}
}

func (s *ExprTagContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JinjaVisitor:
		return t.VisitExprTag(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Jinja) Element() (localctx IElementContext) {
	localctx = NewElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, JinjaRULE_element)
	var _la int

	p.SetState(55)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		localctx = NewRawTextContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(38)
			p.Match(JinjaTEXT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewExprTagContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(39)
			p.Match(JinjaEXPR_START)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(40)
			p.expr(0)
		}
		{
			p.SetState(41)
			p.Match(JinjaEXPR_END)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewIfTagContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(43)
			p.IfBlock()
		}

	case 4:
		localctx = NewForTagContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(44)
			p.ForBlock()
		}

	case 5:
		localctx = NewMacroTagContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(45)
			p.MacroBlock()
		}

	case 6:
		localctx = NewSetTagContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(46)
			p.SetStmt()
		}

	case 7:
		localctx = NewCommentTagContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(47)
			p.Match(JinjaCOMMENT_START)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(51)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == JinjaCOMMENT_BODY {
			{
				p.SetState(48)
				p.Match(JinjaCOMMENT_BODY)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(53)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(54)
			p.Match(JinjaCOMMENT_END)
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

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JinjaRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JinjaRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JinjaRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) CopyAll(ctx *ExprContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type PrimaryExprContext struct {
	ExprContext
}

func NewPrimaryExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrimaryExprContext {
	var p = new(PrimaryExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *PrimaryExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryExprContext) Primary() IPrimaryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryContext)
}

func (s *PrimaryExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JinjaListener); ok {
		listenerT.EnterPrimaryExpr(s)
	}
}

func (s *PrimaryExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JinjaListener); ok {
		listenerT.ExitPrimaryExpr(s)
	}
}

func (s *PrimaryExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JinjaVisitor:
		return t.VisitPrimaryExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type FieldAccessContext struct {
	ExprContext
}

func NewFieldAccessContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FieldAccessContext {
	var p = new(FieldAccessContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *FieldAccessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldAccessContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *FieldAccessContext) E_DOT() antlr.TerminalNode {
	return s.GetToken(JinjaE_DOT, 0)
}

func (s *FieldAccessContext) E_NAME() antlr.TerminalNode {
	return s.GetToken(JinjaE_NAME, 0)
}

func (s *FieldAccessContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JinjaListener); ok {
		listenerT.EnterFieldAccess(s)
	}
}

func (s *FieldAccessContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JinjaListener); ok {
		listenerT.ExitFieldAccess(s)
	}
}

func (s *FieldAccessContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JinjaVisitor:
		return t.VisitFieldAccess(s)

	default:
		return t.VisitChildren(s)
	}
}

type FilterExprContext struct {
	ExprContext
}

func NewFilterExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FilterExprContext {
	var p = new(FilterExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *FilterExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FilterExprContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *FilterExprContext) E_PIPE() antlr.TerminalNode {
	return s.GetToken(JinjaE_PIPE, 0)
}

func (s *FilterExprContext) E_NAME() antlr.TerminalNode {
	return s.GetToken(JinjaE_NAME, 0)
}

func (s *FilterExprContext) E_LPAREN() antlr.TerminalNode {
	return s.GetToken(JinjaE_LPAREN, 0)
}

func (s *FilterExprContext) E_RPAREN() antlr.TerminalNode {
	return s.GetToken(JinjaE_RPAREN, 0)
}

func (s *FilterExprContext) ArgList() IArgListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgListContext)
}

func (s *FilterExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JinjaListener); ok {
		listenerT.EnterFilterExpr(s)
	}
}

func (s *FilterExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JinjaListener); ok {
		listenerT.ExitFilterExpr(s)
	}
}

func (s *FilterExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JinjaVisitor:
		return t.VisitFilterExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type FuncCallContext struct {
	ExprContext
}

func NewFuncCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FuncCallContext {
	var p = new(FuncCallContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *FuncCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncCallContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *FuncCallContext) E_LPAREN() antlr.TerminalNode {
	return s.GetToken(JinjaE_LPAREN, 0)
}

func (s *FuncCallContext) E_RPAREN() antlr.TerminalNode {
	return s.GetToken(JinjaE_RPAREN, 0)
}

func (s *FuncCallContext) ArgList() IArgListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgListContext)
}

func (s *FuncCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JinjaListener); ok {
		listenerT.EnterFuncCall(s)
	}
}

func (s *FuncCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JinjaListener); ok {
		listenerT.ExitFuncCall(s)
	}
}

func (s *FuncCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JinjaVisitor:
		return t.VisitFuncCall(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Jinja) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *Jinja) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 4
	p.EnterRecursionRule(localctx, 4, JinjaRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	localctx = NewPrimaryExprContext(p, localctx)
	p.SetParserRuleContext(localctx)
	_prevctx = localctx

	{
		p.SetState(58)
		p.Primary()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(81)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(79)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) {
			case 1:
				localctx = NewFieldAccessContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, JinjaRULE_expr)
				p.SetState(60)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(61)
					p.Match(JinjaE_DOT)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(62)
					p.Match(JinjaE_NAME)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case 2:
				localctx = NewFuncCallContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, JinjaRULE_expr)
				p.SetState(63)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(64)
					p.Match(JinjaE_LPAREN)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				p.SetState(66)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&960) != 0 {
					{
						p.SetState(65)
						p.ArgList()
					}

				}
				{
					p.SetState(68)
					p.Match(JinjaE_RPAREN)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case 3:
				localctx = NewFilterExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, JinjaRULE_expr)
				p.SetState(69)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				{
					p.SetState(70)
					p.Match(JinjaE_PIPE)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(71)
					p.Match(JinjaE_NAME)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				p.SetState(77)
				p.GetErrorHandler().Sync(p)

				if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) == 1 {
					{
						p.SetState(72)
						p.Match(JinjaE_LPAREN)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}
					p.SetState(74)
					p.GetErrorHandler().Sync(p)
					if p.HasError() {
						goto errorExit
					}
					_la = p.GetTokenStream().LA(1)

					if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&960) != 0 {
						{
							p.SetState(73)
							p.ArgList()
						}

					}
					{
						p.SetState(76)
						p.Match(JinjaE_RPAREN)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				} else if p.HasError() { // JIM
					goto errorExit
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(83)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext())
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

// IPrimaryContext is an interface to support dynamic dispatch.
type IPrimaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsPrimaryContext differentiates from other interfaces.
	IsPrimaryContext()
}

type PrimaryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryContext() *PrimaryContext {
	var p = new(PrimaryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JinjaRULE_primary
	return p
}

func InitEmptyPrimaryContext(p *PrimaryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JinjaRULE_primary
}

func (*PrimaryContext) IsPrimaryContext() {}

func NewPrimaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryContext {
	var p = new(PrimaryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JinjaRULE_primary

	return p
}

func (s *PrimaryContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryContext) CopyAll(ctx *PrimaryContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *PrimaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StringExprContext struct {
	PrimaryContext
}

func NewStringExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringExprContext {
	var p = new(StringExprContext)

	InitEmptyPrimaryContext(&p.PrimaryContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryContext))

	return p
}

func (s *StringExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringExprContext) E_STRING() antlr.TerminalNode {
	return s.GetToken(JinjaE_STRING, 0)
}

func (s *StringExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JinjaListener); ok {
		listenerT.EnterStringExpr(s)
	}
}

func (s *StringExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JinjaListener); ok {
		listenerT.ExitStringExpr(s)
	}
}

func (s *StringExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JinjaVisitor:
		return t.VisitStringExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type NumberExprContext struct {
	PrimaryContext
}

func NewNumberExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumberExprContext {
	var p = new(NumberExprContext)

	InitEmptyPrimaryContext(&p.PrimaryContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryContext))

	return p
}

func (s *NumberExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberExprContext) E_NUMBER() antlr.TerminalNode {
	return s.GetToken(JinjaE_NUMBER, 0)
}

func (s *NumberExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JinjaListener); ok {
		listenerT.EnterNumberExpr(s)
	}
}

func (s *NumberExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JinjaListener); ok {
		listenerT.ExitNumberExpr(s)
	}
}

func (s *NumberExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JinjaVisitor:
		return t.VisitNumberExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type NameExprContext struct {
	PrimaryContext
}

func NewNameExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NameExprContext {
	var p = new(NameExprContext)

	InitEmptyPrimaryContext(&p.PrimaryContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryContext))

	return p
}

func (s *NameExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NameExprContext) E_NAME() antlr.TerminalNode {
	return s.GetToken(JinjaE_NAME, 0)
}

func (s *NameExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JinjaListener); ok {
		listenerT.EnterNameExpr(s)
	}
}

func (s *NameExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JinjaListener); ok {
		listenerT.ExitNameExpr(s)
	}
}

func (s *NameExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JinjaVisitor:
		return t.VisitNameExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParenExprContext struct {
	PrimaryContext
}

func NewParenExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenExprContext {
	var p = new(ParenExprContext)

	InitEmptyPrimaryContext(&p.PrimaryContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryContext))

	return p
}

func (s *ParenExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenExprContext) E_LPAREN() antlr.TerminalNode {
	return s.GetToken(JinjaE_LPAREN, 0)
}

func (s *ParenExprContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ParenExprContext) E_RPAREN() antlr.TerminalNode {
	return s.GetToken(JinjaE_RPAREN, 0)
}

func (s *ParenExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JinjaListener); ok {
		listenerT.EnterParenExpr(s)
	}
}

func (s *ParenExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JinjaListener); ok {
		listenerT.ExitParenExpr(s)
	}
}

func (s *ParenExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JinjaVisitor:
		return t.VisitParenExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Jinja) Primary() (localctx IPrimaryContext) {
	localctx = NewPrimaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, JinjaRULE_primary)
	p.SetState(91)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case JinjaE_NAME:
		localctx = NewNameExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(84)
			p.Match(JinjaE_NAME)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case JinjaE_STRING:
		localctx = NewStringExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(85)
			p.Match(JinjaE_STRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case JinjaE_NUMBER:
		localctx = NewNumberExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(86)
			p.Match(JinjaE_NUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case JinjaE_LPAREN:
		localctx = NewParenExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(87)
			p.Match(JinjaE_LPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(88)
			p.expr(0)
		}
		{
			p.SetState(89)
			p.Match(JinjaE_RPAREN)
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

// IArgListContext is an interface to support dynamic dispatch.
type IArgListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllArg() []IArgContext
	Arg(i int) IArgContext
	AllE_COMMA() []antlr.TerminalNode
	E_COMMA(i int) antlr.TerminalNode

	// IsArgListContext differentiates from other interfaces.
	IsArgListContext()
}

type ArgListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgListContext() *ArgListContext {
	var p = new(ArgListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JinjaRULE_argList
	return p
}

func InitEmptyArgListContext(p *ArgListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JinjaRULE_argList
}

func (*ArgListContext) IsArgListContext() {}

func NewArgListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgListContext {
	var p = new(ArgListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JinjaRULE_argList

	return p
}

func (s *ArgListContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgListContext) AllArg() []IArgContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IArgContext); ok {
			len++
		}
	}

	tst := make([]IArgContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IArgContext); ok {
			tst[i] = t.(IArgContext)
			i++
		}
	}

	return tst
}

func (s *ArgListContext) Arg(i int) IArgContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgContext); ok {
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

	return t.(IArgContext)
}

func (s *ArgListContext) AllE_COMMA() []antlr.TerminalNode {
	return s.GetTokens(JinjaE_COMMA)
}

func (s *ArgListContext) E_COMMA(i int) antlr.TerminalNode {
	return s.GetToken(JinjaE_COMMA, i)
}

func (s *ArgListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JinjaListener); ok {
		listenerT.EnterArgList(s)
	}
}

func (s *ArgListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JinjaListener); ok {
		listenerT.ExitArgList(s)
	}
}

func (s *ArgListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JinjaVisitor:
		return t.VisitArgList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Jinja) ArgList() (localctx IArgListContext) {
	localctx = NewArgListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, JinjaRULE_argList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(93)
		p.Arg()
	}
	p.SetState(98)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == JinjaE_COMMA {
		{
			p.SetState(94)
			p.Match(JinjaE_COMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(95)
			p.Arg()
		}

		p.SetState(100)
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

// IArgContext is an interface to support dynamic dispatch.
type IArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsArgContext differentiates from other interfaces.
	IsArgContext()
}

type ArgContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgContext() *ArgContext {
	var p = new(ArgContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JinjaRULE_arg
	return p
}

func InitEmptyArgContext(p *ArgContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JinjaRULE_arg
}

func (*ArgContext) IsArgContext() {}

func NewArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgContext {
	var p = new(ArgContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JinjaRULE_arg

	return p
}

func (s *ArgContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgContext) CopyAll(ctx *ArgContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type KwargContext struct {
	ArgContext
}

func NewKwargContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *KwargContext {
	var p = new(KwargContext)

	InitEmptyArgContext(&p.ArgContext)
	p.parser = parser
	p.CopyAll(ctx.(*ArgContext))

	return p
}

func (s *KwargContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KwargContext) E_NAME() antlr.TerminalNode {
	return s.GetToken(JinjaE_NAME, 0)
}

func (s *KwargContext) E_ASSIGN() antlr.TerminalNode {
	return s.GetToken(JinjaE_ASSIGN, 0)
}

func (s *KwargContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *KwargContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JinjaListener); ok {
		listenerT.EnterKwarg(s)
	}
}

func (s *KwargContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JinjaListener); ok {
		listenerT.ExitKwarg(s)
	}
}

func (s *KwargContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JinjaVisitor:
		return t.VisitKwarg(s)

	default:
		return t.VisitChildren(s)
	}
}

type PosargContext struct {
	ArgContext
}

func NewPosargContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PosargContext {
	var p = new(PosargContext)

	InitEmptyArgContext(&p.ArgContext)
	p.parser = parser
	p.CopyAll(ctx.(*ArgContext))

	return p
}

func (s *PosargContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PosargContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *PosargContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JinjaListener); ok {
		listenerT.EnterPosarg(s)
	}
}

func (s *PosargContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JinjaListener); ok {
		listenerT.ExitPosarg(s)
	}
}

func (s *PosargContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JinjaVisitor:
		return t.VisitPosarg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Jinja) Arg() (localctx IArgContext) {
	localctx = NewArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, JinjaRULE_arg)
	p.SetState(105)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		localctx = NewKwargContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(101)
			p.Match(JinjaE_NAME)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(102)
			p.Match(JinjaE_ASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(103)
			p.expr(0)
		}

	case 2:
		localctx = NewPosargContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(104)
			p.expr(0)
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

// IIfBlockContext is an interface to support dynamic dispatch.
type IIfBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllBLOCK_START() []antlr.TerminalNode
	BLOCK_START(i int) antlr.TerminalNode
	KW_IF() antlr.TerminalNode
	BlockExpr() IBlockExprContext
	AllBLOCK_END() []antlr.TerminalNode
	BLOCK_END(i int) antlr.TerminalNode
	KW_ENDIF() antlr.TerminalNode
	AllElement() []IElementContext
	Element(i int) IElementContext
	AllElifClause() []IElifClauseContext
	ElifClause(i int) IElifClauseContext
	ElseClause() IElseClauseContext

	// IsIfBlockContext differentiates from other interfaces.
	IsIfBlockContext()
}

type IfBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfBlockContext() *IfBlockContext {
	var p = new(IfBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JinjaRULE_ifBlock
	return p
}

func InitEmptyIfBlockContext(p *IfBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JinjaRULE_ifBlock
}

func (*IfBlockContext) IsIfBlockContext() {}

func NewIfBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfBlockContext {
	var p = new(IfBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JinjaRULE_ifBlock

	return p
}

func (s *IfBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *IfBlockContext) AllBLOCK_START() []antlr.TerminalNode {
	return s.GetTokens(JinjaBLOCK_START)
}

func (s *IfBlockContext) BLOCK_START(i int) antlr.TerminalNode {
	return s.GetToken(JinjaBLOCK_START, i)
}

func (s *IfBlockContext) KW_IF() antlr.TerminalNode {
	return s.GetToken(JinjaKW_IF, 0)
}

func (s *IfBlockContext) BlockExpr() IBlockExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockExprContext)
}

func (s *IfBlockContext) AllBLOCK_END() []antlr.TerminalNode {
	return s.GetTokens(JinjaBLOCK_END)
}

func (s *IfBlockContext) BLOCK_END(i int) antlr.TerminalNode {
	return s.GetToken(JinjaBLOCK_END, i)
}

func (s *IfBlockContext) KW_ENDIF() antlr.TerminalNode {
	return s.GetToken(JinjaKW_ENDIF, 0)
}

func (s *IfBlockContext) AllElement() []IElementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IElementContext); ok {
			len++
		}
	}

	tst := make([]IElementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IElementContext); ok {
			tst[i] = t.(IElementContext)
			i++
		}
	}

	return tst
}

func (s *IfBlockContext) Element(i int) IElementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElementContext); ok {
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

	return t.(IElementContext)
}

func (s *IfBlockContext) AllElifClause() []IElifClauseContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IElifClauseContext); ok {
			len++
		}
	}

	tst := make([]IElifClauseContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IElifClauseContext); ok {
			tst[i] = t.(IElifClauseContext)
			i++
		}
	}

	return tst
}

func (s *IfBlockContext) ElifClause(i int) IElifClauseContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElifClauseContext); ok {
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

	return t.(IElifClauseContext)
}

func (s *IfBlockContext) ElseClause() IElseClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElseClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElseClauseContext)
}

func (s *IfBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JinjaListener); ok {
		listenerT.EnterIfBlock(s)
	}
}

func (s *IfBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JinjaListener); ok {
		listenerT.ExitIfBlock(s)
	}
}

func (s *IfBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JinjaVisitor:
		return t.VisitIfBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Jinja) IfBlock() (localctx IIfBlockContext) {
	localctx = NewIfBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, JinjaRULE_ifBlock)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(107)
		p.Match(JinjaBLOCK_START)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(108)
		p.Match(JinjaKW_IF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(109)
		p.BlockExpr()
	}
	{
		p.SetState(110)
		p.Match(JinjaBLOCK_END)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(114)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(111)
				p.Element()
			}

		}
		p.SetState(116)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(120)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(117)
				p.ElifClause()
			}

		}
		p.SetState(122)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(124)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(123)
			p.ElseClause()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(126)
		p.Match(JinjaBLOCK_START)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(127)
		p.Match(JinjaKW_ENDIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(128)
		p.Match(JinjaBLOCK_END)
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

// IElifClauseContext is an interface to support dynamic dispatch.
type IElifClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BLOCK_START() antlr.TerminalNode
	KW_ELIF() antlr.TerminalNode
	BlockExpr() IBlockExprContext
	BLOCK_END() antlr.TerminalNode
	AllElement() []IElementContext
	Element(i int) IElementContext

	// IsElifClauseContext differentiates from other interfaces.
	IsElifClauseContext()
}

type ElifClauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElifClauseContext() *ElifClauseContext {
	var p = new(ElifClauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JinjaRULE_elifClause
	return p
}

func InitEmptyElifClauseContext(p *ElifClauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JinjaRULE_elifClause
}

func (*ElifClauseContext) IsElifClauseContext() {}

func NewElifClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElifClauseContext {
	var p = new(ElifClauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JinjaRULE_elifClause

	return p
}

func (s *ElifClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *ElifClauseContext) BLOCK_START() antlr.TerminalNode {
	return s.GetToken(JinjaBLOCK_START, 0)
}

func (s *ElifClauseContext) KW_ELIF() antlr.TerminalNode {
	return s.GetToken(JinjaKW_ELIF, 0)
}

func (s *ElifClauseContext) BlockExpr() IBlockExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockExprContext)
}

func (s *ElifClauseContext) BLOCK_END() antlr.TerminalNode {
	return s.GetToken(JinjaBLOCK_END, 0)
}

func (s *ElifClauseContext) AllElement() []IElementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IElementContext); ok {
			len++
		}
	}

	tst := make([]IElementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IElementContext); ok {
			tst[i] = t.(IElementContext)
			i++
		}
	}

	return tst
}

func (s *ElifClauseContext) Element(i int) IElementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElementContext); ok {
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

	return t.(IElementContext)
}

func (s *ElifClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElifClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElifClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JinjaListener); ok {
		listenerT.EnterElifClause(s)
	}
}

func (s *ElifClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JinjaListener); ok {
		listenerT.ExitElifClause(s)
	}
}

func (s *ElifClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JinjaVisitor:
		return t.VisitElifClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Jinja) ElifClause() (localctx IElifClauseContext) {
	localctx = NewElifClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, JinjaRULE_elifClause)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(130)
		p.Match(JinjaBLOCK_START)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(131)
		p.Match(JinjaKW_ELIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(132)
		p.BlockExpr()
	}
	{
		p.SetState(133)
		p.Match(JinjaBLOCK_END)
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
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(134)
				p.Element()
			}

		}
		p.SetState(139)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext())
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
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IElseClauseContext is an interface to support dynamic dispatch.
type IElseClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BLOCK_START() antlr.TerminalNode
	KW_ELSE() antlr.TerminalNode
	BLOCK_END() antlr.TerminalNode
	AllElement() []IElementContext
	Element(i int) IElementContext

	// IsElseClauseContext differentiates from other interfaces.
	IsElseClauseContext()
}

type ElseClauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElseClauseContext() *ElseClauseContext {
	var p = new(ElseClauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JinjaRULE_elseClause
	return p
}

func InitEmptyElseClauseContext(p *ElseClauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JinjaRULE_elseClause
}

func (*ElseClauseContext) IsElseClauseContext() {}

func NewElseClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElseClauseContext {
	var p = new(ElseClauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JinjaRULE_elseClause

	return p
}

func (s *ElseClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *ElseClauseContext) BLOCK_START() antlr.TerminalNode {
	return s.GetToken(JinjaBLOCK_START, 0)
}

func (s *ElseClauseContext) KW_ELSE() antlr.TerminalNode {
	return s.GetToken(JinjaKW_ELSE, 0)
}

func (s *ElseClauseContext) BLOCK_END() antlr.TerminalNode {
	return s.GetToken(JinjaBLOCK_END, 0)
}

func (s *ElseClauseContext) AllElement() []IElementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IElementContext); ok {
			len++
		}
	}

	tst := make([]IElementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IElementContext); ok {
			tst[i] = t.(IElementContext)
			i++
		}
	}

	return tst
}

func (s *ElseClauseContext) Element(i int) IElementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElementContext); ok {
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

	return t.(IElementContext)
}

func (s *ElseClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElseClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JinjaListener); ok {
		listenerT.EnterElseClause(s)
	}
}

func (s *ElseClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JinjaListener); ok {
		listenerT.ExitElseClause(s)
	}
}

func (s *ElseClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JinjaVisitor:
		return t.VisitElseClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Jinja) ElseClause() (localctx IElseClauseContext) {
	localctx = NewElseClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, JinjaRULE_elseClause)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(140)
		p.Match(JinjaBLOCK_START)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(141)
		p.Match(JinjaKW_ELSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(142)
		p.Match(JinjaBLOCK_END)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(146)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(143)
				p.Element()
			}

		}
		p.SetState(148)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext())
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
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBlockExprContext is an interface to support dynamic dispatch.
type IBlockExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllB_NAME() []antlr.TerminalNode
	B_NAME(i int) antlr.TerminalNode
	B_ASSIGN() antlr.TerminalNode
	B_STRING() antlr.TerminalNode

	// IsBlockExprContext differentiates from other interfaces.
	IsBlockExprContext()
}

type BlockExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockExprContext() *BlockExprContext {
	var p = new(BlockExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JinjaRULE_blockExpr
	return p
}

func InitEmptyBlockExprContext(p *BlockExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JinjaRULE_blockExpr
}

func (*BlockExprContext) IsBlockExprContext() {}

func NewBlockExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockExprContext {
	var p = new(BlockExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JinjaRULE_blockExpr

	return p
}

func (s *BlockExprContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockExprContext) AllB_NAME() []antlr.TerminalNode {
	return s.GetTokens(JinjaB_NAME)
}

func (s *BlockExprContext) B_NAME(i int) antlr.TerminalNode {
	return s.GetToken(JinjaB_NAME, i)
}

func (s *BlockExprContext) B_ASSIGN() antlr.TerminalNode {
	return s.GetToken(JinjaB_ASSIGN, 0)
}

func (s *BlockExprContext) B_STRING() antlr.TerminalNode {
	return s.GetToken(JinjaB_STRING, 0)
}

func (s *BlockExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JinjaListener); ok {
		listenerT.EnterBlockExpr(s)
	}
}

func (s *BlockExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JinjaListener); ok {
		listenerT.ExitBlockExpr(s)
	}
}

func (s *BlockExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JinjaVisitor:
		return t.VisitBlockExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Jinja) BlockExpr() (localctx IBlockExprContext) {
	localctx = NewBlockExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, JinjaRULE_blockExpr)
	p.SetState(156)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(149)
			p.Match(JinjaB_NAME)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(150)
			p.Match(JinjaB_NAME)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(151)
			p.Match(JinjaB_ASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(152)
			p.Match(JinjaB_STRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(153)
			p.Match(JinjaB_NAME)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(154)
			p.Match(JinjaB_ASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(155)
			p.Match(JinjaB_NAME)
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

// IForBlockContext is an interface to support dynamic dispatch.
type IForBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllBLOCK_START() []antlr.TerminalNode
	BLOCK_START(i int) antlr.TerminalNode
	KW_FOR() antlr.TerminalNode
	AllB_NAME() []antlr.TerminalNode
	B_NAME(i int) antlr.TerminalNode
	KW_IN() antlr.TerminalNode
	AllBLOCK_END() []antlr.TerminalNode
	BLOCK_END(i int) antlr.TerminalNode
	KW_ENDFOR() antlr.TerminalNode
	AllElement() []IElementContext
	Element(i int) IElementContext

	// IsForBlockContext differentiates from other interfaces.
	IsForBlockContext()
}

type ForBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForBlockContext() *ForBlockContext {
	var p = new(ForBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JinjaRULE_forBlock
	return p
}

func InitEmptyForBlockContext(p *ForBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JinjaRULE_forBlock
}

func (*ForBlockContext) IsForBlockContext() {}

func NewForBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForBlockContext {
	var p = new(ForBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JinjaRULE_forBlock

	return p
}

func (s *ForBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *ForBlockContext) AllBLOCK_START() []antlr.TerminalNode {
	return s.GetTokens(JinjaBLOCK_START)
}

func (s *ForBlockContext) BLOCK_START(i int) antlr.TerminalNode {
	return s.GetToken(JinjaBLOCK_START, i)
}

func (s *ForBlockContext) KW_FOR() antlr.TerminalNode {
	return s.GetToken(JinjaKW_FOR, 0)
}

func (s *ForBlockContext) AllB_NAME() []antlr.TerminalNode {
	return s.GetTokens(JinjaB_NAME)
}

func (s *ForBlockContext) B_NAME(i int) antlr.TerminalNode {
	return s.GetToken(JinjaB_NAME, i)
}

func (s *ForBlockContext) KW_IN() antlr.TerminalNode {
	return s.GetToken(JinjaKW_IN, 0)
}

func (s *ForBlockContext) AllBLOCK_END() []antlr.TerminalNode {
	return s.GetTokens(JinjaBLOCK_END)
}

func (s *ForBlockContext) BLOCK_END(i int) antlr.TerminalNode {
	return s.GetToken(JinjaBLOCK_END, i)
}

func (s *ForBlockContext) KW_ENDFOR() antlr.TerminalNode {
	return s.GetToken(JinjaKW_ENDFOR, 0)
}

func (s *ForBlockContext) AllElement() []IElementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IElementContext); ok {
			len++
		}
	}

	tst := make([]IElementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IElementContext); ok {
			tst[i] = t.(IElementContext)
			i++
		}
	}

	return tst
}

func (s *ForBlockContext) Element(i int) IElementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElementContext); ok {
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

	return t.(IElementContext)
}

func (s *ForBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JinjaListener); ok {
		listenerT.EnterForBlock(s)
	}
}

func (s *ForBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JinjaListener); ok {
		listenerT.ExitForBlock(s)
	}
}

func (s *ForBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JinjaVisitor:
		return t.VisitForBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Jinja) ForBlock() (localctx IForBlockContext) {
	localctx = NewForBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, JinjaRULE_forBlock)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(158)
		p.Match(JinjaBLOCK_START)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(159)
		p.Match(JinjaKW_FOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(160)
		p.Match(JinjaB_NAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(161)
		p.Match(JinjaKW_IN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(162)
		p.Match(JinjaB_NAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(163)
		p.Match(JinjaBLOCK_END)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(167)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(164)
				p.Element()
			}

		}
		p.SetState(169)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	{
		p.SetState(170)
		p.Match(JinjaBLOCK_START)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(171)
		p.Match(JinjaKW_ENDFOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(172)
		p.Match(JinjaBLOCK_END)
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

// IMacroBlockContext is an interface to support dynamic dispatch.
type IMacroBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllBLOCK_START() []antlr.TerminalNode
	BLOCK_START(i int) antlr.TerminalNode
	KW_MACRO() antlr.TerminalNode
	B_NAME() antlr.TerminalNode
	B_LPAREN() antlr.TerminalNode
	B_RPAREN() antlr.TerminalNode
	AllBLOCK_END() []antlr.TerminalNode
	BLOCK_END(i int) antlr.TerminalNode
	KW_ENDMACRO() antlr.TerminalNode
	ParamList() IParamListContext
	AllElement() []IElementContext
	Element(i int) IElementContext

	// IsMacroBlockContext differentiates from other interfaces.
	IsMacroBlockContext()
}

type MacroBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMacroBlockContext() *MacroBlockContext {
	var p = new(MacroBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JinjaRULE_macroBlock
	return p
}

func InitEmptyMacroBlockContext(p *MacroBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JinjaRULE_macroBlock
}

func (*MacroBlockContext) IsMacroBlockContext() {}

func NewMacroBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MacroBlockContext {
	var p = new(MacroBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JinjaRULE_macroBlock

	return p
}

func (s *MacroBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *MacroBlockContext) AllBLOCK_START() []antlr.TerminalNode {
	return s.GetTokens(JinjaBLOCK_START)
}

func (s *MacroBlockContext) BLOCK_START(i int) antlr.TerminalNode {
	return s.GetToken(JinjaBLOCK_START, i)
}

func (s *MacroBlockContext) KW_MACRO() antlr.TerminalNode {
	return s.GetToken(JinjaKW_MACRO, 0)
}

func (s *MacroBlockContext) B_NAME() antlr.TerminalNode {
	return s.GetToken(JinjaB_NAME, 0)
}

func (s *MacroBlockContext) B_LPAREN() antlr.TerminalNode {
	return s.GetToken(JinjaB_LPAREN, 0)
}

func (s *MacroBlockContext) B_RPAREN() antlr.TerminalNode {
	return s.GetToken(JinjaB_RPAREN, 0)
}

func (s *MacroBlockContext) AllBLOCK_END() []antlr.TerminalNode {
	return s.GetTokens(JinjaBLOCK_END)
}

func (s *MacroBlockContext) BLOCK_END(i int) antlr.TerminalNode {
	return s.GetToken(JinjaBLOCK_END, i)
}

func (s *MacroBlockContext) KW_ENDMACRO() antlr.TerminalNode {
	return s.GetToken(JinjaKW_ENDMACRO, 0)
}

func (s *MacroBlockContext) ParamList() IParamListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamListContext)
}

func (s *MacroBlockContext) AllElement() []IElementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IElementContext); ok {
			len++
		}
	}

	tst := make([]IElementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IElementContext); ok {
			tst[i] = t.(IElementContext)
			i++
		}
	}

	return tst
}

func (s *MacroBlockContext) Element(i int) IElementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElementContext); ok {
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

	return t.(IElementContext)
}

func (s *MacroBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MacroBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MacroBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JinjaListener); ok {
		listenerT.EnterMacroBlock(s)
	}
}

func (s *MacroBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JinjaListener); ok {
		listenerT.ExitMacroBlock(s)
	}
}

func (s *MacroBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JinjaVisitor:
		return t.VisitMacroBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Jinja) MacroBlock() (localctx IMacroBlockContext) {
	localctx = NewMacroBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, JinjaRULE_macroBlock)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(174)
		p.Match(JinjaBLOCK_START)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(175)
		p.Match(JinjaKW_MACRO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(176)
		p.Match(JinjaB_NAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(177)
		p.Match(JinjaB_LPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(179)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == JinjaB_NAME {
		{
			p.SetState(178)
			p.ParamList()
		}

	}
	{
		p.SetState(181)
		p.Match(JinjaB_RPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(182)
		p.Match(JinjaBLOCK_END)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(186)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(183)
				p.Element()
			}

		}
		p.SetState(188)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	{
		p.SetState(189)
		p.Match(JinjaBLOCK_START)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(190)
		p.Match(JinjaKW_ENDMACRO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(191)
		p.Match(JinjaBLOCK_END)
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

// IParamListContext is an interface to support dynamic dispatch.
type IParamListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllParam() []IParamContext
	Param(i int) IParamContext
	AllB_COMMA() []antlr.TerminalNode
	B_COMMA(i int) antlr.TerminalNode

	// IsParamListContext differentiates from other interfaces.
	IsParamListContext()
}

type ParamListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamListContext() *ParamListContext {
	var p = new(ParamListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JinjaRULE_paramList
	return p
}

func InitEmptyParamListContext(p *ParamListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JinjaRULE_paramList
}

func (*ParamListContext) IsParamListContext() {}

func NewParamListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamListContext {
	var p = new(ParamListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JinjaRULE_paramList

	return p
}

func (s *ParamListContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamListContext) AllParam() []IParamContext {
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

func (s *ParamListContext) Param(i int) IParamContext {
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

func (s *ParamListContext) AllB_COMMA() []antlr.TerminalNode {
	return s.GetTokens(JinjaB_COMMA)
}

func (s *ParamListContext) B_COMMA(i int) antlr.TerminalNode {
	return s.GetToken(JinjaB_COMMA, i)
}

func (s *ParamListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JinjaListener); ok {
		listenerT.EnterParamList(s)
	}
}

func (s *ParamListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JinjaListener); ok {
		listenerT.ExitParamList(s)
	}
}

func (s *ParamListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JinjaVisitor:
		return t.VisitParamList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Jinja) ParamList() (localctx IParamListContext) {
	localctx = NewParamListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, JinjaRULE_paramList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(193)
		p.Param()
	}
	p.SetState(198)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == JinjaB_COMMA {
		{
			p.SetState(194)
			p.Match(JinjaB_COMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(195)
			p.Param()
		}

		p.SetState(200)
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

// IParamContext is an interface to support dynamic dispatch.
type IParamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsParamContext differentiates from other interfaces.
	IsParamContext()
}

type ParamContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamContext() *ParamContext {
	var p = new(ParamContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JinjaRULE_param
	return p
}

func InitEmptyParamContext(p *ParamContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JinjaRULE_param
}

func (*ParamContext) IsParamContext() {}

func NewParamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamContext {
	var p = new(ParamContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JinjaRULE_param

	return p
}

func (s *ParamContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamContext) CopyAll(ctx *ParamContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ParamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type NamedParamContext struct {
	ParamContext
}

func NewNamedParamContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NamedParamContext {
	var p = new(NamedParamContext)

	InitEmptyParamContext(&p.ParamContext)
	p.parser = parser
	p.CopyAll(ctx.(*ParamContext))

	return p
}

func (s *NamedParamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NamedParamContext) B_NAME() antlr.TerminalNode {
	return s.GetToken(JinjaB_NAME, 0)
}

func (s *NamedParamContext) B_ASSIGN() antlr.TerminalNode {
	return s.GetToken(JinjaB_ASSIGN, 0)
}

func (s *NamedParamContext) B_STRING() antlr.TerminalNode {
	return s.GetToken(JinjaB_STRING, 0)
}

func (s *NamedParamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JinjaListener); ok {
		listenerT.EnterNamedParam(s)
	}
}

func (s *NamedParamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JinjaListener); ok {
		listenerT.ExitNamedParam(s)
	}
}

func (s *NamedParamContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JinjaVisitor:
		return t.VisitNamedParam(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Jinja) Param() (localctx IParamContext) {
	localctx = NewParamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, JinjaRULE_param)
	var _la int

	localctx = NewNamedParamContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(201)
		p.Match(JinjaB_NAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(204)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == JinjaB_ASSIGN {
		{
			p.SetState(202)
			p.Match(JinjaB_ASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(203)
			p.Match(JinjaB_STRING)
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

// ISetStmtContext is an interface to support dynamic dispatch.
type ISetStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BLOCK_START() antlr.TerminalNode
	KW_SET() antlr.TerminalNode
	AllB_NAME() []antlr.TerminalNode
	B_NAME(i int) antlr.TerminalNode
	B_ASSIGN() antlr.TerminalNode
	BLOCK_END() antlr.TerminalNode
	B_STRING() antlr.TerminalNode

	// IsSetStmtContext differentiates from other interfaces.
	IsSetStmtContext()
}

type SetStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySetStmtContext() *SetStmtContext {
	var p = new(SetStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JinjaRULE_setStmt
	return p
}

func InitEmptySetStmtContext(p *SetStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JinjaRULE_setStmt
}

func (*SetStmtContext) IsSetStmtContext() {}

func NewSetStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SetStmtContext {
	var p = new(SetStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JinjaRULE_setStmt

	return p
}

func (s *SetStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *SetStmtContext) BLOCK_START() antlr.TerminalNode {
	return s.GetToken(JinjaBLOCK_START, 0)
}

func (s *SetStmtContext) KW_SET() antlr.TerminalNode {
	return s.GetToken(JinjaKW_SET, 0)
}

func (s *SetStmtContext) AllB_NAME() []antlr.TerminalNode {
	return s.GetTokens(JinjaB_NAME)
}

func (s *SetStmtContext) B_NAME(i int) antlr.TerminalNode {
	return s.GetToken(JinjaB_NAME, i)
}

func (s *SetStmtContext) B_ASSIGN() antlr.TerminalNode {
	return s.GetToken(JinjaB_ASSIGN, 0)
}

func (s *SetStmtContext) BLOCK_END() antlr.TerminalNode {
	return s.GetToken(JinjaBLOCK_END, 0)
}

func (s *SetStmtContext) B_STRING() antlr.TerminalNode {
	return s.GetToken(JinjaB_STRING, 0)
}

func (s *SetStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SetStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SetStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JinjaListener); ok {
		listenerT.EnterSetStmt(s)
	}
}

func (s *SetStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JinjaListener); ok {
		listenerT.ExitSetStmt(s)
	}
}

func (s *SetStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JinjaVisitor:
		return t.VisitSetStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Jinja) SetStmt() (localctx ISetStmtContext) {
	localctx = NewSetStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, JinjaRULE_setStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(206)
		p.Match(JinjaBLOCK_START)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(207)
		p.Match(JinjaKW_SET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(208)
		p.Match(JinjaB_NAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(209)
		p.Match(JinjaB_ASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(210)
		_la = p.GetTokenStream().LA(1)

		if !(_la == JinjaB_NAME || _la == JinjaB_STRING) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(211)
		p.Match(JinjaBLOCK_END)
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

func (p *Jinja) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 2:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *Jinja) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
