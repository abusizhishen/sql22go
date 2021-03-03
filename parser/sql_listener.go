// Code generated from Sql.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // Sql

import "github.com/antlr/antlr4/runtime/Go/antlr"

// SqlListener is a complete listener for a parse tree produced by SqlParser.
type SqlListener interface {
	antlr.ParseTreeListener

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterTableName is called when entering the tableName production.
	EnterTableName(c *TableNameContext)

	// EnterDefaultStatement is called when entering the defaultStatement production.
	EnterDefaultStatement(c *DefaultStatementContext)

	// EnterComment is called when entering the comment production.
	EnterComment(c *CommentContext)

	// EnterLength is called when entering the length production.
	EnterLength(c *LengthContext)

	// EnterField is called when entering the field production.
	EnterField(c *FieldContext)

	// EnterFieldName is called when entering the fieldName production.
	EnterFieldName(c *FieldNameContext)

	// EnterFieldStatement is called when entering the fieldStatement production.
	EnterFieldStatement(c *FieldStatementContext)

	// EnterTable is called when entering the table production.
	EnterTable(c *TableContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitTableName is called when exiting the tableName production.
	ExitTableName(c *TableNameContext)

	// ExitDefaultStatement is called when exiting the defaultStatement production.
	ExitDefaultStatement(c *DefaultStatementContext)

	// ExitComment is called when exiting the comment production.
	ExitComment(c *CommentContext)

	// ExitLength is called when exiting the length production.
	ExitLength(c *LengthContext)

	// ExitField is called when exiting the field production.
	ExitField(c *FieldContext)

	// ExitFieldName is called when exiting the fieldName production.
	ExitFieldName(c *FieldNameContext)

	// ExitFieldStatement is called when exiting the fieldStatement production.
	ExitFieldStatement(c *FieldStatementContext)

	// ExitTable is called when exiting the table production.
	ExitTable(c *TableContext)
}
