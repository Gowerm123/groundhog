package main

import (
	"awesomeProject/internal/eval"
	"awesomeProject/internal/lex"
	"awesomeProject/internal/parse"
	"awesomeProject/internal/printer"
	"fmt"
)

func main() {
	text := "3 + 4 - 5 + 6 * 2"
	lexer := lex.NewLexer(lex.NewSource(text))

	ast, err := parse.Parse(&lexer)
	if err != nil {
		panic(err)
	}

	println(printer.String(&ast))
	println(fmt.Sprintf("= %d", eval.EvalExpression(&ast)))
	println(fmt.Sprintf("= %d", 3+4-5+6*2))
}
