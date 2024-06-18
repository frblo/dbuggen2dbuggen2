package parser

import (
	"database/sql"
	"time"
)

type Issue struct {
	ID             int
	Title          string
	PublishingDate time.Time
	Pdf            sql.NullInt32
	Html           sql.NullInt32
	Coverpage      sql.NullInt32
	Views          int
	Articles       []Article
}

// type mediumArticle struct {
// 	Ok        bool
// 	Date      time.Time
// 	Title     string
// 	Category  int
// 	Order     int
// 	Author    string
// 	NØllesafe bool
// 	Content   string
// }

type Article struct {
	ID         int
	Title      string
	Issue      int
	AuthorText string
	IssueIndex int
	Content    string
	LastEdited time.Time
	NØllesafe  bool
}
