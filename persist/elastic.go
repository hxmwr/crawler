package persist

import (
	"context"
	"crawler/engine"
	"crawler/rpcsupport"
	"fmt"
	"github.com/olivere/elastic"
	"github.com/pkg/errors"
	"log"
)

func ItemSaver(host string) (chan engine.Item, error) {
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		return nil, err
	}
	out := make(chan engine.Item)
	go func() {
		for {
			item := <-out
			fmt.Printf("Got Item: %v", item)
			result := ""
			err := client.Call("ItemSaverService.Save", item, &result)
			if err != nil {
				log.Printf("Item saver: error saving item %v: %v", item , err)
			}
		}
	}()
	return out, nil
}

func Save(client *elastic.Client, index string, item engine.Item) error {
	if item.Type == "" {
		return errors.New("must supply type")
	}

	indexService := client.Index().Index(index).Type(item.Type).BodyJson(item)

	_, err := indexService.Do(context.Background())

	if err != nil {
		return err
	}
	return nil
}
