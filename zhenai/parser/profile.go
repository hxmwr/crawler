package parser

import (
	"crawler/engine"
	"crawler/model"
	"regexp"
	"strconv"
)

var ageRe =regexp.MustCompile(`<td><span class=""></span></td>`)

func ParseProfile(contents []byte) engine.ParseResult {
	profille := model.Profile{}
	match := ageRe.FindSubmatch(contents)
	if match != nil {
		age, err := strconv.Atoi(string(match[1]))
		if err != nil {
			profille.Age = age
		}
	}
	return engine.ParseResult{}
}

func extractString(contents []byte, re *regexp.Regexp) string{
	//match := re.FindSubmatch(contents)
	return ""
}
