// Code generated from Sql.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // Sql

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 18, 82, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	5, 2, 27, 10, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5,
	5, 5, 38, 10, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7,
	48, 10, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 5, 9, 57, 10, 9, 3,
	9, 5, 9, 60, 10, 9, 3, 9, 5, 9, 63, 10, 9, 3, 9, 5, 9, 66, 10, 9, 3, 10,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 7, 10, 75, 10, 10, 12, 10, 14,
	10, 78, 11, 10, 3, 10, 3, 10, 3, 10, 2, 2, 11, 2, 4, 6, 8, 10, 12, 14,
	16, 18, 2, 2, 2, 80, 2, 26, 3, 2, 2, 2, 4, 28, 3, 2, 2, 2, 6, 32, 3, 2,
	2, 2, 8, 35, 3, 2, 2, 2, 10, 39, 3, 2, 2, 2, 12, 47, 3, 2, 2, 2, 14, 49,
	3, 2, 2, 2, 16, 53, 3, 2, 2, 2, 18, 67, 3, 2, 2, 2, 20, 21, 7, 3, 2, 2,
	21, 22, 7, 18, 2, 2, 22, 27, 7, 3, 2, 2, 23, 24, 7, 4, 2, 2, 24, 25, 7,
	18, 2, 2, 25, 27, 7, 4, 2, 2, 26, 20, 3, 2, 2, 2, 26, 23, 3, 2, 2, 2, 27,
	3, 3, 2, 2, 2, 28, 29, 7, 5, 2, 2, 29, 30, 7, 13, 2, 2, 30, 31, 7, 5, 2,
	2, 31, 5, 3, 2, 2, 2, 32, 33, 7, 6, 2, 2, 33, 34, 5, 2, 2, 2, 34, 7, 3,
	2, 2, 2, 35, 37, 7, 7, 2, 2, 36, 38, 7, 18, 2, 2, 37, 36, 3, 2, 2, 2, 37,
	38, 3, 2, 2, 2, 38, 9, 3, 2, 2, 2, 39, 40, 7, 8, 2, 2, 40, 41, 7, 14, 2,
	2, 41, 42, 7, 9, 2, 2, 42, 11, 3, 2, 2, 2, 43, 48, 7, 13, 2, 2, 44, 45,
	7, 5, 2, 2, 45, 46, 7, 13, 2, 2, 46, 48, 7, 5, 2, 2, 47, 43, 3, 2, 2, 2,
	47, 44, 3, 2, 2, 2, 48, 13, 3, 2, 2, 2, 49, 50, 7, 5, 2, 2, 50, 51, 7,
	13, 2, 2, 51, 52, 7, 5, 2, 2, 52, 15, 3, 2, 2, 2, 53, 54, 5, 12, 7, 2,
	54, 56, 7, 13, 2, 2, 55, 57, 5, 10, 6, 2, 56, 55, 3, 2, 2, 2, 56, 57, 3,
	2, 2, 2, 57, 59, 3, 2, 2, 2, 58, 60, 7, 15, 2, 2, 59, 58, 3, 2, 2, 2, 59,
	60, 3, 2, 2, 2, 60, 62, 3, 2, 2, 2, 61, 63, 5, 6, 4, 2, 62, 61, 3, 2, 2,
	2, 62, 63, 3, 2, 2, 2, 63, 65, 3, 2, 2, 2, 64, 66, 5, 8, 5, 2, 65, 64,
	3, 2, 2, 2, 65, 66, 3, 2, 2, 2, 66, 17, 3, 2, 2, 2, 67, 68, 7, 10, 2, 2,
	68, 69, 7, 11, 2, 2, 69, 70, 5, 4, 3, 2, 70, 71, 7, 8, 2, 2, 71, 76, 5,
	16, 9, 2, 72, 73, 7, 12, 2, 2, 73, 75, 5, 16, 9, 2, 74, 72, 3, 2, 2, 2,
	75, 78, 3, 2, 2, 2, 76, 74, 3, 2, 2, 2, 76, 77, 3, 2, 2, 2, 77, 79, 3,
	2, 2, 2, 78, 76, 3, 2, 2, 2, 79, 80, 7, 9, 2, 2, 80, 19, 3, 2, 2, 2, 10,
	26, 37, 47, 56, 59, 62, 65, 76,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'\"'", "'''", "'`'", "'DEFAULT'", "'COMMENT'", "'('", "')'", "'CREATE'",
	"'TABLE'", "','",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "Id", "Int", "NotNull", "WS",
	"FieldType", "STRING",
}

var ruleNames = []string{
	"statement", "tableName", "defaultStatement", "comment", "length", "field",
	"fieldName", "fieldStatement", "table",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type SqlParser struct {
	*antlr.BaseParser
}

func NewSqlParser(input antlr.TokenStream) *SqlParser {
	this := new(SqlParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Sql.g4"

	return this
}

// SqlParser tokens.
const (
	SqlParserEOF       = antlr.TokenEOF
	SqlParserT__0      = 1
	SqlParserT__1      = 2
	SqlParserT__2      = 3
	SqlParserT__3      = 4
	SqlParserT__4      = 5
	SqlParserT__5      = 6
	SqlParserT__6      = 7
	SqlParserT__7      = 8
	SqlParserT__8      = 9
	SqlParserT__9      = 10
	SqlParserId        = 11
	SqlParserInt       = 12
	SqlParserNotNull   = 13
	SqlParserWS        = 14
	SqlParserFieldType = 15
	SqlParserSTRING    = 16
)

// SqlParser rules.
const (
	SqlParserRULE_statement        = 0
	SqlParserRULE_tableName        = 1
	SqlParserRULE_defaultStatement = 2
	SqlParserRULE_comment          = 3
	SqlParserRULE_length           = 4
	SqlParserRULE_field            = 5
	SqlParserRULE_fieldName        = 6
	SqlParserRULE_fieldStatement   = 7
	SqlParserRULE_table            = 8
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
	p.RuleIndex = SqlParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) STRING() antlr.TerminalNode {
	return s.GetToken(SqlParserSTRING, 0)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *SqlParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SqlParserRULE_statement)

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

	p.SetState(24)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SqlParserT__0:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(18)
			p.Match(SqlParserT__0)
		}
		{
			p.SetState(19)
			p.Match(SqlParserSTRING)
		}
		{
			p.SetState(20)
			p.Match(SqlParserT__0)
		}

	case SqlParserT__1:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(21)
			p.Match(SqlParserT__1)
		}
		{
			p.SetState(22)
			p.Match(SqlParserSTRING)
		}
		{
			p.SetState(23)
			p.Match(SqlParserT__1)
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
	p.RuleIndex = SqlParserRULE_tableName
	return p
}

func (*TableNameContext) IsTableNameContext() {}

func NewTableNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TableNameContext {
	var p = new(TableNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_tableName

	return p
}

func (s *TableNameContext) GetParser() antlr.Parser { return s.parser }

func (s *TableNameContext) Id() antlr.TerminalNode {
	return s.GetToken(SqlParserId, 0)
}

func (s *TableNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TableNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TableNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterTableName(s)
	}
}

func (s *TableNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitTableName(s)
	}
}

func (p *SqlParser) TableName() (localctx ITableNameContext) {
	localctx = NewTableNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SqlParserRULE_tableName)

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
		p.SetState(26)
		p.Match(SqlParserT__2)
	}
	{
		p.SetState(27)
		p.Match(SqlParserId)
	}
	{
		p.SetState(28)
		p.Match(SqlParserT__2)
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
	p.RuleIndex = SqlParserRULE_defaultStatement
	return p
}

func (*DefaultStatementContext) IsDefaultStatementContext() {}

func NewDefaultStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefaultStatementContext {
	var p = new(DefaultStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_defaultStatement

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
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterDefaultStatement(s)
	}
}

func (s *DefaultStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitDefaultStatement(s)
	}
}

func (p *SqlParser) DefaultStatement() (localctx IDefaultStatementContext) {
	localctx = NewDefaultStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, SqlParserRULE_defaultStatement)

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
		p.SetState(30)
		p.Match(SqlParserT__3)
	}
	{
		p.SetState(31)
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
	p.RuleIndex = SqlParserRULE_comment
	return p
}

func (*CommentContext) IsCommentContext() {}

func NewCommentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommentContext {
	var p = new(CommentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_comment

	return p
}

func (s *CommentContext) GetParser() antlr.Parser { return s.parser }

func (s *CommentContext) STRING() antlr.TerminalNode {
	return s.GetToken(SqlParserSTRING, 0)
}

func (s *CommentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterComment(s)
	}
}

func (s *CommentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitComment(s)
	}
}

func (p *SqlParser) Comment() (localctx ICommentContext) {
	localctx = NewCommentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, SqlParserRULE_comment)
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
		p.SetState(33)
		p.Match(SqlParserT__4)
	}
	p.SetState(35)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SqlParserSTRING {
		{
			p.SetState(34)
			p.Match(SqlParserSTRING)
		}

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
	p.RuleIndex = SqlParserRULE_length
	return p
}

func (*LengthContext) IsLengthContext() {}

func NewLengthContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LengthContext {
	var p = new(LengthContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_length

	return p
}

func (s *LengthContext) GetParser() antlr.Parser { return s.parser }

func (s *LengthContext) Int() antlr.TerminalNode {
	return s.GetToken(SqlParserInt, 0)
}

func (s *LengthContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LengthContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LengthContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterLength(s)
	}
}

func (s *LengthContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitLength(s)
	}
}

func (p *SqlParser) Length() (localctx ILengthContext) {
	localctx = NewLengthContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, SqlParserRULE_length)

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
		p.SetState(37)
		p.Match(SqlParserT__5)
	}
	{
		p.SetState(38)
		p.Match(SqlParserInt)
	}
	{
		p.SetState(39)
		p.Match(SqlParserT__6)
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
	p.RuleIndex = SqlParserRULE_field
	return p
}

func (*FieldContext) IsFieldContext() {}

func NewFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldContext {
	var p = new(FieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_field

	return p
}

func (s *FieldContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldContext) Id() antlr.TerminalNode {
	return s.GetToken(SqlParserId, 0)
}

func (s *FieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterField(s)
	}
}

func (s *FieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitField(s)
	}
}

func (p *SqlParser) Field() (localctx IFieldContext) {
	localctx = NewFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, SqlParserRULE_field)

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

	p.SetState(45)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SqlParserId:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(41)
			p.Match(SqlParserId)
		}

	case SqlParserT__2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(42)
			p.Match(SqlParserT__2)
		}
		{
			p.SetState(43)
			p.Match(SqlParserId)
		}
		{
			p.SetState(44)
			p.Match(SqlParserT__2)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFieldNameContext is an interface to support dynamic dispatch.
type IFieldNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldNameContext differentiates from other interfaces.
	IsFieldNameContext()
}

type FieldNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldNameContext() *FieldNameContext {
	var p = new(FieldNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_fieldName
	return p
}

func (*FieldNameContext) IsFieldNameContext() {}

func NewFieldNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldNameContext {
	var p = new(FieldNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_fieldName

	return p
}

func (s *FieldNameContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldNameContext) Id() antlr.TerminalNode {
	return s.GetToken(SqlParserId, 0)
}

func (s *FieldNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterFieldName(s)
	}
}

func (s *FieldNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitFieldName(s)
	}
}

func (p *SqlParser) FieldName() (localctx IFieldNameContext) {
	localctx = NewFieldNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, SqlParserRULE_fieldName)

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
		p.SetState(47)
		p.Match(SqlParserT__2)
	}
	{
		p.SetState(48)
		p.Match(SqlParserId)
	}
	{
		p.SetState(49)
		p.Match(SqlParserT__2)
	}

	return localctx
}

// IFieldStatementContext is an interface to support dynamic dispatch.
type IFieldStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetFileType returns the fileType token.
	GetFileType() antlr.Token

	// SetFileType sets the fileType token.
	SetFileType(antlr.Token)

	// IsFieldStatementContext differentiates from other interfaces.
	IsFieldStatementContext()
}

type FieldStatementContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	fileType antlr.Token
}

func NewEmptyFieldStatementContext() *FieldStatementContext {
	var p = new(FieldStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_fieldStatement
	return p
}

func (*FieldStatementContext) IsFieldStatementContext() {}

func NewFieldStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldStatementContext {
	var p = new(FieldStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_fieldStatement

	return p
}

func (s *FieldStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldStatementContext) GetFileType() antlr.Token { return s.fileType }

func (s *FieldStatementContext) SetFileType(v antlr.Token) { s.fileType = v }

func (s *FieldStatementContext) Field() IFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *FieldStatementContext) Id() antlr.TerminalNode {
	return s.GetToken(SqlParserId, 0)
}

func (s *FieldStatementContext) Length() ILengthContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILengthContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILengthContext)
}

func (s *FieldStatementContext) NotNull() antlr.TerminalNode {
	return s.GetToken(SqlParserNotNull, 0)
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
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterFieldStatement(s)
	}
}

func (s *FieldStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitFieldStatement(s)
	}
}

func (p *SqlParser) FieldStatement() (localctx IFieldStatementContext) {
	localctx = NewFieldStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, SqlParserRULE_fieldStatement)
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
		p.SetState(51)
		p.Field()
	}
	{
		p.SetState(52)

		var _m = p.Match(SqlParserId)

		localctx.(*FieldStatementContext).fileType = _m
	}
	p.SetState(54)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SqlParserT__5 {
		{
			p.SetState(53)
			p.Length()
		}

	}
	p.SetState(57)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SqlParserNotNull {
		{
			p.SetState(56)
			p.Match(SqlParserNotNull)
		}

	}
	p.SetState(60)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SqlParserT__3 {
		{
			p.SetState(59)
			p.DefaultStatement()
		}

	}
	p.SetState(63)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SqlParserT__4 {
		{
			p.SetState(62)
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
	p.RuleIndex = SqlParserRULE_table
	return p
}

func (*TableContext) IsTableContext() {}

func NewTableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TableContext {
	var p = new(TableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_table

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
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterTable(s)
	}
}

func (s *TableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitTable(s)
	}
}

func (p *SqlParser) Table() (localctx ITableContext) {
	localctx = NewTableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, SqlParserRULE_table)
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
		p.Match(SqlParserT__7)
	}
	{
		p.SetState(66)
		p.Match(SqlParserT__8)
	}
	{
		p.SetState(67)
		p.TableName()
	}
	{
		p.SetState(68)
		p.Match(SqlParserT__5)
	}
	{
		p.SetState(69)
		p.FieldStatement()
	}
	p.SetState(74)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SqlParserT__9 {
		{
			p.SetState(70)
			p.Match(SqlParserT__9)
		}
		{
			p.SetState(71)
			p.FieldStatement()
		}

		p.SetState(76)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(77)
		p.Match(SqlParserT__6)
	}

	return localctx
}
