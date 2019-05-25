package main

import (
	"crawler/engine"
	"crawler/newdun/parser"
	"crawler/persist"
	"crawler/rpcsupport"
	"crawler/scheduler"
	"crawler/worker/client"
	"flag"
	"log"
	"net/rpc"
	"strings"
)

var workerHosts = flag.String("worker_hosts", "", "worker hosts (comma separated)")

func main() {
	flag.Parse()
	itemChan, err := persist.ItemSaver(":1234")
	if err != nil {
		panic(err)
	}


	pool := createClientPool(strings.Split(*workerHosts, ","))

	processor := client.CreateProcessor(pool)

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


func createClientPool(hosts []string) chan *rpc.Client {
	var clients []*rpc.Client
	for _, h := range hosts	{
		c,  err := rpcsupport.NewClient(h)
		if err == nil {
			clients = append(clients, c)
		} else {
			log.Printf("error connection to %s: %v", h, err)
		}
	}
	out := make(chan *rpc.Client)

	go func() {
		for {
			for _, c := range clients {
				out <- c
			}
		}
	}()

	return out
}