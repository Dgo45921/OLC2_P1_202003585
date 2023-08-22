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
		"'remove'", "", "", "", "'+='", "'-='", "'!='", "'=='", "'!'", "'||'",
		"'&&'", "'='", "'>='", "'<='", "'>'", "'<'", "'*'", "'/'", "'+'", "'-'",
		"'('", "')'", "'{'", "'}'", "'['", "']'", "';'", "':'", "'%'", "','",
		"'?'", "'.'",
	}
	staticData.SymbolicNames = []string{
		"", "RINT", "RFLOAT", "RBOOL", "RSTRING", "RCHARACTER", "RTRUE", "RFALSE",
		"RPRINT", "RIF", "RELSE", "RWHILE", "RVAR", "RLET", "RNIL", "RBREAK",
		"RCONTINUE", "RAPPEND", "RREMOVELAST", "RRAT", "RREMOVEAT", "NUMBER",
		"STRING", "ID", "UNARYPLUS", "UNARYMINUS", "DIF", "IG_IG", "NOT", "OR",
		"AND", "IG", "MAY_IG", "MEN_IG", "MAYOR", "MENOR", "MUL", "DIV", "ADD",
		"SUB", "PARIZQ", "PARDER", "LLAVEIZQ", "LLAVEDER", "OBRA", "CBRA", "PTOCOMA",
		"DOSPTOS", "MODULE", "COMA", "QM", "PTO", "WHITESPACE", "COMMENT", "LINE_COMMENT",
	}
	staticData.RuleNames = []string{
		"s", "block", "arguments", "argument", "instruction", "removeatvec",
		"appendvec", "removelastvec", "vecdec", "breakstatement", "continuestatement",
		"ifstmt", "elseifstatement", "printstmt", "while_statement", "vardec",
		"constdec", "asignation", "unarysum", "unarysub", "expr",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 54, 367, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 1, 4, 1, 48, 8, 1, 11, 1, 12, 1, 49, 1, 1, 1, 1,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 63, 8, 2, 1,
		3, 1, 3, 1, 3, 1, 4, 1, 4, 3, 4, 70, 8, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4,
		76, 8, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 82, 8, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 3, 4, 88, 8, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 94, 8, 4, 1, 4, 1, 4,
		1, 4, 1, 4, 3, 4, 100, 8, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 106, 8, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 3, 4, 112, 8, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 118,
		8, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 124, 8, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		3, 4, 130, 8, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 136, 8, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 146, 8, 4, 1, 5, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1,
		8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1,
		8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1,
		8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 208, 8, 8, 1,
		9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11,
		1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1,
		12, 1, 12, 1, 12, 1, 12, 3, 12, 235, 8, 12, 1, 13, 1, 13, 1, 13, 1, 13,
		1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1,
		15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15,
		1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 3, 15, 270, 8,
		15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16,
		1, 16, 1, 16, 1, 16, 1, 16, 3, 16, 286, 8, 16, 1, 17, 1, 17, 1, 17, 1,
		17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 19,
		1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1,
		20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20,
		1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 3, 20, 330, 8, 20, 1,
		20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20,
		1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1,
		20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 5, 20, 362,
		8, 20, 10, 20, 12, 20, 365, 9, 20, 1, 20, 0, 1, 40, 21, 0, 2, 4, 6, 8,
		10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 0, 6, 1,
		0, 1, 5, 2, 0, 28, 28, 39, 39, 2, 0, 36, 37, 48, 48, 1, 0, 38, 39, 1, 0,
		32, 35, 1, 0, 26, 27, 394, 0, 42, 1, 0, 0, 0, 2, 47, 1, 0, 0, 0, 4, 62,
		1, 0, 0, 0, 6, 64, 1, 0, 0, 0, 8, 145, 1, 0, 0, 0, 10, 147, 1, 0, 0, 0,
		12, 157, 1, 0, 0, 0, 14, 165, 1, 0, 0, 0, 16, 207, 1, 0, 0, 0, 18, 209,
		1, 0, 0, 0, 20, 212, 1, 0, 0, 0, 22, 215, 1, 0, 0, 0, 24, 234, 1, 0, 0,
		0, 26, 236, 1, 0, 0, 0, 28, 242, 1, 0, 0, 0, 30, 269, 1, 0, 0, 0, 32, 285,
		1, 0, 0, 0, 34, 287, 1, 0, 0, 0, 36, 292, 1, 0, 0, 0, 38, 297, 1, 0, 0,
		0, 40, 329, 1, 0, 0, 0, 42, 43, 3, 2, 1, 0, 43, 44, 5, 0, 0, 1, 44, 45,
		6, 0, -1, 0, 45, 1, 1, 0, 0, 0, 46, 48, 3, 8, 4, 0, 47, 46, 1, 0, 0, 0,
		48, 49, 1, 0, 0, 0, 49, 47, 1, 0, 0, 0, 49, 50, 1, 0, 0, 0, 50, 51, 1,
		0, 0, 0, 51, 52, 6, 1, -1, 0, 52, 3, 1, 0, 0, 0, 53, 54, 3, 6, 3, 0, 54,
		55, 5, 49, 0, 0, 55, 56, 3, 4, 2, 0, 56, 57, 6, 2, -1, 0, 57, 63, 1, 0,
		0, 0, 58, 59, 3, 6, 3, 0, 59, 60, 6, 2, -1, 0, 60, 63, 1, 0, 0, 0, 61,
		63, 6, 2, -1, 0, 62, 53, 1, 0, 0, 0, 62, 58, 1, 0, 0, 0, 62, 61, 1, 0,
		0, 0, 63, 5, 1, 0, 0, 0, 64, 65, 3, 40, 20, 0, 65, 66, 6, 3, -1, 0, 66,
		7, 1, 0, 0, 0, 67, 69, 3, 26, 13, 0, 68, 70, 5, 46, 0, 0, 69, 68, 1, 0,
		0, 0, 69, 70, 1, 0, 0, 0, 70, 71, 1, 0, 0, 0, 71, 72, 6, 4, -1, 0, 72,
		146, 1, 0, 0, 0, 73, 75, 3, 30, 15, 0, 74, 76, 5, 46, 0, 0, 75, 74, 1,
		0, 0, 0, 75, 76, 1, 0, 0, 0, 76, 77, 1, 0, 0, 0, 77, 78, 6, 4, -1, 0, 78,
		146, 1, 0, 0, 0, 79, 81, 3, 32, 16, 0, 80, 82, 5, 46, 0, 0, 81, 80, 1,
		0, 0, 0, 81, 82, 1, 0, 0, 0, 82, 83, 1, 0, 0, 0, 83, 84, 6, 4, -1, 0, 84,
		146, 1, 0, 0, 0, 85, 87, 3, 16, 8, 0, 86, 88, 5, 46, 0, 0, 87, 86, 1, 0,
		0, 0, 87, 88, 1, 0, 0, 0, 88, 89, 1, 0, 0, 0, 89, 90, 6, 4, -1, 0, 90,
		146, 1, 0, 0, 0, 91, 93, 3, 12, 6, 0, 92, 94, 5, 46, 0, 0, 93, 92, 1, 0,
		0, 0, 93, 94, 1, 0, 0, 0, 94, 95, 1, 0, 0, 0, 95, 96, 6, 4, -1, 0, 96,
		146, 1, 0, 0, 0, 97, 99, 3, 14, 7, 0, 98, 100, 5, 46, 0, 0, 99, 98, 1,
		0, 0, 0, 99, 100, 1, 0, 0, 0, 100, 101, 1, 0, 0, 0, 101, 102, 6, 4, -1,
		0, 102, 146, 1, 0, 0, 0, 103, 105, 3, 10, 5, 0, 104, 106, 5, 46, 0, 0,
		105, 104, 1, 0, 0, 0, 105, 106, 1, 0, 0, 0, 106, 107, 1, 0, 0, 0, 107,
		108, 6, 4, -1, 0, 108, 146, 1, 0, 0, 0, 109, 111, 3, 34, 17, 0, 110, 112,
		5, 46, 0, 0, 111, 110, 1, 0, 0, 0, 111, 112, 1, 0, 0, 0, 112, 113, 1, 0,
		0, 0, 113, 114, 6, 4, -1, 0, 114, 146, 1, 0, 0, 0, 115, 117, 3, 36, 18,
		0, 116, 118, 5, 46, 0, 0, 117, 116, 1, 0, 0, 0, 117, 118, 1, 0, 0, 0, 118,
		119, 1, 0, 0, 0, 119, 120, 6, 4, -1, 0, 120, 146, 1, 0, 0, 0, 121, 123,
		3, 38, 19, 0, 122, 124, 5, 46, 0, 0, 123, 122, 1, 0, 0, 0, 123, 124, 1,
		0, 0, 0, 124, 125, 1, 0, 0, 0, 125, 126, 6, 4, -1, 0, 126, 146, 1, 0, 0,
		0, 127, 129, 3, 18, 9, 0, 128, 130, 5, 46, 0, 0, 129, 128, 1, 0, 0, 0,
		129, 130, 1, 0, 0, 0, 130, 131, 1, 0, 0, 0, 131, 132, 6, 4, -1, 0, 132,
		146, 1, 0, 0, 0, 133, 135, 3, 20, 10, 0, 134, 136, 5, 46, 0, 0, 135, 134,
		1, 0, 0, 0, 135, 136, 1, 0, 0, 0, 136, 137, 1, 0, 0, 0, 137, 138, 6, 4,
		-1, 0, 138, 146, 1, 0, 0, 0, 139, 140, 3, 22, 11, 0, 140, 141, 6, 4, -1,
		0, 141, 146, 1, 0, 0, 0, 142, 143, 3, 28, 14, 0, 143, 144, 6, 4, -1, 0,
		144, 146, 1, 0, 0, 0, 145, 67, 1, 0, 0, 0, 145, 73, 1, 0, 0, 0, 145, 79,
		1, 0, 0, 0, 145, 85, 1, 0, 0, 0, 145, 91, 1, 0, 0, 0, 145, 97, 1, 0, 0,
		0, 145, 103, 1, 0, 0, 0, 145, 109, 1, 0, 0, 0, 145, 115, 1, 0, 0, 0, 145,
		121, 1, 0, 0, 0, 145, 127, 1, 0, 0, 0, 145, 133, 1, 0, 0, 0, 145, 139,
		1, 0, 0, 0, 145, 142, 1, 0, 0, 0, 146, 9, 1, 0, 0, 0, 147, 148, 5, 23,
		0, 0, 148, 149, 5, 51, 0, 0, 149, 150, 5, 20, 0, 0, 150, 151, 5, 40, 0,
		0, 151, 152, 5, 19, 0, 0, 152, 153, 5, 47, 0, 0, 153, 154, 3, 40, 20, 0,
		154, 155, 5, 41, 0, 0, 155, 156, 6, 5, -1, 0, 156, 11, 1, 0, 0, 0, 157,
		158, 5, 23, 0, 0, 158, 159, 5, 51, 0, 0, 159, 160, 5, 17, 0, 0, 160, 161,
		5, 40, 0, 0, 161, 162, 3, 40, 20, 0, 162, 163, 5, 41, 0, 0, 163, 164, 6,
		6, -1, 0, 164, 13, 1, 0, 0, 0, 165, 166, 5, 23, 0, 0, 166, 167, 5, 51,
		0, 0, 167, 168, 5, 18, 0, 0, 168, 169, 5, 40, 0, 0, 169, 170, 5, 41, 0,
		0, 170, 171, 6, 7, -1, 0, 171, 15, 1, 0, 0, 0, 172, 173, 5, 12, 0, 0, 173,
		174, 5, 23, 0, 0, 174, 175, 5, 47, 0, 0, 175, 176, 5, 44, 0, 0, 176, 177,
		7, 0, 0, 0, 177, 178, 5, 45, 0, 0, 178, 179, 5, 31, 0, 0, 179, 180, 5,
		44, 0, 0, 180, 181, 3, 40, 20, 0, 181, 182, 5, 45, 0, 0, 182, 183, 6, 8,
		-1, 0, 183, 208, 1, 0, 0, 0, 184, 185, 5, 12, 0, 0, 185, 186, 5, 23, 0,
		0, 186, 187, 5, 47, 0, 0, 187, 188, 5, 44, 0, 0, 188, 189, 7, 0, 0, 0,
		189, 190, 5, 45, 0, 0, 190, 191, 5, 31, 0, 0, 191, 192, 5, 44, 0, 0, 192,
		193, 7, 0, 0, 0, 193, 194, 5, 45, 0, 0, 194, 195, 5, 44, 0, 0, 195, 196,
		5, 45, 0, 0, 196, 208, 6, 8, -1, 0, 197, 198, 5, 12, 0, 0, 198, 199, 5,
		23, 0, 0, 199, 200, 5, 47, 0, 0, 200, 201, 5, 44, 0, 0, 201, 202, 7, 0,
		0, 0, 202, 203, 5, 45, 0, 0, 203, 204, 5, 31, 0, 0, 204, 205, 3, 40, 20,
		0, 205, 206, 6, 8, -1, 0, 206, 208, 1, 0, 0, 0, 207, 172, 1, 0, 0, 0, 207,
		184, 1, 0, 0, 0, 207, 197, 1, 0, 0, 0, 208, 17, 1, 0, 0, 0, 209, 210, 5,
		15, 0, 0, 210, 211, 6, 9, -1, 0, 211, 19, 1, 0, 0, 0, 212, 213, 5, 16,
		0, 0, 213, 214, 6, 10, -1, 0, 214, 21, 1, 0, 0, 0, 215, 216, 5, 9, 0, 0,
		216, 217, 3, 40, 20, 0, 217, 218, 5, 42, 0, 0, 218, 219, 3, 2, 1, 0, 219,
		220, 5, 43, 0, 0, 220, 221, 3, 24, 12, 0, 221, 222, 6, 11, -1, 0, 222,
		23, 1, 0, 0, 0, 223, 224, 5, 10, 0, 0, 224, 225, 3, 22, 11, 0, 225, 226,
		6, 12, -1, 0, 226, 235, 1, 0, 0, 0, 227, 228, 5, 10, 0, 0, 228, 229, 5,
		42, 0, 0, 229, 230, 3, 2, 1, 0, 230, 231, 5, 43, 0, 0, 231, 232, 6, 12,
		-1, 0, 232, 235, 1, 0, 0, 0, 233, 235, 6, 12, -1, 0, 234, 223, 1, 0, 0,
		0, 234, 227, 1, 0, 0, 0, 234, 233, 1, 0, 0, 0, 235, 25, 1, 0, 0, 0, 236,
		237, 5, 8, 0, 0, 237, 238, 5, 40, 0, 0, 238, 239, 3, 4, 2, 0, 239, 240,
		5, 41, 0, 0, 240, 241, 6, 13, -1, 0, 241, 27, 1, 0, 0, 0, 242, 243, 5,
		11, 0, 0, 243, 244, 3, 40, 20, 0, 244, 245, 5, 42, 0, 0, 245, 246, 3, 2,
		1, 0, 246, 247, 5, 43, 0, 0, 247, 248, 6, 14, -1, 0, 248, 29, 1, 0, 0,
		0, 249, 250, 5, 12, 0, 0, 250, 251, 5, 23, 0, 0, 251, 252, 5, 47, 0, 0,
		252, 253, 7, 0, 0, 0, 253, 254, 5, 31, 0, 0, 254, 255, 3, 40, 20, 0, 255,
		256, 6, 15, -1, 0, 256, 270, 1, 0, 0, 0, 257, 258, 5, 12, 0, 0, 258, 259,
		5, 23, 0, 0, 259, 260, 5, 31, 0, 0, 260, 261, 3, 40, 20, 0, 261, 262, 6,
		15, -1, 0, 262, 270, 1, 0, 0, 0, 263, 264, 5, 12, 0, 0, 264, 265, 5, 23,
		0, 0, 265, 266, 5, 47, 0, 0, 266, 267, 7, 0, 0, 0, 267, 268, 5, 50, 0,
		0, 268, 270, 6, 15, -1, 0, 269, 249, 1, 0, 0, 0, 269, 257, 1, 0, 0, 0,
		269, 263, 1, 0, 0, 0, 270, 31, 1, 0, 0, 0, 271, 272, 5, 13, 0, 0, 272,
		273, 5, 23, 0, 0, 273, 274, 5, 47, 0, 0, 274, 275, 7, 0, 0, 0, 275, 276,
		5, 31, 0, 0, 276, 277, 3, 40, 20, 0, 277, 278, 6, 16, -1, 0, 278, 286,
		1, 0, 0, 0, 279, 280, 5, 13, 0, 0, 280, 281, 5, 23, 0, 0, 281, 282, 5,
		31, 0, 0, 282, 283, 3, 40, 20, 0, 283, 284, 6, 16, -1, 0, 284, 286, 1,
		0, 0, 0, 285, 271, 1, 0, 0, 0, 285, 279, 1, 0, 0, 0, 286, 33, 1, 0, 0,
		0, 287, 288, 5, 23, 0, 0, 288, 289, 5, 31, 0, 0, 289, 290, 3, 40, 20, 0,
		290, 291, 6, 17, -1, 0, 291, 35, 1, 0, 0, 0, 292, 293, 5, 23, 0, 0, 293,
		294, 5, 24, 0, 0, 294, 295, 3, 40, 20, 0, 295, 296, 6, 18, -1, 0, 296,
		37, 1, 0, 0, 0, 297, 298, 5, 23, 0, 0, 298, 299, 5, 25, 0, 0, 299, 300,
		3, 40, 20, 0, 300, 301, 6, 19, -1, 0, 301, 39, 1, 0, 0, 0, 302, 303, 6,
		20, -1, 0, 303, 304, 5, 40, 0, 0, 304, 305, 3, 40, 20, 0, 305, 306, 5,
		41, 0, 0, 306, 307, 6, 20, -1, 0, 307, 330, 1, 0, 0, 0, 308, 309, 7, 1,
		0, 0, 309, 310, 3, 40, 20, 14, 310, 311, 6, 20, -1, 0, 311, 330, 1, 0,
		0, 0, 312, 313, 5, 44, 0, 0, 313, 314, 3, 4, 2, 0, 314, 315, 5, 45, 0,
		0, 315, 316, 6, 20, -1, 0, 316, 330, 1, 0, 0, 0, 317, 318, 5, 23, 0, 0,
		318, 330, 6, 20, -1, 0, 319, 320, 5, 21, 0, 0, 320, 330, 6, 20, -1, 0,
		321, 322, 5, 22, 0, 0, 322, 330, 6, 20, -1, 0, 323, 324, 5, 6, 0, 0, 324,
		330, 6, 20, -1, 0, 325, 326, 5, 7, 0, 0, 326, 330, 6, 20, -1, 0, 327, 328,
		5, 14, 0, 0, 328, 330, 6, 20, -1, 0, 329, 302, 1, 0, 0, 0, 329, 308, 1,
		0, 0, 0, 329, 312, 1, 0, 0, 0, 329, 317, 1, 0, 0, 0, 329, 319, 1, 0, 0,
		0, 329, 321, 1, 0, 0, 0, 329, 323, 1, 0, 0, 0, 329, 325, 1, 0, 0, 0, 329,
		327, 1, 0, 0, 0, 330, 363, 1, 0, 0, 0, 331, 332, 10, 13, 0, 0, 332, 333,
		7, 2, 0, 0, 333, 334, 3, 40, 20, 14, 334, 335, 6, 20, -1, 0, 335, 362,
		1, 0, 0, 0, 336, 337, 10, 12, 0, 0, 337, 338, 7, 3, 0, 0, 338, 339, 3,
		40, 20, 13, 339, 340, 6, 20, -1, 0, 340, 362, 1, 0, 0, 0, 341, 342, 10,
		11, 0, 0, 342, 343, 7, 4, 0, 0, 343, 344, 3, 40, 20, 12, 344, 345, 6, 20,
		-1, 0, 345, 362, 1, 0, 0, 0, 346, 347, 10, 10, 0, 0, 347, 348, 7, 5, 0,
		0, 348, 349, 3, 40, 20, 11, 349, 350, 6, 20, -1, 0, 350, 362, 1, 0, 0,
		0, 351, 352, 10, 9, 0, 0, 352, 353, 5, 30, 0, 0, 353, 354, 3, 40, 20, 10,
		354, 355, 6, 20, -1, 0, 355, 362, 1, 0, 0, 0, 356, 357, 10, 8, 0, 0, 357,
		358, 5, 29, 0, 0, 358, 359, 3, 40, 20, 9, 359, 360, 6, 20, -1, 0, 360,
		362, 1, 0, 0, 0, 361, 331, 1, 0, 0, 0, 361, 336, 1, 0, 0, 0, 361, 341,
		1, 0, 0, 0, 361, 346, 1, 0, 0, 0, 361, 351, 1, 0, 0, 0, 361, 356, 1, 0,
		0, 0, 362, 365, 1, 0, 0, 0, 363, 361, 1, 0, 0, 0, 363, 364, 1, 0, 0, 0,
		364, 41, 1, 0, 0, 0, 365, 363, 1, 0, 0, 0, 22, 49, 62, 69, 75, 81, 87,
		93, 99, 105, 111, 117, 123, 129, 135, 145, 207, 234, 269, 285, 329, 361,
		363,
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
	SwiftGrammarParserNUMBER       = 21
	SwiftGrammarParserSTRING       = 22
	SwiftGrammarParserID           = 23
	SwiftGrammarParserUNARYPLUS    = 24
	SwiftGrammarParserUNARYMINUS   = 25
	SwiftGrammarParserDIF          = 26
	SwiftGrammarParserIG_IG        = 27
	SwiftGrammarParserNOT          = 28
	SwiftGrammarParserOR           = 29
	SwiftGrammarParserAND          = 30
	SwiftGrammarParserIG           = 31
	SwiftGrammarParserMAY_IG       = 32
	SwiftGrammarParserMEN_IG       = 33
	SwiftGrammarParserMAYOR        = 34
	SwiftGrammarParserMENOR        = 35
	SwiftGrammarParserMUL          = 36
	SwiftGrammarParserDIV          = 37
	SwiftGrammarParserADD          = 38
	SwiftGrammarParserSUB          = 39
	SwiftGrammarParserPARIZQ       = 40
	SwiftGrammarParserPARDER       = 41
	SwiftGrammarParserLLAVEIZQ     = 42
	SwiftGrammarParserLLAVEDER     = 43
	SwiftGrammarParserOBRA         = 44
	SwiftGrammarParserCBRA         = 45
	SwiftGrammarParserPTOCOMA      = 46
	SwiftGrammarParserDOSPTOS      = 47
	SwiftGrammarParserMODULE       = 48
	SwiftGrammarParserCOMA         = 49
	SwiftGrammarParserQM           = 50
	SwiftGrammarParserPTO          = 51
	SwiftGrammarParserWHITESPACE   = 52
	SwiftGrammarParserCOMMENT      = 53
	SwiftGrammarParserLINE_COMMENT = 54
)

// SwiftGrammarParser rules.
const (
	SwiftGrammarParserRULE_s                 = 0
	SwiftGrammarParserRULE_block             = 1
	SwiftGrammarParserRULE_arguments         = 2
	SwiftGrammarParserRULE_argument          = 3
	SwiftGrammarParserRULE_instruction       = 4
	SwiftGrammarParserRULE_removeatvec       = 5
	SwiftGrammarParserRULE_appendvec         = 6
	SwiftGrammarParserRULE_removelastvec     = 7
	SwiftGrammarParserRULE_vecdec            = 8
	SwiftGrammarParserRULE_breakstatement    = 9
	SwiftGrammarParserRULE_continuestatement = 10
	SwiftGrammarParserRULE_ifstmt            = 11
	SwiftGrammarParserRULE_elseifstatement   = 12
	SwiftGrammarParserRULE_printstmt         = 13
	SwiftGrammarParserRULE_while_statement   = 14
	SwiftGrammarParserRULE_vardec            = 15
	SwiftGrammarParserRULE_constdec          = 16
	SwiftGrammarParserRULE_asignation        = 17
	SwiftGrammarParserRULE_unarysum          = 18
	SwiftGrammarParserRULE_unarysub          = 19
	SwiftGrammarParserRULE_expr              = 20
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
		p.SetState(42)

		var _x = p.Block()

		localctx.(*SContext)._block = _x
	}
	{
		p.SetState(43)
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
	p.SetState(47)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8502016) != 0) {
		{
			p.SetState(46)

			var _x = p.Instruction()

			localctx.(*BlockContext)._instruction = _x
		}
		localctx.(*BlockContext).ins = append(localctx.(*BlockContext).ins, localctx.(*BlockContext)._instruction)

		p.SetState(49)
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

	p.SetState(62)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(53)

			var _x = p.Argument()

			localctx.(*ArgumentsContext)._argument = _x
		}
		{
			p.SetState(54)
			p.Match(SwiftGrammarParserCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(55)

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
			p.SetState(58)

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
		p.SetState(64)

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

	p.SetState(145)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(67)

			var _x = p.Printstmt()

			localctx.(*InstructionContext)._printstmt = _x
		}
		p.SetState(69)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserPTOCOMA {
			{
				p.SetState(68)
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
			p.SetState(73)

			var _x = p.Vardec()

			localctx.(*InstructionContext)._vardec = _x
		}
		p.SetState(75)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserPTOCOMA {
			{
				p.SetState(74)
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
			p.SetState(79)

			var _x = p.Constdec()

			localctx.(*InstructionContext)._constdec = _x
		}
		p.SetState(81)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserPTOCOMA {
			{
				p.SetState(80)
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
			p.SetState(85)

			var _x = p.Vecdec()

			localctx.(*InstructionContext)._vecdec = _x
		}
		p.SetState(87)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserPTOCOMA {
			{
				p.SetState(86)
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
			p.SetState(91)

			var _x = p.Appendvec()

			localctx.(*InstructionContext)._appendvec = _x
		}
		p.SetState(93)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserPTOCOMA {
			{
				p.SetState(92)
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
			p.SetState(97)

			var _x = p.Removelastvec()

			localctx.(*InstructionContext)._removelastvec = _x
		}
		p.SetState(99)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserPTOCOMA {
			{
				p.SetState(98)
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
			p.SetState(103)

			var _x = p.Removeatvec()

			localctx.(*InstructionContext)._removeatvec = _x
		}
		p.SetState(105)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserPTOCOMA {
			{
				p.SetState(104)
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
			p.SetState(109)

			var _x = p.Asignation()

			localctx.(*InstructionContext)._asignation = _x
		}
		p.SetState(111)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserPTOCOMA {
			{
				p.SetState(110)
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
			p.SetState(115)

			var _x = p.Unarysum()

			localctx.(*InstructionContext)._unarysum = _x
		}
		p.SetState(117)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserPTOCOMA {
			{
				p.SetState(116)
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
			p.SetState(121)

			var _x = p.Unarysub()

			localctx.(*InstructionContext)._unarysub = _x
		}
		p.SetState(123)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserPTOCOMA {
			{
				p.SetState(122)
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
			p.SetState(127)

			var _x = p.Breakstatement()

			localctx.(*InstructionContext)._breakstatement = _x
		}
		p.SetState(129)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserPTOCOMA {
			{
				p.SetState(128)
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
			p.SetState(133)

			var _x = p.Continuestatement()

			localctx.(*InstructionContext)._continuestatement = _x
		}
		p.SetState(135)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserPTOCOMA {
			{
				p.SetState(134)
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
			p.SetState(139)

			var _x = p.Ifstmt()

			localctx.(*InstructionContext)._ifstmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_ifstmt().GetNewif()

	case 14:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(142)

			var _x = p.While_statement()

			localctx.(*InstructionContext)._while_statement = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_while_statement().GetNewwhile()

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
	p.EnterRule(localctx, 10, SwiftGrammarParserRULE_removeatvec)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(147)

		var _m = p.Match(SwiftGrammarParserID)

		localctx.(*RemoveatvecContext)._ID = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(148)
		p.Match(SwiftGrammarParserPTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(149)
		p.Match(SwiftGrammarParserRREMOVEAT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(150)
		p.Match(SwiftGrammarParserPARIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(151)
		p.Match(SwiftGrammarParserRRAT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(152)
		p.Match(SwiftGrammarParserDOSPTOS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(153)

		var _x = p.expr(0)

		localctx.(*RemoveatvecContext)._expr = _x
	}
	{
		p.SetState(154)
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
	p.EnterRule(localctx, 12, SwiftGrammarParserRULE_appendvec)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(157)

		var _m = p.Match(SwiftGrammarParserID)

		localctx.(*AppendvecContext)._ID = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(158)
		p.Match(SwiftGrammarParserPTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(159)
		p.Match(SwiftGrammarParserRAPPEND)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(160)
		p.Match(SwiftGrammarParserPARIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(161)

		var _x = p.expr(0)

		localctx.(*AppendvecContext)._expr = _x
	}
	{
		p.SetState(162)
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
	p.EnterRule(localctx, 14, SwiftGrammarParserRULE_removelastvec)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(165)

		var _m = p.Match(SwiftGrammarParserID)

		localctx.(*RemovelastvecContext)._ID = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(166)
		p.Match(SwiftGrammarParserPTO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(167)
		p.Match(SwiftGrammarParserRREMOVELAST)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(168)
		p.Match(SwiftGrammarParserPARIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(169)
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
	Expr() IExprContext
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
	_expr     IExprContext
	typpe2    antlr.Token
	firstid   antlr.Token
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
	p.EnterRule(localctx, 16, SwiftGrammarParserRULE_vecdec)
	var _la int

	p.SetState(207)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(172)

			var _m = p.Match(SwiftGrammarParserRVAR)

			localctx.(*VecdecContext)._RVAR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(173)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*VecdecContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(174)
			p.Match(SwiftGrammarParserDOSPTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(175)
			p.Match(SwiftGrammarParserOBRA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(176)

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
			p.SetState(177)
			p.Match(SwiftGrammarParserCBRA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(178)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(179)
			p.Match(SwiftGrammarParserOBRA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(180)

			var _x = p.expr(0)

			localctx.(*VecdecContext)._expr = _x
		}
		{
			p.SetState(181)
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
		}()), nil, localctx.(*VecdecContext).Get_expr().GetE())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(184)

			var _m = p.Match(SwiftGrammarParserRVAR)

			localctx.(*VecdecContext)._RVAR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(185)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*VecdecContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(186)
			p.Match(SwiftGrammarParserDOSPTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(187)
			p.Match(SwiftGrammarParserOBRA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(188)

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
			p.SetState(189)
			p.Match(SwiftGrammarParserCBRA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(190)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(191)
			p.Match(SwiftGrammarParserOBRA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(192)

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
			p.SetState(193)
			p.Match(SwiftGrammarParserCBRA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(194)
			p.Match(SwiftGrammarParserOBRA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(195)
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

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(197)

			var _m = p.Match(SwiftGrammarParserRVAR)

			localctx.(*VecdecContext)._RVAR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(198)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*VecdecContext).firstid = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(199)
			p.Match(SwiftGrammarParserDOSPTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(200)
			p.Match(SwiftGrammarParserOBRA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(201)

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
			p.SetState(202)
			p.Match(SwiftGrammarParserCBRA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(203)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(204)

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
	p.EnterRule(localctx, 18, SwiftGrammarParserRULE_breakstatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(209)

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
	p.EnterRule(localctx, 20, SwiftGrammarParserRULE_continuestatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(212)

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

	// Get_elseifstatement returns the _elseifstatement rule contexts.
	Get_elseifstatement() IElseifstatementContext

	// SetEx sets the ex rule contexts.
	SetEx(IExprContext)

	// SetB sets the b rule contexts.
	SetB(IBlockContext)

	// Set_elseifstatement sets the _elseifstatement rule contexts.
	Set_elseifstatement(IElseifstatementContext)

	// GetNewif returns the newif attribute.
	GetNewif() interfaces.Instruction

	// SetNewif sets the newif attribute.
	SetNewif(interfaces.Instruction)

	// Getter signatures
	RIF() antlr.TerminalNode
	LLAVEIZQ() antlr.TerminalNode
	LLAVEDER() antlr.TerminalNode
	Elseifstatement() IElseifstatementContext
	Expr() IExprContext
	Block() IBlockContext

	// IsIfstmtContext differentiates from other interfaces.
	IsIfstmtContext()
}

type IfstmtContext struct {
	antlr.BaseParserRuleContext
	parser           antlr.Parser
	newif            interfaces.Instruction
	_RIF             antlr.Token
	ex               IExprContext
	b                IBlockContext
	_elseifstatement IElseifstatementContext
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

func (s *IfstmtContext) Get_elseifstatement() IElseifstatementContext { return s._elseifstatement }

func (s *IfstmtContext) SetEx(v IExprContext) { s.ex = v }

func (s *IfstmtContext) SetB(v IBlockContext) { s.b = v }

func (s *IfstmtContext) Set_elseifstatement(v IElseifstatementContext) { s._elseifstatement = v }

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

func (s *IfstmtContext) Elseifstatement() IElseifstatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElseifstatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElseifstatementContext)
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
	p.EnterRule(localctx, 22, SwiftGrammarParserRULE_ifstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(215)

		var _m = p.Match(SwiftGrammarParserRIF)

		localctx.(*IfstmtContext)._RIF = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(216)

		var _x = p.expr(0)

		localctx.(*IfstmtContext).ex = _x
	}
	{
		p.SetState(217)
		p.Match(SwiftGrammarParserLLAVEIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(218)

		var _x = p.Block()

		localctx.(*IfstmtContext).b = _x
	}
	{
		p.SetState(219)
		p.Match(SwiftGrammarParserLLAVEDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(220)

		var _x = p.Elseifstatement()

		localctx.(*IfstmtContext)._elseifstatement = _x
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
	}()), localctx.(*IfstmtContext).GetEx().GetE(), localctx.(*IfstmtContext).GetB().GetBlk(), localctx.(*IfstmtContext).Get_elseifstatement().GetNewelse())

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

// IElseifstatementContext is an interface to support dynamic dispatch.
type IElseifstatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ifstmt returns the _ifstmt rule contexts.
	Get_ifstmt() IIfstmtContext

	// GetB returns the b rule contexts.
	GetB() IBlockContext

	// Set_ifstmt sets the _ifstmt rule contexts.
	Set_ifstmt(IIfstmtContext)

	// SetB sets the b rule contexts.
	SetB(IBlockContext)

	// GetNewelse returns the newelse attribute.
	GetNewelse() []interface{}

	// SetNewelse sets the newelse attribute.
	SetNewelse([]interface{})

	// Getter signatures
	RELSE() antlr.TerminalNode
	Ifstmt() IIfstmtContext
	LLAVEIZQ() antlr.TerminalNode
	LLAVEDER() antlr.TerminalNode
	Block() IBlockContext

	// IsElseifstatementContext differentiates from other interfaces.
	IsElseifstatementContext()
}

type ElseifstatementContext struct {
	antlr.BaseParserRuleContext
	parser  antlr.Parser
	newelse []interface{}
	_ifstmt IIfstmtContext
	b       IBlockContext
}

func NewEmptyElseifstatementContext() *ElseifstatementContext {
	var p = new(ElseifstatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_elseifstatement
	return p
}

func InitEmptyElseifstatementContext(p *ElseifstatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_elseifstatement
}

func (*ElseifstatementContext) IsElseifstatementContext() {}

func NewElseifstatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElseifstatementContext {
	var p = new(ElseifstatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_elseifstatement

	return p
}

func (s *ElseifstatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ElseifstatementContext) Get_ifstmt() IIfstmtContext { return s._ifstmt }

func (s *ElseifstatementContext) GetB() IBlockContext { return s.b }

func (s *ElseifstatementContext) Set_ifstmt(v IIfstmtContext) { s._ifstmt = v }

func (s *ElseifstatementContext) SetB(v IBlockContext) { s.b = v }

func (s *ElseifstatementContext) GetNewelse() []interface{} { return s.newelse }

func (s *ElseifstatementContext) SetNewelse(v []interface{}) { s.newelse = v }

func (s *ElseifstatementContext) RELSE() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserRELSE, 0)
}

func (s *ElseifstatementContext) Ifstmt() IIfstmtContext {
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

func (s *ElseifstatementContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEIZQ, 0)
}

func (s *ElseifstatementContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEDER, 0)
}

func (s *ElseifstatementContext) Block() IBlockContext {
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

func (s *ElseifstatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseifstatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElseifstatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterElseifstatement(s)
	}
}

func (s *ElseifstatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitElseifstatement(s)
	}
}

func (p *SwiftGrammarParser) Elseifstatement() (localctx IElseifstatementContext) {
	localctx = NewElseifstatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, SwiftGrammarParserRULE_elseifstatement)
	p.SetState(234)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(223)
			p.Match(SwiftGrammarParserRELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(224)

			var _x = p.Ifstmt()

			localctx.(*ElseifstatementContext)._ifstmt = _x
		}
		localctx.(*ElseifstatementContext).newelse = append(localctx.(*ElseifstatementContext).newelse, localctx.(*ElseifstatementContext).Get_ifstmt().GetNewif())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(227)
			p.Match(SwiftGrammarParserRELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(228)
			p.Match(SwiftGrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(229)

			var _x = p.Block()

			localctx.(*ElseifstatementContext).b = _x
		}
		{
			p.SetState(230)
			p.Match(SwiftGrammarParserLLAVEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ElseifstatementContext).newelse = append(localctx.(*ElseifstatementContext).newelse, localctx.(*ElseifstatementContext).GetB().GetBlk())

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
	p.EnterRule(localctx, 26, SwiftGrammarParserRULE_printstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(236)

		var _m = p.Match(SwiftGrammarParserRPRINT)

		localctx.(*PrintstmtContext)._RPRINT = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(237)
		p.Match(SwiftGrammarParserPARIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(238)

		var _x = p.Arguments()

		localctx.(*PrintstmtContext)._arguments = _x
	}
	{
		p.SetState(239)
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
	p.EnterRule(localctx, 28, SwiftGrammarParserRULE_while_statement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(242)

		var _m = p.Match(SwiftGrammarParserRWHILE)

		localctx.(*While_statementContext)._RWHILE = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(243)

		var _x = p.expr(0)

		localctx.(*While_statementContext)._expr = _x
	}
	{
		p.SetState(244)
		p.Match(SwiftGrammarParserLLAVEIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(245)

		var _x = p.Block()

		localctx.(*While_statementContext).b = _x
	}
	{
		p.SetState(246)
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
	p.EnterRule(localctx, 30, SwiftGrammarParserRULE_vardec)
	var _la int

	p.SetState(269)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(249)

			var _m = p.Match(SwiftGrammarParserRVAR)

			localctx.(*VardecContext)._RVAR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(250)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*VardecContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(251)
			p.Match(SwiftGrammarParserDOSPTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(252)

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
			p.SetState(253)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(254)

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
			p.SetState(257)

			var _m = p.Match(SwiftGrammarParserRVAR)

			localctx.(*VardecContext)._RVAR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(258)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*VardecContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(259)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(260)

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
			p.SetState(263)

			var _m = p.Match(SwiftGrammarParserRVAR)

			localctx.(*VardecContext)._RVAR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(264)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*VardecContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(265)
			p.Match(SwiftGrammarParserDOSPTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(266)

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
			p.SetState(267)
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
	p.EnterRule(localctx, 32, SwiftGrammarParserRULE_constdec)
	var _la int

	p.SetState(285)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(271)

			var _m = p.Match(SwiftGrammarParserRLET)

			localctx.(*ConstdecContext)._RLET = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(272)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*ConstdecContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(273)
			p.Match(SwiftGrammarParserDOSPTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(274)

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
			p.SetState(275)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(276)

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
			p.SetState(279)

			var _m = p.Match(SwiftGrammarParserRLET)

			localctx.(*ConstdecContext)._RLET = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(280)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*ConstdecContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(281)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(282)

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
	p.EnterRule(localctx, 34, SwiftGrammarParserRULE_asignation)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(287)

		var _m = p.Match(SwiftGrammarParserID)

		localctx.(*AsignationContext)._ID = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(288)
		p.Match(SwiftGrammarParserIG)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(289)

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
	p.EnterRule(localctx, 36, SwiftGrammarParserRULE_unarysum)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(292)

		var _m = p.Match(SwiftGrammarParserID)

		localctx.(*UnarysumContext)._ID = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(293)
		p.Match(SwiftGrammarParserUNARYPLUS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(294)

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
	p.EnterRule(localctx, 38, SwiftGrammarParserRULE_unarysub)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(297)

		var _m = p.Match(SwiftGrammarParserID)

		localctx.(*UnarysubContext)._ID = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(298)
		p.Match(SwiftGrammarParserUNARYMINUS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(299)

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

	// Get_arguments returns the _arguments rule contexts.
	Get_arguments() IArgumentsContext

	// GetRight returns the right rule contexts.
	GetRight() IExprContext

	// SetLeft sets the left rule contexts.
	SetLeft(IExprContext)

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

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
	parser     antlr.Parser
	e          interfaces.Expression
	left       IExprContext
	_expr      IExprContext
	op         antlr.Token
	_OBRA      antlr.Token
	_arguments IArgumentsContext
	_ID        antlr.Token
	_NUMBER    antlr.Token
	_STRING    antlr.Token
	_RTRUE     antlr.Token
	_RFALSE    antlr.Token
	right      IExprContext
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

func (s *ExprContext) Get_arguments() IArgumentsContext { return s._arguments }

func (s *ExprContext) GetRight() IExprContext { return s.right }

func (s *ExprContext) SetLeft(v IExprContext) { s.left = v }

func (s *ExprContext) Set_expr(v IExprContext) { s._expr = v }

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
	_startState := 40
	p.EnterRecursionRule(localctx, 40, SwiftGrammarParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(329)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SwiftGrammarParserPARIZQ:
		{
			p.SetState(303)
			p.Match(SwiftGrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(304)

			var _x = p.expr(0)

			localctx.(*ExprContext)._expr = _x
		}
		{
			p.SetState(305)
			p.Match(SwiftGrammarParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ExprContext).e = localctx.(*ExprContext).Get_expr().GetE()

	case SwiftGrammarParserNOT, SwiftGrammarParserSUB:
		{
			p.SetState(308)

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
			p.SetState(309)

			var _x = p.expr(14)

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

	case SwiftGrammarParserOBRA:
		{
			p.SetState(312)

			var _m = p.Match(SwiftGrammarParserOBRA)

			localctx.(*ExprContext)._OBRA = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(313)

			var _x = p.Arguments()

			localctx.(*ExprContext)._arguments = _x
		}
		{
			p.SetState(314)
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

	case SwiftGrammarParserID:
		{
			p.SetState(317)

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

	case SwiftGrammarParserNUMBER:
		{
			p.SetState(319)

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

	case SwiftGrammarParserSTRING:
		{
			p.SetState(321)

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

	case SwiftGrammarParserRTRUE:
		{
			p.SetState(323)

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

	case SwiftGrammarParserRFALSE:
		{
			p.SetState(325)

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

	case SwiftGrammarParserRNIL:
		{
			p.SetState(327)
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

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(363)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(361)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_expr)
				p.SetState(331)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
					goto errorExit
				}
				{
					p.SetState(332)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&281681135140864) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(333)

					var _x = p.expr(14)

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
				p.SetState(336)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
					goto errorExit
				}
				{
					p.SetState(337)

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
					p.SetState(338)

					var _x = p.expr(13)

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
				p.SetState(341)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
					goto errorExit
				}
				{
					p.SetState(342)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&64424509440) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(343)

					var _x = p.expr(12)

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
				p.SetState(346)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(347)

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
					p.SetState(348)

					var _x = p.expr(11)

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
				p.SetState(351)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(352)

					var _m = p.Match(SwiftGrammarParserAND)

					localctx.(*ExprContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(353)

					var _x = p.expr(10)

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
				p.SetState(356)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(357)

					var _m = p.Match(SwiftGrammarParserOR)

					localctx.(*ExprContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(358)

					var _x = p.expr(9)

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
		p.SetState(365)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext())
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
	case 20:
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
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 8)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
