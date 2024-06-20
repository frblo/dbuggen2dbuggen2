package executor

// Execut...or? I don't know 'er... or?

import (
	"bufio"
	"database/sql"
	"dbuggen2dbuggen2/parser"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func Execute(issues []parser.Issue) {
	path := "dbuggen1data.psql"
	if _, err := os.Stat(path); err == nil {
		log.Printf("File %v already exists. Deleting and rewriting", path)
		os.Remove(path)
	}

	f, err := os.Create(path)

	if err != nil {
		log.Fatalf("Failed to create file, %v", err)
	}
	writer := bufio.NewWriter(f)

	for _, issue := range issues {
		writer.WriteString(fmt.Sprintf("-- Issue %v number %v\n", issue.Title, issue.ID))
		if err := writeIssue(writer, issue); err != nil {
			log.Printf("Failed to write issue %v. Deleting file.", issue.Title)
			os.Remove(path)
			return
		}
		for _, article := range issue.Articles {
			if err := writeArticle(writer, article, issue.ID); err != nil {
				log.Printf("Failed to write article %v. Deleting file.", article.Title)
				os.Remove(path)
				return
			}
			writer.WriteRune('\n')
		}
	}
}

func writeIssue(w *bufio.Writer, issue parser.Issue) error {
	code := fmt.Sprintf("INSERT INTO Archive.Issue VALUES (%v, %v, %v, %v, %v, %v, %v);\n",
		issue.ID,
		fmt.Sprintf(`$tag$%v$tag$`, issue.Title),
		issue.PublishingDate.Format(time.DateOnly),
		sqlNullInt32ToString(issue.Pdf),
		sqlNullInt32ToString(issue.Html),
		sqlNullInt32ToString(issue.Coverpage),
		issue.Views,
	)
	if _, err := w.WriteString(code); err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func writeArticle(w *bufio.Writer, article parser.Article, issueID int) error {
	code := fmt.Sprintf("INSERT INTO Archive.Article VALUES (%v, %v, %v, %v, %v, %v, %v, %v);\n",
		article.ID,
		escapeSqlString(article.Title),
		issueID,
		escapeSqlString(article.AuthorText),
		article.IssueIndex,
		escapeSqlString(article.Content),
		article.LastEdited.Format(time.DateOnly),
		article.NØllesafe,
	)
	if _, err := w.WriteString(code); err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func escapeSqlString(cont string) string {
	tagsymbol := []byte("!")
	tagbuilder := strings.Builder{}
	tagbuilder.Write(tagsymbol)

	tag := dollartag(tagbuilder.String())

	for {
		if !strings.Contains(cont, tag) {
			break
		}

		tagbuilder.Write(tagsymbol)
		tag = dollartag(tagbuilder.String())
	}

	return fmt.Sprintf("%v%v%v", tag, cont, tag)
}

func dollartag(str string) string {
	return fmt.Sprintf("$%v$", str)
}

func sqlNullInt32ToString(nullInt sql.NullInt32) string {
	if nullInt.Valid {
		return string(nullInt.Int32)
	}
	return "NULL"
}
