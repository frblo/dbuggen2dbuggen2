package lexer

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"gopkg.in/yaml.v3"
)

func getArticles(path string) map[int][]Article {
	postsPath := fmt.Sprintf("%v/_posts", path)
	postsDir, err := os.Open(postsPath)
	if err != nil {
		log.Fatal(err)
	}

	posts, err := postsDir.Readdirnames(0)
	if err != nil {
		log.Fatal(err)
	}

	ch := make(chan Article)

	for _, p := range posts {
		go getArticle(postsPath, p, ch)
	}

	articles := make(map[int][]Article)
	for j := 0; j < len(posts); j++ {
		article := <-ch
		if !article.Ok {
			continue
		}
		if i, ok := articles[article.Issue]; ok {
			articles[article.Issue] = append(i, article)
		} else {
			articles[article.Issue] = []Article{article}
		}
	}

	for i, a := range articles {
		articles[i] = sortArticles(a)
	}

	return articles
}

func sortArticles(articles []Article) []Article {
	sort.Slice(articles, func(i, j int) bool {
		return articles[i].Order < articles[j].Order
	})
	return articles
}

var emptyArticle = Article{
	Ok:       false,
	Filename: "",
	Title:    "",
	Issue:    0,
	Order:    0,
	Author:   "",
	Content:  "",
}

func getArticle(repoPath string, filename string, ch chan Article) {
	filePath := fmt.Sprintf("%v/%v", repoPath, filename)
	f, err := os.Open(filePath)
	if err != nil {
		ch <- emptyArticle
		return
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	yamlBuilder := strings.Builder{}
	_, _, errr := reader.ReadLine()
	if errr != nil {
		ch <- emptyArticle
		return
	}

	// Before anyone starts talking smack about this not being lexing, but parsing, I
	// just want to say that I don't care.
	for s, _, err := reader.ReadLine(); string(s) != "---"; s, _, err = reader.ReadLine() {
		if err != nil {
			ch <- emptyArticle
			return
		}
		yamlBuilder.Write(s)
		yamlBuilder.Write([]byte("\n"))
	}
	var articleInfo yamlArticle
	if err := yaml.Unmarshal([]byte(yamlBuilder.String()), &articleInfo); err != nil {
		ch <- emptyArticle
		return
	}

	contentBuilder := strings.Builder{}
	N := 128
	partion := make([]byte, N)
	for n, err := reader.Read(partion); err == nil; n, err = reader.Read(partion) {
		if n != N {
			contentBuilder.Write(partion[:n])
			continue
		}
		contentBuilder.Write(partion)
	}

	content := strings.TrimSpace(contentBuilder.String())

	article := Article{
		Ok:       true,
		Filename: filename,
		Title:    articleInfo.Title,
		Issue:    articleInfo.Category,
		Order:    articleInfo.Order,
		Author:   articleInfo.Author,
		Content:  content,
	}
	ch <- article
}
