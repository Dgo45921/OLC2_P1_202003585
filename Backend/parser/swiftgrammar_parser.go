// Code generated from SwiftGrammar.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // SwiftGrammar
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

import "PY1/interfaces"
import "PY1/environment"
import "PY1/expressions"
import "PY1/instructions"
import "strings"

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type SwiftGrammarParser struct {
	*antlr.BaseParser
}

var SwiftGrammarParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func swiftgrammarParserInit() {
	staticData := &SwiftGrammarParserStaticData
	staticData.LiteralNames = []string{
		"", "'Int'", "'Float'", "'Bool'", "'String'", "'Character'", "'true'",
		"'false'", "'print'", "'if'", "'else'", "'while'", "'var'", "'let'",
		"'nil'", "'break'", "'continue'", "'append'", "'removeLast'", "'at'",
		"'remove'", "'isEmpty'", "'count'", "'switch'", "'case'", "'default'",
		"'for'", "'in'", "'repeating'", "'struct'", "", "", "", "'+='", "'-='",
		"'!='", "'=='", "'!'", "'||'", "'&&'", "'='", "'>='", "'<='", "'>'",
		"'<'", "'*'", "'/'", "'+'", "'-'", "'('", "')'", "'{'", "'}'", "'['",
		"']'", "';'", "':'", "'%'", "','", "'?'", "'.'",
	}
	staticData.SymbolicNames = []string{
		"", "RINT", "RFLOAT", "RBOOL", "RSTRING", "RCHARACTER", "RTRUE", "RFALSE",
		"RPRINT", "RIF", "RELSE", "RWHILE", "RVAR", "RLET", "RNIL", "RBREAK",
		"RCONTINUE", "RAPPEND", "RREMOVELAST", "RRAT", "RREMOVEAT", "RISEMPTY",
		"RCOUNT", "RSWITCH", "RCASE", "RDEFAULT", "RFOR", "RIN", "RREPEATING",
		"RSTRUCT", "NUMBER", "STRING", "ID", "UNARYPLUS", "UNARYMINUS", "DIF",
		"IG_IG", "NOT", "OR", "AND", "IG", "MAY_IG", "MEN_IG", "MAYOR", "MENOR",
		"MUL", "DIV", "ADD", "SUB", "PARIZQ", "PARDER", "LLAVEIZQ", "LLAVEDER",
		"OBRA", "CBRA", "PTOCOMA", "DOSPTOS", "MODULE", "COMA", "QM", "PTO",
		"WHITESPACE", "COMMENT", "LINE_COMMENT",
	}
	staticData.RuleNames = []string{
		"s", "block", "arguments", "argument", "instruction", "structinstruction",
		"structblock", "structdef", "vectormodification", "forloop", "range",
		"removeatvec", "appendvec", "removelastvec", "vecdec", "breakstatement",
		"continuestatement", "switchstatement", "caselist", "case", "defaultstatement",
		"ifstmt", "eliflist", "elif", "elsestament", "printstmt", "while_statement",
		"vardec", "constdec", "asignation", "unarysum", "unarysub", "isemptyvec",
		"countvec", "vectoraccess", "indexesList", "vecac", "matrix_type", "repeatingvector",
		"manualdef", "manualmatrixdef", "values2", "decmatrix", "attrlist",
		"attr", "structaccess", "structexp", "keyvaluelist", "keyvalue", "expr",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 63, 748, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46, 2, 47,
		7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 5, 1,
		106, 8, 1, 10, 1, 12, 1, 109, 9, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 122, 8, 2, 1, 3, 1, 3, 1, 3, 1, 4,
		1, 4, 3, 4, 129, 8, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 135, 8, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 3, 4, 141, 8, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 147, 8, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 153, 8, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4,
		159, 8, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 165, 8, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 3, 4, 171, 8, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 177, 8, 4, 1, 4, 1, 4,
		1, 4, 1, 4, 3, 4, 183, 8, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 189, 8, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 3, 4, 195, 8, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 201,
		8, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 207, 8, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 4, 3, 4, 226, 8, 4, 1, 5, 1, 5, 3, 5, 230, 8, 5, 1, 5, 1, 5, 1, 5, 1,
		5, 3, 5, 236, 8, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 242, 8, 5, 1, 5, 1, 5,
		1, 5, 1, 5, 3, 5, 248, 8, 5, 1, 5, 1, 5, 3, 5, 252, 8, 5, 1, 6, 5, 6, 255,
		8, 6, 10, 6, 12, 6, 258, 9, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 3, 9, 293, 8, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1,
		11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13,
		1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1,
		14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14,
		1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 3, 14, 350, 8, 14, 1,
		15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17,
		1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 3,
		17, 373, 8, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18,
		3, 18, 383, 8, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1,
		20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21,
		1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1,
		21, 3, 21, 414, 8, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22,
		1, 22, 3, 22, 424, 8, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1,
		23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25,
		1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1,
		27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27,
		1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 3, 27, 473,
		8, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1,
		28, 1, 28, 1, 28, 1, 28, 1, 28, 3, 28, 489, 8, 28, 1, 29, 1, 29, 1, 29,
		1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 31, 1, 31, 1, 31, 1,
		31, 1, 31, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 1, 33,
		1, 33, 1, 34, 1, 34, 1, 34, 1, 34, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1,
		35, 1, 35, 3, 35, 527, 8, 35, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 37,
		1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 3, 37, 543, 8,
		37, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38,
		1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1,
		38, 1, 38, 1, 38, 1, 38, 3, 38, 569, 8, 38, 1, 39, 1, 39, 1, 39, 1, 40,
		1, 40, 1, 40, 1, 40, 1, 40, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1,
		41, 3, 41, 586, 8, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 5, 41, 593, 8,
		41, 10, 41, 12, 41, 596, 9, 41, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42,
		1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1,
		42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42,
		1, 42, 3, 42, 626, 8, 42, 1, 43, 1, 43, 1, 43, 1, 43, 1, 43, 1, 43, 1,
		43, 1, 43, 1, 43, 3, 43, 637, 8, 43, 1, 44, 1, 44, 1, 44, 1, 45, 1, 45,
		1, 45, 1, 45, 1, 45, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 47, 1,
		47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 3, 47, 662, 8, 47,
		1, 48, 1, 48, 1, 48, 1, 48, 1, 48, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1,
		49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49,
		1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1,
		49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49,
		1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 3, 49, 711, 8, 49, 1, 49, 1, 49, 1,
		49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49,
		1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1,
		49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 5, 49, 743, 8, 49, 10, 49,
		12, 49, 746, 9, 49, 1, 49, 0, 2, 82, 98, 50, 0, 2, 4, 6, 8, 10, 12, 14,
		16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50,
		52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 86,
		88, 90, 92, 94, 96, 98, 0, 6, 1, 0, 1, 5, 2, 0, 37, 37, 48, 48, 2, 0, 45,
		46, 57, 57, 1, 0, 47, 48, 1, 0, 41, 44, 1, 0, 35, 36, 780, 0, 100, 1, 0,
		0, 0, 2, 107, 1, 0, 0, 0, 4, 121, 1, 0, 0, 0, 6, 123, 1, 0, 0, 0, 8, 225,
		1, 0, 0, 0, 10, 251, 1, 0, 0, 0, 12, 256, 1, 0, 0, 0, 14, 261, 1, 0, 0,
		0, 16, 268, 1, 0, 0, 0, 18, 292, 1, 0, 0, 0, 20, 294, 1, 0, 0, 0, 22, 301,
		1, 0, 0, 0, 24, 311, 1, 0, 0, 0, 26, 319, 1, 0, 0, 0, 28, 349, 1, 0, 0,
		0, 30, 351, 1, 0, 0, 0, 32, 354, 1, 0, 0, 0, 34, 372, 1, 0, 0, 0, 36, 382,
		1, 0, 0, 0, 38, 384, 1, 0, 0, 0, 40, 390, 1, 0, 0, 0, 42, 413, 1, 0, 0,
		0, 44, 423, 1, 0, 0, 0, 46, 425, 1, 0, 0, 0, 48, 433, 1, 0, 0, 0, 50, 439,
		1, 0, 0, 0, 52, 445, 1, 0, 0, 0, 54, 472, 1, 0, 0, 0, 56, 488, 1, 0, 0,
		0, 58, 490, 1, 0, 0, 0, 60, 495, 1, 0, 0, 0, 62, 500, 1, 0, 0, 0, 64, 505,
		1, 0, 0, 0, 66, 510, 1, 0, 0, 0, 68, 515, 1, 0, 0, 0, 70, 526, 1, 0, 0,
		0, 72, 528, 1, 0, 0, 0, 74, 542, 1, 0, 0, 0, 76, 568, 1, 0, 0, 0, 78, 570,
		1, 0, 0, 0, 80, 573, 1, 0, 0, 0, 82, 585, 1, 0, 0, 0, 84, 625, 1, 0, 0,
		0, 86, 636, 1, 0, 0, 0, 88, 638, 1, 0, 0, 0, 90, 641, 1, 0, 0, 0, 92, 646,
		1, 0, 0, 0, 94, 661, 1, 0, 0, 0, 96, 663, 1, 0, 0, 0, 98, 710, 1, 0, 0,
		0, 100, 101, 3, 2, 1, 0, 101, 102, 5, 0, 0, 1, 102, 103, 6, 0, -1, 0, 103,
		1, 1, 0, 0, 0, 104, 106, 3, 8, 4, 0, 105, 104, 1, 0, 0, 0, 106, 109, 1,
		0, 0, 0, 107, 105, 1, 0, 0, 0, 107, 108, 1, 0, 0, 0, 108, 110, 1, 0, 0,
		0, 109, 107, 1, 0, 0, 0, 110, 111, 6, 1, -1, 0, 111, 3, 1, 0, 0, 0, 112,
		113, 3, 6, 3, 0, 113, 114, 5, 58, 0, 0, 114, 115, 3, 4, 2, 0, 115, 116,
		6, 2, -1, 0, 116, 122, 1, 0, 0, 0, 117, 118, 3, 6, 3, 0, 118, 119, 6, 2,
		-1, 0, 119, 122, 1, 0, 0, 0, 120, 122, 6, 2, -1, 0, 121, 112, 1, 0, 0,
		0, 121, 117, 1, 0, 0, 0, 121, 120, 1, 0, 0, 0, 122, 5, 1, 0, 0, 0, 123,
		124, 3, 98, 49, 0, 124, 125, 6, 3, -1, 0, 125, 7, 1, 0, 0, 0, 126, 128,
		3, 50, 25, 0, 127, 129, 5, 55, 0, 0, 128, 127, 1, 0, 0, 0, 128, 129, 1,
		0, 0, 0, 129, 130, 1, 0, 0, 0, 130, 131, 6, 4, -1, 0, 131, 226, 1, 0, 0,
		0, 132, 134, 3, 84, 42, 0, 133, 135, 5, 55, 0, 0, 134, 133, 1, 0, 0, 0,
		134, 135, 1, 0, 0, 0, 135, 136, 1, 0, 0, 0, 136, 137, 6, 4, -1, 0, 137,
		226, 1, 0, 0, 0, 138, 140, 3, 28, 14, 0, 139, 141, 5, 55, 0, 0, 140, 139,
		1, 0, 0, 0, 140, 141, 1, 0, 0, 0, 141, 142, 1, 0, 0, 0, 142, 143, 6, 4,
		-1, 0, 143, 226, 1, 0, 0, 0, 144, 146, 3, 54, 27, 0, 145, 147, 5, 55, 0,
		0, 146, 145, 1, 0, 0, 0, 146, 147, 1, 0, 0, 0, 147, 148, 1, 0, 0, 0, 148,
		149, 6, 4, -1, 0, 149, 226, 1, 0, 0, 0, 150, 152, 3, 56, 28, 0, 151, 153,
		5, 55, 0, 0, 152, 151, 1, 0, 0, 0, 152, 153, 1, 0, 0, 0, 153, 154, 1, 0,
		0, 0, 154, 155, 6, 4, -1, 0, 155, 226, 1, 0, 0, 0, 156, 158, 3, 24, 12,
		0, 157, 159, 5, 55, 0, 0, 158, 157, 1, 0, 0, 0, 158, 159, 1, 0, 0, 0, 159,
		160, 1, 0, 0, 0, 160, 161, 6, 4, -1, 0, 161, 226, 1, 0, 0, 0, 162, 164,
		3, 26, 13, 0, 163, 165, 5, 55, 0, 0, 164, 163, 1, 0, 0, 0, 164, 165, 1,
		0, 0, 0, 165, 166, 1, 0, 0, 0, 166, 167, 6, 4, -1, 0, 167, 226, 1, 0, 0,
		0, 168, 170, 3, 22, 11, 0, 169, 171, 5, 55, 0, 0, 170, 169, 1, 0, 0, 0,
		170, 171, 1, 0, 0, 0, 171, 172, 1, 0, 0, 0, 172, 173, 6, 4, -1, 0, 173,
		226, 1, 0, 0, 0, 174, 176, 3, 58, 29, 0, 175, 177, 5, 55, 0, 0, 176, 175,
		1, 0, 0, 0, 176, 177, 1, 0, 0, 0, 177, 178, 1, 0, 0, 0, 178, 179, 6, 4,
		-1, 0, 179, 226, 1, 0, 0, 0, 180, 182, 3, 60, 30, 0, 181, 183, 5, 55, 0,
		0, 182, 181, 1, 0, 0, 0, 182, 183, 1, 0, 0, 0, 183, 184, 1, 0, 0, 0, 184,
		185, 6, 4, -1, 0, 185, 226, 1, 0, 0, 0, 186, 188, 3, 62, 31, 0, 187, 189,
		5, 55, 0, 0, 188, 187, 1, 0, 0, 0, 188, 189, 1, 0, 0, 0, 189, 190, 1, 0,
		0, 0, 190, 191, 6, 4, -1, 0, 191, 226, 1, 0, 0, 0, 192, 194, 3, 30, 15,
		0, 193, 195, 5, 55, 0, 0, 194, 193, 1, 0, 0, 0, 194, 195, 1, 0, 0, 0, 195,
		196, 1, 0, 0, 0, 196, 197, 6, 4, -1, 0, 197, 226, 1, 0, 0, 0, 198, 200,
		3, 32, 16, 0, 199, 201, 5, 55, 0, 0, 200, 199, 1, 0, 0, 0, 200, 201, 1,
		0, 0, 0, 201, 202, 1, 0, 0, 0, 202, 203, 6, 4, -1, 0, 203, 226, 1, 0, 0,
		0, 204, 206, 3, 16, 8, 0, 205, 207, 5, 55, 0, 0, 206, 205, 1, 0, 0, 0,
		206, 207, 1, 0, 0, 0, 207, 208, 1, 0, 0, 0, 208, 209, 6, 4, -1, 0, 209,
		226, 1, 0, 0, 0, 210, 211, 3, 42, 21, 0, 211, 212, 6, 4, -1, 0, 212, 226,
		1, 0, 0, 0, 213, 214, 3, 52, 26, 0, 214, 215, 6, 4, -1, 0, 215, 226, 1,
		0, 0, 0, 216, 217, 3, 34, 17, 0, 217, 218, 6, 4, -1, 0, 218, 226, 1, 0,
		0, 0, 219, 220, 3, 18, 9, 0, 220, 221, 6, 4, -1, 0, 221, 226, 1, 0, 0,
		0, 222, 223, 3, 14, 7, 0, 223, 224, 6, 4, -1, 0, 224, 226, 1, 0, 0, 0,
		225, 126, 1, 0, 0, 0, 225, 132, 1, 0, 0, 0, 225, 138, 1, 0, 0, 0, 225,
		144, 1, 0, 0, 0, 225, 150, 1, 0, 0, 0, 225, 156, 1, 0, 0, 0, 225, 162,
		1, 0, 0, 0, 225, 168, 1, 0, 0, 0, 225, 174, 1, 0, 0, 0, 225, 180, 1, 0,
		0, 0, 225, 186, 1, 0, 0, 0, 225, 192, 1, 0, 0, 0, 225, 198, 1, 0, 0, 0,
		225, 204, 1, 0, 0, 0, 225, 210, 1, 0, 0, 0, 225, 213, 1, 0, 0, 0, 225,
		216, 1, 0, 0, 0, 225, 219, 1, 0, 0, 0, 225, 222, 1, 0, 0, 0, 226, 9, 1,
		0, 0, 0, 227, 229, 3, 28, 14, 0, 228, 230, 5, 55, 0, 0, 229, 228, 1, 0,
		0, 0, 229, 230, 1, 0, 0, 0, 230, 231, 1, 0, 0, 0, 231, 232, 6, 5, -1, 0,
		232, 252, 1, 0, 0, 0, 233, 235, 3, 54, 27, 0, 234, 236, 5, 55, 0, 0, 235,
		234, 1, 0, 0, 0, 235, 236, 1, 0, 0, 0, 236, 237, 1, 0, 0, 0, 237, 238,
		6, 5, -1, 0, 238, 252, 1, 0, 0, 0, 239, 241, 3, 56, 28, 0, 240, 242, 5,
		55, 0, 0, 241, 240, 1, 0, 0, 0, 241, 242, 1, 0, 0, 0, 242, 243, 1, 0, 0,
		0, 243, 244, 6, 5, -1, 0, 244, 252, 1, 0, 0, 0, 245, 247, 3, 84, 42, 0,
		246, 248, 5, 55, 0, 0, 247, 246, 1, 0, 0, 0, 247, 248, 1, 0, 0, 0, 248,
		249, 1, 0, 0, 0, 249, 250, 6, 5, -1, 0, 250, 252, 1, 0, 0, 0, 251, 227,
		1, 0, 0, 0, 251, 233, 1, 0, 0, 0, 251, 239, 1, 0, 0, 0, 251, 245, 1, 0,
		0, 0, 252, 11, 1, 0, 0, 0, 253, 255, 3, 10, 5, 0, 254, 253, 1, 0, 0, 0,
		255, 258, 1, 0, 0, 0, 256, 254, 1, 0, 0, 0, 256, 257, 1, 0, 0, 0, 257,
		259, 1, 0, 0, 0, 258, 256, 1, 0, 0, 0, 259, 260, 6, 6, -1, 0, 260, 13,
		1, 0, 0, 0, 261, 262, 5, 29, 0, 0, 262, 263, 5, 32, 0, 0, 263, 264, 5,
		51, 0, 0, 264, 265, 3, 12, 6, 0, 265, 266, 5, 52, 0, 0, 266, 267, 6, 7,
		-1, 0, 267, 15, 1, 0, 0, 0, 268, 269, 5, 32, 0, 0, 269, 270, 3, 70, 35,
		0, 270, 271, 5, 40, 0, 0, 271, 272, 3, 98, 49, 0, 272, 273, 6, 8, -1, 0,
		273, 17, 1, 0, 0, 0, 274, 275, 5, 26, 0, 0, 275, 276, 5, 32, 0, 0, 276,
		277, 5, 27, 0, 0, 277, 278, 3, 98, 49, 0, 278, 279, 5, 51, 0, 0, 279, 280,
		3, 2, 1, 0, 280, 281, 5, 52, 0, 0, 281, 282, 6, 9, -1, 0, 282, 293, 1,
		0, 0, 0, 283, 284, 5, 26, 0, 0, 284, 285, 5, 32, 0, 0, 285, 286, 5, 27,
		0, 0, 286, 287, 3, 20, 10, 0, 287, 288, 5, 51, 0, 0, 288, 289, 3, 2, 1,
		0, 289, 290, 5, 52, 0, 0, 290, 291, 6, 9, -1, 0, 291, 293, 1, 0, 0, 0,
		292, 274, 1, 0, 0, 0, 292, 283, 1, 0, 0, 0, 293, 19, 1, 0, 0, 0, 294, 295,
		3, 98, 49, 0, 295, 296, 5, 60, 0, 0, 296, 297, 5, 60, 0, 0, 297, 298, 5,
		60, 0, 0, 298, 299, 3, 98, 49, 0, 299, 300, 6, 10, -1, 0, 300, 21, 1, 0,
		0, 0, 301, 302, 5, 32, 0, 0, 302, 303, 5, 60, 0, 0, 303, 304, 5, 20, 0,
		0, 304, 305, 5, 49, 0, 0, 305, 306, 5, 19, 0, 0, 306, 307, 5, 56, 0, 0,
		307, 308, 3, 98, 49, 0, 308, 309, 5, 50, 0, 0, 309, 310, 6, 11, -1, 0,
		310, 23, 1, 0, 0, 0, 311, 312, 5, 32, 0, 0, 312, 313, 5, 60, 0, 0, 313,
		314, 5, 17, 0, 0, 314, 315, 5, 49, 0, 0, 315, 316, 3, 98, 49, 0, 316, 317,
		5, 50, 0, 0, 317, 318, 6, 12, -1, 0, 318, 25, 1, 0, 0, 0, 319, 320, 5,
		32, 0, 0, 320, 321, 5, 60, 0, 0, 321, 322, 5, 18, 0, 0, 322, 323, 5, 49,
		0, 0, 323, 324, 5, 50, 0, 0, 324, 325, 6, 13, -1, 0, 325, 27, 1, 0, 0,
		0, 326, 327, 5, 12, 0, 0, 327, 328, 5, 32, 0, 0, 328, 329, 5, 56, 0, 0,
		329, 330, 5, 53, 0, 0, 330, 331, 7, 0, 0, 0, 331, 332, 5, 54, 0, 0, 332,
		333, 5, 40, 0, 0, 333, 334, 5, 53, 0, 0, 334, 335, 7, 0, 0, 0, 335, 336,
		5, 54, 0, 0, 336, 337, 5, 53, 0, 0, 337, 338, 5, 54, 0, 0, 338, 350, 6,
		14, -1, 0, 339, 340, 5, 12, 0, 0, 340, 341, 5, 32, 0, 0, 341, 342, 5, 56,
		0, 0, 342, 343, 5, 53, 0, 0, 343, 344, 7, 0, 0, 0, 344, 345, 5, 54, 0,
		0, 345, 346, 5, 40, 0, 0, 346, 347, 3, 98, 49, 0, 347, 348, 6, 14, -1,
		0, 348, 350, 1, 0, 0, 0, 349, 326, 1, 0, 0, 0, 349, 339, 1, 0, 0, 0, 350,
		29, 1, 0, 0, 0, 351, 352, 5, 15, 0, 0, 352, 353, 6, 15, -1, 0, 353, 31,
		1, 0, 0, 0, 354, 355, 5, 16, 0, 0, 355, 356, 6, 16, -1, 0, 356, 33, 1,
		0, 0, 0, 357, 358, 5, 23, 0, 0, 358, 359, 3, 98, 49, 0, 359, 360, 5, 51,
		0, 0, 360, 361, 3, 36, 18, 0, 361, 362, 5, 52, 0, 0, 362, 363, 6, 17, -1,
		0, 363, 373, 1, 0, 0, 0, 364, 365, 5, 23, 0, 0, 365, 366, 3, 98, 49, 0,
		366, 367, 5, 51, 0, 0, 367, 368, 3, 36, 18, 0, 368, 369, 3, 40, 20, 0,
		369, 370, 5, 52, 0, 0, 370, 371, 6, 17, -1, 0, 371, 373, 1, 0, 0, 0, 372,
		357, 1, 0, 0, 0, 372, 364, 1, 0, 0, 0, 373, 35, 1, 0, 0, 0, 374, 375, 3,
		38, 19, 0, 375, 376, 3, 36, 18, 0, 376, 377, 6, 18, -1, 0, 377, 383, 1,
		0, 0, 0, 378, 379, 3, 38, 19, 0, 379, 380, 6, 18, -1, 0, 380, 383, 1, 0,
		0, 0, 381, 383, 6, 18, -1, 0, 382, 374, 1, 0, 0, 0, 382, 378, 1, 0, 0,
		0, 382, 381, 1, 0, 0, 0, 383, 37, 1, 0, 0, 0, 384, 385, 5, 24, 0, 0, 385,
		386, 3, 98, 49, 0, 386, 387, 5, 56, 0, 0, 387, 388, 3, 2, 1, 0, 388, 389,
		6, 19, -1, 0, 389, 39, 1, 0, 0, 0, 390, 391, 5, 25, 0, 0, 391, 392, 5,
		56, 0, 0, 392, 393, 3, 2, 1, 0, 393, 394, 5, 52, 0, 0, 394, 395, 6, 20,
		-1, 0, 395, 41, 1, 0, 0, 0, 396, 397, 5, 9, 0, 0, 397, 398, 3, 98, 49,
		0, 398, 399, 5, 51, 0, 0, 399, 400, 3, 2, 1, 0, 400, 401, 5, 52, 0, 0,
		401, 402, 3, 44, 22, 0, 402, 403, 6, 21, -1, 0, 403, 414, 1, 0, 0, 0, 404,
		405, 5, 9, 0, 0, 405, 406, 3, 98, 49, 0, 406, 407, 5, 51, 0, 0, 407, 408,
		3, 2, 1, 0, 408, 409, 5, 52, 0, 0, 409, 410, 3, 44, 22, 0, 410, 411, 3,
		48, 24, 0, 411, 412, 6, 21, -1, 0, 412, 414, 1, 0, 0, 0, 413, 396, 1, 0,
		0, 0, 413, 404, 1, 0, 0, 0, 414, 43, 1, 0, 0, 0, 415, 416, 3, 46, 23, 0,
		416, 417, 3, 44, 22, 0, 417, 418, 6, 22, -1, 0, 418, 424, 1, 0, 0, 0, 419,
		420, 3, 46, 23, 0, 420, 421, 6, 22, -1, 0, 421, 424, 1, 0, 0, 0, 422, 424,
		6, 22, -1, 0, 423, 415, 1, 0, 0, 0, 423, 419, 1, 0, 0, 0, 423, 422, 1,
		0, 0, 0, 424, 45, 1, 0, 0, 0, 425, 426, 5, 10, 0, 0, 426, 427, 5, 9, 0,
		0, 427, 428, 3, 98, 49, 0, 428, 429, 5, 51, 0, 0, 429, 430, 3, 2, 1, 0,
		430, 431, 5, 52, 0, 0, 431, 432, 6, 23, -1, 0, 432, 47, 1, 0, 0, 0, 433,
		434, 5, 10, 0, 0, 434, 435, 5, 51, 0, 0, 435, 436, 3, 2, 1, 0, 436, 437,
		5, 52, 0, 0, 437, 438, 6, 24, -1, 0, 438, 49, 1, 0, 0, 0, 439, 440, 5,
		8, 0, 0, 440, 441, 5, 49, 0, 0, 441, 442, 3, 4, 2, 0, 442, 443, 5, 50,
		0, 0, 443, 444, 6, 25, -1, 0, 444, 51, 1, 0, 0, 0, 445, 446, 5, 11, 0,
		0, 446, 447, 3, 98, 49, 0, 447, 448, 5, 51, 0, 0, 448, 449, 3, 2, 1, 0,
		449, 450, 5, 52, 0, 0, 450, 451, 6, 26, -1, 0, 451, 53, 1, 0, 0, 0, 452,
		453, 5, 12, 0, 0, 453, 454, 5, 32, 0, 0, 454, 455, 5, 56, 0, 0, 455, 456,
		7, 0, 0, 0, 456, 457, 5, 40, 0, 0, 457, 458, 3, 98, 49, 0, 458, 459, 6,
		27, -1, 0, 459, 473, 1, 0, 0, 0, 460, 461, 5, 12, 0, 0, 461, 462, 5, 32,
		0, 0, 462, 463, 5, 40, 0, 0, 463, 464, 3, 98, 49, 0, 464, 465, 6, 27, -1,
		0, 465, 473, 1, 0, 0, 0, 466, 467, 5, 12, 0, 0, 467, 468, 5, 32, 0, 0,
		468, 469, 5, 56, 0, 0, 469, 470, 7, 0, 0, 0, 470, 471, 5, 59, 0, 0, 471,
		473, 6, 27, -1, 0, 472, 452, 1, 0, 0, 0, 472, 460, 1, 0, 0, 0, 472, 466,
		1, 0, 0, 0, 473, 55, 1, 0, 0, 0, 474, 475, 5, 13, 0, 0, 475, 476, 5, 32,
		0, 0, 476, 477, 5, 56, 0, 0, 477, 478, 7, 0, 0, 0, 478, 479, 5, 40, 0,
		0, 479, 480, 3, 98, 49, 0, 480, 481, 6, 28, -1, 0, 481, 489, 1, 0, 0, 0,
		482, 483, 5, 13, 0, 0, 483, 484, 5, 32, 0, 0, 484, 485, 5, 40, 0, 0, 485,
		486, 3, 98, 49, 0, 486, 487, 6, 28, -1, 0, 487, 489, 1, 0, 0, 0, 488, 474,
		1, 0, 0, 0, 488, 482, 1, 0, 0, 0, 489, 57, 1, 0, 0, 0, 490, 491, 5, 32,
		0, 0, 491, 492, 5, 40, 0, 0, 492, 493, 3, 98, 49, 0, 493, 494, 6, 29, -1,
		0, 494, 59, 1, 0, 0, 0, 495, 496, 5, 32, 0, 0, 496, 497, 5, 33, 0, 0, 497,
		498, 3, 98, 49, 0, 498, 499, 6, 30, -1, 0, 499, 61, 1, 0, 0, 0, 500, 501,
		5, 32, 0, 0, 501, 502, 5, 34, 0, 0, 502, 503, 3, 98, 49, 0, 503, 504, 6,
		31, -1, 0, 504, 63, 1, 0, 0, 0, 505, 506, 5, 32, 0, 0, 506, 507, 5, 60,
		0, 0, 507, 508, 5, 21, 0, 0, 508, 509, 6, 32, -1, 0, 509, 65, 1, 0, 0,
		0, 510, 511, 5, 32, 0, 0, 511, 512, 5, 60, 0, 0, 512, 513, 5, 22, 0, 0,
		513, 514, 6, 33, -1, 0, 514, 67, 1, 0, 0, 0, 515, 516, 5, 32, 0, 0, 516,
		517, 3, 70, 35, 0, 517, 518, 6, 34, -1, 0, 518, 69, 1, 0, 0, 0, 519, 520,
		3, 72, 36, 0, 520, 521, 3, 70, 35, 0, 521, 522, 6, 35, -1, 0, 522, 527,
		1, 0, 0, 0, 523, 524, 3, 72, 36, 0, 524, 525, 6, 35, -1, 0, 525, 527, 1,
		0, 0, 0, 526, 519, 1, 0, 0, 0, 526, 523, 1, 0, 0, 0, 527, 71, 1, 0, 0,
		0, 528, 529, 5, 53, 0, 0, 529, 530, 3, 98, 49, 0, 530, 531, 5, 54, 0, 0,
		531, 532, 6, 36, -1, 0, 532, 73, 1, 0, 0, 0, 533, 534, 5, 53, 0, 0, 534,
		535, 7, 0, 0, 0, 535, 536, 5, 54, 0, 0, 536, 543, 6, 37, -1, 0, 537, 538,
		5, 53, 0, 0, 538, 539, 3, 74, 37, 0, 539, 540, 5, 54, 0, 0, 540, 541, 6,
		37, -1, 0, 541, 543, 1, 0, 0, 0, 542, 533, 1, 0, 0, 0, 542, 537, 1, 0,
		0, 0, 543, 75, 1, 0, 0, 0, 544, 545, 3, 74, 37, 0, 545, 546, 5, 49, 0,
		0, 546, 547, 5, 28, 0, 0, 547, 548, 5, 56, 0, 0, 548, 549, 3, 76, 38, 0,
		549, 550, 5, 58, 0, 0, 550, 551, 5, 22, 0, 0, 551, 552, 5, 56, 0, 0, 552,
		553, 3, 98, 49, 0, 553, 554, 5, 50, 0, 0, 554, 555, 6, 38, -1, 0, 555,
		569, 1, 0, 0, 0, 556, 557, 3, 74, 37, 0, 557, 558, 5, 49, 0, 0, 558, 559,
		5, 28, 0, 0, 559, 560, 5, 56, 0, 0, 560, 561, 3, 98, 49, 0, 561, 562, 5,
		58, 0, 0, 562, 563, 5, 22, 0, 0, 563, 564, 5, 56, 0, 0, 564, 565, 3, 98,
		49, 0, 565, 566, 5, 50, 0, 0, 566, 567, 6, 38, -1, 0, 567, 569, 1, 0, 0,
		0, 568, 544, 1, 0, 0, 0, 568, 556, 1, 0, 0, 0, 569, 77, 1, 0, 0, 0, 570,
		571, 3, 80, 40, 0, 571, 572, 6, 39, -1, 0, 572, 79, 1, 0, 0, 0, 573, 574,
		5, 53, 0, 0, 574, 575, 3, 82, 41, 0, 575, 576, 5, 54, 0, 0, 576, 577, 6,
		40, -1, 0, 577, 81, 1, 0, 0, 0, 578, 579, 6, 41, -1, 0, 579, 580, 3, 80,
		40, 0, 580, 581, 6, 41, -1, 0, 581, 586, 1, 0, 0, 0, 582, 583, 3, 4, 2,
		0, 583, 584, 6, 41, -1, 0, 584, 586, 1, 0, 0, 0, 585, 578, 1, 0, 0, 0,
		585, 582, 1, 0, 0, 0, 586, 594, 1, 0, 0, 0, 587, 588, 10, 3, 0, 0, 588,
		589, 5, 58, 0, 0, 589, 590, 3, 80, 40, 0, 590, 591, 6, 41, -1, 0, 591,
		593, 1, 0, 0, 0, 592, 587, 1, 0, 0, 0, 593, 596, 1, 0, 0, 0, 594, 592,
		1, 0, 0, 0, 594, 595, 1, 0, 0, 0, 595, 83, 1, 0, 0, 0, 596, 594, 1, 0,
		0, 0, 597, 598, 5, 12, 0, 0, 598, 599, 5, 32, 0, 0, 599, 600, 5, 40, 0,
		0, 600, 601, 3, 78, 39, 0, 601, 602, 6, 42, -1, 0, 602, 626, 1, 0, 0, 0,
		603, 604, 5, 12, 0, 0, 604, 605, 5, 32, 0, 0, 605, 606, 5, 56, 0, 0, 606,
		607, 3, 74, 37, 0, 607, 608, 5, 40, 0, 0, 608, 609, 3, 78, 39, 0, 609,
		610, 6, 42, -1, 0, 610, 626, 1, 0, 0, 0, 611, 612, 5, 12, 0, 0, 612, 613,
		5, 32, 0, 0, 613, 614, 5, 40, 0, 0, 614, 615, 3, 76, 38, 0, 615, 616, 6,
		42, -1, 0, 616, 626, 1, 0, 0, 0, 617, 618, 5, 12, 0, 0, 618, 619, 5, 32,
		0, 0, 619, 620, 5, 56, 0, 0, 620, 621, 3, 74, 37, 0, 621, 622, 5, 40, 0,
		0, 622, 623, 3, 76, 38, 0, 623, 624, 6, 42, -1, 0, 624, 626, 1, 0, 0, 0,
		625, 597, 1, 0, 0, 0, 625, 603, 1, 0, 0, 0, 625, 611, 1, 0, 0, 0, 625,
		617, 1, 0, 0, 0, 626, 85, 1, 0, 0, 0, 627, 628, 3, 88, 44, 0, 628, 629,
		6, 43, -1, 0, 629, 630, 5, 60, 0, 0, 630, 631, 3, 86, 43, 0, 631, 632,
		6, 43, -1, 0, 632, 637, 1, 0, 0, 0, 633, 634, 3, 88, 44, 0, 634, 635, 6,
		43, -1, 0, 635, 637, 1, 0, 0, 0, 636, 627, 1, 0, 0, 0, 636, 633, 1, 0,
		0, 0, 637, 87, 1, 0, 0, 0, 638, 639, 5, 32, 0, 0, 639, 640, 6, 44, -1,
		0, 640, 89, 1, 0, 0, 0, 641, 642, 5, 32, 0, 0, 642, 643, 5, 60, 0, 0, 643,
		644, 3, 86, 43, 0, 644, 645, 6, 45, -1, 0, 645, 91, 1, 0, 0, 0, 646, 647,
		5, 32, 0, 0, 647, 648, 5, 49, 0, 0, 648, 649, 3, 94, 47, 0, 649, 650, 5,
		50, 0, 0, 650, 651, 6, 46, -1, 0, 651, 93, 1, 0, 0, 0, 652, 653, 3, 96,
		48, 0, 653, 654, 6, 47, -1, 0, 654, 655, 5, 58, 0, 0, 655, 656, 3, 94,
		47, 0, 656, 657, 6, 47, -1, 0, 657, 662, 1, 0, 0, 0, 658, 659, 3, 96, 48,
		0, 659, 660, 6, 47, -1, 0, 660, 662, 1, 0, 0, 0, 661, 652, 1, 0, 0, 0,
		661, 658, 1, 0, 0, 0, 662, 95, 1, 0, 0, 0, 663, 664, 5, 32, 0, 0, 664,
		665, 5, 56, 0, 0, 665, 666, 3, 98, 49, 0, 666, 667, 6, 48, -1, 0, 667,
		97, 1, 0, 0, 0, 668, 669, 6, 49, -1, 0, 669, 670, 5, 49, 0, 0, 670, 671,
		3, 98, 49, 0, 671, 672, 5, 50, 0, 0, 672, 673, 6, 49, -1, 0, 673, 711,
		1, 0, 0, 0, 674, 675, 7, 1, 0, 0, 675, 676, 3, 98, 49, 19, 676, 677, 6,
		49, -1, 0, 677, 711, 1, 0, 0, 0, 678, 679, 3, 92, 46, 0, 679, 680, 6, 49,
		-1, 0, 680, 711, 1, 0, 0, 0, 681, 682, 3, 90, 45, 0, 682, 683, 6, 49, -1,
		0, 683, 711, 1, 0, 0, 0, 684, 685, 3, 64, 32, 0, 685, 686, 6, 49, -1, 0,
		686, 711, 1, 0, 0, 0, 687, 688, 3, 66, 33, 0, 688, 689, 6, 49, -1, 0, 689,
		711, 1, 0, 0, 0, 690, 691, 3, 68, 34, 0, 691, 692, 6, 49, -1, 0, 692, 711,
		1, 0, 0, 0, 693, 694, 5, 53, 0, 0, 694, 695, 3, 4, 2, 0, 695, 696, 5, 54,
		0, 0, 696, 697, 6, 49, -1, 0, 697, 711, 1, 0, 0, 0, 698, 699, 5, 32, 0,
		0, 699, 711, 6, 49, -1, 0, 700, 701, 5, 30, 0, 0, 701, 711, 6, 49, -1,
		0, 702, 703, 5, 31, 0, 0, 703, 711, 6, 49, -1, 0, 704, 705, 5, 6, 0, 0,
		705, 711, 6, 49, -1, 0, 706, 707, 5, 7, 0, 0, 707, 711, 6, 49, -1, 0, 708,
		709, 5, 14, 0, 0, 709, 711, 6, 49, -1, 0, 710, 668, 1, 0, 0, 0, 710, 674,
		1, 0, 0, 0, 710, 678, 1, 0, 0, 0, 710, 681, 1, 0, 0, 0, 710, 684, 1, 0,
		0, 0, 710, 687, 1, 0, 0, 0, 710, 690, 1, 0, 0, 0, 710, 693, 1, 0, 0, 0,
		710, 698, 1, 0, 0, 0, 710, 700, 1, 0, 0, 0, 710, 702, 1, 0, 0, 0, 710,
		704, 1, 0, 0, 0, 710, 706, 1, 0, 0, 0, 710, 708, 1, 0, 0, 0, 711, 744,
		1, 0, 0, 0, 712, 713, 10, 18, 0, 0, 713, 714, 7, 2, 0, 0, 714, 715, 3,
		98, 49, 19, 715, 716, 6, 49, -1, 0, 716, 743, 1, 0, 0, 0, 717, 718, 10,
		17, 0, 0, 718, 719, 7, 3, 0, 0, 719, 720, 3, 98, 49, 18, 720, 721, 6, 49,
		-1, 0, 721, 743, 1, 0, 0, 0, 722, 723, 10, 16, 0, 0, 723, 724, 7, 4, 0,
		0, 724, 725, 3, 98, 49, 17, 725, 726, 6, 49, -1, 0, 726, 743, 1, 0, 0,
		0, 727, 728, 10, 15, 0, 0, 728, 729, 7, 5, 0, 0, 729, 730, 3, 98, 49, 16,
		730, 731, 6, 49, -1, 0, 731, 743, 1, 0, 0, 0, 732, 733, 10, 14, 0, 0, 733,
		734, 5, 39, 0, 0, 734, 735, 3, 98, 49, 15, 735, 736, 6, 49, -1, 0, 736,
		743, 1, 0, 0, 0, 737, 738, 10, 13, 0, 0, 738, 739, 5, 38, 0, 0, 739, 740,
		3, 98, 49, 14, 740, 741, 6, 49, -1, 0, 741, 743, 1, 0, 0, 0, 742, 712,
		1, 0, 0, 0, 742, 717, 1, 0, 0, 0, 742, 722, 1, 0, 0, 0, 742, 727, 1, 0,
		0, 0, 742, 732, 1, 0, 0, 0, 742, 737, 1, 0, 0, 0, 743, 746, 1, 0, 0, 0,
		744, 742, 1, 0, 0, 0, 744, 745, 1, 0, 0, 0, 745, 99, 1, 0, 0, 0, 746, 744,
		1, 0, 0, 0, 42, 107, 121, 128, 134, 140, 146, 152, 158, 164, 170, 176,
		182, 188, 194, 200, 206, 225, 229, 235, 241, 247, 251, 256, 292, 349, 372,
		382, 413, 423, 472, 488, 526, 542, 568, 585, 594, 625, 636, 661, 710, 742,
		744,
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

// SwiftGrammarParserInit initializes any static state used to implement SwiftGrammarParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewSwiftGrammarParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func SwiftGrammarParserInit() {
	staticData := &SwiftGrammarParserStaticData
	staticData.once.Do(swiftgrammarParserInit)
}

// NewSwiftGrammarParser produces a new parser instance for the optional input antlr.TokenStream.
func NewSwiftGrammarParser(input antlr.TokenStream) *SwiftGrammarParser {
	SwiftGrammarParserInit()
	this := new(SwiftGrammarParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &SwiftGrammarParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "SwiftGrammar.g4"

	return this
}

// SwiftGrammarParser tokens.
const (
	SwiftGrammarParserEOF          = antlr.TokenEOF
	SwiftGrammarParserRINT         = 1
	SwiftGrammarParserRFLOAT       = 2
	SwiftGrammarParserRBOOL        = 3
	SwiftGrammarParserRSTRING      = 4
	SwiftGrammarParserRCHARACTER   = 5
	SwiftGrammarParserRTRUE        = 6
	SwiftGrammarParserRFALSE       = 7
	SwiftGrammarParserRPRINT       = 8
	SwiftGrammarParserRIF          = 9
	SwiftGrammarParserRELSE        = 10
	SwiftGrammarParserRWHILE       = 11
	SwiftGrammarParserRVAR         = 12
	SwiftGrammarParserRLET         = 13
	SwiftGrammarParserRNIL         = 14
	SwiftGrammarParserRBREAK       = 15
	SwiftGrammarParserRCONTINUE    = 16
	SwiftGrammarParserRAPPEND      = 17
	SwiftGrammarParserRREMOVELAST  = 18
	SwiftGrammarParserRRAT         = 19
	SwiftGrammarParserRREMOVEAT    = 20
	SwiftGrammarParserRISEMPTY     = 21
	SwiftGrammarParserRCOUNT       = 22
	SwiftGrammarParserRSWITCH      = 23
	SwiftGrammarParserRCASE        = 24
	SwiftGrammarParserRDEFAULT     = 25
	SwiftGrammarParserRFOR         = 26
	SwiftGrammarParserRIN          = 27
	SwiftGrammarParserRREPEATING   = 28
	SwiftGrammarParserRSTRUCT      = 29
	SwiftGrammarParserNUMBER       = 30
	SwiftGrammarParserSTRING       = 31
	SwiftGrammarParserID           = 32
	SwiftGrammarParserUNARYPLUS    = 33
	SwiftGrammarParserUNARYMINUS   = 34
	SwiftGrammarParserDIF          = 35
	SwiftGrammarParserIG_IG        = 36
	SwiftGrammarParserNOT          = 37
	SwiftGrammarParserOR           = 38
	SwiftGrammarParserAND          = 39
	SwiftGrammarParserIG           = 40
	SwiftGrammarParserMAY_IG       = 41
	SwiftGrammarParserMEN_IG       = 42
	SwiftGrammarParserMAYOR        = 43
	SwiftGrammarParserMENOR        = 44
	SwiftGrammarParserMUL          = 45
	SwiftGrammarParserDIV          = 46
	SwiftGrammarParserADD          = 47
	SwiftGrammarParserSUB          = 48
	SwiftGrammarParserPARIZQ       = 49
	SwiftGrammarParserPARDER       = 50
	SwiftGrammarParserLLAVEIZQ     = 51
	SwiftGrammarParserLLAVEDER     = 52
	SwiftGrammarParserOBRA         = 53
	SwiftGrammarParserCBRA         = 54
	SwiftGrammarParserPTOCOMA      = 55
	SwiftGrammarParserDOSPTOS      = 56
	SwiftGrammarParserMODULE       = 57
	SwiftGrammarParserCOMA         = 58
	SwiftGrammarParserQM           = 59
	SwiftGrammarParserPTO          = 60
	SwiftGrammarParserWHITESPACE   = 61
	SwiftGrammarParserCOMMENT      = 62
	SwiftGrammarParserLINE_COMMENT = 63
)

// SwiftGrammarParser rules.
const (
	SwiftGrammarParserRULE_s                  = 0
	SwiftGrammarParserRULE_block              = 1
	SwiftGrammarParserRULE_arguments          = 2
	SwiftGrammarParserRULE_argument           = 3
	SwiftGrammarParserRULE_instruction        = 4
	SwiftGrammarParserRULE_structinstruction  = 5
	SwiftGrammarParserRULE_structblock        = 6
	SwiftGrammarParserRULE_structdef          = 7
	SwiftGrammarParserRULE_vectormodification = 8
	SwiftGrammarParserRULE_forloop            = 9
	SwiftGrammarParserRULE_range              = 10
	SwiftGrammarParserRULE_removeatvec        = 11
	SwiftGrammarParserRULE_appendvec          = 12
	SwiftGrammarParserRULE_removelastvec      = 13
	SwiftGrammarParserRULE_vecdec             = 14
	SwiftGrammarParserRULE_breakstatement     = 15
	SwiftGrammarParserRULE_continuestatement  = 16
	SwiftGrammarParserRULE_switchstatement    = 17
	SwiftGrammarParserRULE_caselist           = 18
	SwiftGrammarParserRULE_case               = 19
	SwiftGrammarParserRULE_defaultstatement   = 20
	SwiftGrammarParserRULE_ifstmt             = 21
	SwiftGrammarParserRULE_eliflist           = 22
	SwiftGrammarParserRULE_elif               = 23
	SwiftGrammarParserRULE_elsestament        = 24
	SwiftGrammarParserRULE_printstmt          = 25
	SwiftGrammarParserRULE_while_statement    = 26
	SwiftGrammarParserRULE_vardec             = 27
	SwiftGrammarParserRULE_constdec           = 28
	SwiftGrammarParserRULE_asignation         = 29
	SwiftGrammarParserRULE_unarysum           = 30
	SwiftGrammarParserRULE_unarysub           = 31
	SwiftGrammarParserRULE_isemptyvec         = 32
	SwiftGrammarParserRULE_countvec           = 33
	SwiftGrammarParserRULE_vectoraccess       = 34
	SwiftGrammarParserRULE_indexesList        = 35
	SwiftGrammarParserRULE_vecac              = 36
	SwiftGrammarParserRULE_matrix_type        = 37
	SwiftGrammarParserRULE_repeatingvector    = 38
	SwiftGrammarParserRULE_manualdef          = 39
	SwiftGrammarParserRULE_manualmatrixdef    = 40
	SwiftGrammarParserRULE_values2            = 41
	SwiftGrammarParserRULE_decmatrix          = 42
	SwiftGrammarParserRULE_attrlist           = 43
	SwiftGrammarParserRULE_attr               = 44
	SwiftGrammarParserRULE_structaccess       = 45
	SwiftGrammarParserRULE_structexp          = 46
	SwiftGrammarParserRULE_keyvaluelist       = 47
	SwiftGrammarParserRULE_keyvalue           = 48
	SwiftGrammarParserRULE_expr               = 49
)

// ISContext is an interface to support dynamic dispatch.
type ISContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_block returns the _block rule contexts.
	Get_block() IBlockContext

	// Set_block sets the _block rule contexts.
	Set_block(IBlockContext)

	// GetCode returns the code attribute.
	GetCode() []interface{}

	// SetCode sets the code attribute.
	SetCode([]interface{})

	// Getter signatures
	Block() IBlockContext
	EOF() antlr.TerminalNode

	// IsSContext differentiates from other interfaces.
	IsSContext()
}

type SContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	code   []interface{}
	_block IBlockContext
}

func NewEmptySContext() *SContext {
	var p = new(SContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_s
	return p
}

func InitEmptySContext(p *SContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_s
}

func (*SContext) IsSContext() {}

func NewSContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SContext {
	var p = new(SContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_s

	return p
}

func (s *SContext) GetParser() antlr.Parser { return s.parser }

func (s *SContext) Get_block() IBlockContext { return s._block }

func (s *SContext) Set_block(v IBlockContext) { s._block = v }

func (s *SContext) GetCode() []interface{} { return s.code }

func (s *SContext) SetCode(v []interface{}) { s.code = v }

func (s *SContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *SContext) EOF() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserEOF, 0)
}

func (s *SContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterS(s)
	}
}

func (s *SContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitS(s)
	}
}

func (p *SwiftGrammarParser) S() (localctx ISContext) {
	localctx = NewSContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SwiftGrammarParserRULE_s)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(100)

		var _x = p.Block()

		localctx.(*SContext)._block = _x
	}
	{
		p.SetState(101)
		p.Match(SwiftGrammarParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	localctx.(*SContext).code = localctx.(*SContext).Get_block().GetBlk()

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

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_instruction returns the _instruction rule contexts.
	Get_instruction() IInstructionContext

	// Set_instruction sets the _instruction rule contexts.
	Set_instruction(IInstructionContext)

	// GetIns returns the ins rule context list.
	GetIns() []IInstructionContext

	// SetIns sets the ins rule context list.
	SetIns([]IInstructionContext)

	// GetBlk returns the blk attribute.
	GetBlk() []interface{}

	// SetBlk sets the blk attribute.
	SetBlk([]interface{})

	// Getter signatures
	AllInstruction() []IInstructionContext
	Instruction(i int) IInstructionContext

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	antlr.BaseParserRuleContext
	parser       antlr.Parser
	blk          []interface{}
	_instruction IInstructionContext
	ins          []IInstructionContext
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_block
	return p
}

func InitEmptyBlockContext(p *BlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_block
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) Get_instruction() IInstructionContext { return s._instruction }

func (s *BlockContext) Set_instruction(v IInstructionContext) { s._instruction = v }

func (s *BlockContext) GetIns() []IInstructionContext { return s.ins }

func (s *BlockContext) SetIns(v []IInstructionContext) { s.ins = v }

func (s *BlockContext) GetBlk() []interface{} { return s.blk }

func (s *BlockContext) SetBlk(v []interface{}) { s.blk = v }

func (s *BlockContext) AllInstruction() []IInstructionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IInstructionContext); ok {
			len++
		}
	}

	tst := make([]IInstructionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IInstructionContext); ok {
			tst[i] = t.(IInstructionContext)
			i++
		}
	}

	return tst
}

func (s *BlockContext) Instruction(i int) IInstructionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstructionContext); ok {
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

	return t.(IInstructionContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (p *SwiftGrammarParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SwiftGrammarParserRULE_block)

	localctx.(*BlockContext).blk = []interface{}{}
	var listInt []IInstructionContext

	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(107)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4907449088) != 0 {
		{
			p.SetState(104)

			var _x = p.Instruction()

			localctx.(*BlockContext)._instruction = _x
		}
		localctx.(*BlockContext).ins = append(localctx.(*BlockContext).ins, localctx.(*BlockContext)._instruction)

		p.SetState(109)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

	listInt = localctx.(*BlockContext).GetIns()
	for _, e := range listInt {
		localctx.(*BlockContext).blk = append(localctx.(*BlockContext).blk, e.GetInst())
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

// IArgumentsContext is an interface to support dynamic dispatch.
type IArgumentsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_argument returns the _argument rule contexts.
	Get_argument() IArgumentContext

	// Get_arguments returns the _arguments rule contexts.
	Get_arguments() IArgumentsContext

	// Set_argument sets the _argument rule contexts.
	Set_argument(IArgumentContext)

	// Set_arguments sets the _arguments rule contexts.
	Set_arguments(IArgumentsContext)

	// GetArgs returns the args attribute.
	GetArgs() []interface{}

	// SetArgs sets the args attribute.
	SetArgs([]interface{})

	// Getter signatures
	Argument() IArgumentContext
	COMA() antlr.TerminalNode
	Arguments() IArgumentsContext

	// IsArgumentsContext differentiates from other interfaces.
	IsArgumentsContext()
}

type ArgumentsContext struct {
	antlr.BaseParserRuleContext
	parser     antlr.Parser
	args       []interface{}
	_argument  IArgumentContext
	_arguments IArgumentsContext
}

func NewEmptyArgumentsContext() *ArgumentsContext {
	var p = new(ArgumentsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_arguments
	return p
}

func InitEmptyArgumentsContext(p *ArgumentsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_arguments
}

func (*ArgumentsContext) IsArgumentsContext() {}

func NewArgumentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentsContext {
	var p = new(ArgumentsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_arguments

	return p
}

func (s *ArgumentsContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentsContext) Get_argument() IArgumentContext { return s._argument }

func (s *ArgumentsContext) Get_arguments() IArgumentsContext { return s._arguments }

func (s *ArgumentsContext) Set_argument(v IArgumentContext) { s._argument = v }

func (s *ArgumentsContext) Set_arguments(v IArgumentsContext) { s._arguments = v }

func (s *ArgumentsContext) GetArgs() []interface{} { return s.args }

func (s *ArgumentsContext) SetArgs(v []interface{}) { s.args = v }

func (s *ArgumentsContext) Argument() IArgumentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentContext)
}

func (s *ArgumentsContext) COMA() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCOMA, 0)
}

func (s *ArgumentsContext) Arguments() IArgumentsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentsContext)
}

func (s *ArgumentsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterArguments(s)
	}
}

func (s *ArgumentsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitArguments(s)
	}
}

func (p *SwiftGrammarParser) Arguments() (localctx IArgumentsContext) {
	localctx = NewArgumentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, SwiftGrammarParserRULE_arguments)

	localctx.(*ArgumentsContext).args = []interface{}{}

	p.SetState(121)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(112)

			var _x = p.Argument()

			localctx.(*ArgumentsContext)._argument = _x
		}
		{
			p.SetState(113)
			p.Match(SwiftGrammarParserCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(114)

			var _x = p.Arguments()

			localctx.(*ArgumentsContext)._arguments = _x
		}
		localctx.(*ArgumentsContext).args = append(localctx.(*ArgumentsContext).args, localctx.(*ArgumentsContext).Get_argument().GetE())
		for _, arg := range localctx.(*ArgumentsContext).Get_arguments().GetArgs() {
			localctx.(*ArgumentsContext).args = append(localctx.(*ArgumentsContext).args, arg)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(117)

			var _x = p.Argument()

			localctx.(*ArgumentsContext)._argument = _x
		}
		localctx.(*ArgumentsContext).args = append(localctx.(*ArgumentsContext).args, localctx.(*ArgumentsContext).Get_argument().GetE())

	case 3:
		p.EnterOuterAlt(localctx, 3)

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

// IArgumentContext is an interface to support dynamic dispatch.
type IArgumentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// GetE returns the e attribute.
	GetE() interface{}

	// SetE sets the e attribute.
	SetE(interface{})

	// Getter signatures
	Expr() IExprContext

	// IsArgumentContext differentiates from other interfaces.
	IsArgumentContext()
}

type ArgumentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	e      interface{}
	_expr  IExprContext
}

func NewEmptyArgumentContext() *ArgumentContext {
	var p = new(ArgumentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_argument
	return p
}

func InitEmptyArgumentContext(p *ArgumentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_argument
}

func (*ArgumentContext) IsArgumentContext() {}

func NewArgumentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentContext {
	var p = new(ArgumentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_argument

	return p
}

func (s *ArgumentContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentContext) Get_expr() IExprContext { return s._expr }

func (s *ArgumentContext) Set_expr(v IExprContext) { s._expr = v }

func (s *ArgumentContext) GetE() interface{} { return s.e }

func (s *ArgumentContext) SetE(v interface{}) { s.e = v }

func (s *ArgumentContext) Expr() IExprContext {
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

func (s *ArgumentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterArgument(s)
	}
}

func (s *ArgumentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitArgument(s)
	}
}

func (p *SwiftGrammarParser) Argument() (localctx IArgumentContext) {
	localctx = NewArgumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, SwiftGrammarParserRULE_argument)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(123)

		var _x = p.expr(0)

		localctx.(*ArgumentContext)._expr = _x
	}
	localctx.(*ArgumentContext).SetE(localctx.(*ArgumentContext).Get_expr().GetE())

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

// IInstructionContext is an interface to support dynamic dispatch.
type IInstructionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_printstmt returns the _printstmt rule contexts.
	Get_printstmt() IPrintstmtContext

	// Get_decmatrix returns the _decmatrix rule contexts.
	Get_decmatrix() IDecmatrixContext

	// Get_vecdec returns the _vecdec rule contexts.
	Get_vecdec() IVecdecContext

	// Get_vardec returns the _vardec rule contexts.
	Get_vardec() IVardecContext

	// Get_constdec returns the _constdec rule contexts.
	Get_constdec() IConstdecContext

	// Get_appendvec returns the _appendvec rule contexts.
	Get_appendvec() IAppendvecContext

	// Get_removelastvec returns the _removelastvec rule contexts.
	Get_removelastvec() IRemovelastvecContext

	// Get_removeatvec returns the _removeatvec rule contexts.
	Get_removeatvec() IRemoveatvecContext

	// Get_asignation returns the _asignation rule contexts.
	Get_asignation() IAsignationContext

	// Get_unarysum returns the _unarysum rule contexts.
	Get_unarysum() IUnarysumContext

	// Get_unarysub returns the _unarysub rule contexts.
	Get_unarysub() IUnarysubContext

	// Get_breakstatement returns the _breakstatement rule contexts.
	Get_breakstatement() IBreakstatementContext

	// Get_continuestatement returns the _continuestatement rule contexts.
	Get_continuestatement() IContinuestatementContext

	// Get_vectormodification returns the _vectormodification rule contexts.
	Get_vectormodification() IVectormodificationContext

	// Get_ifstmt returns the _ifstmt rule contexts.
	Get_ifstmt() IIfstmtContext

	// Get_while_statement returns the _while_statement rule contexts.
	Get_while_statement() IWhile_statementContext

	// Get_switchstatement returns the _switchstatement rule contexts.
	Get_switchstatement() ISwitchstatementContext

	// Get_forloop returns the _forloop rule contexts.
	Get_forloop() IForloopContext

	// Get_structdef returns the _structdef rule contexts.
	Get_structdef() IStructdefContext

	// Set_printstmt sets the _printstmt rule contexts.
	Set_printstmt(IPrintstmtContext)

	// Set_decmatrix sets the _decmatrix rule contexts.
	Set_decmatrix(IDecmatrixContext)

	// Set_vecdec sets the _vecdec rule contexts.
	Set_vecdec(IVecdecContext)

	// Set_vardec sets the _vardec rule contexts.
	Set_vardec(IVardecContext)

	// Set_constdec sets the _constdec rule contexts.
	Set_constdec(IConstdecContext)

	// Set_appendvec sets the _appendvec rule contexts.
	Set_appendvec(IAppendvecContext)

	// Set_removelastvec sets the _removelastvec rule contexts.
	Set_removelastvec(IRemovelastvecContext)

	// Set_removeatvec sets the _removeatvec rule contexts.
	Set_removeatvec(IRemoveatvecContext)

	// Set_asignation sets the _asignation rule contexts.
	Set_asignation(IAsignationContext)

	// Set_unarysum sets the _unarysum rule contexts.
	Set_unarysum(IUnarysumContext)

	// Set_unarysub sets the _unarysub rule contexts.
	Set_unarysub(IUnarysubContext)

	// Set_breakstatement sets the _breakstatement rule contexts.
	Set_breakstatement(IBreakstatementContext)

	// Set_continuestatement sets the _continuestatement rule contexts.
	Set_continuestatement(IContinuestatementContext)

	// Set_vectormodification sets the _vectormodification rule contexts.
	Set_vectormodification(IVectormodificationContext)

	// Set_ifstmt sets the _ifstmt rule contexts.
	Set_ifstmt(IIfstmtContext)

	// Set_while_statement sets the _while_statement rule contexts.
	Set_while_statement(IWhile_statementContext)

	// Set_switchstatement sets the _switchstatement rule contexts.
	Set_switchstatement(ISwitchstatementContext)

	// Set_forloop sets the _forloop rule contexts.
	Set_forloop(IForloopContext)

	// Set_structdef sets the _structdef rule contexts.
	Set_structdef(IStructdefContext)

	// GetInst returns the inst attribute.
	GetInst() interfaces.Instruction

	// SetInst sets the inst attribute.
	SetInst(interfaces.Instruction)

	// Getter signatures
	Printstmt() IPrintstmtContext
	PTOCOMA() antlr.TerminalNode
	Decmatrix() IDecmatrixContext
	Vecdec() IVecdecContext
	Vardec() IVardecContext
	Constdec() IConstdecContext
	Appendvec() IAppendvecContext
	Removelastvec() IRemovelastvecContext
	Removeatvec() IRemoveatvecContext
	Asignation() IAsignationContext
	Unarysum() IUnarysumContext
	Unarysub() IUnarysubContext
	Breakstatement() IBreakstatementContext
	Continuestatement() IContinuestatementContext
	Vectormodification() IVectormodificationContext
	Ifstmt() IIfstmtContext
	While_statement() IWhile_statementContext
	Switchstatement() ISwitchstatementContext
	Forloop() IForloopContext
	Structdef() IStructdefContext

	// IsInstructionContext differentiates from other interfaces.
	IsInstructionContext()
}

type InstructionContext struct {
	antlr.BaseParserRuleContext
	parser              antlr.Parser
	inst                interfaces.Instruction
	_printstmt          IPrintstmtContext
	_decmatrix          IDecmatrixContext
	_vecdec             IVecdecContext
	_vardec             IVardecContext
	_constdec           IConstdecContext
	_appendvec          IAppendvecContext
	_removelastvec      IRemovelastvecContext
	_removeatvec        IRemoveatvecContext
	_asignation         IAsignationContext
	_unarysum           IUnarysumContext
	_unarysub           IUnarysubContext
	_breakstatement     IBreakstatementContext
	_continuestatement  IContinuestatementContext
	_vectormodification IVectormodificationContext
	_ifstmt             IIfstmtContext
	_while_statement    IWhile_statementContext
	_switchstatement    ISwitchstatementContext
	_forloop            IForloopContext
	_structdef          IStructdefContext
}

func NewEmptyInstructionContext() *InstructionContext {
	var p = new(InstructionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_instruction
	return p
}

func InitEmptyInstructionContext(p *InstructionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_instruction
}

func (*InstructionContext) IsInstructionContext() {}

func NewInstructionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstructionContext {
	var p = new(InstructionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_instruction

	return p
}

func (s *InstructionContext) GetParser() antlr.Parser { return s.parser }

func (s *InstructionContext) Get_printstmt() IPrintstmtContext { return s._printstmt }

func (s *InstructionContext) Get_decmatrix() IDecmatrixContext { return s._decmatrix }

func (s *InstructionContext) Get_vecdec() IVecdecContext { return s._vecdec }

func (s *InstructionContext) Get_vardec() IVardecContext { return s._vardec }

func (s *InstructionContext) Get_constdec() IConstdecContext { return s._constdec }

func (s *InstructionContext) Get_appendvec() IAppendvecContext { return s._appendvec }

func (s *InstructionContext) Get_removelastvec() IRemovelastvecContext { return s._removelastvec }

func (s *InstructionContext) Get_removeatvec() IRemoveatvecContext { return s._removeatvec }

func (s *InstructionContext) Get_asignation() IAsignationContext { return s._asignation }

func (s *InstructionContext) Get_unarysum() IUnarysumContext { return s._unarysum }

func (s *InstructionContext) Get_unarysub() IUnarysubContext { return s._unarysub }

func (s *InstructionContext) Get_breakstatement() IBreakstatementContext { return s._breakstatement }

func (s *InstructionContext) Get_continuestatement() IContinuestatementContext {
	return s._continuestatement
}

func (s *InstructionContext) Get_vectormodification() IVectormodificationContext {
	return s._vectormodification
}

func (s *InstructionContext) Get_ifstmt() IIfstmtContext { return s._ifstmt }

func (s *InstructionContext) Get_while_statement() IWhile_statementContext { return s._while_statement }

func (s *InstructionContext) Get_switchstatement() ISwitchstatementContext { return s._switchstatement }

func (s *InstructionContext) Get_forloop() IForloopContext { return s._forloop }

func (s *InstructionContext) Get_structdef() IStructdefContext { return s._structdef }

func (s *InstructionContext) Set_printstmt(v IPrintstmtContext) { s._printstmt = v }

func (s *InstructionContext) Set_decmatrix(v IDecmatrixContext) { s._decmatrix = v }

func (s *InstructionContext) Set_vecdec(v IVecdecContext) { s._vecdec = v }

func (s *InstructionContext) Set_vardec(v IVardecContext) { s._vardec = v }

func (s *InstructionContext) Set_constdec(v IConstdecContext) { s._constdec = v }

func (s *InstructionContext) Set_appendvec(v IAppendvecContext) { s._appendvec = v }

func (s *InstructionContext) Set_removelastvec(v IRemovelastvecContext) { s._removelastvec = v }

func (s *InstructionContext) Set_removeatvec(v IRemoveatvecContext) { s._removeatvec = v }

func (s *InstructionContext) Set_asignation(v IAsignationContext) { s._asignation = v }

func (s *InstructionContext) Set_unarysum(v IUnarysumContext) { s._unarysum = v }

func (s *InstructionContext) Set_unarysub(v IUnarysubContext) { s._unarysub = v }

func (s *InstructionContext) Set_breakstatement(v IBreakstatementContext) { s._breakstatement = v }

func (s *InstructionContext) Set_continuestatement(v IContinuestatementContext) {
	s._continuestatement = v
}

func (s *InstructionContext) Set_vectormodification(v IVectormodificationContext) {
	s._vectormodification = v
}

func (s *InstructionContext) Set_ifstmt(v IIfstmtContext) { s._ifstmt = v }

func (s *InstructionContext) Set_while_statement(v IWhile_statementContext) { s._while_statement = v }

func (s *InstructionContext) Set_switchstatement(v ISwitchstatementContext) { s._switchstatement = v }

func (s *InstructionContext) Set_forloop(v IForloopContext) { s._forloop = v }

func (s *InstructionContext) Set_structdef(v IStructdefContext) { s._structdef = v }

func (s *InstructionContext) GetInst() interfaces.Instruction { return s.inst }

func (s *InstructionContext) SetInst(v interfaces.Instruction) { s.inst = v }

func (s *InstructionContext) Printstmt() IPrintstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrintstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrintstmtContext)
}

func (s *InstructionContext) PTOCOMA() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPTOCOMA, 0)
}

func (s *InstructionContext) Decmatrix() IDecmatrixContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDecmatrixContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDecmatrixContext)
}

func (s *InstructionContext) Vecdec() IVecdecContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVecdecContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVecdecContext)
}

func (s *InstructionContext) Vardec() IVardecContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVardecContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVardecContext)
}

func (s *InstructionContext) Constdec() IConstdecContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstdecContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstdecContext)
}

func (s *InstructionContext) Appendvec() IAppendvecContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAppendvecContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAppendvecContext)
}

func (s *InstructionContext) Removelastvec() IRemovelastvecContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRemovelastvecContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRemovelastvecContext)
}

func (s *InstructionContext) Removeatvec() IRemoveatvecContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRemoveatvecContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRemoveatvecContext)
}

func (s *InstructionContext) Asignation() IAsignationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAsignationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAsignationContext)
}

func (s *InstructionContext) Unarysum() IUnarysumContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnarysumContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnarysumContext)
}

func (s *InstructionContext) Unarysub() IUnarysubContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnarysubContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnarysubContext)
}

func (s *InstructionContext) Breakstatement() IBreakstatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBreakstatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBreakstatementContext)
}

func (s *InstructionContext) Continuestatement() IContinuestatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IContinuestatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IContinuestatementContext)
}

func (s *InstructionContext) Vectormodification() IVectormodificationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVectormodificationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVectormodificationContext)
}

func (s *InstructionContext) Ifstmt() IIfstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfstmtContext)
}

func (s *InstructionContext) While_statement() IWhile_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhile_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhile_statementContext)
}

func (s *InstructionContext) Switchstatement() ISwitchstatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISwitchstatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISwitchstatementContext)
}

func (s *InstructionContext) Forloop() IForloopContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForloopContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForloopContext)
}

func (s *InstructionContext) Structdef() IStructdefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructdefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructdefContext)
}

func (s *InstructionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstructionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstructionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterInstruction(s)
	}
}

func (s *InstructionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitInstruction(s)
	}
}

func (p *SwiftGrammarParser) Instruction() (localctx IInstructionContext) {
	localctx = NewInstructionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, SwiftGrammarParserRULE_instruction)
	var _la int

	p.SetState(225)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(126)

			var _x = p.Printstmt()

			localctx.(*InstructionContext)._printstmt = _x
		}
		p.SetState(128)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserPTOCOMA {
			{
				p.SetState(127)
				p.Match(SwiftGrammarParserPTOCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_printstmt().GetPrnt()

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(132)

			var _x = p.Decmatrix()

			localctx.(*InstructionContext)._decmatrix = _x
		}
		p.SetState(134)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserPTOCOMA {
			{
				p.SetState(133)
				p.Match(SwiftGrammarParserPTOCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_decmatrix().GetNewmatrix()

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(138)

			var _x = p.Vecdec()

			localctx.(*InstructionContext)._vecdec = _x
		}
		p.SetState(140)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserPTOCOMA {
			{
				p.SetState(139)
				p.Match(SwiftGrammarParserPTOCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_vecdec().GetNewvecdec()

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(144)

			var _x = p.Vardec()

			localctx.(*InstructionContext)._vardec = _x
		}
		p.SetState(146)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserPTOCOMA {
			{
				p.SetState(145)
				p.Match(SwiftGrammarParserPTOCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_vardec().GetNewdec()

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(150)

			var _x = p.Constdec()

			localctx.(*InstructionContext)._constdec = _x
		}
		p.SetState(152)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserPTOCOMA {
			{
				p.SetState(151)
				p.Match(SwiftGrammarParserPTOCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_constdec().GetNewconst()

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(156)

			var _x = p.Appendvec()

			localctx.(*InstructionContext)._appendvec = _x
		}
		p.SetState(158)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserPTOCOMA {
			{
				p.SetState(157)
				p.Match(SwiftGrammarParserPTOCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_appendvec().GetNewappendvec()

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(162)

			var _x = p.Removelastvec()

			localctx.(*InstructionContext)._removelastvec = _x
		}
		p.SetState(164)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserPTOCOMA {
			{
				p.SetState(163)
				p.Match(SwiftGrammarParserPTOCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_removelastvec().GetNewremovelastvec()

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(168)

			var _x = p.Removeatvec()

			localctx.(*InstructionContext)._removeatvec = _x
		}
		p.SetState(170)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserPTOCOMA {
			{
				p.SetState(169)
				p.Match(SwiftGrammarParserPTOCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_removeatvec().GetNewremoveat()

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(174)

			var _x = p.Asignation()

			localctx.(*InstructionContext)._asignation = _x
		}
		p.SetState(176)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserPTOCOMA {
			{
				p.SetState(175)
				p.Match(SwiftGrammarParserPTOCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_asignation().GetNewasignation()

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(180)

			var _x = p.Unarysum()

			localctx.(*InstructionContext)._unarysum = _x
		}
		p.SetState(182)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserPTOCOMA {
			{
				p.SetState(181)
				p.Match(SwiftGrammarParserPTOCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_unarysum().GetNewunarysum()

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(186)

			var _x = p.Unarysub()

			localctx.(*InstructionContext)._unarysub = _x
		}
		p.SetState(188)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserPTOCOMA {
			{
				p.SetState(187)
				p.Match(SwiftGrammarParserPTOCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_unarysub().GetNewunarysub()

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(192)

			var _x = p.Breakstatement()

			localctx.(*InstructionContext)._breakstatement = _x
		}
		p.SetState(194)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserPTOCOMA {
			{
				p.SetState(193)
				p.Match(SwiftGrammarParserPTOCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_breakstatement().GetNewbreak()

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(198)

			var _x = p.Continuestatement()

			localctx.(*InstructionContext)._continuestatement = _x
		}
		p.SetState(200)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserPTOCOMA {
			{
				p.SetState(199)
				p.Match(SwiftGrammarParserPTOCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_continuestatement().GetNewcontinue()

	case 14:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(204)

			var _x = p.Vectormodification()

			localctx.(*InstructionContext)._vectormodification = _x
		}
		p.SetState(206)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserPTOCOMA {
			{
				p.SetState(205)
				p.Match(SwiftGrammarParserPTOCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_vectormodification().GetNewvecmod()

	case 15:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(210)

			var _x = p.Ifstmt()

			localctx.(*InstructionContext)._ifstmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_ifstmt().GetNewif()

	case 16:
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(213)

			var _x = p.While_statement()

			localctx.(*InstructionContext)._while_statement = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_while_statement().GetNewwhile()

	case 17:
		p.EnterOuterAlt(localctx, 17)
		{
			p.SetState(216)

			var _x = p.Switchstatement()

			localctx.(*InstructionContext)._switchstatement = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_switchstatement().GetNewswitch()

	case 18:
		p.EnterOuterAlt(localctx, 18)
		{
			p.SetState(219)

			var _x = p.Forloop()

			localctx.(*InstructionContext)._forloop = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_forloop().GetNewfor()

	case 19:
		p.EnterOuterAlt(localctx, 19)
		{
			p.SetState(222)

			var _x = p.Structdef()

			localctx.(*InstructionContext)._structdef = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_structdef().GetNewstruct()

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

// IStructinstructionContext is an interface to support dynamic dispatch.
type IStructinstructionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_vecdec returns the _vecdec rule contexts.
	Get_vecdec() IVecdecContext

	// Get_vardec returns the _vardec rule contexts.
	Get_vardec() IVardecContext

	// Get_constdec returns the _constdec rule contexts.
	Get_constdec() IConstdecContext

	// Get_decmatrix returns the _decmatrix rule contexts.
	Get_decmatrix() IDecmatrixContext

	// Set_vecdec sets the _vecdec rule contexts.
	Set_vecdec(IVecdecContext)

	// Set_vardec sets the _vardec rule contexts.
	Set_vardec(IVardecContext)

	// Set_constdec sets the _constdec rule contexts.
	Set_constdec(IConstdecContext)

	// Set_decmatrix sets the _decmatrix rule contexts.
	Set_decmatrix(IDecmatrixContext)

	// GetInst returns the inst attribute.
	GetInst() interfaces.Instruction

	// SetInst sets the inst attribute.
	SetInst(interfaces.Instruction)

	// Getter signatures
	Vecdec() IVecdecContext
	PTOCOMA() antlr.TerminalNode
	Vardec() IVardecContext
	Constdec() IConstdecContext
	Decmatrix() IDecmatrixContext

	// IsStructinstructionContext differentiates from other interfaces.
	IsStructinstructionContext()
}

type StructinstructionContext struct {
	antlr.BaseParserRuleContext
	parser     antlr.Parser
	inst       interfaces.Instruction
	_vecdec    IVecdecContext
	_vardec    IVardecContext
	_constdec  IConstdecContext
	_decmatrix IDecmatrixContext
}

func NewEmptyStructinstructionContext() *StructinstructionContext {
	var p = new(StructinstructionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_structinstruction
	return p
}

func InitEmptyStructinstructionContext(p *StructinstructionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_structinstruction
}

func (*StructinstructionContext) IsStructinstructionContext() {}

func NewStructinstructionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructinstructionContext {
	var p = new(StructinstructionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_structinstruction

	return p
}

func (s *StructinstructionContext) GetParser() antlr.Parser { return s.parser }

func (s *StructinstructionContext) Get_vecdec() IVecdecContext { return s._vecdec }

func (s *StructinstructionContext) Get_vardec() IVardecContext { return s._vardec }

func (s *StructinstructionContext) Get_constdec() IConstdecContext { return s._constdec }

func (s *StructinstructionContext) Get_decmatrix() IDecmatrixContext { return s._decmatrix }

func (s *StructinstructionContext) Set_vecdec(v IVecdecContext) { s._vecdec = v }

func (s *StructinstructionContext) Set_vardec(v IVardecContext) { s._vardec = v }

func (s *StructinstructionContext) Set_constdec(v IConstdecContext) { s._constdec = v }

func (s *StructinstructionContext) Set_decmatrix(v IDecmatrixContext) { s._decmatrix = v }

func (s *StructinstructionContext) GetInst() interfaces.Instruction { return s.inst }

func (s *StructinstructionContext) SetInst(v interfaces.Instruction) { s.inst = v }

func (s *StructinstructionContext) Vecdec() IVecdecContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVecdecContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVecdecContext)
}

func (s *StructinstructionContext) PTOCOMA() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPTOCOMA, 0)
}

func (s *StructinstructionContext) Vardec() IVardecContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVardecContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVardecContext)
}

func (s *StructinstructionContext) Constdec() IConstdecContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstdecContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstdecContext)
}

func (s *StructinstructionContext) Decmatrix() IDecmatrixContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDecmatrixContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDecmatrixContext)
}

func (s *StructinstructionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructinstructionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructinstructionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterStructinstruction(s)
	}
}

func (s *StructinstructionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitStructinstruction(s)
	}
}

func (p *SwiftGrammarParser) Structinstruction() (localctx IStructinstructionContext) {
	localctx = NewStructinstructionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, SwiftGrammarParserRULE_structinstruction)
	var _la int

	p.SetState(251)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(227)

			var _x = p.Vecdec()

			localctx.(*StructinstructionContext)._vecdec = _x
		}
		p.SetState(229)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserPTOCOMA {
			{
				p.SetState(228)
				p.Match(SwiftGrammarParserPTOCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		localctx.(*StructinstructionContext).inst = localctx.(*StructinstructionContext).Get_vecdec().GetNewvecdec()

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(233)

			var _x = p.Vardec()

			localctx.(*StructinstructionContext)._vardec = _x
		}
		p.SetState(235)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserPTOCOMA {
			{
				p.SetState(234)
				p.Match(SwiftGrammarParserPTOCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		localctx.(*StructinstructionContext).inst = localctx.(*StructinstructionContext).Get_vardec().GetNewdec()

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(239)

			var _x = p.Constdec()

			localctx.(*StructinstructionContext)._constdec = _x
		}
		p.SetState(241)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserPTOCOMA {
			{
				p.SetState(240)
				p.Match(SwiftGrammarParserPTOCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		localctx.(*StructinstructionContext).inst = localctx.(*StructinstructionContext).Get_constdec().GetNewconst()

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(245)

			var _x = p.Decmatrix()

			localctx.(*StructinstructionContext)._decmatrix = _x
		}
		p.SetState(247)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserPTOCOMA {
			{
				p.SetState(246)
				p.Match(SwiftGrammarParserPTOCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		localctx.(*StructinstructionContext).inst = localctx.(*StructinstructionContext).Get_decmatrix().GetNewmatrix()

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

// IStructblockContext is an interface to support dynamic dispatch.
type IStructblockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_structinstruction returns the _structinstruction rule contexts.
	Get_structinstruction() IStructinstructionContext

	// Set_structinstruction sets the _structinstruction rule contexts.
	Set_structinstruction(IStructinstructionContext)

	// GetIns returns the ins rule context list.
	GetIns() []IStructinstructionContext

	// SetIns sets the ins rule context list.
	SetIns([]IStructinstructionContext)

	// GetSblk returns the sblk attribute.
	GetSblk() []interface{}

	// SetSblk sets the sblk attribute.
	SetSblk([]interface{})

	// Getter signatures
	AllStructinstruction() []IStructinstructionContext
	Structinstruction(i int) IStructinstructionContext

	// IsStructblockContext differentiates from other interfaces.
	IsStructblockContext()
}

type StructblockContext struct {
	antlr.BaseParserRuleContext
	parser             antlr.Parser
	sblk               []interface{}
	_structinstruction IStructinstructionContext
	ins                []IStructinstructionContext
}

func NewEmptyStructblockContext() *StructblockContext {
	var p = new(StructblockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_structblock
	return p
}

func InitEmptyStructblockContext(p *StructblockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_structblock
}

func (*StructblockContext) IsStructblockContext() {}

func NewStructblockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructblockContext {
	var p = new(StructblockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_structblock

	return p
}

func (s *StructblockContext) GetParser() antlr.Parser { return s.parser }

func (s *StructblockContext) Get_structinstruction() IStructinstructionContext {
	return s._structinstruction
}

func (s *StructblockContext) Set_structinstruction(v IStructinstructionContext) {
	s._structinstruction = v
}

func (s *StructblockContext) GetIns() []IStructinstructionContext { return s.ins }

func (s *StructblockContext) SetIns(v []IStructinstructionContext) { s.ins = v }

func (s *StructblockContext) GetSblk() []interface{} { return s.sblk }

func (s *StructblockContext) SetSblk(v []interface{}) { s.sblk = v }

func (s *StructblockContext) AllStructinstruction() []IStructinstructionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStructinstructionContext); ok {
			len++
		}
	}

	tst := make([]IStructinstructionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStructinstructionContext); ok {
			tst[i] = t.(IStructinstructionContext)
			i++
		}
	}

	return tst
}

func (s *StructblockContext) Structinstruction(i int) IStructinstructionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructinstructionContext); ok {
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

	return t.(IStructinstructionContext)
}

func (s *StructblockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructblockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructblockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterStructblock(s)
	}
}

func (s *StructblockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitStructblock(s)
	}
}

func (p *SwiftGrammarParser) Structblock() (localctx IStructblockContext) {
	localctx = NewStructblockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, SwiftGrammarParserRULE_structblock)

	localctx.(*StructblockContext).sblk = []interface{}{}
	var listsinst []IStructinstructionContext

	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(256)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SwiftGrammarParserRVAR || _la == SwiftGrammarParserRLET {
		{
			p.SetState(253)

			var _x = p.Structinstruction()

			localctx.(*StructblockContext)._structinstruction = _x
		}
		localctx.(*StructblockContext).ins = append(localctx.(*StructblockContext).ins, localctx.(*StructblockContext)._structinstruction)

		p.SetState(258)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

	listsinst = localctx.(*StructblockContext).GetIns()
	for _, e := range listsinst {
		localctx.(*StructblockContext).sblk = append(localctx.(*StructblockContext).sblk, e.GetInst())
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

// IStructdefContext is an interface to support dynamic dispatch.
type IStructdefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_structblock returns the _structblock rule contexts.
	Get_structblock() IStructblockContext

	// Set_structblock sets the _structblock rule contexts.
	Set_structblock(IStructblockContext)

	// GetNewstruct returns the newstruct attribute.
	GetNewstruct() interfaces.Instruction

	// SetNewstruct sets the newstruct attribute.
	SetNewstruct(interfaces.Instruction)

	// Getter signatures
	RSTRUCT() antlr.TerminalNode
	ID() antlr.TerminalNode
	LLAVEIZQ() antlr.TerminalNode
	Structblock() IStructblockContext
	LLAVEDER() antlr.TerminalNode

	// IsStructdefContext differentiates from other interfaces.
	IsStructdefContext()
}

type StructdefContext struct {
	antlr.BaseParserRuleContext
	parser       antlr.Parser
	newstruct    interfaces.Instruction
	_ID          antlr.Token
	_structblock IStructblockContext
}

func NewEmptyStructdefContext() *StructdefContext {
	var p = new(StructdefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_structdef
	return p
}

func InitEmptyStructdefContext(p *StructdefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_structdef
}

func (*StructdefContext) IsStructdefContext() {}

func NewStructdefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructdefContext {
	var p = new(StructdefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_structdef

	return p
}

func (s *StructdefContext) GetParser() antlr.Parser { return s.parser }

func (s *StructdefContext) Get_ID() antlr.Token { return s._ID }

func (s *StructdefContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *StructdefContext) Get_structblock() IStructblockContext { return s._structblock }

func (s *StructdefContext) Set_structblock(v IStructblockContext) { s._structblock = v }

func (s *StructdefContext) GetNewstruct() interfaces.Instruction { return s.newstruct }

func (s *StructdefContext) SetNewstruct(v interfaces.Instruction) { s.newstruct = v }

func (s *StructdefContext) RSTRUCT() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserRSTRUCT, 0)
}

func (s *StructdefContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *StructdefContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEIZQ, 0)
}

func (s *StructdefContext) Structblock() IStructblockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructblockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructblockContext)
}

func (s *StructdefContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEDER, 0)
}

func (s *StructdefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructdefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructdefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterStructdef(s)
	}
}

func (s *StructdefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitStructdef(s)
	}
}

func (p *SwiftGrammarParser) Structdef() (localctx IStructdefContext) {
	localctx = NewStructdefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, SwiftGrammarParserRULE_structdef)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(261)
		p.Match(SwiftGrammarParserRSTRUCT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(262)

		var _m = p.Match(SwiftGrammarParserID)

		localctx.(*StructdefContext)._ID = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(263)
		p.Match(SwiftGrammarParserLLAVEIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(264)

		var _x = p.Structblock()

		localctx.(*StructdefContext)._structblock = _x
	}
	{
		p.SetState(265)
		p.Match(SwiftGrammarParserLLAVEDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	localctx.(*StructdefContext).newstruct = instructions.NewStructDef((func() int {
		if localctx.(*StructdefContext).Get_ID() == nil {
			return 0
		} else {
			return localctx.(*StructdefContext).Get_ID().GetLine()
		}
	}()), (func() int {
		if localctx.(*StructdefContext).Get_ID() == nil {
			return 0
		} else {
			return localctx.(*StructdefContext).Get_ID().GetColumn()
		}
	}()), (func() string {
		if localctx.(*StructdefContext).Get_ID() == nil {
			return ""
		} else {
			return localctx.(*StructdefContext).Get_ID().GetText()
		}
	}()), localctx.(*StructdefContext).Get_structblock().GetSblk())

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

// IVectormodificationContext is an interface to support dynamic dispatch.
type IVectormodificationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_indexesList returns the _indexesList rule contexts.
	Get_indexesList() IIndexesListContext

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Set_indexesList sets the _indexesList rule contexts.
	Set_indexesList(IIndexesListContext)

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// GetNewvecmod returns the newvecmod attribute.
	GetNewvecmod() interfaces.Instruction

	// SetNewvecmod sets the newvecmod attribute.
	SetNewvecmod(interfaces.Instruction)

	// Getter signatures
	ID() antlr.TerminalNode
	IndexesList() IIndexesListContext
	IG() antlr.TerminalNode
	Expr() IExprContext

	// IsVectormodificationContext differentiates from other interfaces.
	IsVectormodificationContext()
}

type VectormodificationContext struct {
	antlr.BaseParserRuleContext
	parser       antlr.Parser
	newvecmod    interfaces.Instruction
	_ID          antlr.Token
	_indexesList IIndexesListContext
	_expr        IExprContext
}

func NewEmptyVectormodificationContext() *VectormodificationContext {
	var p = new(VectormodificationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_vectormodification
	return p
}

func InitEmptyVectormodificationContext(p *VectormodificationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_vectormodification
}

func (*VectormodificationContext) IsVectormodificationContext() {}

func NewVectormodificationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VectormodificationContext {
	var p = new(VectormodificationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_vectormodification

	return p
}

func (s *VectormodificationContext) GetParser() antlr.Parser { return s.parser }

func (s *VectormodificationContext) Get_ID() antlr.Token { return s._ID }

func (s *VectormodificationContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *VectormodificationContext) Get_indexesList() IIndexesListContext { return s._indexesList }

func (s *VectormodificationContext) Get_expr() IExprContext { return s._expr }

func (s *VectormodificationContext) Set_indexesList(v IIndexesListContext) { s._indexesList = v }

func (s *VectormodificationContext) Set_expr(v IExprContext) { s._expr = v }

func (s *VectormodificationContext) GetNewvecmod() interfaces.Instruction { return s.newvecmod }

func (s *VectormodificationContext) SetNewvecmod(v interfaces.Instruction) { s.newvecmod = v }

func (s *VectormodificationContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *VectormodificationContext) IndexesList() IIndexesListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIndexesListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIndexesListContext)
}

func (s *VectormodificationContext) IG() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserIG, 0)
}

func (s *VectormodificationContext) Expr() IExprContext {
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

func (s *VectormodificationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VectormodificationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VectormodificationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterVectormodification(s)
	}
}

func (s *VectormodificationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitVectormodification(s)
	}
}

func (p *SwiftGrammarParser) Vectormodification() (localctx IVectormodificationContext) {
	localctx = NewVectormodificationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, SwiftGrammarParserRULE_vectormodification)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(268)

		var _m = p.Match(SwiftGrammarParserID)

		localctx.(*VectormodificationContext)._ID = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(269)

		var _x = p.IndexesList()

		localctx.(*VectormodificationContext)._indexesList = _x
	}
	{
		p.SetState(270)
		p.Match(SwiftGrammarParserIG)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(271)

		var _x = p.expr(0)

		localctx.(*VectormodificationContext)._expr = _x
	}
	localctx.(*VectormodificationContext).newvecmod = instructions.NewVectorModification((func() int {
		if localctx.(*VectormodificationContext).Get_ID() == nil {
			return 0
		} else {
			return localctx.(*VectormodificationContext).Get_ID().GetLine()
		}
	}()), (func() int {
		if localctx.(*VectormodificationContext).Get_ID() == nil {
			return 0
		} else {
			return localctx.(*VectormodificationContext).Get_ID().GetColumn()
		}
	}()), (func() string {
		if localctx.(*VectormodificationContext).Get_ID() == nil {
			return ""
		} else {
			return localctx.(*VectormodificationContext).Get_ID().GetText()
		}
	}()), localctx.(*VectormodificationContext).Get_indexesList().GetIndexes(), localctx.(*VectormodificationContext).Get_expr().GetE())

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

// IForloopContext is an interface to support dynamic dispatch.
type IForloopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_RFOR returns the _RFOR token.
	Get_RFOR() antlr.Token

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_RFOR sets the _RFOR token.
	Set_RFOR(antlr.Token)

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// GetEx returns the ex rule contexts.
	GetEx() IExprContext

	// GetB returns the b rule contexts.
	GetB() IBlockContext

	// Get_range returns the _range rule contexts.
	Get_range() IRangeContext

	// SetEx sets the ex rule contexts.
	SetEx(IExprContext)

	// SetB sets the b rule contexts.
	SetB(IBlockContext)

	// Set_range sets the _range rule contexts.
	Set_range(IRangeContext)

	// GetNewfor returns the newfor attribute.
	GetNewfor() interfaces.Instruction

	// SetNewfor sets the newfor attribute.
	SetNewfor(interfaces.Instruction)

	// Getter signatures
	RFOR() antlr.TerminalNode
	ID() antlr.TerminalNode
	RIN() antlr.TerminalNode
	LLAVEIZQ() antlr.TerminalNode
	LLAVEDER() antlr.TerminalNode
	Expr() IExprContext
	Block() IBlockContext
	Range_() IRangeContext

	// IsForloopContext differentiates from other interfaces.
	IsForloopContext()
}

type ForloopContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	newfor interfaces.Instruction
	_RFOR  antlr.Token
	_ID    antlr.Token
	ex     IExprContext
	b      IBlockContext
	_range IRangeContext
}

func NewEmptyForloopContext() *ForloopContext {
	var p = new(ForloopContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_forloop
	return p
}

func InitEmptyForloopContext(p *ForloopContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_forloop
}

func (*ForloopContext) IsForloopContext() {}

func NewForloopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForloopContext {
	var p = new(ForloopContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_forloop

	return p
}

func (s *ForloopContext) GetParser() antlr.Parser { return s.parser }

func (s *ForloopContext) Get_RFOR() antlr.Token { return s._RFOR }

func (s *ForloopContext) Get_ID() antlr.Token { return s._ID }

func (s *ForloopContext) Set_RFOR(v antlr.Token) { s._RFOR = v }

func (s *ForloopContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *ForloopContext) GetEx() IExprContext { return s.ex }

func (s *ForloopContext) GetB() IBlockContext { return s.b }

func (s *ForloopContext) Get_range() IRangeContext { return s._range }

func (s *ForloopContext) SetEx(v IExprContext) { s.ex = v }

func (s *ForloopContext) SetB(v IBlockContext) { s.b = v }

func (s *ForloopContext) Set_range(v IRangeContext) { s._range = v }

func (s *ForloopContext) GetNewfor() interfaces.Instruction { return s.newfor }

func (s *ForloopContext) SetNewfor(v interfaces.Instruction) { s.newfor = v }

func (s *ForloopContext) RFOR() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserRFOR, 0)
}

func (s *ForloopContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *ForloopContext) RIN() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserRIN, 0)
}

func (s *ForloopContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEIZQ, 0)
}

func (s *ForloopContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEDER, 0)
}

func (s *ForloopContext) Expr() IExprContext {
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

func (s *ForloopContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ForloopContext) Range_() IRangeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRangeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRangeContext)
}

func (s *ForloopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForloopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForloopContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterForloop(s)
	}
}

func (s *ForloopContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitForloop(s)
	}
}

func (p *SwiftGrammarParser) Forloop() (localctx IForloopContext) {
	localctx = NewForloopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, SwiftGrammarParserRULE_forloop)
	p.SetState(292)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 23, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(274)

			var _m = p.Match(SwiftGrammarParserRFOR)

			localctx.(*ForloopContext)._RFOR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(275)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*ForloopContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(276)
			p.Match(SwiftGrammarParserRIN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(277)

			var _x = p.expr(0)

			localctx.(*ForloopContext).ex = _x
		}
		{
			p.SetState(278)
			p.Match(SwiftGrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(279)

			var _x = p.Block()

			localctx.(*ForloopContext).b = _x
		}
		{
			p.SetState(280)
			p.Match(SwiftGrammarParserLLAVEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ForloopContext).newfor = instructions.NewFor((func() int {
			if localctx.(*ForloopContext).Get_RFOR() == nil {
				return 0
			} else {
				return localctx.(*ForloopContext).Get_RFOR().GetLine()
			}
		}()), (func() int {
			if localctx.(*ForloopContext).Get_RFOR() == nil {
				return 0
			} else {
				return localctx.(*ForloopContext).Get_RFOR().GetColumn()
			}
		}()), (func() string {
			if localctx.(*ForloopContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*ForloopContext).Get_ID().GetText()
			}
		}()), localctx.(*ForloopContext).GetEx().GetE(), localctx.(*ForloopContext).GetB().GetBlk())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(283)

			var _m = p.Match(SwiftGrammarParserRFOR)

			localctx.(*ForloopContext)._RFOR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(284)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*ForloopContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(285)
			p.Match(SwiftGrammarParserRIN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(286)

			var _x = p.Range_()

			localctx.(*ForloopContext)._range = _x
		}
		{
			p.SetState(287)
			p.Match(SwiftGrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(288)

			var _x = p.Block()

			localctx.(*ForloopContext).b = _x
		}
		{
			p.SetState(289)
			p.Match(SwiftGrammarParserLLAVEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ForloopContext).newfor = instructions.NewFor((func() int {
			if localctx.(*ForloopContext).Get_RFOR() == nil {
				return 0
			} else {
				return localctx.(*ForloopContext).Get_RFOR().GetLine()
			}
		}()), (func() int {
			if localctx.(*ForloopContext).Get_RFOR() == nil {
				return 0
			} else {
				return localctx.(*ForloopContext).Get_RFOR().GetColumn()
			}
		}()), (func() string {
			if localctx.(*ForloopContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*ForloopContext).Get_ID().GetText()
			}
		}()), localctx.(*ForloopContext).Get_range().GetNewrange(), localctx.(*ForloopContext).GetB().GetBlk())

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

// IRangeContext is an interface to support dynamic dispatch.
type IRangeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetExp1 returns the exp1 rule contexts.
	GetExp1() IExprContext

	// GetExp2 returns the exp2 rule contexts.
	GetExp2() IExprContext

	// SetExp1 sets the exp1 rule contexts.
	SetExp1(IExprContext)

	// SetExp2 sets the exp2 rule contexts.
	SetExp2(IExprContext)

	// GetNewrange returns the newrange attribute.
	GetNewrange() interfaces.Expression

	// SetNewrange sets the newrange attribute.
	SetNewrange(interfaces.Expression)

	// Getter signatures
	AllPTO() []antlr.TerminalNode
	PTO(i int) antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext

	// IsRangeContext differentiates from other interfaces.
	IsRangeContext()
}

type RangeContext struct {
	antlr.BaseParserRuleContext
	parser   antlr.Parser
	newrange interfaces.Expression
	exp1     IExprContext
	exp2     IExprContext
}

func NewEmptyRangeContext() *RangeContext {
	var p = new(RangeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_range
	return p
}

func InitEmptyRangeContext(p *RangeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_range
}

func (*RangeContext) IsRangeContext() {}

func NewRangeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RangeContext {
	var p = new(RangeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_range

	return p
}

func (s *RangeContext) GetParser() antlr.Parser { return s.parser }

func (s *RangeContext) GetExp1() IExprContext { return s.exp1 }

func (s *RangeContext) GetExp2() IExprContext { return s.exp2 }

func (s *RangeContext) SetExp1(v IExprContext) { s.exp1 = v }

func (s *RangeContext) SetExp2(v IExprContext) { s.exp2 = v }

func (s *RangeContext) GetNewrange() interfaces.Expression { return s.newrange }

func (s *RangeContext) SetNewrange(v interfaces.Expression) { s.newrange = v }

func (s *RangeContext) AllPTO() []antlr.TerminalNode {
	return s.GetTokens(SwiftGrammarParserPTO)
}

func (s *RangeContext) PTO(i int) antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPTO, i)
}

func (s *RangeContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *RangeContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *RangeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RangeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RangeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterRange(s)
	}
}

func (s *RangeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitRange(s)
	}
}

func (p *SwiftGrammarParser) Range_() (localctx IRangeContext) {
	localctx = NewRangeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, SwiftGrammarParserRULE_range)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(294)

		var _x = p.expr(0)

		localctx.(*RangeContext).exp1 = _x
	}
	{
		p.SetState(295)
		p.Match(SwiftGrammarParserPTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(296)
		p.Match(SwiftGrammarParserPTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(297)
		p.Match(SwiftGrammarParserPTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(298)

		var _x = p.expr(0)

		localctx.(*RangeContext).exp2 = _x
	}
	localctx.(*RangeContext).newrange = expressions.NewRange((func() antlr.Token {
		if localctx.(*RangeContext).GetExp1() == nil {
			return nil
		} else {
			return localctx.(*RangeContext).GetExp1().GetStart()
		}
	}()).GetLine(), (func() antlr.Token {
		if localctx.(*RangeContext).GetExp1() == nil {
			return nil
		} else {
			return localctx.(*RangeContext).GetExp1().GetStart()
		}
	}()).GetColumn(), localctx.(*RangeContext).GetExp1().GetE(), localctx.(*RangeContext).GetExp2().GetE())

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

// IRemoveatvecContext is an interface to support dynamic dispatch.
type IRemoveatvecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// GetNewremoveat returns the newremoveat attribute.
	GetNewremoveat() interfaces.Instruction

	// SetNewremoveat sets the newremoveat attribute.
	SetNewremoveat(interfaces.Instruction)

	// Getter signatures
	ID() antlr.TerminalNode
	PTO() antlr.TerminalNode
	RREMOVEAT() antlr.TerminalNode
	PARIZQ() antlr.TerminalNode
	RRAT() antlr.TerminalNode
	DOSPTOS() antlr.TerminalNode
	Expr() IExprContext
	PARDER() antlr.TerminalNode

	// IsRemoveatvecContext differentiates from other interfaces.
	IsRemoveatvecContext()
}

type RemoveatvecContext struct {
	antlr.BaseParserRuleContext
	parser      antlr.Parser
	newremoveat interfaces.Instruction
	_ID         antlr.Token
	_expr       IExprContext
}

func NewEmptyRemoveatvecContext() *RemoveatvecContext {
	var p = new(RemoveatvecContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_removeatvec
	return p
}

func InitEmptyRemoveatvecContext(p *RemoveatvecContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_removeatvec
}

func (*RemoveatvecContext) IsRemoveatvecContext() {}

func NewRemoveatvecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RemoveatvecContext {
	var p = new(RemoveatvecContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_removeatvec

	return p
}

func (s *RemoveatvecContext) GetParser() antlr.Parser { return s.parser }

func (s *RemoveatvecContext) Get_ID() antlr.Token { return s._ID }

func (s *RemoveatvecContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *RemoveatvecContext) Get_expr() IExprContext { return s._expr }

func (s *RemoveatvecContext) Set_expr(v IExprContext) { s._expr = v }

func (s *RemoveatvecContext) GetNewremoveat() interfaces.Instruction { return s.newremoveat }

func (s *RemoveatvecContext) SetNewremoveat(v interfaces.Instruction) { s.newremoveat = v }

func (s *RemoveatvecContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *RemoveatvecContext) PTO() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPTO, 0)
}

func (s *RemoveatvecContext) RREMOVEAT() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserRREMOVEAT, 0)
}

func (s *RemoveatvecContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARIZQ, 0)
}

func (s *RemoveatvecContext) RRAT() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserRRAT, 0)
}

func (s *RemoveatvecContext) DOSPTOS() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserDOSPTOS, 0)
}

func (s *RemoveatvecContext) Expr() IExprContext {
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

func (s *RemoveatvecContext) PARDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARDER, 0)
}

func (s *RemoveatvecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RemoveatvecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RemoveatvecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterRemoveatvec(s)
	}
}

func (s *RemoveatvecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitRemoveatvec(s)
	}
}

func (p *SwiftGrammarParser) Removeatvec() (localctx IRemoveatvecContext) {
	localctx = NewRemoveatvecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, SwiftGrammarParserRULE_removeatvec)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(301)

		var _m = p.Match(SwiftGrammarParserID)

		localctx.(*RemoveatvecContext)._ID = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(302)
		p.Match(SwiftGrammarParserPTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(303)
		p.Match(SwiftGrammarParserRREMOVEAT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(304)
		p.Match(SwiftGrammarParserPARIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(305)
		p.Match(SwiftGrammarParserRRAT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(306)
		p.Match(SwiftGrammarParserDOSPTOS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(307)

		var _x = p.expr(0)

		localctx.(*RemoveatvecContext)._expr = _x
	}
	{
		p.SetState(308)
		p.Match(SwiftGrammarParserPARDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	localctx.(*RemoveatvecContext).newremoveat = instructions.NewRemoveAtVector((func() int {
		if localctx.(*RemoveatvecContext).Get_ID() == nil {
			return 0
		} else {
			return localctx.(*RemoveatvecContext).Get_ID().GetLine()
		}
	}()), (func() int {
		if localctx.(*RemoveatvecContext).Get_ID() == nil {
			return 0
		} else {
			return localctx.(*RemoveatvecContext).Get_ID().GetColumn()
		}
	}()), (func() string {
		if localctx.(*RemoveatvecContext).Get_ID() == nil {
			return ""
		} else {
			return localctx.(*RemoveatvecContext).Get_ID().GetText()
		}
	}()), localctx.(*RemoveatvecContext).Get_expr().GetE())

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

// IAppendvecContext is an interface to support dynamic dispatch.
type IAppendvecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// GetNewappendvec returns the newappendvec attribute.
	GetNewappendvec() interfaces.Instruction

	// SetNewappendvec sets the newappendvec attribute.
	SetNewappendvec(interfaces.Instruction)

	// Getter signatures
	ID() antlr.TerminalNode
	PTO() antlr.TerminalNode
	RAPPEND() antlr.TerminalNode
	PARIZQ() antlr.TerminalNode
	Expr() IExprContext
	PARDER() antlr.TerminalNode

	// IsAppendvecContext differentiates from other interfaces.
	IsAppendvecContext()
}

type AppendvecContext struct {
	antlr.BaseParserRuleContext
	parser       antlr.Parser
	newappendvec interfaces.Instruction
	_ID          antlr.Token
	_expr        IExprContext
}

func NewEmptyAppendvecContext() *AppendvecContext {
	var p = new(AppendvecContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_appendvec
	return p
}

func InitEmptyAppendvecContext(p *AppendvecContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_appendvec
}

func (*AppendvecContext) IsAppendvecContext() {}

func NewAppendvecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AppendvecContext {
	var p = new(AppendvecContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_appendvec

	return p
}

func (s *AppendvecContext) GetParser() antlr.Parser { return s.parser }

func (s *AppendvecContext) Get_ID() antlr.Token { return s._ID }

func (s *AppendvecContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *AppendvecContext) Get_expr() IExprContext { return s._expr }

func (s *AppendvecContext) Set_expr(v IExprContext) { s._expr = v }

func (s *AppendvecContext) GetNewappendvec() interfaces.Instruction { return s.newappendvec }

func (s *AppendvecContext) SetNewappendvec(v interfaces.Instruction) { s.newappendvec = v }

func (s *AppendvecContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *AppendvecContext) PTO() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPTO, 0)
}

func (s *AppendvecContext) RAPPEND() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserRAPPEND, 0)
}

func (s *AppendvecContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARIZQ, 0)
}

func (s *AppendvecContext) Expr() IExprContext {
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

func (s *AppendvecContext) PARDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARDER, 0)
}

func (s *AppendvecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AppendvecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AppendvecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterAppendvec(s)
	}
}

func (s *AppendvecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitAppendvec(s)
	}
}

func (p *SwiftGrammarParser) Appendvec() (localctx IAppendvecContext) {
	localctx = NewAppendvecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, SwiftGrammarParserRULE_appendvec)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(311)

		var _m = p.Match(SwiftGrammarParserID)

		localctx.(*AppendvecContext)._ID = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(312)
		p.Match(SwiftGrammarParserPTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(313)
		p.Match(SwiftGrammarParserRAPPEND)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(314)
		p.Match(SwiftGrammarParserPARIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(315)

		var _x = p.expr(0)

		localctx.(*AppendvecContext)._expr = _x
	}
	{
		p.SetState(316)
		p.Match(SwiftGrammarParserPARDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	localctx.(*AppendvecContext).newappendvec = instructions.NewAppendVector((func() int {
		if localctx.(*AppendvecContext).Get_ID() == nil {
			return 0
		} else {
			return localctx.(*AppendvecContext).Get_ID().GetLine()
		}
	}()), (func() int {
		if localctx.(*AppendvecContext).Get_ID() == nil {
			return 0
		} else {
			return localctx.(*AppendvecContext).Get_ID().GetColumn()
		}
	}()), (func() string {
		if localctx.(*AppendvecContext).Get_ID() == nil {
			return ""
		} else {
			return localctx.(*AppendvecContext).Get_ID().GetText()
		}
	}()), localctx.(*AppendvecContext).Get_expr().GetE())

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

// IRemovelastvecContext is an interface to support dynamic dispatch.
type IRemovelastvecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// GetNewremovelastvec returns the newremovelastvec attribute.
	GetNewremovelastvec() interfaces.Instruction

	// SetNewremovelastvec sets the newremovelastvec attribute.
	SetNewremovelastvec(interfaces.Instruction)

	// Getter signatures
	ID() antlr.TerminalNode
	PTO() antlr.TerminalNode
	RREMOVELAST() antlr.TerminalNode
	PARIZQ() antlr.TerminalNode
	PARDER() antlr.TerminalNode

	// IsRemovelastvecContext differentiates from other interfaces.
	IsRemovelastvecContext()
}

type RemovelastvecContext struct {
	antlr.BaseParserRuleContext
	parser           antlr.Parser
	newremovelastvec interfaces.Instruction
	_ID              antlr.Token
}

func NewEmptyRemovelastvecContext() *RemovelastvecContext {
	var p = new(RemovelastvecContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_removelastvec
	return p
}

func InitEmptyRemovelastvecContext(p *RemovelastvecContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_removelastvec
}

func (*RemovelastvecContext) IsRemovelastvecContext() {}

func NewRemovelastvecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RemovelastvecContext {
	var p = new(RemovelastvecContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_removelastvec

	return p
}

func (s *RemovelastvecContext) GetParser() antlr.Parser { return s.parser }

func (s *RemovelastvecContext) Get_ID() antlr.Token { return s._ID }

func (s *RemovelastvecContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *RemovelastvecContext) GetNewremovelastvec() interfaces.Instruction {
	return s.newremovelastvec
}

func (s *RemovelastvecContext) SetNewremovelastvec(v interfaces.Instruction) { s.newremovelastvec = v }

func (s *RemovelastvecContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *RemovelastvecContext) PTO() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPTO, 0)
}

func (s *RemovelastvecContext) RREMOVELAST() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserRREMOVELAST, 0)
}

func (s *RemovelastvecContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARIZQ, 0)
}

func (s *RemovelastvecContext) PARDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARDER, 0)
}

func (s *RemovelastvecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RemovelastvecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RemovelastvecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterRemovelastvec(s)
	}
}

func (s *RemovelastvecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitRemovelastvec(s)
	}
}

func (p *SwiftGrammarParser) Removelastvec() (localctx IRemovelastvecContext) {
	localctx = NewRemovelastvecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, SwiftGrammarParserRULE_removelastvec)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(319)

		var _m = p.Match(SwiftGrammarParserID)

		localctx.(*RemovelastvecContext)._ID = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(320)
		p.Match(SwiftGrammarParserPTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(321)
		p.Match(SwiftGrammarParserRREMOVELAST)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(322)
		p.Match(SwiftGrammarParserPARIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(323)
		p.Match(SwiftGrammarParserPARDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	localctx.(*RemovelastvecContext).newremovelastvec = instructions.NewRemoveLastVector((func() int {
		if localctx.(*RemovelastvecContext).Get_ID() == nil {
			return 0
		} else {
			return localctx.(*RemovelastvecContext).Get_ID().GetLine()
		}
	}()), (func() int {
		if localctx.(*RemovelastvecContext).Get_ID() == nil {
			return 0
		} else {
			return localctx.(*RemovelastvecContext).Get_ID().GetColumn()
		}
	}()), (func() string {
		if localctx.(*RemovelastvecContext).Get_ID() == nil {
			return ""
		} else {
			return localctx.(*RemovelastvecContext).Get_ID().GetText()
		}
	}()))

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

// IVecdecContext is an interface to support dynamic dispatch.
type IVecdecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_RVAR returns the _RVAR token.
	Get_RVAR() antlr.Token

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// GetTyppe returns the typpe token.
	GetTyppe() antlr.Token

	// GetTyppe2 returns the typpe2 token.
	GetTyppe2() antlr.Token

	// GetFirstid returns the firstid token.
	GetFirstid() antlr.Token

	// Set_RVAR sets the _RVAR token.
	Set_RVAR(antlr.Token)

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// SetTyppe sets the typpe token.
	SetTyppe(antlr.Token)

	// SetTyppe2 sets the typpe2 token.
	SetTyppe2(antlr.Token)

	// SetFirstid sets the firstid token.
	SetFirstid(antlr.Token)

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// GetNewvecdec returns the newvecdec attribute.
	GetNewvecdec() interfaces.Instruction

	// SetNewvecdec sets the newvecdec attribute.
	SetNewvecdec(interfaces.Instruction)

	// Getter signatures
	RVAR() antlr.TerminalNode
	ID() antlr.TerminalNode
	DOSPTOS() antlr.TerminalNode
	AllOBRA() []antlr.TerminalNode
	OBRA(i int) antlr.TerminalNode
	AllCBRA() []antlr.TerminalNode
	CBRA(i int) antlr.TerminalNode
	IG() antlr.TerminalNode
	AllRINT() []antlr.TerminalNode
	RINT(i int) antlr.TerminalNode
	AllRFLOAT() []antlr.TerminalNode
	RFLOAT(i int) antlr.TerminalNode
	AllRBOOL() []antlr.TerminalNode
	RBOOL(i int) antlr.TerminalNode
	AllRSTRING() []antlr.TerminalNode
	RSTRING(i int) antlr.TerminalNode
	AllRCHARACTER() []antlr.TerminalNode
	RCHARACTER(i int) antlr.TerminalNode
	Expr() IExprContext

	// IsVecdecContext differentiates from other interfaces.
	IsVecdecContext()
}

type VecdecContext struct {
	antlr.BaseParserRuleContext
	parser    antlr.Parser
	newvecdec interfaces.Instruction
	_RVAR     antlr.Token
	_ID       antlr.Token
	typpe     antlr.Token
	typpe2    antlr.Token
	firstid   antlr.Token
	_expr     IExprContext
}

func NewEmptyVecdecContext() *VecdecContext {
	var p = new(VecdecContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_vecdec
	return p
}

func InitEmptyVecdecContext(p *VecdecContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_vecdec
}

func (*VecdecContext) IsVecdecContext() {}

func NewVecdecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VecdecContext {
	var p = new(VecdecContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_vecdec

	return p
}

func (s *VecdecContext) GetParser() antlr.Parser { return s.parser }

func (s *VecdecContext) Get_RVAR() antlr.Token { return s._RVAR }

func (s *VecdecContext) Get_ID() antlr.Token { return s._ID }

func (s *VecdecContext) GetTyppe() antlr.Token { return s.typpe }

func (s *VecdecContext) GetTyppe2() antlr.Token { return s.typpe2 }

func (s *VecdecContext) GetFirstid() antlr.Token { return s.firstid }

func (s *VecdecContext) Set_RVAR(v antlr.Token) { s._RVAR = v }

func (s *VecdecContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *VecdecContext) SetTyppe(v antlr.Token) { s.typpe = v }

func (s *VecdecContext) SetTyppe2(v antlr.Token) { s.typpe2 = v }

func (s *VecdecContext) SetFirstid(v antlr.Token) { s.firstid = v }

func (s *VecdecContext) Get_expr() IExprContext { return s._expr }

func (s *VecdecContext) Set_expr(v IExprContext) { s._expr = v }

func (s *VecdecContext) GetNewvecdec() interfaces.Instruction { return s.newvecdec }

func (s *VecdecContext) SetNewvecdec(v interfaces.Instruction) { s.newvecdec = v }

func (s *VecdecContext) RVAR() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserRVAR, 0)
}

func (s *VecdecContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *VecdecContext) DOSPTOS() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserDOSPTOS, 0)
}

func (s *VecdecContext) AllOBRA() []antlr.TerminalNode {
	return s.GetTokens(SwiftGrammarParserOBRA)
}

func (s *VecdecContext) OBRA(i int) antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserOBRA, i)
}

func (s *VecdecContext) AllCBRA() []antlr.TerminalNode {
	return s.GetTokens(SwiftGrammarParserCBRA)
}

func (s *VecdecContext) CBRA(i int) antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCBRA, i)
}

func (s *VecdecContext) IG() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserIG, 0)
}

func (s *VecdecContext) AllRINT() []antlr.TerminalNode {
	return s.GetTokens(SwiftGrammarParserRINT)
}

func (s *VecdecContext) RINT(i int) antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserRINT, i)
}

func (s *VecdecContext) AllRFLOAT() []antlr.TerminalNode {
	return s.GetTokens(SwiftGrammarParserRFLOAT)
}

func (s *VecdecContext) RFLOAT(i int) antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserRFLOAT, i)
}

func (s *VecdecContext) AllRBOOL() []antlr.TerminalNode {
	return s.GetTokens(SwiftGrammarParserRBOOL)
}

func (s *VecdecContext) RBOOL(i int) antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserRBOOL, i)
}

func (s *VecdecContext) AllRSTRING() []antlr.TerminalNode {
	return s.GetTokens(SwiftGrammarParserRSTRING)
}

func (s *VecdecContext) RSTRING(i int) antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserRSTRING, i)
}

func (s *VecdecContext) AllRCHARACTER() []antlr.TerminalNode {
	return s.GetTokens(SwiftGrammarParserRCHARACTER)
}

func (s *VecdecContext) RCHARACTER(i int) antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserRCHARACTER, i)
}

func (s *VecdecContext) Expr() IExprContext {
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

func (s *VecdecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VecdecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VecdecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterVecdec(s)
	}
}

func (s *VecdecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitVecdec(s)
	}
}

func (p *SwiftGrammarParser) Vecdec() (localctx IVecdecContext) {
	localctx = NewVecdecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, SwiftGrammarParserRULE_vecdec)
	var _la int

	p.SetState(349)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 24, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(326)

			var _m = p.Match(SwiftGrammarParserRVAR)

			localctx.(*VecdecContext)._RVAR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(327)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*VecdecContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(328)
			p.Match(SwiftGrammarParserDOSPTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(329)
			p.Match(SwiftGrammarParserOBRA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(330)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*VecdecContext).typpe = _lt

			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&62) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*VecdecContext).typpe = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(331)
			p.Match(SwiftGrammarParserCBRA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(332)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(333)
			p.Match(SwiftGrammarParserOBRA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(334)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*VecdecContext).typpe2 = _lt

			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&62) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*VecdecContext).typpe2 = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(335)
			p.Match(SwiftGrammarParserCBRA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(336)
			p.Match(SwiftGrammarParserOBRA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(337)
			p.Match(SwiftGrammarParserCBRA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*VecdecContext).newvecdec = instructions.NewVecDec((func() int {
			if localctx.(*VecdecContext).Get_RVAR() == nil {
				return 0
			} else {
				return localctx.(*VecdecContext).Get_RVAR().GetLine()
			}
		}()), (func() int {
			if localctx.(*VecdecContext).Get_RVAR() == nil {
				return 0
			} else {
				return localctx.(*VecdecContext).Get_RVAR().GetColumn()
			}
		}()), (func() string {
			if localctx.(*VecdecContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*VecdecContext).Get_ID().GetText()
			}
		}()), (func() string {
			if localctx.(*VecdecContext).GetTyppe() == nil {
				return ""
			} else {
				return localctx.(*VecdecContext).GetTyppe().GetText()
			}
		}()), (func() string {
			if localctx.(*VecdecContext).GetTyppe2() == nil {
				return ""
			} else {
				return localctx.(*VecdecContext).GetTyppe2().GetText()
			}
		}()), nil)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(339)

			var _m = p.Match(SwiftGrammarParserRVAR)

			localctx.(*VecdecContext)._RVAR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(340)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*VecdecContext).firstid = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(341)
			p.Match(SwiftGrammarParserDOSPTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(342)
			p.Match(SwiftGrammarParserOBRA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(343)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*VecdecContext).typpe = _lt

			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&62) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*VecdecContext).typpe = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(344)
			p.Match(SwiftGrammarParserCBRA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(345)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(346)

			var _x = p.expr(0)

			localctx.(*VecdecContext)._expr = _x
		}
		localctx.(*VecdecContext).newvecdec = instructions.NewVecDec((func() int {
			if localctx.(*VecdecContext).Get_RVAR() == nil {
				return 0
			} else {
				return localctx.(*VecdecContext).Get_RVAR().GetLine()
			}
		}()), (func() int {
			if localctx.(*VecdecContext).Get_RVAR() == nil {
				return 0
			} else {
				return localctx.(*VecdecContext).Get_RVAR().GetColumn()
			}
		}()), (func() string {
			if localctx.(*VecdecContext).GetFirstid() == nil {
				return ""
			} else {
				return localctx.(*VecdecContext).GetFirstid().GetText()
			}
		}()), (func() string {
			if localctx.(*VecdecContext).GetTyppe() == nil {
				return ""
			} else {
				return localctx.(*VecdecContext).GetTyppe().GetText()
			}
		}()), nil, localctx.(*VecdecContext).Get_expr().GetE())

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

// IBreakstatementContext is an interface to support dynamic dispatch.
type IBreakstatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_RBREAK returns the _RBREAK token.
	Get_RBREAK() antlr.Token

	// Set_RBREAK sets the _RBREAK token.
	Set_RBREAK(antlr.Token)

	// GetNewbreak returns the newbreak attribute.
	GetNewbreak() interfaces.Instruction

	// SetNewbreak sets the newbreak attribute.
	SetNewbreak(interfaces.Instruction)

	// Getter signatures
	RBREAK() antlr.TerminalNode

	// IsBreakstatementContext differentiates from other interfaces.
	IsBreakstatementContext()
}

type BreakstatementContext struct {
	antlr.BaseParserRuleContext
	parser   antlr.Parser
	newbreak interfaces.Instruction
	_RBREAK  antlr.Token
}

func NewEmptyBreakstatementContext() *BreakstatementContext {
	var p = new(BreakstatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_breakstatement
	return p
}

func InitEmptyBreakstatementContext(p *BreakstatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_breakstatement
}

func (*BreakstatementContext) IsBreakstatementContext() {}

func NewBreakstatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BreakstatementContext {
	var p = new(BreakstatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_breakstatement

	return p
}

func (s *BreakstatementContext) GetParser() antlr.Parser { return s.parser }

func (s *BreakstatementContext) Get_RBREAK() antlr.Token { return s._RBREAK }

func (s *BreakstatementContext) Set_RBREAK(v antlr.Token) { s._RBREAK = v }

func (s *BreakstatementContext) GetNewbreak() interfaces.Instruction { return s.newbreak }

func (s *BreakstatementContext) SetNewbreak(v interfaces.Instruction) { s.newbreak = v }

func (s *BreakstatementContext) RBREAK() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserRBREAK, 0)
}

func (s *BreakstatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BreakstatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BreakstatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterBreakstatement(s)
	}
}

func (s *BreakstatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitBreakstatement(s)
	}
}

func (p *SwiftGrammarParser) Breakstatement() (localctx IBreakstatementContext) {
	localctx = NewBreakstatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, SwiftGrammarParserRULE_breakstatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(351)

		var _m = p.Match(SwiftGrammarParserRBREAK)

		localctx.(*BreakstatementContext)._RBREAK = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	localctx.(*BreakstatementContext).newbreak = instructions.NewBreak((func() int {
		if localctx.(*BreakstatementContext).Get_RBREAK() == nil {
			return 0
		} else {
			return localctx.(*BreakstatementContext).Get_RBREAK().GetLine()
		}
	}()), (func() int {
		if localctx.(*BreakstatementContext).Get_RBREAK() == nil {
			return 0
		} else {
			return localctx.(*BreakstatementContext).Get_RBREAK().GetColumn()
		}
	}()))

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

// IContinuestatementContext is an interface to support dynamic dispatch.
type IContinuestatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_RCONTINUE returns the _RCONTINUE token.
	Get_RCONTINUE() antlr.Token

	// Set_RCONTINUE sets the _RCONTINUE token.
	Set_RCONTINUE(antlr.Token)

	// GetNewcontinue returns the newcontinue attribute.
	GetNewcontinue() interfaces.Instruction

	// SetNewcontinue sets the newcontinue attribute.
	SetNewcontinue(interfaces.Instruction)

	// Getter signatures
	RCONTINUE() antlr.TerminalNode

	// IsContinuestatementContext differentiates from other interfaces.
	IsContinuestatementContext()
}

type ContinuestatementContext struct {
	antlr.BaseParserRuleContext
	parser      antlr.Parser
	newcontinue interfaces.Instruction
	_RCONTINUE  antlr.Token
}

func NewEmptyContinuestatementContext() *ContinuestatementContext {
	var p = new(ContinuestatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_continuestatement
	return p
}

func InitEmptyContinuestatementContext(p *ContinuestatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_continuestatement
}

func (*ContinuestatementContext) IsContinuestatementContext() {}

func NewContinuestatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ContinuestatementContext {
	var p = new(ContinuestatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_continuestatement

	return p
}

func (s *ContinuestatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ContinuestatementContext) Get_RCONTINUE() antlr.Token { return s._RCONTINUE }

func (s *ContinuestatementContext) Set_RCONTINUE(v antlr.Token) { s._RCONTINUE = v }

func (s *ContinuestatementContext) GetNewcontinue() interfaces.Instruction { return s.newcontinue }

func (s *ContinuestatementContext) SetNewcontinue(v interfaces.Instruction) { s.newcontinue = v }

func (s *ContinuestatementContext) RCONTINUE() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserRCONTINUE, 0)
}

func (s *ContinuestatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContinuestatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ContinuestatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterContinuestatement(s)
	}
}

func (s *ContinuestatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitContinuestatement(s)
	}
}

func (p *SwiftGrammarParser) Continuestatement() (localctx IContinuestatementContext) {
	localctx = NewContinuestatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, SwiftGrammarParserRULE_continuestatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(354)

		var _m = p.Match(SwiftGrammarParserRCONTINUE)

		localctx.(*ContinuestatementContext)._RCONTINUE = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	localctx.(*ContinuestatementContext).newcontinue = instructions.NewContinue((func() int {
		if localctx.(*ContinuestatementContext).Get_RCONTINUE() == nil {
			return 0
		} else {
			return localctx.(*ContinuestatementContext).Get_RCONTINUE().GetLine()
		}
	}()), (func() int {
		if localctx.(*ContinuestatementContext).Get_RCONTINUE() == nil {
			return 0
		} else {
			return localctx.(*ContinuestatementContext).Get_RCONTINUE().GetColumn()
		}
	}()))

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

// ISwitchstatementContext is an interface to support dynamic dispatch.
type ISwitchstatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_RSWITCH returns the _RSWITCH token.
	Get_RSWITCH() antlr.Token

	// Set_RSWITCH sets the _RSWITCH token.
	Set_RSWITCH(antlr.Token)

	// GetEx returns the ex rule contexts.
	GetEx() IExprContext

	// Get_caselist returns the _caselist rule contexts.
	Get_caselist() ICaselistContext

	// Get_defaultstatement returns the _defaultstatement rule contexts.
	Get_defaultstatement() IDefaultstatementContext

	// SetEx sets the ex rule contexts.
	SetEx(IExprContext)

	// Set_caselist sets the _caselist rule contexts.
	Set_caselist(ICaselistContext)

	// Set_defaultstatement sets the _defaultstatement rule contexts.
	Set_defaultstatement(IDefaultstatementContext)

	// GetNewswitch returns the newswitch attribute.
	GetNewswitch() interfaces.Instruction

	// SetNewswitch sets the newswitch attribute.
	SetNewswitch(interfaces.Instruction)

	// Getter signatures
	RSWITCH() antlr.TerminalNode
	LLAVEIZQ() antlr.TerminalNode
	Caselist() ICaselistContext
	LLAVEDER() antlr.TerminalNode
	Expr() IExprContext
	Defaultstatement() IDefaultstatementContext

	// IsSwitchstatementContext differentiates from other interfaces.
	IsSwitchstatementContext()
}

type SwitchstatementContext struct {
	antlr.BaseParserRuleContext
	parser            antlr.Parser
	newswitch         interfaces.Instruction
	_RSWITCH          antlr.Token
	ex                IExprContext
	_caselist         ICaselistContext
	_defaultstatement IDefaultstatementContext
}

func NewEmptySwitchstatementContext() *SwitchstatementContext {
	var p = new(SwitchstatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_switchstatement
	return p
}

func InitEmptySwitchstatementContext(p *SwitchstatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_switchstatement
}

func (*SwitchstatementContext) IsSwitchstatementContext() {}

func NewSwitchstatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SwitchstatementContext {
	var p = new(SwitchstatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_switchstatement

	return p
}

func (s *SwitchstatementContext) GetParser() antlr.Parser { return s.parser }

func (s *SwitchstatementContext) Get_RSWITCH() antlr.Token { return s._RSWITCH }

func (s *SwitchstatementContext) Set_RSWITCH(v antlr.Token) { s._RSWITCH = v }

func (s *SwitchstatementContext) GetEx() IExprContext { return s.ex }

func (s *SwitchstatementContext) Get_caselist() ICaselistContext { return s._caselist }

func (s *SwitchstatementContext) Get_defaultstatement() IDefaultstatementContext {
	return s._defaultstatement
}

func (s *SwitchstatementContext) SetEx(v IExprContext) { s.ex = v }

func (s *SwitchstatementContext) Set_caselist(v ICaselistContext) { s._caselist = v }

func (s *SwitchstatementContext) Set_defaultstatement(v IDefaultstatementContext) {
	s._defaultstatement = v
}

func (s *SwitchstatementContext) GetNewswitch() interfaces.Instruction { return s.newswitch }

func (s *SwitchstatementContext) SetNewswitch(v interfaces.Instruction) { s.newswitch = v }

func (s *SwitchstatementContext) RSWITCH() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserRSWITCH, 0)
}

func (s *SwitchstatementContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEIZQ, 0)
}

func (s *SwitchstatementContext) Caselist() ICaselistContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICaselistContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICaselistContext)
}

func (s *SwitchstatementContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEDER, 0)
}

func (s *SwitchstatementContext) Expr() IExprContext {
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

func (s *SwitchstatementContext) Defaultstatement() IDefaultstatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDefaultstatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDefaultstatementContext)
}

func (s *SwitchstatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchstatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SwitchstatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterSwitchstatement(s)
	}
}

func (s *SwitchstatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitSwitchstatement(s)
	}
}

func (p *SwiftGrammarParser) Switchstatement() (localctx ISwitchstatementContext) {
	localctx = NewSwitchstatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, SwiftGrammarParserRULE_switchstatement)
	p.SetState(372)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 25, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(357)

			var _m = p.Match(SwiftGrammarParserRSWITCH)

			localctx.(*SwitchstatementContext)._RSWITCH = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(358)

			var _x = p.expr(0)

			localctx.(*SwitchstatementContext).ex = _x
		}
		{
			p.SetState(359)
			p.Match(SwiftGrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(360)

			var _x = p.Caselist()

			localctx.(*SwitchstatementContext)._caselist = _x
		}
		{
			p.SetState(361)
			p.Match(SwiftGrammarParserLLAVEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*SwitchstatementContext).newswitch = instructions.NewSwitch((func() int {
			if localctx.(*SwitchstatementContext).Get_RSWITCH() == nil {
				return 0
			} else {
				return localctx.(*SwitchstatementContext).Get_RSWITCH().GetLine()
			}
		}()), (func() int {
			if localctx.(*SwitchstatementContext).Get_RSWITCH() == nil {
				return 0
			} else {
				return localctx.(*SwitchstatementContext).Get_RSWITCH().GetColumn()
			}
		}()), localctx.(*SwitchstatementContext).GetEx().GetE(), localctx.(*SwitchstatementContext).Get_caselist().GetNewcaselist(), nil)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(364)

			var _m = p.Match(SwiftGrammarParserRSWITCH)

			localctx.(*SwitchstatementContext)._RSWITCH = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(365)

			var _x = p.expr(0)

			localctx.(*SwitchstatementContext).ex = _x
		}
		{
			p.SetState(366)
			p.Match(SwiftGrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(367)

			var _x = p.Caselist()

			localctx.(*SwitchstatementContext)._caselist = _x
		}
		{
			p.SetState(368)

			var _x = p.Defaultstatement()

			localctx.(*SwitchstatementContext)._defaultstatement = _x
		}
		{
			p.SetState(369)
			p.Match(SwiftGrammarParserLLAVEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*SwitchstatementContext).newswitch = instructions.NewSwitch((func() int {
			if localctx.(*SwitchstatementContext).Get_RSWITCH() == nil {
				return 0
			} else {
				return localctx.(*SwitchstatementContext).Get_RSWITCH().GetLine()
			}
		}()), (func() int {
			if localctx.(*SwitchstatementContext).Get_RSWITCH() == nil {
				return 0
			} else {
				return localctx.(*SwitchstatementContext).Get_RSWITCH().GetColumn()
			}
		}()), localctx.(*SwitchstatementContext).GetEx().GetE(), localctx.(*SwitchstatementContext).Get_caselist().GetNewcaselist(), localctx.(*SwitchstatementContext).Get_defaultstatement().GetNewdefault())

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

// ICaselistContext is an interface to support dynamic dispatch.
type ICaselistContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_case returns the _case rule contexts.
	Get_case() ICaseContext

	// Get_caselist returns the _caselist rule contexts.
	Get_caselist() ICaselistContext

	// Set_case sets the _case rule contexts.
	Set_case(ICaseContext)

	// Set_caselist sets the _caselist rule contexts.
	Set_caselist(ICaselistContext)

	// GetNewcaselist returns the newcaselist attribute.
	GetNewcaselist() []interface{}

	// SetNewcaselist sets the newcaselist attribute.
	SetNewcaselist([]interface{})

	// Getter signatures
	Case_() ICaseContext
	Caselist() ICaselistContext

	// IsCaselistContext differentiates from other interfaces.
	IsCaselistContext()
}

type CaselistContext struct {
	antlr.BaseParserRuleContext
	parser      antlr.Parser
	newcaselist []interface{}
	_case       ICaseContext
	_caselist   ICaselistContext
}

func NewEmptyCaselistContext() *CaselistContext {
	var p = new(CaselistContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_caselist
	return p
}

func InitEmptyCaselistContext(p *CaselistContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_caselist
}

func (*CaselistContext) IsCaselistContext() {}

func NewCaselistContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CaselistContext {
	var p = new(CaselistContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_caselist

	return p
}

func (s *CaselistContext) GetParser() antlr.Parser { return s.parser }

func (s *CaselistContext) Get_case() ICaseContext { return s._case }

func (s *CaselistContext) Get_caselist() ICaselistContext { return s._caselist }

func (s *CaselistContext) Set_case(v ICaseContext) { s._case = v }

func (s *CaselistContext) Set_caselist(v ICaselistContext) { s._caselist = v }

func (s *CaselistContext) GetNewcaselist() []interface{} { return s.newcaselist }

func (s *CaselistContext) SetNewcaselist(v []interface{}) { s.newcaselist = v }

func (s *CaselistContext) Case_() ICaseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICaseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICaseContext)
}

func (s *CaselistContext) Caselist() ICaselistContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICaselistContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICaselistContext)
}

func (s *CaselistContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CaselistContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CaselistContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterCaselist(s)
	}
}

func (s *CaselistContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitCaselist(s)
	}
}

func (p *SwiftGrammarParser) Caselist() (localctx ICaselistContext) {
	localctx = NewCaselistContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, SwiftGrammarParserRULE_caselist)

	localctx.(*CaselistContext).newcaselist = []interface{}{}

	p.SetState(382)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 26, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(374)

			var _x = p.Case_()

			localctx.(*CaselistContext)._case = _x
		}
		{
			p.SetState(375)

			var _x = p.Caselist()

			localctx.(*CaselistContext)._caselist = _x
		}
		localctx.(*CaselistContext).newcaselist = append(localctx.(*CaselistContext).newcaselist, localctx.(*CaselistContext).Get_case().GetNewcase())
		for _, arg := range localctx.(*CaselistContext).Get_caselist().GetNewcaselist() {
			localctx.(*CaselistContext).newcaselist = append(localctx.(*CaselistContext).newcaselist, arg)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(378)

			var _x = p.Case_()

			localctx.(*CaselistContext)._case = _x
		}
		localctx.(*CaselistContext).newcaselist = append(localctx.(*CaselistContext).newcaselist, localctx.(*CaselistContext).Get_case().GetNewcase())

	case 3:
		p.EnterOuterAlt(localctx, 3)

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

// ICaseContext is an interface to support dynamic dispatch.
type ICaseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_RCASE returns the _RCASE token.
	Get_RCASE() antlr.Token

	// Set_RCASE sets the _RCASE token.
	Set_RCASE(antlr.Token)

	// GetEx returns the ex rule contexts.
	GetEx() IExprContext

	// GetB returns the b rule contexts.
	GetB() IBlockContext

	// SetEx sets the ex rule contexts.
	SetEx(IExprContext)

	// SetB sets the b rule contexts.
	SetB(IBlockContext)

	// GetNewcase returns the newcase attribute.
	GetNewcase() interfaces.Instruction

	// SetNewcase sets the newcase attribute.
	SetNewcase(interfaces.Instruction)

	// Getter signatures
	RCASE() antlr.TerminalNode
	DOSPTOS() antlr.TerminalNode
	Expr() IExprContext
	Block() IBlockContext

	// IsCaseContext differentiates from other interfaces.
	IsCaseContext()
}

type CaseContext struct {
	antlr.BaseParserRuleContext
	parser  antlr.Parser
	newcase interfaces.Instruction
	_RCASE  antlr.Token
	ex      IExprContext
	b       IBlockContext
}

func NewEmptyCaseContext() *CaseContext {
	var p = new(CaseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_case
	return p
}

func InitEmptyCaseContext(p *CaseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_case
}

func (*CaseContext) IsCaseContext() {}

func NewCaseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CaseContext {
	var p = new(CaseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_case

	return p
}

func (s *CaseContext) GetParser() antlr.Parser { return s.parser }

func (s *CaseContext) Get_RCASE() antlr.Token { return s._RCASE }

func (s *CaseContext) Set_RCASE(v antlr.Token) { s._RCASE = v }

func (s *CaseContext) GetEx() IExprContext { return s.ex }

func (s *CaseContext) GetB() IBlockContext { return s.b }

func (s *CaseContext) SetEx(v IExprContext) { s.ex = v }

func (s *CaseContext) SetB(v IBlockContext) { s.b = v }

func (s *CaseContext) GetNewcase() interfaces.Instruction { return s.newcase }

func (s *CaseContext) SetNewcase(v interfaces.Instruction) { s.newcase = v }

func (s *CaseContext) RCASE() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserRCASE, 0)
}

func (s *CaseContext) DOSPTOS() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserDOSPTOS, 0)
}

func (s *CaseContext) Expr() IExprContext {
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

func (s *CaseContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *CaseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CaseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CaseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterCase(s)
	}
}

func (s *CaseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitCase(s)
	}
}

func (p *SwiftGrammarParser) Case_() (localctx ICaseContext) {
	localctx = NewCaseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, SwiftGrammarParserRULE_case)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(384)

		var _m = p.Match(SwiftGrammarParserRCASE)

		localctx.(*CaseContext)._RCASE = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(385)

		var _x = p.expr(0)

		localctx.(*CaseContext).ex = _x
	}
	{
		p.SetState(386)
		p.Match(SwiftGrammarParserDOSPTOS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(387)

		var _x = p.Block()

		localctx.(*CaseContext).b = _x
	}
	localctx.(*CaseContext).newcase = instructions.NewCase((func() int {
		if localctx.(*CaseContext).Get_RCASE() == nil {
			return 0
		} else {
			return localctx.(*CaseContext).Get_RCASE().GetLine()
		}
	}()), (func() int {
		if localctx.(*CaseContext).Get_RCASE() == nil {
			return 0
		} else {
			return localctx.(*CaseContext).Get_RCASE().GetColumn()
		}
	}()), localctx.(*CaseContext).GetEx().GetE(), localctx.(*CaseContext).GetB().GetBlk())

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

// IDefaultstatementContext is an interface to support dynamic dispatch.
type IDefaultstatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetB returns the b rule contexts.
	GetB() IBlockContext

	// SetB sets the b rule contexts.
	SetB(IBlockContext)

	// GetNewdefault returns the newdefault attribute.
	GetNewdefault() []interface{}

	// SetNewdefault sets the newdefault attribute.
	SetNewdefault([]interface{})

	// Getter signatures
	RDEFAULT() antlr.TerminalNode
	DOSPTOS() antlr.TerminalNode
	LLAVEDER() antlr.TerminalNode
	Block() IBlockContext

	// IsDefaultstatementContext differentiates from other interfaces.
	IsDefaultstatementContext()
}

type DefaultstatementContext struct {
	antlr.BaseParserRuleContext
	parser     antlr.Parser
	newdefault []interface{}
	b          IBlockContext
}

func NewEmptyDefaultstatementContext() *DefaultstatementContext {
	var p = new(DefaultstatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_defaultstatement
	return p
}

func InitEmptyDefaultstatementContext(p *DefaultstatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_defaultstatement
}

func (*DefaultstatementContext) IsDefaultstatementContext() {}

func NewDefaultstatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefaultstatementContext {
	var p = new(DefaultstatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_defaultstatement

	return p
}

func (s *DefaultstatementContext) GetParser() antlr.Parser { return s.parser }

func (s *DefaultstatementContext) GetB() IBlockContext { return s.b }

func (s *DefaultstatementContext) SetB(v IBlockContext) { s.b = v }

func (s *DefaultstatementContext) GetNewdefault() []interface{} { return s.newdefault }

func (s *DefaultstatementContext) SetNewdefault(v []interface{}) { s.newdefault = v }

func (s *DefaultstatementContext) RDEFAULT() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserRDEFAULT, 0)
}

func (s *DefaultstatementContext) DOSPTOS() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserDOSPTOS, 0)
}

func (s *DefaultstatementContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEDER, 0)
}

func (s *DefaultstatementContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *DefaultstatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefaultstatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefaultstatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterDefaultstatement(s)
	}
}

func (s *DefaultstatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitDefaultstatement(s)
	}
}

func (p *SwiftGrammarParser) Defaultstatement() (localctx IDefaultstatementContext) {
	localctx = NewDefaultstatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, SwiftGrammarParserRULE_defaultstatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(390)
		p.Match(SwiftGrammarParserRDEFAULT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(391)
		p.Match(SwiftGrammarParserDOSPTOS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(392)

		var _x = p.Block()

		localctx.(*DefaultstatementContext).b = _x
	}
	{
		p.SetState(393)
		p.Match(SwiftGrammarParserLLAVEDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	localctx.(*DefaultstatementContext).newdefault = localctx.(*DefaultstatementContext).GetB().GetBlk()

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

// IIfstmtContext is an interface to support dynamic dispatch.
type IIfstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_RIF returns the _RIF token.
	Get_RIF() antlr.Token

	// Set_RIF sets the _RIF token.
	Set_RIF(antlr.Token)

	// GetEx returns the ex rule contexts.
	GetEx() IExprContext

	// GetB returns the b rule contexts.
	GetB() IBlockContext

	// Get_eliflist returns the _eliflist rule contexts.
	Get_eliflist() IEliflistContext

	// Get_elsestament returns the _elsestament rule contexts.
	Get_elsestament() IElsestamentContext

	// SetEx sets the ex rule contexts.
	SetEx(IExprContext)

	// SetB sets the b rule contexts.
	SetB(IBlockContext)

	// Set_eliflist sets the _eliflist rule contexts.
	Set_eliflist(IEliflistContext)

	// Set_elsestament sets the _elsestament rule contexts.
	Set_elsestament(IElsestamentContext)

	// GetNewif returns the newif attribute.
	GetNewif() interfaces.Instruction

	// SetNewif sets the newif attribute.
	SetNewif(interfaces.Instruction)

	// Getter signatures
	RIF() antlr.TerminalNode
	LLAVEIZQ() antlr.TerminalNode
	LLAVEDER() antlr.TerminalNode
	Eliflist() IEliflistContext
	Expr() IExprContext
	Block() IBlockContext
	Elsestament() IElsestamentContext

	// IsIfstmtContext differentiates from other interfaces.
	IsIfstmtContext()
}

type IfstmtContext struct {
	antlr.BaseParserRuleContext
	parser       antlr.Parser
	newif        interfaces.Instruction
	_RIF         antlr.Token
	ex           IExprContext
	b            IBlockContext
	_eliflist    IEliflistContext
	_elsestament IElsestamentContext
}

func NewEmptyIfstmtContext() *IfstmtContext {
	var p = new(IfstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_ifstmt
	return p
}

func InitEmptyIfstmtContext(p *IfstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_ifstmt
}

func (*IfstmtContext) IsIfstmtContext() {}

func NewIfstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfstmtContext {
	var p = new(IfstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_ifstmt

	return p
}

func (s *IfstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *IfstmtContext) Get_RIF() antlr.Token { return s._RIF }

func (s *IfstmtContext) Set_RIF(v antlr.Token) { s._RIF = v }

func (s *IfstmtContext) GetEx() IExprContext { return s.ex }

func (s *IfstmtContext) GetB() IBlockContext { return s.b }

func (s *IfstmtContext) Get_eliflist() IEliflistContext { return s._eliflist }

func (s *IfstmtContext) Get_elsestament() IElsestamentContext { return s._elsestament }

func (s *IfstmtContext) SetEx(v IExprContext) { s.ex = v }

func (s *IfstmtContext) SetB(v IBlockContext) { s.b = v }

func (s *IfstmtContext) Set_eliflist(v IEliflistContext) { s._eliflist = v }

func (s *IfstmtContext) Set_elsestament(v IElsestamentContext) { s._elsestament = v }

func (s *IfstmtContext) GetNewif() interfaces.Instruction { return s.newif }

func (s *IfstmtContext) SetNewif(v interfaces.Instruction) { s.newif = v }

func (s *IfstmtContext) RIF() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserRIF, 0)
}

func (s *IfstmtContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEIZQ, 0)
}

func (s *IfstmtContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEDER, 0)
}

func (s *IfstmtContext) Eliflist() IEliflistContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEliflistContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEliflistContext)
}

func (s *IfstmtContext) Expr() IExprContext {
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

func (s *IfstmtContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *IfstmtContext) Elsestament() IElsestamentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElsestamentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElsestamentContext)
}

func (s *IfstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterIfstmt(s)
	}
}

func (s *IfstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitIfstmt(s)
	}
}

func (p *SwiftGrammarParser) Ifstmt() (localctx IIfstmtContext) {
	localctx = NewIfstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, SwiftGrammarParserRULE_ifstmt)
	p.SetState(413)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 27, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(396)

			var _m = p.Match(SwiftGrammarParserRIF)

			localctx.(*IfstmtContext)._RIF = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(397)

			var _x = p.expr(0)

			localctx.(*IfstmtContext).ex = _x
		}
		{
			p.SetState(398)
			p.Match(SwiftGrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(399)

			var _x = p.Block()

			localctx.(*IfstmtContext).b = _x
		}
		{
			p.SetState(400)
			p.Match(SwiftGrammarParserLLAVEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(401)

			var _x = p.Eliflist()

			localctx.(*IfstmtContext)._eliflist = _x
		}
		localctx.(*IfstmtContext).newif = instructions.NewIf((func() int {
			if localctx.(*IfstmtContext).Get_RIF() == nil {
				return 0
			} else {
				return localctx.(*IfstmtContext).Get_RIF().GetLine()
			}
		}()), (func() int {
			if localctx.(*IfstmtContext).Get_RIF() == nil {
				return 0
			} else {
				return localctx.(*IfstmtContext).Get_RIF().GetColumn()
			}
		}()), localctx.(*IfstmtContext).GetEx().GetE(), localctx.(*IfstmtContext).GetB().GetBlk(), localctx.(*IfstmtContext).Get_eliflist().GetNeweliflist(), nil)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(404)

			var _m = p.Match(SwiftGrammarParserRIF)

			localctx.(*IfstmtContext)._RIF = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(405)

			var _x = p.expr(0)

			localctx.(*IfstmtContext).ex = _x
		}
		{
			p.SetState(406)
			p.Match(SwiftGrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(407)

			var _x = p.Block()

			localctx.(*IfstmtContext).b = _x
		}
		{
			p.SetState(408)
			p.Match(SwiftGrammarParserLLAVEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(409)

			var _x = p.Eliflist()

			localctx.(*IfstmtContext)._eliflist = _x
		}
		{
			p.SetState(410)

			var _x = p.Elsestament()

			localctx.(*IfstmtContext)._elsestament = _x
		}
		localctx.(*IfstmtContext).newif = instructions.NewIf((func() int {
			if localctx.(*IfstmtContext).Get_RIF() == nil {
				return 0
			} else {
				return localctx.(*IfstmtContext).Get_RIF().GetLine()
			}
		}()), (func() int {
			if localctx.(*IfstmtContext).Get_RIF() == nil {
				return 0
			} else {
				return localctx.(*IfstmtContext).Get_RIF().GetColumn()
			}
		}()), localctx.(*IfstmtContext).GetEx().GetE(), localctx.(*IfstmtContext).GetB().GetBlk(), localctx.(*IfstmtContext).Get_eliflist().GetNeweliflist(), localctx.(*IfstmtContext).Get_elsestament().GetNewelse())

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

// IEliflistContext is an interface to support dynamic dispatch.
type IEliflistContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_elif returns the _elif rule contexts.
	Get_elif() IElifContext

	// Get_eliflist returns the _eliflist rule contexts.
	Get_eliflist() IEliflistContext

	// Set_elif sets the _elif rule contexts.
	Set_elif(IElifContext)

	// Set_eliflist sets the _eliflist rule contexts.
	Set_eliflist(IEliflistContext)

	// GetNeweliflist returns the neweliflist attribute.
	GetNeweliflist() []interface{}

	// SetNeweliflist sets the neweliflist attribute.
	SetNeweliflist([]interface{})

	// Getter signatures
	Elif() IElifContext
	Eliflist() IEliflistContext

	// IsEliflistContext differentiates from other interfaces.
	IsEliflistContext()
}

type EliflistContext struct {
	antlr.BaseParserRuleContext
	parser      antlr.Parser
	neweliflist []interface{}
	_elif       IElifContext
	_eliflist   IEliflistContext
}

func NewEmptyEliflistContext() *EliflistContext {
	var p = new(EliflistContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_eliflist
	return p
}

func InitEmptyEliflistContext(p *EliflistContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_eliflist
}

func (*EliflistContext) IsEliflistContext() {}

func NewEliflistContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EliflistContext {
	var p = new(EliflistContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_eliflist

	return p
}

func (s *EliflistContext) GetParser() antlr.Parser { return s.parser }

func (s *EliflistContext) Get_elif() IElifContext { return s._elif }

func (s *EliflistContext) Get_eliflist() IEliflistContext { return s._eliflist }

func (s *EliflistContext) Set_elif(v IElifContext) { s._elif = v }

func (s *EliflistContext) Set_eliflist(v IEliflistContext) { s._eliflist = v }

func (s *EliflistContext) GetNeweliflist() []interface{} { return s.neweliflist }

func (s *EliflistContext) SetNeweliflist(v []interface{}) { s.neweliflist = v }

func (s *EliflistContext) Elif() IElifContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElifContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElifContext)
}

func (s *EliflistContext) Eliflist() IEliflistContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEliflistContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEliflistContext)
}

func (s *EliflistContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EliflistContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EliflistContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterEliflist(s)
	}
}

func (s *EliflistContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitEliflist(s)
	}
}

func (p *SwiftGrammarParser) Eliflist() (localctx IEliflistContext) {
	localctx = NewEliflistContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, SwiftGrammarParserRULE_eliflist)

	localctx.(*EliflistContext).neweliflist = []interface{}{}

	p.SetState(423)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 28, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(415)

			var _x = p.Elif()

			localctx.(*EliflistContext)._elif = _x
		}
		{
			p.SetState(416)

			var _x = p.Eliflist()

			localctx.(*EliflistContext)._eliflist = _x
		}
		localctx.(*EliflistContext).neweliflist = append(localctx.(*EliflistContext).neweliflist, localctx.(*EliflistContext).Get_elif().GetNewelif())
		for _, arg := range localctx.(*EliflistContext).Get_eliflist().GetNeweliflist() {
			localctx.(*EliflistContext).neweliflist = append(localctx.(*EliflistContext).neweliflist, arg)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(419)

			var _x = p.Elif()

			localctx.(*EliflistContext)._elif = _x
		}
		localctx.(*EliflistContext).neweliflist = append(localctx.(*EliflistContext).neweliflist, localctx.(*EliflistContext).Get_elif().GetNewelif())

	case 3:
		p.EnterOuterAlt(localctx, 3)

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

// IElifContext is an interface to support dynamic dispatch.
type IElifContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_RELSE returns the _RELSE token.
	Get_RELSE() antlr.Token

	// Set_RELSE sets the _RELSE token.
	Set_RELSE(antlr.Token)

	// GetEx returns the ex rule contexts.
	GetEx() IExprContext

	// GetB returns the b rule contexts.
	GetB() IBlockContext

	// SetEx sets the ex rule contexts.
	SetEx(IExprContext)

	// SetB sets the b rule contexts.
	SetB(IBlockContext)

	// GetNewelif returns the newelif attribute.
	GetNewelif() interfaces.Instruction

	// SetNewelif sets the newelif attribute.
	SetNewelif(interfaces.Instruction)

	// Getter signatures
	RELSE() antlr.TerminalNode
	RIF() antlr.TerminalNode
	LLAVEIZQ() antlr.TerminalNode
	LLAVEDER() antlr.TerminalNode
	Expr() IExprContext
	Block() IBlockContext

	// IsElifContext differentiates from other interfaces.
	IsElifContext()
}

type ElifContext struct {
	antlr.BaseParserRuleContext
	parser  antlr.Parser
	newelif interfaces.Instruction
	_RELSE  antlr.Token
	ex      IExprContext
	b       IBlockContext
}

func NewEmptyElifContext() *ElifContext {
	var p = new(ElifContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_elif
	return p
}

func InitEmptyElifContext(p *ElifContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_elif
}

func (*ElifContext) IsElifContext() {}

func NewElifContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElifContext {
	var p = new(ElifContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_elif

	return p
}

func (s *ElifContext) GetParser() antlr.Parser { return s.parser }

func (s *ElifContext) Get_RELSE() antlr.Token { return s._RELSE }

func (s *ElifContext) Set_RELSE(v antlr.Token) { s._RELSE = v }

func (s *ElifContext) GetEx() IExprContext { return s.ex }

func (s *ElifContext) GetB() IBlockContext { return s.b }

func (s *ElifContext) SetEx(v IExprContext) { s.ex = v }

func (s *ElifContext) SetB(v IBlockContext) { s.b = v }

func (s *ElifContext) GetNewelif() interfaces.Instruction { return s.newelif }

func (s *ElifContext) SetNewelif(v interfaces.Instruction) { s.newelif = v }

func (s *ElifContext) RELSE() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserRELSE, 0)
}

func (s *ElifContext) RIF() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserRIF, 0)
}

func (s *ElifContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEIZQ, 0)
}

func (s *ElifContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEDER, 0)
}

func (s *ElifContext) Expr() IExprContext {
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

func (s *ElifContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ElifContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElifContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElifContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterElif(s)
	}
}

func (s *ElifContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitElif(s)
	}
}

func (p *SwiftGrammarParser) Elif() (localctx IElifContext) {
	localctx = NewElifContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, SwiftGrammarParserRULE_elif)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(425)

		var _m = p.Match(SwiftGrammarParserRELSE)

		localctx.(*ElifContext)._RELSE = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(426)
		p.Match(SwiftGrammarParserRIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(427)

		var _x = p.expr(0)

		localctx.(*ElifContext).ex = _x
	}
	{
		p.SetState(428)
		p.Match(SwiftGrammarParserLLAVEIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(429)

		var _x = p.Block()

		localctx.(*ElifContext).b = _x
	}
	{
		p.SetState(430)
		p.Match(SwiftGrammarParserLLAVEDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	localctx.(*ElifContext).newelif = instructions.NewIf((func() int {
		if localctx.(*ElifContext).Get_RELSE() == nil {
			return 0
		} else {
			return localctx.(*ElifContext).Get_RELSE().GetLine()
		}
	}()), (func() int {
		if localctx.(*ElifContext).Get_RELSE() == nil {
			return 0
		} else {
			return localctx.(*ElifContext).Get_RELSE().GetColumn()
		}
	}()), localctx.(*ElifContext).GetEx().GetE(), localctx.(*ElifContext).GetB().GetBlk(), nil, nil)

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

// IElsestamentContext is an interface to support dynamic dispatch.
type IElsestamentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetB returns the b rule contexts.
	GetB() IBlockContext

	// SetB sets the b rule contexts.
	SetB(IBlockContext)

	// GetNewelse returns the newelse attribute.
	GetNewelse() []interface{}

	// SetNewelse sets the newelse attribute.
	SetNewelse([]interface{})

	// Getter signatures
	RELSE() antlr.TerminalNode
	LLAVEIZQ() antlr.TerminalNode
	LLAVEDER() antlr.TerminalNode
	Block() IBlockContext

	// IsElsestamentContext differentiates from other interfaces.
	IsElsestamentContext()
}

type ElsestamentContext struct {
	antlr.BaseParserRuleContext
	parser  antlr.Parser
	newelse []interface{}
	b       IBlockContext
}

func NewEmptyElsestamentContext() *ElsestamentContext {
	var p = new(ElsestamentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_elsestament
	return p
}

func InitEmptyElsestamentContext(p *ElsestamentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_elsestament
}

func (*ElsestamentContext) IsElsestamentContext() {}

func NewElsestamentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElsestamentContext {
	var p = new(ElsestamentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_elsestament

	return p
}

func (s *ElsestamentContext) GetParser() antlr.Parser { return s.parser }

func (s *ElsestamentContext) GetB() IBlockContext { return s.b }

func (s *ElsestamentContext) SetB(v IBlockContext) { s.b = v }

func (s *ElsestamentContext) GetNewelse() []interface{} { return s.newelse }

func (s *ElsestamentContext) SetNewelse(v []interface{}) { s.newelse = v }

func (s *ElsestamentContext) RELSE() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserRELSE, 0)
}

func (s *ElsestamentContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEIZQ, 0)
}

func (s *ElsestamentContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEDER, 0)
}

func (s *ElsestamentContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ElsestamentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElsestamentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElsestamentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterElsestament(s)
	}
}

func (s *ElsestamentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitElsestament(s)
	}
}

func (p *SwiftGrammarParser) Elsestament() (localctx IElsestamentContext) {
	localctx = NewElsestamentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, SwiftGrammarParserRULE_elsestament)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(433)
		p.Match(SwiftGrammarParserRELSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(434)
		p.Match(SwiftGrammarParserLLAVEIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(435)

		var _x = p.Block()

		localctx.(*ElsestamentContext).b = _x
	}
	{
		p.SetState(436)
		p.Match(SwiftGrammarParserLLAVEDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	localctx.(*ElsestamentContext).newelse = localctx.(*ElsestamentContext).GetB().GetBlk()

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

// IPrintstmtContext is an interface to support dynamic dispatch.
type IPrintstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_RPRINT returns the _RPRINT token.
	Get_RPRINT() antlr.Token

	// Set_RPRINT sets the _RPRINT token.
	Set_RPRINT(antlr.Token)

	// Get_arguments returns the _arguments rule contexts.
	Get_arguments() IArgumentsContext

	// Set_arguments sets the _arguments rule contexts.
	Set_arguments(IArgumentsContext)

	// GetPrnt returns the prnt attribute.
	GetPrnt() interfaces.Instruction

	// SetPrnt sets the prnt attribute.
	SetPrnt(interfaces.Instruction)

	// Getter signatures
	RPRINT() antlr.TerminalNode
	PARIZQ() antlr.TerminalNode
	Arguments() IArgumentsContext
	PARDER() antlr.TerminalNode

	// IsPrintstmtContext differentiates from other interfaces.
	IsPrintstmtContext()
}

type PrintstmtContext struct {
	antlr.BaseParserRuleContext
	parser     antlr.Parser
	prnt       interfaces.Instruction
	_RPRINT    antlr.Token
	_arguments IArgumentsContext
}

func NewEmptyPrintstmtContext() *PrintstmtContext {
	var p = new(PrintstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_printstmt
	return p
}

func InitEmptyPrintstmtContext(p *PrintstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_printstmt
}

func (*PrintstmtContext) IsPrintstmtContext() {}

func NewPrintstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrintstmtContext {
	var p = new(PrintstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_printstmt

	return p
}

func (s *PrintstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *PrintstmtContext) Get_RPRINT() antlr.Token { return s._RPRINT }

func (s *PrintstmtContext) Set_RPRINT(v antlr.Token) { s._RPRINT = v }

func (s *PrintstmtContext) Get_arguments() IArgumentsContext { return s._arguments }

func (s *PrintstmtContext) Set_arguments(v IArgumentsContext) { s._arguments = v }

func (s *PrintstmtContext) GetPrnt() interfaces.Instruction { return s.prnt }

func (s *PrintstmtContext) SetPrnt(v interfaces.Instruction) { s.prnt = v }

func (s *PrintstmtContext) RPRINT() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserRPRINT, 0)
}

func (s *PrintstmtContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARIZQ, 0)
}

func (s *PrintstmtContext) Arguments() IArgumentsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentsContext)
}

func (s *PrintstmtContext) PARDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARDER, 0)
}

func (s *PrintstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrintstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterPrintstmt(s)
	}
}

func (s *PrintstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitPrintstmt(s)
	}
}

func (p *SwiftGrammarParser) Printstmt() (localctx IPrintstmtContext) {
	localctx = NewPrintstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, SwiftGrammarParserRULE_printstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(439)

		var _m = p.Match(SwiftGrammarParserRPRINT)

		localctx.(*PrintstmtContext)._RPRINT = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(440)
		p.Match(SwiftGrammarParserPARIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(441)

		var _x = p.Arguments()

		localctx.(*PrintstmtContext)._arguments = _x
	}
	{
		p.SetState(442)
		p.Match(SwiftGrammarParserPARDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	localctx.(*PrintstmtContext).prnt = instructions.NewPrint((func() int {
		if localctx.(*PrintstmtContext).Get_RPRINT() == nil {
			return 0
		} else {
			return localctx.(*PrintstmtContext).Get_RPRINT().GetLine()
		}
	}()), (func() int {
		if localctx.(*PrintstmtContext).Get_RPRINT() == nil {
			return 0
		} else {
			return localctx.(*PrintstmtContext).Get_RPRINT().GetColumn()
		}
	}()), localctx.(*PrintstmtContext).Get_arguments().GetArgs())

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

// IWhile_statementContext is an interface to support dynamic dispatch.
type IWhile_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_RWHILE returns the _RWHILE token.
	Get_RWHILE() antlr.Token

	// Set_RWHILE sets the _RWHILE token.
	Set_RWHILE(antlr.Token)

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// GetB returns the b rule contexts.
	GetB() IBlockContext

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// SetB sets the b rule contexts.
	SetB(IBlockContext)

	// GetNewwhile returns the newwhile attribute.
	GetNewwhile() interfaces.Instruction

	// SetNewwhile sets the newwhile attribute.
	SetNewwhile(interfaces.Instruction)

	// Getter signatures
	RWHILE() antlr.TerminalNode
	Expr() IExprContext
	LLAVEIZQ() antlr.TerminalNode
	LLAVEDER() antlr.TerminalNode
	Block() IBlockContext

	// IsWhile_statementContext differentiates from other interfaces.
	IsWhile_statementContext()
}

type While_statementContext struct {
	antlr.BaseParserRuleContext
	parser   antlr.Parser
	newwhile interfaces.Instruction
	_RWHILE  antlr.Token
	_expr    IExprContext
	b        IBlockContext
}

func NewEmptyWhile_statementContext() *While_statementContext {
	var p = new(While_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_while_statement
	return p
}

func InitEmptyWhile_statementContext(p *While_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_while_statement
}

func (*While_statementContext) IsWhile_statementContext() {}

func NewWhile_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *While_statementContext {
	var p = new(While_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_while_statement

	return p
}

func (s *While_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *While_statementContext) Get_RWHILE() antlr.Token { return s._RWHILE }

func (s *While_statementContext) Set_RWHILE(v antlr.Token) { s._RWHILE = v }

func (s *While_statementContext) Get_expr() IExprContext { return s._expr }

func (s *While_statementContext) GetB() IBlockContext { return s.b }

func (s *While_statementContext) Set_expr(v IExprContext) { s._expr = v }

func (s *While_statementContext) SetB(v IBlockContext) { s.b = v }

func (s *While_statementContext) GetNewwhile() interfaces.Instruction { return s.newwhile }

func (s *While_statementContext) SetNewwhile(v interfaces.Instruction) { s.newwhile = v }

func (s *While_statementContext) RWHILE() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserRWHILE, 0)
}

func (s *While_statementContext) Expr() IExprContext {
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

func (s *While_statementContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEIZQ, 0)
}

func (s *While_statementContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEDER, 0)
}

func (s *While_statementContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *While_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *While_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *While_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterWhile_statement(s)
	}
}

func (s *While_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitWhile_statement(s)
	}
}

func (p *SwiftGrammarParser) While_statement() (localctx IWhile_statementContext) {
	localctx = NewWhile_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, SwiftGrammarParserRULE_while_statement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(445)

		var _m = p.Match(SwiftGrammarParserRWHILE)

		localctx.(*While_statementContext)._RWHILE = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(446)

		var _x = p.expr(0)

		localctx.(*While_statementContext)._expr = _x
	}
	{
		p.SetState(447)
		p.Match(SwiftGrammarParserLLAVEIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(448)

		var _x = p.Block()

		localctx.(*While_statementContext).b = _x
	}
	{
		p.SetState(449)
		p.Match(SwiftGrammarParserLLAVEDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	localctx.(*While_statementContext).newwhile = instructions.NewWhile((func() int {
		if localctx.(*While_statementContext).Get_RWHILE() == nil {
			return 0
		} else {
			return localctx.(*While_statementContext).Get_RWHILE().GetLine()
		}
	}()), (func() int {
		if localctx.(*While_statementContext).Get_RWHILE() == nil {
			return 0
		} else {
			return localctx.(*While_statementContext).Get_RWHILE().GetColumn()
		}
	}()), localctx.(*While_statementContext).Get_expr().GetE(), localctx.(*While_statementContext).GetB().GetBlk())

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

// IVardecContext is an interface to support dynamic dispatch.
type IVardecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_RVAR returns the _RVAR token.
	Get_RVAR() antlr.Token

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// GetTyppe returns the typpe token.
	GetTyppe() antlr.Token

	// Set_RVAR sets the _RVAR token.
	Set_RVAR(antlr.Token)

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// SetTyppe sets the typpe token.
	SetTyppe(antlr.Token)

	// GetEx returns the ex rule contexts.
	GetEx() IExprContext

	// SetEx sets the ex rule contexts.
	SetEx(IExprContext)

	// GetNewdec returns the newdec attribute.
	GetNewdec() interfaces.Instruction

	// SetNewdec sets the newdec attribute.
	SetNewdec(interfaces.Instruction)

	// Getter signatures
	RVAR() antlr.TerminalNode
	ID() antlr.TerminalNode
	DOSPTOS() antlr.TerminalNode
	IG() antlr.TerminalNode
	Expr() IExprContext
	RINT() antlr.TerminalNode
	RFLOAT() antlr.TerminalNode
	RBOOL() antlr.TerminalNode
	RSTRING() antlr.TerminalNode
	RCHARACTER() antlr.TerminalNode
	QM() antlr.TerminalNode

	// IsVardecContext differentiates from other interfaces.
	IsVardecContext()
}

type VardecContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	newdec interfaces.Instruction
	_RVAR  antlr.Token
	_ID    antlr.Token
	typpe  antlr.Token
	ex     IExprContext
}

func NewEmptyVardecContext() *VardecContext {
	var p = new(VardecContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_vardec
	return p
}

func InitEmptyVardecContext(p *VardecContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_vardec
}

func (*VardecContext) IsVardecContext() {}

func NewVardecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VardecContext {
	var p = new(VardecContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_vardec

	return p
}

func (s *VardecContext) GetParser() antlr.Parser { return s.parser }

func (s *VardecContext) Get_RVAR() antlr.Token { return s._RVAR }

func (s *VardecContext) Get_ID() antlr.Token { return s._ID }

func (s *VardecContext) GetTyppe() antlr.Token { return s.typpe }

func (s *VardecContext) Set_RVAR(v antlr.Token) { s._RVAR = v }

func (s *VardecContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *VardecContext) SetTyppe(v antlr.Token) { s.typpe = v }

func (s *VardecContext) GetEx() IExprContext { return s.ex }

func (s *VardecContext) SetEx(v IExprContext) { s.ex = v }

func (s *VardecContext) GetNewdec() interfaces.Instruction { return s.newdec }

func (s *VardecContext) SetNewdec(v interfaces.Instruction) { s.newdec = v }

func (s *VardecContext) RVAR() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserRVAR, 0)
}

func (s *VardecContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *VardecContext) DOSPTOS() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserDOSPTOS, 0)
}

func (s *VardecContext) IG() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserIG, 0)
}

func (s *VardecContext) Expr() IExprContext {
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

func (s *VardecContext) RINT() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserRINT, 0)
}

func (s *VardecContext) RFLOAT() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserRFLOAT, 0)
}

func (s *VardecContext) RBOOL() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserRBOOL, 0)
}

func (s *VardecContext) RSTRING() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserRSTRING, 0)
}

func (s *VardecContext) RCHARACTER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserRCHARACTER, 0)
}

func (s *VardecContext) QM() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserQM, 0)
}

func (s *VardecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VardecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VardecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterVardec(s)
	}
}

func (s *VardecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitVardec(s)
	}
}

func (p *SwiftGrammarParser) Vardec() (localctx IVardecContext) {
	localctx = NewVardecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, SwiftGrammarParserRULE_vardec)
	var _la int

	p.SetState(472)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 29, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(452)

			var _m = p.Match(SwiftGrammarParserRVAR)

			localctx.(*VardecContext)._RVAR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(453)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*VardecContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(454)
			p.Match(SwiftGrammarParserDOSPTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(455)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*VardecContext).typpe = _lt

			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&62) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*VardecContext).typpe = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(456)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(457)

			var _x = p.expr(0)

			localctx.(*VardecContext).ex = _x
		}
		localctx.(*VardecContext).newdec = instructions.NewVarDec((func() int {
			if localctx.(*VardecContext).Get_RVAR() == nil {
				return 0
			} else {
				return localctx.(*VardecContext).Get_RVAR().GetLine()
			}
		}()), (func() int {
			if localctx.(*VardecContext).Get_RVAR() == nil {
				return 0
			} else {
				return localctx.(*VardecContext).Get_RVAR().GetColumn()
			}
		}()), (func() string {
			if localctx.(*VardecContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*VardecContext).Get_ID().GetText()
			}
		}()), (func() string {
			if localctx.(*VardecContext).GetTyppe() == nil {
				return ""
			} else {
				return localctx.(*VardecContext).GetTyppe().GetText()
			}
		}()), localctx.(*VardecContext).GetEx().GetE())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(460)

			var _m = p.Match(SwiftGrammarParserRVAR)

			localctx.(*VardecContext)._RVAR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(461)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*VardecContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(462)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(463)

			var _x = p.expr(0)

			localctx.(*VardecContext).ex = _x
		}
		localctx.(*VardecContext).newdec = instructions.NewVarDec((func() int {
			if localctx.(*VardecContext).Get_RVAR() == nil {
				return 0
			} else {
				return localctx.(*VardecContext).Get_RVAR().GetLine()
			}
		}()), (func() int {
			if localctx.(*VardecContext).Get_RVAR() == nil {
				return 0
			} else {
				return localctx.(*VardecContext).Get_RVAR().GetColumn()
			}
		}()), (func() string {
			if localctx.(*VardecContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*VardecContext).Get_ID().GetText()
			}
		}()), nil, localctx.(*VardecContext).GetEx().GetE())

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(466)

			var _m = p.Match(SwiftGrammarParserRVAR)

			localctx.(*VardecContext)._RVAR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(467)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*VardecContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(468)
			p.Match(SwiftGrammarParserDOSPTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(469)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*VardecContext).typpe = _lt

			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&62) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*VardecContext).typpe = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(470)
			p.Match(SwiftGrammarParserQM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*VardecContext).newdec = instructions.NewVarDec((func() int {
			if localctx.(*VardecContext).Get_RVAR() == nil {
				return 0
			} else {
				return localctx.(*VardecContext).Get_RVAR().GetLine()
			}
		}()), (func() int {
			if localctx.(*VardecContext).Get_RVAR() == nil {
				return 0
			} else {
				return localctx.(*VardecContext).Get_RVAR().GetColumn()
			}
		}()), (func() string {
			if localctx.(*VardecContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*VardecContext).Get_ID().GetText()
			}
		}()), (func() string {
			if localctx.(*VardecContext).GetTyppe() == nil {
				return ""
			} else {
				return localctx.(*VardecContext).GetTyppe().GetText()
			}
		}()), nil)

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

// IConstdecContext is an interface to support dynamic dispatch.
type IConstdecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_RLET returns the _RLET token.
	Get_RLET() antlr.Token

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// GetTyppe returns the typpe token.
	GetTyppe() antlr.Token

	// Set_RLET sets the _RLET token.
	Set_RLET(antlr.Token)

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// SetTyppe sets the typpe token.
	SetTyppe(antlr.Token)

	// GetEx returns the ex rule contexts.
	GetEx() IExprContext

	// SetEx sets the ex rule contexts.
	SetEx(IExprContext)

	// GetNewconst returns the newconst attribute.
	GetNewconst() interfaces.Instruction

	// SetNewconst sets the newconst attribute.
	SetNewconst(interfaces.Instruction)

	// Getter signatures
	RLET() antlr.TerminalNode
	ID() antlr.TerminalNode
	DOSPTOS() antlr.TerminalNode
	IG() antlr.TerminalNode
	Expr() IExprContext
	RINT() antlr.TerminalNode
	RFLOAT() antlr.TerminalNode
	RBOOL() antlr.TerminalNode
	RSTRING() antlr.TerminalNode
	RCHARACTER() antlr.TerminalNode

	// IsConstdecContext differentiates from other interfaces.
	IsConstdecContext()
}

type ConstdecContext struct {
	antlr.BaseParserRuleContext
	parser   antlr.Parser
	newconst interfaces.Instruction
	_RLET    antlr.Token
	_ID      antlr.Token
	typpe    antlr.Token
	ex       IExprContext
}

func NewEmptyConstdecContext() *ConstdecContext {
	var p = new(ConstdecContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_constdec
	return p
}

func InitEmptyConstdecContext(p *ConstdecContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_constdec
}

func (*ConstdecContext) IsConstdecContext() {}

func NewConstdecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstdecContext {
	var p = new(ConstdecContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_constdec

	return p
}

func (s *ConstdecContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstdecContext) Get_RLET() antlr.Token { return s._RLET }

func (s *ConstdecContext) Get_ID() antlr.Token { return s._ID }

func (s *ConstdecContext) GetTyppe() antlr.Token { return s.typpe }

func (s *ConstdecContext) Set_RLET(v antlr.Token) { s._RLET = v }

func (s *ConstdecContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *ConstdecContext) SetTyppe(v antlr.Token) { s.typpe = v }

func (s *ConstdecContext) GetEx() IExprContext { return s.ex }

func (s *ConstdecContext) SetEx(v IExprContext) { s.ex = v }

func (s *ConstdecContext) GetNewconst() interfaces.Instruction { return s.newconst }

func (s *ConstdecContext) SetNewconst(v interfaces.Instruction) { s.newconst = v }

func (s *ConstdecContext) RLET() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserRLET, 0)
}

func (s *ConstdecContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *ConstdecContext) DOSPTOS() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserDOSPTOS, 0)
}

func (s *ConstdecContext) IG() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserIG, 0)
}

func (s *ConstdecContext) Expr() IExprContext {
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

func (s *ConstdecContext) RINT() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserRINT, 0)
}

func (s *ConstdecContext) RFLOAT() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserRFLOAT, 0)
}

func (s *ConstdecContext) RBOOL() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserRBOOL, 0)
}

func (s *ConstdecContext) RSTRING() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserRSTRING, 0)
}

func (s *ConstdecContext) RCHARACTER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserRCHARACTER, 0)
}

func (s *ConstdecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstdecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstdecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterConstdec(s)
	}
}

func (s *ConstdecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitConstdec(s)
	}
}

func (p *SwiftGrammarParser) Constdec() (localctx IConstdecContext) {
	localctx = NewConstdecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, SwiftGrammarParserRULE_constdec)
	var _la int

	p.SetState(488)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 30, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(474)

			var _m = p.Match(SwiftGrammarParserRLET)

			localctx.(*ConstdecContext)._RLET = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(475)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*ConstdecContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(476)
			p.Match(SwiftGrammarParserDOSPTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(477)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*ConstdecContext).typpe = _lt

			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&62) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*ConstdecContext).typpe = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(478)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(479)

			var _x = p.expr(0)

			localctx.(*ConstdecContext).ex = _x
		}
		localctx.(*ConstdecContext).newconst = instructions.NewConstDec((func() int {
			if localctx.(*ConstdecContext).Get_RLET() == nil {
				return 0
			} else {
				return localctx.(*ConstdecContext).Get_RLET().GetLine()
			}
		}()), (func() int {
			if localctx.(*ConstdecContext).Get_RLET() == nil {
				return 0
			} else {
				return localctx.(*ConstdecContext).Get_RLET().GetColumn()
			}
		}()), (func() string {
			if localctx.(*ConstdecContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*ConstdecContext).Get_ID().GetText()
			}
		}()), (func() string {
			if localctx.(*ConstdecContext).GetTyppe() == nil {
				return ""
			} else {
				return localctx.(*ConstdecContext).GetTyppe().GetText()
			}
		}()), localctx.(*ConstdecContext).GetEx().GetE())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(482)

			var _m = p.Match(SwiftGrammarParserRLET)

			localctx.(*ConstdecContext)._RLET = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(483)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*ConstdecContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(484)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(485)

			var _x = p.expr(0)

			localctx.(*ConstdecContext).ex = _x
		}
		localctx.(*ConstdecContext).newconst = instructions.NewConstDec((func() int {
			if localctx.(*ConstdecContext).Get_RLET() == nil {
				return 0
			} else {
				return localctx.(*ConstdecContext).Get_RLET().GetLine()
			}
		}()), (func() int {
			if localctx.(*ConstdecContext).Get_RLET() == nil {
				return 0
			} else {
				return localctx.(*ConstdecContext).Get_RLET().GetColumn()
			}
		}()), (func() string {
			if localctx.(*ConstdecContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*ConstdecContext).Get_ID().GetText()
			}
		}()), nil, localctx.(*ConstdecContext).GetEx().GetE())

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

// IAsignationContext is an interface to support dynamic dispatch.
type IAsignationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// GetEx returns the ex rule contexts.
	GetEx() IExprContext

	// SetEx sets the ex rule contexts.
	SetEx(IExprContext)

	// GetNewasignation returns the newasignation attribute.
	GetNewasignation() interfaces.Instruction

	// SetNewasignation sets the newasignation attribute.
	SetNewasignation(interfaces.Instruction)

	// Getter signatures
	ID() antlr.TerminalNode
	IG() antlr.TerminalNode
	Expr() IExprContext

	// IsAsignationContext differentiates from other interfaces.
	IsAsignationContext()
}

type AsignationContext struct {
	antlr.BaseParserRuleContext
	parser        antlr.Parser
	newasignation interfaces.Instruction
	_ID           antlr.Token
	ex            IExprContext
}

func NewEmptyAsignationContext() *AsignationContext {
	var p = new(AsignationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_asignation
	return p
}

func InitEmptyAsignationContext(p *AsignationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_asignation
}

func (*AsignationContext) IsAsignationContext() {}

func NewAsignationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AsignationContext {
	var p = new(AsignationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_asignation

	return p
}

func (s *AsignationContext) GetParser() antlr.Parser { return s.parser }

func (s *AsignationContext) Get_ID() antlr.Token { return s._ID }

func (s *AsignationContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *AsignationContext) GetEx() IExprContext { return s.ex }

func (s *AsignationContext) SetEx(v IExprContext) { s.ex = v }

func (s *AsignationContext) GetNewasignation() interfaces.Instruction { return s.newasignation }

func (s *AsignationContext) SetNewasignation(v interfaces.Instruction) { s.newasignation = v }

func (s *AsignationContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *AsignationContext) IG() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserIG, 0)
}

func (s *AsignationContext) Expr() IExprContext {
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

func (s *AsignationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AsignationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AsignationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterAsignation(s)
	}
}

func (s *AsignationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitAsignation(s)
	}
}

func (p *SwiftGrammarParser) Asignation() (localctx IAsignationContext) {
	localctx = NewAsignationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, SwiftGrammarParserRULE_asignation)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(490)

		var _m = p.Match(SwiftGrammarParserID)

		localctx.(*AsignationContext)._ID = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(491)
		p.Match(SwiftGrammarParserIG)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(492)

		var _x = p.expr(0)

		localctx.(*AsignationContext).ex = _x
	}
	localctx.(*AsignationContext).newasignation = instructions.NewAsignation((func() int {
		if localctx.(*AsignationContext).Get_ID() == nil {
			return 0
		} else {
			return localctx.(*AsignationContext).Get_ID().GetLine()
		}
	}()), (func() int {
		if localctx.(*AsignationContext).Get_ID() == nil {
			return 0
		} else {
			return localctx.(*AsignationContext).Get_ID().GetColumn()
		}
	}()), (func() string {
		if localctx.(*AsignationContext).Get_ID() == nil {
			return ""
		} else {
			return localctx.(*AsignationContext).Get_ID().GetText()
		}
	}()), localctx.(*AsignationContext).GetEx().GetE())

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

// IUnarysumContext is an interface to support dynamic dispatch.
type IUnarysumContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// GetEx returns the ex rule contexts.
	GetEx() IExprContext

	// SetEx sets the ex rule contexts.
	SetEx(IExprContext)

	// GetNewunarysum returns the newunarysum attribute.
	GetNewunarysum() interfaces.Instruction

	// SetNewunarysum sets the newunarysum attribute.
	SetNewunarysum(interfaces.Instruction)

	// Getter signatures
	ID() antlr.TerminalNode
	UNARYPLUS() antlr.TerminalNode
	Expr() IExprContext

	// IsUnarysumContext differentiates from other interfaces.
	IsUnarysumContext()
}

type UnarysumContext struct {
	antlr.BaseParserRuleContext
	parser      antlr.Parser
	newunarysum interfaces.Instruction
	_ID         antlr.Token
	ex          IExprContext
}

func NewEmptyUnarysumContext() *UnarysumContext {
	var p = new(UnarysumContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_unarysum
	return p
}

func InitEmptyUnarysumContext(p *UnarysumContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_unarysum
}

func (*UnarysumContext) IsUnarysumContext() {}

func NewUnarysumContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnarysumContext {
	var p = new(UnarysumContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_unarysum

	return p
}

func (s *UnarysumContext) GetParser() antlr.Parser { return s.parser }

func (s *UnarysumContext) Get_ID() antlr.Token { return s._ID }

func (s *UnarysumContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *UnarysumContext) GetEx() IExprContext { return s.ex }

func (s *UnarysumContext) SetEx(v IExprContext) { s.ex = v }

func (s *UnarysumContext) GetNewunarysum() interfaces.Instruction { return s.newunarysum }

func (s *UnarysumContext) SetNewunarysum(v interfaces.Instruction) { s.newunarysum = v }

func (s *UnarysumContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *UnarysumContext) UNARYPLUS() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserUNARYPLUS, 0)
}

func (s *UnarysumContext) Expr() IExprContext {
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

func (s *UnarysumContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnarysumContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnarysumContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterUnarysum(s)
	}
}

func (s *UnarysumContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitUnarysum(s)
	}
}

func (p *SwiftGrammarParser) Unarysum() (localctx IUnarysumContext) {
	localctx = NewUnarysumContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, SwiftGrammarParserRULE_unarysum)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(495)

		var _m = p.Match(SwiftGrammarParserID)

		localctx.(*UnarysumContext)._ID = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(496)
		p.Match(SwiftGrammarParserUNARYPLUS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(497)

		var _x = p.expr(0)

		localctx.(*UnarysumContext).ex = _x
	}
	localctx.(*UnarysumContext).newunarysum = instructions.NewUnarySum((func() int {
		if localctx.(*UnarysumContext).Get_ID() == nil {
			return 0
		} else {
			return localctx.(*UnarysumContext).Get_ID().GetLine()
		}
	}()), (func() int {
		if localctx.(*UnarysumContext).Get_ID() == nil {
			return 0
		} else {
			return localctx.(*UnarysumContext).Get_ID().GetColumn()
		}
	}()), (func() string {
		if localctx.(*UnarysumContext).Get_ID() == nil {
			return ""
		} else {
			return localctx.(*UnarysumContext).Get_ID().GetText()
		}
	}()), "+=", localctx.(*UnarysumContext).GetEx().GetE())

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

// IUnarysubContext is an interface to support dynamic dispatch.
type IUnarysubContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// GetEx returns the ex rule contexts.
	GetEx() IExprContext

	// SetEx sets the ex rule contexts.
	SetEx(IExprContext)

	// GetNewunarysub returns the newunarysub attribute.
	GetNewunarysub() interfaces.Instruction

	// SetNewunarysub sets the newunarysub attribute.
	SetNewunarysub(interfaces.Instruction)

	// Getter signatures
	ID() antlr.TerminalNode
	UNARYMINUS() antlr.TerminalNode
	Expr() IExprContext

	// IsUnarysubContext differentiates from other interfaces.
	IsUnarysubContext()
}

type UnarysubContext struct {
	antlr.BaseParserRuleContext
	parser      antlr.Parser
	newunarysub interfaces.Instruction
	_ID         antlr.Token
	ex          IExprContext
}

func NewEmptyUnarysubContext() *UnarysubContext {
	var p = new(UnarysubContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_unarysub
	return p
}

func InitEmptyUnarysubContext(p *UnarysubContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_unarysub
}

func (*UnarysubContext) IsUnarysubContext() {}

func NewUnarysubContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnarysubContext {
	var p = new(UnarysubContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_unarysub

	return p
}

func (s *UnarysubContext) GetParser() antlr.Parser { return s.parser }

func (s *UnarysubContext) Get_ID() antlr.Token { return s._ID }

func (s *UnarysubContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *UnarysubContext) GetEx() IExprContext { return s.ex }

func (s *UnarysubContext) SetEx(v IExprContext) { s.ex = v }

func (s *UnarysubContext) GetNewunarysub() interfaces.Instruction { return s.newunarysub }

func (s *UnarysubContext) SetNewunarysub(v interfaces.Instruction) { s.newunarysub = v }

func (s *UnarysubContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *UnarysubContext) UNARYMINUS() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserUNARYMINUS, 0)
}

func (s *UnarysubContext) Expr() IExprContext {
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

func (s *UnarysubContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnarysubContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnarysubContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterUnarysub(s)
	}
}

func (s *UnarysubContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitUnarysub(s)
	}
}

func (p *SwiftGrammarParser) Unarysub() (localctx IUnarysubContext) {
	localctx = NewUnarysubContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, SwiftGrammarParserRULE_unarysub)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(500)

		var _m = p.Match(SwiftGrammarParserID)

		localctx.(*UnarysubContext)._ID = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(501)
		p.Match(SwiftGrammarParserUNARYMINUS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(502)

		var _x = p.expr(0)

		localctx.(*UnarysubContext).ex = _x
	}
	localctx.(*UnarysubContext).newunarysub = instructions.NewUnarySum((func() int {
		if localctx.(*UnarysubContext).Get_ID() == nil {
			return 0
		} else {
			return localctx.(*UnarysubContext).Get_ID().GetLine()
		}
	}()), (func() int {
		if localctx.(*UnarysubContext).Get_ID() == nil {
			return 0
		} else {
			return localctx.(*UnarysubContext).Get_ID().GetColumn()
		}
	}()), (func() string {
		if localctx.(*UnarysubContext).Get_ID() == nil {
			return ""
		} else {
			return localctx.(*UnarysubContext).Get_ID().GetText()
		}
	}()), "-=", localctx.(*UnarysubContext).GetEx().GetE())

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

// IIsemptyvecContext is an interface to support dynamic dispatch.
type IIsemptyvecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// GetNewisemptyvec returns the newisemptyvec attribute.
	GetNewisemptyvec() interfaces.Expression

	// SetNewisemptyvec sets the newisemptyvec attribute.
	SetNewisemptyvec(interfaces.Expression)

	// Getter signatures
	ID() antlr.TerminalNode
	PTO() antlr.TerminalNode
	RISEMPTY() antlr.TerminalNode

	// IsIsemptyvecContext differentiates from other interfaces.
	IsIsemptyvecContext()
}

type IsemptyvecContext struct {
	antlr.BaseParserRuleContext
	parser        antlr.Parser
	newisemptyvec interfaces.Expression
	_ID           antlr.Token
}

func NewEmptyIsemptyvecContext() *IsemptyvecContext {
	var p = new(IsemptyvecContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_isemptyvec
	return p
}

func InitEmptyIsemptyvecContext(p *IsemptyvecContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_isemptyvec
}

func (*IsemptyvecContext) IsIsemptyvecContext() {}

func NewIsemptyvecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IsemptyvecContext {
	var p = new(IsemptyvecContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_isemptyvec

	return p
}

func (s *IsemptyvecContext) GetParser() antlr.Parser { return s.parser }

func (s *IsemptyvecContext) Get_ID() antlr.Token { return s._ID }

func (s *IsemptyvecContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *IsemptyvecContext) GetNewisemptyvec() interfaces.Expression { return s.newisemptyvec }

func (s *IsemptyvecContext) SetNewisemptyvec(v interfaces.Expression) { s.newisemptyvec = v }

func (s *IsemptyvecContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *IsemptyvecContext) PTO() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPTO, 0)
}

func (s *IsemptyvecContext) RISEMPTY() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserRISEMPTY, 0)
}

func (s *IsemptyvecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IsemptyvecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IsemptyvecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterIsemptyvec(s)
	}
}

func (s *IsemptyvecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitIsemptyvec(s)
	}
}

func (p *SwiftGrammarParser) Isemptyvec() (localctx IIsemptyvecContext) {
	localctx = NewIsemptyvecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, SwiftGrammarParserRULE_isemptyvec)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(505)

		var _m = p.Match(SwiftGrammarParserID)

		localctx.(*IsemptyvecContext)._ID = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(506)
		p.Match(SwiftGrammarParserPTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(507)
		p.Match(SwiftGrammarParserRISEMPTY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	localctx.(*IsemptyvecContext).newisemptyvec = expressions.NewIsEmptyVector((func() int {
		if localctx.(*IsemptyvecContext).Get_ID() == nil {
			return 0
		} else {
			return localctx.(*IsemptyvecContext).Get_ID().GetLine()
		}
	}()), (func() int {
		if localctx.(*IsemptyvecContext).Get_ID() == nil {
			return 0
		} else {
			return localctx.(*IsemptyvecContext).Get_ID().GetColumn()
		}
	}()), (func() string {
		if localctx.(*IsemptyvecContext).Get_ID() == nil {
			return ""
		} else {
			return localctx.(*IsemptyvecContext).Get_ID().GetText()
		}
	}()))

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

// ICountvecContext is an interface to support dynamic dispatch.
type ICountvecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// GetNewcountvec returns the newcountvec attribute.
	GetNewcountvec() interfaces.Expression

	// SetNewcountvec sets the newcountvec attribute.
	SetNewcountvec(interfaces.Expression)

	// Getter signatures
	ID() antlr.TerminalNode
	PTO() antlr.TerminalNode
	RCOUNT() antlr.TerminalNode

	// IsCountvecContext differentiates from other interfaces.
	IsCountvecContext()
}

type CountvecContext struct {
	antlr.BaseParserRuleContext
	parser      antlr.Parser
	newcountvec interfaces.Expression
	_ID         antlr.Token
}

func NewEmptyCountvecContext() *CountvecContext {
	var p = new(CountvecContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_countvec
	return p
}

func InitEmptyCountvecContext(p *CountvecContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_countvec
}

func (*CountvecContext) IsCountvecContext() {}

func NewCountvecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CountvecContext {
	var p = new(CountvecContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_countvec

	return p
}

func (s *CountvecContext) GetParser() antlr.Parser { return s.parser }

func (s *CountvecContext) Get_ID() antlr.Token { return s._ID }

func (s *CountvecContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *CountvecContext) GetNewcountvec() interfaces.Expression { return s.newcountvec }

func (s *CountvecContext) SetNewcountvec(v interfaces.Expression) { s.newcountvec = v }

func (s *CountvecContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *CountvecContext) PTO() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPTO, 0)
}

func (s *CountvecContext) RCOUNT() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserRCOUNT, 0)
}

func (s *CountvecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CountvecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CountvecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterCountvec(s)
	}
}

func (s *CountvecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitCountvec(s)
	}
}

func (p *SwiftGrammarParser) Countvec() (localctx ICountvecContext) {
	localctx = NewCountvecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, SwiftGrammarParserRULE_countvec)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(510)

		var _m = p.Match(SwiftGrammarParserID)

		localctx.(*CountvecContext)._ID = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(511)
		p.Match(SwiftGrammarParserPTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(512)
		p.Match(SwiftGrammarParserRCOUNT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	localctx.(*CountvecContext).newcountvec = expressions.NewCountVector((func() int {
		if localctx.(*CountvecContext).Get_ID() == nil {
			return 0
		} else {
			return localctx.(*CountvecContext).Get_ID().GetLine()
		}
	}()), (func() int {
		if localctx.(*CountvecContext).Get_ID() == nil {
			return 0
		} else {
			return localctx.(*CountvecContext).Get_ID().GetColumn()
		}
	}()), (func() string {
		if localctx.(*CountvecContext).Get_ID() == nil {
			return ""
		} else {
			return localctx.(*CountvecContext).Get_ID().GetText()
		}
	}()))

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

// IVectoraccessContext is an interface to support dynamic dispatch.
type IVectoraccessContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_indexesList returns the _indexesList rule contexts.
	Get_indexesList() IIndexesListContext

	// Set_indexesList sets the _indexesList rule contexts.
	Set_indexesList(IIndexesListContext)

	// GetNewvecaccess returns the newvecaccess attribute.
	GetNewvecaccess() interfaces.Expression

	// SetNewvecaccess sets the newvecaccess attribute.
	SetNewvecaccess(interfaces.Expression)

	// Getter signatures
	ID() antlr.TerminalNode
	IndexesList() IIndexesListContext

	// IsVectoraccessContext differentiates from other interfaces.
	IsVectoraccessContext()
}

type VectoraccessContext struct {
	antlr.BaseParserRuleContext
	parser       antlr.Parser
	newvecaccess interfaces.Expression
	_ID          antlr.Token
	_indexesList IIndexesListContext
}

func NewEmptyVectoraccessContext() *VectoraccessContext {
	var p = new(VectoraccessContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_vectoraccess
	return p
}

func InitEmptyVectoraccessContext(p *VectoraccessContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_vectoraccess
}

func (*VectoraccessContext) IsVectoraccessContext() {}

func NewVectoraccessContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VectoraccessContext {
	var p = new(VectoraccessContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_vectoraccess

	return p
}

func (s *VectoraccessContext) GetParser() antlr.Parser { return s.parser }

func (s *VectoraccessContext) Get_ID() antlr.Token { return s._ID }

func (s *VectoraccessContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *VectoraccessContext) Get_indexesList() IIndexesListContext { return s._indexesList }

func (s *VectoraccessContext) Set_indexesList(v IIndexesListContext) { s._indexesList = v }

func (s *VectoraccessContext) GetNewvecaccess() interfaces.Expression { return s.newvecaccess }

func (s *VectoraccessContext) SetNewvecaccess(v interfaces.Expression) { s.newvecaccess = v }

func (s *VectoraccessContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *VectoraccessContext) IndexesList() IIndexesListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIndexesListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIndexesListContext)
}

func (s *VectoraccessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VectoraccessContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VectoraccessContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterVectoraccess(s)
	}
}

func (s *VectoraccessContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitVectoraccess(s)
	}
}

func (p *SwiftGrammarParser) Vectoraccess() (localctx IVectoraccessContext) {
	localctx = NewVectoraccessContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, SwiftGrammarParserRULE_vectoraccess)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(515)

		var _m = p.Match(SwiftGrammarParserID)

		localctx.(*VectoraccessContext)._ID = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(516)

		var _x = p.IndexesList()

		localctx.(*VectoraccessContext)._indexesList = _x
	}
	localctx.(*VectoraccessContext).newvecaccess = expressions.NewVectorAccess((func() int {
		if localctx.(*VectoraccessContext).Get_ID() == nil {
			return 0
		} else {
			return localctx.(*VectoraccessContext).Get_ID().GetLine()
		}
	}()), (func() int {
		if localctx.(*VectoraccessContext).Get_ID() == nil {
			return 0
		} else {
			return localctx.(*VectoraccessContext).Get_ID().GetColumn()
		}
	}()), (func() string {
		if localctx.(*VectoraccessContext).Get_ID() == nil {
			return ""
		} else {
			return localctx.(*VectoraccessContext).Get_ID().GetText()
		}
	}()), localctx.(*VectoraccessContext).Get_indexesList().GetIndexes())

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

// IIndexesListContext is an interface to support dynamic dispatch.
type IIndexesListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_vecac returns the _vecac rule contexts.
	Get_vecac() IVecacContext

	// GetI returns the i rule contexts.
	GetI() IIndexesListContext

	// Set_vecac sets the _vecac rule contexts.
	Set_vecac(IVecacContext)

	// SetI sets the i rule contexts.
	SetI(IIndexesListContext)

	// GetIndexes returns the indexes attribute.
	GetIndexes() []interface{}

	// SetIndexes sets the indexes attribute.
	SetIndexes([]interface{})

	// Getter signatures
	Vecac() IVecacContext
	IndexesList() IIndexesListContext

	// IsIndexesListContext differentiates from other interfaces.
	IsIndexesListContext()
}

type IndexesListContext struct {
	antlr.BaseParserRuleContext
	parser  antlr.Parser
	indexes []interface{}
	_vecac  IVecacContext
	i       IIndexesListContext
}

func NewEmptyIndexesListContext() *IndexesListContext {
	var p = new(IndexesListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_indexesList
	return p
}

func InitEmptyIndexesListContext(p *IndexesListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_indexesList
}

func (*IndexesListContext) IsIndexesListContext() {}

func NewIndexesListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IndexesListContext {
	var p = new(IndexesListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_indexesList

	return p
}

func (s *IndexesListContext) GetParser() antlr.Parser { return s.parser }

func (s *IndexesListContext) Get_vecac() IVecacContext { return s._vecac }

func (s *IndexesListContext) GetI() IIndexesListContext { return s.i }

func (s *IndexesListContext) Set_vecac(v IVecacContext) { s._vecac = v }

func (s *IndexesListContext) SetI(v IIndexesListContext) { s.i = v }

func (s *IndexesListContext) GetIndexes() []interface{} { return s.indexes }

func (s *IndexesListContext) SetIndexes(v []interface{}) { s.indexes = v }

func (s *IndexesListContext) Vecac() IVecacContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVecacContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVecacContext)
}

func (s *IndexesListContext) IndexesList() IIndexesListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIndexesListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIndexesListContext)
}

func (s *IndexesListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexesListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IndexesListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterIndexesList(s)
	}
}

func (s *IndexesListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitIndexesList(s)
	}
}

func (p *SwiftGrammarParser) IndexesList() (localctx IIndexesListContext) {
	localctx = NewIndexesListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, SwiftGrammarParserRULE_indexesList)

	localctx.(*IndexesListContext).indexes = []interface{}{}

	p.SetState(526)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 31, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(519)

			var _x = p.Vecac()

			localctx.(*IndexesListContext)._vecac = _x
		}
		{
			p.SetState(520)

			var _x = p.IndexesList()

			localctx.(*IndexesListContext).i = _x
		}
		localctx.(*IndexesListContext).indexes = append(localctx.(*IndexesListContext).indexes, localctx.(*IndexesListContext).Get_vecac().GetNewvecac())
		for _, arg := range localctx.(*IndexesListContext).GetI().GetIndexes() {
			localctx.(*IndexesListContext).indexes = append(localctx.(*IndexesListContext).indexes, arg)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(523)

			var _x = p.Vecac()

			localctx.(*IndexesListContext)._vecac = _x
		}
		localctx.(*IndexesListContext).indexes = append(localctx.(*IndexesListContext).indexes, localctx.(*IndexesListContext).Get_vecac().GetNewvecac())

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

// IVecacContext is an interface to support dynamic dispatch.
type IVecacContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// GetNewvecac returns the newvecac attribute.
	GetNewvecac() interface{}

	// SetNewvecac sets the newvecac attribute.
	SetNewvecac(interface{})

	// Getter signatures
	OBRA() antlr.TerminalNode
	Expr() IExprContext
	CBRA() antlr.TerminalNode

	// IsVecacContext differentiates from other interfaces.
	IsVecacContext()
}

type VecacContext struct {
	antlr.BaseParserRuleContext
	parser   antlr.Parser
	newvecac interface{}
	_expr    IExprContext
}

func NewEmptyVecacContext() *VecacContext {
	var p = new(VecacContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_vecac
	return p
}

func InitEmptyVecacContext(p *VecacContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_vecac
}

func (*VecacContext) IsVecacContext() {}

func NewVecacContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VecacContext {
	var p = new(VecacContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_vecac

	return p
}

func (s *VecacContext) GetParser() antlr.Parser { return s.parser }

func (s *VecacContext) Get_expr() IExprContext { return s._expr }

func (s *VecacContext) Set_expr(v IExprContext) { s._expr = v }

func (s *VecacContext) GetNewvecac() interface{} { return s.newvecac }

func (s *VecacContext) SetNewvecac(v interface{}) { s.newvecac = v }

func (s *VecacContext) OBRA() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserOBRA, 0)
}

func (s *VecacContext) Expr() IExprContext {
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

func (s *VecacContext) CBRA() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCBRA, 0)
}

func (s *VecacContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VecacContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VecacContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterVecac(s)
	}
}

func (s *VecacContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitVecac(s)
	}
}

func (p *SwiftGrammarParser) Vecac() (localctx IVecacContext) {
	localctx = NewVecacContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, SwiftGrammarParserRULE_vecac)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(528)
		p.Match(SwiftGrammarParserOBRA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(529)

		var _x = p.expr(0)

		localctx.(*VecacContext)._expr = _x
	}
	{
		p.SetState(530)
		p.Match(SwiftGrammarParserCBRA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	localctx.(*VecacContext).SetNewvecac(localctx.(*VecacContext).Get_expr().GetE())

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

// IMatrix_typeContext is an interface to support dynamic dispatch.
type IMatrix_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_OBRA returns the _OBRA token.
	Get_OBRA() antlr.Token

	// GetTyppe returns the typpe token.
	GetTyppe() antlr.Token

	// Get_CBRA returns the _CBRA token.
	Get_CBRA() antlr.Token

	// Set_OBRA sets the _OBRA token.
	Set_OBRA(antlr.Token)

	// SetTyppe sets the typpe token.
	SetTyppe(antlr.Token)

	// Set_CBRA sets the _CBRA token.
	Set_CBRA(antlr.Token)

	// Get_matrix_type returns the _matrix_type rule contexts.
	Get_matrix_type() IMatrix_typeContext

	// Set_matrix_type sets the _matrix_type rule contexts.
	Set_matrix_type(IMatrix_typeContext)

	// GetNewmatrixtype returns the newmatrixtype attribute.
	GetNewmatrixtype() string

	// SetNewmatrixtype sets the newmatrixtype attribute.
	SetNewmatrixtype(string)

	// Getter signatures
	OBRA() antlr.TerminalNode
	CBRA() antlr.TerminalNode
	RINT() antlr.TerminalNode
	RFLOAT() antlr.TerminalNode
	RBOOL() antlr.TerminalNode
	RSTRING() antlr.TerminalNode
	RCHARACTER() antlr.TerminalNode
	Matrix_type() IMatrix_typeContext

	// IsMatrix_typeContext differentiates from other interfaces.
	IsMatrix_typeContext()
}

type Matrix_typeContext struct {
	antlr.BaseParserRuleContext
	parser        antlr.Parser
	newmatrixtype string
	_OBRA         antlr.Token
	typpe         antlr.Token
	_CBRA         antlr.Token
	_matrix_type  IMatrix_typeContext
}

func NewEmptyMatrix_typeContext() *Matrix_typeContext {
	var p = new(Matrix_typeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_matrix_type
	return p
}

func InitEmptyMatrix_typeContext(p *Matrix_typeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_matrix_type
}

func (*Matrix_typeContext) IsMatrix_typeContext() {}

func NewMatrix_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Matrix_typeContext {
	var p = new(Matrix_typeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_matrix_type

	return p
}

func (s *Matrix_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Matrix_typeContext) Get_OBRA() antlr.Token { return s._OBRA }

func (s *Matrix_typeContext) GetTyppe() antlr.Token { return s.typpe }

func (s *Matrix_typeContext) Get_CBRA() antlr.Token { return s._CBRA }

func (s *Matrix_typeContext) Set_OBRA(v antlr.Token) { s._OBRA = v }

func (s *Matrix_typeContext) SetTyppe(v antlr.Token) { s.typpe = v }

func (s *Matrix_typeContext) Set_CBRA(v antlr.Token) { s._CBRA = v }

func (s *Matrix_typeContext) Get_matrix_type() IMatrix_typeContext { return s._matrix_type }

func (s *Matrix_typeContext) Set_matrix_type(v IMatrix_typeContext) { s._matrix_type = v }

func (s *Matrix_typeContext) GetNewmatrixtype() string { return s.newmatrixtype }

func (s *Matrix_typeContext) SetNewmatrixtype(v string) { s.newmatrixtype = v }

func (s *Matrix_typeContext) OBRA() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserOBRA, 0)
}

func (s *Matrix_typeContext) CBRA() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCBRA, 0)
}

func (s *Matrix_typeContext) RINT() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserRINT, 0)
}

func (s *Matrix_typeContext) RFLOAT() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserRFLOAT, 0)
}

func (s *Matrix_typeContext) RBOOL() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserRBOOL, 0)
}

func (s *Matrix_typeContext) RSTRING() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserRSTRING, 0)
}

func (s *Matrix_typeContext) RCHARACTER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserRCHARACTER, 0)
}

func (s *Matrix_typeContext) Matrix_type() IMatrix_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMatrix_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMatrix_typeContext)
}

func (s *Matrix_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Matrix_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Matrix_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterMatrix_type(s)
	}
}

func (s *Matrix_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitMatrix_type(s)
	}
}

func (p *SwiftGrammarParser) Matrix_type() (localctx IMatrix_typeContext) {
	localctx = NewMatrix_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, SwiftGrammarParserRULE_matrix_type)
	var _la int

	p.SetState(542)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 32, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(533)

			var _m = p.Match(SwiftGrammarParserOBRA)

			localctx.(*Matrix_typeContext)._OBRA = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(534)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*Matrix_typeContext).typpe = _lt

			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&62) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*Matrix_typeContext).typpe = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(535)

			var _m = p.Match(SwiftGrammarParserCBRA)

			localctx.(*Matrix_typeContext)._CBRA = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*Matrix_typeContext).newmatrixtype = (func() string {
			if localctx.(*Matrix_typeContext).Get_OBRA() == nil {
				return ""
			} else {
				return localctx.(*Matrix_typeContext).Get_OBRA().GetText()
			}
		}()) + (func() string {
			if localctx.(*Matrix_typeContext).GetTyppe() == nil {
				return ""
			} else {
				return localctx.(*Matrix_typeContext).GetTyppe().GetText()
			}
		}()) + (func() string {
			if localctx.(*Matrix_typeContext).Get_CBRA() == nil {
				return ""
			} else {
				return localctx.(*Matrix_typeContext).Get_CBRA().GetText()
			}
		}())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(537)

			var _m = p.Match(SwiftGrammarParserOBRA)

			localctx.(*Matrix_typeContext)._OBRA = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(538)

			var _x = p.Matrix_type()

			localctx.(*Matrix_typeContext)._matrix_type = _x
		}
		{
			p.SetState(539)

			var _m = p.Match(SwiftGrammarParserCBRA)

			localctx.(*Matrix_typeContext)._CBRA = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*Matrix_typeContext).newmatrixtype = (func() string {
			if localctx.(*Matrix_typeContext).Get_OBRA() == nil {
				return ""
			} else {
				return localctx.(*Matrix_typeContext).Get_OBRA().GetText()
			}
		}()) + (func() string {
			if localctx.(*Matrix_typeContext).Get_matrix_type() == nil {
				return ""
			} else {
				return p.GetTokenStream().GetTextFromTokens(localctx.(*Matrix_typeContext).Get_matrix_type().GetStart(), localctx.(*Matrix_typeContext)._matrix_type.GetStop())
			}
		}()) + (func() string {
			if localctx.(*Matrix_typeContext).Get_CBRA() == nil {
				return ""
			} else {
				return localctx.(*Matrix_typeContext).Get_CBRA().GetText()
			}
		}())

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

// IRepeatingvectorContext is an interface to support dynamic dispatch.
type IRepeatingvectorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_matrix_type returns the _matrix_type rule contexts.
	Get_matrix_type() IMatrix_typeContext

	// GetR returns the r rule contexts.
	GetR() IRepeatingvectorContext

	// Get_repeatingvector returns the _repeatingvector rule contexts.
	Get_repeatingvector() IRepeatingvectorContext

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// GetEx1 returns the ex1 rule contexts.
	GetEx1() IExprContext

	// GetEx2 returns the ex2 rule contexts.
	GetEx2() IExprContext

	// Set_matrix_type sets the _matrix_type rule contexts.
	Set_matrix_type(IMatrix_typeContext)

	// SetR sets the r rule contexts.
	SetR(IRepeatingvectorContext)

	// Set_repeatingvector sets the _repeatingvector rule contexts.
	Set_repeatingvector(IRepeatingvectorContext)

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// SetEx1 sets the ex1 rule contexts.
	SetEx1(IExprContext)

	// SetEx2 sets the ex2 rule contexts.
	SetEx2(IExprContext)

	// GetNewrepeatingvec returns the newrepeatingvec attribute.
	GetNewrepeatingvec() interfaces.Expression

	// SetNewrepeatingvec sets the newrepeatingvec attribute.
	SetNewrepeatingvec(interfaces.Expression)

	// Getter signatures
	Matrix_type() IMatrix_typeContext
	PARIZQ() antlr.TerminalNode
	RREPEATING() antlr.TerminalNode
	AllDOSPTOS() []antlr.TerminalNode
	DOSPTOS(i int) antlr.TerminalNode
	COMA() antlr.TerminalNode
	RCOUNT() antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	PARDER() antlr.TerminalNode
	Repeatingvector() IRepeatingvectorContext

	// IsRepeatingvectorContext differentiates from other interfaces.
	IsRepeatingvectorContext()
}

type RepeatingvectorContext struct {
	antlr.BaseParserRuleContext
	parser           antlr.Parser
	newrepeatingvec  interfaces.Expression
	_matrix_type     IMatrix_typeContext
	r                IRepeatingvectorContext
	_repeatingvector IRepeatingvectorContext
	_expr            IExprContext
	ex1              IExprContext
	ex2              IExprContext
}

func NewEmptyRepeatingvectorContext() *RepeatingvectorContext {
	var p = new(RepeatingvectorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_repeatingvector
	return p
}

func InitEmptyRepeatingvectorContext(p *RepeatingvectorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_repeatingvector
}

func (*RepeatingvectorContext) IsRepeatingvectorContext() {}

func NewRepeatingvectorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RepeatingvectorContext {
	var p = new(RepeatingvectorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_repeatingvector

	return p
}

func (s *RepeatingvectorContext) GetParser() antlr.Parser { return s.parser }

func (s *RepeatingvectorContext) Get_matrix_type() IMatrix_typeContext { return s._matrix_type }

func (s *RepeatingvectorContext) GetR() IRepeatingvectorContext { return s.r }

func (s *RepeatingvectorContext) Get_repeatingvector() IRepeatingvectorContext {
	return s._repeatingvector
}

func (s *RepeatingvectorContext) Get_expr() IExprContext { return s._expr }

func (s *RepeatingvectorContext) GetEx1() IExprContext { return s.ex1 }

func (s *RepeatingvectorContext) GetEx2() IExprContext { return s.ex2 }

func (s *RepeatingvectorContext) Set_matrix_type(v IMatrix_typeContext) { s._matrix_type = v }

func (s *RepeatingvectorContext) SetR(v IRepeatingvectorContext) { s.r = v }

func (s *RepeatingvectorContext) Set_repeatingvector(v IRepeatingvectorContext) {
	s._repeatingvector = v
}

func (s *RepeatingvectorContext) Set_expr(v IExprContext) { s._expr = v }

func (s *RepeatingvectorContext) SetEx1(v IExprContext) { s.ex1 = v }

func (s *RepeatingvectorContext) SetEx2(v IExprContext) { s.ex2 = v }

func (s *RepeatingvectorContext) GetNewrepeatingvec() interfaces.Expression { return s.newrepeatingvec }

func (s *RepeatingvectorContext) SetNewrepeatingvec(v interfaces.Expression) { s.newrepeatingvec = v }

func (s *RepeatingvectorContext) Matrix_type() IMatrix_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMatrix_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMatrix_typeContext)
}

func (s *RepeatingvectorContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARIZQ, 0)
}

func (s *RepeatingvectorContext) RREPEATING() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserRREPEATING, 0)
}

func (s *RepeatingvectorContext) AllDOSPTOS() []antlr.TerminalNode {
	return s.GetTokens(SwiftGrammarParserDOSPTOS)
}

func (s *RepeatingvectorContext) DOSPTOS(i int) antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserDOSPTOS, i)
}

func (s *RepeatingvectorContext) COMA() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCOMA, 0)
}

func (s *RepeatingvectorContext) RCOUNT() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserRCOUNT, 0)
}

func (s *RepeatingvectorContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *RepeatingvectorContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *RepeatingvectorContext) PARDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARDER, 0)
}

func (s *RepeatingvectorContext) Repeatingvector() IRepeatingvectorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRepeatingvectorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRepeatingvectorContext)
}

func (s *RepeatingvectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RepeatingvectorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RepeatingvectorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterRepeatingvector(s)
	}
}

func (s *RepeatingvectorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitRepeatingvector(s)
	}
}

func (p *SwiftGrammarParser) Repeatingvector() (localctx IRepeatingvectorContext) {
	localctx = NewRepeatingvectorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, SwiftGrammarParserRULE_repeatingvector)
	p.SetState(568)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 33, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(544)

			var _x = p.Matrix_type()

			localctx.(*RepeatingvectorContext)._matrix_type = _x
		}
		{
			p.SetState(545)
			p.Match(SwiftGrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(546)
			p.Match(SwiftGrammarParserRREPEATING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(547)
			p.Match(SwiftGrammarParserDOSPTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(548)

			var _x = p.Repeatingvector()

			localctx.(*RepeatingvectorContext).r = _x
			localctx.(*RepeatingvectorContext)._repeatingvector = _x
		}
		{
			p.SetState(549)
			p.Match(SwiftGrammarParserCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(550)
			p.Match(SwiftGrammarParserRCOUNT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(551)
			p.Match(SwiftGrammarParserDOSPTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(552)

			var _x = p.expr(0)

			localctx.(*RepeatingvectorContext)._expr = _x
		}
		{
			p.SetState(553)
			p.Match(SwiftGrammarParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*RepeatingvectorContext).newrepeatingvec = expressions.NewRepeatingVector((func() antlr.Token {
			if localctx.(*RepeatingvectorContext).GetR() == nil {
				return nil
			} else {
				return localctx.(*RepeatingvectorContext).GetR().GetStart()
			}
		}()).GetLine(), (func() antlr.Token {
			if localctx.(*RepeatingvectorContext).GetR() == nil {
				return nil
			} else {
				return localctx.(*RepeatingvectorContext).GetR().GetStart()
			}
		}()).GetColumn(), (func() string {
			if localctx.(*RepeatingvectorContext).Get_matrix_type() == nil {
				return ""
			} else {
				return p.GetTokenStream().GetTextFromTokens(localctx.(*RepeatingvectorContext).Get_matrix_type().GetStart(), localctx.(*RepeatingvectorContext)._matrix_type.GetStop())
			}
		}()), localctx.(*RepeatingvectorContext).Get_repeatingvector().GetNewrepeatingvec(), localctx.(*RepeatingvectorContext).Get_expr().GetE())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(556)

			var _x = p.Matrix_type()

			localctx.(*RepeatingvectorContext)._matrix_type = _x
		}
		{
			p.SetState(557)
			p.Match(SwiftGrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(558)
			p.Match(SwiftGrammarParserRREPEATING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(559)
			p.Match(SwiftGrammarParserDOSPTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(560)

			var _x = p.expr(0)

			localctx.(*RepeatingvectorContext).ex1 = _x
		}
		{
			p.SetState(561)
			p.Match(SwiftGrammarParserCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(562)
			p.Match(SwiftGrammarParserRCOUNT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(563)
			p.Match(SwiftGrammarParserDOSPTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(564)

			var _x = p.expr(0)

			localctx.(*RepeatingvectorContext).ex2 = _x
		}
		{
			p.SetState(565)
			p.Match(SwiftGrammarParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*RepeatingvectorContext).newrepeatingvec = expressions.NewRepeatingVector((func() antlr.Token {
			if localctx.(*RepeatingvectorContext).Get_matrix_type() == nil {
				return nil
			} else {
				return localctx.(*RepeatingvectorContext).Get_matrix_type().GetStart()
			}
		}()).GetLine(), (func() antlr.Token {
			if localctx.(*RepeatingvectorContext).Get_matrix_type() == nil {
				return nil
			} else {
				return localctx.(*RepeatingvectorContext).Get_matrix_type().GetStart()
			}
		}()).GetColumn(), (func() string {
			if localctx.(*RepeatingvectorContext).Get_matrix_type() == nil {
				return ""
			} else {
				return p.GetTokenStream().GetTextFromTokens(localctx.(*RepeatingvectorContext).Get_matrix_type().GetStart(), localctx.(*RepeatingvectorContext)._matrix_type.GetStop())
			}
		}()), localctx.(*RepeatingvectorContext).GetEx1().GetE(), localctx.(*RepeatingvectorContext).GetEx2().GetE())

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

// IManualdefContext is an interface to support dynamic dispatch.
type IManualdefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_manualmatrixdef returns the _manualmatrixdef rule contexts.
	Get_manualmatrixdef() IManualmatrixdefContext

	// Set_manualmatrixdef sets the _manualmatrixdef rule contexts.
	Set_manualmatrixdef(IManualmatrixdefContext)

	// GetNewmanualdef returns the newmanualdef attribute.
	GetNewmanualdef() interfaces.Expression

	// SetNewmanualdef sets the newmanualdef attribute.
	SetNewmanualdef(interfaces.Expression)

	// Getter signatures
	Manualmatrixdef() IManualmatrixdefContext

	// IsManualdefContext differentiates from other interfaces.
	IsManualdefContext()
}

type ManualdefContext struct {
	antlr.BaseParserRuleContext
	parser           antlr.Parser
	newmanualdef     interfaces.Expression
	_manualmatrixdef IManualmatrixdefContext
}

func NewEmptyManualdefContext() *ManualdefContext {
	var p = new(ManualdefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_manualdef
	return p
}

func InitEmptyManualdefContext(p *ManualdefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_manualdef
}

func (*ManualdefContext) IsManualdefContext() {}

func NewManualdefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ManualdefContext {
	var p = new(ManualdefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_manualdef

	return p
}

func (s *ManualdefContext) GetParser() antlr.Parser { return s.parser }

func (s *ManualdefContext) Get_manualmatrixdef() IManualmatrixdefContext { return s._manualmatrixdef }

func (s *ManualdefContext) Set_manualmatrixdef(v IManualmatrixdefContext) { s._manualmatrixdef = v }

func (s *ManualdefContext) GetNewmanualdef() interfaces.Expression { return s.newmanualdef }

func (s *ManualdefContext) SetNewmanualdef(v interfaces.Expression) { s.newmanualdef = v }

func (s *ManualdefContext) Manualmatrixdef() IManualmatrixdefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IManualmatrixdefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IManualmatrixdefContext)
}

func (s *ManualdefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ManualdefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ManualdefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterManualdef(s)
	}
}

func (s *ManualdefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitManualdef(s)
	}
}

func (p *SwiftGrammarParser) Manualdef() (localctx IManualdefContext) {
	localctx = NewManualdefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, SwiftGrammarParserRULE_manualdef)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(570)

		var _x = p.Manualmatrixdef()

		localctx.(*ManualdefContext)._manualmatrixdef = _x
	}
	localctx.(*ManualdefContext).newmanualdef = expressions.NewManualMatrixDef((func() antlr.Token {
		if localctx.(*ManualdefContext).Get_manualmatrixdef() == nil {
			return nil
		} else {
			return localctx.(*ManualdefContext).Get_manualmatrixdef().GetStart()
		}
	}()).GetLine(), (func() antlr.Token {
		if localctx.(*ManualdefContext).Get_manualmatrixdef() == nil {
			return nil
		} else {
			return localctx.(*ManualdefContext).Get_manualmatrixdef().GetStart()
		}
	}()).GetColumn(), localctx.(*ManualdefContext).Get_manualmatrixdef().GetNewmanualmatrixdef())

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

// IManualmatrixdefContext is an interface to support dynamic dispatch.
type IManualmatrixdefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_values2 returns the _values2 rule contexts.
	Get_values2() IValues2Context

	// Set_values2 sets the _values2 rule contexts.
	Set_values2(IValues2Context)

	// GetNewmanualmatrixdef returns the newmanualmatrixdef attribute.
	GetNewmanualmatrixdef() []interface{}

	// SetNewmanualmatrixdef sets the newmanualmatrixdef attribute.
	SetNewmanualmatrixdef([]interface{})

	// Getter signatures
	OBRA() antlr.TerminalNode
	Values2() IValues2Context
	CBRA() antlr.TerminalNode

	// IsManualmatrixdefContext differentiates from other interfaces.
	IsManualmatrixdefContext()
}

type ManualmatrixdefContext struct {
	antlr.BaseParserRuleContext
	parser             antlr.Parser
	newmanualmatrixdef []interface{}
	_values2           IValues2Context
}

func NewEmptyManualmatrixdefContext() *ManualmatrixdefContext {
	var p = new(ManualmatrixdefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_manualmatrixdef
	return p
}

func InitEmptyManualmatrixdefContext(p *ManualmatrixdefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_manualmatrixdef
}

func (*ManualmatrixdefContext) IsManualmatrixdefContext() {}

func NewManualmatrixdefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ManualmatrixdefContext {
	var p = new(ManualmatrixdefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_manualmatrixdef

	return p
}

func (s *ManualmatrixdefContext) GetParser() antlr.Parser { return s.parser }

func (s *ManualmatrixdefContext) Get_values2() IValues2Context { return s._values2 }

func (s *ManualmatrixdefContext) Set_values2(v IValues2Context) { s._values2 = v }

func (s *ManualmatrixdefContext) GetNewmanualmatrixdef() []interface{} { return s.newmanualmatrixdef }

func (s *ManualmatrixdefContext) SetNewmanualmatrixdef(v []interface{}) { s.newmanualmatrixdef = v }

func (s *ManualmatrixdefContext) OBRA() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserOBRA, 0)
}

func (s *ManualmatrixdefContext) Values2() IValues2Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValues2Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValues2Context)
}

func (s *ManualmatrixdefContext) CBRA() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCBRA, 0)
}

func (s *ManualmatrixdefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ManualmatrixdefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ManualmatrixdefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterManualmatrixdef(s)
	}
}

func (s *ManualmatrixdefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitManualmatrixdef(s)
	}
}

func (p *SwiftGrammarParser) Manualmatrixdef() (localctx IManualmatrixdefContext) {
	localctx = NewManualmatrixdefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, SwiftGrammarParserRULE_manualmatrixdef)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(573)
		p.Match(SwiftGrammarParserOBRA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(574)

		var _x = p.values2(0)

		localctx.(*ManualmatrixdefContext)._values2 = _x
	}
	{
		p.SetState(575)
		p.Match(SwiftGrammarParserCBRA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	localctx.(*ManualmatrixdefContext).newmanualmatrixdef = append(localctx.(*ManualmatrixdefContext).newmanualmatrixdef, localctx.(*ManualmatrixdefContext).Get_values2().GetNewvalueslist())

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

// IValues2Context is an interface to support dynamic dispatch.
type IValues2Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetV returns the v rule contexts.
	GetV() IValues2Context

	// Get_manualmatrixdef returns the _manualmatrixdef rule contexts.
	Get_manualmatrixdef() IManualmatrixdefContext

	// Get_arguments returns the _arguments rule contexts.
	Get_arguments() IArgumentsContext

	// SetV sets the v rule contexts.
	SetV(IValues2Context)

	// Set_manualmatrixdef sets the _manualmatrixdef rule contexts.
	Set_manualmatrixdef(IManualmatrixdefContext)

	// Set_arguments sets the _arguments rule contexts.
	Set_arguments(IArgumentsContext)

	// GetNewvalueslist returns the newvalueslist attribute.
	GetNewvalueslist() []interface{}

	// SetNewvalueslist sets the newvalueslist attribute.
	SetNewvalueslist([]interface{})

	// Getter signatures
	Manualmatrixdef() IManualmatrixdefContext
	Arguments() IArgumentsContext
	COMA() antlr.TerminalNode
	Values2() IValues2Context

	// IsValues2Context differentiates from other interfaces.
	IsValues2Context()
}

type Values2Context struct {
	antlr.BaseParserRuleContext
	parser           antlr.Parser
	newvalueslist    []interface{}
	v                IValues2Context
	_manualmatrixdef IManualmatrixdefContext
	_arguments       IArgumentsContext
}

func NewEmptyValues2Context() *Values2Context {
	var p = new(Values2Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_values2
	return p
}

func InitEmptyValues2Context(p *Values2Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_values2
}

func (*Values2Context) IsValues2Context() {}

func NewValues2Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Values2Context {
	var p = new(Values2Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_values2

	return p
}

func (s *Values2Context) GetParser() antlr.Parser { return s.parser }

func (s *Values2Context) GetV() IValues2Context { return s.v }

func (s *Values2Context) Get_manualmatrixdef() IManualmatrixdefContext { return s._manualmatrixdef }

func (s *Values2Context) Get_arguments() IArgumentsContext { return s._arguments }

func (s *Values2Context) SetV(v IValues2Context) { s.v = v }

func (s *Values2Context) Set_manualmatrixdef(v IManualmatrixdefContext) { s._manualmatrixdef = v }

func (s *Values2Context) Set_arguments(v IArgumentsContext) { s._arguments = v }

func (s *Values2Context) GetNewvalueslist() []interface{} { return s.newvalueslist }

func (s *Values2Context) SetNewvalueslist(v []interface{}) { s.newvalueslist = v }

func (s *Values2Context) Manualmatrixdef() IManualmatrixdefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IManualmatrixdefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IManualmatrixdefContext)
}

func (s *Values2Context) Arguments() IArgumentsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentsContext)
}

func (s *Values2Context) COMA() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCOMA, 0)
}

func (s *Values2Context) Values2() IValues2Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValues2Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValues2Context)
}

func (s *Values2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Values2Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Values2Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterValues2(s)
	}
}

func (s *Values2Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitValues2(s)
	}
}

func (p *SwiftGrammarParser) Values2() (localctx IValues2Context) {
	return p.values2(0)
}

func (p *SwiftGrammarParser) values2(_p int) (localctx IValues2Context) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewValues2Context(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IValues2Context = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 82
	p.EnterRecursionRule(localctx, 82, SwiftGrammarParserRULE_values2, _p)

	localctx.(*Values2Context).newvalueslist = []interface{}{}

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(585)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 34, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(579)

			var _x = p.Manualmatrixdef()

			localctx.(*Values2Context)._manualmatrixdef = _x
		}

		localctx.(*Values2Context).newvalueslist = append(localctx.(*Values2Context).newvalueslist, localctx.(*Values2Context).Get_manualmatrixdef().GetNewmanualmatrixdef()...)

	case 2:
		{
			p.SetState(582)

			var _x = p.Arguments()

			localctx.(*Values2Context)._arguments = _x
		}

		localctx.(*Values2Context).newvalueslist = append(localctx.(*Values2Context).newvalueslist, localctx.(*Values2Context).Get_arguments().GetArgs()...)

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(594)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 35, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewValues2Context(p, _parentctx, _parentState)
			localctx.(*Values2Context).v = _prevctx
			p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_values2)
			p.SetState(587)

			if !(p.Precpred(p.GetParserRuleContext(), 3)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				goto errorExit
			}
			{
				p.SetState(588)
				p.Match(SwiftGrammarParserCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(589)

				var _x = p.Manualmatrixdef()

				localctx.(*Values2Context)._manualmatrixdef = _x
			}

			localctx.(*Values2Context).newvalueslist = append(localctx.(*Values2Context).GetV().GetNewvalueslist(), localctx.(*Values2Context).Get_manualmatrixdef().GetNewmanualmatrixdef()...)

		}
		p.SetState(596)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 35, p.GetParserRuleContext())
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

// IDecmatrixContext is an interface to support dynamic dispatch.
type IDecmatrixContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_RVAR returns the _RVAR token.
	Get_RVAR() antlr.Token

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_RVAR sets the _RVAR token.
	Set_RVAR(antlr.Token)

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_manualdef returns the _manualdef rule contexts.
	Get_manualdef() IManualdefContext

	// Get_matrix_type returns the _matrix_type rule contexts.
	Get_matrix_type() IMatrix_typeContext

	// Get_repeatingvector returns the _repeatingvector rule contexts.
	Get_repeatingvector() IRepeatingvectorContext

	// Set_manualdef sets the _manualdef rule contexts.
	Set_manualdef(IManualdefContext)

	// Set_matrix_type sets the _matrix_type rule contexts.
	Set_matrix_type(IMatrix_typeContext)

	// Set_repeatingvector sets the _repeatingvector rule contexts.
	Set_repeatingvector(IRepeatingvectorContext)

	// GetNewmatrix returns the newmatrix attribute.
	GetNewmatrix() interfaces.Instruction

	// SetNewmatrix sets the newmatrix attribute.
	SetNewmatrix(interfaces.Instruction)

	// Getter signatures
	RVAR() antlr.TerminalNode
	ID() antlr.TerminalNode
	IG() antlr.TerminalNode
	Manualdef() IManualdefContext
	DOSPTOS() antlr.TerminalNode
	Matrix_type() IMatrix_typeContext
	Repeatingvector() IRepeatingvectorContext

	// IsDecmatrixContext differentiates from other interfaces.
	IsDecmatrixContext()
}

type DecmatrixContext struct {
	antlr.BaseParserRuleContext
	parser           antlr.Parser
	newmatrix        interfaces.Instruction
	_RVAR            antlr.Token
	_ID              antlr.Token
	_manualdef       IManualdefContext
	_matrix_type     IMatrix_typeContext
	_repeatingvector IRepeatingvectorContext
}

func NewEmptyDecmatrixContext() *DecmatrixContext {
	var p = new(DecmatrixContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_decmatrix
	return p
}

func InitEmptyDecmatrixContext(p *DecmatrixContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_decmatrix
}

func (*DecmatrixContext) IsDecmatrixContext() {}

func NewDecmatrixContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DecmatrixContext {
	var p = new(DecmatrixContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_decmatrix

	return p
}

func (s *DecmatrixContext) GetParser() antlr.Parser { return s.parser }

func (s *DecmatrixContext) Get_RVAR() antlr.Token { return s._RVAR }

func (s *DecmatrixContext) Get_ID() antlr.Token { return s._ID }

func (s *DecmatrixContext) Set_RVAR(v antlr.Token) { s._RVAR = v }

func (s *DecmatrixContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *DecmatrixContext) Get_manualdef() IManualdefContext { return s._manualdef }

func (s *DecmatrixContext) Get_matrix_type() IMatrix_typeContext { return s._matrix_type }

func (s *DecmatrixContext) Get_repeatingvector() IRepeatingvectorContext { return s._repeatingvector }

func (s *DecmatrixContext) Set_manualdef(v IManualdefContext) { s._manualdef = v }

func (s *DecmatrixContext) Set_matrix_type(v IMatrix_typeContext) { s._matrix_type = v }

func (s *DecmatrixContext) Set_repeatingvector(v IRepeatingvectorContext) { s._repeatingvector = v }

func (s *DecmatrixContext) GetNewmatrix() interfaces.Instruction { return s.newmatrix }

func (s *DecmatrixContext) SetNewmatrix(v interfaces.Instruction) { s.newmatrix = v }

func (s *DecmatrixContext) RVAR() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserRVAR, 0)
}

func (s *DecmatrixContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *DecmatrixContext) IG() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserIG, 0)
}

func (s *DecmatrixContext) Manualdef() IManualdefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IManualdefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IManualdefContext)
}

func (s *DecmatrixContext) DOSPTOS() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserDOSPTOS, 0)
}

func (s *DecmatrixContext) Matrix_type() IMatrix_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMatrix_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMatrix_typeContext)
}

func (s *DecmatrixContext) Repeatingvector() IRepeatingvectorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRepeatingvectorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRepeatingvectorContext)
}

func (s *DecmatrixContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DecmatrixContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DecmatrixContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterDecmatrix(s)
	}
}

func (s *DecmatrixContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitDecmatrix(s)
	}
}

func (p *SwiftGrammarParser) Decmatrix() (localctx IDecmatrixContext) {
	localctx = NewDecmatrixContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, SwiftGrammarParserRULE_decmatrix)
	p.SetState(625)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 36, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(597)

			var _m = p.Match(SwiftGrammarParserRVAR)

			localctx.(*DecmatrixContext)._RVAR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(598)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*DecmatrixContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(599)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(600)

			var _x = p.Manualdef()

			localctx.(*DecmatrixContext)._manualdef = _x
		}
		localctx.(*DecmatrixContext).newmatrix = instructions.NewMatrixDec((func() int {
			if localctx.(*DecmatrixContext).Get_RVAR() == nil {
				return 0
			} else {
				return localctx.(*DecmatrixContext).Get_RVAR().GetLine()
			}
		}()), (func() int {
			if localctx.(*DecmatrixContext).Get_RVAR() == nil {
				return 0
			} else {
				return localctx.(*DecmatrixContext).Get_RVAR().GetColumn()
			}
		}()), (func() string {
			if localctx.(*DecmatrixContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*DecmatrixContext).Get_ID().GetText()
			}
		}()), nil, localctx.(*DecmatrixContext).Get_manualdef().GetNewmanualdef())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(603)

			var _m = p.Match(SwiftGrammarParserRVAR)

			localctx.(*DecmatrixContext)._RVAR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(604)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*DecmatrixContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(605)
			p.Match(SwiftGrammarParserDOSPTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(606)

			var _x = p.Matrix_type()

			localctx.(*DecmatrixContext)._matrix_type = _x
		}
		{
			p.SetState(607)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(608)

			var _x = p.Manualdef()

			localctx.(*DecmatrixContext)._manualdef = _x
		}
		localctx.(*DecmatrixContext).newmatrix = instructions.NewMatrixDec((func() int {
			if localctx.(*DecmatrixContext).Get_RVAR() == nil {
				return 0
			} else {
				return localctx.(*DecmatrixContext).Get_RVAR().GetLine()
			}
		}()), (func() int {
			if localctx.(*DecmatrixContext).Get_RVAR() == nil {
				return 0
			} else {
				return localctx.(*DecmatrixContext).Get_RVAR().GetColumn()
			}
		}()), (func() string {
			if localctx.(*DecmatrixContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*DecmatrixContext).Get_ID().GetText()
			}
		}()), (func() string {
			if localctx.(*DecmatrixContext).Get_matrix_type() == nil {
				return ""
			} else {
				return p.GetTokenStream().GetTextFromTokens(localctx.(*DecmatrixContext).Get_matrix_type().GetStart(), localctx.(*DecmatrixContext)._matrix_type.GetStop())
			}
		}()), localctx.(*DecmatrixContext).Get_manualdef().GetNewmanualdef())

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(611)

			var _m = p.Match(SwiftGrammarParserRVAR)

			localctx.(*DecmatrixContext)._RVAR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(612)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*DecmatrixContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(613)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(614)

			var _x = p.Repeatingvector()

			localctx.(*DecmatrixContext)._repeatingvector = _x
		}
		localctx.(*DecmatrixContext).newmatrix = instructions.NewMatrixDec((func() int {
			if localctx.(*DecmatrixContext).Get_RVAR() == nil {
				return 0
			} else {
				return localctx.(*DecmatrixContext).Get_RVAR().GetLine()
			}
		}()), (func() int {
			if localctx.(*DecmatrixContext).Get_RVAR() == nil {
				return 0
			} else {
				return localctx.(*DecmatrixContext).Get_RVAR().GetColumn()
			}
		}()), (func() string {
			if localctx.(*DecmatrixContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*DecmatrixContext).Get_ID().GetText()
			}
		}()), nil, localctx.(*DecmatrixContext).Get_repeatingvector().GetNewrepeatingvec())

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(617)

			var _m = p.Match(SwiftGrammarParserRVAR)

			localctx.(*DecmatrixContext)._RVAR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(618)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*DecmatrixContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(619)
			p.Match(SwiftGrammarParserDOSPTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(620)

			var _x = p.Matrix_type()

			localctx.(*DecmatrixContext)._matrix_type = _x
		}
		{
			p.SetState(621)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(622)

			var _x = p.Repeatingvector()

			localctx.(*DecmatrixContext)._repeatingvector = _x
		}
		localctx.(*DecmatrixContext).newmatrix = instructions.NewMatrixDec((func() int {
			if localctx.(*DecmatrixContext).Get_RVAR() == nil {
				return 0
			} else {
				return localctx.(*DecmatrixContext).Get_RVAR().GetLine()
			}
		}()), (func() int {
			if localctx.(*DecmatrixContext).Get_RVAR() == nil {
				return 0
			} else {
				return localctx.(*DecmatrixContext).Get_RVAR().GetColumn()
			}
		}()), (func() string {
			if localctx.(*DecmatrixContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*DecmatrixContext).Get_ID().GetText()
			}
		}()), (func() string {
			if localctx.(*DecmatrixContext).Get_matrix_type() == nil {
				return ""
			} else {
				return p.GetTokenStream().GetTextFromTokens(localctx.(*DecmatrixContext).Get_matrix_type().GetStart(), localctx.(*DecmatrixContext)._matrix_type.GetStop())
			}
		}()), localctx.(*DecmatrixContext).Get_repeatingvector().GetNewrepeatingvec())

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

// IAttrlistContext is an interface to support dynamic dispatch.
type IAttrlistContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_attr returns the _attr rule contexts.
	Get_attr() IAttrContext

	// GetA returns the a rule contexts.
	GetA() IAttrlistContext

	// Set_attr sets the _attr rule contexts.
	Set_attr(IAttrContext)

	// SetA sets the a rule contexts.
	SetA(IAttrlistContext)

	// GetAtrlist returns the atrlist attribute.
	GetAtrlist() []string

	// SetAtrlist sets the atrlist attribute.
	SetAtrlist([]string)

	// Getter signatures
	Attr() IAttrContext
	PTO() antlr.TerminalNode
	Attrlist() IAttrlistContext

	// IsAttrlistContext differentiates from other interfaces.
	IsAttrlistContext()
}

type AttrlistContext struct {
	antlr.BaseParserRuleContext
	parser  antlr.Parser
	atrlist []string
	_attr   IAttrContext
	a       IAttrlistContext
}

func NewEmptyAttrlistContext() *AttrlistContext {
	var p = new(AttrlistContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_attrlist
	return p
}

func InitEmptyAttrlistContext(p *AttrlistContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_attrlist
}

func (*AttrlistContext) IsAttrlistContext() {}

func NewAttrlistContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttrlistContext {
	var p = new(AttrlistContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_attrlist

	return p
}

func (s *AttrlistContext) GetParser() antlr.Parser { return s.parser }

func (s *AttrlistContext) Get_attr() IAttrContext { return s._attr }

func (s *AttrlistContext) GetA() IAttrlistContext { return s.a }

func (s *AttrlistContext) Set_attr(v IAttrContext) { s._attr = v }

func (s *AttrlistContext) SetA(v IAttrlistContext) { s.a = v }

func (s *AttrlistContext) GetAtrlist() []string { return s.atrlist }

func (s *AttrlistContext) SetAtrlist(v []string) { s.atrlist = v }

func (s *AttrlistContext) Attr() IAttrContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttrContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttrContext)
}

func (s *AttrlistContext) PTO() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPTO, 0)
}

func (s *AttrlistContext) Attrlist() IAttrlistContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttrlistContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttrlistContext)
}

func (s *AttrlistContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttrlistContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AttrlistContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterAttrlist(s)
	}
}

func (s *AttrlistContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitAttrlist(s)
	}
}

func (p *SwiftGrammarParser) Attrlist() (localctx IAttrlistContext) {
	localctx = NewAttrlistContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, SwiftGrammarParserRULE_attrlist)

	localctx.(*AttrlistContext).atrlist = []string{}

	p.SetState(636)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 37, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(627)

			var _x = p.Attr()

			localctx.(*AttrlistContext)._attr = _x
		}
		localctx.(*AttrlistContext).atrlist = append(localctx.(*AttrlistContext).atrlist, localctx.(*AttrlistContext).Get_attr().GetAtr())
		{
			p.SetState(629)
			p.Match(SwiftGrammarParserPTO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(630)

			var _x = p.Attrlist()

			localctx.(*AttrlistContext).a = _x
		}
		localctx.(*AttrlistContext).atrlist = append(localctx.(*AttrlistContext).atrlist, localctx.(*AttrlistContext).GetA().GetAtrlist()...)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(633)

			var _x = p.Attr()

			localctx.(*AttrlistContext)._attr = _x
		}
		localctx.(*AttrlistContext).atrlist = append(localctx.(*AttrlistContext).atrlist, localctx.(*AttrlistContext).Get_attr().GetAtr())

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

// IAttrContext is an interface to support dynamic dispatch.
type IAttrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// GetAtr returns the atr attribute.
	GetAtr() string

	// SetAtr sets the atr attribute.
	SetAtr(string)

	// Getter signatures
	ID() antlr.TerminalNode

	// IsAttrContext differentiates from other interfaces.
	IsAttrContext()
}

type AttrContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	atr    string
	_ID    antlr.Token
}

func NewEmptyAttrContext() *AttrContext {
	var p = new(AttrContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_attr
	return p
}

func InitEmptyAttrContext(p *AttrContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_attr
}

func (*AttrContext) IsAttrContext() {}

func NewAttrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttrContext {
	var p = new(AttrContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_attr

	return p
}

func (s *AttrContext) GetParser() antlr.Parser { return s.parser }

func (s *AttrContext) Get_ID() antlr.Token { return s._ID }

func (s *AttrContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *AttrContext) GetAtr() string { return s.atr }

func (s *AttrContext) SetAtr(v string) { s.atr = v }

func (s *AttrContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *AttrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AttrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterAttr(s)
	}
}

func (s *AttrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitAttr(s)
	}
}

func (p *SwiftGrammarParser) Attr() (localctx IAttrContext) {
	localctx = NewAttrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, SwiftGrammarParserRULE_attr)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(638)

		var _m = p.Match(SwiftGrammarParserID)

		localctx.(*AttrContext)._ID = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	localctx.(*AttrContext).atr = (func() string {
		if localctx.(*AttrContext).Get_ID() == nil {
			return ""
		} else {
			return localctx.(*AttrContext).Get_ID().GetText()
		}
	}())

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

// IStructaccessContext is an interface to support dynamic dispatch.
type IStructaccessContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_attrlist returns the _attrlist rule contexts.
	Get_attrlist() IAttrlistContext

	// Set_attrlist sets the _attrlist rule contexts.
	Set_attrlist(IAttrlistContext)

	// GetSaccess returns the saccess attribute.
	GetSaccess() interfaces.Expression

	// SetSaccess sets the saccess attribute.
	SetSaccess(interfaces.Expression)

	// Getter signatures
	ID() antlr.TerminalNode
	PTO() antlr.TerminalNode
	Attrlist() IAttrlistContext

	// IsStructaccessContext differentiates from other interfaces.
	IsStructaccessContext()
}

type StructaccessContext struct {
	antlr.BaseParserRuleContext
	parser    antlr.Parser
	saccess   interfaces.Expression
	_ID       antlr.Token
	_attrlist IAttrlistContext
}

func NewEmptyStructaccessContext() *StructaccessContext {
	var p = new(StructaccessContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_structaccess
	return p
}

func InitEmptyStructaccessContext(p *StructaccessContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_structaccess
}

func (*StructaccessContext) IsStructaccessContext() {}

func NewStructaccessContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructaccessContext {
	var p = new(StructaccessContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_structaccess

	return p
}

func (s *StructaccessContext) GetParser() antlr.Parser { return s.parser }

func (s *StructaccessContext) Get_ID() antlr.Token { return s._ID }

func (s *StructaccessContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *StructaccessContext) Get_attrlist() IAttrlistContext { return s._attrlist }

func (s *StructaccessContext) Set_attrlist(v IAttrlistContext) { s._attrlist = v }

func (s *StructaccessContext) GetSaccess() interfaces.Expression { return s.saccess }

func (s *StructaccessContext) SetSaccess(v interfaces.Expression) { s.saccess = v }

func (s *StructaccessContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *StructaccessContext) PTO() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPTO, 0)
}

func (s *StructaccessContext) Attrlist() IAttrlistContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttrlistContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttrlistContext)
}

func (s *StructaccessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructaccessContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructaccessContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterStructaccess(s)
	}
}

func (s *StructaccessContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitStructaccess(s)
	}
}

func (p *SwiftGrammarParser) Structaccess() (localctx IStructaccessContext) {
	localctx = NewStructaccessContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, SwiftGrammarParserRULE_structaccess)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(641)

		var _m = p.Match(SwiftGrammarParserID)

		localctx.(*StructaccessContext)._ID = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(642)
		p.Match(SwiftGrammarParserPTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(643)

		var _x = p.Attrlist()

		localctx.(*StructaccessContext)._attrlist = _x
	}
	localctx.(*StructaccessContext).saccess = expressions.NewStructAccess((func() int {
		if localctx.(*StructaccessContext).Get_ID() == nil {
			return 0
		} else {
			return localctx.(*StructaccessContext).Get_ID().GetLine()
		}
	}()), (func() int {
		if localctx.(*StructaccessContext).Get_ID() == nil {
			return 0
		} else {
			return localctx.(*StructaccessContext).Get_ID().GetColumn()
		}
	}()), (func() string {
		if localctx.(*StructaccessContext).Get_ID() == nil {
			return ""
		} else {
			return localctx.(*StructaccessContext).Get_ID().GetText()
		}
	}()), localctx.(*StructaccessContext).Get_attrlist().GetAtrlist())

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

// IStructexpContext is an interface to support dynamic dispatch.
type IStructexpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_keyvaluelist returns the _keyvaluelist rule contexts.
	Get_keyvaluelist() IKeyvaluelistContext

	// Set_keyvaluelist sets the _keyvaluelist rule contexts.
	Set_keyvaluelist(IKeyvaluelistContext)

	// GetStructexxp returns the structexxp attribute.
	GetStructexxp() interfaces.Expression

	// SetStructexxp sets the structexxp attribute.
	SetStructexxp(interfaces.Expression)

	// Getter signatures
	ID() antlr.TerminalNode
	PARIZQ() antlr.TerminalNode
	Keyvaluelist() IKeyvaluelistContext
	PARDER() antlr.TerminalNode

	// IsStructexpContext differentiates from other interfaces.
	IsStructexpContext()
}

type StructexpContext struct {
	antlr.BaseParserRuleContext
	parser        antlr.Parser
	structexxp    interfaces.Expression
	_ID           antlr.Token
	_keyvaluelist IKeyvaluelistContext
}

func NewEmptyStructexpContext() *StructexpContext {
	var p = new(StructexpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_structexp
	return p
}

func InitEmptyStructexpContext(p *StructexpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_structexp
}

func (*StructexpContext) IsStructexpContext() {}

func NewStructexpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructexpContext {
	var p = new(StructexpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_structexp

	return p
}

func (s *StructexpContext) GetParser() antlr.Parser { return s.parser }

func (s *StructexpContext) Get_ID() antlr.Token { return s._ID }

func (s *StructexpContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *StructexpContext) Get_keyvaluelist() IKeyvaluelistContext { return s._keyvaluelist }

func (s *StructexpContext) Set_keyvaluelist(v IKeyvaluelistContext) { s._keyvaluelist = v }

func (s *StructexpContext) GetStructexxp() interfaces.Expression { return s.structexxp }

func (s *StructexpContext) SetStructexxp(v interfaces.Expression) { s.structexxp = v }

func (s *StructexpContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *StructexpContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARIZQ, 0)
}

func (s *StructexpContext) Keyvaluelist() IKeyvaluelistContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKeyvaluelistContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IKeyvaluelistContext)
}

func (s *StructexpContext) PARDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARDER, 0)
}

func (s *StructexpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructexpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructexpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterStructexp(s)
	}
}

func (s *StructexpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitStructexp(s)
	}
}

func (p *SwiftGrammarParser) Structexp() (localctx IStructexpContext) {
	localctx = NewStructexpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, SwiftGrammarParserRULE_structexp)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(646)

		var _m = p.Match(SwiftGrammarParserID)

		localctx.(*StructexpContext)._ID = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(647)
		p.Match(SwiftGrammarParserPARIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(648)

		var _x = p.Keyvaluelist()

		localctx.(*StructexpContext)._keyvaluelist = _x
	}
	{
		p.SetState(649)
		p.Match(SwiftGrammarParserPARDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	localctx.(*StructexpContext).structexxp = expressions.NewStructExp((func() int {
		if localctx.(*StructexpContext).Get_ID() == nil {
			return 0
		} else {
			return localctx.(*StructexpContext).Get_ID().GetLine()
		}
	}()), (func() int {
		if localctx.(*StructexpContext).Get_ID() == nil {
			return 0
		} else {
			return localctx.(*StructexpContext).Get_ID().GetColumn()
		}
	}()), (func() string {
		if localctx.(*StructexpContext).Get_ID() == nil {
			return ""
		} else {
			return localctx.(*StructexpContext).Get_ID().GetText()
		}
	}()), localctx.(*StructexpContext).Get_keyvaluelist().GetKvlist())

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

// IKeyvaluelistContext is an interface to support dynamic dispatch.
type IKeyvaluelistContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_keyvalue returns the _keyvalue rule contexts.
	Get_keyvalue() IKeyvalueContext

	// GetA returns the a rule contexts.
	GetA() IKeyvaluelistContext

	// Set_keyvalue sets the _keyvalue rule contexts.
	Set_keyvalue(IKeyvalueContext)

	// SetA sets the a rule contexts.
	SetA(IKeyvaluelistContext)

	// GetKvlist returns the kvlist attribute.
	GetKvlist() []environment.KeyValue

	// SetKvlist sets the kvlist attribute.
	SetKvlist([]environment.KeyValue)

	// Getter signatures
	Keyvalue() IKeyvalueContext
	COMA() antlr.TerminalNode
	Keyvaluelist() IKeyvaluelistContext

	// IsKeyvaluelistContext differentiates from other interfaces.
	IsKeyvaluelistContext()
}

type KeyvaluelistContext struct {
	antlr.BaseParserRuleContext
	parser    antlr.Parser
	kvlist    []environment.KeyValue
	_keyvalue IKeyvalueContext
	a         IKeyvaluelistContext
}

func NewEmptyKeyvaluelistContext() *KeyvaluelistContext {
	var p = new(KeyvaluelistContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_keyvaluelist
	return p
}

func InitEmptyKeyvaluelistContext(p *KeyvaluelistContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_keyvaluelist
}

func (*KeyvaluelistContext) IsKeyvaluelistContext() {}

func NewKeyvaluelistContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeyvaluelistContext {
	var p = new(KeyvaluelistContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_keyvaluelist

	return p
}

func (s *KeyvaluelistContext) GetParser() antlr.Parser { return s.parser }

func (s *KeyvaluelistContext) Get_keyvalue() IKeyvalueContext { return s._keyvalue }

func (s *KeyvaluelistContext) GetA() IKeyvaluelistContext { return s.a }

func (s *KeyvaluelistContext) Set_keyvalue(v IKeyvalueContext) { s._keyvalue = v }

func (s *KeyvaluelistContext) SetA(v IKeyvaluelistContext) { s.a = v }

func (s *KeyvaluelistContext) GetKvlist() []environment.KeyValue { return s.kvlist }

func (s *KeyvaluelistContext) SetKvlist(v []environment.KeyValue) { s.kvlist = v }

func (s *KeyvaluelistContext) Keyvalue() IKeyvalueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKeyvalueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IKeyvalueContext)
}

func (s *KeyvaluelistContext) COMA() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCOMA, 0)
}

func (s *KeyvaluelistContext) Keyvaluelist() IKeyvaluelistContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKeyvaluelistContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IKeyvaluelistContext)
}

func (s *KeyvaluelistContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeyvaluelistContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeyvaluelistContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterKeyvaluelist(s)
	}
}

func (s *KeyvaluelistContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitKeyvaluelist(s)
	}
}

func (p *SwiftGrammarParser) Keyvaluelist() (localctx IKeyvaluelistContext) {
	localctx = NewKeyvaluelistContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, SwiftGrammarParserRULE_keyvaluelist)

	localctx.(*KeyvaluelistContext).kvlist = []environment.KeyValue{}

	p.SetState(661)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 38, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(652)

			var _x = p.Keyvalue()

			localctx.(*KeyvaluelistContext)._keyvalue = _x
		}
		localctx.(*KeyvaluelistContext).kvlist = append(localctx.(*KeyvaluelistContext).kvlist, localctx.(*KeyvaluelistContext).Get_keyvalue().GetKv())
		{
			p.SetState(654)
			p.Match(SwiftGrammarParserCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(655)

			var _x = p.Keyvaluelist()

			localctx.(*KeyvaluelistContext).a = _x
		}
		localctx.(*KeyvaluelistContext).kvlist = append(localctx.(*KeyvaluelistContext).kvlist, localctx.(*KeyvaluelistContext).GetA().GetKvlist()...)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(658)

			var _x = p.Keyvalue()

			localctx.(*KeyvaluelistContext)._keyvalue = _x
		}
		localctx.(*KeyvaluelistContext).kvlist = append(localctx.(*KeyvaluelistContext).kvlist, localctx.(*KeyvaluelistContext).Get_keyvalue().GetKv())

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

// IKeyvalueContext is an interface to support dynamic dispatch.
type IKeyvalueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// GetKv returns the kv attribute.
	GetKv() environment.KeyValue

	// SetKv sets the kv attribute.
	SetKv(environment.KeyValue)

	// Getter signatures
	ID() antlr.TerminalNode
	DOSPTOS() antlr.TerminalNode
	Expr() IExprContext

	// IsKeyvalueContext differentiates from other interfaces.
	IsKeyvalueContext()
}

type KeyvalueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	kv     environment.KeyValue
	_ID    antlr.Token
	_expr  IExprContext
}

func NewEmptyKeyvalueContext() *KeyvalueContext {
	var p = new(KeyvalueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_keyvalue
	return p
}

func InitEmptyKeyvalueContext(p *KeyvalueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_keyvalue
}

func (*KeyvalueContext) IsKeyvalueContext() {}

func NewKeyvalueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeyvalueContext {
	var p = new(KeyvalueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_keyvalue

	return p
}

func (s *KeyvalueContext) GetParser() antlr.Parser { return s.parser }

func (s *KeyvalueContext) Get_ID() antlr.Token { return s._ID }

func (s *KeyvalueContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *KeyvalueContext) Get_expr() IExprContext { return s._expr }

func (s *KeyvalueContext) Set_expr(v IExprContext) { s._expr = v }

func (s *KeyvalueContext) GetKv() environment.KeyValue { return s.kv }

func (s *KeyvalueContext) SetKv(v environment.KeyValue) { s.kv = v }

func (s *KeyvalueContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *KeyvalueContext) DOSPTOS() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserDOSPTOS, 0)
}

func (s *KeyvalueContext) Expr() IExprContext {
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

func (s *KeyvalueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeyvalueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeyvalueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterKeyvalue(s)
	}
}

func (s *KeyvalueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitKeyvalue(s)
	}
}

func (p *SwiftGrammarParser) Keyvalue() (localctx IKeyvalueContext) {
	localctx = NewKeyvalueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 96, SwiftGrammarParserRULE_keyvalue)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(663)

		var _m = p.Match(SwiftGrammarParserID)

		localctx.(*KeyvalueContext)._ID = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(664)
		p.Match(SwiftGrammarParserDOSPTOS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(665)

		var _x = p.expr(0)

		localctx.(*KeyvalueContext)._expr = _x
	}

	localctx.(*KeyvalueContext).kv = environment.KeyValue{
		Key: (func() string {
			if localctx.(*KeyvalueContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*KeyvalueContext).Get_ID().GetText()
			}
		}()),
		Value: localctx.(*KeyvalueContext).Get_expr().GetE(),
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

	// GetOp returns the op token.
	GetOp() antlr.Token

	// Get_OBRA returns the _OBRA token.
	Get_OBRA() antlr.Token

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Get_NUMBER returns the _NUMBER token.
	Get_NUMBER() antlr.Token

	// Get_STRING returns the _STRING token.
	Get_STRING() antlr.Token

	// Get_RTRUE returns the _RTRUE token.
	Get_RTRUE() antlr.Token

	// Get_RFALSE returns the _RFALSE token.
	Get_RFALSE() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// Set_OBRA sets the _OBRA token.
	Set_OBRA(antlr.Token)

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Set_NUMBER sets the _NUMBER token.
	Set_NUMBER(antlr.Token)

	// Set_STRING sets the _STRING token.
	Set_STRING(antlr.Token)

	// Set_RTRUE sets the _RTRUE token.
	Set_RTRUE(antlr.Token)

	// Set_RFALSE sets the _RFALSE token.
	Set_RFALSE(antlr.Token)

	// GetLeft returns the left rule contexts.
	GetLeft() IExprContext

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Get_structexp returns the _structexp rule contexts.
	Get_structexp() IStructexpContext

	// Get_structaccess returns the _structaccess rule contexts.
	Get_structaccess() IStructaccessContext

	// Get_isemptyvec returns the _isemptyvec rule contexts.
	Get_isemptyvec() IIsemptyvecContext

	// Get_countvec returns the _countvec rule contexts.
	Get_countvec() ICountvecContext

	// Get_vectoraccess returns the _vectoraccess rule contexts.
	Get_vectoraccess() IVectoraccessContext

	// Get_arguments returns the _arguments rule contexts.
	Get_arguments() IArgumentsContext

	// GetRight returns the right rule contexts.
	GetRight() IExprContext

	// SetLeft sets the left rule contexts.
	SetLeft(IExprContext)

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// Set_structexp sets the _structexp rule contexts.
	Set_structexp(IStructexpContext)

	// Set_structaccess sets the _structaccess rule contexts.
	Set_structaccess(IStructaccessContext)

	// Set_isemptyvec sets the _isemptyvec rule contexts.
	Set_isemptyvec(IIsemptyvecContext)

	// Set_countvec sets the _countvec rule contexts.
	Set_countvec(ICountvecContext)

	// Set_vectoraccess sets the _vectoraccess rule contexts.
	Set_vectoraccess(IVectoraccessContext)

	// Set_arguments sets the _arguments rule contexts.
	Set_arguments(IArgumentsContext)

	// SetRight sets the right rule contexts.
	SetRight(IExprContext)

	// GetE returns the e attribute.
	GetE() interfaces.Expression

	// SetE sets the e attribute.
	SetE(interfaces.Expression)

	// Getter signatures
	PARIZQ() antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	PARDER() antlr.TerminalNode
	SUB() antlr.TerminalNode
	NOT() antlr.TerminalNode
	Structexp() IStructexpContext
	Structaccess() IStructaccessContext
	Isemptyvec() IIsemptyvecContext
	Countvec() ICountvecContext
	Vectoraccess() IVectoraccessContext
	OBRA() antlr.TerminalNode
	Arguments() IArgumentsContext
	CBRA() antlr.TerminalNode
	ID() antlr.TerminalNode
	NUMBER() antlr.TerminalNode
	STRING() antlr.TerminalNode
	RTRUE() antlr.TerminalNode
	RFALSE() antlr.TerminalNode
	RNIL() antlr.TerminalNode
	MUL() antlr.TerminalNode
	DIV() antlr.TerminalNode
	MODULE() antlr.TerminalNode
	ADD() antlr.TerminalNode
	MAY_IG() antlr.TerminalNode
	MAYOR() antlr.TerminalNode
	MEN_IG() antlr.TerminalNode
	MENOR() antlr.TerminalNode
	IG_IG() antlr.TerminalNode
	DIF() antlr.TerminalNode
	AND() antlr.TerminalNode
	OR() antlr.TerminalNode

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser        antlr.Parser
	e             interfaces.Expression
	left          IExprContext
	_expr         IExprContext
	op            antlr.Token
	_structexp    IStructexpContext
	_structaccess IStructaccessContext
	_isemptyvec   IIsemptyvecContext
	_countvec     ICountvecContext
	_vectoraccess IVectoraccessContext
	_OBRA         antlr.Token
	_arguments    IArgumentsContext
	_ID           antlr.Token
	_NUMBER       antlr.Token
	_STRING       antlr.Token
	_RTRUE        antlr.Token
	_RFALSE       antlr.Token
	right         IExprContext
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) GetOp() antlr.Token { return s.op }

func (s *ExprContext) Get_OBRA() antlr.Token { return s._OBRA }

func (s *ExprContext) Get_ID() antlr.Token { return s._ID }

func (s *ExprContext) Get_NUMBER() antlr.Token { return s._NUMBER }

func (s *ExprContext) Get_STRING() antlr.Token { return s._STRING }

func (s *ExprContext) Get_RTRUE() antlr.Token { return s._RTRUE }

func (s *ExprContext) Get_RFALSE() antlr.Token { return s._RFALSE }

func (s *ExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *ExprContext) Set_OBRA(v antlr.Token) { s._OBRA = v }

func (s *ExprContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *ExprContext) Set_NUMBER(v antlr.Token) { s._NUMBER = v }

func (s *ExprContext) Set_STRING(v antlr.Token) { s._STRING = v }

func (s *ExprContext) Set_RTRUE(v antlr.Token) { s._RTRUE = v }

func (s *ExprContext) Set_RFALSE(v antlr.Token) { s._RFALSE = v }

func (s *ExprContext) GetLeft() IExprContext { return s.left }

func (s *ExprContext) Get_expr() IExprContext { return s._expr }

func (s *ExprContext) Get_structexp() IStructexpContext { return s._structexp }

func (s *ExprContext) Get_structaccess() IStructaccessContext { return s._structaccess }

func (s *ExprContext) Get_isemptyvec() IIsemptyvecContext { return s._isemptyvec }

func (s *ExprContext) Get_countvec() ICountvecContext { return s._countvec }

func (s *ExprContext) Get_vectoraccess() IVectoraccessContext { return s._vectoraccess }

func (s *ExprContext) Get_arguments() IArgumentsContext { return s._arguments }

func (s *ExprContext) GetRight() IExprContext { return s.right }

func (s *ExprContext) SetLeft(v IExprContext) { s.left = v }

func (s *ExprContext) Set_expr(v IExprContext) { s._expr = v }

func (s *ExprContext) Set_structexp(v IStructexpContext) { s._structexp = v }

func (s *ExprContext) Set_structaccess(v IStructaccessContext) { s._structaccess = v }

func (s *ExprContext) Set_isemptyvec(v IIsemptyvecContext) { s._isemptyvec = v }

func (s *ExprContext) Set_countvec(v ICountvecContext) { s._countvec = v }

func (s *ExprContext) Set_vectoraccess(v IVectoraccessContext) { s._vectoraccess = v }

func (s *ExprContext) Set_arguments(v IArgumentsContext) { s._arguments = v }

func (s *ExprContext) SetRight(v IExprContext) { s.right = v }

func (s *ExprContext) GetE() interfaces.Expression { return s.e }

func (s *ExprContext) SetE(v interfaces.Expression) { s.e = v }

func (s *ExprContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARIZQ, 0)
}

func (s *ExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *ExprContext) PARDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARDER, 0)
}

func (s *ExprContext) SUB() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserSUB, 0)
}

func (s *ExprContext) NOT() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserNOT, 0)
}

func (s *ExprContext) Structexp() IStructexpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructexpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructexpContext)
}

func (s *ExprContext) Structaccess() IStructaccessContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructaccessContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructaccessContext)
}

func (s *ExprContext) Isemptyvec() IIsemptyvecContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIsemptyvecContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIsemptyvecContext)
}

func (s *ExprContext) Countvec() ICountvecContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICountvecContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICountvecContext)
}

func (s *ExprContext) Vectoraccess() IVectoraccessContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVectoraccessContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVectoraccessContext)
}

func (s *ExprContext) OBRA() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserOBRA, 0)
}

func (s *ExprContext) Arguments() IArgumentsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentsContext)
}

func (s *ExprContext) CBRA() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCBRA, 0)
}

func (s *ExprContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *ExprContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserNUMBER, 0)
}

func (s *ExprContext) STRING() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserSTRING, 0)
}

func (s *ExprContext) RTRUE() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserRTRUE, 0)
}

func (s *ExprContext) RFALSE() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserRFALSE, 0)
}

func (s *ExprContext) RNIL() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserRNIL, 0)
}

func (s *ExprContext) MUL() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserMUL, 0)
}

func (s *ExprContext) DIV() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserDIV, 0)
}

func (s *ExprContext) MODULE() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserMODULE, 0)
}

func (s *ExprContext) ADD() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserADD, 0)
}

func (s *ExprContext) MAY_IG() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserMAY_IG, 0)
}

func (s *ExprContext) MAYOR() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserMAYOR, 0)
}

func (s *ExprContext) MEN_IG() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserMEN_IG, 0)
}

func (s *ExprContext) MENOR() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserMENOR, 0)
}

func (s *ExprContext) IG_IG() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserIG_IG, 0)
}

func (s *ExprContext) DIF() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserDIF, 0)
}

func (s *ExprContext) AND() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserAND, 0)
}

func (s *ExprContext) OR() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserOR, 0)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitExpr(s)
	}
}

func (p *SwiftGrammarParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *SwiftGrammarParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 98
	p.EnterRecursionRule(localctx, 98, SwiftGrammarParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(710)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 39, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(669)
			p.Match(SwiftGrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(670)

			var _x = p.expr(0)

			localctx.(*ExprContext)._expr = _x
		}
		{
			p.SetState(671)
			p.Match(SwiftGrammarParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ExprContext).e = localctx.(*ExprContext).Get_expr().GetE()

	case 2:
		{
			p.SetState(674)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*ExprContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == SwiftGrammarParserNOT || _la == SwiftGrammarParserSUB) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*ExprContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(675)

			var _x = p.expr(19)

			localctx.(*ExprContext).left = _x
			localctx.(*ExprContext)._expr = _x
		}
		localctx.(*ExprContext).e = expressions.NewUnaryOperation((func() antlr.Token {
			if localctx.(*ExprContext).GetLeft() == nil {
				return nil
			} else {
				return localctx.(*ExprContext).GetLeft().GetStart()
			}
		}()).GetLine(), (func() antlr.Token {
			if localctx.(*ExprContext).GetLeft() == nil {
				return nil
			} else {
				return localctx.(*ExprContext).GetLeft().GetStart()
			}
		}()).GetColumn(), localctx.(*ExprContext).GetLeft().GetE(), (func() string {
			if localctx.(*ExprContext).GetOp() == nil {
				return ""
			} else {
				return localctx.(*ExprContext).GetOp().GetText()
			}
		}()))

	case 3:
		{
			p.SetState(678)

			var _x = p.Structexp()

			localctx.(*ExprContext)._structexp = _x
		}
		localctx.(*ExprContext).e = localctx.(*ExprContext).Get_structexp().GetStructexxp()

	case 4:
		{
			p.SetState(681)

			var _x = p.Structaccess()

			localctx.(*ExprContext)._structaccess = _x
		}
		localctx.(*ExprContext).e = localctx.(*ExprContext).Get_structaccess().GetSaccess()

	case 5:
		{
			p.SetState(684)

			var _x = p.Isemptyvec()

			localctx.(*ExprContext)._isemptyvec = _x
		}
		localctx.(*ExprContext).e = localctx.(*ExprContext).Get_isemptyvec().GetNewisemptyvec()

	case 6:
		{
			p.SetState(687)

			var _x = p.Countvec()

			localctx.(*ExprContext)._countvec = _x
		}
		localctx.(*ExprContext).e = localctx.(*ExprContext).Get_countvec().GetNewcountvec()

	case 7:
		{
			p.SetState(690)

			var _x = p.Vectoraccess()

			localctx.(*ExprContext)._vectoraccess = _x
		}
		localctx.(*ExprContext).e = localctx.(*ExprContext).Get_vectoraccess().GetNewvecaccess()

	case 8:
		{
			p.SetState(693)

			var _m = p.Match(SwiftGrammarParserOBRA)

			localctx.(*ExprContext)._OBRA = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(694)

			var _x = p.Arguments()

			localctx.(*ExprContext)._arguments = _x
		}
		{
			p.SetState(695)
			p.Match(SwiftGrammarParserCBRA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ExprContext).e = expressions.NewVector((func() int {
			if localctx.(*ExprContext).Get_OBRA() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_OBRA().GetLine()
			}
		}()), (func() int {
			if localctx.(*ExprContext).Get_OBRA() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_OBRA().GetColumn()
			}
		}()), localctx.(*ExprContext).Get_arguments().GetArgs())

	case 9:
		{
			p.SetState(698)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*ExprContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ExprContext).e = expressions.NewVariableAccess((func() string {
			if localctx.(*ExprContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*ExprContext).Get_ID().GetText()
			}
		}()))

	case 10:
		{
			p.SetState(700)

			var _m = p.Match(SwiftGrammarParserNUMBER)

			localctx.(*ExprContext)._NUMBER = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		if strings.Contains((func() string {
			if localctx.(*ExprContext).Get_NUMBER() == nil {
				return ""
			} else {
				return localctx.(*ExprContext).Get_NUMBER().GetText()
			}
		}()), ".") {
			num, err := strconv.ParseFloat((func() string {
				if localctx.(*ExprContext).Get_NUMBER() == nil {
					return ""
				} else {
					return localctx.(*ExprContext).Get_NUMBER().GetText()
				}
			}()), 64)
			if err != nil {
				fmt.Println(err)
			}
			localctx.(*ExprContext).e = expressions.NewPrimitive((func() int {
				if localctx.(*ExprContext).Get_NUMBER() == nil {
					return 0
				} else {
					return localctx.(*ExprContext).Get_NUMBER().GetLine()
				}
			}()), (func() int {
				if localctx.(*ExprContext).Get_NUMBER() == nil {
					return 0
				} else {
					return localctx.(*ExprContext).Get_NUMBER().GetColumn()
				}
			}()), num, environment.FLOAT)
		} else {
			num, err := strconv.Atoi((func() string {
				if localctx.(*ExprContext).Get_NUMBER() == nil {
					return ""
				} else {
					return localctx.(*ExprContext).Get_NUMBER().GetText()
				}
			}()))
			if err != nil {
				fmt.Println(err)
			}
			localctx.(*ExprContext).e = expressions.NewPrimitive((func() int {
				if localctx.(*ExprContext).Get_NUMBER() == nil {
					return 0
				} else {
					return localctx.(*ExprContext).Get_NUMBER().GetLine()
				}
			}()), (func() int {
				if localctx.(*ExprContext).Get_NUMBER() == nil {
					return 0
				} else {
					return localctx.(*ExprContext).Get_NUMBER().GetColumn()
				}
			}()), num, environment.INTEGER)
		}

	case 11:
		{
			p.SetState(702)

			var _m = p.Match(SwiftGrammarParserSTRING)

			localctx.(*ExprContext)._STRING = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		str := (func() string {
			if localctx.(*ExprContext).Get_STRING() == nil {
				return ""
			} else {
				return localctx.(*ExprContext).Get_STRING().GetText()
			}
		}())
		var chari = str[1 : len(str)-1]
		chari = strings.ReplaceAll(chari, "\\n", "\n")
		chari = strings.ReplaceAll(chari, "\\r", "\r")
		chari = strings.ReplaceAll(chari, "\\t", "\t")
		chari = strings.ReplaceAll(chari, "\\\\", "\\")
		chari = strings.ReplaceAll(chari, "\\\"", "\"")
		if len(str) == 3 {

			localctx.(*ExprContext).e = expressions.NewPrimitive((func() int {
				if localctx.(*ExprContext).Get_STRING() == nil {
					return 0
				} else {
					return localctx.(*ExprContext).Get_STRING().GetLine()
				}
			}()), (func() int {
				if localctx.(*ExprContext).Get_STRING() == nil {
					return 0
				} else {
					return localctx.(*ExprContext).Get_STRING().GetColumn()
				}
			}()), chari, environment.CHAR)
		} else {

			localctx.(*ExprContext).e = expressions.NewPrimitive((func() int {
				if localctx.(*ExprContext).Get_STRING() == nil {
					return 0
				} else {
					return localctx.(*ExprContext).Get_STRING().GetLine()
				}
			}()), (func() int {
				if localctx.(*ExprContext).Get_STRING() == nil {
					return 0
				} else {
					return localctx.(*ExprContext).Get_STRING().GetColumn()
				}
			}()), chari, environment.STRING)
		}

	case 12:
		{
			p.SetState(704)

			var _m = p.Match(SwiftGrammarParserRTRUE)

			localctx.(*ExprContext)._RTRUE = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ExprContext).e = expressions.NewPrimitive((func() int {
			if localctx.(*ExprContext).Get_RTRUE() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_RTRUE().GetLine()
			}
		}()), (func() int {
			if localctx.(*ExprContext).Get_RTRUE() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_RTRUE().GetColumn()
			}
		}()), true, environment.BOOLEAN)

	case 13:
		{
			p.SetState(706)

			var _m = p.Match(SwiftGrammarParserRFALSE)

			localctx.(*ExprContext)._RFALSE = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ExprContext).e = expressions.NewPrimitive((func() int {
			if localctx.(*ExprContext).Get_RFALSE() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_RFALSE().GetLine()
			}
		}()), (func() int {
			if localctx.(*ExprContext).Get_RFALSE() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_RFALSE().GetColumn()
			}
		}()), false, environment.BOOLEAN)

	case 14:
		{
			p.SetState(708)
			p.Match(SwiftGrammarParserRNIL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ExprContext).e = expressions.NewPrimitive((func() int {
			if localctx.(*ExprContext).Get_RFALSE() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_RFALSE().GetLine()
			}
		}()), (func() int {
			if localctx.(*ExprContext).Get_RFALSE() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_RFALSE().GetColumn()
			}
		}()), nil, environment.NULL)

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(744)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 41, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(742)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 40, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_expr)
				p.SetState(712)

				if !(p.Precpred(p.GetParserRuleContext(), 18)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 18)", ""))
					goto errorExit
				}
				{
					p.SetState(713)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&144220741192122368) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(714)

					var _x = p.expr(19)

					localctx.(*ExprContext).right = _x
					localctx.(*ExprContext)._expr = _x
				}
				localctx.(*ExprContext).e = expressions.NewArithmeticOperation((func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetLine(), (func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetColumn(), localctx.(*ExprContext).GetLeft().GetE(), (func() string {
					if localctx.(*ExprContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExprContext).GetOp().GetText()
					}
				}()), localctx.(*ExprContext).GetRight().GetE())

			case 2:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_expr)
				p.SetState(717)

				if !(p.Precpred(p.GetParserRuleContext(), 17)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 17)", ""))
					goto errorExit
				}
				{
					p.SetState(718)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == SwiftGrammarParserADD || _la == SwiftGrammarParserSUB) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(719)

					var _x = p.expr(18)

					localctx.(*ExprContext).right = _x
					localctx.(*ExprContext)._expr = _x
				}
				localctx.(*ExprContext).e = expressions.NewArithmeticOperation((func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetLine(), (func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetColumn(), localctx.(*ExprContext).GetLeft().GetE(), (func() string {
					if localctx.(*ExprContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExprContext).GetOp().GetText()
					}
				}()), localctx.(*ExprContext).GetRight().GetE())

			case 3:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_expr)
				p.SetState(722)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
					goto errorExit
				}
				{
					p.SetState(723)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&32985348833280) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(724)

					var _x = p.expr(17)

					localctx.(*ExprContext).right = _x
					localctx.(*ExprContext)._expr = _x
				}
				localctx.(*ExprContext).e = expressions.NewRelationalOperation((func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetLine(), (func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetColumn(), localctx.(*ExprContext).GetLeft().GetE(), (func() string {
					if localctx.(*ExprContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExprContext).GetOp().GetText()
					}
				}()), localctx.(*ExprContext).GetRight().GetE())

			case 4:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_expr)
				p.SetState(727)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
					goto errorExit
				}
				{
					p.SetState(728)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == SwiftGrammarParserDIF || _la == SwiftGrammarParserIG_IG) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(729)

					var _x = p.expr(16)

					localctx.(*ExprContext).right = _x
					localctx.(*ExprContext)._expr = _x
				}
				localctx.(*ExprContext).e = expressions.NewRelationalOperation((func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetLine(), (func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetColumn(), localctx.(*ExprContext).GetLeft().GetE(), (func() string {
					if localctx.(*ExprContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExprContext).GetOp().GetText()
					}
				}()), localctx.(*ExprContext).GetRight().GetE())

			case 5:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_expr)
				p.SetState(732)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
					goto errorExit
				}
				{
					p.SetState(733)

					var _m = p.Match(SwiftGrammarParserAND)

					localctx.(*ExprContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(734)

					var _x = p.expr(15)

					localctx.(*ExprContext).right = _x
					localctx.(*ExprContext)._expr = _x
				}
				localctx.(*ExprContext).e = expressions.NewBooleanOperation((func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetLine(), (func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetColumn(), localctx.(*ExprContext).GetLeft().GetE(), (func() string {
					if localctx.(*ExprContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExprContext).GetOp().GetText()
					}
				}()), localctx.(*ExprContext).GetRight().GetE())

			case 6:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_expr)
				p.SetState(737)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
					goto errorExit
				}
				{
					p.SetState(738)

					var _m = p.Match(SwiftGrammarParserOR)

					localctx.(*ExprContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(739)

					var _x = p.expr(14)

					localctx.(*ExprContext).right = _x
					localctx.(*ExprContext)._expr = _x
				}
				localctx.(*ExprContext).e = expressions.NewBooleanOperation((func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetLine(), (func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetColumn(), localctx.(*ExprContext).GetLeft().GetE(), (func() string {
					if localctx.(*ExprContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExprContext).GetOp().GetText()
					}
				}()), localctx.(*ExprContext).GetRight().GetE())

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(746)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 41, p.GetParserRuleContext())
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

func (p *SwiftGrammarParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 41:
		var t *Values2Context = nil
		if localctx != nil {
			t = localctx.(*Values2Context)
		}
		return p.Values2_Sempred(t, predIndex)

	case 49:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *SwiftGrammarParser) Values2_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SwiftGrammarParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 18)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 17)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 16)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 15)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 13)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
