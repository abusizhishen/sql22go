package main

import (
	"github.com/abusizhishen/sql2go/parser"
	"github.com/abusizhishen/sql2go/src"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func main() {
	input,err :=  antlr.NewFileStream("test.sql")
	if err != nil{
		panic(err)
	}
	lexer := parser.NewSqlLexer(input)
	tokens := antlr.NewCommonTokenStream(lexer,antlr.LexerDefaultTokenChannel)
	p := parser.NewSqlParser(tokens)

	antlr.ParseTreeWalkerDefault.Walk(&src.Listener{}, p.Table())
}
