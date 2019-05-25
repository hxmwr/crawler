package main

import (
	"crawler/rpcsupport"
	"crawler/worker"
	"flag"
	"fmt"
	"log"
)


var port = flag.Int("port", 0, "specify a port for me to listen on")

func main()  {
	flag.Parse()
	if *port == 0 {
		log.Printf("must specify a port")
		return
	}
	err := rpcsupport.ServeRpc(
		fmt.Sprintf(":%d", *port),
		worker.CrawlService{})
	if err != nil {
		panic(err)
	}
}
