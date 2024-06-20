package parser

// Parser? I hardly know 'er!

import (
	"database/sql"
	"dbuggen2dbuggen2/lexer"
	"log"
	"time"
)

func Parse(lexedIssues []lexer.Issue, lexedArticles map[int][]lexer.Article) []Issue {
	log.Println("Start parsing...")
	issues := make([]Issue, len(lexedIssues))
	issueCount := 0
	articleCount := 0
	for _, lexedIssue := range lexedIssues {
		if _, ok := lexedArticles[lexedIssue.Number]; !ok {
			continue
		}
		lexedArticles := lexedArticles[lexedIssue.Number]
		if len(lexedArticles) == 0 {
			continue
		}

		articles := makeArticles(lexedArticles, &articleCount)
		issues[issueCount] = makeIssue(lexedIssue, issueCount, articles)
		issueCount++
	}

	log.Println("Parsing complete")
	return issues
}

func makeIssue(lexedIssue lexer.Issue, id int, articles []Article) Issue {
	articleDates := make([]time.Time, len(articles))
	for i, article := range articles {
		articleDates[i] = article.LastEdited
	}
	publishingDate := averageDate(articleDates, lexedIssue.Date)

	null := sql.NullInt32{
		Int32: 0,
		Valid: false,
	}

	return Issue{
		ID:             id,
		Title:          lexedIssue.Name,
		PublishingDate: publishingDate,
		Pdf:            null,
		Html:           null,
		Coverpage:      null,
		Views:          0,
		Articles:       articles,
	}
}

func makeArticles(lexedArticles []lexer.Article, articleID *int) []Article {
	articles := make([]Article, len(lexedArticles))
	order := 0
	for i, article := range lexedArticles {
		articles[i] = makeArticle(article, *articleID, order)
		*articleID++
		order++
	}

	return articles
}

func makeArticle(lexedArticle lexer.Article, articleID int, index int) Article {
	date := extractDate(lexedArticle.Filename)
	nØllesafe := checkNaughtyness(lexedArticle.Content)
	return Article{
		ID:         articleID,
		Title:      lexedArticle.Title,
		AuthorText: lexedArticle.Author,
		IssueIndex: index,
		Content:    lexedArticle.Content,
		LastEdited: date,
		NØllesafe:  nØllesafe,
	}
}
