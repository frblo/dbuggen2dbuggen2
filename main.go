package main

import (
	"dbuggen2dbuggen2/install"
	"dbuggen2dbuggen2/lexer"
	"fmt"
)

func main() {
	path := "dbuggen"
	install.Installdbuggen(path)
	lexedIssues, lexedArticles := lexer.Lex(path)
	// parsedIssue := parser.Parse(lexedIssues, lexedArticles)

	fmt.Println(lexedIssues)
	fmt.Println(lexedArticles)
}
