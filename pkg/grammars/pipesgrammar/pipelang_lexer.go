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
		"DEFAULT_MODE", "DIRECTIVE_VALUE_MODE", "SQL_BODY_MODE",
	}
	staticData.LiteralNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "' '",
	}
	staticData.SymbolicNames = []string{
		"", "COMMENT", "WS_BLANK", "WS", "KW_DESCRIPTION", "KW_TAGS", "KW_TYPE",
		"KW_NODE", "KW_COPY_SCHEDULE", "KW_DATASOURCE", "KW_TARGET_DATASOURCE",
		"KW_SQL_ARROW", "SQL_LINE", "REST_OF_LINE",
	}
	staticData.RuleNames = []string{
		"COMMENT", "WS_BLANK", "WS", "KW_DESCRIPTION", "KW_TAGS", "KW_TYPE",
		"KW_NODE", "KW_COPY_SCHEDULE", "KW_DATASOURCE", "KW_TARGET_DATASOURCE",
		"KW_SQL_ARROW", "SQL_LINE", "REST_OF_LINE", "SQL_KW_DESCRIPTION", "SQL_KW_TAGS",
		"SQL_KW_TYPE", "SQL_KW_NODE", "SQL_KW_DATASOURCE", "SQL_KW_TARGET_DATASOURCE",
		"SQL_KW_COPY_SCHEDULE", "SQL_CHAR", "SQL_NL",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 13, 372, 6, -1, 6, -1, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2,
		2, 3, 7, 3, 2, 4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8,
		2, 9, 7, 9, 2, 10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2,
		14, 7, 14, 2, 15, 7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19,
		7, 19, 2, 20, 7, 20, 2, 21, 7, 21, 1, 0, 1, 0, 5, 0, 50, 8, 0, 10, 0, 12,
		0, 53, 9, 0, 1, 0, 1, 0, 1, 1, 5, 1, 58, 8, 1, 10, 1, 12, 1, 61, 9, 1,
		1, 1, 4, 1, 64, 8, 1, 11, 1, 12, 1, 65, 1, 1, 1, 1, 1, 2, 4, 2, 71, 8,
		2, 11, 2, 12, 2, 72, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 4, 3, 90, 8, 3, 11, 3, 12, 3, 91,
		1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 4, 4, 102, 8, 4, 11, 4,
		12, 4, 103, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 4, 5, 114,
		8, 5, 11, 5, 12, 5, 115, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		4, 6, 126, 8, 6, 11, 6, 12, 6, 127, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 4, 7,
		147, 8, 7, 11, 7, 12, 7, 148, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 4, 8, 165, 8, 8, 11, 8, 12, 8,
		166, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 4, 9, 190,
		8, 9, 11, 9, 12, 9, 191, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		5, 10, 201, 8, 10, 10, 10, 12, 10, 204, 9, 10, 1, 10, 1, 10, 5, 10, 208,
		8, 10, 10, 10, 12, 10, 211, 9, 10, 1, 10, 4, 10, 214, 8, 10, 11, 10, 12,
		10, 215, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 5, 12, 223, 8, 12, 10, 12,
		12, 12, 226, 9, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1,
		13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13,
		4, 13, 246, 8, 13, 11, 13, 12, 13, 247, 1, 13, 1, 13, 1, 13, 1, 14, 1,
		14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 4, 14, 260, 8, 14, 11, 14, 12, 14,
		261, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1,
		15, 4, 15, 274, 8, 15, 11, 15, 12, 15, 275, 1, 15, 1, 15, 1, 15, 1, 16,
		1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 4, 16, 288, 8, 16, 11, 16, 12,
		16, 289, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17,
		1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 4, 17, 308, 8, 17, 11,
		17, 12, 17, 309, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18,
		1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1,
		18, 1, 18, 1, 18, 1, 18, 1, 18, 4, 18, 335, 8, 18, 11, 18, 12, 18, 336,
		1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1,
		19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 4, 19, 358,
		8, 19, 11, 19, 12, 19, 359, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1,
		20, 1, 21, 1, 21, 1, 21, 1, 21, 0, 0, 22, 3, 1, 5, 2, 7, 3, 9, 4, 11, 5,
		13, 6, 15, 7, 17, 8, 19, 9, 21, 10, 23, 11, 25, 12, 27, 13, 29, 0, 31,
		0, 33, 0, 35, 0, 37, 0, 39, 0, 41, 0, 43, 0, 45, 0, 3, 0, 1, 2, 2, 2, 0,
		10, 10, 13, 13, 2, 0, 9, 9, 32, 32, 391, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0,
		0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1,
		0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21,
		1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 1, 27, 1, 0, 0, 0, 2,
		29, 1, 0, 0, 0, 2, 31, 1, 0, 0, 0, 2, 33, 1, 0, 0, 0, 2, 35, 1, 0, 0, 0,
		2, 37, 1, 0, 0, 0, 2, 39, 1, 0, 0, 0, 2, 41, 1, 0, 0, 0, 2, 43, 1, 0, 0,
		0, 2, 45, 1, 0, 0, 0, 3, 47, 1, 0, 0, 0, 5, 59, 1, 0, 0, 0, 7, 70, 1, 0,
		0, 0, 9, 76, 1, 0, 0, 0, 11, 95, 1, 0, 0, 0, 13, 107, 1, 0, 0, 0, 15, 119,
		1, 0, 0, 0, 17, 131, 1, 0, 0, 0, 19, 152, 1, 0, 0, 0, 21, 170, 1, 0, 0,
		0, 23, 195, 1, 0, 0, 0, 25, 219, 1, 0, 0, 0, 27, 224, 1, 0, 0, 0, 29, 231,
		1, 0, 0, 0, 31, 252, 1, 0, 0, 0, 33, 266, 1, 0, 0, 0, 35, 280, 1, 0, 0,
		0, 37, 294, 1, 0, 0, 0, 39, 314, 1, 0, 0, 0, 41, 341, 1, 0, 0, 0, 43, 364,
		1, 0, 0, 0, 45, 368, 1, 0, 0, 0, 47, 51, 5, 35, 0, 0, 48, 50, 8, 0, 0,
		0, 49, 48, 1, 0, 0, 0, 50, 53, 1, 0, 0, 0, 51, 49, 1, 0, 0, 0, 51, 52,
		1, 0, 0, 0, 52, 54, 1, 0, 0, 0, 53, 51, 1, 0, 0, 0, 54, 55, 6, 0, 0, 0,
		55, 4, 1, 0, 0, 0, 56, 58, 7, 1, 0, 0, 57, 56, 1, 0, 0, 0, 58, 61, 1, 0,
		0, 0, 59, 57, 1, 0, 0, 0, 59, 60, 1, 0, 0, 0, 60, 63, 1, 0, 0, 0, 61, 59,
		1, 0, 0, 0, 62, 64, 7, 0, 0, 0, 63, 62, 1, 0, 0, 0, 64, 65, 1, 0, 0, 0,
		65, 63, 1, 0, 0, 0, 65, 66, 1, 0, 0, 0, 66, 67, 1, 0, 0, 0, 67, 68, 6,
		1, 0, 0, 68, 6, 1, 0, 0, 0, 69, 71, 7, 1, 0, 0, 70, 69, 1, 0, 0, 0, 71,
		72, 1, 0, 0, 0, 72, 70, 1, 0, 0, 0, 72, 73, 1, 0, 0, 0, 73, 74, 1, 0, 0,
		0, 74, 75, 6, 2, 0, 0, 75, 8, 1, 0, 0, 0, 76, 77, 5, 68, 0, 0, 77, 78,
		5, 69, 0, 0, 78, 79, 5, 83, 0, 0, 79, 80, 5, 67, 0, 0, 80, 81, 5, 82, 0,
		0, 81, 82, 5, 73, 0, 0, 82, 83, 5, 80, 0, 0, 83, 84, 5, 84, 0, 0, 84, 85,
		5, 73, 0, 0, 85, 86, 5, 79, 0, 0, 86, 87, 5, 78, 0, 0, 87, 89, 1, 0, 0,
		0, 88, 90, 7, 1, 0, 0, 89, 88, 1, 0, 0, 0, 90, 91, 1, 0, 0, 0, 91, 89,
		1, 0, 0, 0, 91, 92, 1, 0, 0, 0, 92, 93, 1, 0, 0, 0, 93, 94, 6, 3, 1, 0,
		94, 10, 1, 0, 0, 0, 95, 96, 5, 84, 0, 0, 96, 97, 5, 65, 0, 0, 97, 98, 5,
		71, 0, 0, 98, 99, 5, 83, 0, 0, 99, 101, 1, 0, 0, 0, 100, 102, 7, 1, 0,
		0, 101, 100, 1, 0, 0, 0, 102, 103, 1, 0, 0, 0, 103, 101, 1, 0, 0, 0, 103,
		104, 1, 0, 0, 0, 104, 105, 1, 0, 0, 0, 105, 106, 6, 4, 1, 0, 106, 12, 1,
		0, 0, 0, 107, 108, 5, 84, 0, 0, 108, 109, 5, 89, 0, 0, 109, 110, 5, 80,
		0, 0, 110, 111, 5, 69, 0, 0, 111, 113, 1, 0, 0, 0, 112, 114, 7, 1, 0, 0,
		113, 112, 1, 0, 0, 0, 114, 115, 1, 0, 0, 0, 115, 113, 1, 0, 0, 0, 115,
		116, 1, 0, 0, 0, 116, 117, 1, 0, 0, 0, 117, 118, 6, 5, 1, 0, 118, 14, 1,
		0, 0, 0, 119, 120, 5, 78, 0, 0, 120, 121, 5, 79, 0, 0, 121, 122, 5, 68,
		0, 0, 122, 123, 5, 69, 0, 0, 123, 125, 1, 0, 0, 0, 124, 126, 7, 1, 0, 0,
		125, 124, 1, 0, 0, 0, 126, 127, 1, 0, 0, 0, 127, 125, 1, 0, 0, 0, 127,
		128, 1, 0, 0, 0, 128, 129, 1, 0, 0, 0, 129, 130, 6, 6, 1, 0, 130, 16, 1,
		0, 0, 0, 131, 132, 5, 67, 0, 0, 132, 133, 5, 79, 0, 0, 133, 134, 5, 80,
		0, 0, 134, 135, 5, 89, 0, 0, 135, 136, 5, 95, 0, 0, 136, 137, 5, 83, 0,
		0, 137, 138, 5, 67, 0, 0, 138, 139, 5, 72, 0, 0, 139, 140, 5, 69, 0, 0,
		140, 141, 5, 68, 0, 0, 141, 142, 5, 85, 0, 0, 142, 143, 5, 76, 0, 0, 143,
		144, 5, 69, 0, 0, 144, 146, 1, 0, 0, 0, 145, 147, 7, 1, 0, 0, 146, 145,
		1, 0, 0, 0, 147, 148, 1, 0, 0, 0, 148, 146, 1, 0, 0, 0, 148, 149, 1, 0,
		0, 0, 149, 150, 1, 0, 0, 0, 150, 151, 6, 7, 1, 0, 151, 18, 1, 0, 0, 0,
		152, 153, 5, 68, 0, 0, 153, 154, 5, 65, 0, 0, 154, 155, 5, 84, 0, 0, 155,
		156, 5, 65, 0, 0, 156, 157, 5, 83, 0, 0, 157, 158, 5, 79, 0, 0, 158, 159,
		5, 85, 0, 0, 159, 160, 5, 82, 0, 0, 160, 161, 5, 67, 0, 0, 161, 162, 5,
		69, 0, 0, 162, 164, 1, 0, 0, 0, 163, 165, 7, 1, 0, 0, 164, 163, 1, 0, 0,
		0, 165, 166, 1, 0, 0, 0, 166, 164, 1, 0, 0, 0, 166, 167, 1, 0, 0, 0, 167,
		168, 1, 0, 0, 0, 168, 169, 6, 8, 1, 0, 169, 20, 1, 0, 0, 0, 170, 171, 5,
		84, 0, 0, 171, 172, 5, 65, 0, 0, 172, 173, 5, 82, 0, 0, 173, 174, 5, 71,
		0, 0, 174, 175, 5, 69, 0, 0, 175, 176, 5, 84, 0, 0, 176, 177, 5, 95, 0,
		0, 177, 178, 5, 68, 0, 0, 178, 179, 5, 65, 0, 0, 179, 180, 5, 84, 0, 0,
		180, 181, 5, 65, 0, 0, 181, 182, 5, 83, 0, 0, 182, 183, 5, 79, 0, 0, 183,
		184, 5, 85, 0, 0, 184, 185, 5, 82, 0, 0, 185, 186, 5, 67, 0, 0, 186, 187,
		5, 69, 0, 0, 187, 189, 1, 0, 0, 0, 188, 190, 7, 1, 0, 0, 189, 188, 1, 0,
		0, 0, 190, 191, 1, 0, 0, 0, 191, 189, 1, 0, 0, 0, 191, 192, 1, 0, 0, 0,
		192, 193, 1, 0, 0, 0, 193, 194, 6, 9, 1, 0, 194, 22, 1, 0, 0, 0, 195, 196,
		5, 83, 0, 0, 196, 197, 5, 81, 0, 0, 197, 198, 5, 76, 0, 0, 198, 202, 1,
		0, 0, 0, 199, 201, 7, 1, 0, 0, 200, 199, 1, 0, 0, 0, 201, 204, 1, 0, 0,
		0, 202, 200, 1, 0, 0, 0, 202, 203, 1, 0, 0, 0, 203, 205, 1, 0, 0, 0, 204,
		202, 1, 0, 0, 0, 205, 209, 5, 62, 0, 0, 206, 208, 8, 0, 0, 0, 207, 206,
		1, 0, 0, 0, 208, 211, 1, 0, 0, 0, 209, 207, 1, 0, 0, 0, 209, 210, 1, 0,
		0, 0, 210, 213, 1, 0, 0, 0, 211, 209, 1, 0, 0, 0, 212, 214, 7, 0, 0, 0,
		213, 212, 1, 0, 0, 0, 214, 215, 1, 0, 0, 0, 215, 213, 1, 0, 0, 0, 215,
		216, 1, 0, 0, 0, 216, 217, 1, 0, 0, 0, 217, 218, 6, 10, 2, 0, 218, 24,
		1, 0, 0, 0, 219, 220, 5, 32, 0, 0, 220, 26, 1, 0, 0, 0, 221, 223, 8, 0,
		0, 0, 222, 221, 1, 0, 0, 0, 223, 226, 1, 0, 0, 0, 224, 222, 1, 0, 0, 0,
		224, 225, 1, 0, 0, 0, 225, 227, 1, 0, 0, 0, 226, 224, 1, 0, 0, 0, 227,
		228, 7, 0, 0, 0, 228, 229, 1, 0, 0, 0, 229, 230, 6, 12, 3, 0, 230, 28,
		1, 0, 0, 0, 231, 232, 4, 13, 0, 0, 232, 233, 5, 68, 0, 0, 233, 234, 5,
		69, 0, 0, 234, 235, 5, 83, 0, 0, 235, 236, 5, 67, 0, 0, 236, 237, 5, 82,
		0, 0, 237, 238, 5, 73, 0, 0, 238, 239, 5, 80, 0, 0, 239, 240, 5, 84, 0,
		0, 240, 241, 5, 73, 0, 0, 241, 242, 5, 79, 0, 0, 242, 243, 5, 78, 0, 0,
		243, 245, 1, 0, 0, 0, 244, 246, 7, 1, 0, 0, 245, 244, 1, 0, 0, 0, 246,
		247, 1, 0, 0, 0, 247, 245, 1, 0, 0, 0, 247, 248, 1, 0, 0, 0, 248, 249,
		1, 0, 0, 0, 249, 250, 6, 13, 4, 0, 250, 251, 6, 13, 5, 0, 251, 30, 1, 0,
		0, 0, 252, 253, 4, 14, 1, 0, 253, 254, 5, 84, 0, 0, 254, 255, 5, 65, 0,
		0, 255, 256, 5, 71, 0, 0, 256, 257, 5, 83, 0, 0, 257, 259, 1, 0, 0, 0,
		258, 260, 7, 1, 0, 0, 259, 258, 1, 0, 0, 0, 260, 261, 1, 0, 0, 0, 261,
		259, 1, 0, 0, 0, 261, 262, 1, 0, 0, 0, 262, 263, 1, 0, 0, 0, 263, 264,
		6, 14, 6, 0, 264, 265, 6, 14, 5, 0, 265, 32, 1, 0, 0, 0, 266, 267, 4, 15,
		2, 0, 267, 268, 5, 84, 0, 0, 268, 269, 5, 89, 0, 0, 269, 270, 5, 80, 0,
		0, 270, 271, 5, 69, 0, 0, 271, 273, 1, 0, 0, 0, 272, 274, 7, 1, 0, 0, 273,
		272, 1, 0, 0, 0, 274, 275, 1, 0, 0, 0, 275, 273, 1, 0, 0, 0, 275, 276,
		1, 0, 0, 0, 276, 277, 1, 0, 0, 0, 277, 278, 6, 15, 7, 0, 278, 279, 6, 15,
		5, 0, 279, 34, 1, 0, 0, 0, 280, 281, 4, 16, 3, 0, 281, 282, 5, 78, 0, 0,
		282, 283, 5, 79, 0, 0, 283, 284, 5, 68, 0, 0, 284, 285, 5, 69, 0, 0, 285,
		287, 1, 0, 0, 0, 286, 288, 7, 1, 0, 0, 287, 286, 1, 0, 0, 0, 288, 289,
		1, 0, 0, 0, 289, 287, 1, 0, 0, 0, 289, 290, 1, 0, 0, 0, 290, 291, 1, 0,
		0, 0, 291, 292, 6, 16, 8, 0, 292, 293, 6, 16, 5, 0, 293, 36, 1, 0, 0, 0,
		294, 295, 4, 17, 4, 0, 295, 296, 5, 68, 0, 0, 296, 297, 5, 65, 0, 0, 297,
		298, 5, 84, 0, 0, 298, 299, 5, 65, 0, 0, 299, 300, 5, 83, 0, 0, 300, 301,
		5, 79, 0, 0, 301, 302, 5, 85, 0, 0, 302, 303, 5, 82, 0, 0, 303, 304, 5,
		67, 0, 0, 304, 305, 5, 69, 0, 0, 305, 307, 1, 0, 0, 0, 306, 308, 7, 1,
		0, 0, 307, 306, 1, 0, 0, 0, 308, 309, 1, 0, 0, 0, 309, 307, 1, 0, 0, 0,
		309, 310, 1, 0, 0, 0, 310, 311, 1, 0, 0, 0, 311, 312, 6, 17, 9, 0, 312,
		313, 6, 17, 5, 0, 313, 38, 1, 0, 0, 0, 314, 315, 4, 18, 5, 0, 315, 316,
		5, 84, 0, 0, 316, 317, 5, 65, 0, 0, 317, 318, 5, 82, 0, 0, 318, 319, 5,
		71, 0, 0, 319, 320, 5, 69, 0, 0, 320, 321, 5, 84, 0, 0, 321, 322, 5, 95,
		0, 0, 322, 323, 5, 68, 0, 0, 323, 324, 5, 65, 0, 0, 324, 325, 5, 84, 0,
		0, 325, 326, 5, 65, 0, 0, 326, 327, 5, 83, 0, 0, 327, 328, 5, 79, 0, 0,
		328, 329, 5, 85, 0, 0, 329, 330, 5, 82, 0, 0, 330, 331, 5, 67, 0, 0, 331,
		332, 5, 69, 0, 0, 332, 334, 1, 0, 0, 0, 333, 335, 7, 1, 0, 0, 334, 333,
		1, 0, 0, 0, 335, 336, 1, 0, 0, 0, 336, 334, 1, 0, 0, 0, 336, 337, 1, 0,
		0, 0, 337, 338, 1, 0, 0, 0, 338, 339, 6, 18, 10, 0, 339, 340, 6, 18, 5,
		0, 340, 40, 1, 0, 0, 0, 341, 342, 4, 19, 6, 0, 342, 343, 5, 67, 0, 0, 343,
		344, 5, 79, 0, 0, 344, 345, 5, 80, 0, 0, 345, 346, 5, 89, 0, 0, 346, 347,
		5, 95, 0, 0, 347, 348, 5, 83, 0, 0, 348, 349, 5, 67, 0, 0, 349, 350, 5,
		72, 0, 0, 350, 351, 5, 69, 0, 0, 351, 352, 5, 68, 0, 0, 352, 353, 5, 85,
		0, 0, 353, 354, 5, 76, 0, 0, 354, 355, 5, 69, 0, 0, 355, 357, 1, 0, 0,
		0, 356, 358, 7, 1, 0, 0, 357, 356, 1, 0, 0, 0, 358, 359, 1, 0, 0, 0, 359,
		357, 1, 0, 0, 0, 359, 360, 1, 0, 0, 0, 360, 361, 1, 0, 0, 0, 361, 362,
		6, 19, 11, 0, 362, 363, 6, 19, 5, 0, 363, 42, 1, 0, 0, 0, 364, 365, 8,
		0, 0, 0, 365, 366, 1, 0, 0, 0, 366, 367, 6, 20, 12, 0, 367, 44, 1, 0, 0,
		0, 368, 369, 7, 0, 0, 0, 369, 370, 1, 0, 0, 0, 370, 371, 6, 21, 13, 0,
		371, 46, 1, 0, 0, 0, 25, 0, 1, 2, 51, 59, 65, 72, 91, 103, 115, 127, 148,
		166, 191, 202, 209, 215, 224, 247, 261, 275, 289, 309, 336, 359, 14, 6,
		0, 0, 5, 1, 0, 5, 2, 0, 4, 0, 0, 7, 4, 0, 2, 1, 0, 7, 5, 0, 7, 6, 0, 7,
		7, 0, 7, 9, 0, 7, 10, 0, 7, 8, 0, 3, 0, 0, 7, 12, 0,
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
	PipeLangLexerCOMMENT              = 1
	PipeLangLexerWS_BLANK             = 2
	PipeLangLexerWS                   = 3
	PipeLangLexerKW_DESCRIPTION       = 4
	PipeLangLexerKW_TAGS              = 5
	PipeLangLexerKW_TYPE              = 6
	PipeLangLexerKW_NODE              = 7
	PipeLangLexerKW_COPY_SCHEDULE     = 8
	PipeLangLexerKW_DATASOURCE        = 9
	PipeLangLexerKW_TARGET_DATASOURCE = 10
	PipeLangLexerKW_SQL_ARROW         = 11
	PipeLangLexerSQL_LINE             = 12
	PipeLangLexerREST_OF_LINE         = 13
)

// PipeLangLexer modes.
const (
	PipeLangLexerDIRECTIVE_VALUE_MODE = iota + 1
	PipeLangLexerSQL_BODY_MODE
)

func (l *PipeLangLexer) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 13:
		return l.SQL_KW_DESCRIPTION_Sempred(localctx, predIndex)

	case 14:
		return l.SQL_KW_TAGS_Sempred(localctx, predIndex)

	case 15:
		return l.SQL_KW_TYPE_Sempred(localctx, predIndex)

	case 16:
		return l.SQL_KW_NODE_Sempred(localctx, predIndex)

	case 17:
		return l.SQL_KW_DATASOURCE_Sempred(localctx, predIndex)

	case 18:
		return l.SQL_KW_TARGET_DATASOURCE_Sempred(localctx, predIndex)

	case 19:
		return l.SQL_KW_COPY_SCHEDULE_Sempred(localctx, predIndex)

	default:
		panic("No registered predicate for: " + fmt.Sprint(ruleIndex))
	}
}

func (p *PipeLangLexer) SQL_KW_DESCRIPTION_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.GetCharPositionInLine() == 0

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *PipeLangLexer) SQL_KW_TAGS_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.GetCharPositionInLine() == 0

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *PipeLangLexer) SQL_KW_TYPE_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 2:
		return p.GetCharPositionInLine() == 0

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *PipeLangLexer) SQL_KW_NODE_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 3:
		return p.GetCharPositionInLine() == 0

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *PipeLangLexer) SQL_KW_DATASOURCE_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 4:
		return p.GetCharPositionInLine() == 0

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *PipeLangLexer) SQL_KW_TARGET_DATASOURCE_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 5:
		return p.GetCharPositionInLine() == 0

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *PipeLangLexer) SQL_KW_COPY_SCHEDULE_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 6:
		return p.GetCharPositionInLine() == 0

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
