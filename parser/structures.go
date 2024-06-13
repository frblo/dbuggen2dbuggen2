package parser

import "time"

type YamlIssue struct {
	Number      int
	Name        string
	Description string
	Date        string
}

type mediumArticle struct {
	Ok        bool
	Date      time.Time
	Title     string
	Category  int
	Order     int
	Author    string
	NØllesafe bool
	Content   string
}

type rawArticle struct {
	Title    string
	Category int
	Order    int
	Author   string
}
