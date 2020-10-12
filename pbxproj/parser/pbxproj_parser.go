// Code generated from PBXProj.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // PBXProj

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 15, 71, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 3, 2, 3, 2, 5, 2, 19, 10, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 4, 3, 4, 5, 4, 30, 10, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5,
	4, 37, 10, 4, 7, 4, 39, 10, 4, 12, 4, 14, 4, 42, 11, 4, 3, 4, 3, 4, 3,
	5, 3, 5, 3, 5, 3, 5, 7, 5, 50, 10, 5, 12, 5, 14, 5, 53, 11, 5, 3, 5, 5,
	5, 56, 10, 5, 3, 5, 3, 5, 3, 6, 3, 6, 5, 6, 62, 10, 6, 3, 7, 3, 7, 3, 8,
	3, 8, 3, 8, 5, 8, 69, 10, 8, 3, 8, 2, 2, 9, 2, 4, 6, 8, 10, 12, 14, 2,
	2, 2, 72, 2, 16, 3, 2, 2, 2, 4, 22, 3, 2, 2, 2, 6, 27, 3, 2, 2, 2, 8, 45,
	3, 2, 2, 2, 10, 59, 3, 2, 2, 2, 12, 63, 3, 2, 2, 2, 14, 68, 3, 2, 2, 2,
	16, 18, 5, 4, 3, 2, 17, 19, 7, 13, 2, 2, 18, 17, 3, 2, 2, 2, 18, 19, 3,
	2, 2, 2, 19, 20, 3, 2, 2, 2, 20, 21, 5, 6, 4, 2, 21, 3, 3, 2, 2, 2, 22,
	23, 7, 3, 2, 2, 23, 24, 7, 4, 2, 2, 24, 25, 7, 14, 2, 2, 25, 26, 7, 5,
	2, 2, 26, 5, 3, 2, 2, 2, 27, 40, 7, 6, 2, 2, 28, 30, 7, 13, 2, 2, 29, 28,
	3, 2, 2, 2, 29, 30, 3, 2, 2, 2, 30, 31, 3, 2, 2, 2, 31, 32, 5, 12, 7, 2,
	32, 33, 7, 7, 2, 2, 33, 34, 5, 14, 8, 2, 34, 36, 7, 8, 2, 2, 35, 37, 7,
	13, 2, 2, 36, 35, 3, 2, 2, 2, 36, 37, 3, 2, 2, 2, 37, 39, 3, 2, 2, 2, 38,
	29, 3, 2, 2, 2, 39, 42, 3, 2, 2, 2, 40, 38, 3, 2, 2, 2, 40, 41, 3, 2, 2,
	2, 41, 43, 3, 2, 2, 2, 42, 40, 3, 2, 2, 2, 43, 44, 7, 9, 2, 2, 44, 7, 3,
	2, 2, 2, 45, 51, 7, 10, 2, 2, 46, 47, 5, 14, 8, 2, 47, 48, 7, 11, 2, 2,
	48, 50, 3, 2, 2, 2, 49, 46, 3, 2, 2, 2, 50, 53, 3, 2, 2, 2, 51, 49, 3,
	2, 2, 2, 51, 52, 3, 2, 2, 2, 52, 55, 3, 2, 2, 2, 53, 51, 3, 2, 2, 2, 54,
	56, 5, 14, 8, 2, 55, 54, 3, 2, 2, 2, 55, 56, 3, 2, 2, 2, 56, 57, 3, 2,
	2, 2, 57, 58, 7, 12, 2, 2, 58, 9, 3, 2, 2, 2, 59, 61, 7, 14, 2, 2, 60,
	62, 7, 13, 2, 2, 61, 60, 3, 2, 2, 2, 61, 62, 3, 2, 2, 2, 62, 11, 3, 2,
	2, 2, 63, 64, 5, 10, 6, 2, 64, 13, 3, 2, 2, 2, 65, 69, 5, 8, 5, 2, 66,
	69, 5, 6, 4, 2, 67, 69, 5, 10, 6, 2, 68, 65, 3, 2, 2, 2, 68, 66, 3, 2,
	2, 2, 68, 67, 3, 2, 2, 2, 69, 15, 3, 2, 2, 2, 10, 18, 29, 36, 40, 51, 55,
	61, 68,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'//'", "'!$*'", "'*$!'", "'{'", "'='", "';'", "'}'", "'('", "','",
	"')'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "Comment", "NString", "WS",
}

var ruleNames = []string{
	"project", "fileEncodeInfo", "valMap", "valArray", "valString", "key",
	"value",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type PBXProjParser struct {
	*antlr.BaseParser
}

func NewPBXProjParser(input antlr.TokenStream) *PBXProjParser {
	this := new(PBXProjParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "PBXProj.g4"

	return this
}

// PBXProjParser tokens.
const (
	PBXProjParserEOF     = antlr.TokenEOF
	PBXProjParserT__0    = 1
	PBXProjParserT__1    = 2
	PBXProjParserT__2    = 3
	PBXProjParserT__3    = 4
	PBXProjParserT__4    = 5
	PBXProjParserT__5    = 6
	PBXProjParserT__6    = 7
	PBXProjParserT__7    = 8
	PBXProjParserT__8    = 9
	PBXProjParserT__9    = 10
	PBXProjParserComment = 11
	PBXProjParserNString = 12
	PBXProjParserWS      = 13
)

// PBXProjParser rules.
const (
	PBXProjParserRULE_project        = 0
	PBXProjParserRULE_fileEncodeInfo = 1
	PBXProjParserRULE_valMap         = 2
	PBXProjParserRULE_valArray       = 3
	PBXProjParserRULE_valString      = 4
	PBXProjParserRULE_key            = 5
	PBXProjParserRULE_value          = 6
)

// IProjectContext is an interface to support dynamic dispatch.
type IProjectContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProjectContext differentiates from other interfaces.
	IsProjectContext()
}

type ProjectContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProjectContext() *ProjectContext {
	var p = new(ProjectContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PBXProjParserRULE_project
	return p
}

func (*ProjectContext) IsProjectContext() {}

func NewProjectContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProjectContext {
	var p = new(ProjectContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PBXProjParserRULE_project

	return p
}

func (s *ProjectContext) GetParser() antlr.Parser { return s.parser }

func (s *ProjectContext) FileEncodeInfo() IFileEncodeInfoContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFileEncodeInfoContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFileEncodeInfoContext)
}

func (s *ProjectContext) ValMap() IValMapContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValMapContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValMapContext)
}

func (s *ProjectContext) Comment() antlr.TerminalNode {
	return s.GetToken(PBXProjParserComment, 0)
}

func (s *ProjectContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProjectContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProjectContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PBXProjListener); ok {
		listenerT.EnterProject(s)
	}
}

func (s *ProjectContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PBXProjListener); ok {
		listenerT.ExitProject(s)
	}
}

func (p *PBXProjParser) Project() (localctx IProjectContext) {
	localctx = NewProjectContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, PBXProjParserRULE_project)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(14)
		p.FileEncodeInfo()
	}
	p.SetState(16)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == PBXProjParserComment {
		{
			p.SetState(15)
			p.Match(PBXProjParserComment)
		}

	}
	{
		p.SetState(18)
		p.ValMap()
	}

	return localctx
}

// IFileEncodeInfoContext is an interface to support dynamic dispatch.
type IFileEncodeInfoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFileEncodeInfoContext differentiates from other interfaces.
	IsFileEncodeInfoContext()
}

type FileEncodeInfoContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFileEncodeInfoContext() *FileEncodeInfoContext {
	var p = new(FileEncodeInfoContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PBXProjParserRULE_fileEncodeInfo
	return p
}

func (*FileEncodeInfoContext) IsFileEncodeInfoContext() {}

func NewFileEncodeInfoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FileEncodeInfoContext {
	var p = new(FileEncodeInfoContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PBXProjParserRULE_fileEncodeInfo

	return p
}

func (s *FileEncodeInfoContext) GetParser() antlr.Parser { return s.parser }

func (s *FileEncodeInfoContext) NString() antlr.TerminalNode {
	return s.GetToken(PBXProjParserNString, 0)
}

func (s *FileEncodeInfoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FileEncodeInfoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FileEncodeInfoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PBXProjListener); ok {
		listenerT.EnterFileEncodeInfo(s)
	}
}

func (s *FileEncodeInfoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PBXProjListener); ok {
		listenerT.ExitFileEncodeInfo(s)
	}
}

func (p *PBXProjParser) FileEncodeInfo() (localctx IFileEncodeInfoContext) {
	localctx = NewFileEncodeInfoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, PBXProjParserRULE_fileEncodeInfo)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(20)
		p.Match(PBXProjParserT__0)
	}
	{
		p.SetState(21)
		p.Match(PBXProjParserT__1)
	}
	{
		p.SetState(22)
		p.Match(PBXProjParserNString)
	}
	{
		p.SetState(23)
		p.Match(PBXProjParserT__2)
	}

	return localctx
}

// IValMapContext is an interface to support dynamic dispatch.
type IValMapContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValMapContext differentiates from other interfaces.
	IsValMapContext()
}

type ValMapContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValMapContext() *ValMapContext {
	var p = new(ValMapContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PBXProjParserRULE_valMap
	return p
}

func (*ValMapContext) IsValMapContext() {}

func NewValMapContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValMapContext {
	var p = new(ValMapContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PBXProjParserRULE_valMap

	return p
}

func (s *ValMapContext) GetParser() antlr.Parser { return s.parser }

func (s *ValMapContext) AllKey() []IKeyContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IKeyContext)(nil)).Elem())
	var tst = make([]IKeyContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IKeyContext)
		}
	}

	return tst
}

func (s *ValMapContext) Key(i int) IKeyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKeyContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IKeyContext)
}

func (s *ValMapContext) AllValue() []IValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IValueContext)(nil)).Elem())
	var tst = make([]IValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IValueContext)
		}
	}

	return tst
}

func (s *ValMapContext) Value(i int) IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *ValMapContext) AllComment() []antlr.TerminalNode {
	return s.GetTokens(PBXProjParserComment)
}

func (s *ValMapContext) Comment(i int) antlr.TerminalNode {
	return s.GetToken(PBXProjParserComment, i)
}

func (s *ValMapContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValMapContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValMapContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PBXProjListener); ok {
		listenerT.EnterValMap(s)
	}
}

func (s *ValMapContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PBXProjListener); ok {
		listenerT.ExitValMap(s)
	}
}

func (p *PBXProjParser) ValMap() (localctx IValMapContext) {
	localctx = NewValMapContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, PBXProjParserRULE_valMap)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(25)
		p.Match(PBXProjParserT__3)
	}
	p.SetState(38)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == PBXProjParserComment || _la == PBXProjParserNString {
		p.SetState(27)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == PBXProjParserComment {
			{
				p.SetState(26)
				p.Match(PBXProjParserComment)
			}

		}
		{
			p.SetState(29)
			p.Key()
		}
		{
			p.SetState(30)
			p.Match(PBXProjParserT__4)
		}
		{
			p.SetState(31)
			p.Value()
		}
		{
			p.SetState(32)
			p.Match(PBXProjParserT__5)
		}
		p.SetState(34)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(33)
				p.Match(PBXProjParserComment)
			}

		}

		p.SetState(40)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(41)
		p.Match(PBXProjParserT__6)
	}

	return localctx
}

// IValArrayContext is an interface to support dynamic dispatch.
type IValArrayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValArrayContext differentiates from other interfaces.
	IsValArrayContext()
}

type ValArrayContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValArrayContext() *ValArrayContext {
	var p = new(ValArrayContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PBXProjParserRULE_valArray
	return p
}

func (*ValArrayContext) IsValArrayContext() {}

func NewValArrayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValArrayContext {
	var p = new(ValArrayContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PBXProjParserRULE_valArray

	return p
}

func (s *ValArrayContext) GetParser() antlr.Parser { return s.parser }

func (s *ValArrayContext) AllValue() []IValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IValueContext)(nil)).Elem())
	var tst = make([]IValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IValueContext)
		}
	}

	return tst
}

func (s *ValArrayContext) Value(i int) IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *ValArrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValArrayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValArrayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PBXProjListener); ok {
		listenerT.EnterValArray(s)
	}
}

func (s *ValArrayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PBXProjListener); ok {
		listenerT.ExitValArray(s)
	}
}

func (p *PBXProjParser) ValArray() (localctx IValArrayContext) {
	localctx = NewValArrayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, PBXProjParserRULE_valArray)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(43)
		p.Match(PBXProjParserT__7)
	}
	p.SetState(49)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(44)
				p.Value()
			}
			{
				p.SetState(45)
				p.Match(PBXProjParserT__8)
			}

		}
		p.SetState(51)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())
	}
	p.SetState(53)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<PBXProjParserT__3)|(1<<PBXProjParserT__7)|(1<<PBXProjParserNString))) != 0 {
		{
			p.SetState(52)
			p.Value()
		}

	}
	{
		p.SetState(55)
		p.Match(PBXProjParserT__9)
	}

	return localctx
}

// IValStringContext is an interface to support dynamic dispatch.
type IValStringContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValStringContext differentiates from other interfaces.
	IsValStringContext()
}

type ValStringContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValStringContext() *ValStringContext {
	var p = new(ValStringContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PBXProjParserRULE_valString
	return p
}

func (*ValStringContext) IsValStringContext() {}

func NewValStringContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValStringContext {
	var p = new(ValStringContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PBXProjParserRULE_valString

	return p
}

func (s *ValStringContext) GetParser() antlr.Parser { return s.parser }

func (s *ValStringContext) NString() antlr.TerminalNode {
	return s.GetToken(PBXProjParserNString, 0)
}

func (s *ValStringContext) Comment() antlr.TerminalNode {
	return s.GetToken(PBXProjParserComment, 0)
}

func (s *ValStringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValStringContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValStringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PBXProjListener); ok {
		listenerT.EnterValString(s)
	}
}

func (s *ValStringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PBXProjListener); ok {
		listenerT.ExitValString(s)
	}
}

func (p *PBXProjParser) ValString() (localctx IValStringContext) {
	localctx = NewValStringContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, PBXProjParserRULE_valString)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(57)
		p.Match(PBXProjParserNString)
	}
	p.SetState(59)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == PBXProjParserComment {
		{
			p.SetState(58)
			p.Match(PBXProjParserComment)
		}

	}

	return localctx
}

// IKeyContext is an interface to support dynamic dispatch.
type IKeyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKeyContext differentiates from other interfaces.
	IsKeyContext()
}

type KeyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeyContext() *KeyContext {
	var p = new(KeyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PBXProjParserRULE_key
	return p
}

func (*KeyContext) IsKeyContext() {}

func NewKeyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeyContext {
	var p = new(KeyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PBXProjParserRULE_key

	return p
}

func (s *KeyContext) GetParser() antlr.Parser { return s.parser }

func (s *KeyContext) ValString() IValStringContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValStringContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValStringContext)
}

func (s *KeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PBXProjListener); ok {
		listenerT.EnterKey(s)
	}
}

func (s *KeyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PBXProjListener); ok {
		listenerT.ExitKey(s)
	}
}

func (p *PBXProjParser) Key() (localctx IKeyContext) {
	localctx = NewKeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, PBXProjParserRULE_key)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(61)
		p.ValString()
	}

	return localctx
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PBXProjParserRULE_value
	return p
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PBXProjParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) ValArray() IValArrayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValArrayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValArrayContext)
}

func (s *ValueContext) ValMap() IValMapContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValMapContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValMapContext)
}

func (s *ValueContext) ValString() IValStringContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValStringContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValStringContext)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PBXProjListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PBXProjListener); ok {
		listenerT.ExitValue(s)
	}
}

func (p *PBXProjParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, PBXProjParserRULE_value)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(66)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case PBXProjParserT__7:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(63)
			p.ValArray()
		}

	case PBXProjParserT__3:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(64)
			p.ValMap()
		}

	case PBXProjParserNString:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(65)
			p.ValString()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}
