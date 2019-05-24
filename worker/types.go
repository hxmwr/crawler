package worker

import (
	"crawler/config"
	"crawler/engine"
	"crawler/newdun/parser"
	"github.com/pkg/errors"
	"log"
)

type SerializedParser struct {
	Name string
	Args interface{}
}

type Request struct {
	Url    string
	Parser SerializedParser
}

func SerializeRequest(r engine.Request) Request {
	name, args := r.Parser.Serialize()
	return Request{
		Url: r.Url,
		Parser: SerializedParser{
			Name: name,
			Args: args,
		},
	}
}

type ParseResult struct {
	Items    []engine.Item
	Requests []Request
}

func SerializedResult(r engine.ParseResult) ParseResult {
	result := ParseResult{
		Items: r.Items,
	}

	for _, req := range r.Requests {
		result.Requests = append(result.Requests, SerializeRequest(req))
	}

	return result
}

func DeserializeResult(r ParseResult) (engine.ParseResult, error) {
	result := engine.ParseResult{
		Items: r.Items,
	}

	for _, req := range r.Requests {
		request, err := DeserializeRequest(req)
		if err != nil {
			log.Printf("error deseralizing request: %v", err)
			continue
		}
		result.Requests = append(result.Requests, request)
	}

	return result, nil
}

func deserializeParser(p SerializedParser) (engine.Parser, error) {
	switch p.Name {
	case config.ParseNewsList:
		return engine.NewFuncParser(parser.ParseNewsList, config.ParseNewsList), nil
	case config.NilParser:
		return engine.NilParser{}, nil
	default:
		return nil, errors.New("unknown parser name")
	}
}

func DeserializeRequest(r Request) (engine.Request, error) {
	p, err := deserializeParser(r.Parser)
	if err != nil {
		return engine.Request{}, err
	}
	return engine.Request{
		Url:    r.Url,
		Parser: p,
	}, nil
}
