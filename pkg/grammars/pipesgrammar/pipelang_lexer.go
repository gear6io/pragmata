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
		"DEFAULT_MODE", "SQL_BODY_MODE",
	}
	staticData.LiteralNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "'\\u0000'",
	}
	staticData.SymbolicNames = []string{
		"", "COMMENT", "WS_BLANK", "WS", "KW_DESCRIPTION", "KW_TAGS", "KW_TYPE",
		"KW_NODE", "KW_COPY_SCHEDULE", "KW_DATASOURCE", "KW_TARGET_DATASOURCE",
		"KW_SQL_ARROW", "REST_OF_LINE", "SQL_LINE",
	}
	staticData.RuleNames = []string{
		"COMMENT", "WS_BLANK", "WS", "KW_DESCRIPTION", "KW_TAGS", "KW_TYPE",
		"KW_NODE", "KW_COPY_SCHEDULE", "KW_DATASOURCE", "KW_TARGET_DATASOURCE",
		"KW_SQL_ARROW", "REST_OF_LINE", "SQL_LINE", "SQL_KW_DESCRIPTION", "SQL_KW_TAGS",
		"SQL_KW_TYPE", "SQL_KW_NODE", "SQL_KW_DATASOURCE", "SQL_KW_TARGET_DATASOURCE",
		"SQL_KW_COPY_SCHEDULE", "SQL_CHAR", "SQL_NL",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 13, 360, 6, -1, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3,
		7, 3, 2, 4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9,
		7, 9, 2, 10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7,
		14, 2, 15, 7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19,
		2, 20, 7, 20, 2, 21, 7, 21, 1, 0, 1, 0, 5, 0, 49, 8, 0, 10, 0, 12, 0, 52,
		9, 0, 1, 0, 1, 0, 1, 1, 5, 1, 57, 8, 1, 10, 1, 12, 1, 60, 9, 1, 1, 1, 4,
		1, 63, 8, 1, 11, 1, 12, 1, 64, 1, 1, 1, 1, 1, 2, 4, 2, 70, 8, 2, 11, 2,
		12, 2, 71, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 4, 3, 89, 8, 3, 11, 3, 12, 3, 90, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 4, 4, 99, 8, 4, 11, 4, 12, 4, 100, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 5, 1, 5, 4, 5, 109, 8, 5, 11, 5, 12, 5, 110, 1, 6, 1,
		6, 1, 6, 1, 6, 1, 6, 1, 6, 4, 6, 119, 8, 6, 11, 6, 12, 6, 120, 1, 7, 1,
		7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1,
		7, 1, 7, 4, 7, 138, 8, 7, 11, 7, 12, 7, 139, 1, 8, 1, 8, 1, 8, 1, 8, 1,
		8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 4, 8, 154, 8, 8, 11, 8, 12,
		8, 155, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1,
		9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 4, 9, 177, 8, 9, 11,
		9, 12, 9, 178, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 5, 10, 186, 8, 10, 10,
		10, 12, 10, 189, 9, 10, 1, 10, 1, 10, 5, 10, 193, 8, 10, 10, 10, 12, 10,
		196, 9, 10, 1, 10, 4, 10, 199, 8, 10, 11, 10, 12, 10, 200, 1, 10, 1, 10,
		1, 11, 1, 11, 5, 11, 207, 8, 11, 10, 11, 12, 11, 210, 9, 11, 1, 11, 5,
		11, 213, 8, 11, 10, 11, 12, 11, 216, 9, 11, 1, 12, 1, 12, 1, 13, 1, 13,
		1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1,
		13, 1, 13, 4, 13, 234, 8, 13, 11, 13, 12, 13, 235, 1, 13, 1, 13, 1, 13,
		1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 4, 14, 248, 8, 14, 11,
		14, 12, 14, 249, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15,
		1, 15, 1, 15, 4, 15, 262, 8, 15, 11, 15, 12, 15, 263, 1, 15, 1, 15, 1,
		15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 4, 16, 276, 8, 16,
		11, 16, 12, 16, 277, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1,
		17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 4, 17, 296,
		8, 17, 11, 17, 12, 17, 297, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1,
		18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18,
		1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 4, 18, 323, 8, 18, 11, 18, 12,
		18, 324, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19,
		1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 4,
		19, 346, 8, 19, 11, 19, 12, 19, 347, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20,
		1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 0, 0, 22, 2, 1, 4, 2, 6, 3, 8,
		4, 10, 5, 12, 6, 14, 7, 16, 8, 18, 9, 20, 10, 22, 11, 24, 12, 26, 13, 28,
		0, 30, 0, 32, 0, 34, 0, 36, 0, 38, 0, 40, 0, 42, 0, 44, 0, 2, 0, 1, 3,
		2, 0, 10, 10, 13, 13, 2, 0, 9, 9, 32, 32, 3, 0, 9, 10, 13, 13, 32, 32,
		381, 0, 2, 1, 0, 0, 0, 0, 4, 1, 0, 0, 0, 0, 6, 1, 0, 0, 0, 0, 8, 1, 0,
		0, 0, 0, 10, 1, 0, 0, 0, 0, 12, 1, 0, 0, 0, 0, 14, 1, 0, 0, 0, 0, 16, 1,
		0, 0, 0, 0, 18, 1, 0, 0, 0, 0, 20, 1, 0, 0, 0, 0, 22, 1, 0, 0, 0, 0, 24,
		1, 0, 0, 0, 0, 26, 1, 0, 0, 0, 1, 28, 1, 0, 0, 0, 1, 30, 1, 0, 0, 0, 1,
		32, 1, 0, 0, 0, 1, 34, 1, 0, 0, 0, 1, 36, 1, 0, 0, 0, 1, 38, 1, 0, 0, 0,
		1, 40, 1, 0, 0, 0, 1, 42, 1, 0, 0, 0, 1, 44, 1, 0, 0, 0, 2, 46, 1, 0, 0,
		0, 4, 58, 1, 0, 0, 0, 6, 69, 1, 0, 0, 0, 8, 75, 1, 0, 0, 0, 10, 92, 1,
		0, 0, 0, 12, 102, 1, 0, 0, 0, 14, 112, 1, 0, 0, 0, 16, 122, 1, 0, 0, 0,
		18, 141, 1, 0, 0, 0, 20, 157, 1, 0, 0, 0, 22, 180, 1, 0, 0, 0, 24, 204,
		1, 0, 0, 0, 26, 217, 1, 0, 0, 0, 28, 219, 1, 0, 0, 0, 30, 240, 1, 0, 0,
		0, 32, 254, 1, 0, 0, 0, 34, 268, 1, 0, 0, 0, 36, 282, 1, 0, 0, 0, 38, 302,
		1, 0, 0, 0, 40, 329, 1, 0, 0, 0, 42, 352, 1, 0, 0, 0, 44, 356, 1, 0, 0,
		0, 46, 50, 5, 35, 0, 0, 47, 49, 8, 0, 0, 0, 48, 47, 1, 0, 0, 0, 49, 52,
		1, 0, 0, 0, 50, 48, 1, 0, 0, 0, 50, 51, 1, 0, 0, 0, 51, 53, 1, 0, 0, 0,
		52, 50, 1, 0, 0, 0, 53, 54, 6, 0, 0, 0, 54, 3, 1, 0, 0, 0, 55, 57, 7, 1,
		0, 0, 56, 55, 1, 0, 0, 0, 57, 60, 1, 0, 0, 0, 58, 56, 1, 0, 0, 0, 58, 59,
		1, 0, 0, 0, 59, 62, 1, 0, 0, 0, 60, 58, 1, 0, 0, 0, 61, 63, 7, 0, 0, 0,
		62, 61, 1, 0, 0, 0, 63, 64, 1, 0, 0, 0, 64, 62, 1, 0, 0, 0, 64, 65, 1,
		0, 0, 0, 65, 66, 1, 0, 0, 0, 66, 67, 6, 1, 0, 0, 67, 5, 1, 0, 0, 0, 68,
		70, 7, 1, 0, 0, 69, 68, 1, 0, 0, 0, 70, 71, 1, 0, 0, 0, 71, 69, 1, 0, 0,
		0, 71, 72, 1, 0, 0, 0, 72, 73, 1, 0, 0, 0, 73, 74, 6, 2, 0, 0, 74, 7, 1,
		0, 0, 0, 75, 76, 5, 68, 0, 0, 76, 77, 5, 69, 0, 0, 77, 78, 5, 83, 0, 0,
		78, 79, 5, 67, 0, 0, 79, 80, 5, 82, 0, 0, 80, 81, 5, 73, 0, 0, 81, 82,
		5, 80, 0, 0, 82, 83, 5, 84, 0, 0, 83, 84, 5, 73, 0, 0, 84, 85, 5, 79, 0,
		0, 85, 86, 5, 78, 0, 0, 86, 88, 1, 0, 0, 0, 87, 89, 7, 1, 0, 0, 88, 87,
		1, 0, 0, 0, 89, 90, 1, 0, 0, 0, 90, 88, 1, 0, 0, 0, 90, 91, 1, 0, 0, 0,
		91, 9, 1, 0, 0, 0, 92, 93, 5, 84, 0, 0, 93, 94, 5, 65, 0, 0, 94, 95, 5,
		71, 0, 0, 95, 96, 5, 83, 0, 0, 96, 98, 1, 0, 0, 0, 97, 99, 7, 1, 0, 0,
		98, 97, 1, 0, 0, 0, 99, 100, 1, 0, 0, 0, 100, 98, 1, 0, 0, 0, 100, 101,
		1, 0, 0, 0, 101, 11, 1, 0, 0, 0, 102, 103, 5, 84, 0, 0, 103, 104, 5, 89,
		0, 0, 104, 105, 5, 80, 0, 0, 105, 106, 5, 69, 0, 0, 106, 108, 1, 0, 0,
		0, 107, 109, 7, 1, 0, 0, 108, 107, 1, 0, 0, 0, 109, 110, 1, 0, 0, 0, 110,
		108, 1, 0, 0, 0, 110, 111, 1, 0, 0, 0, 111, 13, 1, 0, 0, 0, 112, 113, 5,
		78, 0, 0, 113, 114, 5, 79, 0, 0, 114, 115, 5, 68, 0, 0, 115, 116, 5, 69,
		0, 0, 116, 118, 1, 0, 0, 0, 117, 119, 7, 1, 0, 0, 118, 117, 1, 0, 0, 0,
		119, 120, 1, 0, 0, 0, 120, 118, 1, 0, 0, 0, 120, 121, 1, 0, 0, 0, 121,
		15, 1, 0, 0, 0, 122, 123, 5, 67, 0, 0, 123, 124, 5, 79, 0, 0, 124, 125,
		5, 80, 0, 0, 125, 126, 5, 89, 0, 0, 126, 127, 5, 95, 0, 0, 127, 128, 5,
		83, 0, 0, 128, 129, 5, 67, 0, 0, 129, 130, 5, 72, 0, 0, 130, 131, 5, 69,
		0, 0, 131, 132, 5, 68, 0, 0, 132, 133, 5, 85, 0, 0, 133, 134, 5, 76, 0,
		0, 134, 135, 5, 69, 0, 0, 135, 137, 1, 0, 0, 0, 136, 138, 7, 1, 0, 0, 137,
		136, 1, 0, 0, 0, 138, 139, 1, 0, 0, 0, 139, 137, 1, 0, 0, 0, 139, 140,
		1, 0, 0, 0, 140, 17, 1, 0, 0, 0, 141, 142, 5, 68, 0, 0, 142, 143, 5, 65,
		0, 0, 143, 144, 5, 84, 0, 0, 144, 145, 5, 65, 0, 0, 145, 146, 5, 83, 0,
		0, 146, 147, 5, 79, 0, 0, 147, 148, 5, 85, 0, 0, 148, 149, 5, 82, 0, 0,
		149, 150, 5, 67, 0, 0, 150, 151, 5, 69, 0, 0, 151, 153, 1, 0, 0, 0, 152,
		154, 7, 1, 0, 0, 153, 152, 1, 0, 0, 0, 154, 155, 1, 0, 0, 0, 155, 153,
		1, 0, 0, 0, 155, 156, 1, 0, 0, 0, 156, 19, 1, 0, 0, 0, 157, 158, 5, 84,
		0, 0, 158, 159, 5, 65, 0, 0, 159, 160, 5, 82, 0, 0, 160, 161, 5, 71, 0,
		0, 161, 162, 5, 69, 0, 0, 162, 163, 5, 84, 0, 0, 163, 164, 5, 95, 0, 0,
		164, 165, 5, 68, 0, 0, 165, 166, 5, 65, 0, 0, 166, 167, 5, 84, 0, 0, 167,
		168, 5, 65, 0, 0, 168, 169, 5, 83, 0, 0, 169, 170, 5, 79, 0, 0, 170, 171,
		5, 85, 0, 0, 171, 172, 5, 82, 0, 0, 172, 173, 5, 67, 0, 0, 173, 174, 5,
		69, 0, 0, 174, 176, 1, 0, 0, 0, 175, 177, 7, 1, 0, 0, 176, 175, 1, 0, 0,
		0, 177, 178, 1, 0, 0, 0, 178, 176, 1, 0, 0, 0, 178, 179, 1, 0, 0, 0, 179,
		21, 1, 0, 0, 0, 180, 181, 5, 83, 0, 0, 181, 182, 5, 81, 0, 0, 182, 183,
		5, 76, 0, 0, 183, 187, 1, 0, 0, 0, 184, 186, 7, 1, 0, 0, 185, 184, 1, 0,
		0, 0, 186, 189, 1, 0, 0, 0, 187, 185, 1, 0, 0, 0, 187, 188, 1, 0, 0, 0,
		188, 190, 1, 0, 0, 0, 189, 187, 1, 0, 0, 0, 190, 194, 5, 62, 0, 0, 191,
		193, 8, 0, 0, 0, 192, 191, 1, 0, 0, 0, 193, 196, 1, 0, 0, 0, 194, 192,
		1, 0, 0, 0, 194, 195, 1, 0, 0, 0, 195, 198, 1, 0, 0, 0, 196, 194, 1, 0,
		0, 0, 197, 199, 7, 0, 0, 0, 198, 197, 1, 0, 0, 0, 199, 200, 1, 0, 0, 0,
		200, 198, 1, 0, 0, 0, 200, 201, 1, 0, 0, 0, 201, 202, 1, 0, 0, 0, 202,
		203, 6, 10, 1, 0, 203, 23, 1, 0, 0, 0, 204, 208, 8, 2, 0, 0, 205, 207,
		8, 0, 0, 0, 206, 205, 1, 0, 0, 0, 207, 210, 1, 0, 0, 0, 208, 206, 1, 0,
		0, 0, 208, 209, 1, 0, 0, 0, 209, 214, 1, 0, 0, 0, 210, 208, 1, 0, 0, 0,
		211, 213, 7, 0, 0, 0, 212, 211, 1, 0, 0, 0, 213, 216, 1, 0, 0, 0, 214,
		212, 1, 0, 0, 0, 214, 215, 1, 0, 0, 0, 215, 25, 1, 0, 0, 0, 216, 214, 1,
		0, 0, 0, 217, 218, 5, 0, 0, 0, 218, 27, 1, 0, 0, 0, 219, 220, 4, 13, 0,
		0, 220, 221, 5, 68, 0, 0, 221, 222, 5, 69, 0, 0, 222, 223, 5, 83, 0, 0,
		223, 224, 5, 67, 0, 0, 224, 225, 5, 82, 0, 0, 225, 226, 5, 73, 0, 0, 226,
		227, 5, 80, 0, 0, 227, 228, 5, 84, 0, 0, 228, 229, 5, 73, 0, 0, 229, 230,
		5, 79, 0, 0, 230, 231, 5, 78, 0, 0, 231, 233, 1, 0, 0, 0, 232, 234, 7,
		1, 0, 0, 233, 232, 1, 0, 0, 0, 234, 235, 1, 0, 0, 0, 235, 233, 1, 0, 0,
		0, 235, 236, 1, 0, 0, 0, 236, 237, 1, 0, 0, 0, 237, 238, 6, 13, 2, 0, 238,
		239, 6, 13, 3, 0, 239, 29, 1, 0, 0, 0, 240, 241, 4, 14, 1, 0, 241, 242,
		5, 84, 0, 0, 242, 243, 5, 65, 0, 0, 243, 244, 5, 71, 0, 0, 244, 245, 5,
		83, 0, 0, 245, 247, 1, 0, 0, 0, 246, 248, 7, 1, 0, 0, 247, 246, 1, 0, 0,
		0, 248, 249, 1, 0, 0, 0, 249, 247, 1, 0, 0, 0, 249, 250, 1, 0, 0, 0, 250,
		251, 1, 0, 0, 0, 251, 252, 6, 14, 4, 0, 252, 253, 6, 14, 3, 0, 253, 31,
		1, 0, 0, 0, 254, 255, 4, 15, 2, 0, 255, 256, 5, 84, 0, 0, 256, 257, 5,
		89, 0, 0, 257, 258, 5, 80, 0, 0, 258, 259, 5, 69, 0, 0, 259, 261, 1, 0,
		0, 0, 260, 262, 7, 1, 0, 0, 261, 260, 1, 0, 0, 0, 262, 263, 1, 0, 0, 0,
		263, 261, 1, 0, 0, 0, 263, 264, 1, 0, 0, 0, 264, 265, 1, 0, 0, 0, 265,
		266, 6, 15, 5, 0, 266, 267, 6, 15, 3, 0, 267, 33, 1, 0, 0, 0, 268, 269,
		4, 16, 3, 0, 269, 270, 5, 78, 0, 0, 270, 271, 5, 79, 0, 0, 271, 272, 5,
		68, 0, 0, 272, 273, 5, 69, 0, 0, 273, 275, 1, 0, 0, 0, 274, 276, 7, 1,
		0, 0, 275, 274, 1, 0, 0, 0, 276, 277, 1, 0, 0, 0, 277, 275, 1, 0, 0, 0,
		277, 278, 1, 0, 0, 0, 278, 279, 1, 0, 0, 0, 279, 280, 6, 16, 6, 0, 280,
		281, 6, 16, 3, 0, 281, 35, 1, 0, 0, 0, 282, 283, 4, 17, 4, 0, 283, 284,
		5, 68, 0, 0, 284, 285, 5, 65, 0, 0, 285, 286, 5, 84, 0, 0, 286, 287, 5,
		65, 0, 0, 287, 288, 5, 83, 0, 0, 288, 289, 5, 79, 0, 0, 289, 290, 5, 85,
		0, 0, 290, 291, 5, 82, 0, 0, 291, 292, 5, 67, 0, 0, 292, 293, 5, 69, 0,
		0, 293, 295, 1, 0, 0, 0, 294, 296, 7, 1, 0, 0, 295, 294, 1, 0, 0, 0, 296,
		297, 1, 0, 0, 0, 297, 295, 1, 0, 0, 0, 297, 298, 1, 0, 0, 0, 298, 299,
		1, 0, 0, 0, 299, 300, 6, 17, 7, 0, 300, 301, 6, 17, 3, 0, 301, 37, 1, 0,
		0, 0, 302, 303, 4, 18, 5, 0, 303, 304, 5, 84, 0, 0, 304, 305, 5, 65, 0,
		0, 305, 306, 5, 82, 0, 0, 306, 307, 5, 71, 0, 0, 307, 308, 5, 69, 0, 0,
		308, 309, 5, 84, 0, 0, 309, 310, 5, 95, 0, 0, 310, 311, 5, 68, 0, 0, 311,
		312, 5, 65, 0, 0, 312, 313, 5, 84, 0, 0, 313, 314, 5, 65, 0, 0, 314, 315,
		5, 83, 0, 0, 315, 316, 5, 79, 0, 0, 316, 317, 5, 85, 0, 0, 317, 318, 5,
		82, 0, 0, 318, 319, 5, 67, 0, 0, 319, 320, 5, 69, 0, 0, 320, 322, 1, 0,
		0, 0, 321, 323, 7, 1, 0, 0, 322, 321, 1, 0, 0, 0, 323, 324, 1, 0, 0, 0,
		324, 322, 1, 0, 0, 0, 324, 325, 1, 0, 0, 0, 325, 326, 1, 0, 0, 0, 326,
		327, 6, 18, 8, 0, 327, 328, 6, 18, 3, 0, 328, 39, 1, 0, 0, 0, 329, 330,
		4, 19, 6, 0, 330, 331, 5, 67, 0, 0, 331, 332, 5, 79, 0, 0, 332, 333, 5,
		80, 0, 0, 333, 334, 5, 89, 0, 0, 334, 335, 5, 95, 0, 0, 335, 336, 5, 83,
		0, 0, 336, 337, 5, 67, 0, 0, 337, 338, 5, 72, 0, 0, 338, 339, 5, 69, 0,
		0, 339, 340, 5, 68, 0, 0, 340, 341, 5, 85, 0, 0, 341, 342, 5, 76, 0, 0,
		342, 343, 5, 69, 0, 0, 343, 345, 1, 0, 0, 0, 344, 346, 7, 1, 0, 0, 345,
		344, 1, 0, 0, 0, 346, 347, 1, 0, 0, 0, 347, 345, 1, 0, 0, 0, 347, 348,
		1, 0, 0, 0, 348, 349, 1, 0, 0, 0, 349, 350, 6, 19, 9, 0, 350, 351, 6, 19,
		3, 0, 351, 41, 1, 0, 0, 0, 352, 353, 8, 0, 0, 0, 353, 354, 1, 0, 0, 0,
		354, 355, 6, 20, 10, 0, 355, 43, 1, 0, 0, 0, 356, 357, 7, 0, 0, 0, 357,
		358, 1, 0, 0, 0, 358, 359, 6, 21, 11, 0, 359, 45, 1, 0, 0, 0, 25, 0, 1,
		50, 58, 64, 71, 90, 100, 110, 120, 139, 155, 178, 187, 194, 200, 208, 214,
		235, 249, 263, 277, 297, 324, 347, 12, 6, 0, 0, 5, 1, 0, 7, 4, 0, 4, 0,
		0, 7, 5, 0, 7, 6, 0, 7, 7, 0, 7, 9, 0, 7, 10, 0, 7, 8, 0, 3, 0, 0, 7, 13,
		0,
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
	PipeLangLexerREST_OF_LINE         = 12
	PipeLangLexerSQL_LINE             = 13
)

// PipeLangLexerSQL_BODY_MODE is the PipeLangLexer mode.
const PipeLangLexerSQL_BODY_MODE = 1

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
