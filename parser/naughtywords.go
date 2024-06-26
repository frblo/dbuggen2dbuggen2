package parser

import "strings"

var naughty = []string{
	"mottagningen",
	"dadderiet",
	"dadda",
	"nøllan",
	"drifveriet",
	"driveriet",
	"drifvare",
	"mottagare",
	"öfverdrif",
	"öfverdriv",
	"öfverdrifv",
	"n0llan",
	"ekonomeriet",
	"ekonomerist",
	"quisineriet",
}

func checkNaughtyness(str string) bool {
	for _, w := range naughty {
		if strings.Contains(strings.ToLower(str), w) {
			return false
		}
	}
	return true
}
