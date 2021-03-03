// Code generated from Sql.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // Sql

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseSqlListener is a complete listener for a parse tree produced by SqlParser.
type BaseSqlListener struct{}

var _ SqlListener = &BaseSqlListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSqlListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSqlListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSqlListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSqlListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseSqlListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseSqlListener) ExitStatement(ctx *StatementContext) {}

// EnterTableName is called when production tableName is entered.
func (s *BaseSqlListener) EnterTableName(ctx *TableNameContext) {}

// ExitTableName is called when production tableName is exited.
func (s *BaseSqlListener) ExitTableName(ctx *TableNameContext) {}

// EnterDefaultStatement is called when production defaultStatement is entered.
func (s *BaseSqlListener) EnterDefaultStatement(ctx *DefaultStatementContext) {}

// ExitDefaultStatement is called when production defaultStatement is exited.
func (s *BaseSqlListener) ExitDefaultStatement(ctx *DefaultStatementContext) {}

// EnterComment is called when production comment is entered.
func (s *BaseSqlListener) EnterComment(ctx *CommentContext) {}

// ExitComment is called when production comment is exited.
func (s *BaseSqlListener) ExitComment(ctx *CommentContext) {}

// EnterLength is called when production length is entered.
func (s *BaseSqlListener) EnterLength(ctx *LengthContext) {}

// ExitLength is called when production length is exited.
func (s *BaseSqlListener) ExitLength(ctx *LengthContext) {}

// EnterField is called when production field is entered.
func (s *BaseSqlListener) EnterField(ctx *FieldContext) {}

// ExitField is called when production field is exited.
func (s *BaseSqlListener) ExitField(ctx *FieldContext) {}

// EnterFieldName is called when production fieldName is entered.
func (s *BaseSqlListener) EnterFieldName(ctx *FieldNameContext) {}

// ExitFieldName is called when production fieldName is exited.
func (s *BaseSqlListener) ExitFieldName(ctx *FieldNameContext) {}

// EnterFieldStatement is called when production fieldStatement is entered.
func (s *BaseSqlListener) EnterFieldStatement(ctx *FieldStatementContext) {}

// ExitFieldStatement is called when production fieldStatement is exited.
func (s *BaseSqlListener) ExitFieldStatement(ctx *FieldStatementContext) {}

// EnterTable is called when production table is entered.
func (s *BaseSqlListener) EnterTable(ctx *TableContext) {}

// ExitTable is called when production table is exited.
func (s *BaseSqlListener) ExitTable(ctx *TableContext) {}
