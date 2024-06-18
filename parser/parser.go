package parser

// Parser? I hardly know 'er!

import (
	"dbuggen2dbuggen2/lexer"
	"fmt"
	"log"
	"time"
)

func Parse(lexedIssues []lexer.Issue, lexedArticles map[int][]lexer.Article) []Issue {
	return nil // TODO
}

func extractDate(filename string) time.Time {
	date, err := time.Parse("2006-01-02T15", fmt.Sprintf("%vT01", filename[0:10]))
	if err != nil {
		date = time.Now()
		log.Println(err)
	}

	return date
}
