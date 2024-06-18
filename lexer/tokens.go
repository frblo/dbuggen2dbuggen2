package lexer

type Issue struct {
	Number      int
	Name        string
	Description string
	Date        string
}

type yamlArticle struct {
	Title    string
	Category int
	Order    int
	Author   string
}

type Article struct {
	Ok       bool
	Filename string
	Title    string
	Issue    int
	Order    int
	Author   string
	Content  string
}
