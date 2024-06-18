package main

import (
	"dbuggen2dbuggen2/executor"
	"dbuggen2dbuggen2/install"
	"dbuggen2dbuggen2/lexer"
	"dbuggen2dbuggen2/parser"
)

func main() {
	path := "dbuggen"
	install.Installdbuggen(path)
	lexedIssues, lexedArticles := lexer.Lex(path)
	parsedIssue := parser.Parse(lexedIssues, lexedArticles)
	executor.Execute(parsedIssue)
}
