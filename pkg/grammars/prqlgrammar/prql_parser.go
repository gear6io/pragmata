// Code generated from PRQL.g4 by ANTLR 4.13.2. DO NOT EDIT.

package prqlgrammar // PRQL
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

type PRQL struct {
	*antlr.BaseParser
}

var PRQLParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func prqlParserInit() {
	staticData := &PRQLParserStaticData
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
		"query", "clause", "fromClause", "filterClause", "filterBody", "deriveClause",
		"selectClause", "groupClause", "joinClause", "joinSide", "joinCond",
		"joinCondExpr", "joinCondToken", "joinCondInner", "arrayJoinClause",
		"sortClause", "takeClause", "skipClause", "windowClause", "assignmentList",
		"assignment", "selectionList", "selectionItem", "keyList", "keyItem",
		"sortList", "sortItem", "opaqueExpr", "opaqueToken", "opaqueInner",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 48, 290, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 1, 0, 4, 0, 62, 8, 0,
		11, 0, 12, 0, 63, 1, 0, 3, 0, 67, 8, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 80, 8, 1, 1, 2, 1, 2, 1, 2, 3, 2,
		85, 8, 2, 1, 3, 1, 3, 1, 3, 1, 4, 4, 4, 91, 8, 4, 11, 4, 12, 4, 92, 1,
		5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1,
		7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1,
		8, 3, 8, 120, 8, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1,
		10, 1, 10, 1, 10, 3, 10, 133, 8, 10, 1, 11, 4, 11, 136, 8, 11, 11, 11,
		12, 11, 137, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1,
		12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 5, 12, 154, 8, 12, 10, 12, 12, 12,
		157, 9, 12, 1, 12, 3, 12, 160, 8, 12, 1, 13, 1, 13, 3, 13, 164, 8, 13,
		1, 14, 1, 14, 1, 14, 1, 14, 3, 14, 170, 8, 14, 1, 15, 1, 15, 1, 15, 1,
		15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 3, 16, 181, 8, 16, 1, 17, 1, 17,
		1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 5, 19, 194,
		8, 19, 10, 19, 12, 19, 197, 9, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 21, 1,
		21, 1, 21, 5, 21, 206, 8, 21, 10, 21, 12, 21, 209, 9, 21, 1, 22, 1, 22,
		1, 22, 1, 22, 3, 22, 215, 8, 22, 1, 23, 1, 23, 1, 23, 5, 23, 220, 8, 23,
		10, 23, 12, 23, 223, 9, 23, 1, 24, 1, 24, 1, 24, 1, 24, 3, 24, 229, 8,
		24, 1, 25, 1, 25, 1, 25, 5, 25, 234, 8, 25, 10, 25, 12, 25, 237, 9, 25,
		1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 3, 26, 244, 8, 26, 1, 27, 4, 27, 247,
		8, 27, 11, 27, 12, 27, 248, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1,
		28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28,
		1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 5,
		28, 278, 8, 28, 10, 28, 12, 28, 281, 9, 28, 1, 28, 3, 28, 284, 8, 28, 1,
		29, 1, 29, 3, 29, 288, 8, 29, 1, 29, 0, 0, 30, 0, 2, 4, 6, 8, 10, 12, 14,
		16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50,
		52, 54, 56, 58, 0, 1, 1, 0, 16, 19, 328, 0, 61, 1, 0, 0, 0, 2, 79, 1, 0,
		0, 0, 4, 81, 1, 0, 0, 0, 6, 86, 1, 0, 0, 0, 8, 90, 1, 0, 0, 0, 10, 94,
		1, 0, 0, 0, 12, 99, 1, 0, 0, 0, 14, 104, 1, 0, 0, 0, 16, 115, 1, 0, 0,
		0, 18, 126, 1, 0, 0, 0, 20, 132, 1, 0, 0, 0, 22, 135, 1, 0, 0, 0, 24, 159,
		1, 0, 0, 0, 26, 163, 1, 0, 0, 0, 28, 165, 1, 0, 0, 0, 30, 171, 1, 0, 0,
		0, 32, 176, 1, 0, 0, 0, 34, 182, 1, 0, 0, 0, 36, 185, 1, 0, 0, 0, 38, 190,
		1, 0, 0, 0, 40, 198, 1, 0, 0, 0, 42, 202, 1, 0, 0, 0, 44, 214, 1, 0, 0,
		0, 46, 216, 1, 0, 0, 0, 48, 228, 1, 0, 0, 0, 50, 230, 1, 0, 0, 0, 52, 243,
		1, 0, 0, 0, 54, 246, 1, 0, 0, 0, 56, 283, 1, 0, 0, 0, 58, 287, 1, 0, 0,
		0, 60, 62, 3, 2, 1, 0, 61, 60, 1, 0, 0, 0, 62, 63, 1, 0, 0, 0, 63, 61,
		1, 0, 0, 0, 63, 64, 1, 0, 0, 0, 64, 66, 1, 0, 0, 0, 65, 67, 5, 0, 0, 1,
		66, 65, 1, 0, 0, 0, 66, 67, 1, 0, 0, 0, 67, 1, 1, 0, 0, 0, 68, 80, 3, 4,
		2, 0, 69, 80, 3, 6, 3, 0, 70, 80, 3, 10, 5, 0, 71, 80, 3, 12, 6, 0, 72,
		80, 3, 14, 7, 0, 73, 80, 3, 16, 8, 0, 74, 80, 3, 28, 14, 0, 75, 80, 3,
		30, 15, 0, 76, 80, 3, 32, 16, 0, 77, 80, 3, 34, 17, 0, 78, 80, 3, 36, 18,
		0, 79, 68, 1, 0, 0, 0, 79, 69, 1, 0, 0, 0, 79, 70, 1, 0, 0, 0, 79, 71,
		1, 0, 0, 0, 79, 72, 1, 0, 0, 0, 79, 73, 1, 0, 0, 0, 79, 74, 1, 0, 0, 0,
		79, 75, 1, 0, 0, 0, 79, 76, 1, 0, 0, 0, 79, 77, 1, 0, 0, 0, 79, 78, 1,
		0, 0, 0, 80, 3, 1, 0, 0, 0, 81, 82, 5, 3, 0, 0, 82, 84, 5, 46, 0, 0, 83,
		85, 5, 21, 0, 0, 84, 83, 1, 0, 0, 0, 84, 85, 1, 0, 0, 0, 85, 5, 1, 0, 0,
		0, 86, 87, 5, 4, 0, 0, 87, 88, 3, 8, 4, 0, 88, 7, 1, 0, 0, 0, 89, 91, 5,
		48, 0, 0, 90, 89, 1, 0, 0, 0, 91, 92, 1, 0, 0, 0, 92, 90, 1, 0, 0, 0, 92,
		93, 1, 0, 0, 0, 93, 9, 1, 0, 0, 0, 94, 95, 5, 5, 0, 0, 95, 96, 5, 22, 0,
		0, 96, 97, 3, 38, 19, 0, 97, 98, 5, 23, 0, 0, 98, 11, 1, 0, 0, 0, 99, 100,
		5, 6, 0, 0, 100, 101, 5, 22, 0, 0, 101, 102, 3, 42, 21, 0, 102, 103, 5,
		23, 0, 0, 103, 13, 1, 0, 0, 0, 104, 105, 5, 7, 0, 0, 105, 106, 5, 22, 0,
		0, 106, 107, 3, 46, 23, 0, 107, 108, 5, 23, 0, 0, 108, 109, 5, 24, 0, 0,
		109, 110, 5, 8, 0, 0, 110, 111, 5, 22, 0, 0, 111, 112, 3, 38, 19, 0, 112,
		113, 5, 23, 0, 0, 113, 114, 5, 25, 0, 0, 114, 15, 1, 0, 0, 0, 115, 119,
		5, 9, 0, 0, 116, 117, 5, 15, 0, 0, 117, 118, 5, 27, 0, 0, 118, 120, 3,
		18, 9, 0, 119, 116, 1, 0, 0, 0, 119, 120, 1, 0, 0, 0, 120, 121, 1, 0, 0,
		0, 121, 122, 5, 46, 0, 0, 122, 123, 5, 24, 0, 0, 123, 124, 3, 20, 10, 0,
		124, 125, 5, 25, 0, 0, 125, 17, 1, 0, 0, 0, 126, 127, 7, 0, 0, 0, 127,
		19, 1, 0, 0, 0, 128, 129, 5, 37, 0, 0, 129, 130, 5, 37, 0, 0, 130, 133,
		5, 46, 0, 0, 131, 133, 3, 22, 11, 0, 132, 128, 1, 0, 0, 0, 132, 131, 1,
		0, 0, 0, 133, 21, 1, 0, 0, 0, 134, 136, 3, 24, 12, 0, 135, 134, 1, 0, 0,
		0, 136, 137, 1, 0, 0, 0, 137, 135, 1, 0, 0, 0, 137, 138, 1, 0, 0, 0, 138,
		23, 1, 0, 0, 0, 139, 160, 5, 46, 0, 0, 140, 160, 5, 28, 0, 0, 141, 160,
		5, 44, 0, 0, 142, 160, 5, 43, 0, 0, 143, 160, 5, 45, 0, 0, 144, 160, 5,
		37, 0, 0, 145, 160, 5, 38, 0, 0, 146, 160, 5, 41, 0, 0, 147, 160, 5, 42,
		0, 0, 148, 160, 5, 39, 0, 0, 149, 160, 5, 40, 0, 0, 150, 160, 5, 36, 0,
		0, 151, 155, 5, 24, 0, 0, 152, 154, 3, 26, 13, 0, 153, 152, 1, 0, 0, 0,
		154, 157, 1, 0, 0, 0, 155, 153, 1, 0, 0, 0, 155, 156, 1, 0, 0, 0, 156,
		158, 1, 0, 0, 0, 157, 155, 1, 0, 0, 0, 158, 160, 5, 25, 0, 0, 159, 139,
		1, 0, 0, 0, 159, 140, 1, 0, 0, 0, 159, 141, 1, 0, 0, 0, 159, 142, 1, 0,
		0, 0, 159, 143, 1, 0, 0, 0, 159, 144, 1, 0, 0, 0, 159, 145, 1, 0, 0, 0,
		159, 146, 1, 0, 0, 0, 159, 147, 1, 0, 0, 0, 159, 148, 1, 0, 0, 0, 159,
		149, 1, 0, 0, 0, 159, 150, 1, 0, 0, 0, 159, 151, 1, 0, 0, 0, 160, 25, 1,
		0, 0, 0, 161, 164, 3, 24, 12, 0, 162, 164, 5, 26, 0, 0, 163, 161, 1, 0,
		0, 0, 163, 162, 1, 0, 0, 0, 164, 27, 1, 0, 0, 0, 165, 166, 5, 14, 0, 0,
		166, 169, 5, 46, 0, 0, 167, 168, 5, 20, 0, 0, 168, 170, 5, 46, 0, 0, 169,
		167, 1, 0, 0, 0, 169, 170, 1, 0, 0, 0, 170, 29, 1, 0, 0, 0, 171, 172, 5,
		10, 0, 0, 172, 173, 5, 22, 0, 0, 173, 174, 3, 50, 25, 0, 174, 175, 5, 23,
		0, 0, 175, 31, 1, 0, 0, 0, 176, 177, 5, 11, 0, 0, 177, 180, 5, 44, 0, 0,
		178, 179, 5, 29, 0, 0, 179, 181, 5, 44, 0, 0, 180, 178, 1, 0, 0, 0, 180,
		181, 1, 0, 0, 0, 181, 33, 1, 0, 0, 0, 182, 183, 5, 12, 0, 0, 183, 184,
		5, 44, 0, 0, 184, 35, 1, 0, 0, 0, 185, 186, 5, 13, 0, 0, 186, 187, 5, 22,
		0, 0, 187, 188, 3, 38, 19, 0, 188, 189, 5, 23, 0, 0, 189, 37, 1, 0, 0,
		0, 190, 195, 3, 40, 20, 0, 191, 192, 5, 26, 0, 0, 192, 194, 3, 40, 20,
		0, 193, 191, 1, 0, 0, 0, 194, 197, 1, 0, 0, 0, 195, 193, 1, 0, 0, 0, 195,
		196, 1, 0, 0, 0, 196, 39, 1, 0, 0, 0, 197, 195, 1, 0, 0, 0, 198, 199, 5,
		46, 0, 0, 199, 200, 5, 37, 0, 0, 200, 201, 3, 54, 27, 0, 201, 41, 1, 0,
		0, 0, 202, 207, 3, 44, 22, 0, 203, 204, 5, 26, 0, 0, 204, 206, 3, 44, 22,
		0, 205, 203, 1, 0, 0, 0, 206, 209, 1, 0, 0, 0, 207, 205, 1, 0, 0, 0, 207,
		208, 1, 0, 0, 0, 208, 43, 1, 0, 0, 0, 209, 207, 1, 0, 0, 0, 210, 211, 5,
		46, 0, 0, 211, 212, 5, 37, 0, 0, 212, 215, 3, 54, 27, 0, 213, 215, 3, 54,
		27, 0, 214, 210, 1, 0, 0, 0, 214, 213, 1, 0, 0, 0, 215, 45, 1, 0, 0, 0,
		216, 221, 3, 48, 24, 0, 217, 218, 5, 26, 0, 0, 218, 220, 3, 48, 24, 0,
		219, 217, 1, 0, 0, 0, 220, 223, 1, 0, 0, 0, 221, 219, 1, 0, 0, 0, 221,
		222, 1, 0, 0, 0, 222, 47, 1, 0, 0, 0, 223, 221, 1, 0, 0, 0, 224, 225, 5,
		46, 0, 0, 225, 226, 5, 37, 0, 0, 226, 229, 3, 54, 27, 0, 227, 229, 5, 46,
		0, 0, 228, 224, 1, 0, 0, 0, 228, 227, 1, 0, 0, 0, 229, 49, 1, 0, 0, 0,
		230, 235, 3, 52, 26, 0, 231, 232, 5, 26, 0, 0, 232, 234, 3, 52, 26, 0,
		233, 231, 1, 0, 0, 0, 234, 237, 1, 0, 0, 0, 235, 233, 1, 0, 0, 0, 235,
		236, 1, 0, 0, 0, 236, 51, 1, 0, 0, 0, 237, 235, 1, 0, 0, 0, 238, 239, 5,
		30, 0, 0, 239, 244, 5, 46, 0, 0, 240, 241, 5, 31, 0, 0, 241, 244, 5, 46,
		0, 0, 242, 244, 5, 46, 0, 0, 243, 238, 1, 0, 0, 0, 243, 240, 1, 0, 0, 0,
		243, 242, 1, 0, 0, 0, 244, 53, 1, 0, 0, 0, 245, 247, 3, 56, 28, 0, 246,
		245, 1, 0, 0, 0, 247, 248, 1, 0, 0, 0, 248, 246, 1, 0, 0, 0, 248, 249,
		1, 0, 0, 0, 249, 55, 1, 0, 0, 0, 250, 284, 5, 46, 0, 0, 251, 284, 5, 44,
		0, 0, 252, 284, 5, 43, 0, 0, 253, 284, 5, 45, 0, 0, 254, 284, 5, 32, 0,
		0, 255, 284, 5, 31, 0, 0, 256, 284, 5, 30, 0, 0, 257, 284, 5, 33, 0, 0,
		258, 284, 5, 34, 0, 0, 259, 284, 5, 35, 0, 0, 260, 284, 5, 37, 0, 0, 261,
		284, 5, 38, 0, 0, 262, 284, 5, 41, 0, 0, 263, 284, 5, 42, 0, 0, 264, 284,
		5, 39, 0, 0, 265, 284, 5, 40, 0, 0, 266, 284, 5, 36, 0, 0, 267, 284, 5,
		28, 0, 0, 268, 284, 5, 20, 0, 0, 269, 284, 5, 21, 0, 0, 270, 284, 5, 16,
		0, 0, 271, 284, 5, 17, 0, 0, 272, 284, 5, 18, 0, 0, 273, 284, 5, 19, 0,
		0, 274, 284, 5, 15, 0, 0, 275, 279, 5, 24, 0, 0, 276, 278, 3, 58, 29, 0,
		277, 276, 1, 0, 0, 0, 278, 281, 1, 0, 0, 0, 279, 277, 1, 0, 0, 0, 279,
		280, 1, 0, 0, 0, 280, 282, 1, 0, 0, 0, 281, 279, 1, 0, 0, 0, 282, 284,
		5, 25, 0, 0, 283, 250, 1, 0, 0, 0, 283, 251, 1, 0, 0, 0, 283, 252, 1, 0,
		0, 0, 283, 253, 1, 0, 0, 0, 283, 254, 1, 0, 0, 0, 283, 255, 1, 0, 0, 0,
		283, 256, 1, 0, 0, 0, 283, 257, 1, 0, 0, 0, 283, 258, 1, 0, 0, 0, 283,
		259, 1, 0, 0, 0, 283, 260, 1, 0, 0, 0, 283, 261, 1, 0, 0, 0, 283, 262,
		1, 0, 0, 0, 283, 263, 1, 0, 0, 0, 283, 264, 1, 0, 0, 0, 283, 265, 1, 0,
		0, 0, 283, 266, 1, 0, 0, 0, 283, 267, 1, 0, 0, 0, 283, 268, 1, 0, 0, 0,
		283, 269, 1, 0, 0, 0, 283, 270, 1, 0, 0, 0, 283, 271, 1, 0, 0, 0, 283,
		272, 1, 0, 0, 0, 283, 273, 1, 0, 0, 0, 283, 274, 1, 0, 0, 0, 283, 275,
		1, 0, 0, 0, 284, 57, 1, 0, 0, 0, 285, 288, 3, 56, 28, 0, 286, 288, 5, 26,
		0, 0, 287, 285, 1, 0, 0, 0, 287, 286, 1, 0, 0, 0, 288, 59, 1, 0, 0, 0,
		24, 63, 66, 79, 84, 92, 119, 132, 137, 155, 159, 163, 169, 180, 195, 207,
		214, 221, 228, 235, 243, 248, 279, 283, 287,
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

// PRQLInit initializes any static state used to implement PRQL. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewPRQL(). You can call this function if you wish to initialize the static state ahead
// of time.
func PRQLInit() {
	staticData := &PRQLParserStaticData
	staticData.once.Do(prqlParserInit)
}

// NewPRQL produces a new parser instance for the optional input antlr.TokenStream.
func NewPRQL(input antlr.TokenStream) *PRQL {
	PRQLInit()
	this := new(PRQL)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &PRQLParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "PRQL.g4"

	return this
}

// PRQL tokens.
const (
	PRQLEOF           = antlr.TokenEOF
	PRQLCOMMENT       = 1
	PRQLWS            = 2
	PRQLKW_FROM       = 3
	PRQLKW_FILTER     = 4
	PRQLKW_DERIVE     = 5
	PRQLKW_SELECT     = 6
	PRQLKW_GROUP      = 7
	PRQLKW_AGGREGATE  = 8
	PRQLKW_JOIN       = 9
	PRQLKW_SORT       = 10
	PRQLKW_TAKE       = 11
	PRQLKW_SKIP       = 12
	PRQLKW_WINDOW     = 13
	PRQLKW_ARRAY_JOIN = 14
	PRQLKW_SIDE       = 15
	PRQLKW_INNER      = 16
	PRQLKW_LEFT       = 17
	PRQLKW_RIGHT      = 18
	PRQLKW_FULL       = 19
	PRQLKW_AS         = 20
	PRQLKW_FINAL      = 21
	PRQLLBRACE        = 22
	PRQLRBRACE        = 23
	PRQLLPAREN        = 24
	PRQLRPAREN        = 25
	PRQLCOMMA         = 26
	PRQLCOLON         = 27
	PRQLDOT           = 28
	PRQLRANGE         = 29
	PRQLMINUS         = 30
	PRQLPLUS          = 31
	PRQLSTAR          = 32
	PRQLSLASH         = 33
	PRQLPERCENT       = 34
	PRQLPIPE          = 35
	PRQLCAST_OP       = 36
	PRQLEQ            = 37
	PRQLNEQ           = 38
	PRQLLTE           = 39
	PRQLGTE           = 40
	PRQLLT            = 41
	PRQLGT            = 42
	PRQLFLOAT         = 43
	PRQLINTEGER       = 44
	PRQLSTRING        = 45
	PRQLIDENT         = 46
	PRQLNEWLINE       = 47
	PRQLFILTER_LINE   = 48
)

// PRQL rules.
const (
	PRQLRULE_query           = 0
	PRQLRULE_clause          = 1
	PRQLRULE_fromClause      = 2
	PRQLRULE_filterClause    = 3
	PRQLRULE_filterBody      = 4
	PRQLRULE_deriveClause    = 5
	PRQLRULE_selectClause    = 6
	PRQLRULE_groupClause     = 7
	PRQLRULE_joinClause      = 8
	PRQLRULE_joinSide        = 9
	PRQLRULE_joinCond        = 10
	PRQLRULE_joinCondExpr    = 11
	PRQLRULE_joinCondToken   = 12
	PRQLRULE_joinCondInner   = 13
	PRQLRULE_arrayJoinClause = 14
	PRQLRULE_sortClause      = 15
	PRQLRULE_takeClause      = 16
	PRQLRULE_skipClause      = 17
	PRQLRULE_windowClause    = 18
	PRQLRULE_assignmentList  = 19
	PRQLRULE_assignment      = 20
	PRQLRULE_selectionList   = 21
	PRQLRULE_selectionItem   = 22
	PRQLRULE_keyList         = 23
	PRQLRULE_keyItem         = 24
	PRQLRULE_sortList        = 25
	PRQLRULE_sortItem        = 26
	PRQLRULE_opaqueExpr      = 27
	PRQLRULE_opaqueToken     = 28
	PRQLRULE_opaqueInner     = 29
)

// IQueryContext is an interface to support dynamic dispatch.
type IQueryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllClause() []IClauseContext
	Clause(i int) IClauseContext
	EOF() antlr.TerminalNode

	// IsQueryContext differentiates from other interfaces.
	IsQueryContext()
}

type QueryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQueryContext() *QueryContext {
	var p = new(QueryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PRQLRULE_query
	return p
}

func InitEmptyQueryContext(p *QueryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PRQLRULE_query
}

func (*QueryContext) IsQueryContext() {}

func NewQueryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QueryContext {
	var p = new(QueryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PRQLRULE_query

	return p
}

func (s *QueryContext) GetParser() antlr.Parser { return s.parser }

func (s *QueryContext) AllClause() []IClauseContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IClauseContext); ok {
			len++
		}
	}

	tst := make([]IClauseContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IClauseContext); ok {
			tst[i] = t.(IClauseContext)
			i++
		}
	}

	return tst
}

func (s *QueryContext) Clause(i int) IClauseContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IClauseContext); ok {
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

	return t.(IClauseContext)
}

func (s *QueryContext) EOF() antlr.TerminalNode {
	return s.GetToken(PRQLEOF, 0)
}

func (s *QueryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QueryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QueryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PRQLListener); ok {
		listenerT.EnterQuery(s)
	}
}

func (s *QueryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PRQLListener); ok {
		listenerT.ExitQuery(s)
	}
}

func (s *QueryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PRQLVisitor:
		return t.VisitQuery(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PRQL) Query() (localctx IQueryContext) {
	localctx = NewQueryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, PRQLRULE_query)
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
			p.Clause()
		}

		p.SetState(63)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(66)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(65)
			p.Match(PRQLEOF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
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

// IClauseContext is an interface to support dynamic dispatch.
type IClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FromClause() IFromClauseContext
	FilterClause() IFilterClauseContext
	DeriveClause() IDeriveClauseContext
	SelectClause() ISelectClauseContext
	GroupClause() IGroupClauseContext
	JoinClause() IJoinClauseContext
	ArrayJoinClause() IArrayJoinClauseContext
	SortClause() ISortClauseContext
	TakeClause() ITakeClauseContext
	SkipClause() ISkipClauseContext
	WindowClause() IWindowClauseContext

	// IsClauseContext differentiates from other interfaces.
	IsClauseContext()
}

type ClauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClauseContext() *ClauseContext {
	var p = new(ClauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PRQLRULE_clause
	return p
}

func InitEmptyClauseContext(p *ClauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PRQLRULE_clause
}

func (*ClauseContext) IsClauseContext() {}

func NewClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClauseContext {
	var p = new(ClauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PRQLRULE_clause

	return p
}

func (s *ClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *ClauseContext) FromClause() IFromClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFromClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFromClauseContext)
}

func (s *ClauseContext) FilterClause() IFilterClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFilterClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFilterClauseContext)
}

func (s *ClauseContext) DeriveClause() IDeriveClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeriveClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeriveClauseContext)
}

func (s *ClauseContext) SelectClause() ISelectClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelectClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelectClauseContext)
}

func (s *ClauseContext) GroupClause() IGroupClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGroupClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGroupClauseContext)
}

func (s *ClauseContext) JoinClause() IJoinClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJoinClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJoinClauseContext)
}

func (s *ClauseContext) ArrayJoinClause() IArrayJoinClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayJoinClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayJoinClauseContext)
}

func (s *ClauseContext) SortClause() ISortClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISortClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISortClauseContext)
}

func (s *ClauseContext) TakeClause() ITakeClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITakeClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITakeClauseContext)
}

func (s *ClauseContext) SkipClause() ISkipClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISkipClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISkipClauseContext)
}

func (s *ClauseContext) WindowClause() IWindowClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWindowClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWindowClauseContext)
}

func (s *ClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PRQLListener); ok {
		listenerT.EnterClause(s)
	}
}

func (s *ClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PRQLListener); ok {
		listenerT.ExitClause(s)
	}
}

func (s *ClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PRQLVisitor:
		return t.VisitClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PRQL) Clause() (localctx IClauseContext) {
	localctx = NewClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, PRQLRULE_clause)
	p.SetState(79)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PRQLKW_FROM:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(68)
			p.FromClause()
		}

	case PRQLKW_FILTER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(69)
			p.FilterClause()
		}

	case PRQLKW_DERIVE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(70)
			p.DeriveClause()
		}

	case PRQLKW_SELECT:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(71)
			p.SelectClause()
		}

	case PRQLKW_GROUP:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(72)
			p.GroupClause()
		}

	case PRQLKW_JOIN:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(73)
			p.JoinClause()
		}

	case PRQLKW_ARRAY_JOIN:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(74)
			p.ArrayJoinClause()
		}

	case PRQLKW_SORT:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(75)
			p.SortClause()
		}

	case PRQLKW_TAKE:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(76)
			p.TakeClause()
		}

	case PRQLKW_SKIP:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(77)
			p.SkipClause()
		}

	case PRQLKW_WINDOW:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(78)
			p.WindowClause()
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

// IFromClauseContext is an interface to support dynamic dispatch.
type IFromClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	KW_FROM() antlr.TerminalNode
	IDENT() antlr.TerminalNode
	KW_FINAL() antlr.TerminalNode

	// IsFromClauseContext differentiates from other interfaces.
	IsFromClauseContext()
}

type FromClauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFromClauseContext() *FromClauseContext {
	var p = new(FromClauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PRQLRULE_fromClause
	return p
}

func InitEmptyFromClauseContext(p *FromClauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PRQLRULE_fromClause
}

func (*FromClauseContext) IsFromClauseContext() {}

func NewFromClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FromClauseContext {
	var p = new(FromClauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PRQLRULE_fromClause

	return p
}

func (s *FromClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *FromClauseContext) KW_FROM() antlr.TerminalNode {
	return s.GetToken(PRQLKW_FROM, 0)
}

func (s *FromClauseContext) IDENT() antlr.TerminalNode {
	return s.GetToken(PRQLIDENT, 0)
}

func (s *FromClauseContext) KW_FINAL() antlr.TerminalNode {
	return s.GetToken(PRQLKW_FINAL, 0)
}

func (s *FromClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FromClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FromClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PRQLListener); ok {
		listenerT.EnterFromClause(s)
	}
}

func (s *FromClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PRQLListener); ok {
		listenerT.ExitFromClause(s)
	}
}

func (s *FromClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PRQLVisitor:
		return t.VisitFromClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PRQL) FromClause() (localctx IFromClauseContext) {
	localctx = NewFromClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, PRQLRULE_fromClause)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(81)
		p.Match(PRQLKW_FROM)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(82)
		p.Match(PRQLIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(84)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PRQLKW_FINAL {
		{
			p.SetState(83)
			p.Match(PRQLKW_FINAL)
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

// IFilterClauseContext is an interface to support dynamic dispatch.
type IFilterClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	KW_FILTER() antlr.TerminalNode
	FilterBody() IFilterBodyContext

	// IsFilterClauseContext differentiates from other interfaces.
	IsFilterClauseContext()
}

type FilterClauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFilterClauseContext() *FilterClauseContext {
	var p = new(FilterClauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PRQLRULE_filterClause
	return p
}

func InitEmptyFilterClauseContext(p *FilterClauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PRQLRULE_filterClause
}

func (*FilterClauseContext) IsFilterClauseContext() {}

func NewFilterClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FilterClauseContext {
	var p = new(FilterClauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PRQLRULE_filterClause

	return p
}

func (s *FilterClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *FilterClauseContext) KW_FILTER() antlr.TerminalNode {
	return s.GetToken(PRQLKW_FILTER, 0)
}

func (s *FilterClauseContext) FilterBody() IFilterBodyContext {
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

func (s *FilterClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FilterClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FilterClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PRQLListener); ok {
		listenerT.EnterFilterClause(s)
	}
}

func (s *FilterClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PRQLListener); ok {
		listenerT.ExitFilterClause(s)
	}
}

func (s *FilterClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PRQLVisitor:
		return t.VisitFilterClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PRQL) FilterClause() (localctx IFilterClauseContext) {
	localctx = NewFilterClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, PRQLRULE_filterClause)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(86)
		p.Match(PRQLKW_FILTER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(87)
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
	p.RuleIndex = PRQLRULE_filterBody
	return p
}

func InitEmptyFilterBodyContext(p *FilterBodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PRQLRULE_filterBody
}

func (*FilterBodyContext) IsFilterBodyContext() {}

func NewFilterBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FilterBodyContext {
	var p = new(FilterBodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PRQLRULE_filterBody

	return p
}

func (s *FilterBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *FilterBodyContext) AllFILTER_LINE() []antlr.TerminalNode {
	return s.GetTokens(PRQLFILTER_LINE)
}

func (s *FilterBodyContext) FILTER_LINE(i int) antlr.TerminalNode {
	return s.GetToken(PRQLFILTER_LINE, i)
}

func (s *FilterBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FilterBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FilterBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PRQLListener); ok {
		listenerT.EnterFilterBody(s)
	}
}

func (s *FilterBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PRQLListener); ok {
		listenerT.ExitFilterBody(s)
	}
}

func (s *FilterBodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PRQLVisitor:
		return t.VisitFilterBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PRQL) FilterBody() (localctx IFilterBodyContext) {
	localctx = NewFilterBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, PRQLRULE_filterBody)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(90)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == PRQLFILTER_LINE {
		{
			p.SetState(89)
			p.Match(PRQLFILTER_LINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(92)
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

// IDeriveClauseContext is an interface to support dynamic dispatch.
type IDeriveClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	KW_DERIVE() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	AssignmentList() IAssignmentListContext
	RBRACE() antlr.TerminalNode

	// IsDeriveClauseContext differentiates from other interfaces.
	IsDeriveClauseContext()
}

type DeriveClauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeriveClauseContext() *DeriveClauseContext {
	var p = new(DeriveClauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PRQLRULE_deriveClause
	return p
}

func InitEmptyDeriveClauseContext(p *DeriveClauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PRQLRULE_deriveClause
}

func (*DeriveClauseContext) IsDeriveClauseContext() {}

func NewDeriveClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeriveClauseContext {
	var p = new(DeriveClauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PRQLRULE_deriveClause

	return p
}

func (s *DeriveClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *DeriveClauseContext) KW_DERIVE() antlr.TerminalNode {
	return s.GetToken(PRQLKW_DERIVE, 0)
}

func (s *DeriveClauseContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(PRQLLBRACE, 0)
}

func (s *DeriveClauseContext) AssignmentList() IAssignmentListContext {
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

func (s *DeriveClauseContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(PRQLRBRACE, 0)
}

func (s *DeriveClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeriveClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeriveClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PRQLListener); ok {
		listenerT.EnterDeriveClause(s)
	}
}

func (s *DeriveClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PRQLListener); ok {
		listenerT.ExitDeriveClause(s)
	}
}

func (s *DeriveClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PRQLVisitor:
		return t.VisitDeriveClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PRQL) DeriveClause() (localctx IDeriveClauseContext) {
	localctx = NewDeriveClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, PRQLRULE_deriveClause)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(94)
		p.Match(PRQLKW_DERIVE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(95)
		p.Match(PRQLLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(96)
		p.AssignmentList()
	}
	{
		p.SetState(97)
		p.Match(PRQLRBRACE)
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

// ISelectClauseContext is an interface to support dynamic dispatch.
type ISelectClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	KW_SELECT() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	SelectionList() ISelectionListContext
	RBRACE() antlr.TerminalNode

	// IsSelectClauseContext differentiates from other interfaces.
	IsSelectClauseContext()
}

type SelectClauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectClauseContext() *SelectClauseContext {
	var p = new(SelectClauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PRQLRULE_selectClause
	return p
}

func InitEmptySelectClauseContext(p *SelectClauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PRQLRULE_selectClause
}

func (*SelectClauseContext) IsSelectClauseContext() {}

func NewSelectClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectClauseContext {
	var p = new(SelectClauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PRQLRULE_selectClause

	return p
}

func (s *SelectClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectClauseContext) KW_SELECT() antlr.TerminalNode {
	return s.GetToken(PRQLKW_SELECT, 0)
}

func (s *SelectClauseContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(PRQLLBRACE, 0)
}

func (s *SelectClauseContext) SelectionList() ISelectionListContext {
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

func (s *SelectClauseContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(PRQLRBRACE, 0)
}

func (s *SelectClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PRQLListener); ok {
		listenerT.EnterSelectClause(s)
	}
}

func (s *SelectClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PRQLListener); ok {
		listenerT.ExitSelectClause(s)
	}
}

func (s *SelectClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PRQLVisitor:
		return t.VisitSelectClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PRQL) SelectClause() (localctx ISelectClauseContext) {
	localctx = NewSelectClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, PRQLRULE_selectClause)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(99)
		p.Match(PRQLKW_SELECT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(100)
		p.Match(PRQLLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(101)
		p.SelectionList()
	}
	{
		p.SetState(102)
		p.Match(PRQLRBRACE)
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

// IGroupClauseContext is an interface to support dynamic dispatch.
type IGroupClauseContext interface {
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

	// IsGroupClauseContext differentiates from other interfaces.
	IsGroupClauseContext()
}

type GroupClauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGroupClauseContext() *GroupClauseContext {
	var p = new(GroupClauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PRQLRULE_groupClause
	return p
}

func InitEmptyGroupClauseContext(p *GroupClauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PRQLRULE_groupClause
}

func (*GroupClauseContext) IsGroupClauseContext() {}

func NewGroupClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GroupClauseContext {
	var p = new(GroupClauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PRQLRULE_groupClause

	return p
}

func (s *GroupClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *GroupClauseContext) KW_GROUP() antlr.TerminalNode {
	return s.GetToken(PRQLKW_GROUP, 0)
}

func (s *GroupClauseContext) AllLBRACE() []antlr.TerminalNode {
	return s.GetTokens(PRQLLBRACE)
}

func (s *GroupClauseContext) LBRACE(i int) antlr.TerminalNode {
	return s.GetToken(PRQLLBRACE, i)
}

func (s *GroupClauseContext) KeyList() IKeyListContext {
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

func (s *GroupClauseContext) AllRBRACE() []antlr.TerminalNode {
	return s.GetTokens(PRQLRBRACE)
}

func (s *GroupClauseContext) RBRACE(i int) antlr.TerminalNode {
	return s.GetToken(PRQLRBRACE, i)
}

func (s *GroupClauseContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(PRQLLPAREN, 0)
}

func (s *GroupClauseContext) KW_AGGREGATE() antlr.TerminalNode {
	return s.GetToken(PRQLKW_AGGREGATE, 0)
}

func (s *GroupClauseContext) AssignmentList() IAssignmentListContext {
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

func (s *GroupClauseContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(PRQLRPAREN, 0)
}

func (s *GroupClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GroupClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GroupClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PRQLListener); ok {
		listenerT.EnterGroupClause(s)
	}
}

func (s *GroupClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PRQLListener); ok {
		listenerT.ExitGroupClause(s)
	}
}

func (s *GroupClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PRQLVisitor:
		return t.VisitGroupClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PRQL) GroupClause() (localctx IGroupClauseContext) {
	localctx = NewGroupClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, PRQLRULE_groupClause)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(104)
		p.Match(PRQLKW_GROUP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(105)
		p.Match(PRQLLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(106)
		p.KeyList()
	}
	{
		p.SetState(107)
		p.Match(PRQLRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(108)
		p.Match(PRQLLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(109)
		p.Match(PRQLKW_AGGREGATE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(110)
		p.Match(PRQLLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(111)
		p.AssignmentList()
	}
	{
		p.SetState(112)
		p.Match(PRQLRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(113)
		p.Match(PRQLRPAREN)
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

// IJoinClauseContext is an interface to support dynamic dispatch.
type IJoinClauseContext interface {
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

	// IsJoinClauseContext differentiates from other interfaces.
	IsJoinClauseContext()
}

type JoinClauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJoinClauseContext() *JoinClauseContext {
	var p = new(JoinClauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PRQLRULE_joinClause
	return p
}

func InitEmptyJoinClauseContext(p *JoinClauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PRQLRULE_joinClause
}

func (*JoinClauseContext) IsJoinClauseContext() {}

func NewJoinClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JoinClauseContext {
	var p = new(JoinClauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PRQLRULE_joinClause

	return p
}

func (s *JoinClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *JoinClauseContext) KW_JOIN() antlr.TerminalNode {
	return s.GetToken(PRQLKW_JOIN, 0)
}

func (s *JoinClauseContext) IDENT() antlr.TerminalNode {
	return s.GetToken(PRQLIDENT, 0)
}

func (s *JoinClauseContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(PRQLLPAREN, 0)
}

func (s *JoinClauseContext) JoinCond() IJoinCondContext {
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

func (s *JoinClauseContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(PRQLRPAREN, 0)
}

func (s *JoinClauseContext) KW_SIDE() antlr.TerminalNode {
	return s.GetToken(PRQLKW_SIDE, 0)
}

func (s *JoinClauseContext) COLON() antlr.TerminalNode {
	return s.GetToken(PRQLCOLON, 0)
}

func (s *JoinClauseContext) JoinSide() IJoinSideContext {
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

func (s *JoinClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JoinClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JoinClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PRQLListener); ok {
		listenerT.EnterJoinClause(s)
	}
}

func (s *JoinClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PRQLListener); ok {
		listenerT.ExitJoinClause(s)
	}
}

func (s *JoinClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PRQLVisitor:
		return t.VisitJoinClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PRQL) JoinClause() (localctx IJoinClauseContext) {
	localctx = NewJoinClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, PRQLRULE_joinClause)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(115)
		p.Match(PRQLKW_JOIN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(119)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PRQLKW_SIDE {
		{
			p.SetState(116)
			p.Match(PRQLKW_SIDE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(117)
			p.Match(PRQLCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(118)
			p.JoinSide()
		}

	}
	{
		p.SetState(121)
		p.Match(PRQLIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(122)
		p.Match(PRQLLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(123)
		p.JoinCond()
	}
	{
		p.SetState(124)
		p.Match(PRQLRPAREN)
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
	p.RuleIndex = PRQLRULE_joinSide
	return p
}

func InitEmptyJoinSideContext(p *JoinSideContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PRQLRULE_joinSide
}

func (*JoinSideContext) IsJoinSideContext() {}

func NewJoinSideContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JoinSideContext {
	var p = new(JoinSideContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PRQLRULE_joinSide

	return p
}

func (s *JoinSideContext) GetParser() antlr.Parser { return s.parser }

func (s *JoinSideContext) KW_LEFT() antlr.TerminalNode {
	return s.GetToken(PRQLKW_LEFT, 0)
}

func (s *JoinSideContext) KW_RIGHT() antlr.TerminalNode {
	return s.GetToken(PRQLKW_RIGHT, 0)
}

func (s *JoinSideContext) KW_INNER() antlr.TerminalNode {
	return s.GetToken(PRQLKW_INNER, 0)
}

func (s *JoinSideContext) KW_FULL() antlr.TerminalNode {
	return s.GetToken(PRQLKW_FULL, 0)
}

func (s *JoinSideContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JoinSideContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JoinSideContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PRQLListener); ok {
		listenerT.EnterJoinSide(s)
	}
}

func (s *JoinSideContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PRQLListener); ok {
		listenerT.ExitJoinSide(s)
	}
}

func (s *JoinSideContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PRQLVisitor:
		return t.VisitJoinSide(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PRQL) JoinSide() (localctx IJoinSideContext) {
	localctx = NewJoinSideContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, PRQLRULE_joinSide)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(126)
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
	p.RuleIndex = PRQLRULE_joinCond
	return p
}

func InitEmptyJoinCondContext(p *JoinCondContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PRQLRULE_joinCond
}

func (*JoinCondContext) IsJoinCondContext() {}

func NewJoinCondContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JoinCondContext {
	var p = new(JoinCondContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PRQLRULE_joinCond

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
	return s.GetTokens(PRQLEQ)
}

func (s *SelfJoinCondContext) EQ(i int) antlr.TerminalNode {
	return s.GetToken(PRQLEQ, i)
}

func (s *SelfJoinCondContext) IDENT() antlr.TerminalNode {
	return s.GetToken(PRQLIDENT, 0)
}

func (s *SelfJoinCondContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PRQLListener); ok {
		listenerT.EnterSelfJoinCond(s)
	}
}

func (s *SelfJoinCondContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PRQLListener); ok {
		listenerT.ExitSelfJoinCond(s)
	}
}

func (s *SelfJoinCondContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PRQLVisitor:
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
	if listenerT, ok := listener.(PRQLListener); ok {
		listenerT.EnterExplicitJoinCond(s)
	}
}

func (s *ExplicitJoinCondContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PRQLListener); ok {
		listenerT.ExitExplicitJoinCond(s)
	}
}

func (s *ExplicitJoinCondContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PRQLVisitor:
		return t.VisitExplicitJoinCond(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PRQL) JoinCond() (localctx IJoinCondContext) {
	localctx = NewJoinCondContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, PRQLRULE_joinCond)
	p.SetState(132)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		localctx = NewSelfJoinCondContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(128)
			p.Match(PRQLEQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(129)
			p.Match(PRQLEQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(130)
			p.Match(PRQLIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewExplicitJoinCondContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(131)
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
	p.RuleIndex = PRQLRULE_joinCondExpr
	return p
}

func InitEmptyJoinCondExprContext(p *JoinCondExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PRQLRULE_joinCondExpr
}

func (*JoinCondExprContext) IsJoinCondExprContext() {}

func NewJoinCondExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JoinCondExprContext {
	var p = new(JoinCondExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PRQLRULE_joinCondExpr

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
	if listenerT, ok := listener.(PRQLListener); ok {
		listenerT.EnterJoinCondExpr(s)
	}
}

func (s *JoinCondExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PRQLListener); ok {
		listenerT.ExitJoinCondExpr(s)
	}
}

func (s *JoinCondExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PRQLVisitor:
		return t.VisitJoinCondExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PRQL) JoinCondExpr() (localctx IJoinCondExprContext) {
	localctx = NewJoinCondExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, PRQLRULE_joinCondExpr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(135)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&140669054091264) != 0) {
		{
			p.SetState(134)
			p.JoinCondToken()
		}

		p.SetState(137)
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
	p.RuleIndex = PRQLRULE_joinCondToken
	return p
}

func InitEmptyJoinCondTokenContext(p *JoinCondTokenContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PRQLRULE_joinCondToken
}

func (*JoinCondTokenContext) IsJoinCondTokenContext() {}

func NewJoinCondTokenContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JoinCondTokenContext {
	var p = new(JoinCondTokenContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PRQLRULE_joinCondToken

	return p
}

func (s *JoinCondTokenContext) GetParser() antlr.Parser { return s.parser }

func (s *JoinCondTokenContext) IDENT() antlr.TerminalNode {
	return s.GetToken(PRQLIDENT, 0)
}

func (s *JoinCondTokenContext) DOT() antlr.TerminalNode {
	return s.GetToken(PRQLDOT, 0)
}

func (s *JoinCondTokenContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(PRQLINTEGER, 0)
}

func (s *JoinCondTokenContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(PRQLFLOAT, 0)
}

func (s *JoinCondTokenContext) STRING() antlr.TerminalNode {
	return s.GetToken(PRQLSTRING, 0)
}

func (s *JoinCondTokenContext) EQ() antlr.TerminalNode {
	return s.GetToken(PRQLEQ, 0)
}

func (s *JoinCondTokenContext) NEQ() antlr.TerminalNode {
	return s.GetToken(PRQLNEQ, 0)
}

func (s *JoinCondTokenContext) LT() antlr.TerminalNode {
	return s.GetToken(PRQLLT, 0)
}

func (s *JoinCondTokenContext) GT() antlr.TerminalNode {
	return s.GetToken(PRQLGT, 0)
}

func (s *JoinCondTokenContext) LTE() antlr.TerminalNode {
	return s.GetToken(PRQLLTE, 0)
}

func (s *JoinCondTokenContext) GTE() antlr.TerminalNode {
	return s.GetToken(PRQLGTE, 0)
}

func (s *JoinCondTokenContext) CAST_OP() antlr.TerminalNode {
	return s.GetToken(PRQLCAST_OP, 0)
}

func (s *JoinCondTokenContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(PRQLLPAREN, 0)
}

func (s *JoinCondTokenContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(PRQLRPAREN, 0)
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
	if listenerT, ok := listener.(PRQLListener); ok {
		listenerT.EnterJoinCondToken(s)
	}
}

func (s *JoinCondTokenContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PRQLListener); ok {
		listenerT.ExitJoinCondToken(s)
	}
}

func (s *JoinCondTokenContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PRQLVisitor:
		return t.VisitJoinCondToken(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PRQL) JoinCondToken() (localctx IJoinCondTokenContext) {
	localctx = NewJoinCondTokenContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, PRQLRULE_joinCondToken)
	var _la int

	p.SetState(159)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PRQLIDENT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(139)
			p.Match(PRQLIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PRQLDOT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(140)
			p.Match(PRQLDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PRQLINTEGER:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(141)
			p.Match(PRQLINTEGER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PRQLFLOAT:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(142)
			p.Match(PRQLFLOAT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PRQLSTRING:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(143)
			p.Match(PRQLSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PRQLEQ:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(144)
			p.Match(PRQLEQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PRQLNEQ:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(145)
			p.Match(PRQLNEQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PRQLLT:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(146)
			p.Match(PRQLLT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PRQLGT:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(147)
			p.Match(PRQLGT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PRQLLTE:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(148)
			p.Match(PRQLLTE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PRQLGTE:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(149)
			p.Match(PRQLGTE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PRQLCAST_OP:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(150)
			p.Match(PRQLCAST_OP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PRQLLPAREN:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(151)
			p.Match(PRQLLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(155)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&140669121200128) != 0 {
			{
				p.SetState(152)
				p.JoinCondInner()
			}

			p.SetState(157)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(158)
			p.Match(PRQLRPAREN)
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
	p.RuleIndex = PRQLRULE_joinCondInner
	return p
}

func InitEmptyJoinCondInnerContext(p *JoinCondInnerContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PRQLRULE_joinCondInner
}

func (*JoinCondInnerContext) IsJoinCondInnerContext() {}

func NewJoinCondInnerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JoinCondInnerContext {
	var p = new(JoinCondInnerContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PRQLRULE_joinCondInner

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
	return s.GetToken(PRQLCOMMA, 0)
}

func (s *JoinCondInnerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JoinCondInnerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JoinCondInnerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PRQLListener); ok {
		listenerT.EnterJoinCondInner(s)
	}
}

func (s *JoinCondInnerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PRQLListener); ok {
		listenerT.ExitJoinCondInner(s)
	}
}

func (s *JoinCondInnerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PRQLVisitor:
		return t.VisitJoinCondInner(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PRQL) JoinCondInner() (localctx IJoinCondInnerContext) {
	localctx = NewJoinCondInnerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, PRQLRULE_joinCondInner)
	p.SetState(163)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PRQLLPAREN, PRQLDOT, PRQLCAST_OP, PRQLEQ, PRQLNEQ, PRQLLTE, PRQLGTE, PRQLLT, PRQLGT, PRQLFLOAT, PRQLINTEGER, PRQLSTRING, PRQLIDENT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(161)
			p.JoinCondToken()
		}

	case PRQLCOMMA:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(162)
			p.Match(PRQLCOMMA)
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

// IArrayJoinClauseContext is an interface to support dynamic dispatch.
type IArrayJoinClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	KW_ARRAY_JOIN() antlr.TerminalNode
	AllIDENT() []antlr.TerminalNode
	IDENT(i int) antlr.TerminalNode
	KW_AS() antlr.TerminalNode

	// IsArrayJoinClauseContext differentiates from other interfaces.
	IsArrayJoinClauseContext()
}

type ArrayJoinClauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayJoinClauseContext() *ArrayJoinClauseContext {
	var p = new(ArrayJoinClauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PRQLRULE_arrayJoinClause
	return p
}

func InitEmptyArrayJoinClauseContext(p *ArrayJoinClauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PRQLRULE_arrayJoinClause
}

func (*ArrayJoinClauseContext) IsArrayJoinClauseContext() {}

func NewArrayJoinClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayJoinClauseContext {
	var p = new(ArrayJoinClauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PRQLRULE_arrayJoinClause

	return p
}

func (s *ArrayJoinClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayJoinClauseContext) KW_ARRAY_JOIN() antlr.TerminalNode {
	return s.GetToken(PRQLKW_ARRAY_JOIN, 0)
}

func (s *ArrayJoinClauseContext) AllIDENT() []antlr.TerminalNode {
	return s.GetTokens(PRQLIDENT)
}

func (s *ArrayJoinClauseContext) IDENT(i int) antlr.TerminalNode {
	return s.GetToken(PRQLIDENT, i)
}

func (s *ArrayJoinClauseContext) KW_AS() antlr.TerminalNode {
	return s.GetToken(PRQLKW_AS, 0)
}

func (s *ArrayJoinClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayJoinClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayJoinClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PRQLListener); ok {
		listenerT.EnterArrayJoinClause(s)
	}
}

func (s *ArrayJoinClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PRQLListener); ok {
		listenerT.ExitArrayJoinClause(s)
	}
}

func (s *ArrayJoinClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PRQLVisitor:
		return t.VisitArrayJoinClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PRQL) ArrayJoinClause() (localctx IArrayJoinClauseContext) {
	localctx = NewArrayJoinClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, PRQLRULE_arrayJoinClause)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(165)
		p.Match(PRQLKW_ARRAY_JOIN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(166)
		p.Match(PRQLIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(169)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PRQLKW_AS {
		{
			p.SetState(167)
			p.Match(PRQLKW_AS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(168)
			p.Match(PRQLIDENT)
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

// ISortClauseContext is an interface to support dynamic dispatch.
type ISortClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	KW_SORT() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	SortList() ISortListContext
	RBRACE() antlr.TerminalNode

	// IsSortClauseContext differentiates from other interfaces.
	IsSortClauseContext()
}

type SortClauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySortClauseContext() *SortClauseContext {
	var p = new(SortClauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PRQLRULE_sortClause
	return p
}

func InitEmptySortClauseContext(p *SortClauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PRQLRULE_sortClause
}

func (*SortClauseContext) IsSortClauseContext() {}

func NewSortClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SortClauseContext {
	var p = new(SortClauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PRQLRULE_sortClause

	return p
}

func (s *SortClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *SortClauseContext) KW_SORT() antlr.TerminalNode {
	return s.GetToken(PRQLKW_SORT, 0)
}

func (s *SortClauseContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(PRQLLBRACE, 0)
}

func (s *SortClauseContext) SortList() ISortListContext {
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

func (s *SortClauseContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(PRQLRBRACE, 0)
}

func (s *SortClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SortClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SortClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PRQLListener); ok {
		listenerT.EnterSortClause(s)
	}
}

func (s *SortClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PRQLListener); ok {
		listenerT.ExitSortClause(s)
	}
}

func (s *SortClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PRQLVisitor:
		return t.VisitSortClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PRQL) SortClause() (localctx ISortClauseContext) {
	localctx = NewSortClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, PRQLRULE_sortClause)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(171)
		p.Match(PRQLKW_SORT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(172)
		p.Match(PRQLLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(173)
		p.SortList()
	}
	{
		p.SetState(174)
		p.Match(PRQLRBRACE)
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

// ITakeClauseContext is an interface to support dynamic dispatch.
type ITakeClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	KW_TAKE() antlr.TerminalNode
	AllINTEGER() []antlr.TerminalNode
	INTEGER(i int) antlr.TerminalNode
	RANGE() antlr.TerminalNode

	// IsTakeClauseContext differentiates from other interfaces.
	IsTakeClauseContext()
}

type TakeClauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTakeClauseContext() *TakeClauseContext {
	var p = new(TakeClauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PRQLRULE_takeClause
	return p
}

func InitEmptyTakeClauseContext(p *TakeClauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PRQLRULE_takeClause
}

func (*TakeClauseContext) IsTakeClauseContext() {}

func NewTakeClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TakeClauseContext {
	var p = new(TakeClauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PRQLRULE_takeClause

	return p
}

func (s *TakeClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *TakeClauseContext) KW_TAKE() antlr.TerminalNode {
	return s.GetToken(PRQLKW_TAKE, 0)
}

func (s *TakeClauseContext) AllINTEGER() []antlr.TerminalNode {
	return s.GetTokens(PRQLINTEGER)
}

func (s *TakeClauseContext) INTEGER(i int) antlr.TerminalNode {
	return s.GetToken(PRQLINTEGER, i)
}

func (s *TakeClauseContext) RANGE() antlr.TerminalNode {
	return s.GetToken(PRQLRANGE, 0)
}

func (s *TakeClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TakeClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TakeClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PRQLListener); ok {
		listenerT.EnterTakeClause(s)
	}
}

func (s *TakeClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PRQLListener); ok {
		listenerT.ExitTakeClause(s)
	}
}

func (s *TakeClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PRQLVisitor:
		return t.VisitTakeClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PRQL) TakeClause() (localctx ITakeClauseContext) {
	localctx = NewTakeClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, PRQLRULE_takeClause)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(176)
		p.Match(PRQLKW_TAKE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(177)
		p.Match(PRQLINTEGER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(180)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PRQLRANGE {
		{
			p.SetState(178)
			p.Match(PRQLRANGE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(179)
			p.Match(PRQLINTEGER)
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

// ISkipClauseContext is an interface to support dynamic dispatch.
type ISkipClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	KW_SKIP() antlr.TerminalNode
	INTEGER() antlr.TerminalNode

	// IsSkipClauseContext differentiates from other interfaces.
	IsSkipClauseContext()
}

type SkipClauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySkipClauseContext() *SkipClauseContext {
	var p = new(SkipClauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PRQLRULE_skipClause
	return p
}

func InitEmptySkipClauseContext(p *SkipClauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PRQLRULE_skipClause
}

func (*SkipClauseContext) IsSkipClauseContext() {}

func NewSkipClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SkipClauseContext {
	var p = new(SkipClauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PRQLRULE_skipClause

	return p
}

func (s *SkipClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *SkipClauseContext) KW_SKIP() antlr.TerminalNode {
	return s.GetToken(PRQLKW_SKIP, 0)
}

func (s *SkipClauseContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(PRQLINTEGER, 0)
}

func (s *SkipClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SkipClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SkipClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PRQLListener); ok {
		listenerT.EnterSkipClause(s)
	}
}

func (s *SkipClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PRQLListener); ok {
		listenerT.ExitSkipClause(s)
	}
}

func (s *SkipClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PRQLVisitor:
		return t.VisitSkipClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PRQL) SkipClause() (localctx ISkipClauseContext) {
	localctx = NewSkipClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, PRQLRULE_skipClause)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(182)
		p.Match(PRQLKW_SKIP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(183)
		p.Match(PRQLINTEGER)
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

// IWindowClauseContext is an interface to support dynamic dispatch.
type IWindowClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	KW_WINDOW() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	AssignmentList() IAssignmentListContext
	RBRACE() antlr.TerminalNode

	// IsWindowClauseContext differentiates from other interfaces.
	IsWindowClauseContext()
}

type WindowClauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWindowClauseContext() *WindowClauseContext {
	var p = new(WindowClauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PRQLRULE_windowClause
	return p
}

func InitEmptyWindowClauseContext(p *WindowClauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PRQLRULE_windowClause
}

func (*WindowClauseContext) IsWindowClauseContext() {}

func NewWindowClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WindowClauseContext {
	var p = new(WindowClauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PRQLRULE_windowClause

	return p
}

func (s *WindowClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *WindowClauseContext) KW_WINDOW() antlr.TerminalNode {
	return s.GetToken(PRQLKW_WINDOW, 0)
}

func (s *WindowClauseContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(PRQLLBRACE, 0)
}

func (s *WindowClauseContext) AssignmentList() IAssignmentListContext {
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

func (s *WindowClauseContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(PRQLRBRACE, 0)
}

func (s *WindowClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WindowClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WindowClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PRQLListener); ok {
		listenerT.EnterWindowClause(s)
	}
}

func (s *WindowClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PRQLListener); ok {
		listenerT.ExitWindowClause(s)
	}
}

func (s *WindowClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PRQLVisitor:
		return t.VisitWindowClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PRQL) WindowClause() (localctx IWindowClauseContext) {
	localctx = NewWindowClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, PRQLRULE_windowClause)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(185)
		p.Match(PRQLKW_WINDOW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(186)
		p.Match(PRQLLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(187)
		p.AssignmentList()
	}
	{
		p.SetState(188)
		p.Match(PRQLRBRACE)
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
	p.RuleIndex = PRQLRULE_assignmentList
	return p
}

func InitEmptyAssignmentListContext(p *AssignmentListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PRQLRULE_assignmentList
}

func (*AssignmentListContext) IsAssignmentListContext() {}

func NewAssignmentListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentListContext {
	var p = new(AssignmentListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PRQLRULE_assignmentList

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
	return s.GetTokens(PRQLCOMMA)
}

func (s *AssignmentListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(PRQLCOMMA, i)
}

func (s *AssignmentListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PRQLListener); ok {
		listenerT.EnterAssignmentList(s)
	}
}

func (s *AssignmentListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PRQLListener); ok {
		listenerT.ExitAssignmentList(s)
	}
}

func (s *AssignmentListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PRQLVisitor:
		return t.VisitAssignmentList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PRQL) AssignmentList() (localctx IAssignmentListContext) {
	localctx = NewAssignmentListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, PRQLRULE_assignmentList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(190)
		p.Assignment()
	}
	p.SetState(195)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == PRQLCOMMA {
		{
			p.SetState(191)
			p.Match(PRQLCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(192)
			p.Assignment()
		}

		p.SetState(197)
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
	p.RuleIndex = PRQLRULE_assignment
	return p
}

func InitEmptyAssignmentContext(p *AssignmentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PRQLRULE_assignment
}

func (*AssignmentContext) IsAssignmentContext() {}

func NewAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentContext {
	var p = new(AssignmentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PRQLRULE_assignment

	return p
}

func (s *AssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentContext) IDENT() antlr.TerminalNode {
	return s.GetToken(PRQLIDENT, 0)
}

func (s *AssignmentContext) EQ() antlr.TerminalNode {
	return s.GetToken(PRQLEQ, 0)
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
	if listenerT, ok := listener.(PRQLListener); ok {
		listenerT.EnterAssignment(s)
	}
}

func (s *AssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PRQLListener); ok {
		listenerT.ExitAssignment(s)
	}
}

func (s *AssignmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PRQLVisitor:
		return t.VisitAssignment(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PRQL) Assignment() (localctx IAssignmentContext) {
	localctx = NewAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, PRQLRULE_assignment)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(198)
		p.Match(PRQLIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(199)
		p.Match(PRQLEQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(200)
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
	p.RuleIndex = PRQLRULE_selectionList
	return p
}

func InitEmptySelectionListContext(p *SelectionListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PRQLRULE_selectionList
}

func (*SelectionListContext) IsSelectionListContext() {}

func NewSelectionListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectionListContext {
	var p = new(SelectionListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PRQLRULE_selectionList

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
	return s.GetTokens(PRQLCOMMA)
}

func (s *SelectionListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(PRQLCOMMA, i)
}

func (s *SelectionListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectionListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectionListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PRQLListener); ok {
		listenerT.EnterSelectionList(s)
	}
}

func (s *SelectionListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PRQLListener); ok {
		listenerT.ExitSelectionList(s)
	}
}

func (s *SelectionListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PRQLVisitor:
		return t.VisitSelectionList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PRQL) SelectionList() (localctx ISelectionListContext) {
	localctx = NewSelectionListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, PRQLRULE_selectionList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(202)
		p.SelectionItem()
	}
	p.SetState(207)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == PRQLCOMMA {
		{
			p.SetState(203)
			p.Match(PRQLCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(204)
			p.SelectionItem()
		}

		p.SetState(209)
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
	p.RuleIndex = PRQLRULE_selectionItem
	return p
}

func InitEmptySelectionItemContext(p *SelectionItemContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PRQLRULE_selectionItem
}

func (*SelectionItemContext) IsSelectionItemContext() {}

func NewSelectionItemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectionItemContext {
	var p = new(SelectionItemContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PRQLRULE_selectionItem

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
	return s.GetToken(PRQLIDENT, 0)
}

func (s *AliasedSelectionContext) EQ() antlr.TerminalNode {
	return s.GetToken(PRQLEQ, 0)
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
	if listenerT, ok := listener.(PRQLListener); ok {
		listenerT.EnterAliasedSelection(s)
	}
}

func (s *AliasedSelectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PRQLListener); ok {
		listenerT.ExitAliasedSelection(s)
	}
}

func (s *AliasedSelectionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PRQLVisitor:
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
	if listenerT, ok := listener.(PRQLListener); ok {
		listenerT.EnterBareSelection(s)
	}
}

func (s *BareSelectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PRQLListener); ok {
		listenerT.ExitBareSelection(s)
	}
}

func (s *BareSelectionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PRQLVisitor:
		return t.VisitBareSelection(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PRQL) SelectionItem() (localctx ISelectionItemContext) {
	localctx = NewSelectionItemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, PRQLRULE_selectionItem)
	p.SetState(214)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		localctx = NewAliasedSelectionContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(210)
			p.Match(PRQLIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(211)
			p.Match(PRQLEQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(212)
			p.OpaqueExpr()
		}

	case 2:
		localctx = NewBareSelectionContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(213)
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
	p.RuleIndex = PRQLRULE_keyList
	return p
}

func InitEmptyKeyListContext(p *KeyListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PRQLRULE_keyList
}

func (*KeyListContext) IsKeyListContext() {}

func NewKeyListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeyListContext {
	var p = new(KeyListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PRQLRULE_keyList

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
	return s.GetTokens(PRQLCOMMA)
}

func (s *KeyListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(PRQLCOMMA, i)
}

func (s *KeyListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeyListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeyListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PRQLListener); ok {
		listenerT.EnterKeyList(s)
	}
}

func (s *KeyListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PRQLListener); ok {
		listenerT.ExitKeyList(s)
	}
}

func (s *KeyListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PRQLVisitor:
		return t.VisitKeyList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PRQL) KeyList() (localctx IKeyListContext) {
	localctx = NewKeyListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, PRQLRULE_keyList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(216)
		p.KeyItem()
	}
	p.SetState(221)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == PRQLCOMMA {
		{
			p.SetState(217)
			p.Match(PRQLCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(218)
			p.KeyItem()
		}

		p.SetState(223)
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
	p.RuleIndex = PRQLRULE_keyItem
	return p
}

func InitEmptyKeyItemContext(p *KeyItemContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PRQLRULE_keyItem
}

func (*KeyItemContext) IsKeyItemContext() {}

func NewKeyItemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeyItemContext {
	var p = new(KeyItemContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PRQLRULE_keyItem

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
	return s.GetToken(PRQLIDENT, 0)
}

func (s *ComputedKeyContext) EQ() antlr.TerminalNode {
	return s.GetToken(PRQLEQ, 0)
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
	if listenerT, ok := listener.(PRQLListener); ok {
		listenerT.EnterComputedKey(s)
	}
}

func (s *ComputedKeyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PRQLListener); ok {
		listenerT.ExitComputedKey(s)
	}
}

func (s *ComputedKeyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PRQLVisitor:
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
	return s.GetToken(PRQLIDENT, 0)
}

func (s *ColumnKeyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PRQLListener); ok {
		listenerT.EnterColumnKey(s)
	}
}

func (s *ColumnKeyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PRQLListener); ok {
		listenerT.ExitColumnKey(s)
	}
}

func (s *ColumnKeyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PRQLVisitor:
		return t.VisitColumnKey(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PRQL) KeyItem() (localctx IKeyItemContext) {
	localctx = NewKeyItemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, PRQLRULE_keyItem)
	p.SetState(228)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		localctx = NewComputedKeyContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(224)
			p.Match(PRQLIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(225)
			p.Match(PRQLEQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(226)
			p.OpaqueExpr()
		}

	case 2:
		localctx = NewColumnKeyContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(227)
			p.Match(PRQLIDENT)
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
	p.RuleIndex = PRQLRULE_sortList
	return p
}

func InitEmptySortListContext(p *SortListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PRQLRULE_sortList
}

func (*SortListContext) IsSortListContext() {}

func NewSortListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SortListContext {
	var p = new(SortListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PRQLRULE_sortList

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
	return s.GetTokens(PRQLCOMMA)
}

func (s *SortListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(PRQLCOMMA, i)
}

func (s *SortListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SortListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SortListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PRQLListener); ok {
		listenerT.EnterSortList(s)
	}
}

func (s *SortListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PRQLListener); ok {
		listenerT.ExitSortList(s)
	}
}

func (s *SortListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PRQLVisitor:
		return t.VisitSortList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PRQL) SortList() (localctx ISortListContext) {
	localctx = NewSortListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, PRQLRULE_sortList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(230)
		p.SortItem()
	}
	p.SetState(235)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == PRQLCOMMA {
		{
			p.SetState(231)
			p.Match(PRQLCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(232)
			p.SortItem()
		}

		p.SetState(237)
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
	p.RuleIndex = PRQLRULE_sortItem
	return p
}

func InitEmptySortItemContext(p *SortItemContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PRQLRULE_sortItem
}

func (*SortItemContext) IsSortItemContext() {}

func NewSortItemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SortItemContext {
	var p = new(SortItemContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PRQLRULE_sortItem

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
	return s.GetToken(PRQLMINUS, 0)
}

func (s *DescSortContext) IDENT() antlr.TerminalNode {
	return s.GetToken(PRQLIDENT, 0)
}

func (s *DescSortContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PRQLListener); ok {
		listenerT.EnterDescSort(s)
	}
}

func (s *DescSortContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PRQLListener); ok {
		listenerT.ExitDescSort(s)
	}
}

func (s *DescSortContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PRQLVisitor:
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
	return s.GetToken(PRQLIDENT, 0)
}

func (s *AscSortContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PRQLListener); ok {
		listenerT.EnterAscSort(s)
	}
}

func (s *AscSortContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PRQLListener); ok {
		listenerT.ExitAscSort(s)
	}
}

func (s *AscSortContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PRQLVisitor:
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
	return s.GetToken(PRQLPLUS, 0)
}

func (s *AscSortExplicitContext) IDENT() antlr.TerminalNode {
	return s.GetToken(PRQLIDENT, 0)
}

func (s *AscSortExplicitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PRQLListener); ok {
		listenerT.EnterAscSortExplicit(s)
	}
}

func (s *AscSortExplicitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PRQLListener); ok {
		listenerT.ExitAscSortExplicit(s)
	}
}

func (s *AscSortExplicitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PRQLVisitor:
		return t.VisitAscSortExplicit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PRQL) SortItem() (localctx ISortItemContext) {
	localctx = NewSortItemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, PRQLRULE_sortItem)
	p.SetState(243)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PRQLMINUS:
		localctx = NewDescSortContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(238)
			p.Match(PRQLMINUS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(239)
			p.Match(PRQLIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PRQLPLUS:
		localctx = NewAscSortExplicitContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(240)
			p.Match(PRQLPLUS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(241)
			p.Match(PRQLIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PRQLIDENT:
		localctx = NewAscSortContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(242)
			p.Match(PRQLIDENT)
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
	p.RuleIndex = PRQLRULE_opaqueExpr
	return p
}

func InitEmptyOpaqueExprContext(p *OpaqueExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PRQLRULE_opaqueExpr
}

func (*OpaqueExprContext) IsOpaqueExprContext() {}

func NewOpaqueExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OpaqueExprContext {
	var p = new(OpaqueExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PRQLRULE_opaqueExpr

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
	if listenerT, ok := listener.(PRQLListener); ok {
		listenerT.EnterOpaqueExpr(s)
	}
}

func (s *OpaqueExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PRQLListener); ok {
		listenerT.ExitOpaqueExpr(s)
	}
}

func (s *OpaqueExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PRQLVisitor:
		return t.VisitOpaqueExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PRQL) OpaqueExpr() (localctx IOpaqueExprContext) {
	localctx = NewOpaqueExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, PRQLRULE_opaqueExpr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(246)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&140736703987712) != 0) {
		{
			p.SetState(245)
			p.OpaqueToken()
		}

		p.SetState(248)
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
	p.RuleIndex = PRQLRULE_opaqueToken
	return p
}

func InitEmptyOpaqueTokenContext(p *OpaqueTokenContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PRQLRULE_opaqueToken
}

func (*OpaqueTokenContext) IsOpaqueTokenContext() {}

func NewOpaqueTokenContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OpaqueTokenContext {
	var p = new(OpaqueTokenContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PRQLRULE_opaqueToken

	return p
}

func (s *OpaqueTokenContext) GetParser() antlr.Parser { return s.parser }

func (s *OpaqueTokenContext) IDENT() antlr.TerminalNode {
	return s.GetToken(PRQLIDENT, 0)
}

func (s *OpaqueTokenContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(PRQLINTEGER, 0)
}

func (s *OpaqueTokenContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(PRQLFLOAT, 0)
}

func (s *OpaqueTokenContext) STRING() antlr.TerminalNode {
	return s.GetToken(PRQLSTRING, 0)
}

func (s *OpaqueTokenContext) STAR() antlr.TerminalNode {
	return s.GetToken(PRQLSTAR, 0)
}

func (s *OpaqueTokenContext) PLUS() antlr.TerminalNode {
	return s.GetToken(PRQLPLUS, 0)
}

func (s *OpaqueTokenContext) MINUS() antlr.TerminalNode {
	return s.GetToken(PRQLMINUS, 0)
}

func (s *OpaqueTokenContext) SLASH() antlr.TerminalNode {
	return s.GetToken(PRQLSLASH, 0)
}

func (s *OpaqueTokenContext) PERCENT() antlr.TerminalNode {
	return s.GetToken(PRQLPERCENT, 0)
}

func (s *OpaqueTokenContext) PIPE() antlr.TerminalNode {
	return s.GetToken(PRQLPIPE, 0)
}

func (s *OpaqueTokenContext) EQ() antlr.TerminalNode {
	return s.GetToken(PRQLEQ, 0)
}

func (s *OpaqueTokenContext) NEQ() antlr.TerminalNode {
	return s.GetToken(PRQLNEQ, 0)
}

func (s *OpaqueTokenContext) LT() antlr.TerminalNode {
	return s.GetToken(PRQLLT, 0)
}

func (s *OpaqueTokenContext) GT() antlr.TerminalNode {
	return s.GetToken(PRQLGT, 0)
}

func (s *OpaqueTokenContext) LTE() antlr.TerminalNode {
	return s.GetToken(PRQLLTE, 0)
}

func (s *OpaqueTokenContext) GTE() antlr.TerminalNode {
	return s.GetToken(PRQLGTE, 0)
}

func (s *OpaqueTokenContext) CAST_OP() antlr.TerminalNode {
	return s.GetToken(PRQLCAST_OP, 0)
}

func (s *OpaqueTokenContext) DOT() antlr.TerminalNode {
	return s.GetToken(PRQLDOT, 0)
}

func (s *OpaqueTokenContext) KW_AS() antlr.TerminalNode {
	return s.GetToken(PRQLKW_AS, 0)
}

func (s *OpaqueTokenContext) KW_FINAL() antlr.TerminalNode {
	return s.GetToken(PRQLKW_FINAL, 0)
}

func (s *OpaqueTokenContext) KW_INNER() antlr.TerminalNode {
	return s.GetToken(PRQLKW_INNER, 0)
}

func (s *OpaqueTokenContext) KW_LEFT() antlr.TerminalNode {
	return s.GetToken(PRQLKW_LEFT, 0)
}

func (s *OpaqueTokenContext) KW_RIGHT() antlr.TerminalNode {
	return s.GetToken(PRQLKW_RIGHT, 0)
}

func (s *OpaqueTokenContext) KW_FULL() antlr.TerminalNode {
	return s.GetToken(PRQLKW_FULL, 0)
}

func (s *OpaqueTokenContext) KW_SIDE() antlr.TerminalNode {
	return s.GetToken(PRQLKW_SIDE, 0)
}

func (s *OpaqueTokenContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(PRQLLPAREN, 0)
}

func (s *OpaqueTokenContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(PRQLRPAREN, 0)
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
	if listenerT, ok := listener.(PRQLListener); ok {
		listenerT.EnterOpaqueToken(s)
	}
}

func (s *OpaqueTokenContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PRQLListener); ok {
		listenerT.ExitOpaqueToken(s)
	}
}

func (s *OpaqueTokenContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PRQLVisitor:
		return t.VisitOpaqueToken(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PRQL) OpaqueToken() (localctx IOpaqueTokenContext) {
	localctx = NewOpaqueTokenContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, PRQLRULE_opaqueToken)
	var _la int

	p.SetState(283)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PRQLIDENT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(250)
			p.Match(PRQLIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PRQLINTEGER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(251)
			p.Match(PRQLINTEGER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PRQLFLOAT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(252)
			p.Match(PRQLFLOAT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PRQLSTRING:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(253)
			p.Match(PRQLSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PRQLSTAR:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(254)
			p.Match(PRQLSTAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PRQLPLUS:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(255)
			p.Match(PRQLPLUS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PRQLMINUS:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(256)
			p.Match(PRQLMINUS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PRQLSLASH:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(257)
			p.Match(PRQLSLASH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PRQLPERCENT:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(258)
			p.Match(PRQLPERCENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PRQLPIPE:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(259)
			p.Match(PRQLPIPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PRQLEQ:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(260)
			p.Match(PRQLEQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PRQLNEQ:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(261)
			p.Match(PRQLNEQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PRQLLT:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(262)
			p.Match(PRQLLT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PRQLGT:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(263)
			p.Match(PRQLGT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PRQLLTE:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(264)
			p.Match(PRQLLTE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PRQLGTE:
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(265)
			p.Match(PRQLGTE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PRQLCAST_OP:
		p.EnterOuterAlt(localctx, 17)
		{
			p.SetState(266)
			p.Match(PRQLCAST_OP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PRQLDOT:
		p.EnterOuterAlt(localctx, 18)
		{
			p.SetState(267)
			p.Match(PRQLDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PRQLKW_AS:
		p.EnterOuterAlt(localctx, 19)
		{
			p.SetState(268)
			p.Match(PRQLKW_AS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PRQLKW_FINAL:
		p.EnterOuterAlt(localctx, 20)
		{
			p.SetState(269)
			p.Match(PRQLKW_FINAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PRQLKW_INNER:
		p.EnterOuterAlt(localctx, 21)
		{
			p.SetState(270)
			p.Match(PRQLKW_INNER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PRQLKW_LEFT:
		p.EnterOuterAlt(localctx, 22)
		{
			p.SetState(271)
			p.Match(PRQLKW_LEFT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PRQLKW_RIGHT:
		p.EnterOuterAlt(localctx, 23)
		{
			p.SetState(272)
			p.Match(PRQLKW_RIGHT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PRQLKW_FULL:
		p.EnterOuterAlt(localctx, 24)
		{
			p.SetState(273)
			p.Match(PRQLKW_FULL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PRQLKW_SIDE:
		p.EnterOuterAlt(localctx, 25)
		{
			p.SetState(274)
			p.Match(PRQLKW_SIDE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PRQLLPAREN:
		p.EnterOuterAlt(localctx, 26)
		{
			p.SetState(275)
			p.Match(PRQLLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(279)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&140736771096576) != 0 {
			{
				p.SetState(276)
				p.OpaqueInner()
			}

			p.SetState(281)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(282)
			p.Match(PRQLRPAREN)
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
	p.RuleIndex = PRQLRULE_opaqueInner
	return p
}

func InitEmptyOpaqueInnerContext(p *OpaqueInnerContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PRQLRULE_opaqueInner
}

func (*OpaqueInnerContext) IsOpaqueInnerContext() {}

func NewOpaqueInnerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OpaqueInnerContext {
	var p = new(OpaqueInnerContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PRQLRULE_opaqueInner

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
	return s.GetToken(PRQLCOMMA, 0)
}

func (s *OpaqueInnerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OpaqueInnerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OpaqueInnerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PRQLListener); ok {
		listenerT.EnterOpaqueInner(s)
	}
}

func (s *OpaqueInnerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PRQLListener); ok {
		listenerT.ExitOpaqueInner(s)
	}
}

func (s *OpaqueInnerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PRQLVisitor:
		return t.VisitOpaqueInner(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PRQL) OpaqueInner() (localctx IOpaqueInnerContext) {
	localctx = NewOpaqueInnerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, PRQLRULE_opaqueInner)
	p.SetState(287)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PRQLKW_SIDE, PRQLKW_INNER, PRQLKW_LEFT, PRQLKW_RIGHT, PRQLKW_FULL, PRQLKW_AS, PRQLKW_FINAL, PRQLLPAREN, PRQLDOT, PRQLMINUS, PRQLPLUS, PRQLSTAR, PRQLSLASH, PRQLPERCENT, PRQLPIPE, PRQLCAST_OP, PRQLEQ, PRQLNEQ, PRQLLTE, PRQLGTE, PRQLLT, PRQLGT, PRQLFLOAT, PRQLINTEGER, PRQLSTRING, PRQLIDENT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(285)
			p.OpaqueToken()
		}

	case PRQLCOMMA:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(286)
			p.Match(PRQLCOMMA)
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
