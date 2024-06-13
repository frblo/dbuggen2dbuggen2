package reading

import (
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
