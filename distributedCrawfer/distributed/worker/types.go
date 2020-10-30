package worker

import (
	"distributedCrawfer/config"
	"distributedCrawfer/engine"
	"distributedCrawfer/laomassf/parser"
	"distributedCrawfer/model"
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

type SerializedParser struct {
	Name string
	Args interface{}
}

// {"ParseCityList",nil},{"ProfileParser", username}

type Request struct {
	Url      string
	Parser   SerializedParser
	PostData []byte
}

type ParseResult struct {
	Items   []engine.Item
	Request []Request
}

func SerializeRequest(r engine.Request) Request {
	name, args := r.Parser.Serialized()
	return Request{
		Url:      r.Url,
		PostData: r.PostData,
		Parser: SerializedParser{
			Name: name,
			Args: args,
		},
	}
}

func SerializeResult(r engine.ParserResult) ParseResult {
	result := ParseResult{
		Items: r.Items,
	}

	for _, req := range r.Requests {
		result.Request = append(result.Request, SerializeRequest(req))
	}
	return result
}

func DeserializeRequest(r Request) (engine.Request, error) {
	parser, err := deserializeParser(r.Parser)
	if err != nil {
		return engine.Request{}, err
	}
	return engine.Request{
		Url:      r.Url,
		PostData: r.PostData,
		Parser:   parser,
	}, nil
}

func DeserializeResult(r ParseResult) (engine.ParserResult, error) {
	result := engine.ParserResult{
		Items: r.Items,
	}
	for _, req := range r.Request {
		engineReq, err := DeserializeRequest(req)
		if err != nil {
			log.Printf("error deserializing"+"request: %v", err)
			continue
		}
		result.Requests = append(result.Requests, engineReq)
	}
	return result, nil
}

func deserializeParser(p SerializedParser) (engine.Parser, error) {
	switch p.Name {
	case config.GetbookListPage:
		return engine.NewFuncParser(parser.GetBookListPage, config.GetbookListPage), nil
	case config.ParseBookList:
		return engine.NewFuncParser(parser.ParseBookList, config.ParseBookList), nil
	case config.NilParser:
		return engine.NilParser{}, nil
	case config.ParseBookDetail:
		if jsonString, ok := json.Marshal(p.Args); ok == nil {
			bookIndex := model.BookIndex{}
			json.Unmarshal(jsonString, &bookIndex)
			return parser.NewParseBookDetailFormat(bookIndex), nil
		} else {
			return nil, fmt.Errorf("invalid"+"args: %v", p.Args)
		}
	default:
		return nil, errors.New("unDefine Method")
	}
}
