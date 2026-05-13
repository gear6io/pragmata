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
		"DEFAULT_MODE", "VALUE_MODE", "BLOCK_MODE", "PIPELINE_MODE",
	}
	staticData.SymbolicNames = []string{
		"", "COMMENT", "WS_BLANK", "WS", "TYPE", "NAME", "TAGS", "OWNER", "DESTINATION",
		"SCHEDULE", "DESCRIPTION_ML", "DESCRIPTION", "SOURCES", "PARAMS", "PIPELINE",
		"REST_OF_LINE", "BLOCK_LINE", "BLOCK_BLANK", "BLOCK_COMMENT", "NODE_HEADER",
		"PRQL_LINE", "PIPE_BLANK", "PIPE_COMMENT",
	}
	staticData.RuleNames = []string{
		"COLON", "REST", "ID", "COMMENT", "WS_BLANK", "WS", "TYPE", "NAME",
		"TAGS", "OWNER", "DESTINATION", "SCHEDULE", "DESCRIPTION_ML", "DESCRIPTION",
		"SOURCES", "PARAMS", "PIPELINE", "REST_OF_LINE", "BLOCK_TYPE", "BLOCK_NAME",
		"BLOCK_DESC_ML", "BLOCK_DESC", "BLOCK_TAGS", "BLOCK_OWNER", "BLOCK_DESTINATION",
		"BLOCK_SCHEDULE", "BLOCK_SOURCES", "BLOCK_PARAMS", "BLOCK_PIPELINE",
		"BLOCK_LINE", "BLOCK_BLANK", "BLOCK_COMMENT", "PIPE_TYPE", "PIPE_NAME",
		"PIPE_DESC_ML", "PIPE_DESC", "PIPE_TAGS", "PIPE_OWNER", "PIPE_DESTINATION",
		"PIPE_SCHEDULE", "PIPE_SOURCES", "PIPE_PARAMS", "PIPE_PIPELINE", "NODE_HEADER",
		"PRQL_LINE", "PIPE_BLANK", "PIPE_COMMENT",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 22, 709, 6, -1, 6, -1, 6, -1, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2,
		7, 2, 2, 3, 7, 3, 2, 4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8,
		7, 8, 2, 9, 7, 9, 2, 10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13,
		2, 14, 7, 14, 2, 15, 7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2,
		19, 7, 19, 2, 20, 7, 20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24,
		7, 24, 2, 25, 7, 25, 2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7,
		29, 2, 30, 7, 30, 2, 31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34,
		2, 35, 7, 35, 2, 36, 7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2,
		40, 7, 40, 2, 41, 7, 41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45,
		7, 45, 2, 46, 7, 46, 1, 0, 1, 0, 5, 0, 101, 8, 0, 10, 0, 12, 0, 104, 9,
		0, 1, 1, 5, 1, 107, 8, 1, 10, 1, 12, 1, 110, 9, 1, 1, 1, 1, 1, 1, 2, 1,
		2, 5, 2, 116, 8, 2, 10, 2, 12, 2, 119, 9, 2, 1, 3, 1, 3, 5, 3, 123, 8,
		3, 10, 3, 12, 3, 126, 9, 3, 1, 3, 1, 3, 1, 4, 5, 4, 131, 8, 4, 10, 4, 12,
		4, 134, 9, 4, 1, 4, 4, 4, 137, 8, 4, 11, 4, 12, 4, 138, 1, 4, 1, 4, 1,
		5, 4, 5, 144, 8, 5, 11, 5, 12, 5, 145, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1,
		7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1,
		9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1,
		10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11,
		1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1,
		12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12,
		1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1,
		13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14,
		1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1,
		14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15,
		1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1,
		16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 5, 17, 290, 8, 17, 10, 17,
		12, 17, 293, 9, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1,
		18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19,
		1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1,
		20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20,
		1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1,
		21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21,
		1, 21, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1,
		22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23,
		1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1,
		24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24,
		1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1,
		25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26,
		1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1,
		26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27,
		1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1,
		28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28,
		1, 29, 4, 29, 461, 8, 29, 11, 29, 12, 29, 462, 1, 29, 5, 29, 466, 8, 29,
		10, 29, 12, 29, 469, 9, 29, 1, 29, 1, 29, 1, 30, 5, 30, 474, 8, 30, 10,
		30, 12, 30, 477, 9, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 31, 1, 31, 1, 31,
		5, 31, 486, 8, 31, 10, 31, 12, 31, 489, 9, 31, 1, 31, 3, 31, 492, 8, 31,
		1, 31, 1, 31, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1,
		32, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33,
		1, 33, 1, 33, 1, 33, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1,
		34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34,
		1, 34, 1, 34, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1,
		35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 36,
		1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1,
		37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37,
		1, 37, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1,
		38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 39, 1, 39,
		1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1,
		39, 1, 39, 1, 39, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40,
		1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 41, 1, 41, 1, 41, 1,
		41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41,
		1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1,
		42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 43, 4, 43, 658, 8, 43, 11, 43,
		12, 43, 659, 1, 43, 1, 43, 1, 43, 1, 43, 5, 43, 666, 8, 43, 10, 43, 12,
		43, 669, 9, 43, 1, 43, 3, 43, 672, 8, 43, 1, 44, 4, 44, 675, 8, 44, 11,
		44, 12, 44, 676, 1, 44, 4, 44, 680, 8, 44, 11, 44, 12, 44, 681, 1, 44,
		3, 44, 685, 8, 44, 1, 45, 5, 45, 688, 8, 45, 10, 45, 12, 45, 691, 9, 45,
		1, 45, 1, 45, 1, 45, 1, 45, 1, 46, 1, 46, 1, 46, 5, 46, 700, 8, 46, 10,
		46, 12, 46, 703, 9, 46, 1, 46, 3, 46, 706, 8, 46, 1, 46, 1, 46, 0, 0, 47,
		4, 0, 6, 0, 8, 0, 10, 1, 12, 2, 14, 3, 16, 4, 18, 5, 20, 6, 22, 7, 24,
		8, 26, 9, 28, 10, 30, 11, 32, 12, 34, 13, 36, 14, 38, 15, 40, 0, 42, 0,
		44, 0, 46, 0, 48, 0, 50, 0, 52, 0, 54, 0, 56, 0, 58, 0, 60, 0, 62, 16,
		64, 17, 66, 18, 68, 0, 70, 0, 72, 0, 74, 0, 76, 0, 78, 0, 80, 0, 82, 0,
		84, 0, 86, 0, 88, 0, 90, 19, 92, 20, 94, 21, 96, 22, 4, 0, 1, 2, 3, 4,
		2, 0, 9, 9, 32, 32, 2, 0, 10, 10, 13, 13, 3, 0, 65, 90, 95, 95, 97, 122,
		4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 724, 0, 10, 1, 0, 0, 0, 0, 12, 1,
		0, 0, 0, 0, 14, 1, 0, 0, 0, 0, 16, 1, 0, 0, 0, 0, 18, 1, 0, 0, 0, 0, 20,
		1, 0, 0, 0, 0, 22, 1, 0, 0, 0, 0, 24, 1, 0, 0, 0, 0, 26, 1, 0, 0, 0, 0,
		28, 1, 0, 0, 0, 0, 30, 1, 0, 0, 0, 0, 32, 1, 0, 0, 0, 0, 34, 1, 0, 0, 0,
		0, 36, 1, 0, 0, 0, 1, 38, 1, 0, 0, 0, 2, 40, 1, 0, 0, 0, 2, 42, 1, 0, 0,
		0, 2, 44, 1, 0, 0, 0, 2, 46, 1, 0, 0, 0, 2, 48, 1, 0, 0, 0, 2, 50, 1, 0,
		0, 0, 2, 52, 1, 0, 0, 0, 2, 54, 1, 0, 0, 0, 2, 56, 1, 0, 0, 0, 2, 58, 1,
		0, 0, 0, 2, 60, 1, 0, 0, 0, 2, 62, 1, 0, 0, 0, 2, 64, 1, 0, 0, 0, 2, 66,
		1, 0, 0, 0, 3, 68, 1, 0, 0, 0, 3, 70, 1, 0, 0, 0, 3, 72, 1, 0, 0, 0, 3,
		74, 1, 0, 0, 0, 3, 76, 1, 0, 0, 0, 3, 78, 1, 0, 0, 0, 3, 80, 1, 0, 0, 0,
		3, 82, 1, 0, 0, 0, 3, 84, 1, 0, 0, 0, 3, 86, 1, 0, 0, 0, 3, 88, 1, 0, 0,
		0, 3, 90, 1, 0, 0, 0, 3, 92, 1, 0, 0, 0, 3, 94, 1, 0, 0, 0, 3, 96, 1, 0,
		0, 0, 4, 98, 1, 0, 0, 0, 6, 108, 1, 0, 0, 0, 8, 113, 1, 0, 0, 0, 10, 120,
		1, 0, 0, 0, 12, 132, 1, 0, 0, 0, 14, 143, 1, 0, 0, 0, 16, 149, 1, 0, 0,
		0, 18, 158, 1, 0, 0, 0, 20, 167, 1, 0, 0, 0, 22, 176, 1, 0, 0, 0, 24, 186,
		1, 0, 0, 0, 26, 202, 1, 0, 0, 0, 28, 215, 1, 0, 0, 0, 30, 233, 1, 0, 0,
		0, 32, 249, 1, 0, 0, 0, 34, 262, 1, 0, 0, 0, 36, 274, 1, 0, 0, 0, 38, 291,
		1, 0, 0, 0, 40, 298, 1, 0, 0, 0, 42, 309, 1, 0, 0, 0, 44, 320, 1, 0, 0,
		0, 46, 340, 1, 0, 0, 0, 48, 358, 1, 0, 0, 0, 50, 369, 1, 0, 0, 0, 52, 381,
		1, 0, 0, 0, 54, 399, 1, 0, 0, 0, 56, 414, 1, 0, 0, 0, 58, 429, 1, 0, 0,
		0, 60, 443, 1, 0, 0, 0, 62, 460, 1, 0, 0, 0, 64, 475, 1, 0, 0, 0, 66, 482,
		1, 0, 0, 0, 68, 495, 1, 0, 0, 0, 70, 506, 1, 0, 0, 0, 72, 517, 1, 0, 0,
		0, 74, 537, 1, 0, 0, 0, 76, 555, 1, 0, 0, 0, 78, 566, 1, 0, 0, 0, 80, 578,
		1, 0, 0, 0, 82, 596, 1, 0, 0, 0, 84, 611, 1, 0, 0, 0, 86, 626, 1, 0, 0,
		0, 88, 640, 1, 0, 0, 0, 90, 657, 1, 0, 0, 0, 92, 674, 1, 0, 0, 0, 94, 689,
		1, 0, 0, 0, 96, 696, 1, 0, 0, 0, 98, 102, 5, 58, 0, 0, 99, 101, 7, 0, 0,
		0, 100, 99, 1, 0, 0, 0, 101, 104, 1, 0, 0, 0, 102, 100, 1, 0, 0, 0, 102,
		103, 1, 0, 0, 0, 103, 5, 1, 0, 0, 0, 104, 102, 1, 0, 0, 0, 105, 107, 8,
		1, 0, 0, 106, 105, 1, 0, 0, 0, 107, 110, 1, 0, 0, 0, 108, 106, 1, 0, 0,
		0, 108, 109, 1, 0, 0, 0, 109, 111, 1, 0, 0, 0, 110, 108, 1, 0, 0, 0, 111,
		112, 7, 1, 0, 0, 112, 7, 1, 0, 0, 0, 113, 117, 7, 2, 0, 0, 114, 116, 7,
		3, 0, 0, 115, 114, 1, 0, 0, 0, 116, 119, 1, 0, 0, 0, 117, 115, 1, 0, 0,
		0, 117, 118, 1, 0, 0, 0, 118, 9, 1, 0, 0, 0, 119, 117, 1, 0, 0, 0, 120,
		124, 5, 35, 0, 0, 121, 123, 8, 1, 0, 0, 122, 121, 1, 0, 0, 0, 123, 126,
		1, 0, 0, 0, 124, 122, 1, 0, 0, 0, 124, 125, 1, 0, 0, 0, 125, 127, 1, 0,
		0, 0, 126, 124, 1, 0, 0, 0, 127, 128, 6, 3, 0, 0, 128, 11, 1, 0, 0, 0,
		129, 131, 7, 0, 0, 0, 130, 129, 1, 0, 0, 0, 131, 134, 1, 0, 0, 0, 132,
		130, 1, 0, 0, 0, 132, 133, 1, 0, 0, 0, 133, 136, 1, 0, 0, 0, 134, 132,
		1, 0, 0, 0, 135, 137, 7, 1, 0, 0, 136, 135, 1, 0, 0, 0, 137, 138, 1, 0,
		0, 0, 138, 136, 1, 0, 0, 0, 138, 139, 1, 0, 0, 0, 139, 140, 1, 0, 0, 0,
		140, 141, 6, 4, 0, 0, 141, 13, 1, 0, 0, 0, 142, 144, 7, 0, 0, 0, 143, 142,
		1, 0, 0, 0, 144, 145, 1, 0, 0, 0, 145, 143, 1, 0, 0, 0, 145, 146, 1, 0,
		0, 0, 146, 147, 1, 0, 0, 0, 147, 148, 6, 5, 0, 0, 148, 15, 1, 0, 0, 0,
		149, 150, 5, 116, 0, 0, 150, 151, 5, 121, 0, 0, 151, 152, 5, 112, 0, 0,
		152, 153, 5, 101, 0, 0, 153, 154, 1, 0, 0, 0, 154, 155, 3, 4, 0, 0, 155,
		156, 1, 0, 0, 0, 156, 157, 6, 6, 1, 0, 157, 17, 1, 0, 0, 0, 158, 159, 5,
		110, 0, 0, 159, 160, 5, 97, 0, 0, 160, 161, 5, 109, 0, 0, 161, 162, 5,
		101, 0, 0, 162, 163, 1, 0, 0, 0, 163, 164, 3, 4, 0, 0, 164, 165, 1, 0,
		0, 0, 165, 166, 6, 7, 1, 0, 166, 19, 1, 0, 0, 0, 167, 168, 5, 116, 0, 0,
		168, 169, 5, 97, 0, 0, 169, 170, 5, 103, 0, 0, 170, 171, 5, 115, 0, 0,
		171, 172, 1, 0, 0, 0, 172, 173, 3, 4, 0, 0, 173, 174, 1, 0, 0, 0, 174,
		175, 6, 8, 1, 0, 175, 21, 1, 0, 0, 0, 176, 177, 5, 111, 0, 0, 177, 178,
		5, 119, 0, 0, 178, 179, 5, 110, 0, 0, 179, 180, 5, 101, 0, 0, 180, 181,
		5, 114, 0, 0, 181, 182, 1, 0, 0, 0, 182, 183, 3, 4, 0, 0, 183, 184, 1,
		0, 0, 0, 184, 185, 6, 9, 1, 0, 185, 23, 1, 0, 0, 0, 186, 187, 5, 100, 0,
		0, 187, 188, 5, 101, 0, 0, 188, 189, 5, 115, 0, 0, 189, 190, 5, 116, 0,
		0, 190, 191, 5, 105, 0, 0, 191, 192, 5, 110, 0, 0, 192, 193, 5, 97, 0,
		0, 193, 194, 5, 116, 0, 0, 194, 195, 5, 105, 0, 0, 195, 196, 5, 111, 0,
		0, 196, 197, 5, 110, 0, 0, 197, 198, 1, 0, 0, 0, 198, 199, 3, 4, 0, 0,
		199, 200, 1, 0, 0, 0, 200, 201, 6, 10, 1, 0, 201, 25, 1, 0, 0, 0, 202,
		203, 5, 115, 0, 0, 203, 204, 5, 99, 0, 0, 204, 205, 5, 104, 0, 0, 205,
		206, 5, 101, 0, 0, 206, 207, 5, 100, 0, 0, 207, 208, 5, 117, 0, 0, 208,
		209, 5, 108, 0, 0, 209, 210, 5, 101, 0, 0, 210, 211, 1, 0, 0, 0, 211, 212,
		3, 4, 0, 0, 212, 213, 1, 0, 0, 0, 213, 214, 6, 11, 1, 0, 214, 27, 1, 0,
		0, 0, 215, 216, 5, 100, 0, 0, 216, 217, 5, 101, 0, 0, 217, 218, 5, 115,
		0, 0, 218, 219, 5, 99, 0, 0, 219, 220, 5, 114, 0, 0, 220, 221, 5, 105,
		0, 0, 221, 222, 5, 112, 0, 0, 222, 223, 5, 116, 0, 0, 223, 224, 5, 105,
		0, 0, 224, 225, 5, 111, 0, 0, 225, 226, 5, 110, 0, 0, 226, 227, 1, 0, 0,
		0, 227, 228, 3, 4, 0, 0, 228, 229, 5, 124, 0, 0, 229, 230, 3, 6, 1, 0,
		230, 231, 1, 0, 0, 0, 231, 232, 6, 12, 2, 0, 232, 29, 1, 0, 0, 0, 233,
		234, 5, 100, 0, 0, 234, 235, 5, 101, 0, 0, 235, 236, 5, 115, 0, 0, 236,
		237, 5, 99, 0, 0, 237, 238, 5, 114, 0, 0, 238, 239, 5, 105, 0, 0, 239,
		240, 5, 112, 0, 0, 240, 241, 5, 116, 0, 0, 241, 242, 5, 105, 0, 0, 242,
		243, 5, 111, 0, 0, 243, 244, 5, 110, 0, 0, 244, 245, 1, 0, 0, 0, 245, 246,
		3, 4, 0, 0, 246, 247, 1, 0, 0, 0, 247, 248, 6, 13, 1, 0, 248, 31, 1, 0,
		0, 0, 249, 250, 5, 115, 0, 0, 250, 251, 5, 111, 0, 0, 251, 252, 5, 117,
		0, 0, 252, 253, 5, 114, 0, 0, 253, 254, 5, 99, 0, 0, 254, 255, 5, 101,
		0, 0, 255, 256, 5, 115, 0, 0, 256, 257, 1, 0, 0, 0, 257, 258, 3, 4, 0,
		0, 258, 259, 3, 6, 1, 0, 259, 260, 1, 0, 0, 0, 260, 261, 6, 14, 2, 0, 261,
		33, 1, 0, 0, 0, 262, 263, 5, 112, 0, 0, 263, 264, 5, 97, 0, 0, 264, 265,
		5, 114, 0, 0, 265, 266, 5, 97, 0, 0, 266, 267, 5, 109, 0, 0, 267, 268,
		5, 115, 0, 0, 268, 269, 1, 0, 0, 0, 269, 270, 3, 4, 0, 0, 270, 271, 3,
		6, 1, 0, 271, 272, 1, 0, 0, 0, 272, 273, 6, 15, 2, 0, 273, 35, 1, 0, 0,
		0, 274, 275, 5, 112, 0, 0, 275, 276, 5, 105, 0, 0, 276, 277, 5, 112, 0,
		0, 277, 278, 5, 101, 0, 0, 278, 279, 5, 108, 0, 0, 279, 280, 5, 105, 0,
		0, 280, 281, 5, 110, 0, 0, 281, 282, 5, 101, 0, 0, 282, 283, 1, 0, 0, 0,
		283, 284, 3, 4, 0, 0, 284, 285, 3, 6, 1, 0, 285, 286, 1, 0, 0, 0, 286,
		287, 6, 16, 3, 0, 287, 37, 1, 0, 0, 0, 288, 290, 8, 1, 0, 0, 289, 288,
		1, 0, 0, 0, 290, 293, 1, 0, 0, 0, 291, 289, 1, 0, 0, 0, 291, 292, 1, 0,
		0, 0, 292, 294, 1, 0, 0, 0, 293, 291, 1, 0, 0, 0, 294, 295, 7, 1, 0, 0,
		295, 296, 1, 0, 0, 0, 296, 297, 6, 17, 4, 0, 297, 39, 1, 0, 0, 0, 298,
		299, 4, 18, 0, 0, 299, 300, 5, 116, 0, 0, 300, 301, 5, 121, 0, 0, 301,
		302, 5, 112, 0, 0, 302, 303, 5, 101, 0, 0, 303, 304, 1, 0, 0, 0, 304, 305,
		3, 4, 0, 0, 305, 306, 1, 0, 0, 0, 306, 307, 6, 18, 5, 0, 307, 308, 6, 18,
		6, 0, 308, 41, 1, 0, 0, 0, 309, 310, 4, 19, 1, 0, 310, 311, 5, 110, 0,
		0, 311, 312, 5, 97, 0, 0, 312, 313, 5, 109, 0, 0, 313, 314, 5, 101, 0,
		0, 314, 315, 1, 0, 0, 0, 315, 316, 3, 4, 0, 0, 316, 317, 1, 0, 0, 0, 317,
		318, 6, 19, 7, 0, 318, 319, 6, 19, 6, 0, 319, 43, 1, 0, 0, 0, 320, 321,
		4, 20, 2, 0, 321, 322, 5, 100, 0, 0, 322, 323, 5, 101, 0, 0, 323, 324,
		5, 115, 0, 0, 324, 325, 5, 99, 0, 0, 325, 326, 5, 114, 0, 0, 326, 327,
		5, 105, 0, 0, 327, 328, 5, 112, 0, 0, 328, 329, 5, 116, 0, 0, 329, 330,
		5, 105, 0, 0, 330, 331, 5, 111, 0, 0, 331, 332, 5, 110, 0, 0, 332, 333,
		1, 0, 0, 0, 333, 334, 3, 4, 0, 0, 334, 335, 5, 124, 0, 0, 335, 336, 3,
		6, 1, 0, 336, 337, 1, 0, 0, 0, 337, 338, 6, 20, 8, 0, 338, 339, 6, 20,
		9, 0, 339, 45, 1, 0, 0, 0, 340, 341, 4, 21, 3, 0, 341, 342, 5, 100, 0,
		0, 342, 343, 5, 101, 0, 0, 343, 344, 5, 115, 0, 0, 344, 345, 5, 99, 0,
		0, 345, 346, 5, 114, 0, 0, 346, 347, 5, 105, 0, 0, 347, 348, 5, 112, 0,
		0, 348, 349, 5, 116, 0, 0, 349, 350, 5, 105, 0, 0, 350, 351, 5, 111, 0,
		0, 351, 352, 5, 110, 0, 0, 352, 353, 1, 0, 0, 0, 353, 354, 3, 4, 0, 0,
		354, 355, 1, 0, 0, 0, 355, 356, 6, 21, 10, 0, 356, 357, 6, 21, 6, 0, 357,
		47, 1, 0, 0, 0, 358, 359, 4, 22, 4, 0, 359, 360, 5, 116, 0, 0, 360, 361,
		5, 97, 0, 0, 361, 362, 5, 103, 0, 0, 362, 363, 5, 115, 0, 0, 363, 364,
		1, 0, 0, 0, 364, 365, 3, 4, 0, 0, 365, 366, 1, 0, 0, 0, 366, 367, 6, 22,
		11, 0, 367, 368, 6, 22, 6, 0, 368, 49, 1, 0, 0, 0, 369, 370, 4, 23, 5,
		0, 370, 371, 5, 111, 0, 0, 371, 372, 5, 119, 0, 0, 372, 373, 5, 110, 0,
		0, 373, 374, 5, 101, 0, 0, 374, 375, 5, 114, 0, 0, 375, 376, 1, 0, 0, 0,
		376, 377, 3, 4, 0, 0, 377, 378, 1, 0, 0, 0, 378, 379, 6, 23, 12, 0, 379,
		380, 6, 23, 6, 0, 380, 51, 1, 0, 0, 0, 381, 382, 4, 24, 6, 0, 382, 383,
		5, 100, 0, 0, 383, 384, 5, 101, 0, 0, 384, 385, 5, 115, 0, 0, 385, 386,
		5, 116, 0, 0, 386, 387, 5, 105, 0, 0, 387, 388, 5, 110, 0, 0, 388, 389,
		5, 97, 0, 0, 389, 390, 5, 116, 0, 0, 390, 391, 5, 105, 0, 0, 391, 392,
		5, 111, 0, 0, 392, 393, 5, 110, 0, 0, 393, 394, 1, 0, 0, 0, 394, 395, 3,
		4, 0, 0, 395, 396, 1, 0, 0, 0, 396, 397, 6, 24, 13, 0, 397, 398, 6, 24,
		6, 0, 398, 53, 1, 0, 0, 0, 399, 400, 4, 25, 7, 0, 400, 401, 5, 115, 0,
		0, 401, 402, 5, 99, 0, 0, 402, 403, 5, 104, 0, 0, 403, 404, 5, 101, 0,
		0, 404, 405, 5, 100, 0, 0, 405, 406, 5, 117, 0, 0, 406, 407, 5, 108, 0,
		0, 407, 408, 5, 101, 0, 0, 408, 409, 1, 0, 0, 0, 409, 410, 3, 4, 0, 0,
		410, 411, 1, 0, 0, 0, 411, 412, 6, 25, 14, 0, 412, 413, 6, 25, 6, 0, 413,
		55, 1, 0, 0, 0, 414, 415, 4, 26, 8, 0, 415, 416, 5, 115, 0, 0, 416, 417,
		5, 111, 0, 0, 417, 418, 5, 117, 0, 0, 418, 419, 5, 114, 0, 0, 419, 420,
		5, 99, 0, 0, 420, 421, 5, 101, 0, 0, 421, 422, 5, 115, 0, 0, 422, 423,
		1, 0, 0, 0, 423, 424, 3, 4, 0, 0, 424, 425, 3, 6, 1, 0, 425, 426, 1, 0,
		0, 0, 426, 427, 6, 26, 15, 0, 427, 428, 6, 26, 9, 0, 428, 57, 1, 0, 0,
		0, 429, 430, 4, 27, 9, 0, 430, 431, 5, 112, 0, 0, 431, 432, 5, 97, 0, 0,
		432, 433, 5, 114, 0, 0, 433, 434, 5, 97, 0, 0, 434, 435, 5, 109, 0, 0,
		435, 436, 5, 115, 0, 0, 436, 437, 1, 0, 0, 0, 437, 438, 3, 4, 0, 0, 438,
		439, 3, 6, 1, 0, 439, 440, 1, 0, 0, 0, 440, 441, 6, 27, 16, 0, 441, 442,
		6, 27, 9, 0, 442, 59, 1, 0, 0, 0, 443, 444, 4, 28, 10, 0, 444, 445, 5,
		112, 0, 0, 445, 446, 5, 105, 0, 0, 446, 447, 5, 112, 0, 0, 447, 448, 5,
		101, 0, 0, 448, 449, 5, 108, 0, 0, 449, 450, 5, 105, 0, 0, 450, 451, 5,
		110, 0, 0, 451, 452, 5, 101, 0, 0, 452, 453, 1, 0, 0, 0, 453, 454, 3, 4,
		0, 0, 454, 455, 3, 6, 1, 0, 455, 456, 1, 0, 0, 0, 456, 457, 6, 28, 17,
		0, 457, 458, 6, 28, 18, 0, 458, 61, 1, 0, 0, 0, 459, 461, 7, 0, 0, 0, 460,
		459, 1, 0, 0, 0, 461, 462, 1, 0, 0, 0, 462, 460, 1, 0, 0, 0, 462, 463,
		1, 0, 0, 0, 463, 467, 1, 0, 0, 0, 464, 466, 8, 1, 0, 0, 465, 464, 1, 0,
		0, 0, 466, 469, 1, 0, 0, 0, 467, 465, 1, 0, 0, 0, 467, 468, 1, 0, 0, 0,
		468, 470, 1, 0, 0, 0, 469, 467, 1, 0, 0, 0, 470, 471, 7, 1, 0, 0, 471,
		63, 1, 0, 0, 0, 472, 474, 7, 0, 0, 0, 473, 472, 1, 0, 0, 0, 474, 477, 1,
		0, 0, 0, 475, 473, 1, 0, 0, 0, 475, 476, 1, 0, 0, 0, 476, 478, 1, 0, 0,
		0, 477, 475, 1, 0, 0, 0, 478, 479, 7, 1, 0, 0, 479, 480, 1, 0, 0, 0, 480,
		481, 6, 30, 0, 0, 481, 65, 1, 0, 0, 0, 482, 483, 4, 31, 11, 0, 483, 487,
		5, 35, 0, 0, 484, 486, 8, 1, 0, 0, 485, 484, 1, 0, 0, 0, 486, 489, 1, 0,
		0, 0, 487, 485, 1, 0, 0, 0, 487, 488, 1, 0, 0, 0, 488, 491, 1, 0, 0, 0,
		489, 487, 1, 0, 0, 0, 490, 492, 7, 1, 0, 0, 491, 490, 1, 0, 0, 0, 491,
		492, 1, 0, 0, 0, 492, 493, 1, 0, 0, 0, 493, 494, 6, 31, 0, 0, 494, 67,
		1, 0, 0, 0, 495, 496, 4, 32, 12, 0, 496, 497, 5, 116, 0, 0, 497, 498, 5,
		121, 0, 0, 498, 499, 5, 112, 0, 0, 499, 500, 5, 101, 0, 0, 500, 501, 1,
		0, 0, 0, 501, 502, 3, 4, 0, 0, 502, 503, 1, 0, 0, 0, 503, 504, 6, 32, 5,
		0, 504, 505, 6, 32, 6, 0, 505, 69, 1, 0, 0, 0, 506, 507, 4, 33, 13, 0,
		507, 508, 5, 110, 0, 0, 508, 509, 5, 97, 0, 0, 509, 510, 5, 109, 0, 0,
		510, 511, 5, 101, 0, 0, 511, 512, 1, 0, 0, 0, 512, 513, 3, 4, 0, 0, 513,
		514, 1, 0, 0, 0, 514, 515, 6, 33, 7, 0, 515, 516, 6, 33, 6, 0, 516, 71,
		1, 0, 0, 0, 517, 518, 4, 34, 14, 0, 518, 519, 5, 100, 0, 0, 519, 520, 5,
		101, 0, 0, 520, 521, 5, 115, 0, 0, 521, 522, 5, 99, 0, 0, 522, 523, 5,
		114, 0, 0, 523, 524, 5, 105, 0, 0, 524, 525, 5, 112, 0, 0, 525, 526, 5,
		116, 0, 0, 526, 527, 5, 105, 0, 0, 527, 528, 5, 111, 0, 0, 528, 529, 5,
		110, 0, 0, 529, 530, 1, 0, 0, 0, 530, 531, 3, 4, 0, 0, 531, 532, 5, 124,
		0, 0, 532, 533, 3, 6, 1, 0, 533, 534, 1, 0, 0, 0, 534, 535, 6, 34, 8, 0,
		535, 536, 6, 34, 9, 0, 536, 73, 1, 0, 0, 0, 537, 538, 4, 35, 15, 0, 538,
		539, 5, 100, 0, 0, 539, 540, 5, 101, 0, 0, 540, 541, 5, 115, 0, 0, 541,
		542, 5, 99, 0, 0, 542, 543, 5, 114, 0, 0, 543, 544, 5, 105, 0, 0, 544,
		545, 5, 112, 0, 0, 545, 546, 5, 116, 0, 0, 546, 547, 5, 105, 0, 0, 547,
		548, 5, 111, 0, 0, 548, 549, 5, 110, 0, 0, 549, 550, 1, 0, 0, 0, 550, 551,
		3, 4, 0, 0, 551, 552, 1, 0, 0, 0, 552, 553, 6, 35, 10, 0, 553, 554, 6,
		35, 6, 0, 554, 75, 1, 0, 0, 0, 555, 556, 4, 36, 16, 0, 556, 557, 5, 116,
		0, 0, 557, 558, 5, 97, 0, 0, 558, 559, 5, 103, 0, 0, 559, 560, 5, 115,
		0, 0, 560, 561, 1, 0, 0, 0, 561, 562, 3, 4, 0, 0, 562, 563, 1, 0, 0, 0,
		563, 564, 6, 36, 11, 0, 564, 565, 6, 36, 6, 0, 565, 77, 1, 0, 0, 0, 566,
		567, 4, 37, 17, 0, 567, 568, 5, 111, 0, 0, 568, 569, 5, 119, 0, 0, 569,
		570, 5, 110, 0, 0, 570, 571, 5, 101, 0, 0, 571, 572, 5, 114, 0, 0, 572,
		573, 1, 0, 0, 0, 573, 574, 3, 4, 0, 0, 574, 575, 1, 0, 0, 0, 575, 576,
		6, 37, 12, 0, 576, 577, 6, 37, 6, 0, 577, 79, 1, 0, 0, 0, 578, 579, 4,
		38, 18, 0, 579, 580, 5, 100, 0, 0, 580, 581, 5, 101, 0, 0, 581, 582, 5,
		115, 0, 0, 582, 583, 5, 116, 0, 0, 583, 584, 5, 105, 0, 0, 584, 585, 5,
		110, 0, 0, 585, 586, 5, 97, 0, 0, 586, 587, 5, 116, 0, 0, 587, 588, 5,
		105, 0, 0, 588, 589, 5, 111, 0, 0, 589, 590, 5, 110, 0, 0, 590, 591, 1,
		0, 0, 0, 591, 592, 3, 4, 0, 0, 592, 593, 1, 0, 0, 0, 593, 594, 6, 38, 13,
		0, 594, 595, 6, 38, 6, 0, 595, 81, 1, 0, 0, 0, 596, 597, 4, 39, 19, 0,
		597, 598, 5, 115, 0, 0, 598, 599, 5, 99, 0, 0, 599, 600, 5, 104, 0, 0,
		600, 601, 5, 101, 0, 0, 601, 602, 5, 100, 0, 0, 602, 603, 5, 117, 0, 0,
		603, 604, 5, 108, 0, 0, 604, 605, 5, 101, 0, 0, 605, 606, 1, 0, 0, 0, 606,
		607, 3, 4, 0, 0, 607, 608, 1, 0, 0, 0, 608, 609, 6, 39, 14, 0, 609, 610,
		6, 39, 6, 0, 610, 83, 1, 0, 0, 0, 611, 612, 4, 40, 20, 0, 612, 613, 5,
		115, 0, 0, 613, 614, 5, 111, 0, 0, 614, 615, 5, 117, 0, 0, 615, 616, 5,
		114, 0, 0, 616, 617, 5, 99, 0, 0, 617, 618, 5, 101, 0, 0, 618, 619, 5,
		115, 0, 0, 619, 620, 1, 0, 0, 0, 620, 621, 3, 4, 0, 0, 621, 622, 3, 6,
		1, 0, 622, 623, 1, 0, 0, 0, 623, 624, 6, 40, 15, 0, 624, 625, 6, 40, 9,
		0, 625, 85, 1, 0, 0, 0, 626, 627, 4, 41, 21, 0, 627, 628, 5, 112, 0, 0,
		628, 629, 5, 97, 0, 0, 629, 630, 5, 114, 0, 0, 630, 631, 5, 97, 0, 0, 631,
		632, 5, 109, 0, 0, 632, 633, 5, 115, 0, 0, 633, 634, 1, 0, 0, 0, 634, 635,
		3, 4, 0, 0, 635, 636, 3, 6, 1, 0, 636, 637, 1, 0, 0, 0, 637, 638, 6, 41,
		16, 0, 638, 639, 6, 41, 9, 0, 639, 87, 1, 0, 0, 0, 640, 641, 4, 42, 22,
		0, 641, 642, 5, 112, 0, 0, 642, 643, 5, 105, 0, 0, 643, 644, 5, 112, 0,
		0, 644, 645, 5, 101, 0, 0, 645, 646, 5, 108, 0, 0, 646, 647, 5, 105, 0,
		0, 647, 648, 5, 110, 0, 0, 648, 649, 5, 101, 0, 0, 649, 650, 1, 0, 0, 0,
		650, 651, 3, 4, 0, 0, 651, 652, 3, 6, 1, 0, 652, 653, 1, 0, 0, 0, 653,
		654, 6, 42, 17, 0, 654, 655, 6, 42, 18, 0, 655, 89, 1, 0, 0, 0, 656, 658,
		7, 0, 0, 0, 657, 656, 1, 0, 0, 0, 658, 659, 1, 0, 0, 0, 659, 657, 1, 0,
		0, 0, 659, 660, 1, 0, 0, 0, 660, 661, 1, 0, 0, 0, 661, 662, 5, 64, 0, 0,
		662, 663, 3, 8, 2, 0, 663, 667, 3, 4, 0, 0, 664, 666, 8, 1, 0, 0, 665,
		664, 1, 0, 0, 0, 666, 669, 1, 0, 0, 0, 667, 665, 1, 0, 0, 0, 667, 668,
		1, 0, 0, 0, 668, 671, 1, 0, 0, 0, 669, 667, 1, 0, 0, 0, 670, 672, 7, 1,
		0, 0, 671, 670, 1, 0, 0, 0, 671, 672, 1, 0, 0, 0, 672, 91, 1, 0, 0, 0,
		673, 675, 7, 0, 0, 0, 674, 673, 1, 0, 0, 0, 675, 676, 1, 0, 0, 0, 676,
		674, 1, 0, 0, 0, 676, 677, 1, 0, 0, 0, 677, 679, 1, 0, 0, 0, 678, 680,
		8, 1, 0, 0, 679, 678, 1, 0, 0, 0, 680, 681, 1, 0, 0, 0, 681, 679, 1, 0,
		0, 0, 681, 682, 1, 0, 0, 0, 682, 684, 1, 0, 0, 0, 683, 685, 7, 1, 0, 0,
		684, 683, 1, 0, 0, 0, 684, 685, 1, 0, 0, 0, 685, 93, 1, 0, 0, 0, 686, 688,
		7, 0, 0, 0, 687, 686, 1, 0, 0, 0, 688, 691, 1, 0, 0, 0, 689, 687, 1, 0,
		0, 0, 689, 690, 1, 0, 0, 0, 690, 692, 1, 0, 0, 0, 691, 689, 1, 0, 0, 0,
		692, 693, 7, 1, 0, 0, 693, 694, 1, 0, 0, 0, 694, 695, 6, 45, 0, 0, 695,
		95, 1, 0, 0, 0, 696, 697, 4, 46, 23, 0, 697, 701, 5, 35, 0, 0, 698, 700,
		8, 1, 0, 0, 699, 698, 1, 0, 0, 0, 700, 703, 1, 0, 0, 0, 701, 699, 1, 0,
		0, 0, 701, 702, 1, 0, 0, 0, 702, 705, 1, 0, 0, 0, 703, 701, 1, 0, 0, 0,
		704, 706, 7, 1, 0, 0, 705, 704, 1, 0, 0, 0, 705, 706, 1, 0, 0, 0, 706,
		707, 1, 0, 0, 0, 707, 708, 6, 46, 0, 0, 708, 97, 1, 0, 0, 0, 26, 0, 1,
		2, 3, 102, 108, 117, 124, 132, 138, 145, 291, 462, 467, 475, 487, 491,
		659, 667, 671, 676, 681, 684, 689, 701, 705, 19, 6, 0, 0, 5, 1, 0, 5, 2,
		0, 5, 3, 0, 4, 0, 0, 7, 4, 0, 2, 1, 0, 7, 5, 0, 7, 10, 0, 2, 2, 0, 7, 11,
		0, 7, 6, 0, 7, 7, 0, 7, 8, 0, 7, 9, 0, 7, 12, 0, 7, 13, 0, 7, 14, 0, 2,
		3, 0,
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
	PipeLangLexerCOMMENT        = 1
	PipeLangLexerWS_BLANK       = 2
	PipeLangLexerWS             = 3
	PipeLangLexerTYPE           = 4
	PipeLangLexerNAME           = 5
	PipeLangLexerTAGS           = 6
	PipeLangLexerOWNER          = 7
	PipeLangLexerDESTINATION    = 8
	PipeLangLexerSCHEDULE       = 9
	PipeLangLexerDESCRIPTION_ML = 10
	PipeLangLexerDESCRIPTION    = 11
	PipeLangLexerSOURCES        = 12
	PipeLangLexerPARAMS         = 13
	PipeLangLexerPIPELINE       = 14
	PipeLangLexerREST_OF_LINE   = 15
	PipeLangLexerBLOCK_LINE     = 16
	PipeLangLexerBLOCK_BLANK    = 17
	PipeLangLexerBLOCK_COMMENT  = 18
	PipeLangLexerNODE_HEADER    = 19
	PipeLangLexerPRQL_LINE      = 20
	PipeLangLexerPIPE_BLANK     = 21
	PipeLangLexerPIPE_COMMENT   = 22
)

// PipeLangLexer modes.
const (
	PipeLangLexerVALUE_MODE = iota + 1
	PipeLangLexerBLOCK_MODE
	PipeLangLexerPIPELINE_MODE
)

func (l *PipeLangLexer) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 18:
		return l.BLOCK_TYPE_Sempred(localctx, predIndex)

	case 19:
		return l.BLOCK_NAME_Sempred(localctx, predIndex)

	case 20:
		return l.BLOCK_DESC_ML_Sempred(localctx, predIndex)

	case 21:
		return l.BLOCK_DESC_Sempred(localctx, predIndex)

	case 22:
		return l.BLOCK_TAGS_Sempred(localctx, predIndex)

	case 23:
		return l.BLOCK_OWNER_Sempred(localctx, predIndex)

	case 24:
		return l.BLOCK_DESTINATION_Sempred(localctx, predIndex)

	case 25:
		return l.BLOCK_SCHEDULE_Sempred(localctx, predIndex)

	case 26:
		return l.BLOCK_SOURCES_Sempred(localctx, predIndex)

	case 27:
		return l.BLOCK_PARAMS_Sempred(localctx, predIndex)

	case 28:
		return l.BLOCK_PIPELINE_Sempred(localctx, predIndex)

	case 31:
		return l.BLOCK_COMMENT_Sempred(localctx, predIndex)

	case 32:
		return l.PIPE_TYPE_Sempred(localctx, predIndex)

	case 33:
		return l.PIPE_NAME_Sempred(localctx, predIndex)

	case 34:
		return l.PIPE_DESC_ML_Sempred(localctx, predIndex)

	case 35:
		return l.PIPE_DESC_Sempred(localctx, predIndex)

	case 36:
		return l.PIPE_TAGS_Sempred(localctx, predIndex)

	case 37:
		return l.PIPE_OWNER_Sempred(localctx, predIndex)

	case 38:
		return l.PIPE_DESTINATION_Sempred(localctx, predIndex)

	case 39:
		return l.PIPE_SCHEDULE_Sempred(localctx, predIndex)

	case 40:
		return l.PIPE_SOURCES_Sempred(localctx, predIndex)

	case 41:
		return l.PIPE_PARAMS_Sempred(localctx, predIndex)

	case 42:
		return l.PIPE_PIPELINE_Sempred(localctx, predIndex)

	case 46:
		return l.PIPE_COMMENT_Sempred(localctx, predIndex)

	default:
		panic("No registered predicate for: " + fmt.Sprint(ruleIndex))
	}
}

func (p *PipeLangLexer) BLOCK_TYPE_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.GetCharPositionInLine() == 0

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *PipeLangLexer) BLOCK_NAME_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.GetCharPositionInLine() == 0

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *PipeLangLexer) BLOCK_DESC_ML_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 2:
		return p.GetCharPositionInLine() == 0

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *PipeLangLexer) BLOCK_DESC_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 3:
		return p.GetCharPositionInLine() == 0

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *PipeLangLexer) BLOCK_TAGS_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 4:
		return p.GetCharPositionInLine() == 0

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *PipeLangLexer) BLOCK_OWNER_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 5:
		return p.GetCharPositionInLine() == 0

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *PipeLangLexer) BLOCK_DESTINATION_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 6:
		return p.GetCharPositionInLine() == 0

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *PipeLangLexer) BLOCK_SCHEDULE_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 7:
		return p.GetCharPositionInLine() == 0

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *PipeLangLexer) BLOCK_SOURCES_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 8:
		return p.GetCharPositionInLine() == 0

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *PipeLangLexer) BLOCK_PARAMS_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 9:
		return p.GetCharPositionInLine() == 0

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *PipeLangLexer) BLOCK_PIPELINE_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 10:
		return p.GetCharPositionInLine() == 0

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *PipeLangLexer) BLOCK_COMMENT_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 11:
		return p.GetCharPositionInLine() == 0

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *PipeLangLexer) PIPE_TYPE_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 12:
		return p.GetCharPositionInLine() == 0

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *PipeLangLexer) PIPE_NAME_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 13:
		return p.GetCharPositionInLine() == 0

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *PipeLangLexer) PIPE_DESC_ML_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 14:
		return p.GetCharPositionInLine() == 0

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *PipeLangLexer) PIPE_DESC_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 15:
		return p.GetCharPositionInLine() == 0

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *PipeLangLexer) PIPE_TAGS_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 16:
		return p.GetCharPositionInLine() == 0

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *PipeLangLexer) PIPE_OWNER_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 17:
		return p.GetCharPositionInLine() == 0

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *PipeLangLexer) PIPE_DESTINATION_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 18:
		return p.GetCharPositionInLine() == 0

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *PipeLangLexer) PIPE_SCHEDULE_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 19:
		return p.GetCharPositionInLine() == 0

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *PipeLangLexer) PIPE_SOURCES_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 20:
		return p.GetCharPositionInLine() == 0

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *PipeLangLexer) PIPE_PARAMS_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 21:
		return p.GetCharPositionInLine() == 0

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *PipeLangLexer) PIPE_PIPELINE_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 22:
		return p.GetCharPositionInLine() == 0

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *PipeLangLexer) PIPE_COMMENT_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 23:
		return p.GetCharPositionInLine() == 0

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
