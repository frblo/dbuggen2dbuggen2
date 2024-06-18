package parser

// Parser? I hardly know 'er!

import (
	"database/sql"
	"dbuggen2dbuggen2/lexer"
	"fmt"
	"log"
	"time"
)

func Parse(lexedIssues []lexer.Issue, lexedArticles map[int][]lexer.Article) []Issue {
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

func extractDate(filename string) time.Time {
	date, err := time.Parse("2006-01-02T15", fmt.Sprintf("%vT01", filename[0:10]))
	if err != nil {
		date = time.Now()
		log.Println(err)
	}

	return date
}

func averageDate(dates []time.Time, issueMonthString string) time.Time {
	sum := int64(0)
	for _, date := range dates {
		sum += date.Unix()
	}

	mean := sum / int64(len(dates))
	meanDate := time.Unix(mean, 0)

	issueMonth, err := time.Parse("Jan 2006", issueMonthString)
	if err != nil {
		log.Print(err)
		return meanDate
	}

	if meanDate.Month() != issueMonth.Month() || meanDate.Year() != issueMonth.Year() {
		zone, _ := time.LoadLocation("Europe/Stockholm")
		return time.Date(issueMonth.Year(), issueMonth.Month(), 0, 0, 0, 0, 0, zone)
	}

	return meanDate
}
