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
		"'for'", "'in'", "", "", "", "'+='", "'-='", "'!='", "'=='", "'!'",
		"'||'", "'&&'", "'='", "'>='", "'<='", "'>'", "'<'", "'*'", "'/'", "'+'",
		"'-'", "'('", "')'", "'{'", "'}'", "'['", "']'", "';'", "':'", "'%'",
		"','", "'?'", "'.'",
	}
	staticData.SymbolicNames = []string{
		"", "RINT", "RFLOAT", "RBOOL", "RSTRING", "RCHARACTER", "RTRUE", "RFALSE",
		"RPRINT", "RIF", "RELSE", "RWHILE", "RVAR", "RLET", "RNIL", "RBREAK",
		"RCONTINUE", "RAPPEND", "RREMOVELAST", "RRAT", "RREMOVEAT", "RISEMPTY",
		"RCOUNT", "RSWITCH", "RCASE", "RDEFAULT", "RFOR", "RIN", "NUMBER", "STRING",
		"ID", "UNARYPLUS", "UNARYMINUS", "DIF", "IG_IG", "NOT", "OR", "AND",
		"IG", "MAY_IG", "MEN_IG", "MAYOR", "MENOR", "MUL", "DIV", "ADD", "SUB",
		"PARIZQ", "PARDER", "LLAVEIZQ", "LLAVEDER", "OBRA", "CBRA", "PTOCOMA",
		"DOSPTOS", "MODULE", "COMA", "QM", "PTO", "WHITESPACE", "COMMENT", "LINE_COMMENT",
	}
	staticData.RuleNames = []string{
		"s", "block", "arguments", "argument", "instruction", "forloop", "range",
		"removeatvec", "appendvec", "removelastvec", "vecdec", "breakstatement",
		"continuestatement", "switchstatement", "caselist", "case", "defaultstatement",
		"ifstmt", "eliflist", "elif", "elsestament", "printstmt", "while_statement",
		"vardec", "constdec", "asignation", "unarysum", "unarysub", "isemptyvec",
		"countvec", "vectoraccess", "expr",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 61, 497, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 5, 1, 70, 8, 1, 10, 1, 12, 1, 73, 9,
		1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3,
		2, 86, 8, 2, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 3, 4, 93, 8, 4, 1, 4, 1, 4,
		1, 4, 1, 4, 3, 4, 99, 8, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 105, 8, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 3, 4, 111, 8, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 117,
		8, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 123, 8, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		3, 4, 129, 8, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 135, 8, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 3, 4, 141, 8, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 147, 8, 4, 1, 4,
		1, 4, 1, 4, 1, 4, 3, 4, 153, 8, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 159, 8,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 3, 4, 175, 8, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3,
		5, 195, 8, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1,
		7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1,
		8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1,
		10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1,
		10, 3, 10, 252, 8, 10, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 13,
		1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1,
		13, 1, 13, 1, 13, 1, 13, 3, 13, 275, 8, 13, 1, 14, 1, 14, 1, 14, 1, 14,
		1, 14, 1, 14, 1, 14, 1, 14, 3, 14, 285, 8, 14, 1, 15, 1, 15, 1, 15, 1,
		15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17,
		1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1,
		17, 1, 17, 1, 17, 1, 17, 1, 17, 3, 17, 316, 8, 17, 1, 18, 1, 18, 1, 18,
		1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 3, 18, 326, 8, 18, 1, 19, 1, 19, 1,
		19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20,
		1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1,
		22, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23,
		1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1,
		23, 1, 23, 1, 23, 3, 23, 375, 8, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24,
		1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 3, 24, 391,
		8, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1,
		26, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28,
		1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1,
		30, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31,
		1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1,
		31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31,
		1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 3, 31, 460, 8, 31, 1, 31, 1, 31, 1,
		31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31,
		1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1,
		31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 5, 31, 492, 8, 31, 10, 31,
		12, 31, 495, 9, 31, 1, 31, 0, 1, 62, 32, 0, 2, 4, 6, 8, 10, 12, 14, 16,
		18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52,
		54, 56, 58, 60, 62, 0, 6, 1, 0, 1, 5, 2, 0, 35, 35, 46, 46, 2, 0, 43, 44,
		55, 55, 1, 0, 45, 46, 1, 0, 39, 42, 1, 0, 33, 34, 522, 0, 64, 1, 0, 0,
		0, 2, 71, 1, 0, 0, 0, 4, 85, 1, 0, 0, 0, 6, 87, 1, 0, 0, 0, 8, 174, 1,
		0, 0, 0, 10, 194, 1, 0, 0, 0, 12, 196, 1, 0, 0, 0, 14, 203, 1, 0, 0, 0,
		16, 213, 1, 0, 0, 0, 18, 221, 1, 0, 0, 0, 20, 251, 1, 0, 0, 0, 22, 253,
		1, 0, 0, 0, 24, 256, 1, 0, 0, 0, 26, 274, 1, 0, 0, 0, 28, 284, 1, 0, 0,
		0, 30, 286, 1, 0, 0, 0, 32, 292, 1, 0, 0, 0, 34, 315, 1, 0, 0, 0, 36, 325,
		1, 0, 0, 0, 38, 327, 1, 0, 0, 0, 40, 335, 1, 0, 0, 0, 42, 341, 1, 0, 0,
		0, 44, 347, 1, 0, 0, 0, 46, 374, 1, 0, 0, 0, 48, 390, 1, 0, 0, 0, 50, 392,
		1, 0, 0, 0, 52, 397, 1, 0, 0, 0, 54, 402, 1, 0, 0, 0, 56, 407, 1, 0, 0,
		0, 58, 412, 1, 0, 0, 0, 60, 417, 1, 0, 0, 0, 62, 459, 1, 0, 0, 0, 64, 65,
		3, 2, 1, 0, 65, 66, 5, 0, 0, 1, 66, 67, 6, 0, -1, 0, 67, 1, 1, 0, 0, 0,
		68, 70, 3, 8, 4, 0, 69, 68, 1, 0, 0, 0, 70, 73, 1, 0, 0, 0, 71, 69, 1,
		0, 0, 0, 71, 72, 1, 0, 0, 0, 72, 74, 1, 0, 0, 0, 73, 71, 1, 0, 0, 0, 74,
		75, 6, 1, -1, 0, 75, 3, 1, 0, 0, 0, 76, 77, 3, 6, 3, 0, 77, 78, 5, 56,
		0, 0, 78, 79, 3, 4, 2, 0, 79, 80, 6, 2, -1, 0, 80, 86, 1, 0, 0, 0, 81,
		82, 3, 6, 3, 0, 82, 83, 6, 2, -1, 0, 83, 86, 1, 0, 0, 0, 84, 86, 6, 2,
		-1, 0, 85, 76, 1, 0, 0, 0, 85, 81, 1, 0, 0, 0, 85, 84, 1, 0, 0, 0, 86,
		5, 1, 0, 0, 0, 87, 88, 3, 62, 31, 0, 88, 89, 6, 3, -1, 0, 89, 7, 1, 0,
		0, 0, 90, 92, 3, 42, 21, 0, 91, 93, 5, 53, 0, 0, 92, 91, 1, 0, 0, 0, 92,
		93, 1, 0, 0, 0, 93, 94, 1, 0, 0, 0, 94, 95, 6, 4, -1, 0, 95, 175, 1, 0,
		0, 0, 96, 98, 3, 46, 23, 0, 97, 99, 5, 53, 0, 0, 98, 97, 1, 0, 0, 0, 98,
		99, 1, 0, 0, 0, 99, 100, 1, 0, 0, 0, 100, 101, 6, 4, -1, 0, 101, 175, 1,
		0, 0, 0, 102, 104, 3, 48, 24, 0, 103, 105, 5, 53, 0, 0, 104, 103, 1, 0,
		0, 0, 104, 105, 1, 0, 0, 0, 105, 106, 1, 0, 0, 0, 106, 107, 6, 4, -1, 0,
		107, 175, 1, 0, 0, 0, 108, 110, 3, 20, 10, 0, 109, 111, 5, 53, 0, 0, 110,
		109, 1, 0, 0, 0, 110, 111, 1, 0, 0, 0, 111, 112, 1, 0, 0, 0, 112, 113,
		6, 4, -1, 0, 113, 175, 1, 0, 0, 0, 114, 116, 3, 16, 8, 0, 115, 117, 5,
		53, 0, 0, 116, 115, 1, 0, 0, 0, 116, 117, 1, 0, 0, 0, 117, 118, 1, 0, 0,
		0, 118, 119, 6, 4, -1, 0, 119, 175, 1, 0, 0, 0, 120, 122, 3, 18, 9, 0,
		121, 123, 5, 53, 0, 0, 122, 121, 1, 0, 0, 0, 122, 123, 1, 0, 0, 0, 123,
		124, 1, 0, 0, 0, 124, 125, 6, 4, -1, 0, 125, 175, 1, 0, 0, 0, 126, 128,
		3, 14, 7, 0, 127, 129, 5, 53, 0, 0, 128, 127, 1, 0, 0, 0, 128, 129, 1,
		0, 0, 0, 129, 130, 1, 0, 0, 0, 130, 131, 6, 4, -1, 0, 131, 175, 1, 0, 0,
		0, 132, 134, 3, 50, 25, 0, 133, 135, 5, 53, 0, 0, 134, 133, 1, 0, 0, 0,
		134, 135, 1, 0, 0, 0, 135, 136, 1, 0, 0, 0, 136, 137, 6, 4, -1, 0, 137,
		175, 1, 0, 0, 0, 138, 140, 3, 52, 26, 0, 139, 141, 5, 53, 0, 0, 140, 139,
		1, 0, 0, 0, 140, 141, 1, 0, 0, 0, 141, 142, 1, 0, 0, 0, 142, 143, 6, 4,
		-1, 0, 143, 175, 1, 0, 0, 0, 144, 146, 3, 54, 27, 0, 145, 147, 5, 53, 0,
		0, 146, 145, 1, 0, 0, 0, 146, 147, 1, 0, 0, 0, 147, 148, 1, 0, 0, 0, 148,
		149, 6, 4, -1, 0, 149, 175, 1, 0, 0, 0, 150, 152, 3, 22, 11, 0, 151, 153,
		5, 53, 0, 0, 152, 151, 1, 0, 0, 0, 152, 153, 1, 0, 0, 0, 153, 154, 1, 0,
		0, 0, 154, 155, 6, 4, -1, 0, 155, 175, 1, 0, 0, 0, 156, 158, 3, 24, 12,
		0, 157, 159, 5, 53, 0, 0, 158, 157, 1, 0, 0, 0, 158, 159, 1, 0, 0, 0, 159,
		160, 1, 0, 0, 0, 160, 161, 6, 4, -1, 0, 161, 175, 1, 0, 0, 0, 162, 163,
		3, 34, 17, 0, 163, 164, 6, 4, -1, 0, 164, 175, 1, 0, 0, 0, 165, 166, 3,
		44, 22, 0, 166, 167, 6, 4, -1, 0, 167, 175, 1, 0, 0, 0, 168, 169, 3, 26,
		13, 0, 169, 170, 6, 4, -1, 0, 170, 175, 1, 0, 0, 0, 171, 172, 3, 10, 5,
		0, 172, 173, 6, 4, -1, 0, 173, 175, 1, 0, 0, 0, 174, 90, 1, 0, 0, 0, 174,
		96, 1, 0, 0, 0, 174, 102, 1, 0, 0, 0, 174, 108, 1, 0, 0, 0, 174, 114, 1,
		0, 0, 0, 174, 120, 1, 0, 0, 0, 174, 126, 1, 0, 0, 0, 174, 132, 1, 0, 0,
		0, 174, 138, 1, 0, 0, 0, 174, 144, 1, 0, 0, 0, 174, 150, 1, 0, 0, 0, 174,
		156, 1, 0, 0, 0, 174, 162, 1, 0, 0, 0, 174, 165, 1, 0, 0, 0, 174, 168,
		1, 0, 0, 0, 174, 171, 1, 0, 0, 0, 175, 9, 1, 0, 0, 0, 176, 177, 5, 26,
		0, 0, 177, 178, 5, 30, 0, 0, 178, 179, 5, 27, 0, 0, 179, 180, 3, 62, 31,
		0, 180, 181, 5, 49, 0, 0, 181, 182, 3, 2, 1, 0, 182, 183, 5, 50, 0, 0,
		183, 184, 6, 5, -1, 0, 184, 195, 1, 0, 0, 0, 185, 186, 5, 26, 0, 0, 186,
		187, 5, 30, 0, 0, 187, 188, 5, 27, 0, 0, 188, 189, 3, 12, 6, 0, 189, 190,
		5, 49, 0, 0, 190, 191, 3, 2, 1, 0, 191, 192, 5, 50, 0, 0, 192, 193, 6,
		5, -1, 0, 193, 195, 1, 0, 0, 0, 194, 176, 1, 0, 0, 0, 194, 185, 1, 0, 0,
		0, 195, 11, 1, 0, 0, 0, 196, 197, 3, 62, 31, 0, 197, 198, 5, 58, 0, 0,
		198, 199, 5, 58, 0, 0, 199, 200, 5, 58, 0, 0, 200, 201, 3, 62, 31, 0, 201,
		202, 6, 6, -1, 0, 202, 13, 1, 0, 0, 0, 203, 204, 5, 30, 0, 0, 204, 205,
		5, 58, 0, 0, 205, 206, 5, 20, 0, 0, 206, 207, 5, 47, 0, 0, 207, 208, 5,
		19, 0, 0, 208, 209, 5, 54, 0, 0, 209, 210, 3, 62, 31, 0, 210, 211, 5, 48,
		0, 0, 211, 212, 6, 7, -1, 0, 212, 15, 1, 0, 0, 0, 213, 214, 5, 30, 0, 0,
		214, 215, 5, 58, 0, 0, 215, 216, 5, 17, 0, 0, 216, 217, 5, 47, 0, 0, 217,
		218, 3, 62, 31, 0, 218, 219, 5, 48, 0, 0, 219, 220, 6, 8, -1, 0, 220, 17,
		1, 0, 0, 0, 221, 222, 5, 30, 0, 0, 222, 223, 5, 58, 0, 0, 223, 224, 5,
		18, 0, 0, 224, 225, 5, 47, 0, 0, 225, 226, 5, 48, 0, 0, 226, 227, 6, 9,
		-1, 0, 227, 19, 1, 0, 0, 0, 228, 229, 5, 12, 0, 0, 229, 230, 5, 30, 0,
		0, 230, 231, 5, 54, 0, 0, 231, 232, 5, 51, 0, 0, 232, 233, 7, 0, 0, 0,
		233, 234, 5, 52, 0, 0, 234, 235, 5, 38, 0, 0, 235, 236, 5, 51, 0, 0, 236,
		237, 7, 0, 0, 0, 237, 238, 5, 52, 0, 0, 238, 239, 5, 51, 0, 0, 239, 240,
		5, 52, 0, 0, 240, 252, 6, 10, -1, 0, 241, 242, 5, 12, 0, 0, 242, 243, 5,
		30, 0, 0, 243, 244, 5, 54, 0, 0, 244, 245, 5, 51, 0, 0, 245, 246, 7, 0,
		0, 0, 246, 247, 5, 52, 0, 0, 247, 248, 5, 38, 0, 0, 248, 249, 3, 62, 31,
		0, 249, 250, 6, 10, -1, 0, 250, 252, 1, 0, 0, 0, 251, 228, 1, 0, 0, 0,
		251, 241, 1, 0, 0, 0, 252, 21, 1, 0, 0, 0, 253, 254, 5, 15, 0, 0, 254,
		255, 6, 11, -1, 0, 255, 23, 1, 0, 0, 0, 256, 257, 5, 16, 0, 0, 257, 258,
		6, 12, -1, 0, 258, 25, 1, 0, 0, 0, 259, 260, 5, 23, 0, 0, 260, 261, 3,
		62, 31, 0, 261, 262, 5, 49, 0, 0, 262, 263, 3, 28, 14, 0, 263, 264, 5,
		50, 0, 0, 264, 265, 6, 13, -1, 0, 265, 275, 1, 0, 0, 0, 266, 267, 5, 23,
		0, 0, 267, 268, 3, 62, 31, 0, 268, 269, 5, 49, 0, 0, 269, 270, 3, 28, 14,
		0, 270, 271, 3, 32, 16, 0, 271, 272, 5, 50, 0, 0, 272, 273, 6, 13, -1,
		0, 273, 275, 1, 0, 0, 0, 274, 259, 1, 0, 0, 0, 274, 266, 1, 0, 0, 0, 275,
		27, 1, 0, 0, 0, 276, 277, 3, 30, 15, 0, 277, 278, 3, 28, 14, 0, 278, 279,
		6, 14, -1, 0, 279, 285, 1, 0, 0, 0, 280, 281, 3, 30, 15, 0, 281, 282, 6,
		14, -1, 0, 282, 285, 1, 0, 0, 0, 283, 285, 6, 14, -1, 0, 284, 276, 1, 0,
		0, 0, 284, 280, 1, 0, 0, 0, 284, 283, 1, 0, 0, 0, 285, 29, 1, 0, 0, 0,
		286, 287, 5, 24, 0, 0, 287, 288, 3, 62, 31, 0, 288, 289, 5, 54, 0, 0, 289,
		290, 3, 2, 1, 0, 290, 291, 6, 15, -1, 0, 291, 31, 1, 0, 0, 0, 292, 293,
		5, 25, 0, 0, 293, 294, 5, 54, 0, 0, 294, 295, 3, 2, 1, 0, 295, 296, 5,
		50, 0, 0, 296, 297, 6, 16, -1, 0, 297, 33, 1, 0, 0, 0, 298, 299, 5, 9,
		0, 0, 299, 300, 3, 62, 31, 0, 300, 301, 5, 49, 0, 0, 301, 302, 3, 2, 1,
		0, 302, 303, 5, 50, 0, 0, 303, 304, 3, 36, 18, 0, 304, 305, 6, 17, -1,
		0, 305, 316, 1, 0, 0, 0, 306, 307, 5, 9, 0, 0, 307, 308, 3, 62, 31, 0,
		308, 309, 5, 49, 0, 0, 309, 310, 3, 2, 1, 0, 310, 311, 5, 50, 0, 0, 311,
		312, 3, 36, 18, 0, 312, 313, 3, 40, 20, 0, 313, 314, 6, 17, -1, 0, 314,
		316, 1, 0, 0, 0, 315, 298, 1, 0, 0, 0, 315, 306, 1, 0, 0, 0, 316, 35, 1,
		0, 0, 0, 317, 318, 3, 38, 19, 0, 318, 319, 3, 36, 18, 0, 319, 320, 6, 18,
		-1, 0, 320, 326, 1, 0, 0, 0, 321, 322, 3, 38, 19, 0, 322, 323, 6, 18, -1,
		0, 323, 326, 1, 0, 0, 0, 324, 326, 6, 18, -1, 0, 325, 317, 1, 0, 0, 0,
		325, 321, 1, 0, 0, 0, 325, 324, 1, 0, 0, 0, 326, 37, 1, 0, 0, 0, 327, 328,
		5, 10, 0, 0, 328, 329, 5, 9, 0, 0, 329, 330, 3, 62, 31, 0, 330, 331, 5,
		49, 0, 0, 331, 332, 3, 2, 1, 0, 332, 333, 5, 50, 0, 0, 333, 334, 6, 19,
		-1, 0, 334, 39, 1, 0, 0, 0, 335, 336, 5, 10, 0, 0, 336, 337, 5, 49, 0,
		0, 337, 338, 3, 2, 1, 0, 338, 339, 5, 50, 0, 0, 339, 340, 6, 20, -1, 0,
		340, 41, 1, 0, 0, 0, 341, 342, 5, 8, 0, 0, 342, 343, 5, 47, 0, 0, 343,
		344, 3, 4, 2, 0, 344, 345, 5, 48, 0, 0, 345, 346, 6, 21, -1, 0, 346, 43,
		1, 0, 0, 0, 347, 348, 5, 11, 0, 0, 348, 349, 3, 62, 31, 0, 349, 350, 5,
		49, 0, 0, 350, 351, 3, 2, 1, 0, 351, 352, 5, 50, 0, 0, 352, 353, 6, 22,
		-1, 0, 353, 45, 1, 0, 0, 0, 354, 355, 5, 12, 0, 0, 355, 356, 5, 30, 0,
		0, 356, 357, 5, 54, 0, 0, 357, 358, 7, 0, 0, 0, 358, 359, 5, 38, 0, 0,
		359, 360, 3, 62, 31, 0, 360, 361, 6, 23, -1, 0, 361, 375, 1, 0, 0, 0, 362,
		363, 5, 12, 0, 0, 363, 364, 5, 30, 0, 0, 364, 365, 5, 38, 0, 0, 365, 366,
		3, 62, 31, 0, 366, 367, 6, 23, -1, 0, 367, 375, 1, 0, 0, 0, 368, 369, 5,
		12, 0, 0, 369, 370, 5, 30, 0, 0, 370, 371, 5, 54, 0, 0, 371, 372, 7, 0,
		0, 0, 372, 373, 5, 57, 0, 0, 373, 375, 6, 23, -1, 0, 374, 354, 1, 0, 0,
		0, 374, 362, 1, 0, 0, 0, 374, 368, 1, 0, 0, 0, 375, 47, 1, 0, 0, 0, 376,
		377, 5, 13, 0, 0, 377, 378, 5, 30, 0, 0, 378, 379, 5, 54, 0, 0, 379, 380,
		7, 0, 0, 0, 380, 381, 5, 38, 0, 0, 381, 382, 3, 62, 31, 0, 382, 383, 6,
		24, -1, 0, 383, 391, 1, 0, 0, 0, 384, 385, 5, 13, 0, 0, 385, 386, 5, 30,
		0, 0, 386, 387, 5, 38, 0, 0, 387, 388, 3, 62, 31, 0, 388, 389, 6, 24, -1,
		0, 389, 391, 1, 0, 0, 0, 390, 376, 1, 0, 0, 0, 390, 384, 1, 0, 0, 0, 391,
		49, 1, 0, 0, 0, 392, 393, 5, 30, 0, 0, 393, 394, 5, 38, 0, 0, 394, 395,
		3, 62, 31, 0, 395, 396, 6, 25, -1, 0, 396, 51, 1, 0, 0, 0, 397, 398, 5,
		30, 0, 0, 398, 399, 5, 31, 0, 0, 399, 400, 3, 62, 31, 0, 400, 401, 6, 26,
		-1, 0, 401, 53, 1, 0, 0, 0, 402, 403, 5, 30, 0, 0, 403, 404, 5, 32, 0,
		0, 404, 405, 3, 62, 31, 0, 405, 406, 6, 27, -1, 0, 406, 55, 1, 0, 0, 0,
		407, 408, 5, 30, 0, 0, 408, 409, 5, 58, 0, 0, 409, 410, 5, 21, 0, 0, 410,
		411, 6, 28, -1, 0, 411, 57, 1, 0, 0, 0, 412, 413, 5, 30, 0, 0, 413, 414,
		5, 58, 0, 0, 414, 415, 5, 22, 0, 0, 415, 416, 6, 29, -1, 0, 416, 59, 1,
		0, 0, 0, 417, 418, 5, 30, 0, 0, 418, 419, 5, 51, 0, 0, 419, 420, 3, 62,
		31, 0, 420, 421, 5, 52, 0, 0, 421, 422, 6, 30, -1, 0, 422, 61, 1, 0, 0,
		0, 423, 424, 6, 31, -1, 0, 424, 425, 5, 47, 0, 0, 425, 426, 3, 62, 31,
		0, 426, 427, 5, 48, 0, 0, 427, 428, 6, 31, -1, 0, 428, 460, 1, 0, 0, 0,
		429, 430, 7, 1, 0, 0, 430, 431, 3, 62, 31, 17, 431, 432, 6, 31, -1, 0,
		432, 460, 1, 0, 0, 0, 433, 434, 3, 56, 28, 0, 434, 435, 6, 31, -1, 0, 435,
		460, 1, 0, 0, 0, 436, 437, 3, 58, 29, 0, 437, 438, 6, 31, -1, 0, 438, 460,
		1, 0, 0, 0, 439, 440, 3, 60, 30, 0, 440, 441, 6, 31, -1, 0, 441, 460, 1,
		0, 0, 0, 442, 443, 5, 51, 0, 0, 443, 444, 3, 4, 2, 0, 444, 445, 5, 52,
		0, 0, 445, 446, 6, 31, -1, 0, 446, 460, 1, 0, 0, 0, 447, 448, 5, 30, 0,
		0, 448, 460, 6, 31, -1, 0, 449, 450, 5, 28, 0, 0, 450, 460, 6, 31, -1,
		0, 451, 452, 5, 29, 0, 0, 452, 460, 6, 31, -1, 0, 453, 454, 5, 6, 0, 0,
		454, 460, 6, 31, -1, 0, 455, 456, 5, 7, 0, 0, 456, 460, 6, 31, -1, 0, 457,
		458, 5, 14, 0, 0, 458, 460, 6, 31, -1, 0, 459, 423, 1, 0, 0, 0, 459, 429,
		1, 0, 0, 0, 459, 433, 1, 0, 0, 0, 459, 436, 1, 0, 0, 0, 459, 439, 1, 0,
		0, 0, 459, 442, 1, 0, 0, 0, 459, 447, 1, 0, 0, 0, 459, 449, 1, 0, 0, 0,
		459, 451, 1, 0, 0, 0, 459, 453, 1, 0, 0, 0, 459, 455, 1, 0, 0, 0, 459,
		457, 1, 0, 0, 0, 460, 493, 1, 0, 0, 0, 461, 462, 10, 16, 0, 0, 462, 463,
		7, 2, 0, 0, 463, 464, 3, 62, 31, 17, 464, 465, 6, 31, -1, 0, 465, 492,
		1, 0, 0, 0, 466, 467, 10, 15, 0, 0, 467, 468, 7, 3, 0, 0, 468, 469, 3,
		62, 31, 16, 469, 470, 6, 31, -1, 0, 470, 492, 1, 0, 0, 0, 471, 472, 10,
		14, 0, 0, 472, 473, 7, 4, 0, 0, 473, 474, 3, 62, 31, 15, 474, 475, 6, 31,
		-1, 0, 475, 492, 1, 0, 0, 0, 476, 477, 10, 13, 0, 0, 477, 478, 7, 5, 0,
		0, 478, 479, 3, 62, 31, 14, 479, 480, 6, 31, -1, 0, 480, 492, 1, 0, 0,
		0, 481, 482, 10, 12, 0, 0, 482, 483, 5, 37, 0, 0, 483, 484, 3, 62, 31,
		13, 484, 485, 6, 31, -1, 0, 485, 492, 1, 0, 0, 0, 486, 487, 10, 11, 0,
		0, 487, 488, 5, 36, 0, 0, 488, 489, 3, 62, 31, 12, 489, 490, 6, 31, -1,
		0, 490, 492, 1, 0, 0, 0, 491, 461, 1, 0, 0, 0, 491, 466, 1, 0, 0, 0, 491,
		471, 1, 0, 0, 0, 491, 476, 1, 0, 0, 0, 491, 481, 1, 0, 0, 0, 491, 486,
		1, 0, 0, 0, 492, 495, 1, 0, 0, 0, 493, 491, 1, 0, 0, 0, 493, 494, 1, 0,
		0, 0, 494, 63, 1, 0, 0, 0, 495, 493, 1, 0, 0, 0, 26, 71, 85, 92, 98, 104,
		110, 116, 122, 128, 134, 140, 146, 152, 158, 174, 194, 251, 274, 284, 315,
		325, 374, 390, 459, 491, 493,
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
	SwiftGrammarParserNUMBER       = 28
	SwiftGrammarParserSTRING       = 29
	SwiftGrammarParserID           = 30
	SwiftGrammarParserUNARYPLUS    = 31
	SwiftGrammarParserUNARYMINUS   = 32
	SwiftGrammarParserDIF          = 33
	SwiftGrammarParserIG_IG        = 34
	SwiftGrammarParserNOT          = 35
	SwiftGrammarParserOR           = 36
	SwiftGrammarParserAND          = 37
	SwiftGrammarParserIG           = 38
	SwiftGrammarParserMAY_IG       = 39
	SwiftGrammarParserMEN_IG       = 40
	SwiftGrammarParserMAYOR        = 41
	SwiftGrammarParserMENOR        = 42
	SwiftGrammarParserMUL          = 43
	SwiftGrammarParserDIV          = 44
	SwiftGrammarParserADD          = 45
	SwiftGrammarParserSUB          = 46
	SwiftGrammarParserPARIZQ       = 47
	SwiftGrammarParserPARDER       = 48
	SwiftGrammarParserLLAVEIZQ     = 49
	SwiftGrammarParserLLAVEDER     = 50
	SwiftGrammarParserOBRA         = 51
	SwiftGrammarParserCBRA         = 52
	SwiftGrammarParserPTOCOMA      = 53
	SwiftGrammarParserDOSPTOS      = 54
	SwiftGrammarParserMODULE       = 55
	SwiftGrammarParserCOMA         = 56
	SwiftGrammarParserQM           = 57
	SwiftGrammarParserPTO          = 58
	SwiftGrammarParserWHITESPACE   = 59
	SwiftGrammarParserCOMMENT      = 60
	SwiftGrammarParserLINE_COMMENT = 61
)

// SwiftGrammarParser rules.
const (
	SwiftGrammarParserRULE_s                 = 0
	SwiftGrammarParserRULE_block             = 1
	SwiftGrammarParserRULE_arguments         = 2
	SwiftGrammarParserRULE_argument          = 3
	SwiftGrammarParserRULE_instruction       = 4
	SwiftGrammarParserRULE_forloop           = 5
	SwiftGrammarParserRULE_range             = 6
	SwiftGrammarParserRULE_removeatvec       = 7
	SwiftGrammarParserRULE_appendvec         = 8
	SwiftGrammarParserRULE_removelastvec     = 9
	SwiftGrammarParserRULE_vecdec            = 10
	SwiftGrammarParserRULE_breakstatement    = 11
	SwiftGrammarParserRULE_continuestatement = 12
	SwiftGrammarParserRULE_switchstatement   = 13
	SwiftGrammarParserRULE_caselist          = 14
	SwiftGrammarParserRULE_case              = 15
	SwiftGrammarParserRULE_defaultstatement  = 16
	SwiftGrammarParserRULE_ifstmt            = 17
	SwiftGrammarParserRULE_eliflist          = 18
	SwiftGrammarParserRULE_elif              = 19
	SwiftGrammarParserRULE_elsestament       = 20
	SwiftGrammarParserRULE_printstmt         = 21
	SwiftGrammarParserRULE_while_statement   = 22
	SwiftGrammarParserRULE_vardec            = 23
	SwiftGrammarParserRULE_constdec          = 24
	SwiftGrammarParserRULE_asignation        = 25
	SwiftGrammarParserRULE_unarysum          = 26
	SwiftGrammarParserRULE_unarysub          = 27
	SwiftGrammarParserRULE_isemptyvec        = 28
	SwiftGrammarParserRULE_countvec          = 29
	SwiftGrammarParserRULE_vectoraccess      = 30
	SwiftGrammarParserRULE_expr              = 31
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
		p.SetState(64)

		var _x = p.Block()

		localctx.(*SContext)._block = _x
	}
	{
		p.SetState(65)
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
	p.SetState(71)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1149352704) != 0 {
		{
			p.SetState(68)

			var _x = p.Instruction()

			localctx.(*BlockContext)._instruction = _x
		}
		localctx.(*BlockContext).ins = append(localctx.(*BlockContext).ins, localctx.(*BlockContext)._instruction)

		p.SetState(73)
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

	p.SetState(85)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(76)

			var _x = p.Argument()

			localctx.(*ArgumentsContext)._argument = _x
		}
		{
			p.SetState(77)
			p.Match(SwiftGrammarParserCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(78)

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
			p.SetState(81)

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
		p.SetState(87)

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

	// Get_vardec returns the _vardec rule contexts.
	Get_vardec() IVardecContext

	// Get_constdec returns the _constdec rule contexts.
	Get_constdec() IConstdecContext

	// Get_vecdec returns the _vecdec rule contexts.
	Get_vecdec() IVecdecContext

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

	// Get_ifstmt returns the _ifstmt rule contexts.
	Get_ifstmt() IIfstmtContext

	// Get_while_statement returns the _while_statement rule contexts.
	Get_while_statement() IWhile_statementContext

	// Get_switchstatement returns the _switchstatement rule contexts.
	Get_switchstatement() ISwitchstatementContext

	// Get_forloop returns the _forloop rule contexts.
	Get_forloop() IForloopContext

	// Set_printstmt sets the _printstmt rule contexts.
	Set_printstmt(IPrintstmtContext)

	// Set_vardec sets the _vardec rule contexts.
	Set_vardec(IVardecContext)

	// Set_constdec sets the _constdec rule contexts.
	Set_constdec(IConstdecContext)

	// Set_vecdec sets the _vecdec rule contexts.
	Set_vecdec(IVecdecContext)

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

	// Set_ifstmt sets the _ifstmt rule contexts.
	Set_ifstmt(IIfstmtContext)

	// Set_while_statement sets the _while_statement rule contexts.
	Set_while_statement(IWhile_statementContext)

	// Set_switchstatement sets the _switchstatement rule contexts.
	Set_switchstatement(ISwitchstatementContext)

	// Set_forloop sets the _forloop rule contexts.
	Set_forloop(IForloopContext)

	// GetInst returns the inst attribute.
	GetInst() interfaces.Instruction

	// SetInst sets the inst attribute.
	SetInst(interfaces.Instruction)

	// Getter signatures
	Printstmt() IPrintstmtContext
	PTOCOMA() antlr.TerminalNode
	Vardec() IVardecContext
	Constdec() IConstdecContext
	Vecdec() IVecdecContext
	Appendvec() IAppendvecContext
	Removelastvec() IRemovelastvecContext
	Removeatvec() IRemoveatvecContext
	Asignation() IAsignationContext
	Unarysum() IUnarysumContext
	Unarysub() IUnarysubContext
	Breakstatement() IBreakstatementContext
	Continuestatement() IContinuestatementContext
	Ifstmt() IIfstmtContext
	While_statement() IWhile_statementContext
	Switchstatement() ISwitchstatementContext
	Forloop() IForloopContext

	// IsInstructionContext differentiates from other interfaces.
	IsInstructionContext()
}

type InstructionContext struct {
	antlr.BaseParserRuleContext
	parser             antlr.Parser
	inst               interfaces.Instruction
	_printstmt         IPrintstmtContext
	_vardec            IVardecContext
	_constdec          IConstdecContext
	_vecdec            IVecdecContext
	_appendvec         IAppendvecContext
	_removelastvec     IRemovelastvecContext
	_removeatvec       IRemoveatvecContext
	_asignation        IAsignationContext
	_unarysum          IUnarysumContext
	_unarysub          IUnarysubContext
	_breakstatement    IBreakstatementContext
	_continuestatement IContinuestatementContext
	_ifstmt            IIfstmtContext
	_while_statement   IWhile_statementContext
	_switchstatement   ISwitchstatementContext
	_forloop           IForloopContext
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

func (s *InstructionContext) Get_vardec() IVardecContext { return s._vardec }

func (s *InstructionContext) Get_constdec() IConstdecContext { return s._constdec }

func (s *InstructionContext) Get_vecdec() IVecdecContext { return s._vecdec }

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

func (s *InstructionContext) Get_ifstmt() IIfstmtContext { return s._ifstmt }

func (s *InstructionContext) Get_while_statement() IWhile_statementContext { return s._while_statement }

func (s *InstructionContext) Get_switchstatement() ISwitchstatementContext { return s._switchstatement }

func (s *InstructionContext) Get_forloop() IForloopContext { return s._forloop }

func (s *InstructionContext) Set_printstmt(v IPrintstmtContext) { s._printstmt = v }

func (s *InstructionContext) Set_vardec(v IVardecContext) { s._vardec = v }

func (s *InstructionContext) Set_constdec(v IConstdecContext) { s._constdec = v }

func (s *InstructionContext) Set_vecdec(v IVecdecContext) { s._vecdec = v }

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

func (s *InstructionContext) Set_ifstmt(v IIfstmtContext) { s._ifstmt = v }

func (s *InstructionContext) Set_while_statement(v IWhile_statementContext) { s._while_statement = v }

func (s *InstructionContext) Set_switchstatement(v ISwitchstatementContext) { s._switchstatement = v }

func (s *InstructionContext) Set_forloop(v IForloopContext) { s._forloop = v }

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

	p.SetState(174)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(90)

			var _x = p.Printstmt()

			localctx.(*InstructionContext)._printstmt = _x
		}
		p.SetState(92)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserPTOCOMA {
			{
				p.SetState(91)
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
			p.SetState(96)

			var _x = p.Vardec()

			localctx.(*InstructionContext)._vardec = _x
		}
		p.SetState(98)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserPTOCOMA {
			{
				p.SetState(97)
				p.Match(SwiftGrammarParserPTOCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_vardec().GetNewdec()

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(102)

			var _x = p.Constdec()

			localctx.(*InstructionContext)._constdec = _x
		}
		p.SetState(104)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserPTOCOMA {
			{
				p.SetState(103)
				p.Match(SwiftGrammarParserPTOCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_constdec().GetNewconst()

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(108)

			var _x = p.Vecdec()

			localctx.(*InstructionContext)._vecdec = _x
		}
		p.SetState(110)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserPTOCOMA {
			{
				p.SetState(109)
				p.Match(SwiftGrammarParserPTOCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_vecdec().GetNewvecdec()

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(114)

			var _x = p.Appendvec()

			localctx.(*InstructionContext)._appendvec = _x
		}
		p.SetState(116)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserPTOCOMA {
			{
				p.SetState(115)
				p.Match(SwiftGrammarParserPTOCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_appendvec().GetNewappendvec()

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(120)

			var _x = p.Removelastvec()

			localctx.(*InstructionContext)._removelastvec = _x
		}
		p.SetState(122)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserPTOCOMA {
			{
				p.SetState(121)
				p.Match(SwiftGrammarParserPTOCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_removelastvec().GetNewremovelastvec()

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(126)

			var _x = p.Removeatvec()

			localctx.(*InstructionContext)._removeatvec = _x
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
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_removeatvec().GetNewremoveat()

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(132)

			var _x = p.Asignation()

			localctx.(*InstructionContext)._asignation = _x
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
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_asignation().GetNewasignation()

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(138)

			var _x = p.Unarysum()

			localctx.(*InstructionContext)._unarysum = _x
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
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_unarysum().GetNewunarysum()

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(144)

			var _x = p.Unarysub()

			localctx.(*InstructionContext)._unarysub = _x
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
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_unarysub().GetNewunarysub()

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(150)

			var _x = p.Breakstatement()

			localctx.(*InstructionContext)._breakstatement = _x
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
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_breakstatement().GetNewbreak()

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(156)

			var _x = p.Continuestatement()

			localctx.(*InstructionContext)._continuestatement = _x
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
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_continuestatement().GetNewcontinue()

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(162)

			var _x = p.Ifstmt()

			localctx.(*InstructionContext)._ifstmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_ifstmt().GetNewif()

	case 14:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(165)

			var _x = p.While_statement()

			localctx.(*InstructionContext)._while_statement = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_while_statement().GetNewwhile()

	case 15:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(168)

			var _x = p.Switchstatement()

			localctx.(*InstructionContext)._switchstatement = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_switchstatement().GetNewswitch()

	case 16:
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(171)

			var _x = p.Forloop()

			localctx.(*InstructionContext)._forloop = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_forloop().GetNewfor()

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
	p.EnterRule(localctx, 10, SwiftGrammarParserRULE_forloop)
	p.SetState(194)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(176)

			var _m = p.Match(SwiftGrammarParserRFOR)

			localctx.(*ForloopContext)._RFOR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(177)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*ForloopContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(178)
			p.Match(SwiftGrammarParserRIN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(179)

			var _x = p.expr(0)

			localctx.(*ForloopContext).ex = _x
		}
		{
			p.SetState(180)
			p.Match(SwiftGrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(181)

			var _x = p.Block()

			localctx.(*ForloopContext).b = _x
		}
		{
			p.SetState(182)
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
			p.SetState(185)

			var _m = p.Match(SwiftGrammarParserRFOR)

			localctx.(*ForloopContext)._RFOR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(186)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*ForloopContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(187)
			p.Match(SwiftGrammarParserRIN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(188)

			var _x = p.Range_()

			localctx.(*ForloopContext)._range = _x
		}
		{
			p.SetState(189)
			p.Match(SwiftGrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(190)

			var _x = p.Block()

			localctx.(*ForloopContext).b = _x
		}
		{
			p.SetState(191)
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
	p.EnterRule(localctx, 12, SwiftGrammarParserRULE_range)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(196)

		var _x = p.expr(0)

		localctx.(*RangeContext).exp1 = _x
	}
	{
		p.SetState(197)
		p.Match(SwiftGrammarParserPTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(198)
		p.Match(SwiftGrammarParserPTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(199)
		p.Match(SwiftGrammarParserPTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(200)

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
	p.EnterRule(localctx, 14, SwiftGrammarParserRULE_removeatvec)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(203)

		var _m = p.Match(SwiftGrammarParserID)

		localctx.(*RemoveatvecContext)._ID = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(204)
		p.Match(SwiftGrammarParserPTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(205)
		p.Match(SwiftGrammarParserRREMOVEAT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(206)
		p.Match(SwiftGrammarParserPARIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(207)
		p.Match(SwiftGrammarParserRRAT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(208)
		p.Match(SwiftGrammarParserDOSPTOS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(209)

		var _x = p.expr(0)

		localctx.(*RemoveatvecContext)._expr = _x
	}
	{
		p.SetState(210)
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
	p.EnterRule(localctx, 16, SwiftGrammarParserRULE_appendvec)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(213)

		var _m = p.Match(SwiftGrammarParserID)

		localctx.(*AppendvecContext)._ID = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(214)
		p.Match(SwiftGrammarParserPTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(215)
		p.Match(SwiftGrammarParserRAPPEND)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(216)
		p.Match(SwiftGrammarParserPARIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(217)

		var _x = p.expr(0)

		localctx.(*AppendvecContext)._expr = _x
	}
	{
		p.SetState(218)
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
	p.EnterRule(localctx, 18, SwiftGrammarParserRULE_removelastvec)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(221)

		var _m = p.Match(SwiftGrammarParserID)

		localctx.(*RemovelastvecContext)._ID = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(222)
		p.Match(SwiftGrammarParserPTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(223)
		p.Match(SwiftGrammarParserRREMOVELAST)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(224)
		p.Match(SwiftGrammarParserPARIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(225)
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
	p.EnterRule(localctx, 20, SwiftGrammarParserRULE_vecdec)
	var _la int

	p.SetState(251)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(228)

			var _m = p.Match(SwiftGrammarParserRVAR)

			localctx.(*VecdecContext)._RVAR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(229)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*VecdecContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(230)
			p.Match(SwiftGrammarParserDOSPTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(231)
			p.Match(SwiftGrammarParserOBRA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(232)

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
			p.SetState(233)
			p.Match(SwiftGrammarParserCBRA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(234)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(235)
			p.Match(SwiftGrammarParserOBRA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(236)

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
			p.SetState(237)
			p.Match(SwiftGrammarParserCBRA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(238)
			p.Match(SwiftGrammarParserOBRA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(239)
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
			p.SetState(241)

			var _m = p.Match(SwiftGrammarParserRVAR)

			localctx.(*VecdecContext)._RVAR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(242)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*VecdecContext).firstid = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(243)
			p.Match(SwiftGrammarParserDOSPTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(244)
			p.Match(SwiftGrammarParserOBRA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(245)

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
			p.SetState(246)
			p.Match(SwiftGrammarParserCBRA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(247)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(248)

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
	p.EnterRule(localctx, 22, SwiftGrammarParserRULE_breakstatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(253)

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
	p.EnterRule(localctx, 24, SwiftGrammarParserRULE_continuestatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(256)

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
	p.EnterRule(localctx, 26, SwiftGrammarParserRULE_switchstatement)
	p.SetState(274)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(259)

			var _m = p.Match(SwiftGrammarParserRSWITCH)

			localctx.(*SwitchstatementContext)._RSWITCH = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(260)

			var _x = p.expr(0)

			localctx.(*SwitchstatementContext).ex = _x
		}
		{
			p.SetState(261)
			p.Match(SwiftGrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(262)

			var _x = p.Caselist()

			localctx.(*SwitchstatementContext)._caselist = _x
		}
		{
			p.SetState(263)
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
			p.SetState(266)

			var _m = p.Match(SwiftGrammarParserRSWITCH)

			localctx.(*SwitchstatementContext)._RSWITCH = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(267)

			var _x = p.expr(0)

			localctx.(*SwitchstatementContext).ex = _x
		}
		{
			p.SetState(268)
			p.Match(SwiftGrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(269)

			var _x = p.Caselist()

			localctx.(*SwitchstatementContext)._caselist = _x
		}
		{
			p.SetState(270)

			var _x = p.Defaultstatement()

			localctx.(*SwitchstatementContext)._defaultstatement = _x
		}
		{
			p.SetState(271)
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
	p.EnterRule(localctx, 28, SwiftGrammarParserRULE_caselist)

	localctx.(*CaselistContext).newcaselist = []interface{}{}

	p.SetState(284)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(276)

			var _x = p.Case_()

			localctx.(*CaselistContext)._case = _x
		}
		{
			p.SetState(277)

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
			p.SetState(280)

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
	p.EnterRule(localctx, 30, SwiftGrammarParserRULE_case)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(286)

		var _m = p.Match(SwiftGrammarParserRCASE)

		localctx.(*CaseContext)._RCASE = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(287)

		var _x = p.expr(0)

		localctx.(*CaseContext).ex = _x
	}
	{
		p.SetState(288)
		p.Match(SwiftGrammarParserDOSPTOS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(289)

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
	p.EnterRule(localctx, 32, SwiftGrammarParserRULE_defaultstatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(292)
		p.Match(SwiftGrammarParserRDEFAULT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(293)
		p.Match(SwiftGrammarParserDOSPTOS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(294)

		var _x = p.Block()

		localctx.(*DefaultstatementContext).b = _x
	}
	{
		p.SetState(295)
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
	p.EnterRule(localctx, 34, SwiftGrammarParserRULE_ifstmt)
	p.SetState(315)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(298)

			var _m = p.Match(SwiftGrammarParserRIF)

			localctx.(*IfstmtContext)._RIF = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(299)

			var _x = p.expr(0)

			localctx.(*IfstmtContext).ex = _x
		}
		{
			p.SetState(300)
			p.Match(SwiftGrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(301)

			var _x = p.Block()

			localctx.(*IfstmtContext).b = _x
		}
		{
			p.SetState(302)
			p.Match(SwiftGrammarParserLLAVEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(303)

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
			p.SetState(306)

			var _m = p.Match(SwiftGrammarParserRIF)

			localctx.(*IfstmtContext)._RIF = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(307)

			var _x = p.expr(0)

			localctx.(*IfstmtContext).ex = _x
		}
		{
			p.SetState(308)
			p.Match(SwiftGrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(309)

			var _x = p.Block()

			localctx.(*IfstmtContext).b = _x
		}
		{
			p.SetState(310)
			p.Match(SwiftGrammarParserLLAVEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(311)

			var _x = p.Eliflist()

			localctx.(*IfstmtContext)._eliflist = _x
		}
		{
			p.SetState(312)

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
	p.EnterRule(localctx, 36, SwiftGrammarParserRULE_eliflist)

	localctx.(*EliflistContext).neweliflist = []interface{}{}

	p.SetState(325)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(317)

			var _x = p.Elif()

			localctx.(*EliflistContext)._elif = _x
		}
		{
			p.SetState(318)

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
			p.SetState(321)

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
	p.EnterRule(localctx, 38, SwiftGrammarParserRULE_elif)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(327)

		var _m = p.Match(SwiftGrammarParserRELSE)

		localctx.(*ElifContext)._RELSE = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(328)
		p.Match(SwiftGrammarParserRIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(329)

		var _x = p.expr(0)

		localctx.(*ElifContext).ex = _x
	}
	{
		p.SetState(330)
		p.Match(SwiftGrammarParserLLAVEIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(331)

		var _x = p.Block()

		localctx.(*ElifContext).b = _x
	}
	{
		p.SetState(332)
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
	p.EnterRule(localctx, 40, SwiftGrammarParserRULE_elsestament)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(335)
		p.Match(SwiftGrammarParserRELSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(336)
		p.Match(SwiftGrammarParserLLAVEIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(337)

		var _x = p.Block()

		localctx.(*ElsestamentContext).b = _x
	}
	{
		p.SetState(338)
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
	p.EnterRule(localctx, 42, SwiftGrammarParserRULE_printstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(341)

		var _m = p.Match(SwiftGrammarParserRPRINT)

		localctx.(*PrintstmtContext)._RPRINT = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(342)
		p.Match(SwiftGrammarParserPARIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(343)

		var _x = p.Arguments()

		localctx.(*PrintstmtContext)._arguments = _x
	}
	{
		p.SetState(344)
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
	p.EnterRule(localctx, 44, SwiftGrammarParserRULE_while_statement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(347)

		var _m = p.Match(SwiftGrammarParserRWHILE)

		localctx.(*While_statementContext)._RWHILE = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(348)

		var _x = p.expr(0)

		localctx.(*While_statementContext)._expr = _x
	}
	{
		p.SetState(349)
		p.Match(SwiftGrammarParserLLAVEIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(350)

		var _x = p.Block()

		localctx.(*While_statementContext).b = _x
	}
	{
		p.SetState(351)
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
	p.EnterRule(localctx, 46, SwiftGrammarParserRULE_vardec)
	var _la int

	p.SetState(374)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(354)

			var _m = p.Match(SwiftGrammarParserRVAR)

			localctx.(*VardecContext)._RVAR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(355)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*VardecContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(356)
			p.Match(SwiftGrammarParserDOSPTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(357)

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
			p.SetState(358)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(359)

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
			p.SetState(362)

			var _m = p.Match(SwiftGrammarParserRVAR)

			localctx.(*VardecContext)._RVAR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(363)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*VardecContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(364)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(365)

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
			p.SetState(368)

			var _m = p.Match(SwiftGrammarParserRVAR)

			localctx.(*VardecContext)._RVAR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(369)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*VardecContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(370)
			p.Match(SwiftGrammarParserDOSPTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(371)

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
			p.SetState(372)
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
	p.EnterRule(localctx, 48, SwiftGrammarParserRULE_constdec)
	var _la int

	p.SetState(390)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(376)

			var _m = p.Match(SwiftGrammarParserRLET)

			localctx.(*ConstdecContext)._RLET = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(377)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*ConstdecContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(378)
			p.Match(SwiftGrammarParserDOSPTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(379)

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
			p.SetState(380)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(381)

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
			p.SetState(384)

			var _m = p.Match(SwiftGrammarParserRLET)

			localctx.(*ConstdecContext)._RLET = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(385)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*ConstdecContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(386)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(387)

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
	p.EnterRule(localctx, 50, SwiftGrammarParserRULE_asignation)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(392)

		var _m = p.Match(SwiftGrammarParserID)

		localctx.(*AsignationContext)._ID = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(393)
		p.Match(SwiftGrammarParserIG)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(394)

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
	p.EnterRule(localctx, 52, SwiftGrammarParserRULE_unarysum)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(397)

		var _m = p.Match(SwiftGrammarParserID)

		localctx.(*UnarysumContext)._ID = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(398)
		p.Match(SwiftGrammarParserUNARYPLUS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(399)

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
	p.EnterRule(localctx, 54, SwiftGrammarParserRULE_unarysub)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(402)

		var _m = p.Match(SwiftGrammarParserID)

		localctx.(*UnarysubContext)._ID = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(403)
		p.Match(SwiftGrammarParserUNARYMINUS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(404)

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
	p.EnterRule(localctx, 56, SwiftGrammarParserRULE_isemptyvec)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(407)

		var _m = p.Match(SwiftGrammarParserID)

		localctx.(*IsemptyvecContext)._ID = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(408)
		p.Match(SwiftGrammarParserPTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(409)
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
	p.EnterRule(localctx, 58, SwiftGrammarParserRULE_countvec)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(412)

		var _m = p.Match(SwiftGrammarParserID)

		localctx.(*CountvecContext)._ID = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(413)
		p.Match(SwiftGrammarParserPTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(414)
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

	// GetEx returns the ex rule contexts.
	GetEx() IExprContext

	// SetEx sets the ex rule contexts.
	SetEx(IExprContext)

	// GetNewvecaccess returns the newvecaccess attribute.
	GetNewvecaccess() interfaces.Expression

	// SetNewvecaccess sets the newvecaccess attribute.
	SetNewvecaccess(interfaces.Expression)

	// Getter signatures
	ID() antlr.TerminalNode
	OBRA() antlr.TerminalNode
	CBRA() antlr.TerminalNode
	Expr() IExprContext

	// IsVectoraccessContext differentiates from other interfaces.
	IsVectoraccessContext()
}

type VectoraccessContext struct {
	antlr.BaseParserRuleContext
	parser       antlr.Parser
	newvecaccess interfaces.Expression
	_ID          antlr.Token
	ex           IExprContext
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

func (s *VectoraccessContext) GetEx() IExprContext { return s.ex }

func (s *VectoraccessContext) SetEx(v IExprContext) { s.ex = v }

func (s *VectoraccessContext) GetNewvecaccess() interfaces.Expression { return s.newvecaccess }

func (s *VectoraccessContext) SetNewvecaccess(v interfaces.Expression) { s.newvecaccess = v }

func (s *VectoraccessContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *VectoraccessContext) OBRA() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserOBRA, 0)
}

func (s *VectoraccessContext) CBRA() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCBRA, 0)
}

func (s *VectoraccessContext) Expr() IExprContext {
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
	p.EnterRule(localctx, 60, SwiftGrammarParserRULE_vectoraccess)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(417)

		var _m = p.Match(SwiftGrammarParserID)

		localctx.(*VectoraccessContext)._ID = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(418)
		p.Match(SwiftGrammarParserOBRA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(419)

		var _x = p.expr(0)

		localctx.(*VectoraccessContext).ex = _x
	}
	{
		p.SetState(420)
		p.Match(SwiftGrammarParserCBRA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
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
	}()), localctx.(*VectoraccessContext).GetEx().GetE())

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

func (s *ExprContext) Get_isemptyvec() IIsemptyvecContext { return s._isemptyvec }

func (s *ExprContext) Get_countvec() ICountvecContext { return s._countvec }

func (s *ExprContext) Get_vectoraccess() IVectoraccessContext { return s._vectoraccess }

func (s *ExprContext) Get_arguments() IArgumentsContext { return s._arguments }

func (s *ExprContext) GetRight() IExprContext { return s.right }

func (s *ExprContext) SetLeft(v IExprContext) { s.left = v }

func (s *ExprContext) Set_expr(v IExprContext) { s._expr = v }

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
	_startState := 62
	p.EnterRecursionRule(localctx, 62, SwiftGrammarParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(459)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 23, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(424)
			p.Match(SwiftGrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(425)

			var _x = p.expr(0)

			localctx.(*ExprContext)._expr = _x
		}
		{
			p.SetState(426)
			p.Match(SwiftGrammarParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ExprContext).e = localctx.(*ExprContext).Get_expr().GetE()

	case 2:
		{
			p.SetState(429)

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
			p.SetState(430)

			var _x = p.expr(17)

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
			p.SetState(433)

			var _x = p.Isemptyvec()

			localctx.(*ExprContext)._isemptyvec = _x
		}
		localctx.(*ExprContext).e = localctx.(*ExprContext).Get_isemptyvec().GetNewisemptyvec()

	case 4:
		{
			p.SetState(436)

			var _x = p.Countvec()

			localctx.(*ExprContext)._countvec = _x
		}
		localctx.(*ExprContext).e = localctx.(*ExprContext).Get_countvec().GetNewcountvec()

	case 5:
		{
			p.SetState(439)

			var _x = p.Vectoraccess()

			localctx.(*ExprContext)._vectoraccess = _x
		}
		localctx.(*ExprContext).e = localctx.(*ExprContext).Get_vectoraccess().GetNewvecaccess()

	case 6:
		{
			p.SetState(442)

			var _m = p.Match(SwiftGrammarParserOBRA)

			localctx.(*ExprContext)._OBRA = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(443)

			var _x = p.Arguments()

			localctx.(*ExprContext)._arguments = _x
		}
		{
			p.SetState(444)
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

	case 7:
		{
			p.SetState(447)

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

	case 8:
		{
			p.SetState(449)

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

	case 9:
		{
			p.SetState(451)

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

	case 10:
		{
			p.SetState(453)

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

	case 11:
		{
			p.SetState(455)

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

	case 12:
		{
			p.SetState(457)
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
	p.SetState(493)
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
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(491)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 24, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_expr)
				p.SetState(461)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
					goto errorExit
				}
				{
					p.SetState(462)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&36055185298030592) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(463)

					var _x = p.expr(17)

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
				p.SetState(466)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
					goto errorExit
				}
				{
					p.SetState(467)

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
					p.SetState(468)

					var _x = p.expr(16)

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
				p.SetState(471)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
					goto errorExit
				}
				{
					p.SetState(472)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8246337208320) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(473)

					var _x = p.expr(15)

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
				p.SetState(476)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
					goto errorExit
				}
				{
					p.SetState(477)

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
					p.SetState(478)

					var _x = p.expr(14)

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
				p.SetState(481)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
					goto errorExit
				}
				{
					p.SetState(482)

					var _m = p.Match(SwiftGrammarParserAND)

					localctx.(*ExprContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(483)

					var _x = p.expr(13)

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
				p.SetState(486)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
					goto errorExit
				}
				{
					p.SetState(487)

					var _m = p.Match(SwiftGrammarParserOR)

					localctx.(*ExprContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(488)

					var _x = p.expr(12)

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
		p.SetState(495)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 25, p.GetParserRuleContext())
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
	case 31:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *SwiftGrammarParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 16)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 15)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 11)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
