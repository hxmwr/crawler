package main

import (
	"crawler/engine"
	"crawler/newdun/parser"
	"crawler/persist"
	"crawler/scheduler"
	"crawler/worker/client"
)

func main() {
	itemChan, err := persist.ItemSaver()
	if err != nil {
		panic(err)
	}

	processor, err := client.CreateProcessor()
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      100,
		ItemChan:         itemChan,
		RequestProcessor: processor,
	}
	e.Run(engine.Request{
		Url:    "https://www.newdun.com/news/info/company",
		Parser: engine.NewFuncParser(parser.ParseNewsList, "ParseNewsList"),
	})
}
