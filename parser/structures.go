package parser

type YamlIssue struct {
	Number      int
	Name        string
	Description string
	Date        string
}

type RawArticle struct {
	Ok       bool
	Filename string
	Title    string
	Category int
	Order    int
	Author   string
	Content  string
}
