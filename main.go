package main

import (
	"github.com/abusizhishen/sql2Go/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func main() {
	calc()
}

func calc()  {
	input,err :=  antlr.NewFileStream("test.sql")
	if err != nil{
		panic(err)
	}
	lexer := parser.NewsqlLexer(input)
	tokens := antlr.NewCommonTokenStream(lexer,antlr.LexerDefaultTokenChannel)
	p := parser.NewsqlParser(tokens)

	antlr.ParseTreeWalkerDefault.Walk(SqlBaseListen{},p.Table())
}
func strStream(s string)  {
	calc()
}

type SqlBaseListen struct {
	*parser.BasesqlListener
}
