// Code generated from SQLMesh.g4 by ANTLR 4.13.2. DO NOT EDIT.

package grammar // SQLMesh
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

type SQLMesh struct {
	*antlr.BaseParser
}

var SQLMeshParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func sqlmeshParserInit() {
	staticData := &SQLMeshParserStaticData
	staticData.LiteralNames = []string{
		"", "", "", "", "'MODEL'", "'AUDIT'", "'INCREMENTAL_BY_TIME_RANGE'",
		"'INCREMENTAL_BY_UNIQUE_KEY'", "'INCREMENTAL_BY_PARTITION'", "'SCD_TYPE_2'",
		"'EXTERNAL'", "'EMBEDDED'", "'FULL'", "'VIEW'", "'SEED'", "", "", "'column_descriptions'",
		"'physical_properties'", "'on_virtual_update'", "'invalidate_hard_deletes'",
		"'disable_restatement'", "'pre_statements'", "'post_statements'", "'partitioned_by'",
		"'clustered_by'", "'storage_format'", "'table_format'", "'csv_settings'",
		"'unique_key'", "'time_column'", "'description'", "'retention'", "'dialect'",
		"'columns'", "'audits'", "'grain'", "'owner'", "'stamp'", "'tags'",
		"'cron'", "'name'", "'kind'", "'path'", "'query'", "'('", "')'", "'['",
		"']'", "'{'", "'}'", "','", "'='", "':'", "'.'", "", "", "", "", "",
		"", "", "' '",
	}
	staticData.SymbolicNames = []string{
		"", "BLOCK_COMMENT", "LINE_COMMENT", "WS", "MODEL", "AUDIT", "KIND_INCREMENTAL_BY_TIME_RANGE",
		"KIND_INCREMENTAL_BY_UNIQUE_KEY", "KIND_INCREMENTAL_BY_PARTITION", "KIND_SCD_TYPE_2",
		"KIND_EXTERNAL", "KIND_EMBEDDED", "KIND_FULL", "KIND_VIEW", "KIND_SEED",
		"TRUE", "FALSE", "PROP_COLUMN_DESCRIPTIONS", "PROP_PHYSICAL_PROPERTIES",
		"PROP_ON_VIRTUAL_UPDATE", "PROP_INVALIDATE_HARD_DELETES", "PROP_DISABLE_RESTATEMENT",
		"PROP_PRE_STATEMENTS", "PROP_POST_STATEMENTS", "PROP_PARTITIONED_BY",
		"PROP_CLUSTERED_BY", "PROP_STORAGE_FORMAT", "PROP_TABLE_FORMAT", "PROP_CSV_SETTINGS",
		"PROP_UNIQUE_KEY", "PROP_TIME_COLUMN", "PROP_DESCRIPTION", "PROP_RETENTION",
		"PROP_DIALECT", "PROP_COLUMNS", "PROP_AUDITS", "PROP_GRAIN", "PROP_OWNER",
		"PROP_STAMP", "PROP_TAGS", "PROP_CRON", "PROP_NAME", "PROP_KIND", "PROP_PATH",
		"PROP_QUERY", "LPAREN", "RPAREN", "LBRACKET", "RBRACKET", "LBRACE",
		"RBRACE", "COMMA", "EQUALS", "COLON", "DOT", "SEMI", "MACRO_BRACED",
		"MACRO_IDENT", "FLOAT_LITERAL", "INT_LITERAL", "STRING_LITERAL", "IDENTIFIER",
		"SQL_LINE", "SQL_BODY_SEMI",
	}
	staticData.RuleNames = []string{
		"sqlmeshFile", "topStatement", "modelDef", "modelProp", "kindValue",
		"auditDef", "auditProp", "auditArray", "auditEntry", "auditEntryProp",
		"propValue", "literal", "qualifiedIdent", "arrayValue", "dictValue",
		"dictEntry", "macroRef", "sqlBody",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 63, 298, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 1, 0, 5, 0, 38, 8, 0, 10, 0, 12, 0, 41, 9,
		0, 1, 0, 1, 0, 1, 1, 1, 1, 3, 1, 47, 8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		5, 2, 54, 8, 2, 10, 2, 12, 2, 57, 9, 2, 1, 2, 3, 2, 60, 8, 2, 3, 2, 62,
		8, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 3, 3, 148, 8, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 4, 3, 4, 159, 8, 4, 1, 5, 1, 5, 3, 5, 163, 8, 5, 1, 5, 1, 5, 1, 5, 1,
		5, 5, 5, 169, 8, 5, 10, 5, 12, 5, 172, 9, 5, 1, 5, 3, 5, 175, 8, 5, 3,
		5, 177, 8, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3,
		6, 188, 8, 6, 1, 7, 1, 7, 1, 7, 1, 7, 5, 7, 194, 8, 7, 10, 7, 12, 7, 197,
		9, 7, 1, 7, 3, 7, 200, 8, 7, 3, 7, 202, 8, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1,
		8, 1, 8, 5, 8, 210, 8, 8, 10, 8, 12, 8, 213, 9, 8, 1, 8, 3, 8, 216, 8,
		8, 3, 8, 218, 8, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 3,
		9, 228, 8, 9, 1, 10, 1, 10, 1, 10, 1, 10, 3, 10, 234, 8, 10, 1, 11, 1,
		11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 3, 11, 243, 8, 11, 1, 12, 1, 12,
		1, 12, 4, 12, 248, 8, 12, 11, 12, 12, 12, 249, 1, 13, 1, 13, 1, 13, 1,
		13, 5, 13, 256, 8, 13, 10, 13, 12, 13, 259, 9, 13, 1, 13, 3, 13, 262, 8,
		13, 3, 13, 264, 8, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 5, 14,
		272, 8, 14, 10, 14, 12, 14, 275, 9, 14, 1, 14, 3, 14, 278, 8, 14, 3, 14,
		280, 8, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 3,
		16, 290, 8, 16, 1, 17, 5, 17, 293, 8, 17, 10, 17, 12, 17, 296, 9, 17, 1,
		17, 0, 0, 18, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30,
		32, 34, 0, 1, 1, 0, 60, 61, 348, 0, 39, 1, 0, 0, 0, 2, 46, 1, 0, 0, 0,
		4, 48, 1, 0, 0, 0, 6, 147, 1, 0, 0, 0, 8, 158, 1, 0, 0, 0, 10, 160, 1,
		0, 0, 0, 12, 187, 1, 0, 0, 0, 14, 189, 1, 0, 0, 0, 16, 205, 1, 0, 0, 0,
		18, 227, 1, 0, 0, 0, 20, 233, 1, 0, 0, 0, 22, 242, 1, 0, 0, 0, 24, 244,
		1, 0, 0, 0, 26, 251, 1, 0, 0, 0, 28, 267, 1, 0, 0, 0, 30, 283, 1, 0, 0,
		0, 32, 289, 1, 0, 0, 0, 34, 294, 1, 0, 0, 0, 36, 38, 3, 2, 1, 0, 37, 36,
		1, 0, 0, 0, 38, 41, 1, 0, 0, 0, 39, 37, 1, 0, 0, 0, 39, 40, 1, 0, 0, 0,
		40, 42, 1, 0, 0, 0, 41, 39, 1, 0, 0, 0, 42, 43, 5, 0, 0, 1, 43, 1, 1, 0,
		0, 0, 44, 47, 3, 4, 2, 0, 45, 47, 3, 10, 5, 0, 46, 44, 1, 0, 0, 0, 46,
		45, 1, 0, 0, 0, 47, 3, 1, 0, 0, 0, 48, 49, 5, 4, 0, 0, 49, 61, 5, 45, 0,
		0, 50, 55, 3, 6, 3, 0, 51, 52, 5, 51, 0, 0, 52, 54, 3, 6, 3, 0, 53, 51,
		1, 0, 0, 0, 54, 57, 1, 0, 0, 0, 55, 53, 1, 0, 0, 0, 55, 56, 1, 0, 0, 0,
		56, 59, 1, 0, 0, 0, 57, 55, 1, 0, 0, 0, 58, 60, 5, 51, 0, 0, 59, 58, 1,
		0, 0, 0, 59, 60, 1, 0, 0, 0, 60, 62, 1, 0, 0, 0, 61, 50, 1, 0, 0, 0, 61,
		62, 1, 0, 0, 0, 62, 63, 1, 0, 0, 0, 63, 64, 5, 46, 0, 0, 64, 65, 3, 34,
		17, 0, 65, 5, 1, 0, 0, 0, 66, 67, 5, 41, 0, 0, 67, 68, 5, 52, 0, 0, 68,
		148, 3, 20, 10, 0, 69, 70, 5, 42, 0, 0, 70, 71, 5, 52, 0, 0, 71, 148, 3,
		8, 4, 0, 72, 73, 5, 33, 0, 0, 73, 74, 5, 52, 0, 0, 74, 148, 3, 20, 10,
		0, 75, 76, 5, 37, 0, 0, 76, 77, 5, 52, 0, 0, 77, 148, 3, 20, 10, 0, 78,
		79, 5, 40, 0, 0, 79, 80, 5, 52, 0, 0, 80, 148, 3, 20, 10, 0, 81, 82, 5,
		36, 0, 0, 82, 83, 5, 52, 0, 0, 83, 148, 3, 26, 13, 0, 84, 85, 5, 24, 0,
		0, 85, 86, 5, 52, 0, 0, 86, 148, 3, 20, 10, 0, 87, 88, 5, 25, 0, 0, 88,
		89, 5, 52, 0, 0, 89, 148, 3, 26, 13, 0, 90, 91, 5, 26, 0, 0, 91, 92, 5,
		52, 0, 0, 92, 148, 3, 20, 10, 0, 93, 94, 5, 27, 0, 0, 94, 95, 5, 52, 0,
		0, 95, 148, 3, 20, 10, 0, 96, 97, 5, 32, 0, 0, 97, 98, 5, 52, 0, 0, 98,
		148, 3, 20, 10, 0, 99, 100, 5, 39, 0, 0, 100, 101, 5, 52, 0, 0, 101, 148,
		3, 26, 13, 0, 102, 103, 5, 31, 0, 0, 103, 104, 5, 52, 0, 0, 104, 148, 3,
		20, 10, 0, 105, 106, 5, 38, 0, 0, 106, 107, 5, 52, 0, 0, 107, 148, 3, 20,
		10, 0, 108, 109, 5, 34, 0, 0, 109, 110, 5, 52, 0, 0, 110, 148, 3, 28, 14,
		0, 111, 112, 5, 17, 0, 0, 112, 113, 5, 52, 0, 0, 113, 148, 3, 28, 14, 0,
		114, 115, 5, 18, 0, 0, 115, 116, 5, 52, 0, 0, 116, 148, 3, 28, 14, 0, 117,
		118, 5, 22, 0, 0, 118, 119, 5, 52, 0, 0, 119, 148, 3, 26, 13, 0, 120, 121,
		5, 23, 0, 0, 121, 122, 5, 52, 0, 0, 122, 148, 3, 26, 13, 0, 123, 124, 5,
		19, 0, 0, 124, 125, 5, 52, 0, 0, 125, 148, 3, 26, 13, 0, 126, 127, 5, 35,
		0, 0, 127, 128, 5, 52, 0, 0, 128, 148, 3, 14, 7, 0, 129, 130, 5, 30, 0,
		0, 130, 131, 5, 52, 0, 0, 131, 148, 3, 20, 10, 0, 132, 133, 5, 29, 0, 0,
		133, 134, 5, 52, 0, 0, 134, 148, 3, 20, 10, 0, 135, 136, 5, 20, 0, 0, 136,
		137, 5, 52, 0, 0, 137, 148, 3, 20, 10, 0, 138, 139, 5, 21, 0, 0, 139, 140,
		5, 52, 0, 0, 140, 148, 3, 20, 10, 0, 141, 142, 5, 43, 0, 0, 142, 143, 5,
		52, 0, 0, 143, 148, 3, 20, 10, 0, 144, 145, 5, 28, 0, 0, 145, 146, 5, 52,
		0, 0, 146, 148, 3, 28, 14, 0, 147, 66, 1, 0, 0, 0, 147, 69, 1, 0, 0, 0,
		147, 72, 1, 0, 0, 0, 147, 75, 1, 0, 0, 0, 147, 78, 1, 0, 0, 0, 147, 81,
		1, 0, 0, 0, 147, 84, 1, 0, 0, 0, 147, 87, 1, 0, 0, 0, 147, 90, 1, 0, 0,
		0, 147, 93, 1, 0, 0, 0, 147, 96, 1, 0, 0, 0, 147, 99, 1, 0, 0, 0, 147,
		102, 1, 0, 0, 0, 147, 105, 1, 0, 0, 0, 147, 108, 1, 0, 0, 0, 147, 111,
		1, 0, 0, 0, 147, 114, 1, 0, 0, 0, 147, 117, 1, 0, 0, 0, 147, 120, 1, 0,
		0, 0, 147, 123, 1, 0, 0, 0, 147, 126, 1, 0, 0, 0, 147, 129, 1, 0, 0, 0,
		147, 132, 1, 0, 0, 0, 147, 135, 1, 0, 0, 0, 147, 138, 1, 0, 0, 0, 147,
		141, 1, 0, 0, 0, 147, 144, 1, 0, 0, 0, 148, 7, 1, 0, 0, 0, 149, 159, 5,
		12, 0, 0, 150, 159, 5, 6, 0, 0, 151, 159, 5, 7, 0, 0, 152, 159, 5, 8, 0,
		0, 153, 159, 5, 13, 0, 0, 154, 159, 5, 14, 0, 0, 155, 159, 5, 10, 0, 0,
		156, 159, 5, 11, 0, 0, 157, 159, 5, 9, 0, 0, 158, 149, 1, 0, 0, 0, 158,
		150, 1, 0, 0, 0, 158, 151, 1, 0, 0, 0, 158, 152, 1, 0, 0, 0, 158, 153,
		1, 0, 0, 0, 158, 154, 1, 0, 0, 0, 158, 155, 1, 0, 0, 0, 158, 156, 1, 0,
		0, 0, 158, 157, 1, 0, 0, 0, 159, 9, 1, 0, 0, 0, 160, 162, 5, 5, 0, 0, 161,
		163, 5, 61, 0, 0, 162, 161, 1, 0, 0, 0, 162, 163, 1, 0, 0, 0, 163, 164,
		1, 0, 0, 0, 164, 176, 5, 45, 0, 0, 165, 170, 3, 12, 6, 0, 166, 167, 5,
		51, 0, 0, 167, 169, 3, 12, 6, 0, 168, 166, 1, 0, 0, 0, 169, 172, 1, 0,
		0, 0, 170, 168, 1, 0, 0, 0, 170, 171, 1, 0, 0, 0, 171, 174, 1, 0, 0, 0,
		172, 170, 1, 0, 0, 0, 173, 175, 5, 51, 0, 0, 174, 173, 1, 0, 0, 0, 174,
		175, 1, 0, 0, 0, 175, 177, 1, 0, 0, 0, 176, 165, 1, 0, 0, 0, 176, 177,
		1, 0, 0, 0, 177, 178, 1, 0, 0, 0, 178, 179, 5, 46, 0, 0, 179, 180, 3, 34,
		17, 0, 180, 11, 1, 0, 0, 0, 181, 182, 5, 44, 0, 0, 182, 183, 5, 52, 0,
		0, 183, 188, 3, 20, 10, 0, 184, 185, 5, 33, 0, 0, 185, 186, 5, 52, 0, 0,
		186, 188, 3, 20, 10, 0, 187, 181, 1, 0, 0, 0, 187, 184, 1, 0, 0, 0, 188,
		13, 1, 0, 0, 0, 189, 201, 5, 47, 0, 0, 190, 195, 3, 16, 8, 0, 191, 192,
		5, 51, 0, 0, 192, 194, 3, 16, 8, 0, 193, 191, 1, 0, 0, 0, 194, 197, 1,
		0, 0, 0, 195, 193, 1, 0, 0, 0, 195, 196, 1, 0, 0, 0, 196, 199, 1, 0, 0,
		0, 197, 195, 1, 0, 0, 0, 198, 200, 5, 51, 0, 0, 199, 198, 1, 0, 0, 0, 199,
		200, 1, 0, 0, 0, 200, 202, 1, 0, 0, 0, 201, 190, 1, 0, 0, 0, 201, 202,
		1, 0, 0, 0, 202, 203, 1, 0, 0, 0, 203, 204, 5, 48, 0, 0, 204, 15, 1, 0,
		0, 0, 205, 217, 5, 49, 0, 0, 206, 211, 3, 18, 9, 0, 207, 208, 5, 51, 0,
		0, 208, 210, 3, 18, 9, 0, 209, 207, 1, 0, 0, 0, 210, 213, 1, 0, 0, 0, 211,
		209, 1, 0, 0, 0, 211, 212, 1, 0, 0, 0, 212, 215, 1, 0, 0, 0, 213, 211,
		1, 0, 0, 0, 214, 216, 5, 51, 0, 0, 215, 214, 1, 0, 0, 0, 215, 216, 1, 0,
		0, 0, 216, 218, 1, 0, 0, 0, 217, 206, 1, 0, 0, 0, 217, 218, 1, 0, 0, 0,
		218, 219, 1, 0, 0, 0, 219, 220, 5, 50, 0, 0, 220, 17, 1, 0, 0, 0, 221,
		222, 5, 41, 0, 0, 222, 223, 5, 52, 0, 0, 223, 228, 3, 20, 10, 0, 224, 225,
		5, 44, 0, 0, 225, 226, 5, 52, 0, 0, 226, 228, 3, 20, 10, 0, 227, 221, 1,
		0, 0, 0, 227, 224, 1, 0, 0, 0, 228, 19, 1, 0, 0, 0, 229, 234, 3, 22, 11,
		0, 230, 234, 3, 26, 13, 0, 231, 234, 3, 28, 14, 0, 232, 234, 3, 32, 16,
		0, 233, 229, 1, 0, 0, 0, 233, 230, 1, 0, 0, 0, 233, 231, 1, 0, 0, 0, 233,
		232, 1, 0, 0, 0, 234, 21, 1, 0, 0, 0, 235, 243, 5, 60, 0, 0, 236, 243,
		5, 59, 0, 0, 237, 243, 5, 58, 0, 0, 238, 243, 5, 15, 0, 0, 239, 243, 5,
		16, 0, 0, 240, 243, 3, 24, 12, 0, 241, 243, 5, 61, 0, 0, 242, 235, 1, 0,
		0, 0, 242, 236, 1, 0, 0, 0, 242, 237, 1, 0, 0, 0, 242, 238, 1, 0, 0, 0,
		242, 239, 1, 0, 0, 0, 242, 240, 1, 0, 0, 0, 242, 241, 1, 0, 0, 0, 243,
		23, 1, 0, 0, 0, 244, 247, 5, 61, 0, 0, 245, 246, 5, 54, 0, 0, 246, 248,
		5, 61, 0, 0, 247, 245, 1, 0, 0, 0, 248, 249, 1, 0, 0, 0, 249, 247, 1, 0,
		0, 0, 249, 250, 1, 0, 0, 0, 250, 25, 1, 0, 0, 0, 251, 263, 5, 47, 0, 0,
		252, 257, 3, 20, 10, 0, 253, 254, 5, 51, 0, 0, 254, 256, 3, 20, 10, 0,
		255, 253, 1, 0, 0, 0, 256, 259, 1, 0, 0, 0, 257, 255, 1, 0, 0, 0, 257,
		258, 1, 0, 0, 0, 258, 261, 1, 0, 0, 0, 259, 257, 1, 0, 0, 0, 260, 262,
		5, 51, 0, 0, 261, 260, 1, 0, 0, 0, 261, 262, 1, 0, 0, 0, 262, 264, 1, 0,
		0, 0, 263, 252, 1, 0, 0, 0, 263, 264, 1, 0, 0, 0, 264, 265, 1, 0, 0, 0,
		265, 266, 5, 48, 0, 0, 266, 27, 1, 0, 0, 0, 267, 279, 5, 49, 0, 0, 268,
		273, 3, 30, 15, 0, 269, 270, 5, 51, 0, 0, 270, 272, 3, 30, 15, 0, 271,
		269, 1, 0, 0, 0, 272, 275, 1, 0, 0, 0, 273, 271, 1, 0, 0, 0, 273, 274,
		1, 0, 0, 0, 274, 277, 1, 0, 0, 0, 275, 273, 1, 0, 0, 0, 276, 278, 5, 51,
		0, 0, 277, 276, 1, 0, 0, 0, 277, 278, 1, 0, 0, 0, 278, 280, 1, 0, 0, 0,
		279, 268, 1, 0, 0, 0, 279, 280, 1, 0, 0, 0, 280, 281, 1, 0, 0, 0, 281,
		282, 5, 50, 0, 0, 282, 29, 1, 0, 0, 0, 283, 284, 7, 0, 0, 0, 284, 285,
		5, 53, 0, 0, 285, 286, 3, 20, 10, 0, 286, 31, 1, 0, 0, 0, 287, 290, 5,
		56, 0, 0, 288, 290, 5, 57, 0, 0, 289, 287, 1, 0, 0, 0, 289, 288, 1, 0,
		0, 0, 290, 33, 1, 0, 0, 0, 291, 293, 5, 62, 0, 0, 292, 291, 1, 0, 0, 0,
		293, 296, 1, 0, 0, 0, 294, 292, 1, 0, 0, 0, 294, 295, 1, 0, 0, 0, 295,
		35, 1, 0, 0, 0, 296, 294, 1, 0, 0, 0, 30, 39, 46, 55, 59, 61, 147, 158,
		162, 170, 174, 176, 187, 195, 199, 201, 211, 215, 217, 227, 233, 242, 249,
		257, 261, 263, 273, 277, 279, 289, 294,
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

// SQLMeshInit initializes any static state used to implement SQLMesh. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewSQLMesh(). You can call this function if you wish to initialize the static state ahead
// of time.
func SQLMeshInit() {
	staticData := &SQLMeshParserStaticData
	staticData.once.Do(sqlmeshParserInit)
}

// NewSQLMesh produces a new parser instance for the optional input antlr.TokenStream.
func NewSQLMesh(input antlr.TokenStream) *SQLMesh {
	SQLMeshInit()
	this := new(SQLMesh)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &SQLMeshParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "SQLMesh.g4"

	return this
}

// SQLMesh tokens.
const (
	SQLMeshEOF                            = antlr.TokenEOF
	SQLMeshBLOCK_COMMENT                  = 1
	SQLMeshLINE_COMMENT                   = 2
	SQLMeshWS                             = 3
	SQLMeshMODEL                          = 4
	SQLMeshAUDIT                          = 5
	SQLMeshKIND_INCREMENTAL_BY_TIME_RANGE = 6
	SQLMeshKIND_INCREMENTAL_BY_UNIQUE_KEY = 7
	SQLMeshKIND_INCREMENTAL_BY_PARTITION  = 8
	SQLMeshKIND_SCD_TYPE_2                = 9
	SQLMeshKIND_EXTERNAL                  = 10
	SQLMeshKIND_EMBEDDED                  = 11
	SQLMeshKIND_FULL                      = 12
	SQLMeshKIND_VIEW                      = 13
	SQLMeshKIND_SEED                      = 14
	SQLMeshTRUE                           = 15
	SQLMeshFALSE                          = 16
	SQLMeshPROP_COLUMN_DESCRIPTIONS       = 17
	SQLMeshPROP_PHYSICAL_PROPERTIES       = 18
	SQLMeshPROP_ON_VIRTUAL_UPDATE         = 19
	SQLMeshPROP_INVALIDATE_HARD_DELETES   = 20
	SQLMeshPROP_DISABLE_RESTATEMENT       = 21
	SQLMeshPROP_PRE_STATEMENTS            = 22
	SQLMeshPROP_POST_STATEMENTS           = 23
	SQLMeshPROP_PARTITIONED_BY            = 24
	SQLMeshPROP_CLUSTERED_BY              = 25
	SQLMeshPROP_STORAGE_FORMAT            = 26
	SQLMeshPROP_TABLE_FORMAT              = 27
	SQLMeshPROP_CSV_SETTINGS              = 28
	SQLMeshPROP_UNIQUE_KEY                = 29
	SQLMeshPROP_TIME_COLUMN               = 30
	SQLMeshPROP_DESCRIPTION               = 31
	SQLMeshPROP_RETENTION                 = 32
	SQLMeshPROP_DIALECT                   = 33
	SQLMeshPROP_COLUMNS                   = 34
	SQLMeshPROP_AUDITS                    = 35
	SQLMeshPROP_GRAIN                     = 36
	SQLMeshPROP_OWNER                     = 37
	SQLMeshPROP_STAMP                     = 38
	SQLMeshPROP_TAGS                      = 39
	SQLMeshPROP_CRON                      = 40
	SQLMeshPROP_NAME                      = 41
	SQLMeshPROP_KIND                      = 42
	SQLMeshPROP_PATH                      = 43
	SQLMeshPROP_QUERY                     = 44
	SQLMeshLPAREN                         = 45
	SQLMeshRPAREN                         = 46
	SQLMeshLBRACKET                       = 47
	SQLMeshRBRACKET                       = 48
	SQLMeshLBRACE                         = 49
	SQLMeshRBRACE                         = 50
	SQLMeshCOMMA                          = 51
	SQLMeshEQUALS                         = 52
	SQLMeshCOLON                          = 53
	SQLMeshDOT                            = 54
	SQLMeshSEMI                           = 55
	SQLMeshMACRO_BRACED                   = 56
	SQLMeshMACRO_IDENT                    = 57
	SQLMeshFLOAT_LITERAL                  = 58
	SQLMeshINT_LITERAL                    = 59
	SQLMeshSTRING_LITERAL                 = 60
	SQLMeshIDENTIFIER                     = 61
	SQLMeshSQL_LINE                       = 62
	SQLMeshSQL_BODY_SEMI                  = 63
)

// SQLMesh rules.
const (
	SQLMeshRULE_sqlmeshFile    = 0
	SQLMeshRULE_topStatement   = 1
	SQLMeshRULE_modelDef       = 2
	SQLMeshRULE_modelProp      = 3
	SQLMeshRULE_kindValue      = 4
	SQLMeshRULE_auditDef       = 5
	SQLMeshRULE_auditProp      = 6
	SQLMeshRULE_auditArray     = 7
	SQLMeshRULE_auditEntry     = 8
	SQLMeshRULE_auditEntryProp = 9
	SQLMeshRULE_propValue      = 10
	SQLMeshRULE_literal        = 11
	SQLMeshRULE_qualifiedIdent = 12
	SQLMeshRULE_arrayValue     = 13
	SQLMeshRULE_dictValue      = 14
	SQLMeshRULE_dictEntry      = 15
	SQLMeshRULE_macroRef       = 16
	SQLMeshRULE_sqlBody        = 17
)

// ISqlmeshFileContext is an interface to support dynamic dispatch.
type ISqlmeshFileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllTopStatement() []ITopStatementContext
	TopStatement(i int) ITopStatementContext

	// IsSqlmeshFileContext differentiates from other interfaces.
	IsSqlmeshFileContext()
}

type SqlmeshFileContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySqlmeshFileContext() *SqlmeshFileContext {
	var p = new(SqlmeshFileContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLMeshRULE_sqlmeshFile
	return p
}

func InitEmptySqlmeshFileContext(p *SqlmeshFileContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLMeshRULE_sqlmeshFile
}

func (*SqlmeshFileContext) IsSqlmeshFileContext() {}

func NewSqlmeshFileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SqlmeshFileContext {
	var p = new(SqlmeshFileContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLMeshRULE_sqlmeshFile

	return p
}

func (s *SqlmeshFileContext) GetParser() antlr.Parser { return s.parser }

func (s *SqlmeshFileContext) EOF() antlr.TerminalNode {
	return s.GetToken(SQLMeshEOF, 0)
}

func (s *SqlmeshFileContext) AllTopStatement() []ITopStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITopStatementContext); ok {
			len++
		}
	}

	tst := make([]ITopStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITopStatementContext); ok {
			tst[i] = t.(ITopStatementContext)
			i++
		}
	}

	return tst
}

func (s *SqlmeshFileContext) TopStatement(i int) ITopStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITopStatementContext); ok {
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

	return t.(ITopStatementContext)
}

func (s *SqlmeshFileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SqlmeshFileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SqlmeshFileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.EnterSqlmeshFile(s)
	}
}

func (s *SqlmeshFileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.ExitSqlmeshFile(s)
	}
}

func (s *SqlmeshFileContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLMeshVisitor:
		return t.VisitSqlmeshFile(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLMesh) SqlmeshFile() (localctx ISqlmeshFileContext) {
	localctx = NewSqlmeshFileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SQLMeshRULE_sqlmeshFile)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(39)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SQLMeshMODEL || _la == SQLMeshAUDIT {
		{
			p.SetState(36)
			p.TopStatement()
		}

		p.SetState(41)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(42)
		p.Match(SQLMeshEOF)
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

// ITopStatementContext is an interface to support dynamic dispatch.
type ITopStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ModelDef() IModelDefContext
	AuditDef() IAuditDefContext

	// IsTopStatementContext differentiates from other interfaces.
	IsTopStatementContext()
}

type TopStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTopStatementContext() *TopStatementContext {
	var p = new(TopStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLMeshRULE_topStatement
	return p
}

func InitEmptyTopStatementContext(p *TopStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLMeshRULE_topStatement
}

func (*TopStatementContext) IsTopStatementContext() {}

func NewTopStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TopStatementContext {
	var p = new(TopStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLMeshRULE_topStatement

	return p
}

func (s *TopStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *TopStatementContext) ModelDef() IModelDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IModelDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IModelDefContext)
}

func (s *TopStatementContext) AuditDef() IAuditDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAuditDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAuditDefContext)
}

func (s *TopStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TopStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TopStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.EnterTopStatement(s)
	}
}

func (s *TopStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.ExitTopStatement(s)
	}
}

func (s *TopStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLMeshVisitor:
		return t.VisitTopStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLMesh) TopStatement() (localctx ITopStatementContext) {
	localctx = NewTopStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SQLMeshRULE_topStatement)
	p.SetState(46)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SQLMeshMODEL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(44)
			p.ModelDef()
		}

	case SQLMeshAUDIT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(45)
			p.AuditDef()
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

// IModelDefContext is an interface to support dynamic dispatch.
type IModelDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MODEL() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	SqlBody() ISqlBodyContext
	AllModelProp() []IModelPropContext
	ModelProp(i int) IModelPropContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsModelDefContext differentiates from other interfaces.
	IsModelDefContext()
}

type ModelDefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModelDefContext() *ModelDefContext {
	var p = new(ModelDefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLMeshRULE_modelDef
	return p
}

func InitEmptyModelDefContext(p *ModelDefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLMeshRULE_modelDef
}

func (*ModelDefContext) IsModelDefContext() {}

func NewModelDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ModelDefContext {
	var p = new(ModelDefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLMeshRULE_modelDef

	return p
}

func (s *ModelDefContext) GetParser() antlr.Parser { return s.parser }

func (s *ModelDefContext) MODEL() antlr.TerminalNode {
	return s.GetToken(SQLMeshMODEL, 0)
}

func (s *ModelDefContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(SQLMeshLPAREN, 0)
}

func (s *ModelDefContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(SQLMeshRPAREN, 0)
}

func (s *ModelDefContext) SqlBody() ISqlBodyContext {
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

func (s *ModelDefContext) AllModelProp() []IModelPropContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IModelPropContext); ok {
			len++
		}
	}

	tst := make([]IModelPropContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IModelPropContext); ok {
			tst[i] = t.(IModelPropContext)
			i++
		}
	}

	return tst
}

func (s *ModelDefContext) ModelProp(i int) IModelPropContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IModelPropContext); ok {
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

	return t.(IModelPropContext)
}

func (s *ModelDefContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SQLMeshCOMMA)
}

func (s *ModelDefContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SQLMeshCOMMA, i)
}

func (s *ModelDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModelDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ModelDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.EnterModelDef(s)
	}
}

func (s *ModelDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.ExitModelDef(s)
	}
}

func (s *ModelDefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLMeshVisitor:
		return t.VisitModelDef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLMesh) ModelDef() (localctx IModelDefContext) {
	localctx = NewModelDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, SQLMeshRULE_modelDef)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(48)
		p.Match(SQLMeshMODEL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(49)
		p.Match(SQLMeshLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(61)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&17592185913344) != 0 {
		{
			p.SetState(50)
			p.ModelProp()
		}
		p.SetState(55)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(51)
					p.Match(SQLMeshCOMMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(52)
					p.ModelProp()
				}

			}
			p.SetState(57)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		p.SetState(59)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SQLMeshCOMMA {
			{
				p.SetState(58)
				p.Match(SQLMeshCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	}
	{
		p.SetState(63)
		p.Match(SQLMeshRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(64)
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

// IModelPropContext is an interface to support dynamic dispatch.
type IModelPropContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsModelPropContext differentiates from other interfaces.
	IsModelPropContext()
}

type ModelPropContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModelPropContext() *ModelPropContext {
	var p = new(ModelPropContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLMeshRULE_modelProp
	return p
}

func InitEmptyModelPropContext(p *ModelPropContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLMeshRULE_modelProp
}

func (*ModelPropContext) IsModelPropContext() {}

func NewModelPropContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ModelPropContext {
	var p = new(ModelPropContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLMeshRULE_modelProp

	return p
}

func (s *ModelPropContext) GetParser() antlr.Parser { return s.parser }

func (s *ModelPropContext) CopyAll(ctx *ModelPropContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ModelPropContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModelPropContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type DisableRestatementPropContext struct {
	ModelPropContext
}

func NewDisableRestatementPropContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DisableRestatementPropContext {
	var p = new(DisableRestatementPropContext)

	InitEmptyModelPropContext(&p.ModelPropContext)
	p.parser = parser
	p.CopyAll(ctx.(*ModelPropContext))

	return p
}

func (s *DisableRestatementPropContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DisableRestatementPropContext) PROP_DISABLE_RESTATEMENT() antlr.TerminalNode {
	return s.GetToken(SQLMeshPROP_DISABLE_RESTATEMENT, 0)
}

func (s *DisableRestatementPropContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(SQLMeshEQUALS, 0)
}

func (s *DisableRestatementPropContext) PropValue() IPropValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPropValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPropValueContext)
}

func (s *DisableRestatementPropContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.EnterDisableRestatementProp(s)
	}
}

func (s *DisableRestatementPropContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.ExitDisableRestatementProp(s)
	}
}

func (s *DisableRestatementPropContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLMeshVisitor:
		return t.VisitDisableRestatementProp(s)

	default:
		return t.VisitChildren(s)
	}
}

type CsvSettingsPropContext struct {
	ModelPropContext
}

func NewCsvSettingsPropContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CsvSettingsPropContext {
	var p = new(CsvSettingsPropContext)

	InitEmptyModelPropContext(&p.ModelPropContext)
	p.parser = parser
	p.CopyAll(ctx.(*ModelPropContext))

	return p
}

func (s *CsvSettingsPropContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CsvSettingsPropContext) PROP_CSV_SETTINGS() antlr.TerminalNode {
	return s.GetToken(SQLMeshPROP_CSV_SETTINGS, 0)
}

func (s *CsvSettingsPropContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(SQLMeshEQUALS, 0)
}

func (s *CsvSettingsPropContext) DictValue() IDictValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDictValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDictValueContext)
}

func (s *CsvSettingsPropContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.EnterCsvSettingsProp(s)
	}
}

func (s *CsvSettingsPropContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.ExitCsvSettingsProp(s)
	}
}

func (s *CsvSettingsPropContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLMeshVisitor:
		return t.VisitCsvSettingsProp(s)

	default:
		return t.VisitChildren(s)
	}
}

type PartitionedByPropContext struct {
	ModelPropContext
}

func NewPartitionedByPropContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PartitionedByPropContext {
	var p = new(PartitionedByPropContext)

	InitEmptyModelPropContext(&p.ModelPropContext)
	p.parser = parser
	p.CopyAll(ctx.(*ModelPropContext))

	return p
}

func (s *PartitionedByPropContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PartitionedByPropContext) PROP_PARTITIONED_BY() antlr.TerminalNode {
	return s.GetToken(SQLMeshPROP_PARTITIONED_BY, 0)
}

func (s *PartitionedByPropContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(SQLMeshEQUALS, 0)
}

func (s *PartitionedByPropContext) PropValue() IPropValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPropValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPropValueContext)
}

func (s *PartitionedByPropContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.EnterPartitionedByProp(s)
	}
}

func (s *PartitionedByPropContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.ExitPartitionedByProp(s)
	}
}

func (s *PartitionedByPropContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLMeshVisitor:
		return t.VisitPartitionedByProp(s)

	default:
		return t.VisitChildren(s)
	}
}

type TableFormatPropContext struct {
	ModelPropContext
}

func NewTableFormatPropContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TableFormatPropContext {
	var p = new(TableFormatPropContext)

	InitEmptyModelPropContext(&p.ModelPropContext)
	p.parser = parser
	p.CopyAll(ctx.(*ModelPropContext))

	return p
}

func (s *TableFormatPropContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TableFormatPropContext) PROP_TABLE_FORMAT() antlr.TerminalNode {
	return s.GetToken(SQLMeshPROP_TABLE_FORMAT, 0)
}

func (s *TableFormatPropContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(SQLMeshEQUALS, 0)
}

func (s *TableFormatPropContext) PropValue() IPropValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPropValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPropValueContext)
}

func (s *TableFormatPropContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.EnterTableFormatProp(s)
	}
}

func (s *TableFormatPropContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.ExitTableFormatProp(s)
	}
}

func (s *TableFormatPropContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLMeshVisitor:
		return t.VisitTableFormatProp(s)

	default:
		return t.VisitChildren(s)
	}
}

type AuditsPropContext struct {
	ModelPropContext
}

func NewAuditsPropContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AuditsPropContext {
	var p = new(AuditsPropContext)

	InitEmptyModelPropContext(&p.ModelPropContext)
	p.parser = parser
	p.CopyAll(ctx.(*ModelPropContext))

	return p
}

func (s *AuditsPropContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AuditsPropContext) PROP_AUDITS() antlr.TerminalNode {
	return s.GetToken(SQLMeshPROP_AUDITS, 0)
}

func (s *AuditsPropContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(SQLMeshEQUALS, 0)
}

func (s *AuditsPropContext) AuditArray() IAuditArrayContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAuditArrayContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAuditArrayContext)
}

func (s *AuditsPropContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.EnterAuditsProp(s)
	}
}

func (s *AuditsPropContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.ExitAuditsProp(s)
	}
}

func (s *AuditsPropContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLMeshVisitor:
		return t.VisitAuditsProp(s)

	default:
		return t.VisitChildren(s)
	}
}

type StorageFormatPropContext struct {
	ModelPropContext
}

func NewStorageFormatPropContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StorageFormatPropContext {
	var p = new(StorageFormatPropContext)

	InitEmptyModelPropContext(&p.ModelPropContext)
	p.parser = parser
	p.CopyAll(ctx.(*ModelPropContext))

	return p
}

func (s *StorageFormatPropContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StorageFormatPropContext) PROP_STORAGE_FORMAT() antlr.TerminalNode {
	return s.GetToken(SQLMeshPROP_STORAGE_FORMAT, 0)
}

func (s *StorageFormatPropContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(SQLMeshEQUALS, 0)
}

func (s *StorageFormatPropContext) PropValue() IPropValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPropValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPropValueContext)
}

func (s *StorageFormatPropContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.EnterStorageFormatProp(s)
	}
}

func (s *StorageFormatPropContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.ExitStorageFormatProp(s)
	}
}

func (s *StorageFormatPropContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLMeshVisitor:
		return t.VisitStorageFormatProp(s)

	default:
		return t.VisitChildren(s)
	}
}

type GrainPropContext struct {
	ModelPropContext
}

func NewGrainPropContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GrainPropContext {
	var p = new(GrainPropContext)

	InitEmptyModelPropContext(&p.ModelPropContext)
	p.parser = parser
	p.CopyAll(ctx.(*ModelPropContext))

	return p
}

func (s *GrainPropContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GrainPropContext) PROP_GRAIN() antlr.TerminalNode {
	return s.GetToken(SQLMeshPROP_GRAIN, 0)
}

func (s *GrainPropContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(SQLMeshEQUALS, 0)
}

func (s *GrainPropContext) ArrayValue() IArrayValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayValueContext)
}

func (s *GrainPropContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.EnterGrainProp(s)
	}
}

func (s *GrainPropContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.ExitGrainProp(s)
	}
}

func (s *GrainPropContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLMeshVisitor:
		return t.VisitGrainProp(s)

	default:
		return t.VisitChildren(s)
	}
}

type StampPropContext struct {
	ModelPropContext
}

func NewStampPropContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StampPropContext {
	var p = new(StampPropContext)

	InitEmptyModelPropContext(&p.ModelPropContext)
	p.parser = parser
	p.CopyAll(ctx.(*ModelPropContext))

	return p
}

func (s *StampPropContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StampPropContext) PROP_STAMP() antlr.TerminalNode {
	return s.GetToken(SQLMeshPROP_STAMP, 0)
}

func (s *StampPropContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(SQLMeshEQUALS, 0)
}

func (s *StampPropContext) PropValue() IPropValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPropValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPropValueContext)
}

func (s *StampPropContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.EnterStampProp(s)
	}
}

func (s *StampPropContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.ExitStampProp(s)
	}
}

func (s *StampPropContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLMeshVisitor:
		return t.VisitStampProp(s)

	default:
		return t.VisitChildren(s)
	}
}

type PostStatementsPropContext struct {
	ModelPropContext
}

func NewPostStatementsPropContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PostStatementsPropContext {
	var p = new(PostStatementsPropContext)

	InitEmptyModelPropContext(&p.ModelPropContext)
	p.parser = parser
	p.CopyAll(ctx.(*ModelPropContext))

	return p
}

func (s *PostStatementsPropContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PostStatementsPropContext) PROP_POST_STATEMENTS() antlr.TerminalNode {
	return s.GetToken(SQLMeshPROP_POST_STATEMENTS, 0)
}

func (s *PostStatementsPropContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(SQLMeshEQUALS, 0)
}

func (s *PostStatementsPropContext) ArrayValue() IArrayValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayValueContext)
}

func (s *PostStatementsPropContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.EnterPostStatementsProp(s)
	}
}

func (s *PostStatementsPropContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.ExitPostStatementsProp(s)
	}
}

func (s *PostStatementsPropContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLMeshVisitor:
		return t.VisitPostStatementsProp(s)

	default:
		return t.VisitChildren(s)
	}
}

type DialectPropContext struct {
	ModelPropContext
}

func NewDialectPropContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DialectPropContext {
	var p = new(DialectPropContext)

	InitEmptyModelPropContext(&p.ModelPropContext)
	p.parser = parser
	p.CopyAll(ctx.(*ModelPropContext))

	return p
}

func (s *DialectPropContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DialectPropContext) PROP_DIALECT() antlr.TerminalNode {
	return s.GetToken(SQLMeshPROP_DIALECT, 0)
}

func (s *DialectPropContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(SQLMeshEQUALS, 0)
}

func (s *DialectPropContext) PropValue() IPropValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPropValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPropValueContext)
}

func (s *DialectPropContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.EnterDialectProp(s)
	}
}

func (s *DialectPropContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.ExitDialectProp(s)
	}
}

func (s *DialectPropContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLMeshVisitor:
		return t.VisitDialectProp(s)

	default:
		return t.VisitChildren(s)
	}
}

type OnVirtualUpdatePropContext struct {
	ModelPropContext
}

func NewOnVirtualUpdatePropContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OnVirtualUpdatePropContext {
	var p = new(OnVirtualUpdatePropContext)

	InitEmptyModelPropContext(&p.ModelPropContext)
	p.parser = parser
	p.CopyAll(ctx.(*ModelPropContext))

	return p
}

func (s *OnVirtualUpdatePropContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OnVirtualUpdatePropContext) PROP_ON_VIRTUAL_UPDATE() antlr.TerminalNode {
	return s.GetToken(SQLMeshPROP_ON_VIRTUAL_UPDATE, 0)
}

func (s *OnVirtualUpdatePropContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(SQLMeshEQUALS, 0)
}

func (s *OnVirtualUpdatePropContext) ArrayValue() IArrayValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayValueContext)
}

func (s *OnVirtualUpdatePropContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.EnterOnVirtualUpdateProp(s)
	}
}

func (s *OnVirtualUpdatePropContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.ExitOnVirtualUpdateProp(s)
	}
}

func (s *OnVirtualUpdatePropContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLMeshVisitor:
		return t.VisitOnVirtualUpdateProp(s)

	default:
		return t.VisitChildren(s)
	}
}

type InvalidateHardDeletesPropContext struct {
	ModelPropContext
}

func NewInvalidateHardDeletesPropContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InvalidateHardDeletesPropContext {
	var p = new(InvalidateHardDeletesPropContext)

	InitEmptyModelPropContext(&p.ModelPropContext)
	p.parser = parser
	p.CopyAll(ctx.(*ModelPropContext))

	return p
}

func (s *InvalidateHardDeletesPropContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InvalidateHardDeletesPropContext) PROP_INVALIDATE_HARD_DELETES() antlr.TerminalNode {
	return s.GetToken(SQLMeshPROP_INVALIDATE_HARD_DELETES, 0)
}

func (s *InvalidateHardDeletesPropContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(SQLMeshEQUALS, 0)
}

func (s *InvalidateHardDeletesPropContext) PropValue() IPropValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPropValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPropValueContext)
}

func (s *InvalidateHardDeletesPropContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.EnterInvalidateHardDeletesProp(s)
	}
}

func (s *InvalidateHardDeletesPropContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.ExitInvalidateHardDeletesProp(s)
	}
}

func (s *InvalidateHardDeletesPropContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLMeshVisitor:
		return t.VisitInvalidateHardDeletesProp(s)

	default:
		return t.VisitChildren(s)
	}
}

type TagsPropContext struct {
	ModelPropContext
}

func NewTagsPropContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TagsPropContext {
	var p = new(TagsPropContext)

	InitEmptyModelPropContext(&p.ModelPropContext)
	p.parser = parser
	p.CopyAll(ctx.(*ModelPropContext))

	return p
}

func (s *TagsPropContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TagsPropContext) PROP_TAGS() antlr.TerminalNode {
	return s.GetToken(SQLMeshPROP_TAGS, 0)
}

func (s *TagsPropContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(SQLMeshEQUALS, 0)
}

func (s *TagsPropContext) ArrayValue() IArrayValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayValueContext)
}

func (s *TagsPropContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.EnterTagsProp(s)
	}
}

func (s *TagsPropContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.ExitTagsProp(s)
	}
}

func (s *TagsPropContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLMeshVisitor:
		return t.VisitTagsProp(s)

	default:
		return t.VisitChildren(s)
	}
}

type ColumnDescriptionsPropContext struct {
	ModelPropContext
}

func NewColumnDescriptionsPropContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ColumnDescriptionsPropContext {
	var p = new(ColumnDescriptionsPropContext)

	InitEmptyModelPropContext(&p.ModelPropContext)
	p.parser = parser
	p.CopyAll(ctx.(*ModelPropContext))

	return p
}

func (s *ColumnDescriptionsPropContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ColumnDescriptionsPropContext) PROP_COLUMN_DESCRIPTIONS() antlr.TerminalNode {
	return s.GetToken(SQLMeshPROP_COLUMN_DESCRIPTIONS, 0)
}

func (s *ColumnDescriptionsPropContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(SQLMeshEQUALS, 0)
}

func (s *ColumnDescriptionsPropContext) DictValue() IDictValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDictValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDictValueContext)
}

func (s *ColumnDescriptionsPropContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.EnterColumnDescriptionsProp(s)
	}
}

func (s *ColumnDescriptionsPropContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.ExitColumnDescriptionsProp(s)
	}
}

func (s *ColumnDescriptionsPropContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLMeshVisitor:
		return t.VisitColumnDescriptionsProp(s)

	default:
		return t.VisitChildren(s)
	}
}

type RetentionPropContext struct {
	ModelPropContext
}

func NewRetentionPropContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RetentionPropContext {
	var p = new(RetentionPropContext)

	InitEmptyModelPropContext(&p.ModelPropContext)
	p.parser = parser
	p.CopyAll(ctx.(*ModelPropContext))

	return p
}

func (s *RetentionPropContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RetentionPropContext) PROP_RETENTION() antlr.TerminalNode {
	return s.GetToken(SQLMeshPROP_RETENTION, 0)
}

func (s *RetentionPropContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(SQLMeshEQUALS, 0)
}

func (s *RetentionPropContext) PropValue() IPropValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPropValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPropValueContext)
}

func (s *RetentionPropContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.EnterRetentionProp(s)
	}
}

func (s *RetentionPropContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.ExitRetentionProp(s)
	}
}

func (s *RetentionPropContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLMeshVisitor:
		return t.VisitRetentionProp(s)

	default:
		return t.VisitChildren(s)
	}
}

type ColumnsPropContext struct {
	ModelPropContext
}

func NewColumnsPropContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ColumnsPropContext {
	var p = new(ColumnsPropContext)

	InitEmptyModelPropContext(&p.ModelPropContext)
	p.parser = parser
	p.CopyAll(ctx.(*ModelPropContext))

	return p
}

func (s *ColumnsPropContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ColumnsPropContext) PROP_COLUMNS() antlr.TerminalNode {
	return s.GetToken(SQLMeshPROP_COLUMNS, 0)
}

func (s *ColumnsPropContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(SQLMeshEQUALS, 0)
}

func (s *ColumnsPropContext) DictValue() IDictValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDictValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDictValueContext)
}

func (s *ColumnsPropContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.EnterColumnsProp(s)
	}
}

func (s *ColumnsPropContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.ExitColumnsProp(s)
	}
}

func (s *ColumnsPropContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLMeshVisitor:
		return t.VisitColumnsProp(s)

	default:
		return t.VisitChildren(s)
	}
}

type CronPropContext struct {
	ModelPropContext
}

func NewCronPropContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CronPropContext {
	var p = new(CronPropContext)

	InitEmptyModelPropContext(&p.ModelPropContext)
	p.parser = parser
	p.CopyAll(ctx.(*ModelPropContext))

	return p
}

func (s *CronPropContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CronPropContext) PROP_CRON() antlr.TerminalNode {
	return s.GetToken(SQLMeshPROP_CRON, 0)
}

func (s *CronPropContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(SQLMeshEQUALS, 0)
}

func (s *CronPropContext) PropValue() IPropValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPropValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPropValueContext)
}

func (s *CronPropContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.EnterCronProp(s)
	}
}

func (s *CronPropContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.ExitCronProp(s)
	}
}

func (s *CronPropContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLMeshVisitor:
		return t.VisitCronProp(s)

	default:
		return t.VisitChildren(s)
	}
}

type PhysicalPropertiesPropContext struct {
	ModelPropContext
}

func NewPhysicalPropertiesPropContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PhysicalPropertiesPropContext {
	var p = new(PhysicalPropertiesPropContext)

	InitEmptyModelPropContext(&p.ModelPropContext)
	p.parser = parser
	p.CopyAll(ctx.(*ModelPropContext))

	return p
}

func (s *PhysicalPropertiesPropContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PhysicalPropertiesPropContext) PROP_PHYSICAL_PROPERTIES() antlr.TerminalNode {
	return s.GetToken(SQLMeshPROP_PHYSICAL_PROPERTIES, 0)
}

func (s *PhysicalPropertiesPropContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(SQLMeshEQUALS, 0)
}

func (s *PhysicalPropertiesPropContext) DictValue() IDictValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDictValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDictValueContext)
}

func (s *PhysicalPropertiesPropContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.EnterPhysicalPropertiesProp(s)
	}
}

func (s *PhysicalPropertiesPropContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.ExitPhysicalPropertiesProp(s)
	}
}

func (s *PhysicalPropertiesPropContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLMeshVisitor:
		return t.VisitPhysicalPropertiesProp(s)

	default:
		return t.VisitChildren(s)
	}
}

type TimeColumnPropContext struct {
	ModelPropContext
}

func NewTimeColumnPropContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TimeColumnPropContext {
	var p = new(TimeColumnPropContext)

	InitEmptyModelPropContext(&p.ModelPropContext)
	p.parser = parser
	p.CopyAll(ctx.(*ModelPropContext))

	return p
}

func (s *TimeColumnPropContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TimeColumnPropContext) PROP_TIME_COLUMN() antlr.TerminalNode {
	return s.GetToken(SQLMeshPROP_TIME_COLUMN, 0)
}

func (s *TimeColumnPropContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(SQLMeshEQUALS, 0)
}

func (s *TimeColumnPropContext) PropValue() IPropValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPropValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPropValueContext)
}

func (s *TimeColumnPropContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.EnterTimeColumnProp(s)
	}
}

func (s *TimeColumnPropContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.ExitTimeColumnProp(s)
	}
}

func (s *TimeColumnPropContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLMeshVisitor:
		return t.VisitTimeColumnProp(s)

	default:
		return t.VisitChildren(s)
	}
}

type ClusteredByPropContext struct {
	ModelPropContext
}

func NewClusteredByPropContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ClusteredByPropContext {
	var p = new(ClusteredByPropContext)

	InitEmptyModelPropContext(&p.ModelPropContext)
	p.parser = parser
	p.CopyAll(ctx.(*ModelPropContext))

	return p
}

func (s *ClusteredByPropContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClusteredByPropContext) PROP_CLUSTERED_BY() antlr.TerminalNode {
	return s.GetToken(SQLMeshPROP_CLUSTERED_BY, 0)
}

func (s *ClusteredByPropContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(SQLMeshEQUALS, 0)
}

func (s *ClusteredByPropContext) ArrayValue() IArrayValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayValueContext)
}

func (s *ClusteredByPropContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.EnterClusteredByProp(s)
	}
}

func (s *ClusteredByPropContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.ExitClusteredByProp(s)
	}
}

func (s *ClusteredByPropContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLMeshVisitor:
		return t.VisitClusteredByProp(s)

	default:
		return t.VisitChildren(s)
	}
}

type PreStatementsPropContext struct {
	ModelPropContext
}

func NewPreStatementsPropContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PreStatementsPropContext {
	var p = new(PreStatementsPropContext)

	InitEmptyModelPropContext(&p.ModelPropContext)
	p.parser = parser
	p.CopyAll(ctx.(*ModelPropContext))

	return p
}

func (s *PreStatementsPropContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PreStatementsPropContext) PROP_PRE_STATEMENTS() antlr.TerminalNode {
	return s.GetToken(SQLMeshPROP_PRE_STATEMENTS, 0)
}

func (s *PreStatementsPropContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(SQLMeshEQUALS, 0)
}

func (s *PreStatementsPropContext) ArrayValue() IArrayValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayValueContext)
}

func (s *PreStatementsPropContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.EnterPreStatementsProp(s)
	}
}

func (s *PreStatementsPropContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.ExitPreStatementsProp(s)
	}
}

func (s *PreStatementsPropContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLMeshVisitor:
		return t.VisitPreStatementsProp(s)

	default:
		return t.VisitChildren(s)
	}
}

type UniqueKeyPropContext struct {
	ModelPropContext
}

func NewUniqueKeyPropContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UniqueKeyPropContext {
	var p = new(UniqueKeyPropContext)

	InitEmptyModelPropContext(&p.ModelPropContext)
	p.parser = parser
	p.CopyAll(ctx.(*ModelPropContext))

	return p
}

func (s *UniqueKeyPropContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UniqueKeyPropContext) PROP_UNIQUE_KEY() antlr.TerminalNode {
	return s.GetToken(SQLMeshPROP_UNIQUE_KEY, 0)
}

func (s *UniqueKeyPropContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(SQLMeshEQUALS, 0)
}

func (s *UniqueKeyPropContext) PropValue() IPropValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPropValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPropValueContext)
}

func (s *UniqueKeyPropContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.EnterUniqueKeyProp(s)
	}
}

func (s *UniqueKeyPropContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.ExitUniqueKeyProp(s)
	}
}

func (s *UniqueKeyPropContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLMeshVisitor:
		return t.VisitUniqueKeyProp(s)

	default:
		return t.VisitChildren(s)
	}
}

type PathPropContext struct {
	ModelPropContext
}

func NewPathPropContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PathPropContext {
	var p = new(PathPropContext)

	InitEmptyModelPropContext(&p.ModelPropContext)
	p.parser = parser
	p.CopyAll(ctx.(*ModelPropContext))

	return p
}

func (s *PathPropContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PathPropContext) PROP_PATH() antlr.TerminalNode {
	return s.GetToken(SQLMeshPROP_PATH, 0)
}

func (s *PathPropContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(SQLMeshEQUALS, 0)
}

func (s *PathPropContext) PropValue() IPropValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPropValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPropValueContext)
}

func (s *PathPropContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.EnterPathProp(s)
	}
}

func (s *PathPropContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.ExitPathProp(s)
	}
}

func (s *PathPropContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLMeshVisitor:
		return t.VisitPathProp(s)

	default:
		return t.VisitChildren(s)
	}
}

type NamePropContext struct {
	ModelPropContext
}

func NewNamePropContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NamePropContext {
	var p = new(NamePropContext)

	InitEmptyModelPropContext(&p.ModelPropContext)
	p.parser = parser
	p.CopyAll(ctx.(*ModelPropContext))

	return p
}

func (s *NamePropContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NamePropContext) PROP_NAME() antlr.TerminalNode {
	return s.GetToken(SQLMeshPROP_NAME, 0)
}

func (s *NamePropContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(SQLMeshEQUALS, 0)
}

func (s *NamePropContext) PropValue() IPropValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPropValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPropValueContext)
}

func (s *NamePropContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.EnterNameProp(s)
	}
}

func (s *NamePropContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.ExitNameProp(s)
	}
}

func (s *NamePropContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLMeshVisitor:
		return t.VisitNameProp(s)

	default:
		return t.VisitChildren(s)
	}
}

type OwnerPropContext struct {
	ModelPropContext
}

func NewOwnerPropContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OwnerPropContext {
	var p = new(OwnerPropContext)

	InitEmptyModelPropContext(&p.ModelPropContext)
	p.parser = parser
	p.CopyAll(ctx.(*ModelPropContext))

	return p
}

func (s *OwnerPropContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OwnerPropContext) PROP_OWNER() antlr.TerminalNode {
	return s.GetToken(SQLMeshPROP_OWNER, 0)
}

func (s *OwnerPropContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(SQLMeshEQUALS, 0)
}

func (s *OwnerPropContext) PropValue() IPropValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPropValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPropValueContext)
}

func (s *OwnerPropContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.EnterOwnerProp(s)
	}
}

func (s *OwnerPropContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.ExitOwnerProp(s)
	}
}

func (s *OwnerPropContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLMeshVisitor:
		return t.VisitOwnerProp(s)

	default:
		return t.VisitChildren(s)
	}
}

type DescriptionPropContext struct {
	ModelPropContext
}

func NewDescriptionPropContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DescriptionPropContext {
	var p = new(DescriptionPropContext)

	InitEmptyModelPropContext(&p.ModelPropContext)
	p.parser = parser
	p.CopyAll(ctx.(*ModelPropContext))

	return p
}

func (s *DescriptionPropContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DescriptionPropContext) PROP_DESCRIPTION() antlr.TerminalNode {
	return s.GetToken(SQLMeshPROP_DESCRIPTION, 0)
}

func (s *DescriptionPropContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(SQLMeshEQUALS, 0)
}

func (s *DescriptionPropContext) PropValue() IPropValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPropValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPropValueContext)
}

func (s *DescriptionPropContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.EnterDescriptionProp(s)
	}
}

func (s *DescriptionPropContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.ExitDescriptionProp(s)
	}
}

func (s *DescriptionPropContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLMeshVisitor:
		return t.VisitDescriptionProp(s)

	default:
		return t.VisitChildren(s)
	}
}

type KindPropContext struct {
	ModelPropContext
}

func NewKindPropContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *KindPropContext {
	var p = new(KindPropContext)

	InitEmptyModelPropContext(&p.ModelPropContext)
	p.parser = parser
	p.CopyAll(ctx.(*ModelPropContext))

	return p
}

func (s *KindPropContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KindPropContext) PROP_KIND() antlr.TerminalNode {
	return s.GetToken(SQLMeshPROP_KIND, 0)
}

func (s *KindPropContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(SQLMeshEQUALS, 0)
}

func (s *KindPropContext) KindValue() IKindValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKindValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IKindValueContext)
}

func (s *KindPropContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.EnterKindProp(s)
	}
}

func (s *KindPropContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.ExitKindProp(s)
	}
}

func (s *KindPropContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLMeshVisitor:
		return t.VisitKindProp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLMesh) ModelProp() (localctx IModelPropContext) {
	localctx = NewModelPropContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, SQLMeshRULE_modelProp)
	p.SetState(147)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SQLMeshPROP_NAME:
		localctx = NewNamePropContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(66)
			p.Match(SQLMeshPROP_NAME)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(67)
			p.Match(SQLMeshEQUALS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(68)
			p.PropValue()
		}

	case SQLMeshPROP_KIND:
		localctx = NewKindPropContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(69)
			p.Match(SQLMeshPROP_KIND)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(70)
			p.Match(SQLMeshEQUALS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(71)
			p.KindValue()
		}

	case SQLMeshPROP_DIALECT:
		localctx = NewDialectPropContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(72)
			p.Match(SQLMeshPROP_DIALECT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(73)
			p.Match(SQLMeshEQUALS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(74)
			p.PropValue()
		}

	case SQLMeshPROP_OWNER:
		localctx = NewOwnerPropContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(75)
			p.Match(SQLMeshPROP_OWNER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(76)
			p.Match(SQLMeshEQUALS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(77)
			p.PropValue()
		}

	case SQLMeshPROP_CRON:
		localctx = NewCronPropContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(78)
			p.Match(SQLMeshPROP_CRON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(79)
			p.Match(SQLMeshEQUALS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(80)
			p.PropValue()
		}

	case SQLMeshPROP_GRAIN:
		localctx = NewGrainPropContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(81)
			p.Match(SQLMeshPROP_GRAIN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(82)
			p.Match(SQLMeshEQUALS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(83)
			p.ArrayValue()
		}

	case SQLMeshPROP_PARTITIONED_BY:
		localctx = NewPartitionedByPropContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(84)
			p.Match(SQLMeshPROP_PARTITIONED_BY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(85)
			p.Match(SQLMeshEQUALS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(86)
			p.PropValue()
		}

	case SQLMeshPROP_CLUSTERED_BY:
		localctx = NewClusteredByPropContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(87)
			p.Match(SQLMeshPROP_CLUSTERED_BY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(88)
			p.Match(SQLMeshEQUALS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(89)
			p.ArrayValue()
		}

	case SQLMeshPROP_STORAGE_FORMAT:
		localctx = NewStorageFormatPropContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(90)
			p.Match(SQLMeshPROP_STORAGE_FORMAT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(91)
			p.Match(SQLMeshEQUALS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(92)
			p.PropValue()
		}

	case SQLMeshPROP_TABLE_FORMAT:
		localctx = NewTableFormatPropContext(p, localctx)
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(93)
			p.Match(SQLMeshPROP_TABLE_FORMAT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(94)
			p.Match(SQLMeshEQUALS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(95)
			p.PropValue()
		}

	case SQLMeshPROP_RETENTION:
		localctx = NewRetentionPropContext(p, localctx)
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(96)
			p.Match(SQLMeshPROP_RETENTION)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(97)
			p.Match(SQLMeshEQUALS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(98)
			p.PropValue()
		}

	case SQLMeshPROP_TAGS:
		localctx = NewTagsPropContext(p, localctx)
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(99)
			p.Match(SQLMeshPROP_TAGS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(100)
			p.Match(SQLMeshEQUALS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(101)
			p.ArrayValue()
		}

	case SQLMeshPROP_DESCRIPTION:
		localctx = NewDescriptionPropContext(p, localctx)
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(102)
			p.Match(SQLMeshPROP_DESCRIPTION)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(103)
			p.Match(SQLMeshEQUALS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(104)
			p.PropValue()
		}

	case SQLMeshPROP_STAMP:
		localctx = NewStampPropContext(p, localctx)
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(105)
			p.Match(SQLMeshPROP_STAMP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(106)
			p.Match(SQLMeshEQUALS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(107)
			p.PropValue()
		}

	case SQLMeshPROP_COLUMNS:
		localctx = NewColumnsPropContext(p, localctx)
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(108)
			p.Match(SQLMeshPROP_COLUMNS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(109)
			p.Match(SQLMeshEQUALS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(110)
			p.DictValue()
		}

	case SQLMeshPROP_COLUMN_DESCRIPTIONS:
		localctx = NewColumnDescriptionsPropContext(p, localctx)
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(111)
			p.Match(SQLMeshPROP_COLUMN_DESCRIPTIONS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(112)
			p.Match(SQLMeshEQUALS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(113)
			p.DictValue()
		}

	case SQLMeshPROP_PHYSICAL_PROPERTIES:
		localctx = NewPhysicalPropertiesPropContext(p, localctx)
		p.EnterOuterAlt(localctx, 17)
		{
			p.SetState(114)
			p.Match(SQLMeshPROP_PHYSICAL_PROPERTIES)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(115)
			p.Match(SQLMeshEQUALS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(116)
			p.DictValue()
		}

	case SQLMeshPROP_PRE_STATEMENTS:
		localctx = NewPreStatementsPropContext(p, localctx)
		p.EnterOuterAlt(localctx, 18)
		{
			p.SetState(117)
			p.Match(SQLMeshPROP_PRE_STATEMENTS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(118)
			p.Match(SQLMeshEQUALS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(119)
			p.ArrayValue()
		}

	case SQLMeshPROP_POST_STATEMENTS:
		localctx = NewPostStatementsPropContext(p, localctx)
		p.EnterOuterAlt(localctx, 19)
		{
			p.SetState(120)
			p.Match(SQLMeshPROP_POST_STATEMENTS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(121)
			p.Match(SQLMeshEQUALS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(122)
			p.ArrayValue()
		}

	case SQLMeshPROP_ON_VIRTUAL_UPDATE:
		localctx = NewOnVirtualUpdatePropContext(p, localctx)
		p.EnterOuterAlt(localctx, 20)
		{
			p.SetState(123)
			p.Match(SQLMeshPROP_ON_VIRTUAL_UPDATE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(124)
			p.Match(SQLMeshEQUALS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(125)
			p.ArrayValue()
		}

	case SQLMeshPROP_AUDITS:
		localctx = NewAuditsPropContext(p, localctx)
		p.EnterOuterAlt(localctx, 21)
		{
			p.SetState(126)
			p.Match(SQLMeshPROP_AUDITS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(127)
			p.Match(SQLMeshEQUALS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(128)
			p.AuditArray()
		}

	case SQLMeshPROP_TIME_COLUMN:
		localctx = NewTimeColumnPropContext(p, localctx)
		p.EnterOuterAlt(localctx, 22)
		{
			p.SetState(129)
			p.Match(SQLMeshPROP_TIME_COLUMN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(130)
			p.Match(SQLMeshEQUALS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(131)
			p.PropValue()
		}

	case SQLMeshPROP_UNIQUE_KEY:
		localctx = NewUniqueKeyPropContext(p, localctx)
		p.EnterOuterAlt(localctx, 23)
		{
			p.SetState(132)
			p.Match(SQLMeshPROP_UNIQUE_KEY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(133)
			p.Match(SQLMeshEQUALS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(134)
			p.PropValue()
		}

	case SQLMeshPROP_INVALIDATE_HARD_DELETES:
		localctx = NewInvalidateHardDeletesPropContext(p, localctx)
		p.EnterOuterAlt(localctx, 24)
		{
			p.SetState(135)
			p.Match(SQLMeshPROP_INVALIDATE_HARD_DELETES)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(136)
			p.Match(SQLMeshEQUALS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(137)
			p.PropValue()
		}

	case SQLMeshPROP_DISABLE_RESTATEMENT:
		localctx = NewDisableRestatementPropContext(p, localctx)
		p.EnterOuterAlt(localctx, 25)
		{
			p.SetState(138)
			p.Match(SQLMeshPROP_DISABLE_RESTATEMENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(139)
			p.Match(SQLMeshEQUALS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(140)
			p.PropValue()
		}

	case SQLMeshPROP_PATH:
		localctx = NewPathPropContext(p, localctx)
		p.EnterOuterAlt(localctx, 26)
		{
			p.SetState(141)
			p.Match(SQLMeshPROP_PATH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(142)
			p.Match(SQLMeshEQUALS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(143)
			p.PropValue()
		}

	case SQLMeshPROP_CSV_SETTINGS:
		localctx = NewCsvSettingsPropContext(p, localctx)
		p.EnterOuterAlt(localctx, 27)
		{
			p.SetState(144)
			p.Match(SQLMeshPROP_CSV_SETTINGS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(145)
			p.Match(SQLMeshEQUALS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(146)
			p.DictValue()
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

// IKindValueContext is an interface to support dynamic dispatch.
type IKindValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsKindValueContext differentiates from other interfaces.
	IsKindValueContext()
}

type KindValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKindValueContext() *KindValueContext {
	var p = new(KindValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLMeshRULE_kindValue
	return p
}

func InitEmptyKindValueContext(p *KindValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLMeshRULE_kindValue
}

func (*KindValueContext) IsKindValueContext() {}

func NewKindValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KindValueContext {
	var p = new(KindValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLMeshRULE_kindValue

	return p
}

func (s *KindValueContext) GetParser() antlr.Parser { return s.parser }

func (s *KindValueContext) CopyAll(ctx *KindValueContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *KindValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KindValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type KindFullContext struct {
	KindValueContext
}

func NewKindFullContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *KindFullContext {
	var p = new(KindFullContext)

	InitEmptyKindValueContext(&p.KindValueContext)
	p.parser = parser
	p.CopyAll(ctx.(*KindValueContext))

	return p
}

func (s *KindFullContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KindFullContext) KIND_FULL() antlr.TerminalNode {
	return s.GetToken(SQLMeshKIND_FULL, 0)
}

func (s *KindFullContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.EnterKindFull(s)
	}
}

func (s *KindFullContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.ExitKindFull(s)
	}
}

func (s *KindFullContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLMeshVisitor:
		return t.VisitKindFull(s)

	default:
		return t.VisitChildren(s)
	}
}

type KindViewContext struct {
	KindValueContext
}

func NewKindViewContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *KindViewContext {
	var p = new(KindViewContext)

	InitEmptyKindValueContext(&p.KindValueContext)
	p.parser = parser
	p.CopyAll(ctx.(*KindValueContext))

	return p
}

func (s *KindViewContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KindViewContext) KIND_VIEW() antlr.TerminalNode {
	return s.GetToken(SQLMeshKIND_VIEW, 0)
}

func (s *KindViewContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.EnterKindView(s)
	}
}

func (s *KindViewContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.ExitKindView(s)
	}
}

func (s *KindViewContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLMeshVisitor:
		return t.VisitKindView(s)

	default:
		return t.VisitChildren(s)
	}
}

type KindExternalContext struct {
	KindValueContext
}

func NewKindExternalContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *KindExternalContext {
	var p = new(KindExternalContext)

	InitEmptyKindValueContext(&p.KindValueContext)
	p.parser = parser
	p.CopyAll(ctx.(*KindValueContext))

	return p
}

func (s *KindExternalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KindExternalContext) KIND_EXTERNAL() antlr.TerminalNode {
	return s.GetToken(SQLMeshKIND_EXTERNAL, 0)
}

func (s *KindExternalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.EnterKindExternal(s)
	}
}

func (s *KindExternalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.ExitKindExternal(s)
	}
}

func (s *KindExternalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLMeshVisitor:
		return t.VisitKindExternal(s)

	default:
		return t.VisitChildren(s)
	}
}

type KindEmbeddedContext struct {
	KindValueContext
}

func NewKindEmbeddedContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *KindEmbeddedContext {
	var p = new(KindEmbeddedContext)

	InitEmptyKindValueContext(&p.KindValueContext)
	p.parser = parser
	p.CopyAll(ctx.(*KindValueContext))

	return p
}

func (s *KindEmbeddedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KindEmbeddedContext) KIND_EMBEDDED() antlr.TerminalNode {
	return s.GetToken(SQLMeshKIND_EMBEDDED, 0)
}

func (s *KindEmbeddedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.EnterKindEmbedded(s)
	}
}

func (s *KindEmbeddedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.ExitKindEmbedded(s)
	}
}

func (s *KindEmbeddedContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLMeshVisitor:
		return t.VisitKindEmbedded(s)

	default:
		return t.VisitChildren(s)
	}
}

type KindScdType2Context struct {
	KindValueContext
}

func NewKindScdType2Context(parser antlr.Parser, ctx antlr.ParserRuleContext) *KindScdType2Context {
	var p = new(KindScdType2Context)

	InitEmptyKindValueContext(&p.KindValueContext)
	p.parser = parser
	p.CopyAll(ctx.(*KindValueContext))

	return p
}

func (s *KindScdType2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KindScdType2Context) KIND_SCD_TYPE_2() antlr.TerminalNode {
	return s.GetToken(SQLMeshKIND_SCD_TYPE_2, 0)
}

func (s *KindScdType2Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.EnterKindScdType2(s)
	}
}

func (s *KindScdType2Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.ExitKindScdType2(s)
	}
}

func (s *KindScdType2Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLMeshVisitor:
		return t.VisitKindScdType2(s)

	default:
		return t.VisitChildren(s)
	}
}

type KindIncrementalByUniqueKeyContext struct {
	KindValueContext
}

func NewKindIncrementalByUniqueKeyContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *KindIncrementalByUniqueKeyContext {
	var p = new(KindIncrementalByUniqueKeyContext)

	InitEmptyKindValueContext(&p.KindValueContext)
	p.parser = parser
	p.CopyAll(ctx.(*KindValueContext))

	return p
}

func (s *KindIncrementalByUniqueKeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KindIncrementalByUniqueKeyContext) KIND_INCREMENTAL_BY_UNIQUE_KEY() antlr.TerminalNode {
	return s.GetToken(SQLMeshKIND_INCREMENTAL_BY_UNIQUE_KEY, 0)
}

func (s *KindIncrementalByUniqueKeyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.EnterKindIncrementalByUniqueKey(s)
	}
}

func (s *KindIncrementalByUniqueKeyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.ExitKindIncrementalByUniqueKey(s)
	}
}

func (s *KindIncrementalByUniqueKeyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLMeshVisitor:
		return t.VisitKindIncrementalByUniqueKey(s)

	default:
		return t.VisitChildren(s)
	}
}

type KindIncrementalByPartitionContext struct {
	KindValueContext
}

func NewKindIncrementalByPartitionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *KindIncrementalByPartitionContext {
	var p = new(KindIncrementalByPartitionContext)

	InitEmptyKindValueContext(&p.KindValueContext)
	p.parser = parser
	p.CopyAll(ctx.(*KindValueContext))

	return p
}

func (s *KindIncrementalByPartitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KindIncrementalByPartitionContext) KIND_INCREMENTAL_BY_PARTITION() antlr.TerminalNode {
	return s.GetToken(SQLMeshKIND_INCREMENTAL_BY_PARTITION, 0)
}

func (s *KindIncrementalByPartitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.EnterKindIncrementalByPartition(s)
	}
}

func (s *KindIncrementalByPartitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.ExitKindIncrementalByPartition(s)
	}
}

func (s *KindIncrementalByPartitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLMeshVisitor:
		return t.VisitKindIncrementalByPartition(s)

	default:
		return t.VisitChildren(s)
	}
}

type KindSeedContext struct {
	KindValueContext
}

func NewKindSeedContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *KindSeedContext {
	var p = new(KindSeedContext)

	InitEmptyKindValueContext(&p.KindValueContext)
	p.parser = parser
	p.CopyAll(ctx.(*KindValueContext))

	return p
}

func (s *KindSeedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KindSeedContext) KIND_SEED() antlr.TerminalNode {
	return s.GetToken(SQLMeshKIND_SEED, 0)
}

func (s *KindSeedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.EnterKindSeed(s)
	}
}

func (s *KindSeedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.ExitKindSeed(s)
	}
}

func (s *KindSeedContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLMeshVisitor:
		return t.VisitKindSeed(s)

	default:
		return t.VisitChildren(s)
	}
}

type KindIncrementalByTimeRangeContext struct {
	KindValueContext
}

func NewKindIncrementalByTimeRangeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *KindIncrementalByTimeRangeContext {
	var p = new(KindIncrementalByTimeRangeContext)

	InitEmptyKindValueContext(&p.KindValueContext)
	p.parser = parser
	p.CopyAll(ctx.(*KindValueContext))

	return p
}

func (s *KindIncrementalByTimeRangeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KindIncrementalByTimeRangeContext) KIND_INCREMENTAL_BY_TIME_RANGE() antlr.TerminalNode {
	return s.GetToken(SQLMeshKIND_INCREMENTAL_BY_TIME_RANGE, 0)
}

func (s *KindIncrementalByTimeRangeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.EnterKindIncrementalByTimeRange(s)
	}
}

func (s *KindIncrementalByTimeRangeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.ExitKindIncrementalByTimeRange(s)
	}
}

func (s *KindIncrementalByTimeRangeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLMeshVisitor:
		return t.VisitKindIncrementalByTimeRange(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLMesh) KindValue() (localctx IKindValueContext) {
	localctx = NewKindValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, SQLMeshRULE_kindValue)
	p.SetState(158)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SQLMeshKIND_FULL:
		localctx = NewKindFullContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(149)
			p.Match(SQLMeshKIND_FULL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SQLMeshKIND_INCREMENTAL_BY_TIME_RANGE:
		localctx = NewKindIncrementalByTimeRangeContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(150)
			p.Match(SQLMeshKIND_INCREMENTAL_BY_TIME_RANGE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SQLMeshKIND_INCREMENTAL_BY_UNIQUE_KEY:
		localctx = NewKindIncrementalByUniqueKeyContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(151)
			p.Match(SQLMeshKIND_INCREMENTAL_BY_UNIQUE_KEY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SQLMeshKIND_INCREMENTAL_BY_PARTITION:
		localctx = NewKindIncrementalByPartitionContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(152)
			p.Match(SQLMeshKIND_INCREMENTAL_BY_PARTITION)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SQLMeshKIND_VIEW:
		localctx = NewKindViewContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(153)
			p.Match(SQLMeshKIND_VIEW)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SQLMeshKIND_SEED:
		localctx = NewKindSeedContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(154)
			p.Match(SQLMeshKIND_SEED)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SQLMeshKIND_EXTERNAL:
		localctx = NewKindExternalContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(155)
			p.Match(SQLMeshKIND_EXTERNAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SQLMeshKIND_EMBEDDED:
		localctx = NewKindEmbeddedContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(156)
			p.Match(SQLMeshKIND_EMBEDDED)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SQLMeshKIND_SCD_TYPE_2:
		localctx = NewKindScdType2Context(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(157)
			p.Match(SQLMeshKIND_SCD_TYPE_2)
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

// IAuditDefContext is an interface to support dynamic dispatch.
type IAuditDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AUDIT() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	SqlBody() ISqlBodyContext
	IDENTIFIER() antlr.TerminalNode
	AllAuditProp() []IAuditPropContext
	AuditProp(i int) IAuditPropContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsAuditDefContext differentiates from other interfaces.
	IsAuditDefContext()
}

type AuditDefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAuditDefContext() *AuditDefContext {
	var p = new(AuditDefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLMeshRULE_auditDef
	return p
}

func InitEmptyAuditDefContext(p *AuditDefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLMeshRULE_auditDef
}

func (*AuditDefContext) IsAuditDefContext() {}

func NewAuditDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AuditDefContext {
	var p = new(AuditDefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLMeshRULE_auditDef

	return p
}

func (s *AuditDefContext) GetParser() antlr.Parser { return s.parser }

func (s *AuditDefContext) AUDIT() antlr.TerminalNode {
	return s.GetToken(SQLMeshAUDIT, 0)
}

func (s *AuditDefContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(SQLMeshLPAREN, 0)
}

func (s *AuditDefContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(SQLMeshRPAREN, 0)
}

func (s *AuditDefContext) SqlBody() ISqlBodyContext {
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

func (s *AuditDefContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(SQLMeshIDENTIFIER, 0)
}

func (s *AuditDefContext) AllAuditProp() []IAuditPropContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAuditPropContext); ok {
			len++
		}
	}

	tst := make([]IAuditPropContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAuditPropContext); ok {
			tst[i] = t.(IAuditPropContext)
			i++
		}
	}

	return tst
}

func (s *AuditDefContext) AuditProp(i int) IAuditPropContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAuditPropContext); ok {
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

	return t.(IAuditPropContext)
}

func (s *AuditDefContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SQLMeshCOMMA)
}

func (s *AuditDefContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SQLMeshCOMMA, i)
}

func (s *AuditDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AuditDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AuditDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.EnterAuditDef(s)
	}
}

func (s *AuditDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.ExitAuditDef(s)
	}
}

func (s *AuditDefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLMeshVisitor:
		return t.VisitAuditDef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLMesh) AuditDef() (localctx IAuditDefContext) {
	localctx = NewAuditDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, SQLMeshRULE_auditDef)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(160)
		p.Match(SQLMeshAUDIT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(162)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SQLMeshIDENTIFIER {
		{
			p.SetState(161)
			p.Match(SQLMeshIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(164)
		p.Match(SQLMeshLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(176)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SQLMeshPROP_DIALECT || _la == SQLMeshPROP_QUERY {
		{
			p.SetState(165)
			p.AuditProp()
		}
		p.SetState(170)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(166)
					p.Match(SQLMeshCOMMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(167)
					p.AuditProp()
				}

			}
			p.SetState(172)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		p.SetState(174)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SQLMeshCOMMA {
			{
				p.SetState(173)
				p.Match(SQLMeshCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	}
	{
		p.SetState(178)
		p.Match(SQLMeshRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(179)
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

// IAuditPropContext is an interface to support dynamic dispatch.
type IAuditPropContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsAuditPropContext differentiates from other interfaces.
	IsAuditPropContext()
}

type AuditPropContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAuditPropContext() *AuditPropContext {
	var p = new(AuditPropContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLMeshRULE_auditProp
	return p
}

func InitEmptyAuditPropContext(p *AuditPropContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLMeshRULE_auditProp
}

func (*AuditPropContext) IsAuditPropContext() {}

func NewAuditPropContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AuditPropContext {
	var p = new(AuditPropContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLMeshRULE_auditProp

	return p
}

func (s *AuditPropContext) GetParser() antlr.Parser { return s.parser }

func (s *AuditPropContext) CopyAll(ctx *AuditPropContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *AuditPropContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AuditPropContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type QueryPropContext struct {
	AuditPropContext
}

func NewQueryPropContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *QueryPropContext {
	var p = new(QueryPropContext)

	InitEmptyAuditPropContext(&p.AuditPropContext)
	p.parser = parser
	p.CopyAll(ctx.(*AuditPropContext))

	return p
}

func (s *QueryPropContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QueryPropContext) PROP_QUERY() antlr.TerminalNode {
	return s.GetToken(SQLMeshPROP_QUERY, 0)
}

func (s *QueryPropContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(SQLMeshEQUALS, 0)
}

func (s *QueryPropContext) PropValue() IPropValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPropValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPropValueContext)
}

func (s *QueryPropContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.EnterQueryProp(s)
	}
}

func (s *QueryPropContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.ExitQueryProp(s)
	}
}

func (s *QueryPropContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLMeshVisitor:
		return t.VisitQueryProp(s)

	default:
		return t.VisitChildren(s)
	}
}

type AuditDialectPropContext struct {
	AuditPropContext
}

func NewAuditDialectPropContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AuditDialectPropContext {
	var p = new(AuditDialectPropContext)

	InitEmptyAuditPropContext(&p.AuditPropContext)
	p.parser = parser
	p.CopyAll(ctx.(*AuditPropContext))

	return p
}

func (s *AuditDialectPropContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AuditDialectPropContext) PROP_DIALECT() antlr.TerminalNode {
	return s.GetToken(SQLMeshPROP_DIALECT, 0)
}

func (s *AuditDialectPropContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(SQLMeshEQUALS, 0)
}

func (s *AuditDialectPropContext) PropValue() IPropValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPropValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPropValueContext)
}

func (s *AuditDialectPropContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.EnterAuditDialectProp(s)
	}
}

func (s *AuditDialectPropContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.ExitAuditDialectProp(s)
	}
}

func (s *AuditDialectPropContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLMeshVisitor:
		return t.VisitAuditDialectProp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLMesh) AuditProp() (localctx IAuditPropContext) {
	localctx = NewAuditPropContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, SQLMeshRULE_auditProp)
	p.SetState(187)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SQLMeshPROP_QUERY:
		localctx = NewQueryPropContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(181)
			p.Match(SQLMeshPROP_QUERY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(182)
			p.Match(SQLMeshEQUALS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(183)
			p.PropValue()
		}

	case SQLMeshPROP_DIALECT:
		localctx = NewAuditDialectPropContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(184)
			p.Match(SQLMeshPROP_DIALECT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(185)
			p.Match(SQLMeshEQUALS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(186)
			p.PropValue()
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

// IAuditArrayContext is an interface to support dynamic dispatch.
type IAuditArrayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACKET() antlr.TerminalNode
	RBRACKET() antlr.TerminalNode
	AllAuditEntry() []IAuditEntryContext
	AuditEntry(i int) IAuditEntryContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsAuditArrayContext differentiates from other interfaces.
	IsAuditArrayContext()
}

type AuditArrayContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAuditArrayContext() *AuditArrayContext {
	var p = new(AuditArrayContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLMeshRULE_auditArray
	return p
}

func InitEmptyAuditArrayContext(p *AuditArrayContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLMeshRULE_auditArray
}

func (*AuditArrayContext) IsAuditArrayContext() {}

func NewAuditArrayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AuditArrayContext {
	var p = new(AuditArrayContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLMeshRULE_auditArray

	return p
}

func (s *AuditArrayContext) GetParser() antlr.Parser { return s.parser }

func (s *AuditArrayContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(SQLMeshLBRACKET, 0)
}

func (s *AuditArrayContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(SQLMeshRBRACKET, 0)
}

func (s *AuditArrayContext) AllAuditEntry() []IAuditEntryContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAuditEntryContext); ok {
			len++
		}
	}

	tst := make([]IAuditEntryContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAuditEntryContext); ok {
			tst[i] = t.(IAuditEntryContext)
			i++
		}
	}

	return tst
}

func (s *AuditArrayContext) AuditEntry(i int) IAuditEntryContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAuditEntryContext); ok {
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

	return t.(IAuditEntryContext)
}

func (s *AuditArrayContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SQLMeshCOMMA)
}

func (s *AuditArrayContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SQLMeshCOMMA, i)
}

func (s *AuditArrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AuditArrayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AuditArrayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.EnterAuditArray(s)
	}
}

func (s *AuditArrayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.ExitAuditArray(s)
	}
}

func (s *AuditArrayContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLMeshVisitor:
		return t.VisitAuditArray(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLMesh) AuditArray() (localctx IAuditArrayContext) {
	localctx = NewAuditArrayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, SQLMeshRULE_auditArray)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(189)
		p.Match(SQLMeshLBRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(201)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SQLMeshLBRACE {
		{
			p.SetState(190)
			p.AuditEntry()
		}
		p.SetState(195)
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
					p.SetState(191)
					p.Match(SQLMeshCOMMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(192)
					p.AuditEntry()
				}

			}
			p.SetState(197)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		p.SetState(199)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SQLMeshCOMMA {
			{
				p.SetState(198)
				p.Match(SQLMeshCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	}
	{
		p.SetState(203)
		p.Match(SQLMeshRBRACKET)
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

// IAuditEntryContext is an interface to support dynamic dispatch.
type IAuditEntryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllAuditEntryProp() []IAuditEntryPropContext
	AuditEntryProp(i int) IAuditEntryPropContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsAuditEntryContext differentiates from other interfaces.
	IsAuditEntryContext()
}

type AuditEntryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAuditEntryContext() *AuditEntryContext {
	var p = new(AuditEntryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLMeshRULE_auditEntry
	return p
}

func InitEmptyAuditEntryContext(p *AuditEntryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLMeshRULE_auditEntry
}

func (*AuditEntryContext) IsAuditEntryContext() {}

func NewAuditEntryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AuditEntryContext {
	var p = new(AuditEntryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLMeshRULE_auditEntry

	return p
}

func (s *AuditEntryContext) GetParser() antlr.Parser { return s.parser }

func (s *AuditEntryContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(SQLMeshLBRACE, 0)
}

func (s *AuditEntryContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(SQLMeshRBRACE, 0)
}

func (s *AuditEntryContext) AllAuditEntryProp() []IAuditEntryPropContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAuditEntryPropContext); ok {
			len++
		}
	}

	tst := make([]IAuditEntryPropContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAuditEntryPropContext); ok {
			tst[i] = t.(IAuditEntryPropContext)
			i++
		}
	}

	return tst
}

func (s *AuditEntryContext) AuditEntryProp(i int) IAuditEntryPropContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAuditEntryPropContext); ok {
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

	return t.(IAuditEntryPropContext)
}

func (s *AuditEntryContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SQLMeshCOMMA)
}

func (s *AuditEntryContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SQLMeshCOMMA, i)
}

func (s *AuditEntryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AuditEntryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AuditEntryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.EnterAuditEntry(s)
	}
}

func (s *AuditEntryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.ExitAuditEntry(s)
	}
}

func (s *AuditEntryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLMeshVisitor:
		return t.VisitAuditEntry(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLMesh) AuditEntry() (localctx IAuditEntryContext) {
	localctx = NewAuditEntryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, SQLMeshRULE_auditEntry)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(205)
		p.Match(SQLMeshLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(217)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SQLMeshPROP_NAME || _la == SQLMeshPROP_QUERY {
		{
			p.SetState(206)
			p.AuditEntryProp()
		}
		p.SetState(211)
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
					p.SetState(207)
					p.Match(SQLMeshCOMMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(208)
					p.AuditEntryProp()
				}

			}
			p.SetState(213)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		p.SetState(215)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SQLMeshCOMMA {
			{
				p.SetState(214)
				p.Match(SQLMeshCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	}
	{
		p.SetState(219)
		p.Match(SQLMeshRBRACE)
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

// IAuditEntryPropContext is an interface to support dynamic dispatch.
type IAuditEntryPropContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsAuditEntryPropContext differentiates from other interfaces.
	IsAuditEntryPropContext()
}

type AuditEntryPropContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAuditEntryPropContext() *AuditEntryPropContext {
	var p = new(AuditEntryPropContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLMeshRULE_auditEntryProp
	return p
}

func InitEmptyAuditEntryPropContext(p *AuditEntryPropContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLMeshRULE_auditEntryProp
}

func (*AuditEntryPropContext) IsAuditEntryPropContext() {}

func NewAuditEntryPropContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AuditEntryPropContext {
	var p = new(AuditEntryPropContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLMeshRULE_auditEntryProp

	return p
}

func (s *AuditEntryPropContext) GetParser() antlr.Parser { return s.parser }

func (s *AuditEntryPropContext) CopyAll(ctx *AuditEntryPropContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *AuditEntryPropContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AuditEntryPropContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type AuditEntryQueryPropContext struct {
	AuditEntryPropContext
}

func NewAuditEntryQueryPropContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AuditEntryQueryPropContext {
	var p = new(AuditEntryQueryPropContext)

	InitEmptyAuditEntryPropContext(&p.AuditEntryPropContext)
	p.parser = parser
	p.CopyAll(ctx.(*AuditEntryPropContext))

	return p
}

func (s *AuditEntryQueryPropContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AuditEntryQueryPropContext) PROP_QUERY() antlr.TerminalNode {
	return s.GetToken(SQLMeshPROP_QUERY, 0)
}

func (s *AuditEntryQueryPropContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(SQLMeshEQUALS, 0)
}

func (s *AuditEntryQueryPropContext) PropValue() IPropValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPropValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPropValueContext)
}

func (s *AuditEntryQueryPropContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.EnterAuditEntryQueryProp(s)
	}
}

func (s *AuditEntryQueryPropContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.ExitAuditEntryQueryProp(s)
	}
}

func (s *AuditEntryQueryPropContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLMeshVisitor:
		return t.VisitAuditEntryQueryProp(s)

	default:
		return t.VisitChildren(s)
	}
}

type AuditEntryNamePropContext struct {
	AuditEntryPropContext
}

func NewAuditEntryNamePropContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AuditEntryNamePropContext {
	var p = new(AuditEntryNamePropContext)

	InitEmptyAuditEntryPropContext(&p.AuditEntryPropContext)
	p.parser = parser
	p.CopyAll(ctx.(*AuditEntryPropContext))

	return p
}

func (s *AuditEntryNamePropContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AuditEntryNamePropContext) PROP_NAME() antlr.TerminalNode {
	return s.GetToken(SQLMeshPROP_NAME, 0)
}

func (s *AuditEntryNamePropContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(SQLMeshEQUALS, 0)
}

func (s *AuditEntryNamePropContext) PropValue() IPropValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPropValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPropValueContext)
}

func (s *AuditEntryNamePropContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.EnterAuditEntryNameProp(s)
	}
}

func (s *AuditEntryNamePropContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.ExitAuditEntryNameProp(s)
	}
}

func (s *AuditEntryNamePropContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLMeshVisitor:
		return t.VisitAuditEntryNameProp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLMesh) AuditEntryProp() (localctx IAuditEntryPropContext) {
	localctx = NewAuditEntryPropContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, SQLMeshRULE_auditEntryProp)
	p.SetState(227)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SQLMeshPROP_NAME:
		localctx = NewAuditEntryNamePropContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(221)
			p.Match(SQLMeshPROP_NAME)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(222)
			p.Match(SQLMeshEQUALS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(223)
			p.PropValue()
		}

	case SQLMeshPROP_QUERY:
		localctx = NewAuditEntryQueryPropContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(224)
			p.Match(SQLMeshPROP_QUERY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(225)
			p.Match(SQLMeshEQUALS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(226)
			p.PropValue()
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

// IPropValueContext is an interface to support dynamic dispatch.
type IPropValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Literal() ILiteralContext
	ArrayValue() IArrayValueContext
	DictValue() IDictValueContext
	MacroRef() IMacroRefContext

	// IsPropValueContext differentiates from other interfaces.
	IsPropValueContext()
}

type PropValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPropValueContext() *PropValueContext {
	var p = new(PropValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLMeshRULE_propValue
	return p
}

func InitEmptyPropValueContext(p *PropValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLMeshRULE_propValue
}

func (*PropValueContext) IsPropValueContext() {}

func NewPropValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PropValueContext {
	var p = new(PropValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLMeshRULE_propValue

	return p
}

func (s *PropValueContext) GetParser() antlr.Parser { return s.parser }

func (s *PropValueContext) Literal() ILiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *PropValueContext) ArrayValue() IArrayValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayValueContext)
}

func (s *PropValueContext) DictValue() IDictValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDictValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDictValueContext)
}

func (s *PropValueContext) MacroRef() IMacroRefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMacroRefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMacroRefContext)
}

func (s *PropValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PropValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.EnterPropValue(s)
	}
}

func (s *PropValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.ExitPropValue(s)
	}
}

func (s *PropValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLMeshVisitor:
		return t.VisitPropValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLMesh) PropValue() (localctx IPropValueContext) {
	localctx = NewPropValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, SQLMeshRULE_propValue)
	p.SetState(233)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SQLMeshTRUE, SQLMeshFALSE, SQLMeshFLOAT_LITERAL, SQLMeshINT_LITERAL, SQLMeshSTRING_LITERAL, SQLMeshIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(229)
			p.Literal()
		}

	case SQLMeshLBRACKET:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(230)
			p.ArrayValue()
		}

	case SQLMeshLBRACE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(231)
			p.DictValue()
		}

	case SQLMeshMACRO_BRACED, SQLMeshMACRO_IDENT:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(232)
			p.MacroRef()
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

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLMeshRULE_literal
	return p
}

func InitEmptyLiteralContext(p *LiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLMeshRULE_literal
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLMeshRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) CopyAll(ctx *LiteralContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type QualifiedIdentLitContext struct {
	LiteralContext
}

func NewQualifiedIdentLitContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *QualifiedIdentLitContext {
	var p = new(QualifiedIdentLitContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *QualifiedIdentLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QualifiedIdentLitContext) QualifiedIdent() IQualifiedIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IQualifiedIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IQualifiedIdentContext)
}

func (s *QualifiedIdentLitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.EnterQualifiedIdentLit(s)
	}
}

func (s *QualifiedIdentLitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.ExitQualifiedIdentLit(s)
	}
}

func (s *QualifiedIdentLitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLMeshVisitor:
		return t.VisitQualifiedIdentLit(s)

	default:
		return t.VisitChildren(s)
	}
}

type StringLitContext struct {
	LiteralContext
}

func NewStringLitContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringLitContext {
	var p = new(StringLitContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *StringLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringLitContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(SQLMeshSTRING_LITERAL, 0)
}

func (s *StringLitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.EnterStringLit(s)
	}
}

func (s *StringLitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.ExitStringLit(s)
	}
}

func (s *StringLitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLMeshVisitor:
		return t.VisitStringLit(s)

	default:
		return t.VisitChildren(s)
	}
}

type TrueLitContext struct {
	LiteralContext
}

func NewTrueLitContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TrueLitContext {
	var p = new(TrueLitContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *TrueLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TrueLitContext) TRUE() antlr.TerminalNode {
	return s.GetToken(SQLMeshTRUE, 0)
}

func (s *TrueLitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.EnterTrueLit(s)
	}
}

func (s *TrueLitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.ExitTrueLit(s)
	}
}

func (s *TrueLitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLMeshVisitor:
		return t.VisitTrueLit(s)

	default:
		return t.VisitChildren(s)
	}
}

type FloatLitContext struct {
	LiteralContext
}

func NewFloatLitContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FloatLitContext {
	var p = new(FloatLitContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *FloatLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatLitContext) FLOAT_LITERAL() antlr.TerminalNode {
	return s.GetToken(SQLMeshFLOAT_LITERAL, 0)
}

func (s *FloatLitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.EnterFloatLit(s)
	}
}

func (s *FloatLitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.ExitFloatLit(s)
	}
}

func (s *FloatLitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLMeshVisitor:
		return t.VisitFloatLit(s)

	default:
		return t.VisitChildren(s)
	}
}

type FalseLitContext struct {
	LiteralContext
}

func NewFalseLitContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FalseLitContext {
	var p = new(FalseLitContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *FalseLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FalseLitContext) FALSE() antlr.TerminalNode {
	return s.GetToken(SQLMeshFALSE, 0)
}

func (s *FalseLitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.EnterFalseLit(s)
	}
}

func (s *FalseLitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.ExitFalseLit(s)
	}
}

func (s *FalseLitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLMeshVisitor:
		return t.VisitFalseLit(s)

	default:
		return t.VisitChildren(s)
	}
}

type IntLitContext struct {
	LiteralContext
}

func NewIntLitContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntLitContext {
	var p = new(IntLitContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *IntLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntLitContext) INT_LITERAL() antlr.TerminalNode {
	return s.GetToken(SQLMeshINT_LITERAL, 0)
}

func (s *IntLitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.EnterIntLit(s)
	}
}

func (s *IntLitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.ExitIntLit(s)
	}
}

func (s *IntLitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLMeshVisitor:
		return t.VisitIntLit(s)

	default:
		return t.VisitChildren(s)
	}
}

type IdentLitContext struct {
	LiteralContext
}

func NewIdentLitContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdentLitContext {
	var p = new(IdentLitContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *IdentLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentLitContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(SQLMeshIDENTIFIER, 0)
}

func (s *IdentLitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.EnterIdentLit(s)
	}
}

func (s *IdentLitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.ExitIdentLit(s)
	}
}

func (s *IdentLitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLMeshVisitor:
		return t.VisitIdentLit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLMesh) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, SQLMeshRULE_literal)
	p.SetState(242)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext()) {
	case 1:
		localctx = NewStringLitContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(235)
			p.Match(SQLMeshSTRING_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewIntLitContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(236)
			p.Match(SQLMeshINT_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewFloatLitContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(237)
			p.Match(SQLMeshFLOAT_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewTrueLitContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(238)
			p.Match(SQLMeshTRUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		localctx = NewFalseLitContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(239)
			p.Match(SQLMeshFALSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		localctx = NewQualifiedIdentLitContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(240)
			p.QualifiedIdent()
		}

	case 7:
		localctx = NewIdentLitContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(241)
			p.Match(SQLMeshIDENTIFIER)
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

// IQualifiedIdentContext is an interface to support dynamic dispatch.
type IQualifiedIdentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIDENTIFIER() []antlr.TerminalNode
	IDENTIFIER(i int) antlr.TerminalNode
	AllDOT() []antlr.TerminalNode
	DOT(i int) antlr.TerminalNode

	// IsQualifiedIdentContext differentiates from other interfaces.
	IsQualifiedIdentContext()
}

type QualifiedIdentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQualifiedIdentContext() *QualifiedIdentContext {
	var p = new(QualifiedIdentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLMeshRULE_qualifiedIdent
	return p
}

func InitEmptyQualifiedIdentContext(p *QualifiedIdentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLMeshRULE_qualifiedIdent
}

func (*QualifiedIdentContext) IsQualifiedIdentContext() {}

func NewQualifiedIdentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QualifiedIdentContext {
	var p = new(QualifiedIdentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLMeshRULE_qualifiedIdent

	return p
}

func (s *QualifiedIdentContext) GetParser() antlr.Parser { return s.parser }

func (s *QualifiedIdentContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(SQLMeshIDENTIFIER)
}

func (s *QualifiedIdentContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(SQLMeshIDENTIFIER, i)
}

func (s *QualifiedIdentContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(SQLMeshDOT)
}

func (s *QualifiedIdentContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(SQLMeshDOT, i)
}

func (s *QualifiedIdentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QualifiedIdentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QualifiedIdentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.EnterQualifiedIdent(s)
	}
}

func (s *QualifiedIdentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.ExitQualifiedIdent(s)
	}
}

func (s *QualifiedIdentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLMeshVisitor:
		return t.VisitQualifiedIdent(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLMesh) QualifiedIdent() (localctx IQualifiedIdentContext) {
	localctx = NewQualifiedIdentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, SQLMeshRULE_qualifiedIdent)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(244)
		p.Match(SQLMeshIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(247)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == SQLMeshDOT {
		{
			p.SetState(245)
			p.Match(SQLMeshDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(246)
			p.Match(SQLMeshIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(249)
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

// IArrayValueContext is an interface to support dynamic dispatch.
type IArrayValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACKET() antlr.TerminalNode
	RBRACKET() antlr.TerminalNode
	AllPropValue() []IPropValueContext
	PropValue(i int) IPropValueContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsArrayValueContext differentiates from other interfaces.
	IsArrayValueContext()
}

type ArrayValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayValueContext() *ArrayValueContext {
	var p = new(ArrayValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLMeshRULE_arrayValue
	return p
}

func InitEmptyArrayValueContext(p *ArrayValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLMeshRULE_arrayValue
}

func (*ArrayValueContext) IsArrayValueContext() {}

func NewArrayValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayValueContext {
	var p = new(ArrayValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLMeshRULE_arrayValue

	return p
}

func (s *ArrayValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayValueContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(SQLMeshLBRACKET, 0)
}

func (s *ArrayValueContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(SQLMeshRBRACKET, 0)
}

func (s *ArrayValueContext) AllPropValue() []IPropValueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPropValueContext); ok {
			len++
		}
	}

	tst := make([]IPropValueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPropValueContext); ok {
			tst[i] = t.(IPropValueContext)
			i++
		}
	}

	return tst
}

func (s *ArrayValueContext) PropValue(i int) IPropValueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPropValueContext); ok {
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

	return t.(IPropValueContext)
}

func (s *ArrayValueContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SQLMeshCOMMA)
}

func (s *ArrayValueContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SQLMeshCOMMA, i)
}

func (s *ArrayValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.EnterArrayValue(s)
	}
}

func (s *ArrayValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.ExitArrayValue(s)
	}
}

func (s *ArrayValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLMeshVisitor:
		return t.VisitArrayValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLMesh) ArrayValue() (localctx IArrayValueContext) {
	localctx = NewArrayValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, SQLMeshRULE_arrayValue)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(251)
		p.Match(SQLMeshLBRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(263)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4540332111831334912) != 0 {
		{
			p.SetState(252)
			p.PropValue()
		}
		p.SetState(257)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(253)
					p.Match(SQLMeshCOMMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(254)
					p.PropValue()
				}

			}
			p.SetState(259)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		p.SetState(261)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SQLMeshCOMMA {
			{
				p.SetState(260)
				p.Match(SQLMeshCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	}
	{
		p.SetState(265)
		p.Match(SQLMeshRBRACKET)
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

// IDictValueContext is an interface to support dynamic dispatch.
type IDictValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllDictEntry() []IDictEntryContext
	DictEntry(i int) IDictEntryContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsDictValueContext differentiates from other interfaces.
	IsDictValueContext()
}

type DictValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDictValueContext() *DictValueContext {
	var p = new(DictValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLMeshRULE_dictValue
	return p
}

func InitEmptyDictValueContext(p *DictValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLMeshRULE_dictValue
}

func (*DictValueContext) IsDictValueContext() {}

func NewDictValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DictValueContext {
	var p = new(DictValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLMeshRULE_dictValue

	return p
}

func (s *DictValueContext) GetParser() antlr.Parser { return s.parser }

func (s *DictValueContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(SQLMeshLBRACE, 0)
}

func (s *DictValueContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(SQLMeshRBRACE, 0)
}

func (s *DictValueContext) AllDictEntry() []IDictEntryContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDictEntryContext); ok {
			len++
		}
	}

	tst := make([]IDictEntryContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDictEntryContext); ok {
			tst[i] = t.(IDictEntryContext)
			i++
		}
	}

	return tst
}

func (s *DictValueContext) DictEntry(i int) IDictEntryContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDictEntryContext); ok {
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

	return t.(IDictEntryContext)
}

func (s *DictValueContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SQLMeshCOMMA)
}

func (s *DictValueContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SQLMeshCOMMA, i)
}

func (s *DictValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DictValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DictValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.EnterDictValue(s)
	}
}

func (s *DictValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.ExitDictValue(s)
	}
}

func (s *DictValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLMeshVisitor:
		return t.VisitDictValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLMesh) DictValue() (localctx IDictValueContext) {
	localctx = NewDictValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, SQLMeshRULE_dictValue)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(267)
		p.Match(SQLMeshLBRACE)
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

	if _la == SQLMeshSTRING_LITERAL || _la == SQLMeshIDENTIFIER {
		{
			p.SetState(268)
			p.DictEntry()
		}
		p.SetState(273)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 25, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(269)
					p.Match(SQLMeshCOMMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(270)
					p.DictEntry()
				}

			}
			p.SetState(275)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 25, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		p.SetState(277)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SQLMeshCOMMA {
			{
				p.SetState(276)
				p.Match(SQLMeshCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	}
	{
		p.SetState(281)
		p.Match(SQLMeshRBRACE)
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

// IDictEntryContext is an interface to support dynamic dispatch.
type IDictEntryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	COLON() antlr.TerminalNode
	PropValue() IPropValueContext
	STRING_LITERAL() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode

	// IsDictEntryContext differentiates from other interfaces.
	IsDictEntryContext()
}

type DictEntryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDictEntryContext() *DictEntryContext {
	var p = new(DictEntryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLMeshRULE_dictEntry
	return p
}

func InitEmptyDictEntryContext(p *DictEntryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLMeshRULE_dictEntry
}

func (*DictEntryContext) IsDictEntryContext() {}

func NewDictEntryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DictEntryContext {
	var p = new(DictEntryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLMeshRULE_dictEntry

	return p
}

func (s *DictEntryContext) GetParser() antlr.Parser { return s.parser }

func (s *DictEntryContext) COLON() antlr.TerminalNode {
	return s.GetToken(SQLMeshCOLON, 0)
}

func (s *DictEntryContext) PropValue() IPropValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPropValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPropValueContext)
}

func (s *DictEntryContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(SQLMeshSTRING_LITERAL, 0)
}

func (s *DictEntryContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(SQLMeshIDENTIFIER, 0)
}

func (s *DictEntryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DictEntryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DictEntryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.EnterDictEntry(s)
	}
}

func (s *DictEntryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.ExitDictEntry(s)
	}
}

func (s *DictEntryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLMeshVisitor:
		return t.VisitDictEntry(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLMesh) DictEntry() (localctx IDictEntryContext) {
	localctx = NewDictEntryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, SQLMeshRULE_dictEntry)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(283)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SQLMeshSTRING_LITERAL || _la == SQLMeshIDENTIFIER) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(284)
		p.Match(SQLMeshCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(285)
		p.PropValue()
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

// IMacroRefContext is an interface to support dynamic dispatch.
type IMacroRefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsMacroRefContext differentiates from other interfaces.
	IsMacroRefContext()
}

type MacroRefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMacroRefContext() *MacroRefContext {
	var p = new(MacroRefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLMeshRULE_macroRef
	return p
}

func InitEmptyMacroRefContext(p *MacroRefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLMeshRULE_macroRef
}

func (*MacroRefContext) IsMacroRefContext() {}

func NewMacroRefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MacroRefContext {
	var p = new(MacroRefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLMeshRULE_macroRef

	return p
}

func (s *MacroRefContext) GetParser() antlr.Parser { return s.parser }

func (s *MacroRefContext) CopyAll(ctx *MacroRefContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *MacroRefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MacroRefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type MacroIdentRefContext struct {
	MacroRefContext
}

func NewMacroIdentRefContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MacroIdentRefContext {
	var p = new(MacroIdentRefContext)

	InitEmptyMacroRefContext(&p.MacroRefContext)
	p.parser = parser
	p.CopyAll(ctx.(*MacroRefContext))

	return p
}

func (s *MacroIdentRefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MacroIdentRefContext) MACRO_IDENT() antlr.TerminalNode {
	return s.GetToken(SQLMeshMACRO_IDENT, 0)
}

func (s *MacroIdentRefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.EnterMacroIdentRef(s)
	}
}

func (s *MacroIdentRefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.ExitMacroIdentRef(s)
	}
}

func (s *MacroIdentRefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLMeshVisitor:
		return t.VisitMacroIdentRef(s)

	default:
		return t.VisitChildren(s)
	}
}

type MacroBracedRefContext struct {
	MacroRefContext
}

func NewMacroBracedRefContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MacroBracedRefContext {
	var p = new(MacroBracedRefContext)

	InitEmptyMacroRefContext(&p.MacroRefContext)
	p.parser = parser
	p.CopyAll(ctx.(*MacroRefContext))

	return p
}

func (s *MacroBracedRefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MacroBracedRefContext) MACRO_BRACED() antlr.TerminalNode {
	return s.GetToken(SQLMeshMACRO_BRACED, 0)
}

func (s *MacroBracedRefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.EnterMacroBracedRef(s)
	}
}

func (s *MacroBracedRefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.ExitMacroBracedRef(s)
	}
}

func (s *MacroBracedRefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLMeshVisitor:
		return t.VisitMacroBracedRef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLMesh) MacroRef() (localctx IMacroRefContext) {
	localctx = NewMacroRefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, SQLMeshRULE_macroRef)
	p.SetState(289)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SQLMeshMACRO_BRACED:
		localctx = NewMacroBracedRefContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(287)
			p.Match(SQLMeshMACRO_BRACED)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SQLMeshMACRO_IDENT:
		localctx = NewMacroIdentRefContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(288)
			p.Match(SQLMeshMACRO_IDENT)
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
	p.RuleIndex = SQLMeshRULE_sqlBody
	return p
}

func InitEmptySqlBodyContext(p *SqlBodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SQLMeshRULE_sqlBody
}

func (*SqlBodyContext) IsSqlBodyContext() {}

func NewSqlBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SqlBodyContext {
	var p = new(SqlBodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SQLMeshRULE_sqlBody

	return p
}

func (s *SqlBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *SqlBodyContext) AllSQL_LINE() []antlr.TerminalNode {
	return s.GetTokens(SQLMeshSQL_LINE)
}

func (s *SqlBodyContext) SQL_LINE(i int) antlr.TerminalNode {
	return s.GetToken(SQLMeshSQL_LINE, i)
}

func (s *SqlBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SqlBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SqlBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.EnterSqlBody(s)
	}
}

func (s *SqlBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SQLMeshListener); ok {
		listenerT.ExitSqlBody(s)
	}
}

func (s *SqlBodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SQLMeshVisitor:
		return t.VisitSqlBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SQLMesh) SqlBody() (localctx ISqlBodyContext) {
	localctx = NewSqlBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, SQLMeshRULE_sqlBody)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(294)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SQLMeshSQL_LINE {
		{
			p.SetState(291)
			p.Match(SQLMeshSQL_LINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(296)
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
