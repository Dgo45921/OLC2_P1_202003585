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
		"", "'int'", "'float'", "'bool'", "'true'", "'false'", "'print'", "'if'",
		"'else'", "'while'", "", "", "", "'!='", "'=='", "'!'", "'||'", "'&&'",
		"'='", "'>='", "'<='", "'>'", "'<'", "'*'", "'/'", "'+'", "'-'", "'('",
		"')'", "'{'", "'}'", "';'", "'%'", "','",
	}
	staticData.SymbolicNames = []string{
		"", "RINT", "RFLOAT", "RBOOL", "RTRUE", "RFALSE", "RPRINT", "RIF", "RELSE",
		"RWHILE", "NUMBER", "STRING", "ID", "DIF", "IG_IG", "NOT", "OR", "AND",
		"IG", "MAY_IG", "MEN_IG", "MAYOR", "MENOR", "MUL", "DIV", "ADD", "SUB",
		"PARIZQ", "PARDER", "LLAVEIZQ", "LLAVEDER", "PTOCOMA", "MODULE", "COMA",
		"WHITESPACE", "COMMENT", "LINE_COMMENT",
	}
	staticData.RuleNames = []string{
		"RINT", "RFLOAT", "RBOOL", "RTRUE", "RFALSE", "RPRINT", "RIF", "RELSE",
		"RWHILE", "NUMBER", "STRING", "ID", "DIF", "IG_IG", "NOT", "OR", "AND",
		"IG", "MAY_IG", "MEN_IG", "MAYOR", "MENOR", "MUL", "DIV", "ADD", "SUB",
		"PARIZQ", "PARDER", "LLAVEIZQ", "LLAVEDER", "PTOCOMA", "MODULE", "COMA",
		"WHITESPACE", "COMMENT", "LINE_COMMENT", "ESC_SEQ",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 36, 233, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9,
		4, 9, 123, 8, 9, 11, 9, 12, 9, 124, 1, 9, 1, 9, 4, 9, 129, 8, 9, 11, 9,
		12, 9, 130, 3, 9, 133, 8, 9, 1, 10, 1, 10, 5, 10, 137, 8, 10, 10, 10, 12,
		10, 140, 9, 10, 1, 10, 1, 10, 1, 11, 1, 11, 5, 11, 146, 8, 11, 10, 11,
		12, 11, 149, 9, 11, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 14, 1,
		14, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18,
		1, 18, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 1, 22, 1, 22, 1,
		23, 1, 23, 1, 24, 1, 24, 1, 25, 1, 25, 1, 26, 1, 26, 1, 27, 1, 27, 1, 28,
		1, 28, 1, 29, 1, 29, 1, 30, 1, 30, 1, 31, 1, 31, 1, 32, 1, 32, 1, 33, 4,
		33, 200, 8, 33, 11, 33, 12, 33, 201, 1, 33, 1, 33, 1, 34, 1, 34, 1, 34,
		1, 34, 5, 34, 210, 8, 34, 10, 34, 12, 34, 213, 9, 34, 1, 34, 1, 34, 1,
		34, 1, 34, 1, 34, 1, 35, 1, 35, 1, 35, 1, 35, 5, 35, 224, 8, 35, 10, 35,
		12, 35, 227, 9, 35, 1, 35, 1, 35, 1, 36, 1, 36, 1, 36, 1, 211, 0, 37, 1,
		1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11,
		23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20,
		41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29,
		59, 30, 61, 31, 63, 32, 65, 33, 67, 34, 69, 35, 71, 36, 73, 0, 1, 0, 7,
		1, 0, 48, 57, 1, 0, 34, 34, 2, 0, 65, 90, 97, 122, 4, 0, 48, 57, 65, 90,
		95, 95, 97, 122, 4, 0, 9, 10, 13, 13, 32, 32, 92, 92, 2, 0, 10, 10, 13,
		13, 7, 0, 32, 33, 35, 35, 43, 43, 45, 46, 58, 58, 64, 64, 91, 93, 239,
		0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0,
		0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0,
		0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0,
		0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1,
		0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39,
		1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0,
		47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0,
		0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0,
		0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0,
		0, 0, 0, 71, 1, 0, 0, 0, 1, 75, 1, 0, 0, 0, 3, 79, 1, 0, 0, 0, 5, 85, 1,
		0, 0, 0, 7, 90, 1, 0, 0, 0, 9, 95, 1, 0, 0, 0, 11, 101, 1, 0, 0, 0, 13,
		107, 1, 0, 0, 0, 15, 110, 1, 0, 0, 0, 17, 115, 1, 0, 0, 0, 19, 122, 1,
		0, 0, 0, 21, 134, 1, 0, 0, 0, 23, 143, 1, 0, 0, 0, 25, 150, 1, 0, 0, 0,
		27, 153, 1, 0, 0, 0, 29, 156, 1, 0, 0, 0, 31, 158, 1, 0, 0, 0, 33, 161,
		1, 0, 0, 0, 35, 164, 1, 0, 0, 0, 37, 166, 1, 0, 0, 0, 39, 169, 1, 0, 0,
		0, 41, 172, 1, 0, 0, 0, 43, 174, 1, 0, 0, 0, 45, 176, 1, 0, 0, 0, 47, 178,
		1, 0, 0, 0, 49, 180, 1, 0, 0, 0, 51, 182, 1, 0, 0, 0, 53, 184, 1, 0, 0,
		0, 55, 186, 1, 0, 0, 0, 57, 188, 1, 0, 0, 0, 59, 190, 1, 0, 0, 0, 61, 192,
		1, 0, 0, 0, 63, 194, 1, 0, 0, 0, 65, 196, 1, 0, 0, 0, 67, 199, 1, 0, 0,
		0, 69, 205, 1, 0, 0, 0, 71, 219, 1, 0, 0, 0, 73, 230, 1, 0, 0, 0, 75, 76,
		5, 105, 0, 0, 76, 77, 5, 110, 0, 0, 77, 78, 5, 116, 0, 0, 78, 2, 1, 0,
		0, 0, 79, 80, 5, 102, 0, 0, 80, 81, 5, 108, 0, 0, 81, 82, 5, 111, 0, 0,
		82, 83, 5, 97, 0, 0, 83, 84, 5, 116, 0, 0, 84, 4, 1, 0, 0, 0, 85, 86, 5,
		98, 0, 0, 86, 87, 5, 111, 0, 0, 87, 88, 5, 111, 0, 0, 88, 89, 5, 108, 0,
		0, 89, 6, 1, 0, 0, 0, 90, 91, 5, 116, 0, 0, 91, 92, 5, 114, 0, 0, 92, 93,
		5, 117, 0, 0, 93, 94, 5, 101, 0, 0, 94, 8, 1, 0, 0, 0, 95, 96, 5, 102,
		0, 0, 96, 97, 5, 97, 0, 0, 97, 98, 5, 108, 0, 0, 98, 99, 5, 115, 0, 0,
		99, 100, 5, 101, 0, 0, 100, 10, 1, 0, 0, 0, 101, 102, 5, 112, 0, 0, 102,
		103, 5, 114, 0, 0, 103, 104, 5, 105, 0, 0, 104, 105, 5, 110, 0, 0, 105,
		106, 5, 116, 0, 0, 106, 12, 1, 0, 0, 0, 107, 108, 5, 105, 0, 0, 108, 109,
		5, 102, 0, 0, 109, 14, 1, 0, 0, 0, 110, 111, 5, 101, 0, 0, 111, 112, 5,
		108, 0, 0, 112, 113, 5, 115, 0, 0, 113, 114, 5, 101, 0, 0, 114, 16, 1,
		0, 0, 0, 115, 116, 5, 119, 0, 0, 116, 117, 5, 104, 0, 0, 117, 118, 5, 105,
		0, 0, 118, 119, 5, 108, 0, 0, 119, 120, 5, 101, 0, 0, 120, 18, 1, 0, 0,
		0, 121, 123, 7, 0, 0, 0, 122, 121, 1, 0, 0, 0, 123, 124, 1, 0, 0, 0, 124,
		122, 1, 0, 0, 0, 124, 125, 1, 0, 0, 0, 125, 132, 1, 0, 0, 0, 126, 128,
		5, 46, 0, 0, 127, 129, 7, 0, 0, 0, 128, 127, 1, 0, 0, 0, 129, 130, 1, 0,
		0, 0, 130, 128, 1, 0, 0, 0, 130, 131, 1, 0, 0, 0, 131, 133, 1, 0, 0, 0,
		132, 126, 1, 0, 0, 0, 132, 133, 1, 0, 0, 0, 133, 20, 1, 0, 0, 0, 134, 138,
		5, 34, 0, 0, 135, 137, 8, 1, 0, 0, 136, 135, 1, 0, 0, 0, 137, 140, 1, 0,
		0, 0, 138, 136, 1, 0, 0, 0, 138, 139, 1, 0, 0, 0, 139, 141, 1, 0, 0, 0,
		140, 138, 1, 0, 0, 0, 141, 142, 5, 34, 0, 0, 142, 22, 1, 0, 0, 0, 143,
		147, 7, 2, 0, 0, 144, 146, 7, 3, 0, 0, 145, 144, 1, 0, 0, 0, 146, 149,
		1, 0, 0, 0, 147, 145, 1, 0, 0, 0, 147, 148, 1, 0, 0, 0, 148, 24, 1, 0,
		0, 0, 149, 147, 1, 0, 0, 0, 150, 151, 5, 33, 0, 0, 151, 152, 5, 61, 0,
		0, 152, 26, 1, 0, 0, 0, 153, 154, 5, 61, 0, 0, 154, 155, 5, 61, 0, 0, 155,
		28, 1, 0, 0, 0, 156, 157, 5, 33, 0, 0, 157, 30, 1, 0, 0, 0, 158, 159, 5,
		124, 0, 0, 159, 160, 5, 124, 0, 0, 160, 32, 1, 0, 0, 0, 161, 162, 5, 38,
		0, 0, 162, 163, 5, 38, 0, 0, 163, 34, 1, 0, 0, 0, 164, 165, 5, 61, 0, 0,
		165, 36, 1, 0, 0, 0, 166, 167, 5, 62, 0, 0, 167, 168, 5, 61, 0, 0, 168,
		38, 1, 0, 0, 0, 169, 170, 5, 60, 0, 0, 170, 171, 5, 61, 0, 0, 171, 40,
		1, 0, 0, 0, 172, 173, 5, 62, 0, 0, 173, 42, 1, 0, 0, 0, 174, 175, 5, 60,
		0, 0, 175, 44, 1, 0, 0, 0, 176, 177, 5, 42, 0, 0, 177, 46, 1, 0, 0, 0,
		178, 179, 5, 47, 0, 0, 179, 48, 1, 0, 0, 0, 180, 181, 5, 43, 0, 0, 181,
		50, 1, 0, 0, 0, 182, 183, 5, 45, 0, 0, 183, 52, 1, 0, 0, 0, 184, 185, 5,
		40, 0, 0, 185, 54, 1, 0, 0, 0, 186, 187, 5, 41, 0, 0, 187, 56, 1, 0, 0,
		0, 188, 189, 5, 123, 0, 0, 189, 58, 1, 0, 0, 0, 190, 191, 5, 125, 0, 0,
		191, 60, 1, 0, 0, 0, 192, 193, 5, 59, 0, 0, 193, 62, 1, 0, 0, 0, 194, 195,
		5, 37, 0, 0, 195, 64, 1, 0, 0, 0, 196, 197, 5, 44, 0, 0, 197, 66, 1, 0,
		0, 0, 198, 200, 7, 4, 0, 0, 199, 198, 1, 0, 0, 0, 200, 201, 1, 0, 0, 0,
		201, 199, 1, 0, 0, 0, 201, 202, 1, 0, 0, 0, 202, 203, 1, 0, 0, 0, 203,
		204, 6, 33, 0, 0, 204, 68, 1, 0, 0, 0, 205, 206, 5, 47, 0, 0, 206, 207,
		5, 42, 0, 0, 207, 211, 1, 0, 0, 0, 208, 210, 9, 0, 0, 0, 209, 208, 1, 0,
		0, 0, 210, 213, 1, 0, 0, 0, 211, 212, 1, 0, 0, 0, 211, 209, 1, 0, 0, 0,
		212, 214, 1, 0, 0, 0, 213, 211, 1, 0, 0, 0, 214, 215, 5, 42, 0, 0, 215,
		216, 5, 47, 0, 0, 216, 217, 1, 0, 0, 0, 217, 218, 6, 34, 0, 0, 218, 70,
		1, 0, 0, 0, 219, 220, 5, 47, 0, 0, 220, 221, 5, 47, 0, 0, 221, 225, 1,
		0, 0, 0, 222, 224, 8, 5, 0, 0, 223, 222, 1, 0, 0, 0, 224, 227, 1, 0, 0,
		0, 225, 223, 1, 0, 0, 0, 225, 226, 1, 0, 0, 0, 226, 228, 1, 0, 0, 0, 227,
		225, 1, 0, 0, 0, 228, 229, 6, 35, 0, 0, 229, 72, 1, 0, 0, 0, 230, 231,
		5, 92, 0, 0, 231, 232, 7, 6, 0, 0, 232, 74, 1, 0, 0, 0, 9, 0, 124, 130,
		132, 138, 147, 201, 211, 225, 1, 6, 0, 0,
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
	SwiftLexerRTRUE        = 4
	SwiftLexerRFALSE       = 5
	SwiftLexerRPRINT       = 6
	SwiftLexerRIF          = 7
	SwiftLexerRELSE        = 8
	SwiftLexerRWHILE       = 9
	SwiftLexerNUMBER       = 10
	SwiftLexerSTRING       = 11
	SwiftLexerID           = 12
	SwiftLexerDIF          = 13
	SwiftLexerIG_IG        = 14
	SwiftLexerNOT          = 15
	SwiftLexerOR           = 16
	SwiftLexerAND          = 17
	SwiftLexerIG           = 18
	SwiftLexerMAY_IG       = 19
	SwiftLexerMEN_IG       = 20
	SwiftLexerMAYOR        = 21
	SwiftLexerMENOR        = 22
	SwiftLexerMUL          = 23
	SwiftLexerDIV          = 24
	SwiftLexerADD          = 25
	SwiftLexerSUB          = 26
	SwiftLexerPARIZQ       = 27
	SwiftLexerPARDER       = 28
	SwiftLexerLLAVEIZQ     = 29
	SwiftLexerLLAVEDER     = 30
	SwiftLexerPTOCOMA      = 31
	SwiftLexerMODULE       = 32
	SwiftLexerCOMA         = 33
	SwiftLexerWHITESPACE   = 34
	SwiftLexerCOMMENT      = 35
	SwiftLexerLINE_COMMENT = 36
)
