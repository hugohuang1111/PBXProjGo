// Code generated from PBXProj.g4 by ANTLR 4.8. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 15, 104,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 3,
	2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3,
	5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11,
	3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 7, 12, 65, 10, 12, 12, 12, 14, 12, 68,
	11, 12, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 5, 13, 75, 10, 13, 3, 14, 3,
	14, 3, 14, 7, 14, 80, 10, 14, 12, 14, 14, 14, 83, 11, 14, 3, 14, 3, 14,
	3, 15, 6, 15, 88, 10, 15, 13, 15, 14, 15, 89, 3, 16, 3, 16, 3, 16, 3, 16,
	5, 16, 96, 10, 16, 3, 17, 6, 17, 99, 10, 17, 13, 17, 14, 17, 100, 3, 17,
	3, 17, 4, 66, 81, 2, 18, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17,
	10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 2, 29, 2, 31, 2, 33, 15, 3, 2,
	4, 6, 2, 47, 59, 65, 92, 97, 97, 99, 124, 5, 2, 11, 12, 14, 15, 34, 34,
	2, 107, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3,
	2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17,
	3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2,
	25, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 3, 35, 3, 2, 2, 2, 5, 38, 3, 2, 2, 2,
	7, 42, 3, 2, 2, 2, 9, 46, 3, 2, 2, 2, 11, 48, 3, 2, 2, 2, 13, 50, 3, 2,
	2, 2, 15, 52, 3, 2, 2, 2, 17, 54, 3, 2, 2, 2, 19, 56, 3, 2, 2, 2, 21, 58,
	3, 2, 2, 2, 23, 60, 3, 2, 2, 2, 25, 74, 3, 2, 2, 2, 27, 76, 3, 2, 2, 2,
	29, 87, 3, 2, 2, 2, 31, 95, 3, 2, 2, 2, 33, 98, 3, 2, 2, 2, 35, 36, 7,
	49, 2, 2, 36, 37, 7, 49, 2, 2, 37, 4, 3, 2, 2, 2, 38, 39, 7, 35, 2, 2,
	39, 40, 7, 38, 2, 2, 40, 41, 7, 44, 2, 2, 41, 6, 3, 2, 2, 2, 42, 43, 7,
	44, 2, 2, 43, 44, 7, 38, 2, 2, 44, 45, 7, 35, 2, 2, 45, 8, 3, 2, 2, 2,
	46, 47, 7, 125, 2, 2, 47, 10, 3, 2, 2, 2, 48, 49, 7, 63, 2, 2, 49, 12,
	3, 2, 2, 2, 50, 51, 7, 61, 2, 2, 51, 14, 3, 2, 2, 2, 52, 53, 7, 127, 2,
	2, 53, 16, 3, 2, 2, 2, 54, 55, 7, 42, 2, 2, 55, 18, 3, 2, 2, 2, 56, 57,
	7, 46, 2, 2, 57, 20, 3, 2, 2, 2, 58, 59, 7, 43, 2, 2, 59, 22, 3, 2, 2,
	2, 60, 61, 7, 49, 2, 2, 61, 62, 7, 44, 2, 2, 62, 66, 3, 2, 2, 2, 63, 65,
	11, 2, 2, 2, 64, 63, 3, 2, 2, 2, 65, 68, 3, 2, 2, 2, 66, 67, 3, 2, 2, 2,
	66, 64, 3, 2, 2, 2, 67, 69, 3, 2, 2, 2, 68, 66, 3, 2, 2, 2, 69, 70, 7,
	44, 2, 2, 70, 71, 7, 49, 2, 2, 71, 24, 3, 2, 2, 2, 72, 75, 5, 27, 14, 2,
	73, 75, 5, 29, 15, 2, 74, 72, 3, 2, 2, 2, 74, 73, 3, 2, 2, 2, 75, 26, 3,
	2, 2, 2, 76, 81, 7, 36, 2, 2, 77, 80, 5, 31, 16, 2, 78, 80, 11, 2, 2, 2,
	79, 77, 3, 2, 2, 2, 79, 78, 3, 2, 2, 2, 80, 83, 3, 2, 2, 2, 81, 82, 3,
	2, 2, 2, 81, 79, 3, 2, 2, 2, 82, 84, 3, 2, 2, 2, 83, 81, 3, 2, 2, 2, 84,
	85, 7, 36, 2, 2, 85, 28, 3, 2, 2, 2, 86, 88, 9, 2, 2, 2, 87, 86, 3, 2,
	2, 2, 88, 89, 3, 2, 2, 2, 89, 87, 3, 2, 2, 2, 89, 90, 3, 2, 2, 2, 90, 30,
	3, 2, 2, 2, 91, 92, 7, 94, 2, 2, 92, 96, 7, 36, 2, 2, 93, 94, 7, 94, 2,
	2, 94, 96, 7, 94, 2, 2, 95, 91, 3, 2, 2, 2, 95, 93, 3, 2, 2, 2, 96, 32,
	3, 2, 2, 2, 97, 99, 9, 3, 2, 2, 98, 97, 3, 2, 2, 2, 99, 100, 3, 2, 2, 2,
	100, 98, 3, 2, 2, 2, 100, 101, 3, 2, 2, 2, 101, 102, 3, 2, 2, 2, 102, 103,
	8, 17, 2, 2, 103, 34, 3, 2, 2, 2, 10, 2, 66, 74, 79, 81, 89, 95, 100, 3,
	8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'//'", "'!$*'", "'*$!'", "'{'", "'='", "';'", "'}'", "'('", "','",
	"')'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "Comment", "NString", "WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"T__9", "Comment", "NString", "String1", "String2", "ESC", "WS",
}

type PBXProjLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewPBXProjLexer(input antlr.CharStream) *PBXProjLexer {

	l := new(PBXProjLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "PBXProj.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// PBXProjLexer tokens.
const (
	PBXProjLexerT__0    = 1
	PBXProjLexerT__1    = 2
	PBXProjLexerT__2    = 3
	PBXProjLexerT__3    = 4
	PBXProjLexerT__4    = 5
	PBXProjLexerT__5    = 6
	PBXProjLexerT__6    = 7
	PBXProjLexerT__7    = 8
	PBXProjLexerT__8    = 9
	PBXProjLexerT__9    = 10
	PBXProjLexerComment = 11
	PBXProjLexerNString = 12
	PBXProjLexerWS      = 13
)
