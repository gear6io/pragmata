// Code generated from PQL.g4 by ANTLR 4.13.2. DO NOT EDIT.

package pqlgrammar // PQL
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

type PQL struct {
	*antlr.BaseParser
}

var PQLParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func pqlParserInit() {
	staticData := &PQLParserStaticData
	staticData.LiteralNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "'{'", "'}'", "'('", "')'", "','", "':'", "'.'",
		"'..'", "'-'", "'+'", "'*'", "'/'", "'%'", "'|'", "'::'", "'='", "'!='",
		"'<='", "'>='", "'<'", "'>'", "", "", "", "", "", "' '",
	}
	staticData.SymbolicNames = []string{
		"", "COMMENT", "WS", "KW_FROM", "KW_FILTER", "KW_DERIVE", "KW_SELECT",
		"KW_GROUP", "KW_AGGREGATE", "KW_JOIN", "KW_SORT", "KW_TAKE", "KW_SKIP",
		"KW_WINDOW", "KW_ARRAY_JOIN", "KW_SIDE", "KW_INNER", "KW_LEFT", "KW_RIGHT",
		"KW_FULL", "KW_AS", "KW_FINAL", "LBRACE", "RBRACE", "LPAREN", "RPAREN",
		"COMMA", "COLON", "DOT", "RANGE", "MINUS", "PLUS", "STAR", "SLASH",
		"PERCENT", "PIPE", "CAST_OP", "EQ", "NEQ", "LTE", "GTE", "LT", "GT",
		"FLOAT", "INTEGER", "STRING", "IDENT", "NEWLINE", "FILTER_LINE",
	}
	staticData.RuleNames = []string{
		"pipeline", "transform", "fromTransform", "filterTransform", "filterBody",
		"deriveTransform", "selectTransform", "groupTransform", "joinTransform",
		"joinSide", "joinCond", "joinCondExpr", "joinCondToken", "joinCondInner",
		"arrayJoinTransform", "sortTransform", "takeTransform", "skipTransform",
		"windowTransform", "assignmentList", "assignment", "selectionList",
		"selectionItem", "keyList", "keyItem", "sortList", "sortItem", "opaqueExpr",
		"opaqueToken", "opaqueInner",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 48, 289, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 1, 0, 4, 0, 62, 8, 0,
		11, 0, 12, 0, 63, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 79, 8, 1, 1, 2, 1, 2, 1, 2, 3, 2, 84, 8,
		2, 1, 3, 1, 3, 1, 3, 1, 4, 4, 4, 90, 8, 4, 11, 4, 12, 4, 91, 1, 5, 1, 5,
		1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8,
		119, 8, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10,
		1, 10, 3, 10, 132, 8, 10, 1, 11, 4, 11, 135, 8, 11, 11, 11, 12, 11, 136,
		1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1,
		12, 1, 12, 1, 12, 1, 12, 5, 12, 153, 8, 12, 10, 12, 12, 12, 156, 9, 12,
		1, 12, 3, 12, 159, 8, 12, 1, 13, 1, 13, 3, 13, 163, 8, 13, 1, 14, 1, 14,
		1, 14, 1, 14, 3, 14, 169, 8, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1,
		16, 1, 16, 1, 16, 1, 16, 3, 16, 180, 8, 16, 1, 17, 1, 17, 1, 17, 1, 18,
		1, 18, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 5, 19, 193, 8, 19, 10,
		19, 12, 19, 196, 9, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21,
		5, 21, 205, 8, 21, 10, 21, 12, 21, 208, 9, 21, 1, 22, 1, 22, 1, 22, 1,
		22, 3, 22, 214, 8, 22, 1, 23, 1, 23, 1, 23, 5, 23, 219, 8, 23, 10, 23,
		12, 23, 222, 9, 23, 1, 24, 1, 24, 1, 24, 1, 24, 3, 24, 228, 8, 24, 1, 25,
		1, 25, 1, 25, 5, 25, 233, 8, 25, 10, 25, 12, 25, 236, 9, 25, 1, 26, 1,
		26, 1, 26, 1, 26, 1, 26, 3, 26, 243, 8, 26, 1, 27, 4, 27, 246, 8, 27, 11,
		27, 12, 27, 247, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28,
		1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1,
		28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 5, 28, 277,
		8, 28, 10, 28, 12, 28, 280, 9, 28, 1, 28, 3, 28, 283, 8, 28, 1, 29, 1,
		29, 3, 29, 287, 8, 29, 1, 29, 0, 0, 30, 0, 2, 4, 6, 8, 10, 12, 14, 16,
		18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52,
		54, 56, 58, 0, 1, 1, 0, 16, 19, 326, 0, 61, 1, 0, 0, 0, 2, 78, 1, 0, 0,
		0, 4, 80, 1, 0, 0, 0, 6, 85, 1, 0, 0, 0, 8, 89, 1, 0, 0, 0, 10, 93, 1,
		0, 0, 0, 12, 98, 1, 0, 0, 0, 14, 103, 1, 0, 0, 0, 16, 114, 1, 0, 0, 0,
		18, 125, 1, 0, 0, 0, 20, 131, 1, 0, 0, 0, 22, 134, 1, 0, 0, 0, 24, 158,
		1, 0, 0, 0, 26, 162, 1, 0, 0, 0, 28, 164, 1, 0, 0, 0, 30, 170, 1, 0, 0,
		0, 32, 175, 1, 0, 0, 0, 34, 181, 1, 0, 0, 0, 36, 184, 1, 0, 0, 0, 38, 189,
		1, 0, 0, 0, 40, 197, 1, 0, 0, 0, 42, 201, 1, 0, 0, 0, 44, 213, 1, 0, 0,
		0, 46, 215, 1, 0, 0, 0, 48, 227, 1, 0, 0, 0, 50, 229, 1, 0, 0, 0, 52, 242,
		1, 0, 0, 0, 54, 245, 1, 0, 0, 0, 56, 282, 1, 0, 0, 0, 58, 286, 1, 0, 0,
		0, 60, 62, 3, 2, 1, 0, 61, 60, 1, 0, 0, 0, 62, 63, 1, 0, 0, 0, 63, 61,
		1, 0, 0, 0, 63, 64, 1, 0, 0, 0, 64, 65, 1, 0, 0, 0, 65, 66, 5, 0, 0, 1,
		66, 1, 1, 0, 0, 0, 67, 79, 3, 4, 2, 0, 68, 79, 3, 6, 3, 0, 69, 79, 3, 10,
		5, 0, 70, 79, 3, 12, 6, 0, 71, 79, 3, 14, 7, 0, 72, 79, 3, 16, 8, 0, 73,
		79, 3, 28, 14, 0, 74, 79, 3, 30, 15, 0, 75, 79, 3, 32, 16, 0, 76, 79, 3,
		34, 17, 0, 77, 79, 3, 36, 18, 0, 78, 67, 1, 0, 0, 0, 78, 68, 1, 0, 0, 0,
		78, 69, 1, 0, 0, 0, 78, 70, 1, 0, 0, 0, 78, 71, 1, 0, 0, 0, 78, 72, 1,
		0, 0, 0, 78, 73, 1, 0, 0, 0, 78, 74, 1, 0, 0, 0, 78, 75, 1, 0, 0, 0, 78,
		76, 1, 0, 0, 0, 78, 77, 1, 0, 0, 0, 79, 3, 1, 0, 0, 0, 80, 81, 5, 3, 0,
		0, 81, 83, 5, 46, 0, 0, 82, 84, 5, 21, 0, 0, 83, 82, 1, 0, 0, 0, 83, 84,
		1, 0, 0, 0, 84, 5, 1, 0, 0, 0, 85, 86, 5, 4, 0, 0, 86, 87, 3, 8, 4, 0,
		87, 7, 1, 0, 0, 0, 88, 90, 5, 48, 0, 0, 89, 88, 1, 0, 0, 0, 90, 91, 1,
		0, 0, 0, 91, 89, 1, 0, 0, 0, 91, 92, 1, 0, 0, 0, 92, 9, 1, 0, 0, 0, 93,
		94, 5, 5, 0, 0, 94, 95, 5, 22, 0, 0, 95, 96, 3, 38, 19, 0, 96, 97, 5, 23,
		0, 0, 97, 11, 1, 0, 0, 0, 98, 99, 5, 6, 0, 0, 99, 100, 5, 22, 0, 0, 100,
		101, 3, 42, 21, 0, 101, 102, 5, 23, 0, 0, 102, 13, 1, 0, 0, 0, 103, 104,
		5, 7, 0, 0, 104, 105, 5, 22, 0, 0, 105, 106, 3, 46, 23, 0, 106, 107, 5,
		23, 0, 0, 107, 108, 5, 24, 0, 0, 108, 109, 5, 8, 0, 0, 109, 110, 5, 22,
		0, 0, 110, 111, 3, 38, 19, 0, 111, 112, 5, 23, 0, 0, 112, 113, 5, 25, 0,
		0, 113, 15, 1, 0, 0, 0, 114, 118, 5, 9, 0, 0, 115, 116, 5, 15, 0, 0, 116,
		117, 5, 27, 0, 0, 117, 119, 3, 18, 9, 0, 118, 115, 1, 0, 0, 0, 118, 119,
		1, 0, 0, 0, 119, 120, 1, 0, 0, 0, 120, 121, 5, 46, 0, 0, 121, 122, 5, 24,
		0, 0, 122, 123, 3, 20, 10, 0, 123, 124, 5, 25, 0, 0, 124, 17, 1, 0, 0,
		0, 125, 126, 7, 0, 0, 0, 126, 19, 1, 0, 0, 0, 127, 128, 5, 37, 0, 0, 128,
		129, 5, 37, 0, 0, 129, 132, 5, 46, 0, 0, 130, 132, 3, 22, 11, 0, 131, 127,
		1, 0, 0, 0, 131, 130, 1, 0, 0, 0, 132, 21, 1, 0, 0, 0, 133, 135, 3, 24,
		12, 0, 134, 133, 1, 0, 0, 0, 135, 136, 1, 0, 0, 0, 136, 134, 1, 0, 0, 0,
		136, 137, 1, 0, 0, 0, 137, 23, 1, 0, 0, 0, 138, 159, 5, 46, 0, 0, 139,
		159, 5, 28, 0, 0, 140, 159, 5, 44, 0, 0, 141, 159, 5, 43, 0, 0, 142, 159,
		5, 45, 0, 0, 143, 159, 5, 37, 0, 0, 144, 159, 5, 38, 0, 0, 145, 159, 5,
		41, 0, 0, 146, 159, 5, 42, 0, 0, 147, 159, 5, 39, 0, 0, 148, 159, 5, 40,
		0, 0, 149, 159, 5, 36, 0, 0, 150, 154, 5, 24, 0, 0, 151, 153, 3, 26, 13,
		0, 152, 151, 1, 0, 0, 0, 153, 156, 1, 0, 0, 0, 154, 152, 1, 0, 0, 0, 154,
		155, 1, 0, 0, 0, 155, 157, 1, 0, 0, 0, 156, 154, 1, 0, 0, 0, 157, 159,
		5, 25, 0, 0, 158, 138, 1, 0, 0, 0, 158, 139, 1, 0, 0, 0, 158, 140, 1, 0,
		0, 0, 158, 141, 1, 0, 0, 0, 158, 142, 1, 0, 0, 0, 158, 143, 1, 0, 0, 0,
		158, 144, 1, 0, 0, 0, 158, 145, 1, 0, 0, 0, 158, 146, 1, 0, 0, 0, 158,
		147, 1, 0, 0, 0, 158, 148, 1, 0, 0, 0, 158, 149, 1, 0, 0, 0, 158, 150,
		1, 0, 0, 0, 159, 25, 1, 0, 0, 0, 160, 163, 3, 24, 12, 0, 161, 163, 5, 26,
		0, 0, 162, 160, 1, 0, 0, 0, 162, 161, 1, 0, 0, 0, 163, 27, 1, 0, 0, 0,
		164, 165, 5, 14, 0, 0, 165, 168, 5, 46, 0, 0, 166, 167, 5, 20, 0, 0, 167,
		169, 5, 46, 0, 0, 168, 166, 1, 0, 0, 0, 168, 169, 1, 0, 0, 0, 169, 29,
		1, 0, 0, 0, 170, 171, 5, 10, 0, 0, 171, 172, 5, 22, 0, 0, 172, 173, 3,
		50, 25, 0, 173, 174, 5, 23, 0, 0, 174, 31, 1, 0, 0, 0, 175, 176, 5, 11,
		0, 0, 176, 179, 5, 44, 0, 0, 177, 178, 5, 29, 0, 0, 178, 180, 5, 44, 0,
		0, 179, 177, 1, 0, 0, 0, 179, 180, 1, 0, 0, 0, 180, 33, 1, 0, 0, 0, 181,
		182, 5, 12, 0, 0, 182, 183, 5, 44, 0, 0, 183, 35, 1, 0, 0, 0, 184, 185,
		5, 13, 0, 0, 185, 186, 5, 22, 0, 0, 186, 187, 3, 38, 19, 0, 187, 188, 5,
		23, 0, 0, 188, 37, 1, 0, 0, 0, 189, 194, 3, 40, 20, 0, 190, 191, 5, 26,
		0, 0, 191, 193, 3, 40, 20, 0, 192, 190, 1, 0, 0, 0, 193, 196, 1, 0, 0,
		0, 194, 192, 1, 0, 0, 0, 194, 195, 1, 0, 0, 0, 195, 39, 1, 0, 0, 0, 196,
		194, 1, 0, 0, 0, 197, 198, 5, 46, 0, 0, 198, 199, 5, 37, 0, 0, 199, 200,
		3, 54, 27, 0, 200, 41, 1, 0, 0, 0, 201, 206, 3, 44, 22, 0, 202, 203, 5,
		26, 0, 0, 203, 205, 3, 44, 22, 0, 204, 202, 1, 0, 0, 0, 205, 208, 1, 0,
		0, 0, 206, 204, 1, 0, 0, 0, 206, 207, 1, 0, 0, 0, 207, 43, 1, 0, 0, 0,
		208, 206, 1, 0, 0, 0, 209, 210, 5, 46, 0, 0, 210, 211, 5, 37, 0, 0, 211,
		214, 3, 54, 27, 0, 212, 214, 3, 54, 27, 0, 213, 209, 1, 0, 0, 0, 213, 212,
		1, 0, 0, 0, 214, 45, 1, 0, 0, 0, 215, 220, 3, 48, 24, 0, 216, 217, 5, 26,
		0, 0, 217, 219, 3, 48, 24, 0, 218, 216, 1, 0, 0, 0, 219, 222, 1, 0, 0,
		0, 220, 218, 1, 0, 0, 0, 220, 221, 1, 0, 0, 0, 221, 47, 1, 0, 0, 0, 222,
		220, 1, 0, 0, 0, 223, 224, 5, 46, 0, 0, 224, 225, 5, 37, 0, 0, 225, 228,
		3, 54, 27, 0, 226, 228, 5, 46, 0, 0, 227, 223, 1, 0, 0, 0, 227, 226, 1,
		0, 0, 0, 228, 49, 1, 0, 0, 0, 229, 234, 3, 52, 26, 0, 230, 231, 5, 26,
		0, 0, 231, 233, 3, 52, 26, 0, 232, 230, 1, 0, 0, 0, 233, 236, 1, 0, 0,
		0, 234, 232, 1, 0, 0, 0, 234, 235, 1, 0, 0, 0, 235, 51, 1, 0, 0, 0, 236,
		234, 1, 0, 0, 0, 237, 238, 5, 30, 0, 0, 238, 243, 5, 46, 0, 0, 239, 240,
		5, 31, 0, 0, 240, 243, 5, 46, 0, 0, 241, 243, 5, 46, 0, 0, 242, 237, 1,
		0, 0, 0, 242, 239, 1, 0, 0, 0, 242, 241, 1, 0, 0, 0, 243, 53, 1, 0, 0,
		0, 244, 246, 3, 56, 28, 0, 245, 244, 1, 0, 0, 0, 246, 247, 1, 0, 0, 0,
		247, 245, 1, 0, 0, 0, 247, 248, 1, 0, 0, 0, 248, 55, 1, 0, 0, 0, 249, 283,
		5, 46, 0, 0, 250, 283, 5, 44, 0, 0, 251, 283, 5, 43, 0, 0, 252, 283, 5,
		45, 0, 0, 253, 283, 5, 32, 0, 0, 254, 283, 5, 31, 0, 0, 255, 283, 5, 30,
		0, 0, 256, 283, 5, 33, 0, 0, 257, 283, 5, 34, 0, 0, 258, 283, 5, 35, 0,
		0, 259, 283, 5, 37, 0, 0, 260, 283, 5, 38, 0, 0, 261, 283, 5, 41, 0, 0,
		262, 283, 5, 42, 0, 0, 263, 283, 5, 39, 0, 0, 264, 283, 5, 40, 0, 0, 265,
		283, 5, 36, 0, 0, 266, 283, 5, 28, 0, 0, 267, 283, 5, 20, 0, 0, 268, 283,
		5, 21, 0, 0, 269, 283, 5, 16, 0, 0, 270, 283, 5, 17, 0, 0, 271, 283, 5,
		18, 0, 0, 272, 283, 5, 19, 0, 0, 273, 283, 5, 15, 0, 0, 274, 278, 5, 24,
		0, 0, 275, 277, 3, 58, 29, 0, 276, 275, 1, 0, 0, 0, 277, 280, 1, 0, 0,
		0, 278, 276, 1, 0, 0, 0, 278, 279, 1, 0, 0, 0, 279, 281, 1, 0, 0, 0, 280,
		278, 1, 0, 0, 0, 281, 283, 5, 25, 0, 0, 282, 249, 1, 0, 0, 0, 282, 250,
		1, 0, 0, 0, 282, 251, 1, 0, 0, 0, 282, 252, 1, 0, 0, 0, 282, 253, 1, 0,
		0, 0, 282, 254, 1, 0, 0, 0, 282, 255, 1, 0, 0, 0, 282, 256, 1, 0, 0, 0,
		282, 257, 1, 0, 0, 0, 282, 258, 1, 0, 0, 0, 282, 259, 1, 0, 0, 0, 282,
		260, 1, 0, 0, 0, 282, 261, 1, 0, 0, 0, 282, 262, 1, 0, 0, 0, 282, 263,
		1, 0, 0, 0, 282, 264, 1, 0, 0, 0, 282, 265, 1, 0, 0, 0, 282, 266, 1, 0,
		0, 0, 282, 267, 1, 0, 0, 0, 282, 268, 1, 0, 0, 0, 282, 269, 1, 0, 0, 0,
		282, 270, 1, 0, 0, 0, 282, 271, 1, 0, 0, 0, 282, 272, 1, 0, 0, 0, 282,
		273, 1, 0, 0, 0, 282, 274, 1, 0, 0, 0, 283, 57, 1, 0, 0, 0, 284, 287, 3,
		56, 28, 0, 285, 287, 5, 26, 0, 0, 286, 284, 1, 0, 0, 0, 286, 285, 1, 0,
		0, 0, 287, 59, 1, 0, 0, 0, 23, 63, 78, 83, 91, 118, 131, 136, 154, 158,
		162, 168, 179, 194, 206, 213, 220, 227, 234, 242, 247, 278, 282, 286,
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

// PQLInit initializes any static state used to implement PQL. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewPQL(). You can call this function if you wish to initialize the static state ahead
// of time.
func PQLInit() {
	staticData := &PQLParserStaticData
	staticData.once.Do(pqlParserInit)
}

// NewPQL produces a new parser instance for the optional input antlr.TokenStream.
func NewPQL(input antlr.TokenStream) *PQL {
	PQLInit()
	this := new(PQL)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &PQLParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "PQL.g4"

	return this
}

// PQL tokens.
const (
	PQLEOF           = antlr.TokenEOF
	PQLCOMMENT       = 1
	PQLWS            = 2
	PQLKW_FROM       = 3
	PQLKW_FILTER     = 4
	PQLKW_DERIVE     = 5
	PQLKW_SELECT     = 6
	PQLKW_GROUP      = 7
	PQLKW_AGGREGATE  = 8
	PQLKW_JOIN       = 9
	PQLKW_SORT       = 10
	PQLKW_TAKE       = 11
	PQLKW_SKIP       = 12
	PQLKW_WINDOW     = 13
	PQLKW_ARRAY_JOIN = 14
	PQLKW_SIDE       = 15
	PQLKW_INNER      = 16
	PQLKW_LEFT       = 17
	PQLKW_RIGHT      = 18
	PQLKW_FULL       = 19
	PQLKW_AS         = 20
	PQLKW_FINAL      = 21
	PQLLBRACE        = 22
	PQLRBRACE        = 23
	PQLLPAREN        = 24
	PQLRPAREN        = 25
	PQLCOMMA         = 26
	PQLCOLON         = 27
	PQLDOT           = 28
	PQLRANGE         = 29
	PQLMINUS         = 30
	PQLPLUS          = 31
	PQLSTAR          = 32
	PQLSLASH         = 33
	PQLPERCENT       = 34
	PQLPIPE          = 35
	PQLCAST_OP       = 36
	PQLEQ            = 37
	PQLNEQ           = 38
	PQLLTE           = 39
	PQLGTE           = 40
	PQLLT            = 41
	PQLGT            = 42
	PQLFLOAT         = 43
	PQLINTEGER       = 44
	PQLSTRING        = 45
	PQLIDENT         = 46
	PQLNEWLINE       = 47
	PQLFILTER_LINE   = 48
)

// PQL rules.
const (
	PQLRULE_pipeline           = 0
	PQLRULE_transform          = 1
	PQLRULE_fromTransform      = 2
	PQLRULE_filterTransform    = 3
	PQLRULE_filterBody         = 4
	PQLRULE_deriveTransform    = 5
	PQLRULE_selectTransform    = 6
	PQLRULE_groupTransform     = 7
	PQLRULE_joinTransform      = 8
	PQLRULE_joinSide           = 9
	PQLRULE_joinCond           = 10
	PQLRULE_joinCondExpr       = 11
	PQLRULE_joinCondToken      = 12
	PQLRULE_joinCondInner      = 13
	PQLRULE_arrayJoinTransform = 14
	PQLRULE_sortTransform      = 15
	PQLRULE_takeTransform      = 16
	PQLRULE_skipTransform      = 17
	PQLRULE_windowTransform    = 18
	PQLRULE_assignmentList     = 19
	PQLRULE_assignment         = 20
	PQLRULE_selectionList      = 21
	PQLRULE_selectionItem      = 22
	PQLRULE_keyList            = 23
	PQLRULE_keyItem            = 24
	PQLRULE_sortList           = 25
	PQLRULE_sortItem           = 26
	PQLRULE_opaqueExpr         = 27
	PQLRULE_opaqueToken        = 28
	PQLRULE_opaqueInner        = 29
)

// IPipelineContext is an interface to support dynamic dispatch.
type IPipelineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllTransform() []ITransformContext
	Transform(i int) ITransformContext

	// IsPipelineContext differentiates from other interfaces.
	IsPipelineContext()
}

type PipelineContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPipelineContext() *PipelineContext {
	var p = new(PipelineContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PQLRULE_pipeline
	return p
}

func InitEmptyPipelineContext(p *PipelineContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PQLRULE_pipeline
}

func (*PipelineContext) IsPipelineContext() {}

func NewPipelineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PipelineContext {
	var p = new(PipelineContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PQLRULE_pipeline

	return p
}

func (s *PipelineContext) GetParser() antlr.Parser { return s.parser }

func (s *PipelineContext) EOF() antlr.TerminalNode {
	return s.GetToken(PQLEOF, 0)
}

func (s *PipelineContext) AllTransform() []ITransformContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITransformContext); ok {
			len++
		}
	}

	tst := make([]ITransformContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITransformContext); ok {
			tst[i] = t.(ITransformContext)
			i++
		}
	}

	return tst
}

func (s *PipelineContext) Transform(i int) ITransformContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITransformContext); ok {
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

	return t.(ITransformContext)
}

func (s *PipelineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PipelineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PipelineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PQLListener); ok {
		listenerT.EnterPipeline(s)
	}
}

func (s *PipelineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PQLListener); ok {
		listenerT.ExitPipeline(s)
	}
}

func (s *PipelineContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PQLVisitor:
		return t.VisitPipeline(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PQL) Pipeline() (localctx IPipelineContext) {
	localctx = NewPipelineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, PQLRULE_pipeline)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(61)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&32504) != 0) {
		{
			p.SetState(60)
			p.Transform()
		}

		p.SetState(63)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(65)
		p.Match(PQLEOF)
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

// ITransformContext is an interface to support dynamic dispatch.
type ITransformContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FromTransform() IFromTransformContext
	FilterTransform() IFilterTransformContext
	DeriveTransform() IDeriveTransformContext
	SelectTransform() ISelectTransformContext
	GroupTransform() IGroupTransformContext
	JoinTransform() IJoinTransformContext
	ArrayJoinTransform() IArrayJoinTransformContext
	SortTransform() ISortTransformContext
	TakeTransform() ITakeTransformContext
	SkipTransform() ISkipTransformContext
	WindowTransform() IWindowTransformContext

	// IsTransformContext differentiates from other interfaces.
	IsTransformContext()
}

type TransformContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTransformContext() *TransformContext {
	var p = new(TransformContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PQLRULE_transform
	return p
}

func InitEmptyTransformContext(p *TransformContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PQLRULE_transform
}

func (*TransformContext) IsTransformContext() {}

func NewTransformContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TransformContext {
	var p = new(TransformContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PQLRULE_transform

	return p
}

func (s *TransformContext) GetParser() antlr.Parser { return s.parser }

func (s *TransformContext) FromTransform() IFromTransformContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFromTransformContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFromTransformContext)
}

func (s *TransformContext) FilterTransform() IFilterTransformContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFilterTransformContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFilterTransformContext)
}

func (s *TransformContext) DeriveTransform() IDeriveTransformContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeriveTransformContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeriveTransformContext)
}

func (s *TransformContext) SelectTransform() ISelectTransformContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelectTransformContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelectTransformContext)
}

func (s *TransformContext) GroupTransform() IGroupTransformContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGroupTransformContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGroupTransformContext)
}

func (s *TransformContext) JoinTransform() IJoinTransformContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJoinTransformContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJoinTransformContext)
}

func (s *TransformContext) ArrayJoinTransform() IArrayJoinTransformContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayJoinTransformContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayJoinTransformContext)
}

func (s *TransformContext) SortTransform() ISortTransformContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISortTransformContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISortTransformContext)
}

func (s *TransformContext) TakeTransform() ITakeTransformContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITakeTransformContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITakeTransformContext)
}

func (s *TransformContext) SkipTransform() ISkipTransformContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISkipTransformContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISkipTransformContext)
}

func (s *TransformContext) WindowTransform() IWindowTransformContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWindowTransformContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWindowTransformContext)
}

func (s *TransformContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TransformContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TransformContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PQLListener); ok {
		listenerT.EnterTransform(s)
	}
}

func (s *TransformContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PQLListener); ok {
		listenerT.ExitTransform(s)
	}
}

func (s *TransformContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PQLVisitor:
		return t.VisitTransform(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PQL) Transform() (localctx ITransformContext) {
	localctx = NewTransformContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, PQLRULE_transform)
	p.SetState(78)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PQLKW_FROM:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(67)
			p.FromTransform()
		}

	case PQLKW_FILTER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(68)
			p.FilterTransform()
		}

	case PQLKW_DERIVE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(69)
			p.DeriveTransform()
		}

	case PQLKW_SELECT:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(70)
			p.SelectTransform()
		}

	case PQLKW_GROUP:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(71)
			p.GroupTransform()
		}

	case PQLKW_JOIN:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(72)
			p.JoinTransform()
		}

	case PQLKW_ARRAY_JOIN:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(73)
			p.ArrayJoinTransform()
		}

	case PQLKW_SORT:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(74)
			p.SortTransform()
		}

	case PQLKW_TAKE:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(75)
			p.TakeTransform()
		}

	case PQLKW_SKIP:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(76)
			p.SkipTransform()
		}

	case PQLKW_WINDOW:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(77)
			p.WindowTransform()
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

// IFromTransformContext is an interface to support dynamic dispatch.
type IFromTransformContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	KW_FROM() antlr.TerminalNode
	IDENT() antlr.TerminalNode
	KW_FINAL() antlr.TerminalNode

	// IsFromTransformContext differentiates from other interfaces.
	IsFromTransformContext()
}

type FromTransformContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFromTransformContext() *FromTransformContext {
	var p = new(FromTransformContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PQLRULE_fromTransform
	return p
}

func InitEmptyFromTransformContext(p *FromTransformContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PQLRULE_fromTransform
}

func (*FromTransformContext) IsFromTransformContext() {}

func NewFromTransformContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FromTransformContext {
	var p = new(FromTransformContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PQLRULE_fromTransform

	return p
}

func (s *FromTransformContext) GetParser() antlr.Parser { return s.parser }

func (s *FromTransformContext) KW_FROM() antlr.TerminalNode {
	return s.GetToken(PQLKW_FROM, 0)
}

func (s *FromTransformContext) IDENT() antlr.TerminalNode {
	return s.GetToken(PQLIDENT, 0)
}

func (s *FromTransformContext) KW_FINAL() antlr.TerminalNode {
	return s.GetToken(PQLKW_FINAL, 0)
}

func (s *FromTransformContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FromTransformContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FromTransformContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PQLListener); ok {
		listenerT.EnterFromTransform(s)
	}
}

func (s *FromTransformContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PQLListener); ok {
		listenerT.ExitFromTransform(s)
	}
}

func (s *FromTransformContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PQLVisitor:
		return t.VisitFromTransform(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PQL) FromTransform() (localctx IFromTransformContext) {
	localctx = NewFromTransformContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, PQLRULE_fromTransform)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(80)
		p.Match(PQLKW_FROM)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(81)
		p.Match(PQLIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(83)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PQLKW_FINAL {
		{
			p.SetState(82)
			p.Match(PQLKW_FINAL)
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

// IFilterTransformContext is an interface to support dynamic dispatch.
type IFilterTransformContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	KW_FILTER() antlr.TerminalNode
	FilterBody() IFilterBodyContext

	// IsFilterTransformContext differentiates from other interfaces.
	IsFilterTransformContext()
}

type FilterTransformContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFilterTransformContext() *FilterTransformContext {
	var p = new(FilterTransformContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PQLRULE_filterTransform
	return p
}

func InitEmptyFilterTransformContext(p *FilterTransformContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PQLRULE_filterTransform
}

func (*FilterTransformContext) IsFilterTransformContext() {}

func NewFilterTransformContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FilterTransformContext {
	var p = new(FilterTransformContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PQLRULE_filterTransform

	return p
}

func (s *FilterTransformContext) GetParser() antlr.Parser { return s.parser }

func (s *FilterTransformContext) KW_FILTER() antlr.TerminalNode {
	return s.GetToken(PQLKW_FILTER, 0)
}

func (s *FilterTransformContext) FilterBody() IFilterBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFilterBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFilterBodyContext)
}

func (s *FilterTransformContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FilterTransformContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FilterTransformContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PQLListener); ok {
		listenerT.EnterFilterTransform(s)
	}
}

func (s *FilterTransformContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PQLListener); ok {
		listenerT.ExitFilterTransform(s)
	}
}

func (s *FilterTransformContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PQLVisitor:
		return t.VisitFilterTransform(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PQL) FilterTransform() (localctx IFilterTransformContext) {
	localctx = NewFilterTransformContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, PQLRULE_filterTransform)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(85)
		p.Match(PQLKW_FILTER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(86)
		p.FilterBody()
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

// IFilterBodyContext is an interface to support dynamic dispatch.
type IFilterBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllFILTER_LINE() []antlr.TerminalNode
	FILTER_LINE(i int) antlr.TerminalNode

	// IsFilterBodyContext differentiates from other interfaces.
	IsFilterBodyContext()
}

type FilterBodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFilterBodyContext() *FilterBodyContext {
	var p = new(FilterBodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PQLRULE_filterBody
	return p
}

func InitEmptyFilterBodyContext(p *FilterBodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PQLRULE_filterBody
}

func (*FilterBodyContext) IsFilterBodyContext() {}

func NewFilterBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FilterBodyContext {
	var p = new(FilterBodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PQLRULE_filterBody

	return p
}

func (s *FilterBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *FilterBodyContext) AllFILTER_LINE() []antlr.TerminalNode {
	return s.GetTokens(PQLFILTER_LINE)
}

func (s *FilterBodyContext) FILTER_LINE(i int) antlr.TerminalNode {
	return s.GetToken(PQLFILTER_LINE, i)
}

func (s *FilterBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FilterBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FilterBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PQLListener); ok {
		listenerT.EnterFilterBody(s)
	}
}

func (s *FilterBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PQLListener); ok {
		listenerT.ExitFilterBody(s)
	}
}

func (s *FilterBodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PQLVisitor:
		return t.VisitFilterBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PQL) FilterBody() (localctx IFilterBodyContext) {
	localctx = NewFilterBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, PQLRULE_filterBody)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(89)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == PQLFILTER_LINE {
		{
			p.SetState(88)
			p.Match(PQLFILTER_LINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(91)
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

// IDeriveTransformContext is an interface to support dynamic dispatch.
type IDeriveTransformContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	KW_DERIVE() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	AssignmentList() IAssignmentListContext
	RBRACE() antlr.TerminalNode

	// IsDeriveTransformContext differentiates from other interfaces.
	IsDeriveTransformContext()
}

type DeriveTransformContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeriveTransformContext() *DeriveTransformContext {
	var p = new(DeriveTransformContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PQLRULE_deriveTransform
	return p
}

func InitEmptyDeriveTransformContext(p *DeriveTransformContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PQLRULE_deriveTransform
}

func (*DeriveTransformContext) IsDeriveTransformContext() {}

func NewDeriveTransformContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeriveTransformContext {
	var p = new(DeriveTransformContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PQLRULE_deriveTransform

	return p
}

func (s *DeriveTransformContext) GetParser() antlr.Parser { return s.parser }

func (s *DeriveTransformContext) KW_DERIVE() antlr.TerminalNode {
	return s.GetToken(PQLKW_DERIVE, 0)
}

func (s *DeriveTransformContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(PQLLBRACE, 0)
}

func (s *DeriveTransformContext) AssignmentList() IAssignmentListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentListContext)
}

func (s *DeriveTransformContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(PQLRBRACE, 0)
}

func (s *DeriveTransformContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeriveTransformContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeriveTransformContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PQLListener); ok {
		listenerT.EnterDeriveTransform(s)
	}
}

func (s *DeriveTransformContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PQLListener); ok {
		listenerT.ExitDeriveTransform(s)
	}
}

func (s *DeriveTransformContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PQLVisitor:
		return t.VisitDeriveTransform(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PQL) DeriveTransform() (localctx IDeriveTransformContext) {
	localctx = NewDeriveTransformContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, PQLRULE_deriveTransform)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(93)
		p.Match(PQLKW_DERIVE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(94)
		p.Match(PQLLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(95)
		p.AssignmentList()
	}
	{
		p.SetState(96)
		p.Match(PQLRBRACE)
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

// ISelectTransformContext is an interface to support dynamic dispatch.
type ISelectTransformContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	KW_SELECT() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	SelectionList() ISelectionListContext
	RBRACE() antlr.TerminalNode

	// IsSelectTransformContext differentiates from other interfaces.
	IsSelectTransformContext()
}

type SelectTransformContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectTransformContext() *SelectTransformContext {
	var p = new(SelectTransformContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PQLRULE_selectTransform
	return p
}

func InitEmptySelectTransformContext(p *SelectTransformContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PQLRULE_selectTransform
}

func (*SelectTransformContext) IsSelectTransformContext() {}

func NewSelectTransformContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectTransformContext {
	var p = new(SelectTransformContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PQLRULE_selectTransform

	return p
}

func (s *SelectTransformContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectTransformContext) KW_SELECT() antlr.TerminalNode {
	return s.GetToken(PQLKW_SELECT, 0)
}

func (s *SelectTransformContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(PQLLBRACE, 0)
}

func (s *SelectTransformContext) SelectionList() ISelectionListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelectionListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelectionListContext)
}

func (s *SelectTransformContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(PQLRBRACE, 0)
}

func (s *SelectTransformContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectTransformContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectTransformContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PQLListener); ok {
		listenerT.EnterSelectTransform(s)
	}
}

func (s *SelectTransformContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PQLListener); ok {
		listenerT.ExitSelectTransform(s)
	}
}

func (s *SelectTransformContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PQLVisitor:
		return t.VisitSelectTransform(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PQL) SelectTransform() (localctx ISelectTransformContext) {
	localctx = NewSelectTransformContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, PQLRULE_selectTransform)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(98)
		p.Match(PQLKW_SELECT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(99)
		p.Match(PQLLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(100)
		p.SelectionList()
	}
	{
		p.SetState(101)
		p.Match(PQLRBRACE)
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

// IGroupTransformContext is an interface to support dynamic dispatch.
type IGroupTransformContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	KW_GROUP() antlr.TerminalNode
	AllLBRACE() []antlr.TerminalNode
	LBRACE(i int) antlr.TerminalNode
	KeyList() IKeyListContext
	AllRBRACE() []antlr.TerminalNode
	RBRACE(i int) antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	KW_AGGREGATE() antlr.TerminalNode
	AssignmentList() IAssignmentListContext
	RPAREN() antlr.TerminalNode

	// IsGroupTransformContext differentiates from other interfaces.
	IsGroupTransformContext()
}

type GroupTransformContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGroupTransformContext() *GroupTransformContext {
	var p = new(GroupTransformContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PQLRULE_groupTransform
	return p
}

func InitEmptyGroupTransformContext(p *GroupTransformContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PQLRULE_groupTransform
}

func (*GroupTransformContext) IsGroupTransformContext() {}

func NewGroupTransformContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GroupTransformContext {
	var p = new(GroupTransformContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PQLRULE_groupTransform

	return p
}

func (s *GroupTransformContext) GetParser() antlr.Parser { return s.parser }

func (s *GroupTransformContext) KW_GROUP() antlr.TerminalNode {
	return s.GetToken(PQLKW_GROUP, 0)
}

func (s *GroupTransformContext) AllLBRACE() []antlr.TerminalNode {
	return s.GetTokens(PQLLBRACE)
}

func (s *GroupTransformContext) LBRACE(i int) antlr.TerminalNode {
	return s.GetToken(PQLLBRACE, i)
}

func (s *GroupTransformContext) KeyList() IKeyListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKeyListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IKeyListContext)
}

func (s *GroupTransformContext) AllRBRACE() []antlr.TerminalNode {
	return s.GetTokens(PQLRBRACE)
}

func (s *GroupTransformContext) RBRACE(i int) antlr.TerminalNode {
	return s.GetToken(PQLRBRACE, i)
}

func (s *GroupTransformContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(PQLLPAREN, 0)
}

func (s *GroupTransformContext) KW_AGGREGATE() antlr.TerminalNode {
	return s.GetToken(PQLKW_AGGREGATE, 0)
}

func (s *GroupTransformContext) AssignmentList() IAssignmentListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentListContext)
}

func (s *GroupTransformContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(PQLRPAREN, 0)
}

func (s *GroupTransformContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GroupTransformContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GroupTransformContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PQLListener); ok {
		listenerT.EnterGroupTransform(s)
	}
}

func (s *GroupTransformContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PQLListener); ok {
		listenerT.ExitGroupTransform(s)
	}
}

func (s *GroupTransformContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PQLVisitor:
		return t.VisitGroupTransform(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PQL) GroupTransform() (localctx IGroupTransformContext) {
	localctx = NewGroupTransformContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, PQLRULE_groupTransform)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(103)
		p.Match(PQLKW_GROUP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(104)
		p.Match(PQLLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(105)
		p.KeyList()
	}
	{
		p.SetState(106)
		p.Match(PQLRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(107)
		p.Match(PQLLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(108)
		p.Match(PQLKW_AGGREGATE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(109)
		p.Match(PQLLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(110)
		p.AssignmentList()
	}
	{
		p.SetState(111)
		p.Match(PQLRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(112)
		p.Match(PQLRPAREN)
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

// IJoinTransformContext is an interface to support dynamic dispatch.
type IJoinTransformContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	KW_JOIN() antlr.TerminalNode
	IDENT() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	JoinCond() IJoinCondContext
	RPAREN() antlr.TerminalNode
	KW_SIDE() antlr.TerminalNode
	COLON() antlr.TerminalNode
	JoinSide() IJoinSideContext

	// IsJoinTransformContext differentiates from other interfaces.
	IsJoinTransformContext()
}

type JoinTransformContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJoinTransformContext() *JoinTransformContext {
	var p = new(JoinTransformContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PQLRULE_joinTransform
	return p
}

func InitEmptyJoinTransformContext(p *JoinTransformContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PQLRULE_joinTransform
}

func (*JoinTransformContext) IsJoinTransformContext() {}

func NewJoinTransformContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JoinTransformContext {
	var p = new(JoinTransformContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PQLRULE_joinTransform

	return p
}

func (s *JoinTransformContext) GetParser() antlr.Parser { return s.parser }

func (s *JoinTransformContext) KW_JOIN() antlr.TerminalNode {
	return s.GetToken(PQLKW_JOIN, 0)
}

func (s *JoinTransformContext) IDENT() antlr.TerminalNode {
	return s.GetToken(PQLIDENT, 0)
}

func (s *JoinTransformContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(PQLLPAREN, 0)
}

func (s *JoinTransformContext) JoinCond() IJoinCondContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJoinCondContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJoinCondContext)
}

func (s *JoinTransformContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(PQLRPAREN, 0)
}

func (s *JoinTransformContext) KW_SIDE() antlr.TerminalNode {
	return s.GetToken(PQLKW_SIDE, 0)
}

func (s *JoinTransformContext) COLON() antlr.TerminalNode {
	return s.GetToken(PQLCOLON, 0)
}

func (s *JoinTransformContext) JoinSide() IJoinSideContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJoinSideContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJoinSideContext)
}

func (s *JoinTransformContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JoinTransformContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JoinTransformContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PQLListener); ok {
		listenerT.EnterJoinTransform(s)
	}
}

func (s *JoinTransformContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PQLListener); ok {
		listenerT.ExitJoinTransform(s)
	}
}

func (s *JoinTransformContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PQLVisitor:
		return t.VisitJoinTransform(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PQL) JoinTransform() (localctx IJoinTransformContext) {
	localctx = NewJoinTransformContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, PQLRULE_joinTransform)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(114)
		p.Match(PQLKW_JOIN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(118)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PQLKW_SIDE {
		{
			p.SetState(115)
			p.Match(PQLKW_SIDE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(116)
			p.Match(PQLCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(117)
			p.JoinSide()
		}

	}
	{
		p.SetState(120)
		p.Match(PQLIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(121)
		p.Match(PQLLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(122)
		p.JoinCond()
	}
	{
		p.SetState(123)
		p.Match(PQLRPAREN)
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

// IJoinSideContext is an interface to support dynamic dispatch.
type IJoinSideContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	KW_LEFT() antlr.TerminalNode
	KW_RIGHT() antlr.TerminalNode
	KW_INNER() antlr.TerminalNode
	KW_FULL() antlr.TerminalNode

	// IsJoinSideContext differentiates from other interfaces.
	IsJoinSideContext()
}

type JoinSideContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJoinSideContext() *JoinSideContext {
	var p = new(JoinSideContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PQLRULE_joinSide
	return p
}

func InitEmptyJoinSideContext(p *JoinSideContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PQLRULE_joinSide
}

func (*JoinSideContext) IsJoinSideContext() {}

func NewJoinSideContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JoinSideContext {
	var p = new(JoinSideContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PQLRULE_joinSide

	return p
}

func (s *JoinSideContext) GetParser() antlr.Parser { return s.parser }

func (s *JoinSideContext) KW_LEFT() antlr.TerminalNode {
	return s.GetToken(PQLKW_LEFT, 0)
}

func (s *JoinSideContext) KW_RIGHT() antlr.TerminalNode {
	return s.GetToken(PQLKW_RIGHT, 0)
}

func (s *JoinSideContext) KW_INNER() antlr.TerminalNode {
	return s.GetToken(PQLKW_INNER, 0)
}

func (s *JoinSideContext) KW_FULL() antlr.TerminalNode {
	return s.GetToken(PQLKW_FULL, 0)
}

func (s *JoinSideContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JoinSideContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JoinSideContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PQLListener); ok {
		listenerT.EnterJoinSide(s)
	}
}

func (s *JoinSideContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PQLListener); ok {
		listenerT.ExitJoinSide(s)
	}
}

func (s *JoinSideContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PQLVisitor:
		return t.VisitJoinSide(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PQL) JoinSide() (localctx IJoinSideContext) {
	localctx = NewJoinSideContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, PQLRULE_joinSide)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(125)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&983040) != 0) {
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

// IJoinCondContext is an interface to support dynamic dispatch.
type IJoinCondContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsJoinCondContext differentiates from other interfaces.
	IsJoinCondContext()
}

type JoinCondContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJoinCondContext() *JoinCondContext {
	var p = new(JoinCondContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PQLRULE_joinCond
	return p
}

func InitEmptyJoinCondContext(p *JoinCondContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PQLRULE_joinCond
}

func (*JoinCondContext) IsJoinCondContext() {}

func NewJoinCondContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JoinCondContext {
	var p = new(JoinCondContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PQLRULE_joinCond

	return p
}

func (s *JoinCondContext) GetParser() antlr.Parser { return s.parser }

func (s *JoinCondContext) CopyAll(ctx *JoinCondContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *JoinCondContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JoinCondContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SelfJoinCondContext struct {
	JoinCondContext
}

func NewSelfJoinCondContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SelfJoinCondContext {
	var p = new(SelfJoinCondContext)

	InitEmptyJoinCondContext(&p.JoinCondContext)
	p.parser = parser
	p.CopyAll(ctx.(*JoinCondContext))

	return p
}

func (s *SelfJoinCondContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelfJoinCondContext) AllEQ() []antlr.TerminalNode {
	return s.GetTokens(PQLEQ)
}

func (s *SelfJoinCondContext) EQ(i int) antlr.TerminalNode {
	return s.GetToken(PQLEQ, i)
}

func (s *SelfJoinCondContext) IDENT() antlr.TerminalNode {
	return s.GetToken(PQLIDENT, 0)
}

func (s *SelfJoinCondContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PQLListener); ok {
		listenerT.EnterSelfJoinCond(s)
	}
}

func (s *SelfJoinCondContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PQLListener); ok {
		listenerT.ExitSelfJoinCond(s)
	}
}

func (s *SelfJoinCondContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PQLVisitor:
		return t.VisitSelfJoinCond(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExplicitJoinCondContext struct {
	JoinCondContext
}

func NewExplicitJoinCondContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExplicitJoinCondContext {
	var p = new(ExplicitJoinCondContext)

	InitEmptyJoinCondContext(&p.JoinCondContext)
	p.parser = parser
	p.CopyAll(ctx.(*JoinCondContext))

	return p
}

func (s *ExplicitJoinCondContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExplicitJoinCondContext) JoinCondExpr() IJoinCondExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJoinCondExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJoinCondExprContext)
}

func (s *ExplicitJoinCondContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PQLListener); ok {
		listenerT.EnterExplicitJoinCond(s)
	}
}

func (s *ExplicitJoinCondContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PQLListener); ok {
		listenerT.ExitExplicitJoinCond(s)
	}
}

func (s *ExplicitJoinCondContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PQLVisitor:
		return t.VisitExplicitJoinCond(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PQL) JoinCond() (localctx IJoinCondContext) {
	localctx = NewJoinCondContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, PQLRULE_joinCond)
	p.SetState(131)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		localctx = NewSelfJoinCondContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(127)
			p.Match(PQLEQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(128)
			p.Match(PQLEQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(129)
			p.Match(PQLIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewExplicitJoinCondContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(130)
			p.JoinCondExpr()
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

// IJoinCondExprContext is an interface to support dynamic dispatch.
type IJoinCondExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllJoinCondToken() []IJoinCondTokenContext
	JoinCondToken(i int) IJoinCondTokenContext

	// IsJoinCondExprContext differentiates from other interfaces.
	IsJoinCondExprContext()
}

type JoinCondExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJoinCondExprContext() *JoinCondExprContext {
	var p = new(JoinCondExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PQLRULE_joinCondExpr
	return p
}

func InitEmptyJoinCondExprContext(p *JoinCondExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PQLRULE_joinCondExpr
}

func (*JoinCondExprContext) IsJoinCondExprContext() {}

func NewJoinCondExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JoinCondExprContext {
	var p = new(JoinCondExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PQLRULE_joinCondExpr

	return p
}

func (s *JoinCondExprContext) GetParser() antlr.Parser { return s.parser }

func (s *JoinCondExprContext) AllJoinCondToken() []IJoinCondTokenContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IJoinCondTokenContext); ok {
			len++
		}
	}

	tst := make([]IJoinCondTokenContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IJoinCondTokenContext); ok {
			tst[i] = t.(IJoinCondTokenContext)
			i++
		}
	}

	return tst
}

func (s *JoinCondExprContext) JoinCondToken(i int) IJoinCondTokenContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJoinCondTokenContext); ok {
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

	return t.(IJoinCondTokenContext)
}

func (s *JoinCondExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JoinCondExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JoinCondExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PQLListener); ok {
		listenerT.EnterJoinCondExpr(s)
	}
}

func (s *JoinCondExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PQLListener); ok {
		listenerT.ExitJoinCondExpr(s)
	}
}

func (s *JoinCondExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PQLVisitor:
		return t.VisitJoinCondExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PQL) JoinCondExpr() (localctx IJoinCondExprContext) {
	localctx = NewJoinCondExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, PQLRULE_joinCondExpr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(134)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&140669054091264) != 0) {
		{
			p.SetState(133)
			p.JoinCondToken()
		}

		p.SetState(136)
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

// IJoinCondTokenContext is an interface to support dynamic dispatch.
type IJoinCondTokenContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENT() antlr.TerminalNode
	DOT() antlr.TerminalNode
	INTEGER() antlr.TerminalNode
	FLOAT() antlr.TerminalNode
	STRING() antlr.TerminalNode
	EQ() antlr.TerminalNode
	NEQ() antlr.TerminalNode
	LT() antlr.TerminalNode
	GT() antlr.TerminalNode
	LTE() antlr.TerminalNode
	GTE() antlr.TerminalNode
	CAST_OP() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	AllJoinCondInner() []IJoinCondInnerContext
	JoinCondInner(i int) IJoinCondInnerContext

	// IsJoinCondTokenContext differentiates from other interfaces.
	IsJoinCondTokenContext()
}

type JoinCondTokenContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJoinCondTokenContext() *JoinCondTokenContext {
	var p = new(JoinCondTokenContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PQLRULE_joinCondToken
	return p
}

func InitEmptyJoinCondTokenContext(p *JoinCondTokenContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PQLRULE_joinCondToken
}

func (*JoinCondTokenContext) IsJoinCondTokenContext() {}

func NewJoinCondTokenContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JoinCondTokenContext {
	var p = new(JoinCondTokenContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PQLRULE_joinCondToken

	return p
}

func (s *JoinCondTokenContext) GetParser() antlr.Parser { return s.parser }

func (s *JoinCondTokenContext) IDENT() antlr.TerminalNode {
	return s.GetToken(PQLIDENT, 0)
}

func (s *JoinCondTokenContext) DOT() antlr.TerminalNode {
	return s.GetToken(PQLDOT, 0)
}

func (s *JoinCondTokenContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(PQLINTEGER, 0)
}

func (s *JoinCondTokenContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(PQLFLOAT, 0)
}

func (s *JoinCondTokenContext) STRING() antlr.TerminalNode {
	return s.GetToken(PQLSTRING, 0)
}

func (s *JoinCondTokenContext) EQ() antlr.TerminalNode {
	return s.GetToken(PQLEQ, 0)
}

func (s *JoinCondTokenContext) NEQ() antlr.TerminalNode {
	return s.GetToken(PQLNEQ, 0)
}

func (s *JoinCondTokenContext) LT() antlr.TerminalNode {
	return s.GetToken(PQLLT, 0)
}

func (s *JoinCondTokenContext) GT() antlr.TerminalNode {
	return s.GetToken(PQLGT, 0)
}

func (s *JoinCondTokenContext) LTE() antlr.TerminalNode {
	return s.GetToken(PQLLTE, 0)
}

func (s *JoinCondTokenContext) GTE() antlr.TerminalNode {
	return s.GetToken(PQLGTE, 0)
}

func (s *JoinCondTokenContext) CAST_OP() antlr.TerminalNode {
	return s.GetToken(PQLCAST_OP, 0)
}

func (s *JoinCondTokenContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(PQLLPAREN, 0)
}

func (s *JoinCondTokenContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(PQLRPAREN, 0)
}

func (s *JoinCondTokenContext) AllJoinCondInner() []IJoinCondInnerContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IJoinCondInnerContext); ok {
			len++
		}
	}

	tst := make([]IJoinCondInnerContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IJoinCondInnerContext); ok {
			tst[i] = t.(IJoinCondInnerContext)
			i++
		}
	}

	return tst
}

func (s *JoinCondTokenContext) JoinCondInner(i int) IJoinCondInnerContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJoinCondInnerContext); ok {
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

	return t.(IJoinCondInnerContext)
}

func (s *JoinCondTokenContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JoinCondTokenContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JoinCondTokenContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PQLListener); ok {
		listenerT.EnterJoinCondToken(s)
	}
}

func (s *JoinCondTokenContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PQLListener); ok {
		listenerT.ExitJoinCondToken(s)
	}
}

func (s *JoinCondTokenContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PQLVisitor:
		return t.VisitJoinCondToken(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PQL) JoinCondToken() (localctx IJoinCondTokenContext) {
	localctx = NewJoinCondTokenContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, PQLRULE_joinCondToken)
	var _la int

	p.SetState(158)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PQLIDENT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(138)
			p.Match(PQLIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PQLDOT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(139)
			p.Match(PQLDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PQLINTEGER:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(140)
			p.Match(PQLINTEGER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PQLFLOAT:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(141)
			p.Match(PQLFLOAT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PQLSTRING:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(142)
			p.Match(PQLSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PQLEQ:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(143)
			p.Match(PQLEQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PQLNEQ:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(144)
			p.Match(PQLNEQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PQLLT:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(145)
			p.Match(PQLLT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PQLGT:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(146)
			p.Match(PQLGT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PQLLTE:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(147)
			p.Match(PQLLTE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PQLGTE:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(148)
			p.Match(PQLGTE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PQLCAST_OP:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(149)
			p.Match(PQLCAST_OP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PQLLPAREN:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(150)
			p.Match(PQLLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(154)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&140669121200128) != 0 {
			{
				p.SetState(151)
				p.JoinCondInner()
			}

			p.SetState(156)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(157)
			p.Match(PQLRPAREN)
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

// IJoinCondInnerContext is an interface to support dynamic dispatch.
type IJoinCondInnerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	JoinCondToken() IJoinCondTokenContext
	COMMA() antlr.TerminalNode

	// IsJoinCondInnerContext differentiates from other interfaces.
	IsJoinCondInnerContext()
}

type JoinCondInnerContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJoinCondInnerContext() *JoinCondInnerContext {
	var p = new(JoinCondInnerContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PQLRULE_joinCondInner
	return p
}

func InitEmptyJoinCondInnerContext(p *JoinCondInnerContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PQLRULE_joinCondInner
}

func (*JoinCondInnerContext) IsJoinCondInnerContext() {}

func NewJoinCondInnerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JoinCondInnerContext {
	var p = new(JoinCondInnerContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PQLRULE_joinCondInner

	return p
}

func (s *JoinCondInnerContext) GetParser() antlr.Parser { return s.parser }

func (s *JoinCondInnerContext) JoinCondToken() IJoinCondTokenContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJoinCondTokenContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJoinCondTokenContext)
}

func (s *JoinCondInnerContext) COMMA() antlr.TerminalNode {
	return s.GetToken(PQLCOMMA, 0)
}

func (s *JoinCondInnerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JoinCondInnerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JoinCondInnerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PQLListener); ok {
		listenerT.EnterJoinCondInner(s)
	}
}

func (s *JoinCondInnerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PQLListener); ok {
		listenerT.ExitJoinCondInner(s)
	}
}

func (s *JoinCondInnerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PQLVisitor:
		return t.VisitJoinCondInner(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PQL) JoinCondInner() (localctx IJoinCondInnerContext) {
	localctx = NewJoinCondInnerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, PQLRULE_joinCondInner)
	p.SetState(162)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PQLLPAREN, PQLDOT, PQLCAST_OP, PQLEQ, PQLNEQ, PQLLTE, PQLGTE, PQLLT, PQLGT, PQLFLOAT, PQLINTEGER, PQLSTRING, PQLIDENT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(160)
			p.JoinCondToken()
		}

	case PQLCOMMA:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(161)
			p.Match(PQLCOMMA)
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

// IArrayJoinTransformContext is an interface to support dynamic dispatch.
type IArrayJoinTransformContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	KW_ARRAY_JOIN() antlr.TerminalNode
	AllIDENT() []antlr.TerminalNode
	IDENT(i int) antlr.TerminalNode
	KW_AS() antlr.TerminalNode

	// IsArrayJoinTransformContext differentiates from other interfaces.
	IsArrayJoinTransformContext()
}

type ArrayJoinTransformContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayJoinTransformContext() *ArrayJoinTransformContext {
	var p = new(ArrayJoinTransformContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PQLRULE_arrayJoinTransform
	return p
}

func InitEmptyArrayJoinTransformContext(p *ArrayJoinTransformContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PQLRULE_arrayJoinTransform
}

func (*ArrayJoinTransformContext) IsArrayJoinTransformContext() {}

func NewArrayJoinTransformContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayJoinTransformContext {
	var p = new(ArrayJoinTransformContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PQLRULE_arrayJoinTransform

	return p
}

func (s *ArrayJoinTransformContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayJoinTransformContext) KW_ARRAY_JOIN() antlr.TerminalNode {
	return s.GetToken(PQLKW_ARRAY_JOIN, 0)
}

func (s *ArrayJoinTransformContext) AllIDENT() []antlr.TerminalNode {
	return s.GetTokens(PQLIDENT)
}

func (s *ArrayJoinTransformContext) IDENT(i int) antlr.TerminalNode {
	return s.GetToken(PQLIDENT, i)
}

func (s *ArrayJoinTransformContext) KW_AS() antlr.TerminalNode {
	return s.GetToken(PQLKW_AS, 0)
}

func (s *ArrayJoinTransformContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayJoinTransformContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayJoinTransformContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PQLListener); ok {
		listenerT.EnterArrayJoinTransform(s)
	}
}

func (s *ArrayJoinTransformContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PQLListener); ok {
		listenerT.ExitArrayJoinTransform(s)
	}
}

func (s *ArrayJoinTransformContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PQLVisitor:
		return t.VisitArrayJoinTransform(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PQL) ArrayJoinTransform() (localctx IArrayJoinTransformContext) {
	localctx = NewArrayJoinTransformContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, PQLRULE_arrayJoinTransform)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(164)
		p.Match(PQLKW_ARRAY_JOIN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(165)
		p.Match(PQLIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(168)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PQLKW_AS {
		{
			p.SetState(166)
			p.Match(PQLKW_AS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(167)
			p.Match(PQLIDENT)
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

// ISortTransformContext is an interface to support dynamic dispatch.
type ISortTransformContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	KW_SORT() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	SortList() ISortListContext
	RBRACE() antlr.TerminalNode

	// IsSortTransformContext differentiates from other interfaces.
	IsSortTransformContext()
}

type SortTransformContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySortTransformContext() *SortTransformContext {
	var p = new(SortTransformContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PQLRULE_sortTransform
	return p
}

func InitEmptySortTransformContext(p *SortTransformContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PQLRULE_sortTransform
}

func (*SortTransformContext) IsSortTransformContext() {}

func NewSortTransformContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SortTransformContext {
	var p = new(SortTransformContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PQLRULE_sortTransform

	return p
}

func (s *SortTransformContext) GetParser() antlr.Parser { return s.parser }

func (s *SortTransformContext) KW_SORT() antlr.TerminalNode {
	return s.GetToken(PQLKW_SORT, 0)
}

func (s *SortTransformContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(PQLLBRACE, 0)
}

func (s *SortTransformContext) SortList() ISortListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISortListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISortListContext)
}

func (s *SortTransformContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(PQLRBRACE, 0)
}

func (s *SortTransformContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SortTransformContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SortTransformContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PQLListener); ok {
		listenerT.EnterSortTransform(s)
	}
}

func (s *SortTransformContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PQLListener); ok {
		listenerT.ExitSortTransform(s)
	}
}

func (s *SortTransformContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PQLVisitor:
		return t.VisitSortTransform(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PQL) SortTransform() (localctx ISortTransformContext) {
	localctx = NewSortTransformContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, PQLRULE_sortTransform)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(170)
		p.Match(PQLKW_SORT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(171)
		p.Match(PQLLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(172)
		p.SortList()
	}
	{
		p.SetState(173)
		p.Match(PQLRBRACE)
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

// ITakeTransformContext is an interface to support dynamic dispatch.
type ITakeTransformContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	KW_TAKE() antlr.TerminalNode
	AllINTEGER() []antlr.TerminalNode
	INTEGER(i int) antlr.TerminalNode
	RANGE() antlr.TerminalNode

	// IsTakeTransformContext differentiates from other interfaces.
	IsTakeTransformContext()
}

type TakeTransformContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTakeTransformContext() *TakeTransformContext {
	var p = new(TakeTransformContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PQLRULE_takeTransform
	return p
}

func InitEmptyTakeTransformContext(p *TakeTransformContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PQLRULE_takeTransform
}

func (*TakeTransformContext) IsTakeTransformContext() {}

func NewTakeTransformContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TakeTransformContext {
	var p = new(TakeTransformContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PQLRULE_takeTransform

	return p
}

func (s *TakeTransformContext) GetParser() antlr.Parser { return s.parser }

func (s *TakeTransformContext) KW_TAKE() antlr.TerminalNode {
	return s.GetToken(PQLKW_TAKE, 0)
}

func (s *TakeTransformContext) AllINTEGER() []antlr.TerminalNode {
	return s.GetTokens(PQLINTEGER)
}

func (s *TakeTransformContext) INTEGER(i int) antlr.TerminalNode {
	return s.GetToken(PQLINTEGER, i)
}

func (s *TakeTransformContext) RANGE() antlr.TerminalNode {
	return s.GetToken(PQLRANGE, 0)
}

func (s *TakeTransformContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TakeTransformContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TakeTransformContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PQLListener); ok {
		listenerT.EnterTakeTransform(s)
	}
}

func (s *TakeTransformContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PQLListener); ok {
		listenerT.ExitTakeTransform(s)
	}
}

func (s *TakeTransformContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PQLVisitor:
		return t.VisitTakeTransform(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PQL) TakeTransform() (localctx ITakeTransformContext) {
	localctx = NewTakeTransformContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, PQLRULE_takeTransform)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(175)
		p.Match(PQLKW_TAKE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(176)
		p.Match(PQLINTEGER)
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

	if _la == PQLRANGE {
		{
			p.SetState(177)
			p.Match(PQLRANGE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(178)
			p.Match(PQLINTEGER)
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

// ISkipTransformContext is an interface to support dynamic dispatch.
type ISkipTransformContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	KW_SKIP() antlr.TerminalNode
	INTEGER() antlr.TerminalNode

	// IsSkipTransformContext differentiates from other interfaces.
	IsSkipTransformContext()
}

type SkipTransformContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySkipTransformContext() *SkipTransformContext {
	var p = new(SkipTransformContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PQLRULE_skipTransform
	return p
}

func InitEmptySkipTransformContext(p *SkipTransformContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PQLRULE_skipTransform
}

func (*SkipTransformContext) IsSkipTransformContext() {}

func NewSkipTransformContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SkipTransformContext {
	var p = new(SkipTransformContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PQLRULE_skipTransform

	return p
}

func (s *SkipTransformContext) GetParser() antlr.Parser { return s.parser }

func (s *SkipTransformContext) KW_SKIP() antlr.TerminalNode {
	return s.GetToken(PQLKW_SKIP, 0)
}

func (s *SkipTransformContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(PQLINTEGER, 0)
}

func (s *SkipTransformContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SkipTransformContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SkipTransformContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PQLListener); ok {
		listenerT.EnterSkipTransform(s)
	}
}

func (s *SkipTransformContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PQLListener); ok {
		listenerT.ExitSkipTransform(s)
	}
}

func (s *SkipTransformContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PQLVisitor:
		return t.VisitSkipTransform(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PQL) SkipTransform() (localctx ISkipTransformContext) {
	localctx = NewSkipTransformContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, PQLRULE_skipTransform)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(181)
		p.Match(PQLKW_SKIP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(182)
		p.Match(PQLINTEGER)
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

// IWindowTransformContext is an interface to support dynamic dispatch.
type IWindowTransformContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	KW_WINDOW() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	AssignmentList() IAssignmentListContext
	RBRACE() antlr.TerminalNode

	// IsWindowTransformContext differentiates from other interfaces.
	IsWindowTransformContext()
}

type WindowTransformContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWindowTransformContext() *WindowTransformContext {
	var p = new(WindowTransformContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PQLRULE_windowTransform
	return p
}

func InitEmptyWindowTransformContext(p *WindowTransformContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PQLRULE_windowTransform
}

func (*WindowTransformContext) IsWindowTransformContext() {}

func NewWindowTransformContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WindowTransformContext {
	var p = new(WindowTransformContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PQLRULE_windowTransform

	return p
}

func (s *WindowTransformContext) GetParser() antlr.Parser { return s.parser }

func (s *WindowTransformContext) KW_WINDOW() antlr.TerminalNode {
	return s.GetToken(PQLKW_WINDOW, 0)
}

func (s *WindowTransformContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(PQLLBRACE, 0)
}

func (s *WindowTransformContext) AssignmentList() IAssignmentListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentListContext)
}

func (s *WindowTransformContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(PQLRBRACE, 0)
}

func (s *WindowTransformContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WindowTransformContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WindowTransformContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PQLListener); ok {
		listenerT.EnterWindowTransform(s)
	}
}

func (s *WindowTransformContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PQLListener); ok {
		listenerT.ExitWindowTransform(s)
	}
}

func (s *WindowTransformContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PQLVisitor:
		return t.VisitWindowTransform(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PQL) WindowTransform() (localctx IWindowTransformContext) {
	localctx = NewWindowTransformContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, PQLRULE_windowTransform)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(184)
		p.Match(PQLKW_WINDOW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(185)
		p.Match(PQLLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(186)
		p.AssignmentList()
	}
	{
		p.SetState(187)
		p.Match(PQLRBRACE)
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

// IAssignmentListContext is an interface to support dynamic dispatch.
type IAssignmentListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllAssignment() []IAssignmentContext
	Assignment(i int) IAssignmentContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsAssignmentListContext differentiates from other interfaces.
	IsAssignmentListContext()
}

type AssignmentListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentListContext() *AssignmentListContext {
	var p = new(AssignmentListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PQLRULE_assignmentList
	return p
}

func InitEmptyAssignmentListContext(p *AssignmentListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PQLRULE_assignmentList
}

func (*AssignmentListContext) IsAssignmentListContext() {}

func NewAssignmentListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentListContext {
	var p = new(AssignmentListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PQLRULE_assignmentList

	return p
}

func (s *AssignmentListContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentListContext) AllAssignment() []IAssignmentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAssignmentContext); ok {
			len++
		}
	}

	tst := make([]IAssignmentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAssignmentContext); ok {
			tst[i] = t.(IAssignmentContext)
			i++
		}
	}

	return tst
}

func (s *AssignmentListContext) Assignment(i int) IAssignmentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentContext); ok {
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

	return t.(IAssignmentContext)
}

func (s *AssignmentListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(PQLCOMMA)
}

func (s *AssignmentListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(PQLCOMMA, i)
}

func (s *AssignmentListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PQLListener); ok {
		listenerT.EnterAssignmentList(s)
	}
}

func (s *AssignmentListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PQLListener); ok {
		listenerT.ExitAssignmentList(s)
	}
}

func (s *AssignmentListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PQLVisitor:
		return t.VisitAssignmentList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PQL) AssignmentList() (localctx IAssignmentListContext) {
	localctx = NewAssignmentListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, PQLRULE_assignmentList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(189)
		p.Assignment()
	}
	p.SetState(194)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == PQLCOMMA {
		{
			p.SetState(190)
			p.Match(PQLCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(191)
			p.Assignment()
		}

		p.SetState(196)
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

// IAssignmentContext is an interface to support dynamic dispatch.
type IAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENT() antlr.TerminalNode
	EQ() antlr.TerminalNode
	OpaqueExpr() IOpaqueExprContext

	// IsAssignmentContext differentiates from other interfaces.
	IsAssignmentContext()
}

type AssignmentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentContext() *AssignmentContext {
	var p = new(AssignmentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PQLRULE_assignment
	return p
}

func InitEmptyAssignmentContext(p *AssignmentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PQLRULE_assignment
}

func (*AssignmentContext) IsAssignmentContext() {}

func NewAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentContext {
	var p = new(AssignmentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PQLRULE_assignment

	return p
}

func (s *AssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentContext) IDENT() antlr.TerminalNode {
	return s.GetToken(PQLIDENT, 0)
}

func (s *AssignmentContext) EQ() antlr.TerminalNode {
	return s.GetToken(PQLEQ, 0)
}

func (s *AssignmentContext) OpaqueExpr() IOpaqueExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOpaqueExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOpaqueExprContext)
}

func (s *AssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PQLListener); ok {
		listenerT.EnterAssignment(s)
	}
}

func (s *AssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PQLListener); ok {
		listenerT.ExitAssignment(s)
	}
}

func (s *AssignmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PQLVisitor:
		return t.VisitAssignment(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PQL) Assignment() (localctx IAssignmentContext) {
	localctx = NewAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, PQLRULE_assignment)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(197)
		p.Match(PQLIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(198)
		p.Match(PQLEQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(199)
		p.OpaqueExpr()
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

// ISelectionListContext is an interface to support dynamic dispatch.
type ISelectionListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllSelectionItem() []ISelectionItemContext
	SelectionItem(i int) ISelectionItemContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsSelectionListContext differentiates from other interfaces.
	IsSelectionListContext()
}

type SelectionListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectionListContext() *SelectionListContext {
	var p = new(SelectionListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PQLRULE_selectionList
	return p
}

func InitEmptySelectionListContext(p *SelectionListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PQLRULE_selectionList
}

func (*SelectionListContext) IsSelectionListContext() {}

func NewSelectionListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectionListContext {
	var p = new(SelectionListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PQLRULE_selectionList

	return p
}

func (s *SelectionListContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectionListContext) AllSelectionItem() []ISelectionItemContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISelectionItemContext); ok {
			len++
		}
	}

	tst := make([]ISelectionItemContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISelectionItemContext); ok {
			tst[i] = t.(ISelectionItemContext)
			i++
		}
	}

	return tst
}

func (s *SelectionListContext) SelectionItem(i int) ISelectionItemContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelectionItemContext); ok {
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

	return t.(ISelectionItemContext)
}

func (s *SelectionListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(PQLCOMMA)
}

func (s *SelectionListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(PQLCOMMA, i)
}

func (s *SelectionListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectionListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectionListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PQLListener); ok {
		listenerT.EnterSelectionList(s)
	}
}

func (s *SelectionListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PQLListener); ok {
		listenerT.ExitSelectionList(s)
	}
}

func (s *SelectionListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PQLVisitor:
		return t.VisitSelectionList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PQL) SelectionList() (localctx ISelectionListContext) {
	localctx = NewSelectionListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, PQLRULE_selectionList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(201)
		p.SelectionItem()
	}
	p.SetState(206)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == PQLCOMMA {
		{
			p.SetState(202)
			p.Match(PQLCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(203)
			p.SelectionItem()
		}

		p.SetState(208)
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

// ISelectionItemContext is an interface to support dynamic dispatch.
type ISelectionItemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsSelectionItemContext differentiates from other interfaces.
	IsSelectionItemContext()
}

type SelectionItemContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectionItemContext() *SelectionItemContext {
	var p = new(SelectionItemContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PQLRULE_selectionItem
	return p
}

func InitEmptySelectionItemContext(p *SelectionItemContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PQLRULE_selectionItem
}

func (*SelectionItemContext) IsSelectionItemContext() {}

func NewSelectionItemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectionItemContext {
	var p = new(SelectionItemContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PQLRULE_selectionItem

	return p
}

func (s *SelectionItemContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectionItemContext) CopyAll(ctx *SelectionItemContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *SelectionItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectionItemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type AliasedSelectionContext struct {
	SelectionItemContext
}

func NewAliasedSelectionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AliasedSelectionContext {
	var p = new(AliasedSelectionContext)

	InitEmptySelectionItemContext(&p.SelectionItemContext)
	p.parser = parser
	p.CopyAll(ctx.(*SelectionItemContext))

	return p
}

func (s *AliasedSelectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AliasedSelectionContext) IDENT() antlr.TerminalNode {
	return s.GetToken(PQLIDENT, 0)
}

func (s *AliasedSelectionContext) EQ() antlr.TerminalNode {
	return s.GetToken(PQLEQ, 0)
}

func (s *AliasedSelectionContext) OpaqueExpr() IOpaqueExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOpaqueExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOpaqueExprContext)
}

func (s *AliasedSelectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PQLListener); ok {
		listenerT.EnterAliasedSelection(s)
	}
}

func (s *AliasedSelectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PQLListener); ok {
		listenerT.ExitAliasedSelection(s)
	}
}

func (s *AliasedSelectionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PQLVisitor:
		return t.VisitAliasedSelection(s)

	default:
		return t.VisitChildren(s)
	}
}

type BareSelectionContext struct {
	SelectionItemContext
}

func NewBareSelectionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BareSelectionContext {
	var p = new(BareSelectionContext)

	InitEmptySelectionItemContext(&p.SelectionItemContext)
	p.parser = parser
	p.CopyAll(ctx.(*SelectionItemContext))

	return p
}

func (s *BareSelectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BareSelectionContext) OpaqueExpr() IOpaqueExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOpaqueExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOpaqueExprContext)
}

func (s *BareSelectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PQLListener); ok {
		listenerT.EnterBareSelection(s)
	}
}

func (s *BareSelectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PQLListener); ok {
		listenerT.ExitBareSelection(s)
	}
}

func (s *BareSelectionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PQLVisitor:
		return t.VisitBareSelection(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PQL) SelectionItem() (localctx ISelectionItemContext) {
	localctx = NewSelectionItemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, PQLRULE_selectionItem)
	p.SetState(213)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		localctx = NewAliasedSelectionContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(209)
			p.Match(PQLIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(210)
			p.Match(PQLEQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(211)
			p.OpaqueExpr()
		}

	case 2:
		localctx = NewBareSelectionContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(212)
			p.OpaqueExpr()
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

// IKeyListContext is an interface to support dynamic dispatch.
type IKeyListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllKeyItem() []IKeyItemContext
	KeyItem(i int) IKeyItemContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsKeyListContext differentiates from other interfaces.
	IsKeyListContext()
}

type KeyListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeyListContext() *KeyListContext {
	var p = new(KeyListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PQLRULE_keyList
	return p
}

func InitEmptyKeyListContext(p *KeyListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PQLRULE_keyList
}

func (*KeyListContext) IsKeyListContext() {}

func NewKeyListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeyListContext {
	var p = new(KeyListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PQLRULE_keyList

	return p
}

func (s *KeyListContext) GetParser() antlr.Parser { return s.parser }

func (s *KeyListContext) AllKeyItem() []IKeyItemContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IKeyItemContext); ok {
			len++
		}
	}

	tst := make([]IKeyItemContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IKeyItemContext); ok {
			tst[i] = t.(IKeyItemContext)
			i++
		}
	}

	return tst
}

func (s *KeyListContext) KeyItem(i int) IKeyItemContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKeyItemContext); ok {
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

	return t.(IKeyItemContext)
}

func (s *KeyListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(PQLCOMMA)
}

func (s *KeyListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(PQLCOMMA, i)
}

func (s *KeyListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeyListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeyListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PQLListener); ok {
		listenerT.EnterKeyList(s)
	}
}

func (s *KeyListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PQLListener); ok {
		listenerT.ExitKeyList(s)
	}
}

func (s *KeyListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PQLVisitor:
		return t.VisitKeyList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PQL) KeyList() (localctx IKeyListContext) {
	localctx = NewKeyListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, PQLRULE_keyList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(215)
		p.KeyItem()
	}
	p.SetState(220)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == PQLCOMMA {
		{
			p.SetState(216)
			p.Match(PQLCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(217)
			p.KeyItem()
		}

		p.SetState(222)
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

// IKeyItemContext is an interface to support dynamic dispatch.
type IKeyItemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsKeyItemContext differentiates from other interfaces.
	IsKeyItemContext()
}

type KeyItemContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeyItemContext() *KeyItemContext {
	var p = new(KeyItemContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PQLRULE_keyItem
	return p
}

func InitEmptyKeyItemContext(p *KeyItemContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PQLRULE_keyItem
}

func (*KeyItemContext) IsKeyItemContext() {}

func NewKeyItemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeyItemContext {
	var p = new(KeyItemContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PQLRULE_keyItem

	return p
}

func (s *KeyItemContext) GetParser() antlr.Parser { return s.parser }

func (s *KeyItemContext) CopyAll(ctx *KeyItemContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *KeyItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeyItemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ComputedKeyContext struct {
	KeyItemContext
}

func NewComputedKeyContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ComputedKeyContext {
	var p = new(ComputedKeyContext)

	InitEmptyKeyItemContext(&p.KeyItemContext)
	p.parser = parser
	p.CopyAll(ctx.(*KeyItemContext))

	return p
}

func (s *ComputedKeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComputedKeyContext) IDENT() antlr.TerminalNode {
	return s.GetToken(PQLIDENT, 0)
}

func (s *ComputedKeyContext) EQ() antlr.TerminalNode {
	return s.GetToken(PQLEQ, 0)
}

func (s *ComputedKeyContext) OpaqueExpr() IOpaqueExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOpaqueExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOpaqueExprContext)
}

func (s *ComputedKeyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PQLListener); ok {
		listenerT.EnterComputedKey(s)
	}
}

func (s *ComputedKeyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PQLListener); ok {
		listenerT.ExitComputedKey(s)
	}
}

func (s *ComputedKeyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PQLVisitor:
		return t.VisitComputedKey(s)

	default:
		return t.VisitChildren(s)
	}
}

type ColumnKeyContext struct {
	KeyItemContext
}

func NewColumnKeyContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ColumnKeyContext {
	var p = new(ColumnKeyContext)

	InitEmptyKeyItemContext(&p.KeyItemContext)
	p.parser = parser
	p.CopyAll(ctx.(*KeyItemContext))

	return p
}

func (s *ColumnKeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ColumnKeyContext) IDENT() antlr.TerminalNode {
	return s.GetToken(PQLIDENT, 0)
}

func (s *ColumnKeyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PQLListener); ok {
		listenerT.EnterColumnKey(s)
	}
}

func (s *ColumnKeyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PQLListener); ok {
		listenerT.ExitColumnKey(s)
	}
}

func (s *ColumnKeyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PQLVisitor:
		return t.VisitColumnKey(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PQL) KeyItem() (localctx IKeyItemContext) {
	localctx = NewKeyItemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, PQLRULE_keyItem)
	p.SetState(227)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		localctx = NewComputedKeyContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(223)
			p.Match(PQLIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(224)
			p.Match(PQLEQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(225)
			p.OpaqueExpr()
		}

	case 2:
		localctx = NewColumnKeyContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(226)
			p.Match(PQLIDENT)
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

// ISortListContext is an interface to support dynamic dispatch.
type ISortListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllSortItem() []ISortItemContext
	SortItem(i int) ISortItemContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsSortListContext differentiates from other interfaces.
	IsSortListContext()
}

type SortListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySortListContext() *SortListContext {
	var p = new(SortListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PQLRULE_sortList
	return p
}

func InitEmptySortListContext(p *SortListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PQLRULE_sortList
}

func (*SortListContext) IsSortListContext() {}

func NewSortListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SortListContext {
	var p = new(SortListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PQLRULE_sortList

	return p
}

func (s *SortListContext) GetParser() antlr.Parser { return s.parser }

func (s *SortListContext) AllSortItem() []ISortItemContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISortItemContext); ok {
			len++
		}
	}

	tst := make([]ISortItemContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISortItemContext); ok {
			tst[i] = t.(ISortItemContext)
			i++
		}
	}

	return tst
}

func (s *SortListContext) SortItem(i int) ISortItemContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISortItemContext); ok {
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

	return t.(ISortItemContext)
}

func (s *SortListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(PQLCOMMA)
}

func (s *SortListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(PQLCOMMA, i)
}

func (s *SortListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SortListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SortListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PQLListener); ok {
		listenerT.EnterSortList(s)
	}
}

func (s *SortListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PQLListener); ok {
		listenerT.ExitSortList(s)
	}
}

func (s *SortListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PQLVisitor:
		return t.VisitSortList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PQL) SortList() (localctx ISortListContext) {
	localctx = NewSortListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, PQLRULE_sortList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(229)
		p.SortItem()
	}
	p.SetState(234)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == PQLCOMMA {
		{
			p.SetState(230)
			p.Match(PQLCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(231)
			p.SortItem()
		}

		p.SetState(236)
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

// ISortItemContext is an interface to support dynamic dispatch.
type ISortItemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsSortItemContext differentiates from other interfaces.
	IsSortItemContext()
}

type SortItemContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySortItemContext() *SortItemContext {
	var p = new(SortItemContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PQLRULE_sortItem
	return p
}

func InitEmptySortItemContext(p *SortItemContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PQLRULE_sortItem
}

func (*SortItemContext) IsSortItemContext() {}

func NewSortItemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SortItemContext {
	var p = new(SortItemContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PQLRULE_sortItem

	return p
}

func (s *SortItemContext) GetParser() antlr.Parser { return s.parser }

func (s *SortItemContext) CopyAll(ctx *SortItemContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *SortItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SortItemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type DescSortContext struct {
	SortItemContext
}

func NewDescSortContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DescSortContext {
	var p = new(DescSortContext)

	InitEmptySortItemContext(&p.SortItemContext)
	p.parser = parser
	p.CopyAll(ctx.(*SortItemContext))

	return p
}

func (s *DescSortContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DescSortContext) MINUS() antlr.TerminalNode {
	return s.GetToken(PQLMINUS, 0)
}

func (s *DescSortContext) IDENT() antlr.TerminalNode {
	return s.GetToken(PQLIDENT, 0)
}

func (s *DescSortContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PQLListener); ok {
		listenerT.EnterDescSort(s)
	}
}

func (s *DescSortContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PQLListener); ok {
		listenerT.ExitDescSort(s)
	}
}

func (s *DescSortContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PQLVisitor:
		return t.VisitDescSort(s)

	default:
		return t.VisitChildren(s)
	}
}

type AscSortContext struct {
	SortItemContext
}

func NewAscSortContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AscSortContext {
	var p = new(AscSortContext)

	InitEmptySortItemContext(&p.SortItemContext)
	p.parser = parser
	p.CopyAll(ctx.(*SortItemContext))

	return p
}

func (s *AscSortContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AscSortContext) IDENT() antlr.TerminalNode {
	return s.GetToken(PQLIDENT, 0)
}

func (s *AscSortContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PQLListener); ok {
		listenerT.EnterAscSort(s)
	}
}

func (s *AscSortContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PQLListener); ok {
		listenerT.ExitAscSort(s)
	}
}

func (s *AscSortContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PQLVisitor:
		return t.VisitAscSort(s)

	default:
		return t.VisitChildren(s)
	}
}

type AscSortExplicitContext struct {
	SortItemContext
}

func NewAscSortExplicitContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AscSortExplicitContext {
	var p = new(AscSortExplicitContext)

	InitEmptySortItemContext(&p.SortItemContext)
	p.parser = parser
	p.CopyAll(ctx.(*SortItemContext))

	return p
}

func (s *AscSortExplicitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AscSortExplicitContext) PLUS() antlr.TerminalNode {
	return s.GetToken(PQLPLUS, 0)
}

func (s *AscSortExplicitContext) IDENT() antlr.TerminalNode {
	return s.GetToken(PQLIDENT, 0)
}

func (s *AscSortExplicitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PQLListener); ok {
		listenerT.EnterAscSortExplicit(s)
	}
}

func (s *AscSortExplicitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PQLListener); ok {
		listenerT.ExitAscSortExplicit(s)
	}
}

func (s *AscSortExplicitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PQLVisitor:
		return t.VisitAscSortExplicit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PQL) SortItem() (localctx ISortItemContext) {
	localctx = NewSortItemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, PQLRULE_sortItem)
	p.SetState(242)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PQLMINUS:
		localctx = NewDescSortContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(237)
			p.Match(PQLMINUS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(238)
			p.Match(PQLIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PQLPLUS:
		localctx = NewAscSortExplicitContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(239)
			p.Match(PQLPLUS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(240)
			p.Match(PQLIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PQLIDENT:
		localctx = NewAscSortContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(241)
			p.Match(PQLIDENT)
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

// IOpaqueExprContext is an interface to support dynamic dispatch.
type IOpaqueExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllOpaqueToken() []IOpaqueTokenContext
	OpaqueToken(i int) IOpaqueTokenContext

	// IsOpaqueExprContext differentiates from other interfaces.
	IsOpaqueExprContext()
}

type OpaqueExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOpaqueExprContext() *OpaqueExprContext {
	var p = new(OpaqueExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PQLRULE_opaqueExpr
	return p
}

func InitEmptyOpaqueExprContext(p *OpaqueExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PQLRULE_opaqueExpr
}

func (*OpaqueExprContext) IsOpaqueExprContext() {}

func NewOpaqueExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OpaqueExprContext {
	var p = new(OpaqueExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PQLRULE_opaqueExpr

	return p
}

func (s *OpaqueExprContext) GetParser() antlr.Parser { return s.parser }

func (s *OpaqueExprContext) AllOpaqueToken() []IOpaqueTokenContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOpaqueTokenContext); ok {
			len++
		}
	}

	tst := make([]IOpaqueTokenContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOpaqueTokenContext); ok {
			tst[i] = t.(IOpaqueTokenContext)
			i++
		}
	}

	return tst
}

func (s *OpaqueExprContext) OpaqueToken(i int) IOpaqueTokenContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOpaqueTokenContext); ok {
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

	return t.(IOpaqueTokenContext)
}

func (s *OpaqueExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OpaqueExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OpaqueExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PQLListener); ok {
		listenerT.EnterOpaqueExpr(s)
	}
}

func (s *OpaqueExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PQLListener); ok {
		listenerT.ExitOpaqueExpr(s)
	}
}

func (s *OpaqueExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PQLVisitor:
		return t.VisitOpaqueExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PQL) OpaqueExpr() (localctx IOpaqueExprContext) {
	localctx = NewOpaqueExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, PQLRULE_opaqueExpr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(245)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&140736703987712) != 0) {
		{
			p.SetState(244)
			p.OpaqueToken()
		}

		p.SetState(247)
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

// IOpaqueTokenContext is an interface to support dynamic dispatch.
type IOpaqueTokenContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENT() antlr.TerminalNode
	INTEGER() antlr.TerminalNode
	FLOAT() antlr.TerminalNode
	STRING() antlr.TerminalNode
	STAR() antlr.TerminalNode
	PLUS() antlr.TerminalNode
	MINUS() antlr.TerminalNode
	SLASH() antlr.TerminalNode
	PERCENT() antlr.TerminalNode
	PIPE() antlr.TerminalNode
	EQ() antlr.TerminalNode
	NEQ() antlr.TerminalNode
	LT() antlr.TerminalNode
	GT() antlr.TerminalNode
	LTE() antlr.TerminalNode
	GTE() antlr.TerminalNode
	CAST_OP() antlr.TerminalNode
	DOT() antlr.TerminalNode
	KW_AS() antlr.TerminalNode
	KW_FINAL() antlr.TerminalNode
	KW_INNER() antlr.TerminalNode
	KW_LEFT() antlr.TerminalNode
	KW_RIGHT() antlr.TerminalNode
	KW_FULL() antlr.TerminalNode
	KW_SIDE() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	AllOpaqueInner() []IOpaqueInnerContext
	OpaqueInner(i int) IOpaqueInnerContext

	// IsOpaqueTokenContext differentiates from other interfaces.
	IsOpaqueTokenContext()
}

type OpaqueTokenContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOpaqueTokenContext() *OpaqueTokenContext {
	var p = new(OpaqueTokenContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PQLRULE_opaqueToken
	return p
}

func InitEmptyOpaqueTokenContext(p *OpaqueTokenContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PQLRULE_opaqueToken
}

func (*OpaqueTokenContext) IsOpaqueTokenContext() {}

func NewOpaqueTokenContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OpaqueTokenContext {
	var p = new(OpaqueTokenContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PQLRULE_opaqueToken

	return p
}

func (s *OpaqueTokenContext) GetParser() antlr.Parser { return s.parser }

func (s *OpaqueTokenContext) IDENT() antlr.TerminalNode {
	return s.GetToken(PQLIDENT, 0)
}

func (s *OpaqueTokenContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(PQLINTEGER, 0)
}

func (s *OpaqueTokenContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(PQLFLOAT, 0)
}

func (s *OpaqueTokenContext) STRING() antlr.TerminalNode {
	return s.GetToken(PQLSTRING, 0)
}

func (s *OpaqueTokenContext) STAR() antlr.TerminalNode {
	return s.GetToken(PQLSTAR, 0)
}

func (s *OpaqueTokenContext) PLUS() antlr.TerminalNode {
	return s.GetToken(PQLPLUS, 0)
}

func (s *OpaqueTokenContext) MINUS() antlr.TerminalNode {
	return s.GetToken(PQLMINUS, 0)
}

func (s *OpaqueTokenContext) SLASH() antlr.TerminalNode {
	return s.GetToken(PQLSLASH, 0)
}

func (s *OpaqueTokenContext) PERCENT() antlr.TerminalNode {
	return s.GetToken(PQLPERCENT, 0)
}

func (s *OpaqueTokenContext) PIPE() antlr.TerminalNode {
	return s.GetToken(PQLPIPE, 0)
}

func (s *OpaqueTokenContext) EQ() antlr.TerminalNode {
	return s.GetToken(PQLEQ, 0)
}

func (s *OpaqueTokenContext) NEQ() antlr.TerminalNode {
	return s.GetToken(PQLNEQ, 0)
}

func (s *OpaqueTokenContext) LT() antlr.TerminalNode {
	return s.GetToken(PQLLT, 0)
}

func (s *OpaqueTokenContext) GT() antlr.TerminalNode {
	return s.GetToken(PQLGT, 0)
}

func (s *OpaqueTokenContext) LTE() antlr.TerminalNode {
	return s.GetToken(PQLLTE, 0)
}

func (s *OpaqueTokenContext) GTE() antlr.TerminalNode {
	return s.GetToken(PQLGTE, 0)
}

func (s *OpaqueTokenContext) CAST_OP() antlr.TerminalNode {
	return s.GetToken(PQLCAST_OP, 0)
}

func (s *OpaqueTokenContext) DOT() antlr.TerminalNode {
	return s.GetToken(PQLDOT, 0)
}

func (s *OpaqueTokenContext) KW_AS() antlr.TerminalNode {
	return s.GetToken(PQLKW_AS, 0)
}

func (s *OpaqueTokenContext) KW_FINAL() antlr.TerminalNode {
	return s.GetToken(PQLKW_FINAL, 0)
}

func (s *OpaqueTokenContext) KW_INNER() antlr.TerminalNode {
	return s.GetToken(PQLKW_INNER, 0)
}

func (s *OpaqueTokenContext) KW_LEFT() antlr.TerminalNode {
	return s.GetToken(PQLKW_LEFT, 0)
}

func (s *OpaqueTokenContext) KW_RIGHT() antlr.TerminalNode {
	return s.GetToken(PQLKW_RIGHT, 0)
}

func (s *OpaqueTokenContext) KW_FULL() antlr.TerminalNode {
	return s.GetToken(PQLKW_FULL, 0)
}

func (s *OpaqueTokenContext) KW_SIDE() antlr.TerminalNode {
	return s.GetToken(PQLKW_SIDE, 0)
}

func (s *OpaqueTokenContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(PQLLPAREN, 0)
}

func (s *OpaqueTokenContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(PQLRPAREN, 0)
}

func (s *OpaqueTokenContext) AllOpaqueInner() []IOpaqueInnerContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOpaqueInnerContext); ok {
			len++
		}
	}

	tst := make([]IOpaqueInnerContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOpaqueInnerContext); ok {
			tst[i] = t.(IOpaqueInnerContext)
			i++
		}
	}

	return tst
}

func (s *OpaqueTokenContext) OpaqueInner(i int) IOpaqueInnerContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOpaqueInnerContext); ok {
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

	return t.(IOpaqueInnerContext)
}

func (s *OpaqueTokenContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OpaqueTokenContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OpaqueTokenContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PQLListener); ok {
		listenerT.EnterOpaqueToken(s)
	}
}

func (s *OpaqueTokenContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PQLListener); ok {
		listenerT.ExitOpaqueToken(s)
	}
}

func (s *OpaqueTokenContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PQLVisitor:
		return t.VisitOpaqueToken(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PQL) OpaqueToken() (localctx IOpaqueTokenContext) {
	localctx = NewOpaqueTokenContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, PQLRULE_opaqueToken)
	var _la int

	p.SetState(282)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PQLIDENT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(249)
			p.Match(PQLIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PQLINTEGER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(250)
			p.Match(PQLINTEGER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PQLFLOAT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(251)
			p.Match(PQLFLOAT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PQLSTRING:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(252)
			p.Match(PQLSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PQLSTAR:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(253)
			p.Match(PQLSTAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PQLPLUS:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(254)
			p.Match(PQLPLUS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PQLMINUS:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(255)
			p.Match(PQLMINUS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PQLSLASH:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(256)
			p.Match(PQLSLASH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PQLPERCENT:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(257)
			p.Match(PQLPERCENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PQLPIPE:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(258)
			p.Match(PQLPIPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PQLEQ:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(259)
			p.Match(PQLEQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PQLNEQ:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(260)
			p.Match(PQLNEQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PQLLT:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(261)
			p.Match(PQLLT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PQLGT:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(262)
			p.Match(PQLGT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PQLLTE:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(263)
			p.Match(PQLLTE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PQLGTE:
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(264)
			p.Match(PQLGTE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PQLCAST_OP:
		p.EnterOuterAlt(localctx, 17)
		{
			p.SetState(265)
			p.Match(PQLCAST_OP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PQLDOT:
		p.EnterOuterAlt(localctx, 18)
		{
			p.SetState(266)
			p.Match(PQLDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PQLKW_AS:
		p.EnterOuterAlt(localctx, 19)
		{
			p.SetState(267)
			p.Match(PQLKW_AS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PQLKW_FINAL:
		p.EnterOuterAlt(localctx, 20)
		{
			p.SetState(268)
			p.Match(PQLKW_FINAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PQLKW_INNER:
		p.EnterOuterAlt(localctx, 21)
		{
			p.SetState(269)
			p.Match(PQLKW_INNER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PQLKW_LEFT:
		p.EnterOuterAlt(localctx, 22)
		{
			p.SetState(270)
			p.Match(PQLKW_LEFT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PQLKW_RIGHT:
		p.EnterOuterAlt(localctx, 23)
		{
			p.SetState(271)
			p.Match(PQLKW_RIGHT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PQLKW_FULL:
		p.EnterOuterAlt(localctx, 24)
		{
			p.SetState(272)
			p.Match(PQLKW_FULL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PQLKW_SIDE:
		p.EnterOuterAlt(localctx, 25)
		{
			p.SetState(273)
			p.Match(PQLKW_SIDE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PQLLPAREN:
		p.EnterOuterAlt(localctx, 26)
		{
			p.SetState(274)
			p.Match(PQLLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(278)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&140736771096576) != 0 {
			{
				p.SetState(275)
				p.OpaqueInner()
			}

			p.SetState(280)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(281)
			p.Match(PQLRPAREN)
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

// IOpaqueInnerContext is an interface to support dynamic dispatch.
type IOpaqueInnerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OpaqueToken() IOpaqueTokenContext
	COMMA() antlr.TerminalNode

	// IsOpaqueInnerContext differentiates from other interfaces.
	IsOpaqueInnerContext()
}

type OpaqueInnerContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOpaqueInnerContext() *OpaqueInnerContext {
	var p = new(OpaqueInnerContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PQLRULE_opaqueInner
	return p
}

func InitEmptyOpaqueInnerContext(p *OpaqueInnerContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PQLRULE_opaqueInner
}

func (*OpaqueInnerContext) IsOpaqueInnerContext() {}

func NewOpaqueInnerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OpaqueInnerContext {
	var p = new(OpaqueInnerContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PQLRULE_opaqueInner

	return p
}

func (s *OpaqueInnerContext) GetParser() antlr.Parser { return s.parser }

func (s *OpaqueInnerContext) OpaqueToken() IOpaqueTokenContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOpaqueTokenContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOpaqueTokenContext)
}

func (s *OpaqueInnerContext) COMMA() antlr.TerminalNode {
	return s.GetToken(PQLCOMMA, 0)
}

func (s *OpaqueInnerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OpaqueInnerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OpaqueInnerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PQLListener); ok {
		listenerT.EnterOpaqueInner(s)
	}
}

func (s *OpaqueInnerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PQLListener); ok {
		listenerT.ExitOpaqueInner(s)
	}
}

func (s *OpaqueInnerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PQLVisitor:
		return t.VisitOpaqueInner(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PQL) OpaqueInner() (localctx IOpaqueInnerContext) {
	localctx = NewOpaqueInnerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, PQLRULE_opaqueInner)
	p.SetState(286)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PQLKW_SIDE, PQLKW_INNER, PQLKW_LEFT, PQLKW_RIGHT, PQLKW_FULL, PQLKW_AS, PQLKW_FINAL, PQLLPAREN, PQLDOT, PQLMINUS, PQLPLUS, PQLSTAR, PQLSLASH, PQLPERCENT, PQLPIPE, PQLCAST_OP, PQLEQ, PQLNEQ, PQLLTE, PQLGTE, PQLLT, PQLGT, PQLFLOAT, PQLINTEGER, PQLSTRING, PQLIDENT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(284)
			p.OpaqueToken()
		}

	case PQLCOMMA:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(285)
			p.Match(PQLCOMMA)
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
