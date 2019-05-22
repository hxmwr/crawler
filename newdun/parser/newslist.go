package parser

import (
	"crawler/engine"
	"regexp"
)

const newsListRe = `<a href="(https://www.newdun.com/news/(\d+).html)" target="_blank">(.+?)</a>`

func ParseNewsList(contents []byte, _ string) engine.ParseResult {
	re := regexp.MustCompile(newsListRe)
	res := re.FindAllSubmatch(contents, -1)
	var requests []engine.Request
	for _, v := range res {
		requests = append(requests, engine.Request{
			Url:    string(v[1]),
			Parser: engine.NilParser{},
		})
		println(string(v[1]))
	}
	return engine.ParseResult{
		Requests: requests,
	}
}
