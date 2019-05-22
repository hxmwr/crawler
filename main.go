package main

import (
	"crawler/engine"
	"crawler/newdun/parser"
	"crawler/worker/client"
)

func main() {
	itemChan, err := persist.ItemSaver("dating_profile")
	if err != nil {
		panic(err)
	}

	processor, err := client.CreateProcessor()
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:        &engine.Scheduler(),
		WorkerCount:      100,
		ItemChan:         itemChan,
		RequestProcessor: processor,
	}
	e.Run(engine.Request{
		Url:    "https://www.newdun.com/news/info/company",
		Parser: engine.NewFuncParser(parser.ParseNewsList, "ParseNewsList"),
	})
}
