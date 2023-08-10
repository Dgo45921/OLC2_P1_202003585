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
		"')'", "'{'", "'}'", "';'", "'%'",
	}
	staticData.SymbolicNames = []string{
		"", "RINT", "RFLOAT", "RBOOL", "RTRUE", "RFALSE", "RPRINT", "RIF", "RELSE",
		"RWHILE", "NUMBER", "STRING", "ID", "DIF", "IG_IG", "NOT", "OR", "AND",
		"IG", "MAY_IG", "MEN_IG", "MAYOR", "MENOR", "MUL", "DIV", "ADD", "SUB",
		"PARIZQ", "PARDER", "LLAVEIZQ", "LLAVEDER", "PTOCOMA", "MODULE", "WHITESPACE",
		"COMMENT", "LINE_COMMENT",
	}
	staticData.RuleNames = []string{
		"RINT", "RFLOAT", "RBOOL", "RTRUE", "RFALSE", "RPRINT", "RIF", "RELSE",
		"RWHILE", "NUMBER", "STRING", "ID", "DIF", "IG_IG", "NOT", "OR", "AND",
		"IG", "MAY_IG", "MEN_IG", "MAYOR", "MENOR", "MUL", "DIV", "ADD", "SUB",
		"PARIZQ", "PARDER", "LLAVEIZQ", "LLAVEDER", "PTOCOMA", "MODULE", "WHITESPACE",
		"COMMENT", "LINE_COMMENT", "ESC_SEQ",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 35, 229, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 1, 0,
		1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 4, 9, 121,
		8, 9, 11, 9, 12, 9, 122, 1, 9, 1, 9, 4, 9, 127, 8, 9, 11, 9, 12, 9, 128,
		3, 9, 131, 8, 9, 1, 10, 1, 10, 5, 10, 135, 8, 10, 10, 10, 12, 10, 138,
		9, 10, 1, 10, 1, 10, 1, 11, 1, 11, 5, 11, 144, 8, 11, 10, 11, 12, 11, 147,
		9, 11, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1,
		15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 19,
		1, 19, 1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 1, 22, 1, 22, 1, 23, 1, 23, 1,
		24, 1, 24, 1, 25, 1, 25, 1, 26, 1, 26, 1, 27, 1, 27, 1, 28, 1, 28, 1, 29,
		1, 29, 1, 30, 1, 30, 1, 31, 1, 31, 1, 32, 4, 32, 196, 8, 32, 11, 32, 12,
		32, 197, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 1, 33, 5, 33, 206, 8, 33, 10,
		33, 12, 33, 209, 9, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 34, 1, 34,
		1, 34, 1, 34, 5, 34, 220, 8, 34, 10, 34, 12, 34, 223, 9, 34, 1, 34, 1,
		34, 1, 35, 1, 35, 1, 35, 1, 207, 0, 36, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11,
		6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15,
		31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24,
		49, 25, 51, 26, 53, 27, 55, 28, 57, 29, 59, 30, 61, 31, 63, 32, 65, 33,
		67, 34, 69, 35, 71, 0, 1, 0, 7, 1, 0, 48, 57, 1, 0, 34, 34, 2, 0, 65, 90,
		97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 4, 0, 9, 10, 13, 13, 32,
		32, 92, 92, 2, 0, 10, 10, 13, 13, 7, 0, 32, 33, 35, 35, 43, 43, 45, 46,
		58, 58, 64, 64, 91, 93, 235, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5,
		1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13,
		1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0,
		21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0,
		0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0,
		0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0,
		0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1,
		0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59,
		1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0,
		67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 1, 73, 1, 0, 0, 0, 3, 77, 1, 0, 0, 0,
		5, 83, 1, 0, 0, 0, 7, 88, 1, 0, 0, 0, 9, 93, 1, 0, 0, 0, 11, 99, 1, 0,
		0, 0, 13, 105, 1, 0, 0, 0, 15, 108, 1, 0, 0, 0, 17, 113, 1, 0, 0, 0, 19,
		120, 1, 0, 0, 0, 21, 132, 1, 0, 0, 0, 23, 141, 1, 0, 0, 0, 25, 148, 1,
		0, 0, 0, 27, 151, 1, 0, 0, 0, 29, 154, 1, 0, 0, 0, 31, 156, 1, 0, 0, 0,
		33, 159, 1, 0, 0, 0, 35, 162, 1, 0, 0, 0, 37, 164, 1, 0, 0, 0, 39, 167,
		1, 0, 0, 0, 41, 170, 1, 0, 0, 0, 43, 172, 1, 0, 0, 0, 45, 174, 1, 0, 0,
		0, 47, 176, 1, 0, 0, 0, 49, 178, 1, 0, 0, 0, 51, 180, 1, 0, 0, 0, 53, 182,
		1, 0, 0, 0, 55, 184, 1, 0, 0, 0, 57, 186, 1, 0, 0, 0, 59, 188, 1, 0, 0,
		0, 61, 190, 1, 0, 0, 0, 63, 192, 1, 0, 0, 0, 65, 195, 1, 0, 0, 0, 67, 201,
		1, 0, 0, 0, 69, 215, 1, 0, 0, 0, 71, 226, 1, 0, 0, 0, 73, 74, 5, 105, 0,
		0, 74, 75, 5, 110, 0, 0, 75, 76, 5, 116, 0, 0, 76, 2, 1, 0, 0, 0, 77, 78,
		5, 102, 0, 0, 78, 79, 5, 108, 0, 0, 79, 80, 5, 111, 0, 0, 80, 81, 5, 97,
		0, 0, 81, 82, 5, 116, 0, 0, 82, 4, 1, 0, 0, 0, 83, 84, 5, 98, 0, 0, 84,
		85, 5, 111, 0, 0, 85, 86, 5, 111, 0, 0, 86, 87, 5, 108, 0, 0, 87, 6, 1,
		0, 0, 0, 88, 89, 5, 116, 0, 0, 89, 90, 5, 114, 0, 0, 90, 91, 5, 117, 0,
		0, 91, 92, 5, 101, 0, 0, 92, 8, 1, 0, 0, 0, 93, 94, 5, 102, 0, 0, 94, 95,
		5, 97, 0, 0, 95, 96, 5, 108, 0, 0, 96, 97, 5, 115, 0, 0, 97, 98, 5, 101,
		0, 0, 98, 10, 1, 0, 0, 0, 99, 100, 5, 112, 0, 0, 100, 101, 5, 114, 0, 0,
		101, 102, 5, 105, 0, 0, 102, 103, 5, 110, 0, 0, 103, 104, 5, 116, 0, 0,
		104, 12, 1, 0, 0, 0, 105, 106, 5, 105, 0, 0, 106, 107, 5, 102, 0, 0, 107,
		14, 1, 0, 0, 0, 108, 109, 5, 101, 0, 0, 109, 110, 5, 108, 0, 0, 110, 111,
		5, 115, 0, 0, 111, 112, 5, 101, 0, 0, 112, 16, 1, 0, 0, 0, 113, 114, 5,
		119, 0, 0, 114, 115, 5, 104, 0, 0, 115, 116, 5, 105, 0, 0, 116, 117, 5,
		108, 0, 0, 117, 118, 5, 101, 0, 0, 118, 18, 1, 0, 0, 0, 119, 121, 7, 0,
		0, 0, 120, 119, 1, 0, 0, 0, 121, 122, 1, 0, 0, 0, 122, 120, 1, 0, 0, 0,
		122, 123, 1, 0, 0, 0, 123, 130, 1, 0, 0, 0, 124, 126, 5, 46, 0, 0, 125,
		127, 7, 0, 0, 0, 126, 125, 1, 0, 0, 0, 127, 128, 1, 0, 0, 0, 128, 126,
		1, 0, 0, 0, 128, 129, 1, 0, 0, 0, 129, 131, 1, 0, 0, 0, 130, 124, 1, 0,
		0, 0, 130, 131, 1, 0, 0, 0, 131, 20, 1, 0, 0, 0, 132, 136, 5, 34, 0, 0,
		133, 135, 8, 1, 0, 0, 134, 133, 1, 0, 0, 0, 135, 138, 1, 0, 0, 0, 136,
		134, 1, 0, 0, 0, 136, 137, 1, 0, 0, 0, 137, 139, 1, 0, 0, 0, 138, 136,
		1, 0, 0, 0, 139, 140, 5, 34, 0, 0, 140, 22, 1, 0, 0, 0, 141, 145, 7, 2,
		0, 0, 142, 144, 7, 3, 0, 0, 143, 142, 1, 0, 0, 0, 144, 147, 1, 0, 0, 0,
		145, 143, 1, 0, 0, 0, 145, 146, 1, 0, 0, 0, 146, 24, 1, 0, 0, 0, 147, 145,
		1, 0, 0, 0, 148, 149, 5, 33, 0, 0, 149, 150, 5, 61, 0, 0, 150, 26, 1, 0,
		0, 0, 151, 152, 5, 61, 0, 0, 152, 153, 5, 61, 0, 0, 153, 28, 1, 0, 0, 0,
		154, 155, 5, 33, 0, 0, 155, 30, 1, 0, 0, 0, 156, 157, 5, 124, 0, 0, 157,
		158, 5, 124, 0, 0, 158, 32, 1, 0, 0, 0, 159, 160, 5, 38, 0, 0, 160, 161,
		5, 38, 0, 0, 161, 34, 1, 0, 0, 0, 162, 163, 5, 61, 0, 0, 163, 36, 1, 0,
		0, 0, 164, 165, 5, 62, 0, 0, 165, 166, 5, 61, 0, 0, 166, 38, 1, 0, 0, 0,
		167, 168, 5, 60, 0, 0, 168, 169, 5, 61, 0, 0, 169, 40, 1, 0, 0, 0, 170,
		171, 5, 62, 0, 0, 171, 42, 1, 0, 0, 0, 172, 173, 5, 60, 0, 0, 173, 44,
		1, 0, 0, 0, 174, 175, 5, 42, 0, 0, 175, 46, 1, 0, 0, 0, 176, 177, 5, 47,
		0, 0, 177, 48, 1, 0, 0, 0, 178, 179, 5, 43, 0, 0, 179, 50, 1, 0, 0, 0,
		180, 181, 5, 45, 0, 0, 181, 52, 1, 0, 0, 0, 182, 183, 5, 40, 0, 0, 183,
		54, 1, 0, 0, 0, 184, 185, 5, 41, 0, 0, 185, 56, 1, 0, 0, 0, 186, 187, 5,
		123, 0, 0, 187, 58, 1, 0, 0, 0, 188, 189, 5, 125, 0, 0, 189, 60, 1, 0,
		0, 0, 190, 191, 5, 59, 0, 0, 191, 62, 1, 0, 0, 0, 192, 193, 5, 37, 0, 0,
		193, 64, 1, 0, 0, 0, 194, 196, 7, 4, 0, 0, 195, 194, 1, 0, 0, 0, 196, 197,
		1, 0, 0, 0, 197, 195, 1, 0, 0, 0, 197, 198, 1, 0, 0, 0, 198, 199, 1, 0,
		0, 0, 199, 200, 6, 32, 0, 0, 200, 66, 1, 0, 0, 0, 201, 202, 5, 47, 0, 0,
		202, 203, 5, 42, 0, 0, 203, 207, 1, 0, 0, 0, 204, 206, 9, 0, 0, 0, 205,
		204, 1, 0, 0, 0, 206, 209, 1, 0, 0, 0, 207, 208, 1, 0, 0, 0, 207, 205,
		1, 0, 0, 0, 208, 210, 1, 0, 0, 0, 209, 207, 1, 0, 0, 0, 210, 211, 5, 42,
		0, 0, 211, 212, 5, 47, 0, 0, 212, 213, 1, 0, 0, 0, 213, 214, 6, 33, 0,
		0, 214, 68, 1, 0, 0, 0, 215, 216, 5, 47, 0, 0, 216, 217, 5, 47, 0, 0, 217,
		221, 1, 0, 0, 0, 218, 220, 8, 5, 0, 0, 219, 218, 1, 0, 0, 0, 220, 223,
		1, 0, 0, 0, 221, 219, 1, 0, 0, 0, 221, 222, 1, 0, 0, 0, 222, 224, 1, 0,
		0, 0, 223, 221, 1, 0, 0, 0, 224, 225, 6, 34, 0, 0, 225, 70, 1, 0, 0, 0,
		226, 227, 5, 92, 0, 0, 227, 228, 7, 6, 0, 0, 228, 72, 1, 0, 0, 0, 9, 0,
		122, 128, 130, 136, 145, 197, 207, 221, 1, 6, 0, 0,
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
	SwiftLexerWHITESPACE   = 33
	SwiftLexerCOMMENT      = 34
	SwiftLexerLINE_COMMENT = 35
)
