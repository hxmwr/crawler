package persist

import (
	"crawler/engine"
	"fmt"
)

func ItemSaver() (chan engine.Item, error) {
	itemChan := make(chan engine.Item)
	go func() {
		item := <-itemChan
		fmt.Printf("Got Item: %v", item)
	}()
	return itemChan, nil
}
