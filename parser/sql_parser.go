// Code generated from sql.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // sql

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 18, 75, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 4, 9, 9, 9, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 5, 2, 25, 10,
	2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 45, 10, 7, 3, 8, 3, 8, 3,
	8, 5, 8, 50, 10, 8, 3, 8, 5, 8, 53, 10, 8, 3, 8, 5, 8, 56, 10, 8, 3, 8,
	5, 8, 59, 10, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 7, 9, 68, 10,
	9, 12, 9, 14, 9, 71, 11, 9, 3, 9, 3, 9, 3, 9, 2, 2, 10, 2, 4, 6, 8, 10,
	12, 14, 16, 2, 2, 2, 73, 2, 24, 3, 2, 2, 2, 4, 26, 3, 2, 2, 2, 6, 30, 3,
	2, 2, 2, 8, 33, 3, 2, 2, 2, 10, 36, 3, 2, 2, 2, 12, 44, 3, 2, 2, 2, 14,
	46, 3, 2, 2, 2, 16, 60, 3, 2, 2, 2, 18, 19, 7, 3, 2, 2, 19, 20, 7, 18,
	2, 2, 20, 25, 7, 3, 2, 2, 21, 22, 7, 4, 2, 2, 22, 23, 7, 18, 2, 2, 23,
	25, 7, 4, 2, 2, 24, 18, 3, 2, 2, 2, 24, 21, 3, 2, 2, 2, 25, 3, 3, 2, 2,
	2, 26, 27, 7, 5, 2, 2, 27, 28, 7, 13, 2, 2, 28, 29, 7, 5, 2, 2, 29, 5,
	3, 2, 2, 2, 30, 31, 7, 6, 2, 2, 31, 32, 5, 2, 2, 2, 32, 7, 3, 2, 2, 2,
	33, 34, 7, 7, 2, 2, 34, 35, 5, 2, 2, 2, 35, 9, 3, 2, 2, 2, 36, 37, 7, 8,
	2, 2, 37, 38, 7, 14, 2, 2, 38, 39, 7, 9, 2, 2, 39, 11, 3, 2, 2, 2, 40,
	45, 7, 13, 2, 2, 41, 42, 7, 5, 2, 2, 42, 43, 7, 13, 2, 2, 43, 45, 7, 5,
	2, 2, 44, 40, 3, 2, 2, 2, 44, 41, 3, 2, 2, 2, 45, 13, 3, 2, 2, 2, 46, 47,
	5, 12, 7, 2, 47, 49, 7, 17, 2, 2, 48, 50, 5, 10, 6, 2, 49, 48, 3, 2, 2,
	2, 49, 50, 3, 2, 2, 2, 50, 52, 3, 2, 2, 2, 51, 53, 7, 15, 2, 2, 52, 51,
	3, 2, 2, 2, 52, 53, 3, 2, 2, 2, 53, 55, 3, 2, 2, 2, 54, 56, 5, 6, 4, 2,
	55, 54, 3, 2, 2, 2, 55, 56, 3, 2, 2, 2, 56, 58, 3, 2, 2, 2, 57, 59, 5,
	8, 5, 2, 58, 57, 3, 2, 2, 2, 58, 59, 3, 2, 2, 2, 59, 15, 3, 2, 2, 2, 60,
	61, 7, 10, 2, 2, 61, 62, 7, 11, 2, 2, 62, 63, 5, 4, 3, 2, 63, 64, 7, 8,
	2, 2, 64, 69, 5, 14, 8, 2, 65, 66, 7, 12, 2, 2, 66, 68, 5, 14, 8, 2, 67,
	65, 3, 2, 2, 2, 68, 71, 3, 2, 2, 2, 69, 67, 3, 2, 2, 2, 69, 70, 3, 2, 2,
	2, 70, 72, 3, 2, 2, 2, 71, 69, 3, 2, 2, 2, 72, 73, 7, 9, 2, 2, 73, 17,
	3, 2, 2, 2, 9, 24, 44, 49, 52, 55, 58, 69,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'\"'", "'''", "'`'", "'DEFAULT'", "'COMMENT'", "'('", "')'", "'CREATE'",
	"'TABLE'", "','",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "Id", "Int", "NotNull", "WS",
	"Type", "STRING",
}

var ruleNames = []string{
	"statement", "tableName", "defaultStatement", "comment", "length", "field",
	"fieldStatement", "table",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type sqlParser struct {
	*antlr.BaseParser
}

func NewsqlParser(input antlr.TokenStream) *sqlParser {
	this := new(sqlParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "sql.g4"

	return this
}

// sqlParser tokens.
const (
	sqlParserEOF     = antlr.TokenEOF
	sqlParserT__0    = 1
	sqlParserT__1    = 2
	sqlParserT__2    = 3
	sqlParserT__3    = 4
	sqlParserT__4    = 5
	sqlParserT__5    = 6
	sqlParserT__6    = 7
	sqlParserT__7    = 8
	sqlParserT__8    = 9
	sqlParserT__9    = 10
	sqlParserId      = 11
	sqlParserInt     = 12
	sqlParserNotNull = 13
	sqlParserWS      = 14
	sqlParserType    = 15
	sqlParserSTRING  = 16
)

// sqlParser rules.
const (
	sqlParserRULE_statement        = 0
	sqlParserRULE_tableName        = 1
	sqlParserRULE_defaultStatement = 2
	sqlParserRULE_comment          = 3
	sqlParserRULE_length           = 4
	sqlParserRULE_field            = 5
	sqlParserRULE_fieldStatement   = 6
	sqlParserRULE_table            = 7
)

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = sqlParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = sqlParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) STRING() antlr.TerminalNode {
	return s.GetToken(sqlParserSTRING, 0)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sqlListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sqlListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *sqlParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, sqlParserRULE_statement)

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

	p.SetState(22)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case sqlParserT__0:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(16)
			p.Match(sqlParserT__0)
		}
		{
			p.SetState(17)
			p.Match(sqlParserSTRING)
		}
		{
			p.SetState(18)
			p.Match(sqlParserT__0)
		}

	case sqlParserT__1:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(19)
			p.Match(sqlParserT__1)
		}
		{
			p.SetState(20)
			p.Match(sqlParserSTRING)
		}
		{
			p.SetState(21)
			p.Match(sqlParserT__1)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ITableNameContext is an interface to support dynamic dispatch.
type ITableNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTableNameContext differentiates from other interfaces.
	IsTableNameContext()
}

type TableNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTableNameContext() *TableNameContext {
	var p = new(TableNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = sqlParserRULE_tableName
	return p
}

func (*TableNameContext) IsTableNameContext() {}

func NewTableNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TableNameContext {
	var p = new(TableNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = sqlParserRULE_tableName

	return p
}

func (s *TableNameContext) GetParser() antlr.Parser { return s.parser }

func (s *TableNameContext) Id() antlr.TerminalNode {
	return s.GetToken(sqlParserId, 0)
}

func (s *TableNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TableNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TableNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sqlListener); ok {
		listenerT.EnterTableName(s)
	}
}

func (s *TableNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sqlListener); ok {
		listenerT.ExitTableName(s)
	}
}

func (p *sqlParser) TableName() (localctx ITableNameContext) {
	localctx = NewTableNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, sqlParserRULE_tableName)

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
		p.Match(sqlParserT__2)
	}
	{
		p.SetState(25)
		p.Match(sqlParserId)
	}
	{
		p.SetState(26)
		p.Match(sqlParserT__2)
	}

	return localctx
}

// IDefaultStatementContext is an interface to support dynamic dispatch.
type IDefaultStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDefaultStatementContext differentiates from other interfaces.
	IsDefaultStatementContext()
}

type DefaultStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefaultStatementContext() *DefaultStatementContext {
	var p = new(DefaultStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = sqlParserRULE_defaultStatement
	return p
}

func (*DefaultStatementContext) IsDefaultStatementContext() {}

func NewDefaultStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefaultStatementContext {
	var p = new(DefaultStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = sqlParserRULE_defaultStatement

	return p
}

func (s *DefaultStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *DefaultStatementContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *DefaultStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefaultStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefaultStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sqlListener); ok {
		listenerT.EnterDefaultStatement(s)
	}
}

func (s *DefaultStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sqlListener); ok {
		listenerT.ExitDefaultStatement(s)
	}
}

func (p *sqlParser) DefaultStatement() (localctx IDefaultStatementContext) {
	localctx = NewDefaultStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, sqlParserRULE_defaultStatement)

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
		p.SetState(28)
		p.Match(sqlParserT__3)
	}
	{
		p.SetState(29)
		p.Statement()
	}

	return localctx
}

// ICommentContext is an interface to support dynamic dispatch.
type ICommentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCommentContext differentiates from other interfaces.
	IsCommentContext()
}

type CommentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommentContext() *CommentContext {
	var p = new(CommentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = sqlParserRULE_comment
	return p
}

func (*CommentContext) IsCommentContext() {}

func NewCommentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommentContext {
	var p = new(CommentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = sqlParserRULE_comment

	return p
}

func (s *CommentContext) GetParser() antlr.Parser { return s.parser }

func (s *CommentContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *CommentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sqlListener); ok {
		listenerT.EnterComment(s)
	}
}

func (s *CommentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sqlListener); ok {
		listenerT.ExitComment(s)
	}
}

func (p *sqlParser) Comment() (localctx ICommentContext) {
	localctx = NewCommentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, sqlParserRULE_comment)

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
		p.SetState(31)
		p.Match(sqlParserT__4)
	}
	{
		p.SetState(32)
		p.Statement()
	}

	return localctx
}

// ILengthContext is an interface to support dynamic dispatch.
type ILengthContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLengthContext differentiates from other interfaces.
	IsLengthContext()
}

type LengthContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLengthContext() *LengthContext {
	var p = new(LengthContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = sqlParserRULE_length
	return p
}

func (*LengthContext) IsLengthContext() {}

func NewLengthContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LengthContext {
	var p = new(LengthContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = sqlParserRULE_length

	return p
}

func (s *LengthContext) GetParser() antlr.Parser { return s.parser }

func (s *LengthContext) Int() antlr.TerminalNode {
	return s.GetToken(sqlParserInt, 0)
}

func (s *LengthContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LengthContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LengthContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sqlListener); ok {
		listenerT.EnterLength(s)
	}
}

func (s *LengthContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sqlListener); ok {
		listenerT.ExitLength(s)
	}
}

func (p *sqlParser) Length() (localctx ILengthContext) {
	localctx = NewLengthContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, sqlParserRULE_length)

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
		p.SetState(34)
		p.Match(sqlParserT__5)
	}
	{
		p.SetState(35)
		p.Match(sqlParserInt)
	}
	{
		p.SetState(36)
		p.Match(sqlParserT__6)
	}

	return localctx
}

// IFieldContext is an interface to support dynamic dispatch.
type IFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldContext differentiates from other interfaces.
	IsFieldContext()
}

type FieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldContext() *FieldContext {
	var p = new(FieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = sqlParserRULE_field
	return p
}

func (*FieldContext) IsFieldContext() {}

func NewFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldContext {
	var p = new(FieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = sqlParserRULE_field

	return p
}

func (s *FieldContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldContext) Id() antlr.TerminalNode {
	return s.GetToken(sqlParserId, 0)
}

func (s *FieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sqlListener); ok {
		listenerT.EnterField(s)
	}
}

func (s *FieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sqlListener); ok {
		listenerT.ExitField(s)
	}
}

func (p *sqlParser) Field() (localctx IFieldContext) {
	localctx = NewFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, sqlParserRULE_field)

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

	p.SetState(42)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case sqlParserId:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(38)
			p.Match(sqlParserId)
		}

	case sqlParserT__2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(39)
			p.Match(sqlParserT__2)
		}
		{
			p.SetState(40)
			p.Match(sqlParserId)
		}
		{
			p.SetState(41)
			p.Match(sqlParserT__2)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFieldStatementContext is an interface to support dynamic dispatch.
type IFieldStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldStatementContext differentiates from other interfaces.
	IsFieldStatementContext()
}

type FieldStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldStatementContext() *FieldStatementContext {
	var p = new(FieldStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = sqlParserRULE_fieldStatement
	return p
}

func (*FieldStatementContext) IsFieldStatementContext() {}

func NewFieldStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldStatementContext {
	var p = new(FieldStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = sqlParserRULE_fieldStatement

	return p
}

func (s *FieldStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldStatementContext) Field() IFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *FieldStatementContext) Type() antlr.TerminalNode {
	return s.GetToken(sqlParserType, 0)
}

func (s *FieldStatementContext) Length() ILengthContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILengthContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILengthContext)
}

func (s *FieldStatementContext) NotNull() antlr.TerminalNode {
	return s.GetToken(sqlParserNotNull, 0)
}

func (s *FieldStatementContext) DefaultStatement() IDefaultStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefaultStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDefaultStatementContext)
}

func (s *FieldStatementContext) Comment() ICommentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommentContext)
}

func (s *FieldStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sqlListener); ok {
		listenerT.EnterFieldStatement(s)
	}
}

func (s *FieldStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sqlListener); ok {
		listenerT.ExitFieldStatement(s)
	}
}

func (p *sqlParser) FieldStatement() (localctx IFieldStatementContext) {
	localctx = NewFieldStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, sqlParserRULE_fieldStatement)
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
		p.SetState(44)
		p.Field()
	}
	{
		p.SetState(45)
		p.Match(sqlParserType)
	}
	p.SetState(47)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == sqlParserT__5 {
		{
			p.SetState(46)
			p.Length()
		}

	}
	p.SetState(50)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == sqlParserNotNull {
		{
			p.SetState(49)
			p.Match(sqlParserNotNull)
		}

	}
	p.SetState(53)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == sqlParserT__3 {
		{
			p.SetState(52)
			p.DefaultStatement()
		}

	}
	p.SetState(56)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == sqlParserT__4 {
		{
			p.SetState(55)
			p.Comment()
		}

	}

	return localctx
}

// ITableContext is an interface to support dynamic dispatch.
type ITableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTableContext differentiates from other interfaces.
	IsTableContext()
}

type TableContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTableContext() *TableContext {
	var p = new(TableContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = sqlParserRULE_table
	return p
}

func (*TableContext) IsTableContext() {}

func NewTableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TableContext {
	var p = new(TableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = sqlParserRULE_table

	return p
}

func (s *TableContext) GetParser() antlr.Parser { return s.parser }

func (s *TableContext) TableName() ITableNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITableNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITableNameContext)
}

func (s *TableContext) AllFieldStatement() []IFieldStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFieldStatementContext)(nil)).Elem())
	var tst = make([]IFieldStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFieldStatementContext)
		}
	}

	return tst
}

func (s *TableContext) FieldStatement(i int) IFieldStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFieldStatementContext)
}

func (s *TableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sqlListener); ok {
		listenerT.EnterTable(s)
	}
}

func (s *TableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sqlListener); ok {
		listenerT.ExitTable(s)
	}
}

func (p *sqlParser) Table() (localctx ITableContext) {
	localctx = NewTableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, sqlParserRULE_table)
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
		p.SetState(58)
		p.Match(sqlParserT__7)
	}
	{
		p.SetState(59)
		p.Match(sqlParserT__8)
	}
	{
		p.SetState(60)
		p.TableName()
	}
	{
		p.SetState(61)
		p.Match(sqlParserT__5)
	}
	{
		p.SetState(62)
		p.FieldStatement()
	}
	p.SetState(67)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == sqlParserT__9 {
		{
			p.SetState(63)
			p.Match(sqlParserT__9)
		}
		{
			p.SetState(64)
			p.FieldStatement()
		}

		p.SetState(69)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(70)
		p.Match(sqlParserT__6)
	}

	return localctx
}
