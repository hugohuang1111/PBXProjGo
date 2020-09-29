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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 17, 90, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 3, 2, 3, 2, 5, 2, 23, 10, 2, 3, 2, 3,
	2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 6, 4, 34, 10, 4, 13, 4, 14,
	4, 35, 3, 4, 5, 4, 39, 10, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 7, 5,
	47, 10, 5, 12, 5, 14, 5, 50, 11, 5, 3, 5, 5, 5, 53, 10, 5, 3, 5, 3, 5,
	3, 6, 3, 6, 5, 6, 59, 10, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 5, 8, 66, 10,
	8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 6, 9, 74, 10, 9, 13, 9, 14, 9, 75,
	3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 7, 10, 85, 10, 10, 12, 10,
	14, 10, 88, 11, 10, 3, 10, 2, 2, 11, 2, 4, 6, 8, 10, 12, 14, 16, 18, 2,
	2, 2, 90, 2, 20, 3, 2, 2, 2, 4, 26, 3, 2, 2, 2, 6, 31, 3, 2, 2, 2, 8, 42,
	3, 2, 2, 2, 10, 56, 3, 2, 2, 2, 12, 60, 3, 2, 2, 2, 14, 65, 3, 2, 2, 2,
	16, 67, 3, 2, 2, 2, 18, 86, 3, 2, 2, 2, 20, 22, 5, 4, 3, 2, 21, 23, 7,
	15, 2, 2, 22, 21, 3, 2, 2, 2, 22, 23, 3, 2, 2, 2, 23, 24, 3, 2, 2, 2, 24,
	25, 5, 6, 4, 2, 25, 3, 3, 2, 2, 2, 26, 27, 7, 3, 2, 2, 27, 28, 7, 4, 2,
	2, 28, 29, 7, 16, 2, 2, 29, 30, 7, 5, 2, 2, 30, 5, 3, 2, 2, 2, 31, 38,
	7, 6, 2, 2, 32, 34, 5, 16, 9, 2, 33, 32, 3, 2, 2, 2, 34, 35, 3, 2, 2, 2,
	35, 33, 3, 2, 2, 2, 35, 36, 3, 2, 2, 2, 36, 39, 3, 2, 2, 2, 37, 39, 5,
	18, 10, 2, 38, 33, 3, 2, 2, 2, 38, 37, 3, 2, 2, 2, 39, 40, 3, 2, 2, 2,
	40, 41, 7, 7, 2, 2, 41, 7, 3, 2, 2, 2, 42, 48, 7, 8, 2, 2, 43, 44, 5, 14,
	8, 2, 44, 45, 7, 9, 2, 2, 45, 47, 3, 2, 2, 2, 46, 43, 3, 2, 2, 2, 47, 50,
	3, 2, 2, 2, 48, 46, 3, 2, 2, 2, 48, 49, 3, 2, 2, 2, 49, 52, 3, 2, 2, 2,
	50, 48, 3, 2, 2, 2, 51, 53, 5, 14, 8, 2, 52, 51, 3, 2, 2, 2, 52, 53, 3,
	2, 2, 2, 53, 54, 3, 2, 2, 2, 54, 55, 7, 10, 2, 2, 55, 9, 3, 2, 2, 2, 56,
	58, 7, 16, 2, 2, 57, 59, 7, 15, 2, 2, 58, 57, 3, 2, 2, 2, 58, 59, 3, 2,
	2, 2, 59, 11, 3, 2, 2, 2, 60, 61, 5, 10, 6, 2, 61, 13, 3, 2, 2, 2, 62,
	66, 5, 8, 5, 2, 63, 66, 5, 6, 4, 2, 64, 66, 5, 10, 6, 2, 65, 62, 3, 2,
	2, 2, 65, 63, 3, 2, 2, 2, 65, 64, 3, 2, 2, 2, 66, 15, 3, 2, 2, 2, 67, 73,
	7, 13, 2, 2, 68, 69, 5, 12, 7, 2, 69, 70, 7, 11, 2, 2, 70, 71, 5, 14, 8,
	2, 71, 72, 7, 12, 2, 2, 72, 74, 3, 2, 2, 2, 73, 68, 3, 2, 2, 2, 74, 75,
	3, 2, 2, 2, 75, 73, 3, 2, 2, 2, 75, 76, 3, 2, 2, 2, 76, 77, 3, 2, 2, 2,
	77, 78, 7, 14, 2, 2, 78, 17, 3, 2, 2, 2, 79, 80, 5, 12, 7, 2, 80, 81, 7,
	11, 2, 2, 81, 82, 5, 14, 8, 2, 82, 83, 7, 12, 2, 2, 83, 85, 3, 2, 2, 2,
	84, 79, 3, 2, 2, 2, 85, 88, 3, 2, 2, 2, 86, 84, 3, 2, 2, 2, 86, 87, 3,
	2, 2, 2, 87, 19, 3, 2, 2, 2, 88, 86, 3, 2, 2, 2, 11, 22, 35, 38, 48, 52,
	58, 65, 75, 86,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'//'", "'!$*'", "'*$!'", "'{'", "'}'", "'('", "','", "')'", "'='",
	"';'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "SectionB", "SectionE", "Comment",
	"NString", "WS",
}

var ruleNames = []string{
	"project", "fileEncodeInfo", "valMap", "valArray", "valString", "key",
	"value", "sectionName", "sectionNoName",
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
	PBXProjParserEOF      = antlr.TokenEOF
	PBXProjParserT__0     = 1
	PBXProjParserT__1     = 2
	PBXProjParserT__2     = 3
	PBXProjParserT__3     = 4
	PBXProjParserT__4     = 5
	PBXProjParserT__5     = 6
	PBXProjParserT__6     = 7
	PBXProjParserT__7     = 8
	PBXProjParserT__8     = 9
	PBXProjParserT__9     = 10
	PBXProjParserSectionB = 11
	PBXProjParserSectionE = 12
	PBXProjParserComment  = 13
	PBXProjParserNString  = 14
	PBXProjParserWS       = 15
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
	PBXProjParserRULE_sectionName    = 7
	PBXProjParserRULE_sectionNoName  = 8
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
		p.SetState(18)
		p.FileEncodeInfo()
	}
	p.SetState(20)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == PBXProjParserComment {
		{
			p.SetState(19)
			p.Match(PBXProjParserComment)
		}

	}
	{
		p.SetState(22)
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
		p.SetState(24)
		p.Match(PBXProjParserT__0)
	}
	{
		p.SetState(25)
		p.Match(PBXProjParserT__1)
	}
	{
		p.SetState(26)
		p.Match(PBXProjParserNString)
	}
	{
		p.SetState(27)
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

func (s *ValMapContext) SectionNoName() ISectionNoNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISectionNoNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISectionNoNameContext)
}

func (s *ValMapContext) AllSectionName() []ISectionNameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISectionNameContext)(nil)).Elem())
	var tst = make([]ISectionNameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISectionNameContext)
		}
	}

	return tst
}

func (s *ValMapContext) SectionName(i int) ISectionNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISectionNameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISectionNameContext)
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
		p.SetState(29)
		p.Match(PBXProjParserT__3)
	}
	p.SetState(36)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case PBXProjParserSectionB:
		p.SetState(31)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == PBXProjParserSectionB {
			{
				p.SetState(30)
				p.SectionName()
			}

			p.SetState(33)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case PBXProjParserT__4, PBXProjParserNString:
		{
			p.SetState(35)
			p.SectionNoName()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(38)
		p.Match(PBXProjParserT__4)
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
		p.SetState(40)
		p.Match(PBXProjParserT__5)
	}
	p.SetState(46)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(41)
				p.Value()
			}
			{
				p.SetState(42)
				p.Match(PBXProjParserT__6)
			}

		}
		p.SetState(48)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
	}
	p.SetState(50)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<PBXProjParserT__3)|(1<<PBXProjParserT__5)|(1<<PBXProjParserNString))) != 0 {
		{
			p.SetState(49)
			p.Value()
		}

	}
	{
		p.SetState(52)
		p.Match(PBXProjParserT__7)
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
		p.SetState(54)
		p.Match(PBXProjParserNString)
	}
	p.SetState(56)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == PBXProjParserComment {
		{
			p.SetState(55)
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
		p.SetState(58)
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

	p.SetState(63)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case PBXProjParserT__5:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(60)
			p.ValArray()
		}

	case PBXProjParserT__3:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(61)
			p.ValMap()
		}

	case PBXProjParserNString:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(62)
			p.ValString()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ISectionNameContext is an interface to support dynamic dispatch.
type ISectionNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSectionNameContext differentiates from other interfaces.
	IsSectionNameContext()
}

type SectionNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySectionNameContext() *SectionNameContext {
	var p = new(SectionNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PBXProjParserRULE_sectionName
	return p
}

func (*SectionNameContext) IsSectionNameContext() {}

func NewSectionNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SectionNameContext {
	var p = new(SectionNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PBXProjParserRULE_sectionName

	return p
}

func (s *SectionNameContext) GetParser() antlr.Parser { return s.parser }

func (s *SectionNameContext) SectionB() antlr.TerminalNode {
	return s.GetToken(PBXProjParserSectionB, 0)
}

func (s *SectionNameContext) SectionE() antlr.TerminalNode {
	return s.GetToken(PBXProjParserSectionE, 0)
}

func (s *SectionNameContext) AllKey() []IKeyContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IKeyContext)(nil)).Elem())
	var tst = make([]IKeyContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IKeyContext)
		}
	}

	return tst
}

func (s *SectionNameContext) Key(i int) IKeyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKeyContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IKeyContext)
}

func (s *SectionNameContext) AllValue() []IValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IValueContext)(nil)).Elem())
	var tst = make([]IValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IValueContext)
		}
	}

	return tst
}

func (s *SectionNameContext) Value(i int) IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *SectionNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SectionNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SectionNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PBXProjListener); ok {
		listenerT.EnterSectionName(s)
	}
}

func (s *SectionNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PBXProjListener); ok {
		listenerT.ExitSectionName(s)
	}
}

func (p *PBXProjParser) SectionName() (localctx ISectionNameContext) {
	localctx = NewSectionNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, PBXProjParserRULE_sectionName)
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
		p.SetState(65)
		p.Match(PBXProjParserSectionB)
	}
	p.SetState(71)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == PBXProjParserNString {
		{
			p.SetState(66)
			p.Key()
		}
		{
			p.SetState(67)
			p.Match(PBXProjParserT__8)
		}
		{
			p.SetState(68)
			p.Value()
		}
		{
			p.SetState(69)
			p.Match(PBXProjParserT__9)
		}

		p.SetState(73)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(75)
		p.Match(PBXProjParserSectionE)
	}

	return localctx
}

// ISectionNoNameContext is an interface to support dynamic dispatch.
type ISectionNoNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSectionNoNameContext differentiates from other interfaces.
	IsSectionNoNameContext()
}

type SectionNoNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySectionNoNameContext() *SectionNoNameContext {
	var p = new(SectionNoNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PBXProjParserRULE_sectionNoName
	return p
}

func (*SectionNoNameContext) IsSectionNoNameContext() {}

func NewSectionNoNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SectionNoNameContext {
	var p = new(SectionNoNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PBXProjParserRULE_sectionNoName

	return p
}

func (s *SectionNoNameContext) GetParser() antlr.Parser { return s.parser }

func (s *SectionNoNameContext) AllKey() []IKeyContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IKeyContext)(nil)).Elem())
	var tst = make([]IKeyContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IKeyContext)
		}
	}

	return tst
}

func (s *SectionNoNameContext) Key(i int) IKeyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKeyContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IKeyContext)
}

func (s *SectionNoNameContext) AllValue() []IValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IValueContext)(nil)).Elem())
	var tst = make([]IValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IValueContext)
		}
	}

	return tst
}

func (s *SectionNoNameContext) Value(i int) IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *SectionNoNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SectionNoNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SectionNoNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PBXProjListener); ok {
		listenerT.EnterSectionNoName(s)
	}
}

func (s *SectionNoNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PBXProjListener); ok {
		listenerT.ExitSectionNoName(s)
	}
}

func (p *PBXProjParser) SectionNoName() (localctx ISectionNoNameContext) {
	localctx = NewSectionNoNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, PBXProjParserRULE_sectionNoName)
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
	p.SetState(84)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == PBXProjParserNString {
		{
			p.SetState(77)
			p.Key()
		}
		{
			p.SetState(78)
			p.Match(PBXProjParserT__8)
		}
		{
			p.SetState(79)
			p.Value()
		}
		{
			p.SetState(80)
			p.Match(PBXProjParserT__9)
		}

		p.SetState(86)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}
