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
		"'false'", "'print'", "'if'", "'else'", "'while'", "'var'", "", "",
		"", "'!='", "'=='", "'!'", "'||'", "'&&'", "'='", "'>='", "'<='", "'>'",
		"'<'", "'*'", "'/'", "'+'", "'-'", "'('", "')'", "'{'", "'}'", "';'",
		"':'", "'%'", "','",
	}
	staticData.SymbolicNames = []string{
		"", "RINT", "RFLOAT", "RBOOL", "RSTRING", "RCHARACTER", "RTRUE", "RFALSE",
		"RPRINT", "RIF", "RELSE", "RWHILE", "RVAR", "NUMBER", "STRING", "ID",
		"DIF", "IG_IG", "NOT", "OR", "AND", "IG", "MAY_IG", "MEN_IG", "MAYOR",
		"MENOR", "MUL", "DIV", "ADD", "SUB", "PARIZQ", "PARDER", "LLAVEIZQ",
		"LLAVEDER", "PTOCOMA", "DOSPTOS", "MODULE", "COMA", "WHITESPACE", "COMMENT",
		"LINE_COMMENT",
	}
	staticData.RuleNames = []string{
		"RINT", "RFLOAT", "RBOOL", "RSTRING", "RCHARACTER", "RTRUE", "RFALSE",
		"RPRINT", "RIF", "RELSE", "RWHILE", "RVAR", "NUMBER", "STRING", "ID",
		"DIF", "IG_IG", "NOT", "OR", "AND", "IG", "MAY_IG", "MEN_IG", "MAYOR",
		"MENOR", "MUL", "DIV", "ADD", "SUB", "PARIZQ", "PARDER", "LLAVEIZQ",
		"LLAVEDER", "PTOCOMA", "DOSPTOS", "MODULE", "COMA", "WHITESPACE", "COMMENT",
		"LINE_COMMENT",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 40, 259, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 1, 0, 1, 0, 1, 0, 1, 0,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 4, 12, 150, 8, 12, 11, 12, 12, 12, 151,
		1, 12, 1, 12, 4, 12, 156, 8, 12, 11, 12, 12, 12, 157, 3, 12, 160, 8, 12,
		1, 13, 1, 13, 5, 13, 164, 8, 13, 10, 13, 12, 13, 167, 9, 13, 1, 13, 1,
		13, 1, 14, 1, 14, 5, 14, 173, 8, 14, 10, 14, 12, 14, 176, 9, 14, 1, 15,
		1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1,
		19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22,
		1, 23, 1, 23, 1, 24, 1, 24, 1, 25, 1, 25, 1, 26, 1, 26, 1, 27, 1, 27, 1,
		28, 1, 28, 1, 29, 1, 29, 1, 30, 1, 30, 1, 31, 1, 31, 1, 32, 1, 32, 1, 33,
		1, 33, 1, 34, 1, 34, 1, 35, 1, 35, 1, 36, 1, 36, 1, 37, 4, 37, 229, 8,
		37, 11, 37, 12, 37, 230, 1, 37, 1, 37, 1, 38, 1, 38, 1, 38, 1, 38, 5, 38,
		239, 8, 38, 10, 38, 12, 38, 242, 9, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1,
		38, 1, 39, 1, 39, 1, 39, 1, 39, 5, 39, 253, 8, 39, 10, 39, 12, 39, 256,
		9, 39, 1, 39, 1, 39, 1, 240, 0, 40, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6,
		13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31,
		16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49,
		25, 51, 26, 53, 27, 55, 28, 57, 29, 59, 30, 61, 31, 63, 32, 65, 33, 67,
		34, 69, 35, 71, 36, 73, 37, 75, 38, 77, 39, 79, 40, 1, 0, 6, 1, 0, 48,
		57, 1, 0, 34, 34, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90,
		95, 95, 97, 122, 4, 0, 9, 10, 13, 13, 32, 32, 92, 92, 2, 0, 10, 10, 13,
		13, 266, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1,
		0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15,
		1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0,
		23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0,
		0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0,
		0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0,
		0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1,
		0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61,
		1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0,
		69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0, 0,
		0, 77, 1, 0, 0, 0, 0, 79, 1, 0, 0, 0, 1, 81, 1, 0, 0, 0, 3, 85, 1, 0, 0,
		0, 5, 91, 1, 0, 0, 0, 7, 96, 1, 0, 0, 0, 9, 103, 1, 0, 0, 0, 11, 113, 1,
		0, 0, 0, 13, 118, 1, 0, 0, 0, 15, 124, 1, 0, 0, 0, 17, 130, 1, 0, 0, 0,
		19, 133, 1, 0, 0, 0, 21, 138, 1, 0, 0, 0, 23, 144, 1, 0, 0, 0, 25, 149,
		1, 0, 0, 0, 27, 161, 1, 0, 0, 0, 29, 170, 1, 0, 0, 0, 31, 177, 1, 0, 0,
		0, 33, 180, 1, 0, 0, 0, 35, 183, 1, 0, 0, 0, 37, 185, 1, 0, 0, 0, 39, 188,
		1, 0, 0, 0, 41, 191, 1, 0, 0, 0, 43, 193, 1, 0, 0, 0, 45, 196, 1, 0, 0,
		0, 47, 199, 1, 0, 0, 0, 49, 201, 1, 0, 0, 0, 51, 203, 1, 0, 0, 0, 53, 205,
		1, 0, 0, 0, 55, 207, 1, 0, 0, 0, 57, 209, 1, 0, 0, 0, 59, 211, 1, 0, 0,
		0, 61, 213, 1, 0, 0, 0, 63, 215, 1, 0, 0, 0, 65, 217, 1, 0, 0, 0, 67, 219,
		1, 0, 0, 0, 69, 221, 1, 0, 0, 0, 71, 223, 1, 0, 0, 0, 73, 225, 1, 0, 0,
		0, 75, 228, 1, 0, 0, 0, 77, 234, 1, 0, 0, 0, 79, 248, 1, 0, 0, 0, 81, 82,
		5, 73, 0, 0, 82, 83, 5, 110, 0, 0, 83, 84, 5, 116, 0, 0, 84, 2, 1, 0, 0,
		0, 85, 86, 5, 70, 0, 0, 86, 87, 5, 108, 0, 0, 87, 88, 5, 111, 0, 0, 88,
		89, 5, 97, 0, 0, 89, 90, 5, 116, 0, 0, 90, 4, 1, 0, 0, 0, 91, 92, 5, 66,
		0, 0, 92, 93, 5, 111, 0, 0, 93, 94, 5, 111, 0, 0, 94, 95, 5, 108, 0, 0,
		95, 6, 1, 0, 0, 0, 96, 97, 5, 83, 0, 0, 97, 98, 5, 116, 0, 0, 98, 99, 5,
		114, 0, 0, 99, 100, 5, 105, 0, 0, 100, 101, 5, 110, 0, 0, 101, 102, 5,
		103, 0, 0, 102, 8, 1, 0, 0, 0, 103, 104, 5, 67, 0, 0, 104, 105, 5, 104,
		0, 0, 105, 106, 5, 97, 0, 0, 106, 107, 5, 114, 0, 0, 107, 108, 5, 97, 0,
		0, 108, 109, 5, 99, 0, 0, 109, 110, 5, 116, 0, 0, 110, 111, 5, 101, 0,
		0, 111, 112, 5, 114, 0, 0, 112, 10, 1, 0, 0, 0, 113, 114, 5, 116, 0, 0,
		114, 115, 5, 114, 0, 0, 115, 116, 5, 117, 0, 0, 116, 117, 5, 101, 0, 0,
		117, 12, 1, 0, 0, 0, 118, 119, 5, 102, 0, 0, 119, 120, 5, 97, 0, 0, 120,
		121, 5, 108, 0, 0, 121, 122, 5, 115, 0, 0, 122, 123, 5, 101, 0, 0, 123,
		14, 1, 0, 0, 0, 124, 125, 5, 112, 0, 0, 125, 126, 5, 114, 0, 0, 126, 127,
		5, 105, 0, 0, 127, 128, 5, 110, 0, 0, 128, 129, 5, 116, 0, 0, 129, 16,
		1, 0, 0, 0, 130, 131, 5, 105, 0, 0, 131, 132, 5, 102, 0, 0, 132, 18, 1,
		0, 0, 0, 133, 134, 5, 101, 0, 0, 134, 135, 5, 108, 0, 0, 135, 136, 5, 115,
		0, 0, 136, 137, 5, 101, 0, 0, 137, 20, 1, 0, 0, 0, 138, 139, 5, 119, 0,
		0, 139, 140, 5, 104, 0, 0, 140, 141, 5, 105, 0, 0, 141, 142, 5, 108, 0,
		0, 142, 143, 5, 101, 0, 0, 143, 22, 1, 0, 0, 0, 144, 145, 5, 118, 0, 0,
		145, 146, 5, 97, 0, 0, 146, 147, 5, 114, 0, 0, 147, 24, 1, 0, 0, 0, 148,
		150, 7, 0, 0, 0, 149, 148, 1, 0, 0, 0, 150, 151, 1, 0, 0, 0, 151, 149,
		1, 0, 0, 0, 151, 152, 1, 0, 0, 0, 152, 159, 1, 0, 0, 0, 153, 155, 5, 46,
		0, 0, 154, 156, 7, 0, 0, 0, 155, 154, 1, 0, 0, 0, 156, 157, 1, 0, 0, 0,
		157, 155, 1, 0, 0, 0, 157, 158, 1, 0, 0, 0, 158, 160, 1, 0, 0, 0, 159,
		153, 1, 0, 0, 0, 159, 160, 1, 0, 0, 0, 160, 26, 1, 0, 0, 0, 161, 165, 5,
		34, 0, 0, 162, 164, 8, 1, 0, 0, 163, 162, 1, 0, 0, 0, 164, 167, 1, 0, 0,
		0, 165, 163, 1, 0, 0, 0, 165, 166, 1, 0, 0, 0, 166, 168, 1, 0, 0, 0, 167,
		165, 1, 0, 0, 0, 168, 169, 5, 34, 0, 0, 169, 28, 1, 0, 0, 0, 170, 174,
		7, 2, 0, 0, 171, 173, 7, 3, 0, 0, 172, 171, 1, 0, 0, 0, 173, 176, 1, 0,
		0, 0, 174, 172, 1, 0, 0, 0, 174, 175, 1, 0, 0, 0, 175, 30, 1, 0, 0, 0,
		176, 174, 1, 0, 0, 0, 177, 178, 5, 33, 0, 0, 178, 179, 5, 61, 0, 0, 179,
		32, 1, 0, 0, 0, 180, 181, 5, 61, 0, 0, 181, 182, 5, 61, 0, 0, 182, 34,
		1, 0, 0, 0, 183, 184, 5, 33, 0, 0, 184, 36, 1, 0, 0, 0, 185, 186, 5, 124,
		0, 0, 186, 187, 5, 124, 0, 0, 187, 38, 1, 0, 0, 0, 188, 189, 5, 38, 0,
		0, 189, 190, 5, 38, 0, 0, 190, 40, 1, 0, 0, 0, 191, 192, 5, 61, 0, 0, 192,
		42, 1, 0, 0, 0, 193, 194, 5, 62, 0, 0, 194, 195, 5, 61, 0, 0, 195, 44,
		1, 0, 0, 0, 196, 197, 5, 60, 0, 0, 197, 198, 5, 61, 0, 0, 198, 46, 1, 0,
		0, 0, 199, 200, 5, 62, 0, 0, 200, 48, 1, 0, 0, 0, 201, 202, 5, 60, 0, 0,
		202, 50, 1, 0, 0, 0, 203, 204, 5, 42, 0, 0, 204, 52, 1, 0, 0, 0, 205, 206,
		5, 47, 0, 0, 206, 54, 1, 0, 0, 0, 207, 208, 5, 43, 0, 0, 208, 56, 1, 0,
		0, 0, 209, 210, 5, 45, 0, 0, 210, 58, 1, 0, 0, 0, 211, 212, 5, 40, 0, 0,
		212, 60, 1, 0, 0, 0, 213, 214, 5, 41, 0, 0, 214, 62, 1, 0, 0, 0, 215, 216,
		5, 123, 0, 0, 216, 64, 1, 0, 0, 0, 217, 218, 5, 125, 0, 0, 218, 66, 1,
		0, 0, 0, 219, 220, 5, 59, 0, 0, 220, 68, 1, 0, 0, 0, 221, 222, 5, 58, 0,
		0, 222, 70, 1, 0, 0, 0, 223, 224, 5, 37, 0, 0, 224, 72, 1, 0, 0, 0, 225,
		226, 5, 44, 0, 0, 226, 74, 1, 0, 0, 0, 227, 229, 7, 4, 0, 0, 228, 227,
		1, 0, 0, 0, 229, 230, 1, 0, 0, 0, 230, 228, 1, 0, 0, 0, 230, 231, 1, 0,
		0, 0, 231, 232, 1, 0, 0, 0, 232, 233, 6, 37, 0, 0, 233, 76, 1, 0, 0, 0,
		234, 235, 5, 47, 0, 0, 235, 236, 5, 42, 0, 0, 236, 240, 1, 0, 0, 0, 237,
		239, 9, 0, 0, 0, 238, 237, 1, 0, 0, 0, 239, 242, 1, 0, 0, 0, 240, 241,
		1, 0, 0, 0, 240, 238, 1, 0, 0, 0, 241, 243, 1, 0, 0, 0, 242, 240, 1, 0,
		0, 0, 243, 244, 5, 42, 0, 0, 244, 245, 5, 47, 0, 0, 245, 246, 1, 0, 0,
		0, 246, 247, 6, 38, 0, 0, 247, 78, 1, 0, 0, 0, 248, 249, 5, 47, 0, 0, 249,
		250, 5, 47, 0, 0, 250, 254, 1, 0, 0, 0, 251, 253, 8, 5, 0, 0, 252, 251,
		1, 0, 0, 0, 253, 256, 1, 0, 0, 0, 254, 252, 1, 0, 0, 0, 254, 255, 1, 0,
		0, 0, 255, 257, 1, 0, 0, 0, 256, 254, 1, 0, 0, 0, 257, 258, 6, 39, 0, 0,
		258, 80, 1, 0, 0, 0, 9, 0, 151, 157, 159, 165, 174, 230, 240, 254, 1, 6,
		0, 0,
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
	SwiftLexerNUMBER       = 13
	SwiftLexerSTRING       = 14
	SwiftLexerID           = 15
	SwiftLexerDIF          = 16
	SwiftLexerIG_IG        = 17
	SwiftLexerNOT          = 18
	SwiftLexerOR           = 19
	SwiftLexerAND          = 20
	SwiftLexerIG           = 21
	SwiftLexerMAY_IG       = 22
	SwiftLexerMEN_IG       = 23
	SwiftLexerMAYOR        = 24
	SwiftLexerMENOR        = 25
	SwiftLexerMUL          = 26
	SwiftLexerDIV          = 27
	SwiftLexerADD          = 28
	SwiftLexerSUB          = 29
	SwiftLexerPARIZQ       = 30
	SwiftLexerPARDER       = 31
	SwiftLexerLLAVEIZQ     = 32
	SwiftLexerLLAVEDER     = 33
	SwiftLexerPTOCOMA      = 34
	SwiftLexerDOSPTOS      = 35
	SwiftLexerMODULE       = 36
	SwiftLexerCOMA         = 37
	SwiftLexerWHITESPACE   = 38
	SwiftLexerCOMMENT      = 39
	SwiftLexerLINE_COMMENT = 40
)
