package main

import (
	"crawler/persist"
	"crawler/rpcsupport"
	"github.com/olivere/elastic"
)

func main() {
	err := serveRpc(":1234", "newdun")
	if err != nil {
		panic(err)
	}
}

func serveRpc(host, index string) error {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	return rpcsupport.ServeRpc(host, &persist.ItemSaverService{
		Client: client,
		Index:  index,
	})
}
