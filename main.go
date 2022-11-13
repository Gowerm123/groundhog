package main

import (
	"awesomeProject/internal/eval"
	"awesomeProject/internal/lex"
	"awesomeProject/internal/parse"
	"fmt"
	"os"
)

func main() {
	text := os.Args[1]
	lexer := lex.NewLexer(lex.NewSource(text))

	ast, err := parse.Parse(&lexer)
	if err != nil {
		panic(err)
	}

	println(fmt.Sprintf("= %f", eval.EvalExpression(&ast)))
}
