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
		"'<='", "'>='", "'<'", "'>'",
	}
	staticData.SymbolicNames = []string{
		"", "COMMENT", "WS", "KW_FROM", "KW_FILTER", "KW_DERIVE", "KW_SELECT",
		"KW_GROUP", "KW_AGGREGATE", "KW_JOIN", "KW_SORT", "KW_TAKE", "KW_SKIP",
		"KW_WINDOW", "KW_ARRAY_JOIN", "KW_SIDE", "KW_INNER", "KW_LEFT", "KW_RIGHT",
		"KW_FULL", "KW_AS", "KW_FINAL", "LBRACE", "RBRACE", "LPAREN", "RPAREN",
		"COMMA", "COLON", "DOT", "RANGE", "MINUS", "PLUS", "STAR", "SLASH",
		"PERCENT", "PIPE", "CAST_OP", "EQ", "NEQ", "LTE", "GTE", "LT", "GT",
		"FLOAT", "INTEGER", "STRING", "IDENT", "NEWLINE",
	}
	staticData.RuleNames = []string{
		"query", "clause", "fromClause", "filterClause", "filterBody", "filterToken",
		"filterInner", "deriveClause", "selectClause", "groupClause", "joinClause",
		"joinSide", "joinCond", "joinCondExpr", "joinCondToken", "joinCondInner",
		"arrayJoinClause", "sortClause", "takeClause", "skipClause", "windowClause",
		"assignmentList", "assignment", "selectionList", "selectionItem", "keyList",
		"keyItem", "sortList", "sortItem", "opaqueExpr", "opaqueToken", "opaqueInner",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 47, 334, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 1, 0, 4, 0, 66, 8, 0, 11, 0, 12, 0, 67, 1, 0, 3, 0, 71, 8, 0, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 84, 8,
		1, 1, 2, 1, 2, 1, 2, 3, 2, 89, 8, 2, 1, 3, 1, 3, 1, 3, 1, 4, 4, 4, 95,
		8, 4, 11, 4, 12, 4, 96, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 5, 5, 127, 8, 5, 10, 5,
		12, 5, 130, 9, 5, 1, 5, 3, 5, 133, 8, 5, 1, 6, 1, 6, 3, 6, 137, 8, 6, 1,
		7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1,
		9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10,
		1, 10, 3, 10, 164, 8, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1,
		11, 1, 12, 1, 12, 1, 12, 1, 12, 3, 12, 177, 8, 12, 1, 13, 4, 13, 180, 8,
		13, 11, 13, 12, 13, 181, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14,
		1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 5, 14, 198, 8, 14, 10,
		14, 12, 14, 201, 9, 14, 1, 14, 3, 14, 204, 8, 14, 1, 15, 1, 15, 3, 15,
		208, 8, 15, 1, 16, 1, 16, 1, 16, 1, 16, 3, 16, 214, 8, 16, 1, 17, 1, 17,
		1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 3, 18, 225, 8, 18, 1,
		19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21,
		5, 21, 238, 8, 21, 10, 21, 12, 21, 241, 9, 21, 1, 22, 1, 22, 1, 22, 1,
		22, 1, 23, 1, 23, 1, 23, 5, 23, 250, 8, 23, 10, 23, 12, 23, 253, 9, 23,
		1, 24, 1, 24, 1, 24, 1, 24, 3, 24, 259, 8, 24, 1, 25, 1, 25, 1, 25, 5,
		25, 264, 8, 25, 10, 25, 12, 25, 267, 9, 25, 1, 26, 1, 26, 1, 26, 1, 26,
		3, 26, 273, 8, 26, 1, 27, 1, 27, 1, 27, 5, 27, 278, 8, 27, 10, 27, 12,
		27, 281, 9, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 3, 28, 288, 8, 28, 1,
		29, 4, 29, 291, 8, 29, 11, 29, 12, 29, 292, 1, 30, 1, 30, 1, 30, 1, 30,
		1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1,
		30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30,
		1, 30, 1, 30, 5, 30, 322, 8, 30, 10, 30, 12, 30, 325, 9, 30, 1, 30, 3,
		30, 328, 8, 30, 1, 31, 1, 31, 3, 31, 332, 8, 31, 1, 31, 0, 0, 32, 0, 2,
		4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40,
		42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 0, 1, 1, 0, 16, 19, 398, 0,
		65, 1, 0, 0, 0, 2, 83, 1, 0, 0, 0, 4, 85, 1, 0, 0, 0, 6, 90, 1, 0, 0, 0,
		8, 94, 1, 0, 0, 0, 10, 132, 1, 0, 0, 0, 12, 136, 1, 0, 0, 0, 14, 138, 1,
		0, 0, 0, 16, 143, 1, 0, 0, 0, 18, 148, 1, 0, 0, 0, 20, 159, 1, 0, 0, 0,
		22, 170, 1, 0, 0, 0, 24, 176, 1, 0, 0, 0, 26, 179, 1, 0, 0, 0, 28, 203,
		1, 0, 0, 0, 30, 207, 1, 0, 0, 0, 32, 209, 1, 0, 0, 0, 34, 215, 1, 0, 0,
		0, 36, 220, 1, 0, 0, 0, 38, 226, 1, 0, 0, 0, 40, 229, 1, 0, 0, 0, 42, 234,
		1, 0, 0, 0, 44, 242, 1, 0, 0, 0, 46, 246, 1, 0, 0, 0, 48, 258, 1, 0, 0,
		0, 50, 260, 1, 0, 0, 0, 52, 272, 1, 0, 0, 0, 54, 274, 1, 0, 0, 0, 56, 287,
		1, 0, 0, 0, 58, 290, 1, 0, 0, 0, 60, 327, 1, 0, 0, 0, 62, 331, 1, 0, 0,
		0, 64, 66, 3, 2, 1, 0, 65, 64, 1, 0, 0, 0, 66, 67, 1, 0, 0, 0, 67, 65,
		1, 0, 0, 0, 67, 68, 1, 0, 0, 0, 68, 70, 1, 0, 0, 0, 69, 71, 5, 0, 0, 1,
		70, 69, 1, 0, 0, 0, 70, 71, 1, 0, 0, 0, 71, 1, 1, 0, 0, 0, 72, 84, 3, 4,
		2, 0, 73, 84, 3, 6, 3, 0, 74, 84, 3, 14, 7, 0, 75, 84, 3, 16, 8, 0, 76,
		84, 3, 18, 9, 0, 77, 84, 3, 20, 10, 0, 78, 84, 3, 32, 16, 0, 79, 84, 3,
		34, 17, 0, 80, 84, 3, 36, 18, 0, 81, 84, 3, 38, 19, 0, 82, 84, 3, 40, 20,
		0, 83, 72, 1, 0, 0, 0, 83, 73, 1, 0, 0, 0, 83, 74, 1, 0, 0, 0, 83, 75,
		1, 0, 0, 0, 83, 76, 1, 0, 0, 0, 83, 77, 1, 0, 0, 0, 83, 78, 1, 0, 0, 0,
		83, 79, 1, 0, 0, 0, 83, 80, 1, 0, 0, 0, 83, 81, 1, 0, 0, 0, 83, 82, 1,
		0, 0, 0, 84, 3, 1, 0, 0, 0, 85, 86, 5, 3, 0, 0, 86, 88, 5, 46, 0, 0, 87,
		89, 5, 21, 0, 0, 88, 87, 1, 0, 0, 0, 88, 89, 1, 0, 0, 0, 89, 5, 1, 0, 0,
		0, 90, 91, 5, 4, 0, 0, 91, 92, 3, 8, 4, 0, 92, 7, 1, 0, 0, 0, 93, 95, 3,
		10, 5, 0, 94, 93, 1, 0, 0, 0, 95, 96, 1, 0, 0, 0, 96, 94, 1, 0, 0, 0, 96,
		97, 1, 0, 0, 0, 97, 9, 1, 0, 0, 0, 98, 133, 5, 46, 0, 0, 99, 133, 5, 44,
		0, 0, 100, 133, 5, 43, 0, 0, 101, 133, 5, 45, 0, 0, 102, 133, 5, 32, 0,
		0, 103, 133, 5, 31, 0, 0, 104, 133, 5, 30, 0, 0, 105, 133, 5, 33, 0, 0,
		106, 133, 5, 34, 0, 0, 107, 133, 5, 35, 0, 0, 108, 133, 5, 37, 0, 0, 109,
		133, 5, 38, 0, 0, 110, 133, 5, 41, 0, 0, 111, 133, 5, 42, 0, 0, 112, 133,
		5, 39, 0, 0, 113, 133, 5, 40, 0, 0, 114, 133, 5, 36, 0, 0, 115, 133, 5,
		28, 0, 0, 116, 133, 5, 29, 0, 0, 117, 133, 5, 20, 0, 0, 118, 133, 5, 21,
		0, 0, 119, 133, 5, 16, 0, 0, 120, 133, 5, 17, 0, 0, 121, 133, 5, 18, 0,
		0, 122, 133, 5, 19, 0, 0, 123, 133, 5, 15, 0, 0, 124, 128, 5, 24, 0, 0,
		125, 127, 3, 12, 6, 0, 126, 125, 1, 0, 0, 0, 127, 130, 1, 0, 0, 0, 128,
		126, 1, 0, 0, 0, 128, 129, 1, 0, 0, 0, 129, 131, 1, 0, 0, 0, 130, 128,
		1, 0, 0, 0, 131, 133, 5, 25, 0, 0, 132, 98, 1, 0, 0, 0, 132, 99, 1, 0,
		0, 0, 132, 100, 1, 0, 0, 0, 132, 101, 1, 0, 0, 0, 132, 102, 1, 0, 0, 0,
		132, 103, 1, 0, 0, 0, 132, 104, 1, 0, 0, 0, 132, 105, 1, 0, 0, 0, 132,
		106, 1, 0, 0, 0, 132, 107, 1, 0, 0, 0, 132, 108, 1, 0, 0, 0, 132, 109,
		1, 0, 0, 0, 132, 110, 1, 0, 0, 0, 132, 111, 1, 0, 0, 0, 132, 112, 1, 0,
		0, 0, 132, 113, 1, 0, 0, 0, 132, 114, 1, 0, 0, 0, 132, 115, 1, 0, 0, 0,
		132, 116, 1, 0, 0, 0, 132, 117, 1, 0, 0, 0, 132, 118, 1, 0, 0, 0, 132,
		119, 1, 0, 0, 0, 132, 120, 1, 0, 0, 0, 132, 121, 1, 0, 0, 0, 132, 122,
		1, 0, 0, 0, 132, 123, 1, 0, 0, 0, 132, 124, 1, 0, 0, 0, 133, 11, 1, 0,
		0, 0, 134, 137, 3, 10, 5, 0, 135, 137, 5, 26, 0, 0, 136, 134, 1, 0, 0,
		0, 136, 135, 1, 0, 0, 0, 137, 13, 1, 0, 0, 0, 138, 139, 5, 5, 0, 0, 139,
		140, 5, 22, 0, 0, 140, 141, 3, 42, 21, 0, 141, 142, 5, 23, 0, 0, 142, 15,
		1, 0, 0, 0, 143, 144, 5, 6, 0, 0, 144, 145, 5, 22, 0, 0, 145, 146, 3, 46,
		23, 0, 146, 147, 5, 23, 0, 0, 147, 17, 1, 0, 0, 0, 148, 149, 5, 7, 0, 0,
		149, 150, 5, 22, 0, 0, 150, 151, 3, 50, 25, 0, 151, 152, 5, 23, 0, 0, 152,
		153, 5, 24, 0, 0, 153, 154, 5, 8, 0, 0, 154, 155, 5, 22, 0, 0, 155, 156,
		3, 42, 21, 0, 156, 157, 5, 23, 0, 0, 157, 158, 5, 25, 0, 0, 158, 19, 1,
		0, 0, 0, 159, 163, 5, 9, 0, 0, 160, 161, 5, 15, 0, 0, 161, 162, 5, 27,
		0, 0, 162, 164, 3, 22, 11, 0, 163, 160, 1, 0, 0, 0, 163, 164, 1, 0, 0,
		0, 164, 165, 1, 0, 0, 0, 165, 166, 5, 46, 0, 0, 166, 167, 5, 24, 0, 0,
		167, 168, 3, 24, 12, 0, 168, 169, 5, 25, 0, 0, 169, 21, 1, 0, 0, 0, 170,
		171, 7, 0, 0, 0, 171, 23, 1, 0, 0, 0, 172, 173, 5, 37, 0, 0, 173, 174,
		5, 37, 0, 0, 174, 177, 5, 46, 0, 0, 175, 177, 3, 26, 13, 0, 176, 172, 1,
		0, 0, 0, 176, 175, 1, 0, 0, 0, 177, 25, 1, 0, 0, 0, 178, 180, 3, 28, 14,
		0, 179, 178, 1, 0, 0, 0, 180, 181, 1, 0, 0, 0, 181, 179, 1, 0, 0, 0, 181,
		182, 1, 0, 0, 0, 182, 27, 1, 0, 0, 0, 183, 204, 5, 46, 0, 0, 184, 204,
		5, 28, 0, 0, 185, 204, 5, 44, 0, 0, 186, 204, 5, 43, 0, 0, 187, 204, 5,
		45, 0, 0, 188, 204, 5, 37, 0, 0, 189, 204, 5, 38, 0, 0, 190, 204, 5, 41,
		0, 0, 191, 204, 5, 42, 0, 0, 192, 204, 5, 39, 0, 0, 193, 204, 5, 40, 0,
		0, 194, 204, 5, 36, 0, 0, 195, 199, 5, 24, 0, 0, 196, 198, 3, 30, 15, 0,
		197, 196, 1, 0, 0, 0, 198, 201, 1, 0, 0, 0, 199, 197, 1, 0, 0, 0, 199,
		200, 1, 0, 0, 0, 200, 202, 1, 0, 0, 0, 201, 199, 1, 0, 0, 0, 202, 204,
		5, 25, 0, 0, 203, 183, 1, 0, 0, 0, 203, 184, 1, 0, 0, 0, 203, 185, 1, 0,
		0, 0, 203, 186, 1, 0, 0, 0, 203, 187, 1, 0, 0, 0, 203, 188, 1, 0, 0, 0,
		203, 189, 1, 0, 0, 0, 203, 190, 1, 0, 0, 0, 203, 191, 1, 0, 0, 0, 203,
		192, 1, 0, 0, 0, 203, 193, 1, 0, 0, 0, 203, 194, 1, 0, 0, 0, 203, 195,
		1, 0, 0, 0, 204, 29, 1, 0, 0, 0, 205, 208, 3, 28, 14, 0, 206, 208, 5, 26,
		0, 0, 207, 205, 1, 0, 0, 0, 207, 206, 1, 0, 0, 0, 208, 31, 1, 0, 0, 0,
		209, 210, 5, 14, 0, 0, 210, 213, 5, 46, 0, 0, 211, 212, 5, 20, 0, 0, 212,
		214, 5, 46, 0, 0, 213, 211, 1, 0, 0, 0, 213, 214, 1, 0, 0, 0, 214, 33,
		1, 0, 0, 0, 215, 216, 5, 10, 0, 0, 216, 217, 5, 22, 0, 0, 217, 218, 3,
		54, 27, 0, 218, 219, 5, 23, 0, 0, 219, 35, 1, 0, 0, 0, 220, 221, 5, 11,
		0, 0, 221, 224, 5, 44, 0, 0, 222, 223, 5, 29, 0, 0, 223, 225, 5, 44, 0,
		0, 224, 222, 1, 0, 0, 0, 224, 225, 1, 0, 0, 0, 225, 37, 1, 0, 0, 0, 226,
		227, 5, 12, 0, 0, 227, 228, 5, 44, 0, 0, 228, 39, 1, 0, 0, 0, 229, 230,
		5, 13, 0, 0, 230, 231, 5, 22, 0, 0, 231, 232, 3, 42, 21, 0, 232, 233, 5,
		23, 0, 0, 233, 41, 1, 0, 0, 0, 234, 239, 3, 44, 22, 0, 235, 236, 5, 26,
		0, 0, 236, 238, 3, 44, 22, 0, 237, 235, 1, 0, 0, 0, 238, 241, 1, 0, 0,
		0, 239, 237, 1, 0, 0, 0, 239, 240, 1, 0, 0, 0, 240, 43, 1, 0, 0, 0, 241,
		239, 1, 0, 0, 0, 242, 243, 5, 46, 0, 0, 243, 244, 5, 37, 0, 0, 244, 245,
		3, 58, 29, 0, 245, 45, 1, 0, 0, 0, 246, 251, 3, 48, 24, 0, 247, 248, 5,
		26, 0, 0, 248, 250, 3, 48, 24, 0, 249, 247, 1, 0, 0, 0, 250, 253, 1, 0,
		0, 0, 251, 249, 1, 0, 0, 0, 251, 252, 1, 0, 0, 0, 252, 47, 1, 0, 0, 0,
		253, 251, 1, 0, 0, 0, 254, 255, 5, 46, 0, 0, 255, 256, 5, 37, 0, 0, 256,
		259, 3, 58, 29, 0, 257, 259, 3, 58, 29, 0, 258, 254, 1, 0, 0, 0, 258, 257,
		1, 0, 0, 0, 259, 49, 1, 0, 0, 0, 260, 265, 3, 52, 26, 0, 261, 262, 5, 26,
		0, 0, 262, 264, 3, 52, 26, 0, 263, 261, 1, 0, 0, 0, 264, 267, 1, 0, 0,
		0, 265, 263, 1, 0, 0, 0, 265, 266, 1, 0, 0, 0, 266, 51, 1, 0, 0, 0, 267,
		265, 1, 0, 0, 0, 268, 269, 5, 46, 0, 0, 269, 270, 5, 37, 0, 0, 270, 273,
		3, 58, 29, 0, 271, 273, 5, 46, 0, 0, 272, 268, 1, 0, 0, 0, 272, 271, 1,
		0, 0, 0, 273, 53, 1, 0, 0, 0, 274, 279, 3, 56, 28, 0, 275, 276, 5, 26,
		0, 0, 276, 278, 3, 56, 28, 0, 277, 275, 1, 0, 0, 0, 278, 281, 1, 0, 0,
		0, 279, 277, 1, 0, 0, 0, 279, 280, 1, 0, 0, 0, 280, 55, 1, 0, 0, 0, 281,
		279, 1, 0, 0, 0, 282, 283, 5, 30, 0, 0, 283, 288, 5, 46, 0, 0, 284, 285,
		5, 31, 0, 0, 285, 288, 5, 46, 0, 0, 286, 288, 5, 46, 0, 0, 287, 282, 1,
		0, 0, 0, 287, 284, 1, 0, 0, 0, 287, 286, 1, 0, 0, 0, 288, 57, 1, 0, 0,
		0, 289, 291, 3, 60, 30, 0, 290, 289, 1, 0, 0, 0, 291, 292, 1, 0, 0, 0,
		292, 290, 1, 0, 0, 0, 292, 293, 1, 0, 0, 0, 293, 59, 1, 0, 0, 0, 294, 328,
		5, 46, 0, 0, 295, 328, 5, 44, 0, 0, 296, 328, 5, 43, 0, 0, 297, 328, 5,
		45, 0, 0, 298, 328, 5, 32, 0, 0, 299, 328, 5, 31, 0, 0, 300, 328, 5, 30,
		0, 0, 301, 328, 5, 33, 0, 0, 302, 328, 5, 34, 0, 0, 303, 328, 5, 35, 0,
		0, 304, 328, 5, 37, 0, 0, 305, 328, 5, 38, 0, 0, 306, 328, 5, 41, 0, 0,
		307, 328, 5, 42, 0, 0, 308, 328, 5, 39, 0, 0, 309, 328, 5, 40, 0, 0, 310,
		328, 5, 36, 0, 0, 311, 328, 5, 28, 0, 0, 312, 328, 5, 20, 0, 0, 313, 328,
		5, 21, 0, 0, 314, 328, 5, 16, 0, 0, 315, 328, 5, 17, 0, 0, 316, 328, 5,
		18, 0, 0, 317, 328, 5, 19, 0, 0, 318, 328, 5, 15, 0, 0, 319, 323, 5, 24,
		0, 0, 320, 322, 3, 62, 31, 0, 321, 320, 1, 0, 0, 0, 322, 325, 1, 0, 0,
		0, 323, 321, 1, 0, 0, 0, 323, 324, 1, 0, 0, 0, 324, 326, 1, 0, 0, 0, 325,
		323, 1, 0, 0, 0, 326, 328, 5, 25, 0, 0, 327, 294, 1, 0, 0, 0, 327, 295,
		1, 0, 0, 0, 327, 296, 1, 0, 0, 0, 327, 297, 1, 0, 0, 0, 327, 298, 1, 0,
		0, 0, 327, 299, 1, 0, 0, 0, 327, 300, 1, 0, 0, 0, 327, 301, 1, 0, 0, 0,
		327, 302, 1, 0, 0, 0, 327, 303, 1, 0, 0, 0, 327, 304, 1, 0, 0, 0, 327,
		305, 1, 0, 0, 0, 327, 306, 1, 0, 0, 0, 327, 307, 1, 0, 0, 0, 327, 308,
		1, 0, 0, 0, 327, 309, 1, 0, 0, 0, 327, 310, 1, 0, 0, 0, 327, 311, 1, 0,
		0, 0, 327, 312, 1, 0, 0, 0, 327, 313, 1, 0, 0, 0, 327, 314, 1, 0, 0, 0,
		327, 315, 1, 0, 0, 0, 327, 316, 1, 0, 0, 0, 327, 317, 1, 0, 0, 0, 327,
		318, 1, 0, 0, 0, 327, 319, 1, 0, 0, 0, 328, 61, 1, 0, 0, 0, 329, 332, 3,
		60, 30, 0, 330, 332, 5, 26, 0, 0, 331, 329, 1, 0, 0, 0, 331, 330, 1, 0,
		0, 0, 332, 63, 1, 0, 0, 0, 27, 67, 70, 83, 88, 96, 128, 132, 136, 163,
		176, 181, 199, 203, 207, 213, 224, 239, 251, 258, 265, 272, 279, 287, 292,
		323, 327, 331,
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
)

// PRQL rules.
const (
	PRQLRULE_query           = 0
	PRQLRULE_clause          = 1
	PRQLRULE_fromClause      = 2
	PRQLRULE_filterClause    = 3
	PRQLRULE_filterBody      = 4
	PRQLRULE_filterToken     = 5
	PRQLRULE_filterInner     = 6
	PRQLRULE_deriveClause    = 7
	PRQLRULE_selectClause    = 8
	PRQLRULE_groupClause     = 9
	PRQLRULE_joinClause      = 10
	PRQLRULE_joinSide        = 11
	PRQLRULE_joinCond        = 12
	PRQLRULE_joinCondExpr    = 13
	PRQLRULE_joinCondToken   = 14
	PRQLRULE_joinCondInner   = 15
	PRQLRULE_arrayJoinClause = 16
	PRQLRULE_sortClause      = 17
	PRQLRULE_takeClause      = 18
	PRQLRULE_skipClause      = 19
	PRQLRULE_windowClause    = 20
	PRQLRULE_assignmentList  = 21
	PRQLRULE_assignment      = 22
	PRQLRULE_selectionList   = 23
	PRQLRULE_selectionItem   = 24
	PRQLRULE_keyList         = 25
	PRQLRULE_keyItem         = 26
	PRQLRULE_sortList        = 27
	PRQLRULE_sortItem        = 28
	PRQLRULE_opaqueExpr      = 29
	PRQLRULE_opaqueToken     = 30
	PRQLRULE_opaqueInner     = 31
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
	p.SetState(65)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&32504) != 0) {
		{
			p.SetState(64)
			p.Clause()
		}

		p.SetState(67)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(70)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(69)
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
	p.SetState(83)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PRQLKW_FROM:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(72)
			p.FromClause()
		}

	case PRQLKW_FILTER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(73)
			p.FilterClause()
		}

	case PRQLKW_DERIVE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(74)
			p.DeriveClause()
		}

	case PRQLKW_SELECT:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(75)
			p.SelectClause()
		}

	case PRQLKW_GROUP:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(76)
			p.GroupClause()
		}

	case PRQLKW_JOIN:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(77)
			p.JoinClause()
		}

	case PRQLKW_ARRAY_JOIN:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(78)
			p.ArrayJoinClause()
		}

	case PRQLKW_SORT:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(79)
			p.SortClause()
		}

	case PRQLKW_TAKE:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(80)
			p.TakeClause()
		}

	case PRQLKW_SKIP:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(81)
			p.SkipClause()
		}

	case PRQLKW_WINDOW:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(82)
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
		p.SetState(85)
		p.Match(PRQLKW_FROM)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(86)
		p.Match(PRQLIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(88)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PRQLKW_FINAL {
		{
			p.SetState(87)
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
		p.SetState(90)
		p.Match(PRQLKW_FILTER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(91)
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
	AllFilterToken() []IFilterTokenContext
	FilterToken(i int) IFilterTokenContext

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

func (s *FilterBodyContext) AllFilterToken() []IFilterTokenContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFilterTokenContext); ok {
			len++
		}
	}

	tst := make([]IFilterTokenContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFilterTokenContext); ok {
			tst[i] = t.(IFilterTokenContext)
			i++
		}
	}

	return tst
}

func (s *FilterBodyContext) FilterToken(i int) IFilterTokenContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFilterTokenContext); ok {
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

	return t.(IFilterTokenContext)
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
	p.SetState(94)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&140737240858624) != 0) {
		{
			p.SetState(93)
			p.FilterToken()
		}

		p.SetState(96)
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

// IFilterTokenContext is an interface to support dynamic dispatch.
type IFilterTokenContext interface {
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
	RANGE() antlr.TerminalNode
	KW_AS() antlr.TerminalNode
	KW_FINAL() antlr.TerminalNode
	KW_INNER() antlr.TerminalNode
	KW_LEFT() antlr.TerminalNode
	KW_RIGHT() antlr.TerminalNode
	KW_FULL() antlr.TerminalNode
	KW_SIDE() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	AllFilterInner() []IFilterInnerContext
	FilterInner(i int) IFilterInnerContext

	// IsFilterTokenContext differentiates from other interfaces.
	IsFilterTokenContext()
}

type FilterTokenContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFilterTokenContext() *FilterTokenContext {
	var p = new(FilterTokenContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PRQLRULE_filterToken
	return p
}

func InitEmptyFilterTokenContext(p *FilterTokenContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PRQLRULE_filterToken
}

func (*FilterTokenContext) IsFilterTokenContext() {}

func NewFilterTokenContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FilterTokenContext {
	var p = new(FilterTokenContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PRQLRULE_filterToken

	return p
}

func (s *FilterTokenContext) GetParser() antlr.Parser { return s.parser }

func (s *FilterTokenContext) IDENT() antlr.TerminalNode {
	return s.GetToken(PRQLIDENT, 0)
}

func (s *FilterTokenContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(PRQLINTEGER, 0)
}

func (s *FilterTokenContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(PRQLFLOAT, 0)
}

func (s *FilterTokenContext) STRING() antlr.TerminalNode {
	return s.GetToken(PRQLSTRING, 0)
}

func (s *FilterTokenContext) STAR() antlr.TerminalNode {
	return s.GetToken(PRQLSTAR, 0)
}

func (s *FilterTokenContext) PLUS() antlr.TerminalNode {
	return s.GetToken(PRQLPLUS, 0)
}

func (s *FilterTokenContext) MINUS() antlr.TerminalNode {
	return s.GetToken(PRQLMINUS, 0)
}

func (s *FilterTokenContext) SLASH() antlr.TerminalNode {
	return s.GetToken(PRQLSLASH, 0)
}

func (s *FilterTokenContext) PERCENT() antlr.TerminalNode {
	return s.GetToken(PRQLPERCENT, 0)
}

func (s *FilterTokenContext) PIPE() antlr.TerminalNode {
	return s.GetToken(PRQLPIPE, 0)
}

func (s *FilterTokenContext) EQ() antlr.TerminalNode {
	return s.GetToken(PRQLEQ, 0)
}

func (s *FilterTokenContext) NEQ() antlr.TerminalNode {
	return s.GetToken(PRQLNEQ, 0)
}

func (s *FilterTokenContext) LT() antlr.TerminalNode {
	return s.GetToken(PRQLLT, 0)
}

func (s *FilterTokenContext) GT() antlr.TerminalNode {
	return s.GetToken(PRQLGT, 0)
}

func (s *FilterTokenContext) LTE() antlr.TerminalNode {
	return s.GetToken(PRQLLTE, 0)
}

func (s *FilterTokenContext) GTE() antlr.TerminalNode {
	return s.GetToken(PRQLGTE, 0)
}

func (s *FilterTokenContext) CAST_OP() antlr.TerminalNode {
	return s.GetToken(PRQLCAST_OP, 0)
}

func (s *FilterTokenContext) DOT() antlr.TerminalNode {
	return s.GetToken(PRQLDOT, 0)
}

func (s *FilterTokenContext) RANGE() antlr.TerminalNode {
	return s.GetToken(PRQLRANGE, 0)
}

func (s *FilterTokenContext) KW_AS() antlr.TerminalNode {
	return s.GetToken(PRQLKW_AS, 0)
}

func (s *FilterTokenContext) KW_FINAL() antlr.TerminalNode {
	return s.GetToken(PRQLKW_FINAL, 0)
}

func (s *FilterTokenContext) KW_INNER() antlr.TerminalNode {
	return s.GetToken(PRQLKW_INNER, 0)
}

func (s *FilterTokenContext) KW_LEFT() antlr.TerminalNode {
	return s.GetToken(PRQLKW_LEFT, 0)
}

func (s *FilterTokenContext) KW_RIGHT() antlr.TerminalNode {
	return s.GetToken(PRQLKW_RIGHT, 0)
}

func (s *FilterTokenContext) KW_FULL() antlr.TerminalNode {
	return s.GetToken(PRQLKW_FULL, 0)
}

func (s *FilterTokenContext) KW_SIDE() antlr.TerminalNode {
	return s.GetToken(PRQLKW_SIDE, 0)
}

func (s *FilterTokenContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(PRQLLPAREN, 0)
}

func (s *FilterTokenContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(PRQLRPAREN, 0)
}

func (s *FilterTokenContext) AllFilterInner() []IFilterInnerContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFilterInnerContext); ok {
			len++
		}
	}

	tst := make([]IFilterInnerContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFilterInnerContext); ok {
			tst[i] = t.(IFilterInnerContext)
			i++
		}
	}

	return tst
}

func (s *FilterTokenContext) FilterInner(i int) IFilterInnerContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFilterInnerContext); ok {
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

	return t.(IFilterInnerContext)
}

func (s *FilterTokenContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FilterTokenContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FilterTokenContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PRQLListener); ok {
		listenerT.EnterFilterToken(s)
	}
}

func (s *FilterTokenContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PRQLListener); ok {
		listenerT.ExitFilterToken(s)
	}
}

func (s *FilterTokenContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PRQLVisitor:
		return t.VisitFilterToken(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PRQL) FilterToken() (localctx IFilterTokenContext) {
	localctx = NewFilterTokenContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, PRQLRULE_filterToken)
	var _la int

	p.SetState(132)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PRQLIDENT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(98)
			p.Match(PRQLIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PRQLINTEGER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(99)
			p.Match(PRQLINTEGER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PRQLFLOAT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(100)
			p.Match(PRQLFLOAT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PRQLSTRING:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(101)
			p.Match(PRQLSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PRQLSTAR:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(102)
			p.Match(PRQLSTAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PRQLPLUS:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(103)
			p.Match(PRQLPLUS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PRQLMINUS:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(104)
			p.Match(PRQLMINUS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PRQLSLASH:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(105)
			p.Match(PRQLSLASH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PRQLPERCENT:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(106)
			p.Match(PRQLPERCENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PRQLPIPE:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(107)
			p.Match(PRQLPIPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PRQLEQ:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(108)
			p.Match(PRQLEQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PRQLNEQ:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(109)
			p.Match(PRQLNEQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PRQLLT:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(110)
			p.Match(PRQLLT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PRQLGT:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(111)
			p.Match(PRQLGT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PRQLLTE:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(112)
			p.Match(PRQLLTE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PRQLGTE:
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(113)
			p.Match(PRQLGTE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PRQLCAST_OP:
		p.EnterOuterAlt(localctx, 17)
		{
			p.SetState(114)
			p.Match(PRQLCAST_OP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PRQLDOT:
		p.EnterOuterAlt(localctx, 18)
		{
			p.SetState(115)
			p.Match(PRQLDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PRQLRANGE:
		p.EnterOuterAlt(localctx, 19)
		{
			p.SetState(116)
			p.Match(PRQLRANGE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PRQLKW_AS:
		p.EnterOuterAlt(localctx, 20)
		{
			p.SetState(117)
			p.Match(PRQLKW_AS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PRQLKW_FINAL:
		p.EnterOuterAlt(localctx, 21)
		{
			p.SetState(118)
			p.Match(PRQLKW_FINAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PRQLKW_INNER:
		p.EnterOuterAlt(localctx, 22)
		{
			p.SetState(119)
			p.Match(PRQLKW_INNER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PRQLKW_LEFT:
		p.EnterOuterAlt(localctx, 23)
		{
			p.SetState(120)
			p.Match(PRQLKW_LEFT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PRQLKW_RIGHT:
		p.EnterOuterAlt(localctx, 24)
		{
			p.SetState(121)
			p.Match(PRQLKW_RIGHT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PRQLKW_FULL:
		p.EnterOuterAlt(localctx, 25)
		{
			p.SetState(122)
			p.Match(PRQLKW_FULL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PRQLKW_SIDE:
		p.EnterOuterAlt(localctx, 26)
		{
			p.SetState(123)
			p.Match(PRQLKW_SIDE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PRQLLPAREN:
		p.EnterOuterAlt(localctx, 27)
		{
			p.SetState(124)
			p.Match(PRQLLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(128)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&140737307967488) != 0 {
			{
				p.SetState(125)
				p.FilterInner()
			}

			p.SetState(130)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(131)
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

// IFilterInnerContext is an interface to support dynamic dispatch.
type IFilterInnerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FilterToken() IFilterTokenContext
	COMMA() antlr.TerminalNode

	// IsFilterInnerContext differentiates from other interfaces.
	IsFilterInnerContext()
}

type FilterInnerContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFilterInnerContext() *FilterInnerContext {
	var p = new(FilterInnerContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PRQLRULE_filterInner
	return p
}

func InitEmptyFilterInnerContext(p *FilterInnerContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PRQLRULE_filterInner
}

func (*FilterInnerContext) IsFilterInnerContext() {}

func NewFilterInnerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FilterInnerContext {
	var p = new(FilterInnerContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PRQLRULE_filterInner

	return p
}

func (s *FilterInnerContext) GetParser() antlr.Parser { return s.parser }

func (s *FilterInnerContext) FilterToken() IFilterTokenContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFilterTokenContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFilterTokenContext)
}

func (s *FilterInnerContext) COMMA() antlr.TerminalNode {
	return s.GetToken(PRQLCOMMA, 0)
}

func (s *FilterInnerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FilterInnerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FilterInnerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PRQLListener); ok {
		listenerT.EnterFilterInner(s)
	}
}

func (s *FilterInnerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PRQLListener); ok {
		listenerT.ExitFilterInner(s)
	}
}

func (s *FilterInnerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PRQLVisitor:
		return t.VisitFilterInner(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PRQL) FilterInner() (localctx IFilterInnerContext) {
	localctx = NewFilterInnerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, PRQLRULE_filterInner)
	p.SetState(136)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PRQLKW_SIDE, PRQLKW_INNER, PRQLKW_LEFT, PRQLKW_RIGHT, PRQLKW_FULL, PRQLKW_AS, PRQLKW_FINAL, PRQLLPAREN, PRQLDOT, PRQLRANGE, PRQLMINUS, PRQLPLUS, PRQLSTAR, PRQLSLASH, PRQLPERCENT, PRQLPIPE, PRQLCAST_OP, PRQLEQ, PRQLNEQ, PRQLLTE, PRQLGTE, PRQLLT, PRQLGT, PRQLFLOAT, PRQLINTEGER, PRQLSTRING, PRQLIDENT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(134)
			p.FilterToken()
		}

	case PRQLCOMMA:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(135)
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
	p.EnterRule(localctx, 14, PRQLRULE_deriveClause)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(138)
		p.Match(PRQLKW_DERIVE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(139)
		p.Match(PRQLLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(140)
		p.AssignmentList()
	}
	{
		p.SetState(141)
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
	p.EnterRule(localctx, 16, PRQLRULE_selectClause)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(143)
		p.Match(PRQLKW_SELECT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(144)
		p.Match(PRQLLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(145)
		p.SelectionList()
	}
	{
		p.SetState(146)
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
	p.EnterRule(localctx, 18, PRQLRULE_groupClause)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(148)
		p.Match(PRQLKW_GROUP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(149)
		p.Match(PRQLLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(150)
		p.KeyList()
	}
	{
		p.SetState(151)
		p.Match(PRQLRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(152)
		p.Match(PRQLLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(153)
		p.Match(PRQLKW_AGGREGATE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(154)
		p.Match(PRQLLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(155)
		p.AssignmentList()
	}
	{
		p.SetState(156)
		p.Match(PRQLRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(157)
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
	p.EnterRule(localctx, 20, PRQLRULE_joinClause)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(159)
		p.Match(PRQLKW_JOIN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(163)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PRQLKW_SIDE {
		{
			p.SetState(160)
			p.Match(PRQLKW_SIDE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(161)
			p.Match(PRQLCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(162)
			p.JoinSide()
		}

	}
	{
		p.SetState(165)
		p.Match(PRQLIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(166)
		p.Match(PRQLLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(167)
		p.JoinCond()
	}
	{
		p.SetState(168)
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
	p.EnterRule(localctx, 22, PRQLRULE_joinSide)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(170)
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
	p.EnterRule(localctx, 24, PRQLRULE_joinCond)
	p.SetState(176)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		localctx = NewSelfJoinCondContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(172)
			p.Match(PRQLEQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(173)
			p.Match(PRQLEQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(174)
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
			p.SetState(175)
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
	p.EnterRule(localctx, 26, PRQLRULE_joinCondExpr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(179)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&140669054091264) != 0) {
		{
			p.SetState(178)
			p.JoinCondToken()
		}

		p.SetState(181)
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
	p.EnterRule(localctx, 28, PRQLRULE_joinCondToken)
	var _la int

	p.SetState(203)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PRQLIDENT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(183)
			p.Match(PRQLIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PRQLDOT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(184)
			p.Match(PRQLDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PRQLINTEGER:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(185)
			p.Match(PRQLINTEGER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PRQLFLOAT:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(186)
			p.Match(PRQLFLOAT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PRQLSTRING:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(187)
			p.Match(PRQLSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PRQLEQ:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(188)
			p.Match(PRQLEQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PRQLNEQ:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(189)
			p.Match(PRQLNEQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PRQLLT:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(190)
			p.Match(PRQLLT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PRQLGT:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(191)
			p.Match(PRQLGT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PRQLLTE:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(192)
			p.Match(PRQLLTE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PRQLGTE:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(193)
			p.Match(PRQLGTE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PRQLCAST_OP:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(194)
			p.Match(PRQLCAST_OP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PRQLLPAREN:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(195)
			p.Match(PRQLLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(199)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&140669121200128) != 0 {
			{
				p.SetState(196)
				p.JoinCondInner()
			}

			p.SetState(201)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(202)
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
	p.EnterRule(localctx, 30, PRQLRULE_joinCondInner)
	p.SetState(207)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PRQLLPAREN, PRQLDOT, PRQLCAST_OP, PRQLEQ, PRQLNEQ, PRQLLTE, PRQLGTE, PRQLLT, PRQLGT, PRQLFLOAT, PRQLINTEGER, PRQLSTRING, PRQLIDENT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(205)
			p.JoinCondToken()
		}

	case PRQLCOMMA:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(206)
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
	p.EnterRule(localctx, 32, PRQLRULE_arrayJoinClause)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(209)
		p.Match(PRQLKW_ARRAY_JOIN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(210)
		p.Match(PRQLIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(213)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PRQLKW_AS {
		{
			p.SetState(211)
			p.Match(PRQLKW_AS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(212)
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
	p.EnterRule(localctx, 34, PRQLRULE_sortClause)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(215)
		p.Match(PRQLKW_SORT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(216)
		p.Match(PRQLLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(217)
		p.SortList()
	}
	{
		p.SetState(218)
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
	p.EnterRule(localctx, 36, PRQLRULE_takeClause)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(220)
		p.Match(PRQLKW_TAKE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(221)
		p.Match(PRQLINTEGER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(224)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PRQLRANGE {
		{
			p.SetState(222)
			p.Match(PRQLRANGE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(223)
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
	p.EnterRule(localctx, 38, PRQLRULE_skipClause)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(226)
		p.Match(PRQLKW_SKIP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(227)
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
	p.EnterRule(localctx, 40, PRQLRULE_windowClause)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(229)
		p.Match(PRQLKW_WINDOW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(230)
		p.Match(PRQLLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(231)
		p.AssignmentList()
	}
	{
		p.SetState(232)
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
	p.EnterRule(localctx, 42, PRQLRULE_assignmentList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(234)
		p.Assignment()
	}
	p.SetState(239)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == PRQLCOMMA {
		{
			p.SetState(235)
			p.Match(PRQLCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(236)
			p.Assignment()
		}

		p.SetState(241)
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
	p.EnterRule(localctx, 44, PRQLRULE_assignment)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(242)
		p.Match(PRQLIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(243)
		p.Match(PRQLEQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(244)
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
	p.EnterRule(localctx, 46, PRQLRULE_selectionList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(246)
		p.SelectionItem()
	}
	p.SetState(251)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == PRQLCOMMA {
		{
			p.SetState(247)
			p.Match(PRQLCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(248)
			p.SelectionItem()
		}

		p.SetState(253)
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
	p.EnterRule(localctx, 48, PRQLRULE_selectionItem)
	p.SetState(258)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		localctx = NewAliasedSelectionContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(254)
			p.Match(PRQLIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(255)
			p.Match(PRQLEQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(256)
			p.OpaqueExpr()
		}

	case 2:
		localctx = NewBareSelectionContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(257)
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
	p.EnterRule(localctx, 50, PRQLRULE_keyList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(260)
		p.KeyItem()
	}
	p.SetState(265)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == PRQLCOMMA {
		{
			p.SetState(261)
			p.Match(PRQLCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(262)
			p.KeyItem()
		}

		p.SetState(267)
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
	p.EnterRule(localctx, 52, PRQLRULE_keyItem)
	p.SetState(272)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext()) {
	case 1:
		localctx = NewComputedKeyContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(268)
			p.Match(PRQLIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(269)
			p.Match(PRQLEQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(270)
			p.OpaqueExpr()
		}

	case 2:
		localctx = NewColumnKeyContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(271)
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
	p.EnterRule(localctx, 54, PRQLRULE_sortList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(274)
		p.SortItem()
	}
	p.SetState(279)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == PRQLCOMMA {
		{
			p.SetState(275)
			p.Match(PRQLCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(276)
			p.SortItem()
		}

		p.SetState(281)
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
	p.EnterRule(localctx, 56, PRQLRULE_sortItem)
	p.SetState(287)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PRQLMINUS:
		localctx = NewDescSortContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(282)
			p.Match(PRQLMINUS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(283)
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
			p.SetState(284)
			p.Match(PRQLPLUS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(285)
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
			p.SetState(286)
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
	p.EnterRule(localctx, 58, PRQLRULE_opaqueExpr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(290)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&140736703987712) != 0) {
		{
			p.SetState(289)
			p.OpaqueToken()
		}

		p.SetState(292)
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
	p.EnterRule(localctx, 60, PRQLRULE_opaqueToken)
	var _la int

	p.SetState(327)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PRQLIDENT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(294)
			p.Match(PRQLIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PRQLINTEGER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(295)
			p.Match(PRQLINTEGER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PRQLFLOAT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(296)
			p.Match(PRQLFLOAT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PRQLSTRING:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(297)
			p.Match(PRQLSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PRQLSTAR:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(298)
			p.Match(PRQLSTAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PRQLPLUS:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(299)
			p.Match(PRQLPLUS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PRQLMINUS:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(300)
			p.Match(PRQLMINUS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PRQLSLASH:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(301)
			p.Match(PRQLSLASH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PRQLPERCENT:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(302)
			p.Match(PRQLPERCENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PRQLPIPE:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(303)
			p.Match(PRQLPIPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PRQLEQ:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(304)
			p.Match(PRQLEQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PRQLNEQ:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(305)
			p.Match(PRQLNEQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PRQLLT:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(306)
			p.Match(PRQLLT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PRQLGT:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(307)
			p.Match(PRQLGT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PRQLLTE:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(308)
			p.Match(PRQLLTE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PRQLGTE:
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(309)
			p.Match(PRQLGTE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PRQLCAST_OP:
		p.EnterOuterAlt(localctx, 17)
		{
			p.SetState(310)
			p.Match(PRQLCAST_OP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PRQLDOT:
		p.EnterOuterAlt(localctx, 18)
		{
			p.SetState(311)
			p.Match(PRQLDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PRQLKW_AS:
		p.EnterOuterAlt(localctx, 19)
		{
			p.SetState(312)
			p.Match(PRQLKW_AS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PRQLKW_FINAL:
		p.EnterOuterAlt(localctx, 20)
		{
			p.SetState(313)
			p.Match(PRQLKW_FINAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PRQLKW_INNER:
		p.EnterOuterAlt(localctx, 21)
		{
			p.SetState(314)
			p.Match(PRQLKW_INNER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PRQLKW_LEFT:
		p.EnterOuterAlt(localctx, 22)
		{
			p.SetState(315)
			p.Match(PRQLKW_LEFT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PRQLKW_RIGHT:
		p.EnterOuterAlt(localctx, 23)
		{
			p.SetState(316)
			p.Match(PRQLKW_RIGHT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PRQLKW_FULL:
		p.EnterOuterAlt(localctx, 24)
		{
			p.SetState(317)
			p.Match(PRQLKW_FULL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PRQLKW_SIDE:
		p.EnterOuterAlt(localctx, 25)
		{
			p.SetState(318)
			p.Match(PRQLKW_SIDE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PRQLLPAREN:
		p.EnterOuterAlt(localctx, 26)
		{
			p.SetState(319)
			p.Match(PRQLLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(323)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&140736771096576) != 0 {
			{
				p.SetState(320)
				p.OpaqueInner()
			}

			p.SetState(325)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(326)
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
	p.EnterRule(localctx, 62, PRQLRULE_opaqueInner)
	p.SetState(331)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PRQLKW_SIDE, PRQLKW_INNER, PRQLKW_LEFT, PRQLKW_RIGHT, PRQLKW_FULL, PRQLKW_AS, PRQLKW_FINAL, PRQLLPAREN, PRQLDOT, PRQLMINUS, PRQLPLUS, PRQLSTAR, PRQLSLASH, PRQLPERCENT, PRQLPIPE, PRQLCAST_OP, PRQLEQ, PRQLNEQ, PRQLLTE, PRQLGTE, PRQLLT, PRQLGT, PRQLFLOAT, PRQLINTEGER, PRQLSTRING, PRQLIDENT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(329)
			p.OpaqueToken()
		}

	case PRQLCOMMA:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(330)
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
