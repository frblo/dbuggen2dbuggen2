package parser

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

func GetIssues() []YamlIssue {
	yfile, err := os.ReadFile("dbuggen/_data/issues.yaml")
	if err != nil {
		log.Fatal(err)
	}

	var issues []YamlIssue
	if err := yaml.Unmarshal(yfile, &issues); err != nil {
		log.Fatal(err)
	}

	return issues
}

func GetArticles() []RawArticle {
	postsDir, err := os.Open("dbuggen/_posts")
	if err != nil {
		log.Fatal(err)
	}

	posts, err := postsDir.Readdirnames(0)
	if err != nil {
		log.Fatal(err)
	}

	ch := make(chan RawArticle)
	articles := make([]RawArticle, len(posts))

	for _, p := range posts {
		go getArticle(p, ch)
	}

	for i := 0; i < len(posts); i++ {
		a := <-ch
		if a.Ok {
			articles[i] = a
		}
	}

	return articles
}

var emptyArticle = RawArticle{
	Ok:       false,
	Filename: "",
	Title:    "",
	Category: 0,
	Order:    0,
	Author:   "",
	Content:  "",
}

func getArticle(filename string, ch chan RawArticle) {
	f, err := os.Open(fmt.Sprintf("dbuggen/_posts/%v", filename))
	if err != nil {
		ch <- emptyArticle
	}
	defer f.Close()

	// reader := bufio.NewReader(f)

}
