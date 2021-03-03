package src

import (
	"fmt"
	"github.com/abusizhishen/sql2go/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type Listener struct {
	*parser.BaseSqlListener
}

func (l *Listener)EnterTable(ctx *parser.TableContext)  {
	fmt.Println(ctx.GetText())
	fmt.Println("type struct " + l.String(ctx.TableName()) + "{")
}

func (l *Listener)ExitTable(ctx *parser.TableContext)  {
	fmt.Println("}")
}

func (l *Listener)EnterFieldStatement(ctx *parser.FieldStatementContext)  {
	var Id = l.String(ctx.Field())
	fmt.Print(ctx.GetFileType().GetText())
	fmt.Println("%s    %s   json ")
}

func (l *Listener)ExitFieldStatement(ctx *parser.FieldStatementContext)  {
	fmt.Println("")
}

func (l *Listener)String(tree antlr.ParseTree) string {
	string2 := tree.GetText()
	if len(string2) <2{
		panic(string2)
	}

	return string2[1:len(string2)-1]
}