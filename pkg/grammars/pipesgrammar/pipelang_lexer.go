// Code generated from PipeLangLexer.g4 by ANTLR 4.13.2. DO NOT EDIT.

package grammar

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

type PipeLangLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var PipeLangLexerLexerStaticData struct {
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

func pipelanglexerLexerInit() {
	staticData := &PipeLangLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE", "VALUE_MODE", "SECTION_MODE",
	}
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
		"COLON", "REST", "ID", "COMMENT", "WS_BLANK", "WS", "TYPE", "NAME",
		"TAGS", "OWNER", "DESTINATION", "SCHEDULE", "UNIQUE_KEY", "DESCRIPTION_ML",
		"DESCRIPTION", "SOURCES", "PARAMS", "PIPELINE", "IDENTIFIER", "COLON_TOK",
		"LBRACE", "RBRACE", "COMMA", "DASH", "STRING_LIT", "NUMBER", "INLINE_CMT",
		"VALUE", "VALUE_NL", "SEC_TYPE", "SEC_NAME", "SEC_DESC_ML", "SEC_DESC",
		"SEC_TAGS", "SEC_OWNER", "SEC_DESTINATION", "SEC_SCHEDULE", "SEC_UNIQUE_KEY",
		"SEC_SOURCES", "SEC_PARAMS", "SEC_PIPELINE", "SECTION_LINE", "SECTION_BLANK",
		"SECTION_COMMENT",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 29, 585, 6, -1, 6, -1, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2,
		2, 3, 7, 3, 2, 4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8,
		2, 9, 7, 9, 2, 10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2,
		14, 7, 14, 2, 15, 7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19,
		7, 19, 2, 20, 7, 20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7,
		24, 2, 25, 7, 25, 2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29,
		2, 30, 7, 30, 2, 31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2,
		35, 7, 35, 2, 36, 7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40,
		7, 40, 2, 41, 7, 41, 2, 42, 7, 42, 2, 43, 7, 43, 1, 0, 1, 0, 5, 0, 94,
		8, 0, 10, 0, 12, 0, 97, 9, 0, 1, 1, 5, 1, 100, 8, 1, 10, 1, 12, 1, 103,
		9, 1, 1, 1, 1, 1, 1, 2, 1, 2, 5, 2, 109, 8, 2, 10, 2, 12, 2, 112, 9, 2,
		1, 3, 1, 3, 5, 3, 116, 8, 3, 10, 3, 12, 3, 119, 9, 3, 1, 3, 1, 3, 1, 4,
		5, 4, 124, 8, 4, 10, 4, 12, 4, 127, 9, 4, 1, 4, 4, 4, 130, 8, 4, 11, 4,
		12, 4, 131, 1, 4, 1, 4, 1, 5, 4, 5, 137, 8, 5, 11, 5, 12, 5, 138, 1, 5,
		1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1,
		10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11,
		1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1,
		12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12,
		1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1,
		13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14,
		1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1,
		14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15,
		1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1,
		16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17,
		1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1,
		17, 1, 18, 1, 18, 5, 18, 307, 8, 18, 10, 18, 12, 18, 310, 9, 18, 1, 19,
		1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 1, 22, 1, 22, 1, 23, 1, 23, 1, 24, 1,
		24, 5, 24, 324, 8, 24, 10, 24, 12, 24, 327, 9, 24, 1, 24, 1, 24, 1, 25,
		3, 25, 332, 8, 25, 1, 25, 4, 25, 335, 8, 25, 11, 25, 12, 25, 336, 1, 25,
		1, 25, 5, 25, 341, 8, 25, 10, 25, 12, 25, 344, 9, 25, 3, 25, 346, 8, 25,
		1, 26, 1, 26, 5, 26, 350, 8, 26, 10, 26, 12, 26, 353, 9, 26, 1, 26, 1,
		26, 1, 27, 4, 27, 358, 8, 27, 11, 27, 12, 27, 359, 1, 27, 1, 27, 1, 27,
		1, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 1, 29, 1,
		29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 30,
		1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 31, 1, 31, 1, 31, 1,
		31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31,
		1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 32, 1, 32, 1, 32, 1, 32, 1,
		32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32,
		1, 32, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1,
		33, 1, 33, 1, 33, 1, 33, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34,
		1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1,
		35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35,
		1, 35, 1, 35, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1,
		36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 37, 1, 37, 1, 37, 1, 37,
		1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1,
		37, 1, 37, 1, 37, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38,
		1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 39, 1, 39, 1, 39, 1,
		39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39,
		1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1,
		40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 41, 4, 41, 550, 8, 41, 11, 41,
		12, 41, 551, 1, 41, 5, 41, 555, 8, 41, 10, 41, 12, 41, 558, 9, 41, 1, 41,
		3, 41, 561, 8, 41, 1, 42, 5, 42, 564, 8, 42, 10, 42, 12, 42, 567, 9, 42,
		1, 42, 1, 42, 1, 42, 1, 42, 1, 43, 1, 43, 1, 43, 5, 43, 576, 8, 43, 10,
		43, 12, 43, 579, 9, 43, 1, 43, 3, 43, 582, 8, 43, 1, 43, 1, 43, 0, 0, 44,
		3, 0, 5, 0, 7, 0, 9, 1, 11, 2, 13, 3, 15, 4, 17, 5, 19, 6, 21, 7, 23, 8,
		25, 9, 27, 10, 29, 11, 31, 12, 33, 13, 35, 14, 37, 15, 39, 16, 41, 17,
		43, 18, 45, 19, 47, 20, 49, 21, 51, 22, 53, 23, 55, 24, 57, 25, 59, 26,
		61, 0, 63, 0, 65, 0, 67, 0, 69, 0, 71, 0, 73, 0, 75, 0, 77, 0, 79, 0, 81,
		0, 83, 0, 85, 27, 87, 28, 89, 29, 3, 0, 1, 2, 8, 2, 0, 9, 9, 32, 32, 2,
		0, 10, 10, 13, 13, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90,
		95, 95, 97, 122, 4, 0, 45, 57, 65, 90, 95, 95, 97, 122, 1, 0, 34, 34, 1,
		0, 48, 57, 3, 0, 10, 10, 13, 13, 35, 35, 600, 0, 9, 1, 0, 0, 0, 0, 11,
		1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0,
		19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0,
		0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0,
		0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0,
		0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1,
		0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 1, 55, 1, 0, 0, 0, 1, 57,
		1, 0, 0, 0, 1, 59, 1, 0, 0, 0, 2, 61, 1, 0, 0, 0, 2, 63, 1, 0, 0, 0, 2,
		65, 1, 0, 0, 0, 2, 67, 1, 0, 0, 0, 2, 69, 1, 0, 0, 0, 2, 71, 1, 0, 0, 0,
		2, 73, 1, 0, 0, 0, 2, 75, 1, 0, 0, 0, 2, 77, 1, 0, 0, 0, 2, 79, 1, 0, 0,
		0, 2, 81, 1, 0, 0, 0, 2, 83, 1, 0, 0, 0, 2, 85, 1, 0, 0, 0, 2, 87, 1, 0,
		0, 0, 2, 89, 1, 0, 0, 0, 3, 91, 1, 0, 0, 0, 5, 101, 1, 0, 0, 0, 7, 106,
		1, 0, 0, 0, 9, 113, 1, 0, 0, 0, 11, 125, 1, 0, 0, 0, 13, 136, 1, 0, 0,
		0, 15, 142, 1, 0, 0, 0, 17, 152, 1, 0, 0, 0, 19, 162, 1, 0, 0, 0, 21, 172,
		1, 0, 0, 0, 23, 183, 1, 0, 0, 0, 25, 200, 1, 0, 0, 0, 27, 214, 1, 0, 0,
		0, 29, 230, 1, 0, 0, 0, 31, 249, 1, 0, 0, 0, 33, 266, 1, 0, 0, 0, 35, 278,
		1, 0, 0, 0, 37, 289, 1, 0, 0, 0, 39, 304, 1, 0, 0, 0, 41, 311, 1, 0, 0,
		0, 43, 313, 1, 0, 0, 0, 45, 315, 1, 0, 0, 0, 47, 317, 1, 0, 0, 0, 49, 319,
		1, 0, 0, 0, 51, 321, 1, 0, 0, 0, 53, 331, 1, 0, 0, 0, 55, 347, 1, 0, 0,
		0, 57, 357, 1, 0, 0, 0, 59, 365, 1, 0, 0, 0, 61, 370, 1, 0, 0, 0, 63, 381,
		1, 0, 0, 0, 65, 392, 1, 0, 0, 0, 67, 412, 1, 0, 0, 0, 69, 430, 1, 0, 0,
		0, 71, 441, 1, 0, 0, 0, 73, 453, 1, 0, 0, 0, 75, 471, 1, 0, 0, 0, 77, 486,
		1, 0, 0, 0, 79, 503, 1, 0, 0, 0, 81, 518, 1, 0, 0, 0, 83, 532, 1, 0, 0,
		0, 85, 549, 1, 0, 0, 0, 87, 565, 1, 0, 0, 0, 89, 572, 1, 0, 0, 0, 91, 95,
		5, 58, 0, 0, 92, 94, 7, 0, 0, 0, 93, 92, 1, 0, 0, 0, 94, 97, 1, 0, 0, 0,
		95, 93, 1, 0, 0, 0, 95, 96, 1, 0, 0, 0, 96, 4, 1, 0, 0, 0, 97, 95, 1, 0,
		0, 0, 98, 100, 8, 1, 0, 0, 99, 98, 1, 0, 0, 0, 100, 103, 1, 0, 0, 0, 101,
		99, 1, 0, 0, 0, 101, 102, 1, 0, 0, 0, 102, 104, 1, 0, 0, 0, 103, 101, 1,
		0, 0, 0, 104, 105, 7, 1, 0, 0, 105, 6, 1, 0, 0, 0, 106, 110, 7, 2, 0, 0,
		107, 109, 7, 3, 0, 0, 108, 107, 1, 0, 0, 0, 109, 112, 1, 0, 0, 0, 110,
		108, 1, 0, 0, 0, 110, 111, 1, 0, 0, 0, 111, 8, 1, 0, 0, 0, 112, 110, 1,
		0, 0, 0, 113, 117, 5, 35, 0, 0, 114, 116, 8, 1, 0, 0, 115, 114, 1, 0, 0,
		0, 116, 119, 1, 0, 0, 0, 117, 115, 1, 0, 0, 0, 117, 118, 1, 0, 0, 0, 118,
		120, 1, 0, 0, 0, 119, 117, 1, 0, 0, 0, 120, 121, 6, 3, 0, 0, 121, 10, 1,
		0, 0, 0, 122, 124, 7, 0, 0, 0, 123, 122, 1, 0, 0, 0, 124, 127, 1, 0, 0,
		0, 125, 123, 1, 0, 0, 0, 125, 126, 1, 0, 0, 0, 126, 129, 1, 0, 0, 0, 127,
		125, 1, 0, 0, 0, 128, 130, 7, 1, 0, 0, 129, 128, 1, 0, 0, 0, 130, 131,
		1, 0, 0, 0, 131, 129, 1, 0, 0, 0, 131, 132, 1, 0, 0, 0, 132, 133, 1, 0,
		0, 0, 133, 134, 6, 4, 0, 0, 134, 12, 1, 0, 0, 0, 135, 137, 7, 0, 0, 0,
		136, 135, 1, 0, 0, 0, 137, 138, 1, 0, 0, 0, 138, 136, 1, 0, 0, 0, 138,
		139, 1, 0, 0, 0, 139, 140, 1, 0, 0, 0, 140, 141, 6, 5, 0, 0, 141, 14, 1,
		0, 0, 0, 142, 143, 4, 6, 0, 0, 143, 144, 5, 116, 0, 0, 144, 145, 5, 121,
		0, 0, 145, 146, 5, 112, 0, 0, 146, 147, 5, 101, 0, 0, 147, 148, 1, 0, 0,
		0, 148, 149, 3, 3, 0, 0, 149, 150, 1, 0, 0, 0, 150, 151, 6, 6, 1, 0, 151,
		16, 1, 0, 0, 0, 152, 153, 4, 7, 1, 0, 153, 154, 5, 110, 0, 0, 154, 155,
		5, 97, 0, 0, 155, 156, 5, 109, 0, 0, 156, 157, 5, 101, 0, 0, 157, 158,
		1, 0, 0, 0, 158, 159, 3, 3, 0, 0, 159, 160, 1, 0, 0, 0, 160, 161, 6, 7,
		1, 0, 161, 18, 1, 0, 0, 0, 162, 163, 4, 8, 2, 0, 163, 164, 5, 116, 0, 0,
		164, 165, 5, 97, 0, 0, 165, 166, 5, 103, 0, 0, 166, 167, 5, 115, 0, 0,
		167, 168, 1, 0, 0, 0, 168, 169, 3, 3, 0, 0, 169, 170, 1, 0, 0, 0, 170,
		171, 6, 8, 1, 0, 171, 20, 1, 0, 0, 0, 172, 173, 4, 9, 3, 0, 173, 174, 5,
		111, 0, 0, 174, 175, 5, 119, 0, 0, 175, 176, 5, 110, 0, 0, 176, 177, 5,
		101, 0, 0, 177, 178, 5, 114, 0, 0, 178, 179, 1, 0, 0, 0, 179, 180, 3, 3,
		0, 0, 180, 181, 1, 0, 0, 0, 181, 182, 6, 9, 1, 0, 182, 22, 1, 0, 0, 0,
		183, 184, 4, 10, 4, 0, 184, 185, 5, 100, 0, 0, 185, 186, 5, 101, 0, 0,
		186, 187, 5, 115, 0, 0, 187, 188, 5, 116, 0, 0, 188, 189, 5, 105, 0, 0,
		189, 190, 5, 110, 0, 0, 190, 191, 5, 97, 0, 0, 191, 192, 5, 116, 0, 0,
		192, 193, 5, 105, 0, 0, 193, 194, 5, 111, 0, 0, 194, 195, 5, 110, 0, 0,
		195, 196, 1, 0, 0, 0, 196, 197, 3, 3, 0, 0, 197, 198, 1, 0, 0, 0, 198,
		199, 6, 10, 1, 0, 199, 24, 1, 0, 0, 0, 200, 201, 4, 11, 5, 0, 201, 202,
		5, 115, 0, 0, 202, 203, 5, 99, 0, 0, 203, 204, 5, 104, 0, 0, 204, 205,
		5, 101, 0, 0, 205, 206, 5, 100, 0, 0, 206, 207, 5, 117, 0, 0, 207, 208,
		5, 108, 0, 0, 208, 209, 5, 101, 0, 0, 209, 210, 1, 0, 0, 0, 210, 211, 3,
		3, 0, 0, 211, 212, 1, 0, 0, 0, 212, 213, 6, 11, 1, 0, 213, 26, 1, 0, 0,
		0, 214, 215, 4, 12, 6, 0, 215, 216, 5, 117, 0, 0, 216, 217, 5, 110, 0,
		0, 217, 218, 5, 105, 0, 0, 218, 219, 5, 113, 0, 0, 219, 220, 5, 117, 0,
		0, 220, 221, 5, 101, 0, 0, 221, 222, 5, 95, 0, 0, 222, 223, 5, 107, 0,
		0, 223, 224, 5, 101, 0, 0, 224, 225, 5, 121, 0, 0, 225, 226, 1, 0, 0, 0,
		226, 227, 3, 3, 0, 0, 227, 228, 1, 0, 0, 0, 228, 229, 6, 12, 1, 0, 229,
		28, 1, 0, 0, 0, 230, 231, 4, 13, 7, 0, 231, 232, 5, 100, 0, 0, 232, 233,
		5, 101, 0, 0, 233, 234, 5, 115, 0, 0, 234, 235, 5, 99, 0, 0, 235, 236,
		5, 114, 0, 0, 236, 237, 5, 105, 0, 0, 237, 238, 5, 112, 0, 0, 238, 239,
		5, 116, 0, 0, 239, 240, 5, 105, 0, 0, 240, 241, 5, 111, 0, 0, 241, 242,
		5, 110, 0, 0, 242, 243, 1, 0, 0, 0, 243, 244, 3, 3, 0, 0, 244, 245, 5,
		124, 0, 0, 245, 246, 3, 5, 1, 0, 246, 247, 1, 0, 0, 0, 247, 248, 6, 13,
		2, 0, 248, 30, 1, 0, 0, 0, 249, 250, 4, 14, 8, 0, 250, 251, 5, 100, 0,
		0, 251, 252, 5, 101, 0, 0, 252, 253, 5, 115, 0, 0, 253, 254, 5, 99, 0,
		0, 254, 255, 5, 114, 0, 0, 255, 256, 5, 105, 0, 0, 256, 257, 5, 112, 0,
		0, 257, 258, 5, 116, 0, 0, 258, 259, 5, 105, 0, 0, 259, 260, 5, 111, 0,
		0, 260, 261, 5, 110, 0, 0, 261, 262, 1, 0, 0, 0, 262, 263, 3, 3, 0, 0,
		263, 264, 1, 0, 0, 0, 264, 265, 6, 14, 1, 0, 265, 32, 1, 0, 0, 0, 266,
		267, 4, 15, 9, 0, 267, 268, 5, 115, 0, 0, 268, 269, 5, 111, 0, 0, 269,
		270, 5, 117, 0, 0, 270, 271, 5, 114, 0, 0, 271, 272, 5, 99, 0, 0, 272,
		273, 5, 101, 0, 0, 273, 274, 5, 115, 0, 0, 274, 275, 1, 0, 0, 0, 275, 276,
		3, 3, 0, 0, 276, 277, 3, 5, 1, 0, 277, 34, 1, 0, 0, 0, 278, 279, 4, 16,
		10, 0, 279, 280, 5, 112, 0, 0, 280, 281, 5, 97, 0, 0, 281, 282, 5, 114,
		0, 0, 282, 283, 5, 97, 0, 0, 283, 284, 5, 109, 0, 0, 284, 285, 5, 115,
		0, 0, 285, 286, 1, 0, 0, 0, 286, 287, 3, 3, 0, 0, 287, 288, 3, 5, 1, 0,
		288, 36, 1, 0, 0, 0, 289, 290, 4, 17, 11, 0, 290, 291, 5, 112, 0, 0, 291,
		292, 5, 105, 0, 0, 292, 293, 5, 112, 0, 0, 293, 294, 5, 101, 0, 0, 294,
		295, 5, 108, 0, 0, 295, 296, 5, 105, 0, 0, 296, 297, 5, 110, 0, 0, 297,
		298, 5, 101, 0, 0, 298, 299, 1, 0, 0, 0, 299, 300, 3, 3, 0, 0, 300, 301,
		3, 5, 1, 0, 301, 302, 1, 0, 0, 0, 302, 303, 6, 17, 2, 0, 303, 38, 1, 0,
		0, 0, 304, 308, 7, 2, 0, 0, 305, 307, 7, 4, 0, 0, 306, 305, 1, 0, 0, 0,
		307, 310, 1, 0, 0, 0, 308, 306, 1, 0, 0, 0, 308, 309, 1, 0, 0, 0, 309,
		40, 1, 0, 0, 0, 310, 308, 1, 0, 0, 0, 311, 312, 5, 58, 0, 0, 312, 42, 1,
		0, 0, 0, 313, 314, 5, 123, 0, 0, 314, 44, 1, 0, 0, 0, 315, 316, 5, 125,
		0, 0, 316, 46, 1, 0, 0, 0, 317, 318, 5, 44, 0, 0, 318, 48, 1, 0, 0, 0,
		319, 320, 5, 45, 0, 0, 320, 50, 1, 0, 0, 0, 321, 325, 5, 34, 0, 0, 322,
		324, 8, 5, 0, 0, 323, 322, 1, 0, 0, 0, 324, 327, 1, 0, 0, 0, 325, 323,
		1, 0, 0, 0, 325, 326, 1, 0, 0, 0, 326, 328, 1, 0, 0, 0, 327, 325, 1, 0,
		0, 0, 328, 329, 5, 34, 0, 0, 329, 52, 1, 0, 0, 0, 330, 332, 5, 45, 0, 0,
		331, 330, 1, 0, 0, 0, 331, 332, 1, 0, 0, 0, 332, 334, 1, 0, 0, 0, 333,
		335, 7, 6, 0, 0, 334, 333, 1, 0, 0, 0, 335, 336, 1, 0, 0, 0, 336, 334,
		1, 0, 0, 0, 336, 337, 1, 0, 0, 0, 337, 345, 1, 0, 0, 0, 338, 342, 5, 46,
		0, 0, 339, 341, 7, 6, 0, 0, 340, 339, 1, 0, 0, 0, 341, 344, 1, 0, 0, 0,
		342, 340, 1, 0, 0, 0, 342, 343, 1, 0, 0, 0, 343, 346, 1, 0, 0, 0, 344,
		342, 1, 0, 0, 0, 345, 338, 1, 0, 0, 0, 345, 346, 1, 0, 0, 0, 346, 54, 1,
		0, 0, 0, 347, 351, 5, 35, 0, 0, 348, 350, 8, 1, 0, 0, 349, 348, 1, 0, 0,
		0, 350, 353, 1, 0, 0, 0, 351, 349, 1, 0, 0, 0, 351, 352, 1, 0, 0, 0, 352,
		354, 1, 0, 0, 0, 353, 351, 1, 0, 0, 0, 354, 355, 6, 26, 0, 0, 355, 56,
		1, 0, 0, 0, 356, 358, 8, 7, 0, 0, 357, 356, 1, 0, 0, 0, 358, 359, 1, 0,
		0, 0, 359, 357, 1, 0, 0, 0, 359, 360, 1, 0, 0, 0, 360, 361, 1, 0, 0, 0,
		361, 362, 6, 27, 3, 0, 362, 363, 1, 0, 0, 0, 363, 364, 6, 27, 4, 0, 364,
		58, 1, 0, 0, 0, 365, 366, 7, 1, 0, 0, 366, 367, 1, 0, 0, 0, 367, 368, 6,
		28, 0, 0, 368, 369, 6, 28, 4, 0, 369, 60, 1, 0, 0, 0, 370, 371, 4, 29,
		12, 0, 371, 372, 5, 116, 0, 0, 372, 373, 5, 121, 0, 0, 373, 374, 5, 112,
		0, 0, 374, 375, 5, 101, 0, 0, 375, 376, 1, 0, 0, 0, 376, 377, 3, 3, 0,
		0, 377, 378, 1, 0, 0, 0, 378, 379, 6, 29, 5, 0, 379, 380, 6, 29, 6, 0,
		380, 62, 1, 0, 0, 0, 381, 382, 4, 30, 13, 0, 382, 383, 5, 110, 0, 0, 383,
		384, 5, 97, 0, 0, 384, 385, 5, 109, 0, 0, 385, 386, 5, 101, 0, 0, 386,
		387, 1, 0, 0, 0, 387, 388, 3, 3, 0, 0, 388, 389, 1, 0, 0, 0, 389, 390,
		6, 30, 7, 0, 390, 391, 6, 30, 6, 0, 391, 64, 1, 0, 0, 0, 392, 393, 4, 31,
		14, 0, 393, 394, 5, 100, 0, 0, 394, 395, 5, 101, 0, 0, 395, 396, 5, 115,
		0, 0, 396, 397, 5, 99, 0, 0, 397, 398, 5, 114, 0, 0, 398, 399, 5, 105,
		0, 0, 399, 400, 5, 112, 0, 0, 400, 401, 5, 116, 0, 0, 401, 402, 5, 105,
		0, 0, 402, 403, 5, 111, 0, 0, 403, 404, 5, 110, 0, 0, 404, 405, 1, 0, 0,
		0, 405, 406, 3, 3, 0, 0, 406, 407, 5, 124, 0, 0, 407, 408, 3, 5, 1, 0,
		408, 409, 1, 0, 0, 0, 409, 410, 6, 31, 8, 0, 410, 411, 6, 31, 9, 0, 411,
		66, 1, 0, 0, 0, 412, 413, 4, 32, 15, 0, 413, 414, 5, 100, 0, 0, 414, 415,
		5, 101, 0, 0, 415, 416, 5, 115, 0, 0, 416, 417, 5, 99, 0, 0, 417, 418,
		5, 114, 0, 0, 418, 419, 5, 105, 0, 0, 419, 420, 5, 112, 0, 0, 420, 421,
		5, 116, 0, 0, 421, 422, 5, 105, 0, 0, 422, 423, 5, 111, 0, 0, 423, 424,
		5, 110, 0, 0, 424, 425, 1, 0, 0, 0, 425, 426, 3, 3, 0, 0, 426, 427, 1,
		0, 0, 0, 427, 428, 6, 32, 10, 0, 428, 429, 6, 32, 6, 0, 429, 68, 1, 0,
		0, 0, 430, 431, 4, 33, 16, 0, 431, 432, 5, 116, 0, 0, 432, 433, 5, 97,
		0, 0, 433, 434, 5, 103, 0, 0, 434, 435, 5, 115, 0, 0, 435, 436, 1, 0, 0,
		0, 436, 437, 3, 3, 0, 0, 437, 438, 1, 0, 0, 0, 438, 439, 6, 33, 11, 0,
		439, 440, 6, 33, 6, 0, 440, 70, 1, 0, 0, 0, 441, 442, 4, 34, 17, 0, 442,
		443, 5, 111, 0, 0, 443, 444, 5, 119, 0, 0, 444, 445, 5, 110, 0, 0, 445,
		446, 5, 101, 0, 0, 446, 447, 5, 114, 0, 0, 447, 448, 1, 0, 0, 0, 448, 449,
		3, 3, 0, 0, 449, 450, 1, 0, 0, 0, 450, 451, 6, 34, 12, 0, 451, 452, 6,
		34, 6, 0, 452, 72, 1, 0, 0, 0, 453, 454, 4, 35, 18, 0, 454, 455, 5, 100,
		0, 0, 455, 456, 5, 101, 0, 0, 456, 457, 5, 115, 0, 0, 457, 458, 5, 116,
		0, 0, 458, 459, 5, 105, 0, 0, 459, 460, 5, 110, 0, 0, 460, 461, 5, 97,
		0, 0, 461, 462, 5, 116, 0, 0, 462, 463, 5, 105, 0, 0, 463, 464, 5, 111,
		0, 0, 464, 465, 5, 110, 0, 0, 465, 466, 1, 0, 0, 0, 466, 467, 3, 3, 0,
		0, 467, 468, 1, 0, 0, 0, 468, 469, 6, 35, 13, 0, 469, 470, 6, 35, 6, 0,
		470, 74, 1, 0, 0, 0, 471, 472, 4, 36, 19, 0, 472, 473, 5, 115, 0, 0, 473,
		474, 5, 99, 0, 0, 474, 475, 5, 104, 0, 0, 475, 476, 5, 101, 0, 0, 476,
		477, 5, 100, 0, 0, 477, 478, 5, 117, 0, 0, 478, 479, 5, 108, 0, 0, 479,
		480, 5, 101, 0, 0, 480, 481, 1, 0, 0, 0, 481, 482, 3, 3, 0, 0, 482, 483,
		1, 0, 0, 0, 483, 484, 6, 36, 14, 0, 484, 485, 6, 36, 6, 0, 485, 76, 1,
		0, 0, 0, 486, 487, 4, 37, 20, 0, 487, 488, 5, 117, 0, 0, 488, 489, 5, 110,
		0, 0, 489, 490, 5, 105, 0, 0, 490, 491, 5, 113, 0, 0, 491, 492, 5, 117,
		0, 0, 492, 493, 5, 101, 0, 0, 493, 494, 5, 95, 0, 0, 494, 495, 5, 107,
		0, 0, 495, 496, 5, 101, 0, 0, 496, 497, 5, 121, 0, 0, 497, 498, 1, 0, 0,
		0, 498, 499, 3, 3, 0, 0, 499, 500, 1, 0, 0, 0, 500, 501, 6, 37, 15, 0,
		501, 502, 6, 37, 6, 0, 502, 78, 1, 0, 0, 0, 503, 504, 4, 38, 21, 0, 504,
		505, 5, 115, 0, 0, 505, 506, 5, 111, 0, 0, 506, 507, 5, 117, 0, 0, 507,
		508, 5, 114, 0, 0, 508, 509, 5, 99, 0, 0, 509, 510, 5, 101, 0, 0, 510,
		511, 5, 115, 0, 0, 511, 512, 1, 0, 0, 0, 512, 513, 3, 3, 0, 0, 513, 514,
		3, 5, 1, 0, 514, 515, 1, 0, 0, 0, 515, 516, 6, 38, 16, 0, 516, 517, 6,
		38, 4, 0, 517, 80, 1, 0, 0, 0, 518, 519, 4, 39, 22, 0, 519, 520, 5, 112,
		0, 0, 520, 521, 5, 97, 0, 0, 521, 522, 5, 114, 0, 0, 522, 523, 5, 97, 0,
		0, 523, 524, 5, 109, 0, 0, 524, 525, 5, 115, 0, 0, 525, 526, 1, 0, 0, 0,
		526, 527, 3, 3, 0, 0, 527, 528, 3, 5, 1, 0, 528, 529, 1, 0, 0, 0, 529,
		530, 6, 39, 17, 0, 530, 531, 6, 39, 4, 0, 531, 82, 1, 0, 0, 0, 532, 533,
		4, 40, 23, 0, 533, 534, 5, 112, 0, 0, 534, 535, 5, 105, 0, 0, 535, 536,
		5, 112, 0, 0, 536, 537, 5, 101, 0, 0, 537, 538, 5, 108, 0, 0, 538, 539,
		5, 105, 0, 0, 539, 540, 5, 110, 0, 0, 540, 541, 5, 101, 0, 0, 541, 542,
		1, 0, 0, 0, 542, 543, 3, 3, 0, 0, 543, 544, 3, 5, 1, 0, 544, 545, 1, 0,
		0, 0, 545, 546, 6, 40, 18, 0, 546, 547, 6, 40, 9, 0, 547, 84, 1, 0, 0,
		0, 548, 550, 7, 0, 0, 0, 549, 548, 1, 0, 0, 0, 550, 551, 1, 0, 0, 0, 551,
		549, 1, 0, 0, 0, 551, 552, 1, 0, 0, 0, 552, 556, 1, 0, 0, 0, 553, 555,
		8, 1, 0, 0, 554, 553, 1, 0, 0, 0, 555, 558, 1, 0, 0, 0, 556, 554, 1, 0,
		0, 0, 556, 557, 1, 0, 0, 0, 557, 560, 1, 0, 0, 0, 558, 556, 1, 0, 0, 0,
		559, 561, 7, 1, 0, 0, 560, 559, 1, 0, 0, 0, 560, 561, 1, 0, 0, 0, 561,
		86, 1, 0, 0, 0, 562, 564, 7, 0, 0, 0, 563, 562, 1, 0, 0, 0, 564, 567, 1,
		0, 0, 0, 565, 563, 1, 0, 0, 0, 565, 566, 1, 0, 0, 0, 566, 568, 1, 0, 0,
		0, 567, 565, 1, 0, 0, 0, 568, 569, 7, 1, 0, 0, 569, 570, 1, 0, 0, 0, 570,
		571, 6, 42, 0, 0, 571, 88, 1, 0, 0, 0, 572, 573, 4, 43, 24, 0, 573, 577,
		5, 35, 0, 0, 574, 576, 8, 1, 0, 0, 575, 574, 1, 0, 0, 0, 576, 579, 1, 0,
		0, 0, 577, 575, 1, 0, 0, 0, 577, 578, 1, 0, 0, 0, 578, 581, 1, 0, 0, 0,
		579, 577, 1, 0, 0, 0, 580, 582, 7, 1, 0, 0, 581, 580, 1, 0, 0, 0, 581,
		582, 1, 0, 0, 0, 582, 583, 1, 0, 0, 0, 583, 584, 6, 43, 0, 0, 584, 90,
		1, 0, 0, 0, 24, 0, 1, 2, 95, 101, 110, 117, 125, 131, 138, 308, 325, 331,
		336, 342, 345, 351, 359, 551, 556, 560, 565, 577, 581, 19, 6, 0, 0, 5,
		1, 0, 5, 2, 0, 1, 27, 0, 4, 0, 0, 7, 4, 0, 2, 1, 0, 7, 5, 0, 7, 11, 0,
		2, 2, 0, 7, 12, 0, 7, 6, 0, 7, 7, 0, 7, 8, 0, 7, 9, 0, 7, 10, 0, 7, 13,
		0, 7, 14, 0, 7, 15, 0,
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

// PipeLangLexerInit initializes any static state used to implement PipeLangLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewPipeLangLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func PipeLangLexerInit() {
	staticData := &PipeLangLexerLexerStaticData
	staticData.once.Do(pipelanglexerLexerInit)
}

// NewPipeLangLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewPipeLangLexer(input antlr.CharStream) *PipeLangLexer {
	PipeLangLexerInit()
	l := new(PipeLangLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &PipeLangLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "PipeLangLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// PipeLangLexer tokens.
const (
	PipeLangLexerCOMMENT         = 1
	PipeLangLexerWS_BLANK        = 2
	PipeLangLexerWS              = 3
	PipeLangLexerTYPE            = 4
	PipeLangLexerNAME            = 5
	PipeLangLexerTAGS            = 6
	PipeLangLexerOWNER           = 7
	PipeLangLexerDESTINATION     = 8
	PipeLangLexerSCHEDULE        = 9
	PipeLangLexerUNIQUE_KEY      = 10
	PipeLangLexerDESCRIPTION_ML  = 11
	PipeLangLexerDESCRIPTION     = 12
	PipeLangLexerSOURCES         = 13
	PipeLangLexerPARAMS          = 14
	PipeLangLexerPIPELINE        = 15
	PipeLangLexerIDENTIFIER      = 16
	PipeLangLexerCOLON_TOK       = 17
	PipeLangLexerLBRACE          = 18
	PipeLangLexerRBRACE          = 19
	PipeLangLexerCOMMA           = 20
	PipeLangLexerDASH            = 21
	PipeLangLexerSTRING_LIT      = 22
	PipeLangLexerNUMBER          = 23
	PipeLangLexerINLINE_CMT      = 24
	PipeLangLexerVALUE           = 25
	PipeLangLexerVALUE_NL        = 26
	PipeLangLexerSECTION_LINE    = 27
	PipeLangLexerSECTION_BLANK   = 28
	PipeLangLexerSECTION_COMMENT = 29
)

// PipeLangLexer modes.
const (
	PipeLangLexerVALUE_MODE = iota + 1
	PipeLangLexerSECTION_MODE
)

func (l *PipeLangLexer) cleanValue() {
	t := l.GetText()
	s, e := 0, len(t)
	for s < e && (t[s] == ' ' || t[s] == '\t') {
		s++
	}
	for e > s && (t[e-1] == ' ' || t[e-1] == '\t') {
		e--
	}
	l.SetText(t[s:e])
}

func (l *PipeLangLexer) Action(localctx antlr.RuleContext, ruleIndex, actionIndex int) {
	switch ruleIndex {
	case 27:
		l.VALUE_Action(localctx, actionIndex)

	default:
		panic("No registered action for: " + fmt.Sprint(ruleIndex))
	}
}

func (l *PipeLangLexer) VALUE_Action(localctx antlr.RuleContext, actionIndex int) {
	switch actionIndex {
	case 0:
		l.cleanValue()

	default:
		panic("No registered action for: " + fmt.Sprint(actionIndex))
	}
}

func (l *PipeLangLexer) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 6:
		return l.TYPE_Sempred(localctx, predIndex)

	case 7:
		return l.NAME_Sempred(localctx, predIndex)

	case 8:
		return l.TAGS_Sempred(localctx, predIndex)

	case 9:
		return l.OWNER_Sempred(localctx, predIndex)

	case 10:
		return l.DESTINATION_Sempred(localctx, predIndex)

	case 11:
		return l.SCHEDULE_Sempred(localctx, predIndex)

	case 12:
		return l.UNIQUE_KEY_Sempred(localctx, predIndex)

	case 13:
		return l.DESCRIPTION_ML_Sempred(localctx, predIndex)

	case 14:
		return l.DESCRIPTION_Sempred(localctx, predIndex)

	case 15:
		return l.SOURCES_Sempred(localctx, predIndex)

	case 16:
		return l.PARAMS_Sempred(localctx, predIndex)

	case 17:
		return l.PIPELINE_Sempred(localctx, predIndex)

	case 29:
		return l.SEC_TYPE_Sempred(localctx, predIndex)

	case 30:
		return l.SEC_NAME_Sempred(localctx, predIndex)

	case 31:
		return l.SEC_DESC_ML_Sempred(localctx, predIndex)

	case 32:
		return l.SEC_DESC_Sempred(localctx, predIndex)

	case 33:
		return l.SEC_TAGS_Sempred(localctx, predIndex)

	case 34:
		return l.SEC_OWNER_Sempred(localctx, predIndex)

	case 35:
		return l.SEC_DESTINATION_Sempred(localctx, predIndex)

	case 36:
		return l.SEC_SCHEDULE_Sempred(localctx, predIndex)

	case 37:
		return l.SEC_UNIQUE_KEY_Sempred(localctx, predIndex)

	case 38:
		return l.SEC_SOURCES_Sempred(localctx, predIndex)

	case 39:
		return l.SEC_PARAMS_Sempred(localctx, predIndex)

	case 40:
		return l.SEC_PIPELINE_Sempred(localctx, predIndex)

	case 43:
		return l.SECTION_COMMENT_Sempred(localctx, predIndex)

	default:
		panic("No registered predicate for: " + fmt.Sprint(ruleIndex))
	}
}

func (p *PipeLangLexer) TYPE_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.GetCharPositionInLine() == 0

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *PipeLangLexer) NAME_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.GetCharPositionInLine() == 0

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *PipeLangLexer) TAGS_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 2:
		return p.GetCharPositionInLine() == 0

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *PipeLangLexer) OWNER_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 3:
		return p.GetCharPositionInLine() == 0

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *PipeLangLexer) DESTINATION_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 4:
		return p.GetCharPositionInLine() == 0

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *PipeLangLexer) SCHEDULE_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 5:
		return p.GetCharPositionInLine() == 0

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *PipeLangLexer) UNIQUE_KEY_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 6:
		return p.GetCharPositionInLine() == 0

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *PipeLangLexer) DESCRIPTION_ML_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 7:
		return p.GetCharPositionInLine() == 0

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *PipeLangLexer) DESCRIPTION_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 8:
		return p.GetCharPositionInLine() == 0

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *PipeLangLexer) SOURCES_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 9:
		return p.GetCharPositionInLine() == 0

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *PipeLangLexer) PARAMS_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 10:
		return p.GetCharPositionInLine() == 0

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *PipeLangLexer) PIPELINE_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 11:
		return p.GetCharPositionInLine() == 0

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *PipeLangLexer) SEC_TYPE_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 12:
		return p.GetCharPositionInLine() == 0

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *PipeLangLexer) SEC_NAME_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 13:
		return p.GetCharPositionInLine() == 0

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *PipeLangLexer) SEC_DESC_ML_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 14:
		return p.GetCharPositionInLine() == 0

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *PipeLangLexer) SEC_DESC_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 15:
		return p.GetCharPositionInLine() == 0

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *PipeLangLexer) SEC_TAGS_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 16:
		return p.GetCharPositionInLine() == 0

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *PipeLangLexer) SEC_OWNER_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 17:
		return p.GetCharPositionInLine() == 0

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *PipeLangLexer) SEC_DESTINATION_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 18:
		return p.GetCharPositionInLine() == 0

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *PipeLangLexer) SEC_SCHEDULE_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 19:
		return p.GetCharPositionInLine() == 0

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *PipeLangLexer) SEC_UNIQUE_KEY_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 20:
		return p.GetCharPositionInLine() == 0

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *PipeLangLexer) SEC_SOURCES_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 21:
		return p.GetCharPositionInLine() == 0

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *PipeLangLexer) SEC_PARAMS_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 22:
		return p.GetCharPositionInLine() == 0

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *PipeLangLexer) SEC_PIPELINE_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 23:
		return p.GetCharPositionInLine() == 0

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *PipeLangLexer) SECTION_COMMENT_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 24:
		return p.GetCharPositionInLine() == 0

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
