package executor

import (
	"bufio"
	"database/sql"
	"dbuggen2dbuggen2/parser"
	"fmt"
	"log"
	"os"
	"time"
)

// Execut...or? I don't know 'er... or?

func Execute(issues []parser.Issue) {
	f, err := os.Create("dbuggen1data.psql")
	if err != nil {
		log.Fatalf("Failed to create file, %v", err)
	}
	writer := bufio.NewWriter(f)

	for _, issue := range issues {
		writeIssue(writer, issue)
		for _, article := range issue.Articles {
			writeArticle(writer, article, issue.ID)
		}
	}
}

func writeIssue(w *bufio.Writer, issue parser.Issue) {
	code := fmt.Sprintf("INSERT INTO Archive.Issue VALUES (%v, %v, %v, %v, %v, %v, %v);\n",
		issue.ID,
		fmt.Sprintf(`'%v'`, issue.Title),
		issue.PublishingDate.Format(time.DateOnly),
		sqlNullInt32ToString(issue.Pdf),
		sqlNullInt32ToString(issue.Html),
		sqlNullInt32ToString(issue.Coverpage),
		issue.Views,
	)
	if _, err := w.WriteString(code); err != nil {
		log.Fatal(err)
	}
}

func writeArticle(w *bufio.Writer, article parser.Article, issueID int) {
	code := fmt.Sprintf("INSERT INTO Archive.Article VALUES (%v, %v, %v, %v, %v, %v, %v, %v);\n",
		article.ID,
		fmt.Sprintf(`'%v'`, article.Title),
		issueID,
		fmt.Sprintf(`'%v'`, article.AuthorText),
		article.IssueIndex,
		fmt.Sprintf(`'%v'`, article.Content),
		article.LastEdited.Format(time.DateOnly),
		article.NØllesafe,
	)
	if _, err := w.WriteString(code); err != nil {
		log.Fatal(err)
	}
}

func sqlNullInt32ToString(nullInt sql.NullInt32) string {
	if nullInt.Valid {
		return string(nullInt.Int32)
	}
	return "NULL"
}
