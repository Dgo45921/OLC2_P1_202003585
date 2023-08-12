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
		"s", "block", "arguments", "argument", "instruction", "printstmt", "vardec",
		"constdec", "asignation", "ifstmt", "expr",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 42, 191, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 4, 1, 28, 8, 1, 11, 1, 12, 1, 29, 1,
		1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 43,
		8, 2, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 3, 4, 50, 8, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 3, 4, 56, 8, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 62, 8, 4, 1, 4, 1, 4,
		1, 4, 1, 4, 3, 4, 68, 8, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 75, 8,
		4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 6, 1, 6, 3, 6, 103, 8, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1,
		7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 119, 8, 7, 1, 8, 1,
		8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1,
		10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 3, 10, 154,
		8, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1,
		10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 5,
		10, 186, 8, 10, 10, 10, 12, 10, 189, 9, 10, 1, 10, 0, 1, 20, 11, 0, 2,
		4, 6, 8, 10, 12, 14, 16, 18, 20, 0, 6, 1, 0, 1, 5, 2, 0, 19, 19, 30, 30,
		2, 0, 27, 28, 37, 37, 1, 0, 29, 30, 1, 0, 23, 26, 1, 0, 17, 18, 205, 0,
		22, 1, 0, 0, 0, 2, 27, 1, 0, 0, 0, 4, 42, 1, 0, 0, 0, 6, 44, 1, 0, 0, 0,
		8, 74, 1, 0, 0, 0, 10, 76, 1, 0, 0, 0, 12, 102, 1, 0, 0, 0, 14, 118, 1,
		0, 0, 0, 16, 120, 1, 0, 0, 0, 18, 125, 1, 0, 0, 0, 20, 153, 1, 0, 0, 0,
		22, 23, 3, 2, 1, 0, 23, 24, 5, 0, 0, 1, 24, 25, 6, 0, -1, 0, 25, 1, 1,
		0, 0, 0, 26, 28, 3, 8, 4, 0, 27, 26, 1, 0, 0, 0, 28, 29, 1, 0, 0, 0, 29,
		27, 1, 0, 0, 0, 29, 30, 1, 0, 0, 0, 30, 31, 1, 0, 0, 0, 31, 32, 6, 1, -1,
		0, 32, 3, 1, 0, 0, 0, 33, 34, 3, 6, 3, 0, 34, 35, 5, 38, 0, 0, 35, 36,
		3, 4, 2, 0, 36, 37, 6, 2, -1, 0, 37, 43, 1, 0, 0, 0, 38, 39, 3, 6, 3, 0,
		39, 40, 6, 2, -1, 0, 40, 43, 1, 0, 0, 0, 41, 43, 6, 2, -1, 0, 42, 33, 1,
		0, 0, 0, 42, 38, 1, 0, 0, 0, 42, 41, 1, 0, 0, 0, 43, 5, 1, 0, 0, 0, 44,
		45, 3, 20, 10, 0, 45, 46, 6, 3, -1, 0, 46, 7, 1, 0, 0, 0, 47, 49, 3, 10,
		5, 0, 48, 50, 5, 35, 0, 0, 49, 48, 1, 0, 0, 0, 49, 50, 1, 0, 0, 0, 50,
		51, 1, 0, 0, 0, 51, 52, 6, 4, -1, 0, 52, 75, 1, 0, 0, 0, 53, 55, 3, 12,
		6, 0, 54, 56, 5, 35, 0, 0, 55, 54, 1, 0, 0, 0, 55, 56, 1, 0, 0, 0, 56,
		57, 1, 0, 0, 0, 57, 58, 6, 4, -1, 0, 58, 75, 1, 0, 0, 0, 59, 61, 3, 14,
		7, 0, 60, 62, 5, 35, 0, 0, 61, 60, 1, 0, 0, 0, 61, 62, 1, 0, 0, 0, 62,
		63, 1, 0, 0, 0, 63, 64, 6, 4, -1, 0, 64, 75, 1, 0, 0, 0, 65, 67, 3, 16,
		8, 0, 66, 68, 5, 35, 0, 0, 67, 66, 1, 0, 0, 0, 67, 68, 1, 0, 0, 0, 68,
		69, 1, 0, 0, 0, 69, 70, 6, 4, -1, 0, 70, 75, 1, 0, 0, 0, 71, 72, 3, 18,
		9, 0, 72, 73, 6, 4, -1, 0, 73, 75, 1, 0, 0, 0, 74, 47, 1, 0, 0, 0, 74,
		53, 1, 0, 0, 0, 74, 59, 1, 0, 0, 0, 74, 65, 1, 0, 0, 0, 74, 71, 1, 0, 0,
		0, 75, 9, 1, 0, 0, 0, 76, 77, 5, 8, 0, 0, 77, 78, 5, 31, 0, 0, 78, 79,
		3, 4, 2, 0, 79, 80, 5, 32, 0, 0, 80, 81, 6, 5, -1, 0, 81, 11, 1, 0, 0,
		0, 82, 83, 5, 12, 0, 0, 83, 84, 5, 16, 0, 0, 84, 85, 5, 36, 0, 0, 85, 86,
		7, 0, 0, 0, 86, 87, 5, 22, 0, 0, 87, 88, 3, 20, 10, 0, 88, 89, 6, 6, -1,
		0, 89, 103, 1, 0, 0, 0, 90, 91, 5, 12, 0, 0, 91, 92, 5, 16, 0, 0, 92, 93,
		5, 22, 0, 0, 93, 94, 3, 20, 10, 0, 94, 95, 6, 6, -1, 0, 95, 103, 1, 0,
		0, 0, 96, 97, 5, 12, 0, 0, 97, 98, 5, 16, 0, 0, 98, 99, 5, 36, 0, 0, 99,
		100, 7, 0, 0, 0, 100, 101, 5, 39, 0, 0, 101, 103, 6, 6, -1, 0, 102, 82,
		1, 0, 0, 0, 102, 90, 1, 0, 0, 0, 102, 96, 1, 0, 0, 0, 103, 13, 1, 0, 0,
		0, 104, 105, 5, 13, 0, 0, 105, 106, 5, 16, 0, 0, 106, 107, 5, 36, 0, 0,
		107, 108, 7, 0, 0, 0, 108, 109, 5, 22, 0, 0, 109, 110, 3, 20, 10, 0, 110,
		111, 6, 7, -1, 0, 111, 119, 1, 0, 0, 0, 112, 113, 5, 13, 0, 0, 113, 114,
		5, 16, 0, 0, 114, 115, 5, 22, 0, 0, 115, 116, 3, 20, 10, 0, 116, 117, 6,
		7, -1, 0, 117, 119, 1, 0, 0, 0, 118, 104, 1, 0, 0, 0, 118, 112, 1, 0, 0,
		0, 119, 15, 1, 0, 0, 0, 120, 121, 5, 16, 0, 0, 121, 122, 5, 22, 0, 0, 122,
		123, 3, 20, 10, 0, 123, 124, 6, 8, -1, 0, 124, 17, 1, 0, 0, 0, 125, 126,
		5, 9, 0, 0, 126, 127, 5, 31, 0, 0, 127, 128, 3, 20, 10, 0, 128, 129, 5,
		32, 0, 0, 129, 130, 5, 33, 0, 0, 130, 131, 3, 2, 1, 0, 131, 132, 5, 34,
		0, 0, 132, 19, 1, 0, 0, 0, 133, 134, 6, 10, -1, 0, 134, 135, 5, 31, 0,
		0, 135, 136, 3, 20, 10, 0, 136, 137, 5, 32, 0, 0, 137, 138, 6, 10, -1,
		0, 138, 154, 1, 0, 0, 0, 139, 140, 7, 1, 0, 0, 140, 141, 3, 20, 10, 12,
		141, 142, 6, 10, -1, 0, 142, 154, 1, 0, 0, 0, 143, 144, 5, 16, 0, 0, 144,
		154, 6, 10, -1, 0, 145, 146, 5, 14, 0, 0, 146, 154, 6, 10, -1, 0, 147,
		148, 5, 15, 0, 0, 148, 154, 6, 10, -1, 0, 149, 150, 5, 6, 0, 0, 150, 154,
		6, 10, -1, 0, 151, 152, 5, 7, 0, 0, 152, 154, 6, 10, -1, 0, 153, 133, 1,
		0, 0, 0, 153, 139, 1, 0, 0, 0, 153, 143, 1, 0, 0, 0, 153, 145, 1, 0, 0,
		0, 153, 147, 1, 0, 0, 0, 153, 149, 1, 0, 0, 0, 153, 151, 1, 0, 0, 0, 154,
		187, 1, 0, 0, 0, 155, 156, 10, 11, 0, 0, 156, 157, 7, 2, 0, 0, 157, 158,
		3, 20, 10, 12, 158, 159, 6, 10, -1, 0, 159, 186, 1, 0, 0, 0, 160, 161,
		10, 10, 0, 0, 161, 162, 7, 3, 0, 0, 162, 163, 3, 20, 10, 11, 163, 164,
		6, 10, -1, 0, 164, 186, 1, 0, 0, 0, 165, 166, 10, 9, 0, 0, 166, 167, 7,
		4, 0, 0, 167, 168, 3, 20, 10, 10, 168, 169, 6, 10, -1, 0, 169, 186, 1,
		0, 0, 0, 170, 171, 10, 8, 0, 0, 171, 172, 7, 5, 0, 0, 172, 173, 3, 20,
		10, 9, 173, 174, 6, 10, -1, 0, 174, 186, 1, 0, 0, 0, 175, 176, 10, 7, 0,
		0, 176, 177, 5, 21, 0, 0, 177, 178, 3, 20, 10, 8, 178, 179, 6, 10, -1,
		0, 179, 186, 1, 0, 0, 0, 180, 181, 10, 6, 0, 0, 181, 182, 5, 20, 0, 0,
		182, 183, 3, 20, 10, 7, 183, 184, 6, 10, -1, 0, 184, 186, 1, 0, 0, 0, 185,
		155, 1, 0, 0, 0, 185, 160, 1, 0, 0, 0, 185, 165, 1, 0, 0, 0, 185, 170,
		1, 0, 0, 0, 185, 175, 1, 0, 0, 0, 185, 180, 1, 0, 0, 0, 186, 189, 1, 0,
		0, 0, 187, 185, 1, 0, 0, 0, 187, 188, 1, 0, 0, 0, 188, 21, 1, 0, 0, 0,
		189, 187, 1, 0, 0, 0, 12, 29, 42, 49, 55, 61, 67, 74, 102, 118, 153, 185,
		187,
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
	SwiftGrammarParserNUMBER       = 14
	SwiftGrammarParserSTRING       = 15
	SwiftGrammarParserID           = 16
	SwiftGrammarParserDIF          = 17
	SwiftGrammarParserIG_IG        = 18
	SwiftGrammarParserNOT          = 19
	SwiftGrammarParserOR           = 20
	SwiftGrammarParserAND          = 21
	SwiftGrammarParserIG           = 22
	SwiftGrammarParserMAY_IG       = 23
	SwiftGrammarParserMEN_IG       = 24
	SwiftGrammarParserMAYOR        = 25
	SwiftGrammarParserMENOR        = 26
	SwiftGrammarParserMUL          = 27
	SwiftGrammarParserDIV          = 28
	SwiftGrammarParserADD          = 29
	SwiftGrammarParserSUB          = 30
	SwiftGrammarParserPARIZQ       = 31
	SwiftGrammarParserPARDER       = 32
	SwiftGrammarParserLLAVEIZQ     = 33
	SwiftGrammarParserLLAVEDER     = 34
	SwiftGrammarParserPTOCOMA      = 35
	SwiftGrammarParserDOSPTOS      = 36
	SwiftGrammarParserMODULE       = 37
	SwiftGrammarParserCOMA         = 38
	SwiftGrammarParserQM           = 39
	SwiftGrammarParserWHITESPACE   = 40
	SwiftGrammarParserCOMMENT      = 41
	SwiftGrammarParserLINE_COMMENT = 42
)

// SwiftGrammarParser rules.
const (
	SwiftGrammarParserRULE_s           = 0
	SwiftGrammarParserRULE_block       = 1
	SwiftGrammarParserRULE_arguments   = 2
	SwiftGrammarParserRULE_argument    = 3
	SwiftGrammarParserRULE_instruction = 4
	SwiftGrammarParserRULE_printstmt   = 5
	SwiftGrammarParserRULE_vardec      = 6
	SwiftGrammarParserRULE_constdec    = 7
	SwiftGrammarParserRULE_asignation  = 8
	SwiftGrammarParserRULE_ifstmt      = 9
	SwiftGrammarParserRULE_expr        = 10
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
		p.SetState(22)

		var _x = p.Block()

		localctx.(*SContext)._block = _x
	}
	{
		p.SetState(23)
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
	p.SetState(27)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&78592) != 0) {
		{
			p.SetState(26)

			var _x = p.Instruction()

			localctx.(*BlockContext)._instruction = _x
		}
		localctx.(*BlockContext).ins = append(localctx.(*BlockContext).ins, localctx.(*BlockContext)._instruction)

		p.SetState(29)
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

	p.SetState(42)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(33)

			var _x = p.Argument()

			localctx.(*ArgumentsContext)._argument = _x
		}
		{
			p.SetState(34)
			p.Match(SwiftGrammarParserCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(35)

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
			p.SetState(38)

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
		p.SetState(44)

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

	// Get_asignation returns the _asignation rule contexts.
	Get_asignation() IAsignationContext

	// Set_printstmt sets the _printstmt rule contexts.
	Set_printstmt(IPrintstmtContext)

	// Set_vardec sets the _vardec rule contexts.
	Set_vardec(IVardecContext)

	// Set_constdec sets the _constdec rule contexts.
	Set_constdec(IConstdecContext)

	// Set_asignation sets the _asignation rule contexts.
	Set_asignation(IAsignationContext)

	// GetInst returns the inst attribute.
	GetInst() interfaces.Instruction

	// SetInst sets the inst attribute.
	SetInst(interfaces.Instruction)

	// Getter signatures
	Printstmt() IPrintstmtContext
	PTOCOMA() antlr.TerminalNode
	Vardec() IVardecContext
	Constdec() IConstdecContext
	Asignation() IAsignationContext
	Ifstmt() IIfstmtContext

	// IsInstructionContext differentiates from other interfaces.
	IsInstructionContext()
}

type InstructionContext struct {
	antlr.BaseParserRuleContext
	parser      antlr.Parser
	inst        interfaces.Instruction
	_printstmt  IPrintstmtContext
	_vardec     IVardecContext
	_constdec   IConstdecContext
	_asignation IAsignationContext
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

func (s *InstructionContext) Get_asignation() IAsignationContext { return s._asignation }

func (s *InstructionContext) Set_printstmt(v IPrintstmtContext) { s._printstmt = v }

func (s *InstructionContext) Set_vardec(v IVardecContext) { s._vardec = v }

func (s *InstructionContext) Set_constdec(v IConstdecContext) { s._constdec = v }

func (s *InstructionContext) Set_asignation(v IAsignationContext) { s._asignation = v }

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

	p.SetState(74)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SwiftGrammarParserRPRINT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(47)

			var _x = p.Printstmt()

			localctx.(*InstructionContext)._printstmt = _x
		}
		p.SetState(49)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserPTOCOMA {
			{
				p.SetState(48)
				p.Match(SwiftGrammarParserPTOCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_printstmt().GetPrnt()

	case SwiftGrammarParserRVAR:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(53)

			var _x = p.Vardec()

			localctx.(*InstructionContext)._vardec = _x
		}
		p.SetState(55)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserPTOCOMA {
			{
				p.SetState(54)
				p.Match(SwiftGrammarParserPTOCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_vardec().GetNewdec()

	case SwiftGrammarParserRLET:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(59)

			var _x = p.Constdec()

			localctx.(*InstructionContext)._constdec = _x
		}
		p.SetState(61)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserPTOCOMA {
			{
				p.SetState(60)
				p.Match(SwiftGrammarParserPTOCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_constdec().GetNewconst()

	case SwiftGrammarParserID:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(65)

			var _x = p.Asignation()

			localctx.(*InstructionContext)._asignation = _x
		}
		p.SetState(67)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SwiftGrammarParserPTOCOMA {
			{
				p.SetState(66)
				p.Match(SwiftGrammarParserPTOCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_asignation().GetNewasignation()

	case SwiftGrammarParserRIF:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(71)
			p.Ifstmt()
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
	p.EnterRule(localctx, 10, SwiftGrammarParserRULE_printstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(76)

		var _m = p.Match(SwiftGrammarParserRPRINT)

		localctx.(*PrintstmtContext)._RPRINT = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(77)
		p.Match(SwiftGrammarParserPARIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(78)

		var _x = p.Arguments()

		localctx.(*PrintstmtContext)._arguments = _x
	}
	{
		p.SetState(79)
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
	p.EnterRule(localctx, 12, SwiftGrammarParserRULE_vardec)
	var _la int

	p.SetState(102)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(82)

			var _m = p.Match(SwiftGrammarParserRVAR)

			localctx.(*VardecContext)._RVAR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(83)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*VardecContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(84)
			p.Match(SwiftGrammarParserDOSPTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(85)

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
			p.SetState(86)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(87)

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
			p.SetState(90)

			var _m = p.Match(SwiftGrammarParserRVAR)

			localctx.(*VardecContext)._RVAR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(91)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*VardecContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(92)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(93)

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
			p.SetState(96)

			var _m = p.Match(SwiftGrammarParserRVAR)

			localctx.(*VardecContext)._RVAR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(97)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*VardecContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(98)
			p.Match(SwiftGrammarParserDOSPTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(99)

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
			p.SetState(100)
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
	p.EnterRule(localctx, 14, SwiftGrammarParserRULE_constdec)
	var _la int

	p.SetState(118)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(104)

			var _m = p.Match(SwiftGrammarParserRLET)

			localctx.(*ConstdecContext)._RLET = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(105)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*ConstdecContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(106)
			p.Match(SwiftGrammarParserDOSPTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(107)

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
			p.SetState(108)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(109)

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
			p.SetState(112)

			var _m = p.Match(SwiftGrammarParserRLET)

			localctx.(*ConstdecContext)._RLET = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(113)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*ConstdecContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(114)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(115)

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
	p.EnterRule(localctx, 16, SwiftGrammarParserRULE_asignation)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(120)

		var _m = p.Match(SwiftGrammarParserID)

		localctx.(*AsignationContext)._ID = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(121)
		p.Match(SwiftGrammarParserIG)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(122)

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

// IIfstmtContext is an interface to support dynamic dispatch.
type IIfstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RIF() antlr.TerminalNode
	PARIZQ() antlr.TerminalNode
	Expr() IExprContext
	PARDER() antlr.TerminalNode
	LLAVEIZQ() antlr.TerminalNode
	Block() IBlockContext
	LLAVEDER() antlr.TerminalNode

	// IsIfstmtContext differentiates from other interfaces.
	IsIfstmtContext()
}

type IfstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
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

func (s *IfstmtContext) RIF() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserRIF, 0)
}

func (s *IfstmtContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARIZQ, 0)
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

func (s *IfstmtContext) PARDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARDER, 0)
}

func (s *IfstmtContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEIZQ, 0)
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

func (s *IfstmtContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEDER, 0)
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
	p.EnterRule(localctx, 18, SwiftGrammarParserRULE_ifstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(125)
		p.Match(SwiftGrammarParserRIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(126)
		p.Match(SwiftGrammarParserPARIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(127)
		p.expr(0)
	}
	{
		p.SetState(128)
		p.Match(SwiftGrammarParserPARDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(129)
		p.Match(SwiftGrammarParserLLAVEIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(130)
		p.Block()
	}
	{
		p.SetState(131)
		p.Match(SwiftGrammarParserLLAVEDER)
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

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

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

	// GetRight returns the right rule contexts.
	GetRight() IExprContext

	// SetLeft sets the left rule contexts.
	SetLeft(IExprContext)

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

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
	ID() antlr.TerminalNode
	NUMBER() antlr.TerminalNode
	STRING() antlr.TerminalNode
	RTRUE() antlr.TerminalNode
	RFALSE() antlr.TerminalNode
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
	parser  antlr.Parser
	e       interfaces.Expression
	left    IExprContext
	_expr   IExprContext
	op      antlr.Token
	_ID     antlr.Token
	_NUMBER antlr.Token
	_STRING antlr.Token
	_RTRUE  antlr.Token
	_RFALSE antlr.Token
	right   IExprContext
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

func (s *ExprContext) Get_ID() antlr.Token { return s._ID }

func (s *ExprContext) Get_NUMBER() antlr.Token { return s._NUMBER }

func (s *ExprContext) Get_STRING() antlr.Token { return s._STRING }

func (s *ExprContext) Get_RTRUE() antlr.Token { return s._RTRUE }

func (s *ExprContext) Get_RFALSE() antlr.Token { return s._RFALSE }

func (s *ExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *ExprContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *ExprContext) Set_NUMBER(v antlr.Token) { s._NUMBER = v }

func (s *ExprContext) Set_STRING(v antlr.Token) { s._STRING = v }

func (s *ExprContext) Set_RTRUE(v antlr.Token) { s._RTRUE = v }

func (s *ExprContext) Set_RFALSE(v antlr.Token) { s._RFALSE = v }

func (s *ExprContext) GetLeft() IExprContext { return s.left }

func (s *ExprContext) Get_expr() IExprContext { return s._expr }

func (s *ExprContext) GetRight() IExprContext { return s.right }

func (s *ExprContext) SetLeft(v IExprContext) { s.left = v }

func (s *ExprContext) Set_expr(v IExprContext) { s._expr = v }

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
	_startState := 20
	p.EnterRecursionRule(localctx, 20, SwiftGrammarParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(153)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SwiftGrammarParserPARIZQ:
		{
			p.SetState(134)
			p.Match(SwiftGrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(135)

			var _x = p.expr(0)

			localctx.(*ExprContext)._expr = _x
		}
		{
			p.SetState(136)
			p.Match(SwiftGrammarParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ExprContext).e = localctx.(*ExprContext).Get_expr().GetE()

	case SwiftGrammarParserNOT, SwiftGrammarParserSUB:
		{
			p.SetState(139)

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
			p.SetState(140)

			var _x = p.expr(12)

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

	case SwiftGrammarParserID:
		{
			p.SetState(143)

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
			p.SetState(145)

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
			p.SetState(147)

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
			}()), str[1:len(str)-1], environment.CHAR)
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
			}()), str[1:len(str)-1], environment.STRING)
		}

	case SwiftGrammarParserRTRUE:
		{
			p.SetState(149)

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
			p.SetState(151)

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

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(187)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(185)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_expr)
				p.SetState(155)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
					goto errorExit
				}
				{
					p.SetState(156)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&137841606656) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(157)

					var _x = p.expr(12)

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
				p.SetState(160)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(161)

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
					p.SetState(162)

					var _x = p.expr(11)

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
				p.SetState(165)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(166)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&125829120) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(167)

					var _x = p.expr(10)

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
				p.SetState(170)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(171)

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
					p.SetState(172)

					var _x = p.expr(9)

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
				p.SetState(175)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(176)

					var _m = p.Match(SwiftGrammarParserAND)

					localctx.(*ExprContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(177)

					var _x = p.expr(8)

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
				p.SetState(180)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(181)

					var _m = p.Match(SwiftGrammarParserOR)

					localctx.(*ExprContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(182)

					var _x = p.expr(7)

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
		p.SetState(189)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext())
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
	case 10:
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
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 6)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
