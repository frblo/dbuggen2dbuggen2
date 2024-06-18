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

type Article struct {
	ID         int
	Title      string
	AuthorText string
	IssueIndex int
	Content    string
	LastEdited time.Time
	NØllesafe  bool
}
