package main

import (
	"dbuggen2dbuggen2/install"
	"dbuggen2dbuggen2/parser"
	"fmt"
)

func main() {
	install.Installdbuggen()
	// parser.GetIssues()

	artics := parser.GetArticles()
	// for _, a := range artics {
	// 	fmt.Println(a.Title)
	// }
	fmt.Println(artics[0].Title)
	fmt.Println(artics[0].Author)
	fmt.Println(artics[0].Content)
}
