package main

import (
	"crawler/config"
	"crawler/rpcsupport"
	"crawler/worker"
	"fmt"
)

func main()  {
	err := rpcsupport.ServeRpc(fmt.Sprintf(":%d", config.WorkerPort0), worker.CrawlService{})
	if err != nil {
		panic(err)
	}
}
