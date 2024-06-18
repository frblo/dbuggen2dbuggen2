package lexer

// Lexer? I barely know 'er

func Lex(path string) ([]Issue, map[int][]Article) {
	issues := getIssues(path)
	articles := getArticles(path)
	return issues, articles
}
