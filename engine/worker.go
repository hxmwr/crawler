package engine

import (
	"crawler/fetcher"
	"log"
)

func Worker(r Request) (ParseResult, error) {
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher: error fetching url %s: %v", r.Url, err)
	}
	return r.Parser.Parse(body, r.Url), nil
}