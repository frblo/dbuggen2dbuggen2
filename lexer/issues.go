package lexer

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

func getIssues(path string) []Issue {
	yfile, err := os.ReadFile(fmt.Sprintf("%v/_data/issues.yaml", path))
	if err != nil {
		log.Fatal(err)
	}

	var rawIssues []Issue
	if err := yaml.Unmarshal(yfile, &rawIssues); err != nil {
		log.Fatal(err)
	}

	return rawIssues
}
