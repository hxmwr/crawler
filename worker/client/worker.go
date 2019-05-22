package client

import (
	"crawler/config"
	"crawler/engine"
	"crawler/rpcsupport"
	"crawler/worker"
	"fmt"
)

func CreateProcessor() (engine.Processor, error) {
	client, err := rpcsupport.NewClient(fmt.Sprintf(":%d", config.WorkerPort0))
	if err != nil {
		return nil, err
	}

	return func(request engine.Request) (result engine.ParseResult, e error) {
		sReq :=  worker.SerializeRequest(request)
		var sResult worker.ParseResult
		err := client.Call(config.CralServiceRpc)
		if err != nil {
			return engine.ParseResult{}, err
		}
		return worker.DeserializeResult(sResult)
	}, nil
}
