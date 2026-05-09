// Code generated from JinjaLexer.g4 by ANTLR 4.13.2. DO NOT EDIT.

package jinjagrammar

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type JinjaLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var JinjaLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func jinjalexerLexerInit() {
	staticData := &JinjaLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE", "EXPR_MODE", "BLOCK_MODE", "COMMENT_MODE",
	}
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
		"EXPR_START", "BLOCK_START", "COMMENT_START", "TEXT", "EXPR_END", "E_NAME",
		"E_STRING", "E_NUMBER", "E_LPAREN", "E_RPAREN", "E_COMMA", "E_DOT",
		"E_PIPE", "E_EQ", "E_NEQ", "E_ASSIGN", "E_WS", "BLOCK_END", "KW_IF",
		"KW_ELIF", "KW_ELSE", "KW_ENDIF", "KW_FOR", "KW_IN", "KW_ENDFOR", "KW_MACRO",
		"KW_ENDMACRO", "KW_SET", "B_NAME", "B_STRING", "B_ASSIGN", "B_LPAREN",
		"B_RPAREN", "B_COMMA", "B_WS", "COMMENT_END", "COMMENT_BODY",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 37, 277, 6, -1, 6, -1, 6, -1, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2,
		7, 2, 2, 3, 7, 3, 2, 4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8,
		7, 8, 2, 9, 7, 9, 2, 10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13,
		2, 14, 7, 14, 2, 15, 7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2,
		19, 7, 19, 2, 20, 7, 20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24,
		7, 24, 2, 25, 7, 25, 2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7,
		29, 2, 30, 7, 30, 2, 31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34,
		2, 35, 7, 35, 2, 36, 7, 36, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 4, 3, 97,
		8, 3, 11, 3, 12, 3, 98, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 5, 5,
		108, 8, 5, 10, 5, 12, 5, 111, 9, 5, 1, 6, 1, 6, 5, 6, 115, 8, 6, 10, 6,
		12, 6, 118, 9, 6, 1, 6, 1, 6, 1, 6, 5, 6, 123, 8, 6, 10, 6, 12, 6, 126,
		9, 6, 1, 6, 3, 6, 129, 8, 6, 1, 7, 4, 7, 132, 8, 7, 11, 7, 12, 7, 133,
		1, 7, 1, 7, 4, 7, 138, 8, 7, 11, 7, 12, 7, 139, 3, 7, 142, 8, 7, 1, 8,
		1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13,
		1, 13, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 4, 16, 163, 8, 16, 11,
		16, 12, 16, 164, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 18,
		1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1,
		20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22,
		1, 22, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1,
		24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 26,
		1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1, 28, 1,
		28, 5, 28, 228, 8, 28, 10, 28, 12, 28, 231, 9, 28, 1, 29, 1, 29, 5, 29,
		235, 8, 29, 10, 29, 12, 29, 238, 9, 29, 1, 29, 1, 29, 1, 29, 5, 29, 243,
		8, 29, 10, 29, 12, 29, 246, 9, 29, 1, 29, 3, 29, 249, 8, 29, 1, 30, 1,
		30, 1, 31, 1, 31, 1, 32, 1, 32, 1, 33, 1, 33, 1, 34, 4, 34, 260, 8, 34,
		11, 34, 12, 34, 261, 1, 34, 1, 34, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1,
		36, 1, 36, 1, 36, 4, 36, 274, 8, 36, 11, 36, 12, 36, 275, 0, 0, 37, 4,
		1, 6, 2, 8, 3, 10, 4, 12, 5, 14, 6, 16, 7, 18, 8, 20, 9, 22, 10, 24, 11,
		26, 12, 28, 13, 30, 14, 32, 15, 34, 16, 36, 17, 38, 18, 40, 19, 42, 20,
		44, 21, 46, 22, 48, 23, 50, 24, 52, 25, 54, 26, 56, 27, 58, 28, 60, 29,
		62, 30, 64, 31, 66, 32, 68, 33, 70, 34, 72, 35, 74, 36, 76, 37, 4, 0, 1,
		2, 3, 10, 1, 0, 123, 123, 4, 0, 35, 35, 37, 37, 39, 39, 123, 123, 3, 0,
		65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 4, 0, 10,
		10, 13, 13, 39, 39, 92, 92, 4, 0, 10, 10, 13, 13, 34, 34, 92, 92, 1, 0,
		48, 57, 3, 0, 9, 10, 13, 13, 32, 32, 1, 0, 35, 35, 1, 0, 125, 125, 290,
		0, 4, 1, 0, 0, 0, 0, 6, 1, 0, 0, 0, 0, 8, 1, 0, 0, 0, 0, 10, 1, 0, 0, 0,
		1, 12, 1, 0, 0, 0, 1, 14, 1, 0, 0, 0, 1, 16, 1, 0, 0, 0, 1, 18, 1, 0, 0,
		0, 1, 20, 1, 0, 0, 0, 1, 22, 1, 0, 0, 0, 1, 24, 1, 0, 0, 0, 1, 26, 1, 0,
		0, 0, 1, 28, 1, 0, 0, 0, 1, 30, 1, 0, 0, 0, 1, 32, 1, 0, 0, 0, 1, 34, 1,
		0, 0, 0, 1, 36, 1, 0, 0, 0, 2, 38, 1, 0, 0, 0, 2, 40, 1, 0, 0, 0, 2, 42,
		1, 0, 0, 0, 2, 44, 1, 0, 0, 0, 2, 46, 1, 0, 0, 0, 2, 48, 1, 0, 0, 0, 2,
		50, 1, 0, 0, 0, 2, 52, 1, 0, 0, 0, 2, 54, 1, 0, 0, 0, 2, 56, 1, 0, 0, 0,
		2, 58, 1, 0, 0, 0, 2, 60, 1, 0, 0, 0, 2, 62, 1, 0, 0, 0, 2, 64, 1, 0, 0,
		0, 2, 66, 1, 0, 0, 0, 2, 68, 1, 0, 0, 0, 2, 70, 1, 0, 0, 0, 2, 72, 1, 0,
		0, 0, 3, 74, 1, 0, 0, 0, 3, 76, 1, 0, 0, 0, 4, 78, 1, 0, 0, 0, 6, 83, 1,
		0, 0, 0, 8, 88, 1, 0, 0, 0, 10, 96, 1, 0, 0, 0, 12, 100, 1, 0, 0, 0, 14,
		105, 1, 0, 0, 0, 16, 128, 1, 0, 0, 0, 18, 131, 1, 0, 0, 0, 20, 143, 1,
		0, 0, 0, 22, 145, 1, 0, 0, 0, 24, 147, 1, 0, 0, 0, 26, 149, 1, 0, 0, 0,
		28, 151, 1, 0, 0, 0, 30, 153, 1, 0, 0, 0, 32, 156, 1, 0, 0, 0, 34, 159,
		1, 0, 0, 0, 36, 162, 1, 0, 0, 0, 38, 168, 1, 0, 0, 0, 40, 173, 1, 0, 0,
		0, 42, 176, 1, 0, 0, 0, 44, 181, 1, 0, 0, 0, 46, 186, 1, 0, 0, 0, 48, 192,
		1, 0, 0, 0, 50, 196, 1, 0, 0, 0, 52, 199, 1, 0, 0, 0, 54, 206, 1, 0, 0,
		0, 56, 212, 1, 0, 0, 0, 58, 221, 1, 0, 0, 0, 60, 225, 1, 0, 0, 0, 62, 248,
		1, 0, 0, 0, 64, 250, 1, 0, 0, 0, 66, 252, 1, 0, 0, 0, 68, 254, 1, 0, 0,
		0, 70, 256, 1, 0, 0, 0, 72, 259, 1, 0, 0, 0, 74, 265, 1, 0, 0, 0, 76, 273,
		1, 0, 0, 0, 78, 79, 5, 123, 0, 0, 79, 80, 5, 123, 0, 0, 80, 81, 1, 0, 0,
		0, 81, 82, 6, 0, 0, 0, 82, 5, 1, 0, 0, 0, 83, 84, 5, 123, 0, 0, 84, 85,
		5, 37, 0, 0, 85, 86, 1, 0, 0, 0, 86, 87, 6, 1, 1, 0, 87, 7, 1, 0, 0, 0,
		88, 89, 5, 123, 0, 0, 89, 90, 5, 35, 0, 0, 90, 91, 1, 0, 0, 0, 91, 92,
		6, 2, 2, 0, 92, 9, 1, 0, 0, 0, 93, 97, 8, 0, 0, 0, 94, 95, 5, 123, 0, 0,
		95, 97, 8, 1, 0, 0, 96, 93, 1, 0, 0, 0, 96, 94, 1, 0, 0, 0, 97, 98, 1,
		0, 0, 0, 98, 96, 1, 0, 0, 0, 98, 99, 1, 0, 0, 0, 99, 11, 1, 0, 0, 0, 100,
		101, 5, 125, 0, 0, 101, 102, 5, 125, 0, 0, 102, 103, 1, 0, 0, 0, 103, 104,
		6, 4, 3, 0, 104, 13, 1, 0, 0, 0, 105, 109, 7, 2, 0, 0, 106, 108, 7, 3,
		0, 0, 107, 106, 1, 0, 0, 0, 108, 111, 1, 0, 0, 0, 109, 107, 1, 0, 0, 0,
		109, 110, 1, 0, 0, 0, 110, 15, 1, 0, 0, 0, 111, 109, 1, 0, 0, 0, 112, 116,
		5, 39, 0, 0, 113, 115, 8, 4, 0, 0, 114, 113, 1, 0, 0, 0, 115, 118, 1, 0,
		0, 0, 116, 114, 1, 0, 0, 0, 116, 117, 1, 0, 0, 0, 117, 119, 1, 0, 0, 0,
		118, 116, 1, 0, 0, 0, 119, 129, 5, 39, 0, 0, 120, 124, 5, 34, 0, 0, 121,
		123, 8, 5, 0, 0, 122, 121, 1, 0, 0, 0, 123, 126, 1, 0, 0, 0, 124, 122,
		1, 0, 0, 0, 124, 125, 1, 0, 0, 0, 125, 127, 1, 0, 0, 0, 126, 124, 1, 0,
		0, 0, 127, 129, 5, 34, 0, 0, 128, 112, 1, 0, 0, 0, 128, 120, 1, 0, 0, 0,
		129, 17, 1, 0, 0, 0, 130, 132, 7, 6, 0, 0, 131, 130, 1, 0, 0, 0, 132, 133,
		1, 0, 0, 0, 133, 131, 1, 0, 0, 0, 133, 134, 1, 0, 0, 0, 134, 141, 1, 0,
		0, 0, 135, 137, 5, 46, 0, 0, 136, 138, 7, 6, 0, 0, 137, 136, 1, 0, 0, 0,
		138, 139, 1, 0, 0, 0, 139, 137, 1, 0, 0, 0, 139, 140, 1, 0, 0, 0, 140,
		142, 1, 0, 0, 0, 141, 135, 1, 0, 0, 0, 141, 142, 1, 0, 0, 0, 142, 19, 1,
		0, 0, 0, 143, 144, 5, 40, 0, 0, 144, 21, 1, 0, 0, 0, 145, 146, 5, 41, 0,
		0, 146, 23, 1, 0, 0, 0, 147, 148, 5, 44, 0, 0, 148, 25, 1, 0, 0, 0, 149,
		150, 5, 46, 0, 0, 150, 27, 1, 0, 0, 0, 151, 152, 5, 124, 0, 0, 152, 29,
		1, 0, 0, 0, 153, 154, 5, 61, 0, 0, 154, 155, 5, 61, 0, 0, 155, 31, 1, 0,
		0, 0, 156, 157, 5, 33, 0, 0, 157, 158, 5, 61, 0, 0, 158, 33, 1, 0, 0, 0,
		159, 160, 5, 61, 0, 0, 160, 35, 1, 0, 0, 0, 161, 163, 7, 7, 0, 0, 162,
		161, 1, 0, 0, 0, 163, 164, 1, 0, 0, 0, 164, 162, 1, 0, 0, 0, 164, 165,
		1, 0, 0, 0, 165, 166, 1, 0, 0, 0, 166, 167, 6, 16, 4, 0, 167, 37, 1, 0,
		0, 0, 168, 169, 5, 37, 0, 0, 169, 170, 5, 125, 0, 0, 170, 171, 1, 0, 0,
		0, 171, 172, 6, 17, 3, 0, 172, 39, 1, 0, 0, 0, 173, 174, 5, 105, 0, 0,
		174, 175, 5, 102, 0, 0, 175, 41, 1, 0, 0, 0, 176, 177, 5, 101, 0, 0, 177,
		178, 5, 108, 0, 0, 178, 179, 5, 105, 0, 0, 179, 180, 5, 102, 0, 0, 180,
		43, 1, 0, 0, 0, 181, 182, 5, 101, 0, 0, 182, 183, 5, 108, 0, 0, 183, 184,
		5, 115, 0, 0, 184, 185, 5, 101, 0, 0, 185, 45, 1, 0, 0, 0, 186, 187, 5,
		101, 0, 0, 187, 188, 5, 110, 0, 0, 188, 189, 5, 100, 0, 0, 189, 190, 5,
		105, 0, 0, 190, 191, 5, 102, 0, 0, 191, 47, 1, 0, 0, 0, 192, 193, 5, 102,
		0, 0, 193, 194, 5, 111, 0, 0, 194, 195, 5, 114, 0, 0, 195, 49, 1, 0, 0,
		0, 196, 197, 5, 105, 0, 0, 197, 198, 5, 110, 0, 0, 198, 51, 1, 0, 0, 0,
		199, 200, 5, 101, 0, 0, 200, 201, 5, 110, 0, 0, 201, 202, 5, 100, 0, 0,
		202, 203, 5, 102, 0, 0, 203, 204, 5, 111, 0, 0, 204, 205, 5, 114, 0, 0,
		205, 53, 1, 0, 0, 0, 206, 207, 5, 109, 0, 0, 207, 208, 5, 97, 0, 0, 208,
		209, 5, 99, 0, 0, 209, 210, 5, 114, 0, 0, 210, 211, 5, 111, 0, 0, 211,
		55, 1, 0, 0, 0, 212, 213, 5, 101, 0, 0, 213, 214, 5, 110, 0, 0, 214, 215,
		5, 100, 0, 0, 215, 216, 5, 109, 0, 0, 216, 217, 5, 97, 0, 0, 217, 218,
		5, 99, 0, 0, 218, 219, 5, 114, 0, 0, 219, 220, 5, 111, 0, 0, 220, 57, 1,
		0, 0, 0, 221, 222, 5, 115, 0, 0, 222, 223, 5, 101, 0, 0, 223, 224, 5, 116,
		0, 0, 224, 59, 1, 0, 0, 0, 225, 229, 7, 2, 0, 0, 226, 228, 7, 3, 0, 0,
		227, 226, 1, 0, 0, 0, 228, 231, 1, 0, 0, 0, 229, 227, 1, 0, 0, 0, 229,
		230, 1, 0, 0, 0, 230, 61, 1, 0, 0, 0, 231, 229, 1, 0, 0, 0, 232, 236, 5,
		39, 0, 0, 233, 235, 8, 4, 0, 0, 234, 233, 1, 0, 0, 0, 235, 238, 1, 0, 0,
		0, 236, 234, 1, 0, 0, 0, 236, 237, 1, 0, 0, 0, 237, 239, 1, 0, 0, 0, 238,
		236, 1, 0, 0, 0, 239, 249, 5, 39, 0, 0, 240, 244, 5, 34, 0, 0, 241, 243,
		8, 5, 0, 0, 242, 241, 1, 0, 0, 0, 243, 246, 1, 0, 0, 0, 244, 242, 1, 0,
		0, 0, 244, 245, 1, 0, 0, 0, 245, 247, 1, 0, 0, 0, 246, 244, 1, 0, 0, 0,
		247, 249, 5, 34, 0, 0, 248, 232, 1, 0, 0, 0, 248, 240, 1, 0, 0, 0, 249,
		63, 1, 0, 0, 0, 250, 251, 5, 61, 0, 0, 251, 65, 1, 0, 0, 0, 252, 253, 5,
		40, 0, 0, 253, 67, 1, 0, 0, 0, 254, 255, 5, 41, 0, 0, 255, 69, 1, 0, 0,
		0, 256, 257, 5, 44, 0, 0, 257, 71, 1, 0, 0, 0, 258, 260, 7, 7, 0, 0, 259,
		258, 1, 0, 0, 0, 260, 261, 1, 0, 0, 0, 261, 259, 1, 0, 0, 0, 261, 262,
		1, 0, 0, 0, 262, 263, 1, 0, 0, 0, 263, 264, 6, 34, 4, 0, 264, 73, 1, 0,
		0, 0, 265, 266, 5, 35, 0, 0, 266, 267, 5, 125, 0, 0, 267, 268, 1, 0, 0,
		0, 268, 269, 6, 35, 3, 0, 269, 75, 1, 0, 0, 0, 270, 274, 8, 8, 0, 0, 271,
		272, 5, 35, 0, 0, 272, 274, 8, 9, 0, 0, 273, 270, 1, 0, 0, 0, 273, 271,
		1, 0, 0, 0, 274, 275, 1, 0, 0, 0, 275, 273, 1, 0, 0, 0, 275, 276, 1, 0,
		0, 0, 276, 77, 1, 0, 0, 0, 21, 0, 1, 2, 3, 96, 98, 109, 116, 124, 128,
		133, 139, 141, 164, 229, 236, 244, 248, 261, 273, 275, 5, 5, 1, 0, 5, 2,
		0, 5, 3, 0, 4, 0, 0, 6, 0, 0,
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

// JinjaLexerInit initializes any static state used to implement JinjaLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewJinjaLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func JinjaLexerInit() {
	staticData := &JinjaLexerLexerStaticData
	staticData.once.Do(jinjalexerLexerInit)
}

// NewJinjaLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewJinjaLexer(input antlr.CharStream) *JinjaLexer {
	JinjaLexerInit()
	l := new(JinjaLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &JinjaLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "JinjaLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// JinjaLexer tokens.
const (
	JinjaLexerEXPR_START    = 1
	JinjaLexerBLOCK_START   = 2
	JinjaLexerCOMMENT_START = 3
	JinjaLexerTEXT          = 4
	JinjaLexerEXPR_END      = 5
	JinjaLexerE_NAME        = 6
	JinjaLexerE_STRING      = 7
	JinjaLexerE_NUMBER      = 8
	JinjaLexerE_LPAREN      = 9
	JinjaLexerE_RPAREN      = 10
	JinjaLexerE_COMMA       = 11
	JinjaLexerE_DOT         = 12
	JinjaLexerE_PIPE        = 13
	JinjaLexerE_EQ          = 14
	JinjaLexerE_NEQ         = 15
	JinjaLexerE_ASSIGN      = 16
	JinjaLexerE_WS          = 17
	JinjaLexerBLOCK_END     = 18
	JinjaLexerKW_IF         = 19
	JinjaLexerKW_ELIF       = 20
	JinjaLexerKW_ELSE       = 21
	JinjaLexerKW_ENDIF      = 22
	JinjaLexerKW_FOR        = 23
	JinjaLexerKW_IN         = 24
	JinjaLexerKW_ENDFOR     = 25
	JinjaLexerKW_MACRO      = 26
	JinjaLexerKW_ENDMACRO   = 27
	JinjaLexerKW_SET        = 28
	JinjaLexerB_NAME        = 29
	JinjaLexerB_STRING      = 30
	JinjaLexerB_ASSIGN      = 31
	JinjaLexerB_LPAREN      = 32
	JinjaLexerB_RPAREN      = 33
	JinjaLexerB_COMMA       = 34
	JinjaLexerB_WS          = 35
	JinjaLexerCOMMENT_END   = 36
	JinjaLexerCOMMENT_BODY  = 37
)

// JinjaLexer modes.
const (
	JinjaLexerEXPR_MODE = iota + 1
	JinjaLexerBLOCK_MODE
	JinjaLexerCOMMENT_MODE
)
