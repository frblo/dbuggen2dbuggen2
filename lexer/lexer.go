package lexer

import "log"

// Lexer? I barely know 'er

func Lex(path string) ([]Issue, map[int][]Article) {
	log.Println("Start lexing")
	issues := getIssues(path)
	articles := getArticles(path)
	log.Println("Lexing complete")
	return issues, articles
}
