// Code generated from Tebot.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // Tebot

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type TebotParser struct {
	*antlr.BaseParser
}

var TebotParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func tebotParserInit() {
	staticData := &TebotParserStaticData
	staticData.LiteralNames = []string{
		"", "'click'", "'input'", "'select'", "'capture'", "'assert'", "'sleep'",
		"'goto'", "'move'", "'scroll'", "'save'", "'paste'", "'='", "'('", "')'",
		"','",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "TEBOT_WHITESPACE",
		"TEBOT_COMMENT", "TEBOT_QUOTED_STRING",
	}
	staticData.RuleNames = []string{
		"tebot_selector", "tebot_value", "tebot_action", "tebot_assign", "tebot_operation",
		"tebot_assert", "tebot_operations",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 18, 54, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1,
		3, 1, 3, 3, 3, 24, 8, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 32,
		8, 4, 1, 4, 1, 4, 1, 4, 3, 4, 37, 8, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 5, 1, 6, 5, 6, 49, 8, 6, 10, 6, 12, 6, 52, 9, 6, 1, 6,
		0, 0, 7, 0, 2, 4, 6, 8, 10, 12, 0, 1, 1, 0, 1, 11, 50, 0, 14, 1, 0, 0,
		0, 2, 16, 1, 0, 0, 0, 4, 18, 1, 0, 0, 0, 6, 23, 1, 0, 0, 0, 8, 36, 1, 0,
		0, 0, 10, 38, 1, 0, 0, 0, 12, 50, 1, 0, 0, 0, 14, 15, 5, 18, 0, 0, 15,
		1, 1, 0, 0, 0, 16, 17, 5, 18, 0, 0, 17, 3, 1, 0, 0, 0, 18, 19, 7, 0, 0,
		0, 19, 5, 1, 0, 0, 0, 20, 21, 3, 0, 0, 0, 21, 22, 5, 12, 0, 0, 22, 24,
		1, 0, 0, 0, 23, 20, 1, 0, 0, 0, 23, 24, 1, 0, 0, 0, 24, 25, 1, 0, 0, 0,
		25, 26, 3, 2, 1, 0, 26, 7, 1, 0, 0, 0, 27, 28, 3, 4, 2, 0, 28, 31, 5, 13,
		0, 0, 29, 32, 3, 0, 0, 0, 30, 32, 3, 6, 3, 0, 31, 29, 1, 0, 0, 0, 31, 30,
		1, 0, 0, 0, 32, 33, 1, 0, 0, 0, 33, 34, 5, 14, 0, 0, 34, 37, 1, 0, 0, 0,
		35, 37, 3, 10, 5, 0, 36, 27, 1, 0, 0, 0, 36, 35, 1, 0, 0, 0, 37, 9, 1,
		0, 0, 0, 38, 39, 5, 5, 0, 0, 39, 40, 5, 13, 0, 0, 40, 41, 3, 2, 1, 0, 41,
		42, 5, 15, 0, 0, 42, 43, 3, 2, 1, 0, 43, 44, 5, 15, 0, 0, 44, 45, 3, 2,
		1, 0, 45, 46, 5, 14, 0, 0, 46, 11, 1, 0, 0, 0, 47, 49, 3, 8, 4, 0, 48,
		47, 1, 0, 0, 0, 49, 52, 1, 0, 0, 0, 50, 48, 1, 0, 0, 0, 50, 51, 1, 0, 0,
		0, 51, 13, 1, 0, 0, 0, 52, 50, 1, 0, 0, 0, 4, 23, 31, 36, 50,
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

// TebotParserInit initializes any static state used to implement TebotParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewTebotParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func TebotParserInit() {
	staticData := &TebotParserStaticData
	staticData.once.Do(tebotParserInit)
}

// NewTebotParser produces a new parser instance for the optional input antlr.TokenStream.
func NewTebotParser(input antlr.TokenStream) *TebotParser {
	TebotParserInit()
	this := new(TebotParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &TebotParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Tebot.g4"

	return this
}

// TebotParser tokens.
const (
	TebotParserEOF                 = antlr.TokenEOF
	TebotParserT__0                = 1
	TebotParserT__1                = 2
	TebotParserT__2                = 3
	TebotParserT__3                = 4
	TebotParserT__4                = 5
	TebotParserT__5                = 6
	TebotParserT__6                = 7
	TebotParserT__7                = 8
	TebotParserT__8                = 9
	TebotParserT__9                = 10
	TebotParserT__10               = 11
	TebotParserT__11               = 12
	TebotParserT__12               = 13
	TebotParserT__13               = 14
	TebotParserT__14               = 15
	TebotParserTEBOT_WHITESPACE    = 16
	TebotParserTEBOT_COMMENT       = 17
	TebotParserTEBOT_QUOTED_STRING = 18
)

// TebotParser rules.
const (
	TebotParserRULE_tebot_selector   = 0
	TebotParserRULE_tebot_value      = 1
	TebotParserRULE_tebot_action     = 2
	TebotParserRULE_tebot_assign     = 3
	TebotParserRULE_tebot_operation  = 4
	TebotParserRULE_tebot_assert     = 5
	TebotParserRULE_tebot_operations = 6
)

// ITebot_selectorContext is an interface to support dynamic dispatch.
type ITebot_selectorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TEBOT_QUOTED_STRING() antlr.TerminalNode

	// IsTebot_selectorContext differentiates from other interfaces.
	IsTebot_selectorContext()
}

type Tebot_selectorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTebot_selectorContext() *Tebot_selectorContext {
	var p = new(Tebot_selectorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TebotParserRULE_tebot_selector
	return p
}

func InitEmptyTebot_selectorContext(p *Tebot_selectorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TebotParserRULE_tebot_selector
}

func (*Tebot_selectorContext) IsTebot_selectorContext() {}

func NewTebot_selectorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Tebot_selectorContext {
	var p = new(Tebot_selectorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TebotParserRULE_tebot_selector

	return p
}

func (s *Tebot_selectorContext) GetParser() antlr.Parser { return s.parser }

func (s *Tebot_selectorContext) TEBOT_QUOTED_STRING() antlr.TerminalNode {
	return s.GetToken(TebotParserTEBOT_QUOTED_STRING, 0)
}

func (s *Tebot_selectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Tebot_selectorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Tebot_selectorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TebotListener); ok {
		listenerT.EnterTebot_selector(s)
	}
}

func (s *Tebot_selectorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TebotListener); ok {
		listenerT.ExitTebot_selector(s)
	}
}

func (p *TebotParser) Tebot_selector() (localctx ITebot_selectorContext) {
	localctx = NewTebot_selectorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, TebotParserRULE_tebot_selector)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(14)
		p.Match(TebotParserTEBOT_QUOTED_STRING)
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

// ITebot_valueContext is an interface to support dynamic dispatch.
type ITebot_valueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TEBOT_QUOTED_STRING() antlr.TerminalNode

	// IsTebot_valueContext differentiates from other interfaces.
	IsTebot_valueContext()
}

type Tebot_valueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTebot_valueContext() *Tebot_valueContext {
	var p = new(Tebot_valueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TebotParserRULE_tebot_value
	return p
}

func InitEmptyTebot_valueContext(p *Tebot_valueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TebotParserRULE_tebot_value
}

func (*Tebot_valueContext) IsTebot_valueContext() {}

func NewTebot_valueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Tebot_valueContext {
	var p = new(Tebot_valueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TebotParserRULE_tebot_value

	return p
}

func (s *Tebot_valueContext) GetParser() antlr.Parser { return s.parser }

func (s *Tebot_valueContext) TEBOT_QUOTED_STRING() antlr.TerminalNode {
	return s.GetToken(TebotParserTEBOT_QUOTED_STRING, 0)
}

func (s *Tebot_valueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Tebot_valueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Tebot_valueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TebotListener); ok {
		listenerT.EnterTebot_value(s)
	}
}

func (s *Tebot_valueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TebotListener); ok {
		listenerT.ExitTebot_value(s)
	}
}

func (p *TebotParser) Tebot_value() (localctx ITebot_valueContext) {
	localctx = NewTebot_valueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, TebotParserRULE_tebot_value)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(16)
		p.Match(TebotParserTEBOT_QUOTED_STRING)
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

// ITebot_actionContext is an interface to support dynamic dispatch.
type ITebot_actionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsTebot_actionContext differentiates from other interfaces.
	IsTebot_actionContext()
}

type Tebot_actionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTebot_actionContext() *Tebot_actionContext {
	var p = new(Tebot_actionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TebotParserRULE_tebot_action
	return p
}

func InitEmptyTebot_actionContext(p *Tebot_actionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TebotParserRULE_tebot_action
}

func (*Tebot_actionContext) IsTebot_actionContext() {}

func NewTebot_actionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Tebot_actionContext {
	var p = new(Tebot_actionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TebotParserRULE_tebot_action

	return p
}

func (s *Tebot_actionContext) GetParser() antlr.Parser { return s.parser }
func (s *Tebot_actionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Tebot_actionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Tebot_actionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TebotListener); ok {
		listenerT.EnterTebot_action(s)
	}
}

func (s *Tebot_actionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TebotListener); ok {
		listenerT.ExitTebot_action(s)
	}
}

func (p *TebotParser) Tebot_action() (localctx ITebot_actionContext) {
	localctx = NewTebot_actionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, TebotParserRULE_tebot_action)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(18)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4094) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
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

// ITebot_assignContext is an interface to support dynamic dispatch.
type ITebot_assignContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetSelector returns the selector rule contexts.
	GetSelector() ITebot_selectorContext

	// GetValue returns the value rule contexts.
	GetValue() ITebot_valueContext

	// SetSelector sets the selector rule contexts.
	SetSelector(ITebot_selectorContext)

	// SetValue sets the value rule contexts.
	SetValue(ITebot_valueContext)

	// Getter signatures
	Tebot_value() ITebot_valueContext
	Tebot_selector() ITebot_selectorContext

	// IsTebot_assignContext differentiates from other interfaces.
	IsTebot_assignContext()
}

type Tebot_assignContext struct {
	antlr.BaseParserRuleContext
	parser   antlr.Parser
	selector ITebot_selectorContext
	value    ITebot_valueContext
}

func NewEmptyTebot_assignContext() *Tebot_assignContext {
	var p = new(Tebot_assignContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TebotParserRULE_tebot_assign
	return p
}

func InitEmptyTebot_assignContext(p *Tebot_assignContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TebotParserRULE_tebot_assign
}

func (*Tebot_assignContext) IsTebot_assignContext() {}

func NewTebot_assignContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Tebot_assignContext {
	var p = new(Tebot_assignContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TebotParserRULE_tebot_assign

	return p
}

func (s *Tebot_assignContext) GetParser() antlr.Parser { return s.parser }

func (s *Tebot_assignContext) GetSelector() ITebot_selectorContext { return s.selector }

func (s *Tebot_assignContext) GetValue() ITebot_valueContext { return s.value }

func (s *Tebot_assignContext) SetSelector(v ITebot_selectorContext) { s.selector = v }

func (s *Tebot_assignContext) SetValue(v ITebot_valueContext) { s.value = v }

func (s *Tebot_assignContext) Tebot_value() ITebot_valueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITebot_valueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITebot_valueContext)
}

func (s *Tebot_assignContext) Tebot_selector() ITebot_selectorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITebot_selectorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITebot_selectorContext)
}

func (s *Tebot_assignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Tebot_assignContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Tebot_assignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TebotListener); ok {
		listenerT.EnterTebot_assign(s)
	}
}

func (s *Tebot_assignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TebotListener); ok {
		listenerT.ExitTebot_assign(s)
	}
}

func (p *TebotParser) Tebot_assign() (localctx ITebot_assignContext) {
	localctx = NewTebot_assignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, TebotParserRULE_tebot_assign)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(23)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(20)

			var _x = p.Tebot_selector()

			localctx.(*Tebot_assignContext).selector = _x
		}
		{
			p.SetState(21)
			p.Match(TebotParserT__11)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(25)

		var _x = p.Tebot_value()

		localctx.(*Tebot_assignContext).value = _x
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

// ITebot_operationContext is an interface to support dynamic dispatch.
type ITebot_operationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Tebot_action() ITebot_actionContext
	Tebot_selector() ITebot_selectorContext
	Tebot_assign() ITebot_assignContext
	Tebot_assert() ITebot_assertContext

	// IsTebot_operationContext differentiates from other interfaces.
	IsTebot_operationContext()
}

type Tebot_operationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTebot_operationContext() *Tebot_operationContext {
	var p = new(Tebot_operationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TebotParserRULE_tebot_operation
	return p
}

func InitEmptyTebot_operationContext(p *Tebot_operationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TebotParserRULE_tebot_operation
}

func (*Tebot_operationContext) IsTebot_operationContext() {}

func NewTebot_operationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Tebot_operationContext {
	var p = new(Tebot_operationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TebotParserRULE_tebot_operation

	return p
}

func (s *Tebot_operationContext) GetParser() antlr.Parser { return s.parser }

func (s *Tebot_operationContext) Tebot_action() ITebot_actionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITebot_actionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITebot_actionContext)
}

func (s *Tebot_operationContext) Tebot_selector() ITebot_selectorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITebot_selectorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITebot_selectorContext)
}

func (s *Tebot_operationContext) Tebot_assign() ITebot_assignContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITebot_assignContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITebot_assignContext)
}

func (s *Tebot_operationContext) Tebot_assert() ITebot_assertContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITebot_assertContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITebot_assertContext)
}

func (s *Tebot_operationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Tebot_operationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Tebot_operationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TebotListener); ok {
		listenerT.EnterTebot_operation(s)
	}
}

func (s *Tebot_operationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TebotListener); ok {
		listenerT.ExitTebot_operation(s)
	}
}

func (p *TebotParser) Tebot_operation() (localctx ITebot_operationContext) {
	localctx = NewTebot_operationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, TebotParserRULE_tebot_operation)
	p.SetState(36)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(27)
			p.Tebot_action()
		}
		{
			p.SetState(28)
			p.Match(TebotParserT__12)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(31)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(29)
				p.Tebot_selector()
			}

		case 2:
			{
				p.SetState(30)
				p.Tebot_assign()
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}
		{
			p.SetState(33)
			p.Match(TebotParserT__13)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(35)
			p.Tebot_assert()
		}

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

// ITebot_assertContext is an interface to support dynamic dispatch.
type ITebot_assertContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetDsn returns the dsn rule contexts.
	GetDsn() ITebot_valueContext

	// GetSql returns the sql rule contexts.
	GetSql() ITebot_valueContext

	// GetExpected returns the expected rule contexts.
	GetExpected() ITebot_valueContext

	// SetDsn sets the dsn rule contexts.
	SetDsn(ITebot_valueContext)

	// SetSql sets the sql rule contexts.
	SetSql(ITebot_valueContext)

	// SetExpected sets the expected rule contexts.
	SetExpected(ITebot_valueContext)

	// Getter signatures
	AllTebot_value() []ITebot_valueContext
	Tebot_value(i int) ITebot_valueContext

	// IsTebot_assertContext differentiates from other interfaces.
	IsTebot_assertContext()
}

type Tebot_assertContext struct {
	antlr.BaseParserRuleContext
	parser   antlr.Parser
	dsn      ITebot_valueContext
	sql      ITebot_valueContext
	expected ITebot_valueContext
}

func NewEmptyTebot_assertContext() *Tebot_assertContext {
	var p = new(Tebot_assertContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TebotParserRULE_tebot_assert
	return p
}

func InitEmptyTebot_assertContext(p *Tebot_assertContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TebotParserRULE_tebot_assert
}

func (*Tebot_assertContext) IsTebot_assertContext() {}

func NewTebot_assertContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Tebot_assertContext {
	var p = new(Tebot_assertContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TebotParserRULE_tebot_assert

	return p
}

func (s *Tebot_assertContext) GetParser() antlr.Parser { return s.parser }

func (s *Tebot_assertContext) GetDsn() ITebot_valueContext { return s.dsn }

func (s *Tebot_assertContext) GetSql() ITebot_valueContext { return s.sql }

func (s *Tebot_assertContext) GetExpected() ITebot_valueContext { return s.expected }

func (s *Tebot_assertContext) SetDsn(v ITebot_valueContext) { s.dsn = v }

func (s *Tebot_assertContext) SetSql(v ITebot_valueContext) { s.sql = v }

func (s *Tebot_assertContext) SetExpected(v ITebot_valueContext) { s.expected = v }

func (s *Tebot_assertContext) AllTebot_value() []ITebot_valueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITebot_valueContext); ok {
			len++
		}
	}

	tst := make([]ITebot_valueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITebot_valueContext); ok {
			tst[i] = t.(ITebot_valueContext)
			i++
		}
	}

	return tst
}

func (s *Tebot_assertContext) Tebot_value(i int) ITebot_valueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITebot_valueContext); ok {
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

	return t.(ITebot_valueContext)
}

func (s *Tebot_assertContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Tebot_assertContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Tebot_assertContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TebotListener); ok {
		listenerT.EnterTebot_assert(s)
	}
}

func (s *Tebot_assertContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TebotListener); ok {
		listenerT.ExitTebot_assert(s)
	}
}

func (p *TebotParser) Tebot_assert() (localctx ITebot_assertContext) {
	localctx = NewTebot_assertContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, TebotParserRULE_tebot_assert)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(38)
		p.Match(TebotParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(39)
		p.Match(TebotParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(40)

		var _x = p.Tebot_value()

		localctx.(*Tebot_assertContext).dsn = _x
	}
	{
		p.SetState(41)
		p.Match(TebotParserT__14)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(42)

		var _x = p.Tebot_value()

		localctx.(*Tebot_assertContext).sql = _x
	}
	{
		p.SetState(43)
		p.Match(TebotParserT__14)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(44)

		var _x = p.Tebot_value()

		localctx.(*Tebot_assertContext).expected = _x
	}
	{
		p.SetState(45)
		p.Match(TebotParserT__13)
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

// ITebot_operationsContext is an interface to support dynamic dispatch.
type ITebot_operationsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllTebot_operation() []ITebot_operationContext
	Tebot_operation(i int) ITebot_operationContext

	// IsTebot_operationsContext differentiates from other interfaces.
	IsTebot_operationsContext()
}

type Tebot_operationsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTebot_operationsContext() *Tebot_operationsContext {
	var p = new(Tebot_operationsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TebotParserRULE_tebot_operations
	return p
}

func InitEmptyTebot_operationsContext(p *Tebot_operationsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TebotParserRULE_tebot_operations
}

func (*Tebot_operationsContext) IsTebot_operationsContext() {}

func NewTebot_operationsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Tebot_operationsContext {
	var p = new(Tebot_operationsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TebotParserRULE_tebot_operations

	return p
}

func (s *Tebot_operationsContext) GetParser() antlr.Parser { return s.parser }

func (s *Tebot_operationsContext) AllTebot_operation() []ITebot_operationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITebot_operationContext); ok {
			len++
		}
	}

	tst := make([]ITebot_operationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITebot_operationContext); ok {
			tst[i] = t.(ITebot_operationContext)
			i++
		}
	}

	return tst
}

func (s *Tebot_operationsContext) Tebot_operation(i int) ITebot_operationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITebot_operationContext); ok {
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

	return t.(ITebot_operationContext)
}

func (s *Tebot_operationsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Tebot_operationsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Tebot_operationsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TebotListener); ok {
		listenerT.EnterTebot_operations(s)
	}
}

func (s *Tebot_operationsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TebotListener); ok {
		listenerT.ExitTebot_operations(s)
	}
}

func (p *TebotParser) Tebot_operations() (localctx ITebot_operationsContext) {
	localctx = NewTebot_operationsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, TebotParserRULE_tebot_operations)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(50)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4094) != 0 {
		{
			p.SetState(47)
			p.Tebot_operation()
		}

		p.SetState(52)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
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
