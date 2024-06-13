package parser

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

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

func GetArticles() []mediumArticle {
	postsDir, err := os.Open("dbuggen/_posts")
	if err != nil {
		log.Fatal(err)
	}

	posts, err := postsDir.Readdirnames(0)
	if err != nil {
		log.Fatal(err)
	}

	ch := make(chan mediumArticle)
	articles := make([]mediumArticle, len(posts))

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

var emptyArticle = mediumArticle{
	Ok:        false,
	Date:      time.Now(),
	Title:     "",
	Category:  0,
	Order:     0,
	Author:    "",
	NØllesafe: false,
	Content:   "",
}

func getArticle(filename string, ch chan mediumArticle) {
	path := fmt.Sprintf("dbuggen/_posts/%v", filename)
	f, err := os.Open(path)
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

	for s, _, err := reader.ReadLine(); string(s) != "---"; s, _, err = reader.ReadLine() {
		if err != nil {
			ch <- emptyArticle
			return
		}
		yamlBuilder.Write(s)
		yamlBuilder.Write([]byte("\n"))
	}
	var rA rawArticle
	if err := yaml.Unmarshal([]byte(yamlBuilder.String()), &rA); err != nil {
		ch <- emptyArticle
		return
	}

	// offset, err := f.Seek(0, io.SeekCurrent)
	// if err != nil {
	// 	ch <- emptyArticle
	// 	return
	// }

	conB, err := os.ReadFile(path)
	if err != nil {
		ch <- emptyArticle
		return
	}

	// cont := string(conB[offset:(len(conB) - 1)])
	cont := string(conB)

	article := mediumArticle{
		Ok:        true,
		Date:      time.Now(),
		Title:     rA.Title,
		Category:  rA.Category,
		Order:     rA.Order,
		Author:    rA.Author,
		NØllesafe: checkNaughtyness(cont),
		Content:   cont,
	}
	ch <- article
}
