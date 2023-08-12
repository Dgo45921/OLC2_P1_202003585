// Code generated from SwiftLexer.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser

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

type SwiftLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var SwiftLexerLexerStaticData struct {
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

func swiftlexerLexerInit() {
	staticData := &SwiftLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'Int'", "'Float'", "'Bool'", "'String'", "'Character'", "'true'",
		"'false'", "'print'", "'if'", "'else'", "'while'", "'var'", "'let'",
		"", "", "", "'!='", "'=='", "'!'", "'||'", "'&&'", "'='", "'>='", "'<='",
		"'>'", "'<'", "'*'", "'/'", "'+'", "'-'", "'('", "')'", "'{'", "'}'",
		"';'", "':'", "'%'", "','", "'?'",
	}
	staticData.SymbolicNames = []string{
		"", "RINT", "RFLOAT", "RBOOL", "RSTRING", "RCHARACTER", "RTRUE", "RFALSE",
		"RPRINT", "RIF", "RELSE", "RWHILE", "RVAR", "RLET", "NUMBER", "STRING",
		"ID", "DIF", "IG_IG", "NOT", "OR", "AND", "IG", "MAY_IG", "MEN_IG",
		"MAYOR", "MENOR", "MUL", "DIV", "ADD", "SUB", "PARIZQ", "PARDER", "LLAVEIZQ",
		"LLAVEDER", "PTOCOMA", "DOSPTOS", "MODULE", "COMA", "QM", "WHITESPACE",
		"COMMENT", "LINE_COMMENT",
	}
	staticData.RuleNames = []string{
		"RINT", "RFLOAT", "RBOOL", "RSTRING", "RCHARACTER", "RTRUE", "RFALSE",
		"RPRINT", "RIF", "RELSE", "RWHILE", "RVAR", "RLET", "NUMBER", "STRING",
		"ID", "DIF", "IG_IG", "NOT", "OR", "AND", "IG", "MAY_IG", "MEN_IG",
		"MAYOR", "MENOR", "MUL", "DIV", "ADD", "SUB", "PARIZQ", "PARDER", "LLAVEIZQ",
		"LLAVEDER", "PTOCOMA", "DOSPTOS", "MODULE", "COMA", "QM", "WHITESPACE",
		"COMMENT", "LINE_COMMENT",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 42, 269, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1,
		7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1,
		12, 1, 12, 1, 13, 4, 13, 158, 8, 13, 11, 13, 12, 13, 159, 1, 13, 1, 13,
		4, 13, 164, 8, 13, 11, 13, 12, 13, 165, 3, 13, 168, 8, 13, 1, 14, 1, 14,
		5, 14, 172, 8, 14, 10, 14, 12, 14, 175, 9, 14, 1, 14, 1, 14, 1, 15, 1,
		15, 5, 15, 181, 8, 15, 10, 15, 12, 15, 184, 9, 15, 1, 16, 1, 16, 1, 16,
		1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1,
		20, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24,
		1, 25, 1, 25, 1, 26, 1, 26, 1, 27, 1, 27, 1, 28, 1, 28, 1, 29, 1, 29, 1,
		30, 1, 30, 1, 31, 1, 31, 1, 32, 1, 32, 1, 33, 1, 33, 1, 34, 1, 34, 1, 35,
		1, 35, 1, 36, 1, 36, 1, 37, 1, 37, 1, 38, 1, 38, 1, 39, 4, 39, 239, 8,
		39, 11, 39, 12, 39, 240, 1, 39, 1, 39, 1, 40, 1, 40, 1, 40, 1, 40, 5, 40,
		249, 8, 40, 10, 40, 12, 40, 252, 9, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1,
		40, 1, 41, 1, 41, 1, 41, 1, 41, 5, 41, 263, 8, 41, 10, 41, 12, 41, 266,
		9, 41, 1, 41, 1, 41, 1, 250, 0, 42, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6,
		13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31,
		16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49,
		25, 51, 26, 53, 27, 55, 28, 57, 29, 59, 30, 61, 31, 63, 32, 65, 33, 67,
		34, 69, 35, 71, 36, 73, 37, 75, 38, 77, 39, 79, 40, 81, 41, 83, 42, 1,
		0, 6, 1, 0, 48, 57, 1, 0, 34, 34, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0,
		48, 57, 65, 90, 95, 95, 97, 122, 4, 0, 9, 10, 13, 13, 32, 32, 92, 92, 2,
		0, 10, 10, 13, 13, 276, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0,
		0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1,
		0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21,
		1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0,
		29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0,
		0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0,
		0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0,
		0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1,
		0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67,
		1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0, 0,
		75, 1, 0, 0, 0, 0, 77, 1, 0, 0, 0, 0, 79, 1, 0, 0, 0, 0, 81, 1, 0, 0, 0,
		0, 83, 1, 0, 0, 0, 1, 85, 1, 0, 0, 0, 3, 89, 1, 0, 0, 0, 5, 95, 1, 0, 0,
		0, 7, 100, 1, 0, 0, 0, 9, 107, 1, 0, 0, 0, 11, 117, 1, 0, 0, 0, 13, 122,
		1, 0, 0, 0, 15, 128, 1, 0, 0, 0, 17, 134, 1, 0, 0, 0, 19, 137, 1, 0, 0,
		0, 21, 142, 1, 0, 0, 0, 23, 148, 1, 0, 0, 0, 25, 152, 1, 0, 0, 0, 27, 157,
		1, 0, 0, 0, 29, 169, 1, 0, 0, 0, 31, 178, 1, 0, 0, 0, 33, 185, 1, 0, 0,
		0, 35, 188, 1, 0, 0, 0, 37, 191, 1, 0, 0, 0, 39, 193, 1, 0, 0, 0, 41, 196,
		1, 0, 0, 0, 43, 199, 1, 0, 0, 0, 45, 201, 1, 0, 0, 0, 47, 204, 1, 0, 0,
		0, 49, 207, 1, 0, 0, 0, 51, 209, 1, 0, 0, 0, 53, 211, 1, 0, 0, 0, 55, 213,
		1, 0, 0, 0, 57, 215, 1, 0, 0, 0, 59, 217, 1, 0, 0, 0, 61, 219, 1, 0, 0,
		0, 63, 221, 1, 0, 0, 0, 65, 223, 1, 0, 0, 0, 67, 225, 1, 0, 0, 0, 69, 227,
		1, 0, 0, 0, 71, 229, 1, 0, 0, 0, 73, 231, 1, 0, 0, 0, 75, 233, 1, 0, 0,
		0, 77, 235, 1, 0, 0, 0, 79, 238, 1, 0, 0, 0, 81, 244, 1, 0, 0, 0, 83, 258,
		1, 0, 0, 0, 85, 86, 5, 73, 0, 0, 86, 87, 5, 110, 0, 0, 87, 88, 5, 116,
		0, 0, 88, 2, 1, 0, 0, 0, 89, 90, 5, 70, 0, 0, 90, 91, 5, 108, 0, 0, 91,
		92, 5, 111, 0, 0, 92, 93, 5, 97, 0, 0, 93, 94, 5, 116, 0, 0, 94, 4, 1,
		0, 0, 0, 95, 96, 5, 66, 0, 0, 96, 97, 5, 111, 0, 0, 97, 98, 5, 111, 0,
		0, 98, 99, 5, 108, 0, 0, 99, 6, 1, 0, 0, 0, 100, 101, 5, 83, 0, 0, 101,
		102, 5, 116, 0, 0, 102, 103, 5, 114, 0, 0, 103, 104, 5, 105, 0, 0, 104,
		105, 5, 110, 0, 0, 105, 106, 5, 103, 0, 0, 106, 8, 1, 0, 0, 0, 107, 108,
		5, 67, 0, 0, 108, 109, 5, 104, 0, 0, 109, 110, 5, 97, 0, 0, 110, 111, 5,
		114, 0, 0, 111, 112, 5, 97, 0, 0, 112, 113, 5, 99, 0, 0, 113, 114, 5, 116,
		0, 0, 114, 115, 5, 101, 0, 0, 115, 116, 5, 114, 0, 0, 116, 10, 1, 0, 0,
		0, 117, 118, 5, 116, 0, 0, 118, 119, 5, 114, 0, 0, 119, 120, 5, 117, 0,
		0, 120, 121, 5, 101, 0, 0, 121, 12, 1, 0, 0, 0, 122, 123, 5, 102, 0, 0,
		123, 124, 5, 97, 0, 0, 124, 125, 5, 108, 0, 0, 125, 126, 5, 115, 0, 0,
		126, 127, 5, 101, 0, 0, 127, 14, 1, 0, 0, 0, 128, 129, 5, 112, 0, 0, 129,
		130, 5, 114, 0, 0, 130, 131, 5, 105, 0, 0, 131, 132, 5, 110, 0, 0, 132,
		133, 5, 116, 0, 0, 133, 16, 1, 0, 0, 0, 134, 135, 5, 105, 0, 0, 135, 136,
		5, 102, 0, 0, 136, 18, 1, 0, 0, 0, 137, 138, 5, 101, 0, 0, 138, 139, 5,
		108, 0, 0, 139, 140, 5, 115, 0, 0, 140, 141, 5, 101, 0, 0, 141, 20, 1,
		0, 0, 0, 142, 143, 5, 119, 0, 0, 143, 144, 5, 104, 0, 0, 144, 145, 5, 105,
		0, 0, 145, 146, 5, 108, 0, 0, 146, 147, 5, 101, 0, 0, 147, 22, 1, 0, 0,
		0, 148, 149, 5, 118, 0, 0, 149, 150, 5, 97, 0, 0, 150, 151, 5, 114, 0,
		0, 151, 24, 1, 0, 0, 0, 152, 153, 5, 108, 0, 0, 153, 154, 5, 101, 0, 0,
		154, 155, 5, 116, 0, 0, 155, 26, 1, 0, 0, 0, 156, 158, 7, 0, 0, 0, 157,
		156, 1, 0, 0, 0, 158, 159, 1, 0, 0, 0, 159, 157, 1, 0, 0, 0, 159, 160,
		1, 0, 0, 0, 160, 167, 1, 0, 0, 0, 161, 163, 5, 46, 0, 0, 162, 164, 7, 0,
		0, 0, 163, 162, 1, 0, 0, 0, 164, 165, 1, 0, 0, 0, 165, 163, 1, 0, 0, 0,
		165, 166, 1, 0, 0, 0, 166, 168, 1, 0, 0, 0, 167, 161, 1, 0, 0, 0, 167,
		168, 1, 0, 0, 0, 168, 28, 1, 0, 0, 0, 169, 173, 5, 34, 0, 0, 170, 172,
		8, 1, 0, 0, 171, 170, 1, 0, 0, 0, 172, 175, 1, 0, 0, 0, 173, 171, 1, 0,
		0, 0, 173, 174, 1, 0, 0, 0, 174, 176, 1, 0, 0, 0, 175, 173, 1, 0, 0, 0,
		176, 177, 5, 34, 0, 0, 177, 30, 1, 0, 0, 0, 178, 182, 7, 2, 0, 0, 179,
		181, 7, 3, 0, 0, 180, 179, 1, 0, 0, 0, 181, 184, 1, 0, 0, 0, 182, 180,
		1, 0, 0, 0, 182, 183, 1, 0, 0, 0, 183, 32, 1, 0, 0, 0, 184, 182, 1, 0,
		0, 0, 185, 186, 5, 33, 0, 0, 186, 187, 5, 61, 0, 0, 187, 34, 1, 0, 0, 0,
		188, 189, 5, 61, 0, 0, 189, 190, 5, 61, 0, 0, 190, 36, 1, 0, 0, 0, 191,
		192, 5, 33, 0, 0, 192, 38, 1, 0, 0, 0, 193, 194, 5, 124, 0, 0, 194, 195,
		5, 124, 0, 0, 195, 40, 1, 0, 0, 0, 196, 197, 5, 38, 0, 0, 197, 198, 5,
		38, 0, 0, 198, 42, 1, 0, 0, 0, 199, 200, 5, 61, 0, 0, 200, 44, 1, 0, 0,
		0, 201, 202, 5, 62, 0, 0, 202, 203, 5, 61, 0, 0, 203, 46, 1, 0, 0, 0, 204,
		205, 5, 60, 0, 0, 205, 206, 5, 61, 0, 0, 206, 48, 1, 0, 0, 0, 207, 208,
		5, 62, 0, 0, 208, 50, 1, 0, 0, 0, 209, 210, 5, 60, 0, 0, 210, 52, 1, 0,
		0, 0, 211, 212, 5, 42, 0, 0, 212, 54, 1, 0, 0, 0, 213, 214, 5, 47, 0, 0,
		214, 56, 1, 0, 0, 0, 215, 216, 5, 43, 0, 0, 216, 58, 1, 0, 0, 0, 217, 218,
		5, 45, 0, 0, 218, 60, 1, 0, 0, 0, 219, 220, 5, 40, 0, 0, 220, 62, 1, 0,
		0, 0, 221, 222, 5, 41, 0, 0, 222, 64, 1, 0, 0, 0, 223, 224, 5, 123, 0,
		0, 224, 66, 1, 0, 0, 0, 225, 226, 5, 125, 0, 0, 226, 68, 1, 0, 0, 0, 227,
		228, 5, 59, 0, 0, 228, 70, 1, 0, 0, 0, 229, 230, 5, 58, 0, 0, 230, 72,
		1, 0, 0, 0, 231, 232, 5, 37, 0, 0, 232, 74, 1, 0, 0, 0, 233, 234, 5, 44,
		0, 0, 234, 76, 1, 0, 0, 0, 235, 236, 5, 63, 0, 0, 236, 78, 1, 0, 0, 0,
		237, 239, 7, 4, 0, 0, 238, 237, 1, 0, 0, 0, 239, 240, 1, 0, 0, 0, 240,
		238, 1, 0, 0, 0, 240, 241, 1, 0, 0, 0, 241, 242, 1, 0, 0, 0, 242, 243,
		6, 39, 0, 0, 243, 80, 1, 0, 0, 0, 244, 245, 5, 47, 0, 0, 245, 246, 5, 42,
		0, 0, 246, 250, 1, 0, 0, 0, 247, 249, 9, 0, 0, 0, 248, 247, 1, 0, 0, 0,
		249, 252, 1, 0, 0, 0, 250, 251, 1, 0, 0, 0, 250, 248, 1, 0, 0, 0, 251,
		253, 1, 0, 0, 0, 252, 250, 1, 0, 0, 0, 253, 254, 5, 42, 0, 0, 254, 255,
		5, 47, 0, 0, 255, 256, 1, 0, 0, 0, 256, 257, 6, 40, 0, 0, 257, 82, 1, 0,
		0, 0, 258, 259, 5, 47, 0, 0, 259, 260, 5, 47, 0, 0, 260, 264, 1, 0, 0,
		0, 261, 263, 8, 5, 0, 0, 262, 261, 1, 0, 0, 0, 263, 266, 1, 0, 0, 0, 264,
		262, 1, 0, 0, 0, 264, 265, 1, 0, 0, 0, 265, 267, 1, 0, 0, 0, 266, 264,
		1, 0, 0, 0, 267, 268, 6, 41, 0, 0, 268, 84, 1, 0, 0, 0, 9, 0, 159, 165,
		167, 173, 182, 240, 250, 264, 1, 6, 0, 0,
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

// SwiftLexerInit initializes any static state used to implement SwiftLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewSwiftLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func SwiftLexerInit() {
	staticData := &SwiftLexerLexerStaticData
	staticData.once.Do(swiftlexerLexerInit)
}

// NewSwiftLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewSwiftLexer(input antlr.CharStream) *SwiftLexer {
	SwiftLexerInit()
	l := new(SwiftLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &SwiftLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "SwiftLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// SwiftLexer tokens.
const (
	SwiftLexerRINT         = 1
	SwiftLexerRFLOAT       = 2
	SwiftLexerRBOOL        = 3
	SwiftLexerRSTRING      = 4
	SwiftLexerRCHARACTER   = 5
	SwiftLexerRTRUE        = 6
	SwiftLexerRFALSE       = 7
	SwiftLexerRPRINT       = 8
	SwiftLexerRIF          = 9
	SwiftLexerRELSE        = 10
	SwiftLexerRWHILE       = 11
	SwiftLexerRVAR         = 12
	SwiftLexerRLET         = 13
	SwiftLexerNUMBER       = 14
	SwiftLexerSTRING       = 15
	SwiftLexerID           = 16
	SwiftLexerDIF          = 17
	SwiftLexerIG_IG        = 18
	SwiftLexerNOT          = 19
	SwiftLexerOR           = 20
	SwiftLexerAND          = 21
	SwiftLexerIG           = 22
	SwiftLexerMAY_IG       = 23
	SwiftLexerMEN_IG       = 24
	SwiftLexerMAYOR        = 25
	SwiftLexerMENOR        = 26
	SwiftLexerMUL          = 27
	SwiftLexerDIV          = 28
	SwiftLexerADD          = 29
	SwiftLexerSUB          = 30
	SwiftLexerPARIZQ       = 31
	SwiftLexerPARDER       = 32
	SwiftLexerLLAVEIZQ     = 33
	SwiftLexerLLAVEDER     = 34
	SwiftLexerPTOCOMA      = 35
	SwiftLexerDOSPTOS      = 36
	SwiftLexerMODULE       = 37
	SwiftLexerCOMA         = 38
	SwiftLexerQM           = 39
	SwiftLexerWHITESPACE   = 40
	SwiftLexerCOMMENT      = 41
	SwiftLexerLINE_COMMENT = 42
)
