// Code generated from Tebot.g4 by ANTLR 4.13.0. DO NOT EDIT.

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

type TebotLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var TebotLexerLexerStaticData struct {
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

func tebotlexerLexerInit() {
	staticData := &TebotLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'click'", "'input'", "'select'", "'capture'", "'assert'", "'sleep'",
		"'goto'", "'move'", "'scroll'", "'save'", "'paste'", "'='", "'('", "')'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "TEBOT_WHITESPACE",
		"TEBOT_COMMENT", "TEBOT_QUOTED_STRING",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "T__12", "T__13", "TEBOT_WHITESPACE", "TEBOT_COMMENT",
		"TEBOT_QUOTED_STRING",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 17, 134, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1,
		8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1,
		10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 4, 14, 111, 8, 14,
		11, 14, 12, 14, 112, 1, 14, 1, 14, 1, 15, 1, 15, 5, 15, 119, 8, 15, 10,
		15, 12, 15, 122, 9, 15, 1, 15, 1, 15, 1, 16, 1, 16, 5, 16, 128, 8, 16,
		10, 16, 12, 16, 131, 9, 16, 1, 16, 1, 16, 0, 0, 17, 1, 1, 3, 2, 5, 3, 7,
		4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27,
		14, 29, 15, 31, 16, 33, 17, 1, 0, 3, 3, 0, 9, 10, 13, 13, 32, 32, 2, 0,
		10, 10, 13, 13, 3, 0, 10, 10, 13, 13, 34, 34, 136, 0, 1, 1, 0, 0, 0, 0,
		3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0,
		11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0,
		0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0,
		0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0,
		0, 0, 1, 35, 1, 0, 0, 0, 3, 41, 1, 0, 0, 0, 5, 47, 1, 0, 0, 0, 7, 54, 1,
		0, 0, 0, 9, 62, 1, 0, 0, 0, 11, 69, 1, 0, 0, 0, 13, 75, 1, 0, 0, 0, 15,
		80, 1, 0, 0, 0, 17, 85, 1, 0, 0, 0, 19, 92, 1, 0, 0, 0, 21, 97, 1, 0, 0,
		0, 23, 103, 1, 0, 0, 0, 25, 105, 1, 0, 0, 0, 27, 107, 1, 0, 0, 0, 29, 110,
		1, 0, 0, 0, 31, 116, 1, 0, 0, 0, 33, 125, 1, 0, 0, 0, 35, 36, 5, 99, 0,
		0, 36, 37, 5, 108, 0, 0, 37, 38, 5, 105, 0, 0, 38, 39, 5, 99, 0, 0, 39,
		40, 5, 107, 0, 0, 40, 2, 1, 0, 0, 0, 41, 42, 5, 105, 0, 0, 42, 43, 5, 110,
		0, 0, 43, 44, 5, 112, 0, 0, 44, 45, 5, 117, 0, 0, 45, 46, 5, 116, 0, 0,
		46, 4, 1, 0, 0, 0, 47, 48, 5, 115, 0, 0, 48, 49, 5, 101, 0, 0, 49, 50,
		5, 108, 0, 0, 50, 51, 5, 101, 0, 0, 51, 52, 5, 99, 0, 0, 52, 53, 5, 116,
		0, 0, 53, 6, 1, 0, 0, 0, 54, 55, 5, 99, 0, 0, 55, 56, 5, 97, 0, 0, 56,
		57, 5, 112, 0, 0, 57, 58, 5, 116, 0, 0, 58, 59, 5, 117, 0, 0, 59, 60, 5,
		114, 0, 0, 60, 61, 5, 101, 0, 0, 61, 8, 1, 0, 0, 0, 62, 63, 5, 97, 0, 0,
		63, 64, 5, 115, 0, 0, 64, 65, 5, 115, 0, 0, 65, 66, 5, 101, 0, 0, 66, 67,
		5, 114, 0, 0, 67, 68, 5, 116, 0, 0, 68, 10, 1, 0, 0, 0, 69, 70, 5, 115,
		0, 0, 70, 71, 5, 108, 0, 0, 71, 72, 5, 101, 0, 0, 72, 73, 5, 101, 0, 0,
		73, 74, 5, 112, 0, 0, 74, 12, 1, 0, 0, 0, 75, 76, 5, 103, 0, 0, 76, 77,
		5, 111, 0, 0, 77, 78, 5, 116, 0, 0, 78, 79, 5, 111, 0, 0, 79, 14, 1, 0,
		0, 0, 80, 81, 5, 109, 0, 0, 81, 82, 5, 111, 0, 0, 82, 83, 5, 118, 0, 0,
		83, 84, 5, 101, 0, 0, 84, 16, 1, 0, 0, 0, 85, 86, 5, 115, 0, 0, 86, 87,
		5, 99, 0, 0, 87, 88, 5, 114, 0, 0, 88, 89, 5, 111, 0, 0, 89, 90, 5, 108,
		0, 0, 90, 91, 5, 108, 0, 0, 91, 18, 1, 0, 0, 0, 92, 93, 5, 115, 0, 0, 93,
		94, 5, 97, 0, 0, 94, 95, 5, 118, 0, 0, 95, 96, 5, 101, 0, 0, 96, 20, 1,
		0, 0, 0, 97, 98, 5, 112, 0, 0, 98, 99, 5, 97, 0, 0, 99, 100, 5, 115, 0,
		0, 100, 101, 5, 116, 0, 0, 101, 102, 5, 101, 0, 0, 102, 22, 1, 0, 0, 0,
		103, 104, 5, 61, 0, 0, 104, 24, 1, 0, 0, 0, 105, 106, 5, 40, 0, 0, 106,
		26, 1, 0, 0, 0, 107, 108, 5, 41, 0, 0, 108, 28, 1, 0, 0, 0, 109, 111, 7,
		0, 0, 0, 110, 109, 1, 0, 0, 0, 111, 112, 1, 0, 0, 0, 112, 110, 1, 0, 0,
		0, 112, 113, 1, 0, 0, 0, 113, 114, 1, 0, 0, 0, 114, 115, 6, 14, 0, 0, 115,
		30, 1, 0, 0, 0, 116, 120, 5, 35, 0, 0, 117, 119, 8, 1, 0, 0, 118, 117,
		1, 0, 0, 0, 119, 122, 1, 0, 0, 0, 120, 118, 1, 0, 0, 0, 120, 121, 1, 0,
		0, 0, 121, 123, 1, 0, 0, 0, 122, 120, 1, 0, 0, 0, 123, 124, 6, 15, 0, 0,
		124, 32, 1, 0, 0, 0, 125, 129, 5, 34, 0, 0, 126, 128, 8, 2, 0, 0, 127,
		126, 1, 0, 0, 0, 128, 131, 1, 0, 0, 0, 129, 127, 1, 0, 0, 0, 129, 130,
		1, 0, 0, 0, 130, 132, 1, 0, 0, 0, 131, 129, 1, 0, 0, 0, 132, 133, 5, 34,
		0, 0, 133, 34, 1, 0, 0, 0, 4, 0, 112, 120, 129, 1, 0, 1, 0,
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

// TebotLexerInit initializes any static state used to implement TebotLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewTebotLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func TebotLexerInit() {
	staticData := &TebotLexerLexerStaticData
	staticData.once.Do(tebotlexerLexerInit)
}

// NewTebotLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewTebotLexer(input antlr.CharStream) *TebotLexer {
	TebotLexerInit()
	l := new(TebotLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &TebotLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "Tebot.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// TebotLexer tokens.
const (
	TebotLexerT__0                = 1
	TebotLexerT__1                = 2
	TebotLexerT__2                = 3
	TebotLexerT__3                = 4
	TebotLexerT__4                = 5
	TebotLexerT__5                = 6
	TebotLexerT__6                = 7
	TebotLexerT__7                = 8
	TebotLexerT__8                = 9
	TebotLexerT__9                = 10
	TebotLexerT__10               = 11
	TebotLexerT__11               = 12
	TebotLexerT__12               = 13
	TebotLexerT__13               = 14
	TebotLexerTEBOT_WHITESPACE    = 15
	TebotLexerTEBOT_COMMENT       = 16
	TebotLexerTEBOT_QUOTED_STRING = 17
)
