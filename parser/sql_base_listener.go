// Code generated from sql.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // sql

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasesqlListener is a complete listener for a parse tree produced by sqlParser.
type BasesqlListener struct{}

var _ sqlListener = &BasesqlListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasesqlListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasesqlListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasesqlListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasesqlListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStatement is called when production statement is entered.
func (s *BasesqlListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BasesqlListener) ExitStatement(ctx *StatementContext) {}

// EnterTableName is called when production tableName is entered.
func (s *BasesqlListener) EnterTableName(ctx *TableNameContext) {}

// ExitTableName is called when production tableName is exited.
func (s *BasesqlListener) ExitTableName(ctx *TableNameContext) {}

// EnterDefaultStatement is called when production defaultStatement is entered.
func (s *BasesqlListener) EnterDefaultStatement(ctx *DefaultStatementContext) {}

// ExitDefaultStatement is called when production defaultStatement is exited.
func (s *BasesqlListener) ExitDefaultStatement(ctx *DefaultStatementContext) {}

// EnterComment is called when production comment is entered.
func (s *BasesqlListener) EnterComment(ctx *CommentContext) {}

// ExitComment is called when production comment is exited.
func (s *BasesqlListener) ExitComment(ctx *CommentContext) {}

// EnterLength is called when production length is entered.
func (s *BasesqlListener) EnterLength(ctx *LengthContext) {}

// ExitLength is called when production length is exited.
func (s *BasesqlListener) ExitLength(ctx *LengthContext) {}

// EnterField is called when production field is entered.
func (s *BasesqlListener) EnterField(ctx *FieldContext) {}

// ExitField is called when production field is exited.
func (s *BasesqlListener) ExitField(ctx *FieldContext) {}

// EnterFieldStatement is called when production fieldStatement is entered.
func (s *BasesqlListener) EnterFieldStatement(ctx *FieldStatementContext) {}

// ExitFieldStatement is called when production fieldStatement is exited.
func (s *BasesqlListener) ExitFieldStatement(ctx *FieldStatementContext) {}

// EnterTable is called when production table is entered.
func (s *BasesqlListener) EnterTable(ctx *TableContext) {}

// ExitTable is called when production table is exited.
func (s *BasesqlListener) ExitTable(ctx *TableContext) {}
