package engine

import (
	"errors"
	"log"
	"reflect"
	"singleTaskCrawfer/fetcher"
)

func Run(seeds ...Request) error {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		var body []byte
		if r.Url != "" {
			var err error
			log.Printf("Fetching %s %s", r.Url, r.PostData)
			body, err = fetcher.FetchPost(r.Url, r.PostData)
			if err != nil {
				log.Printf("Fetcher: error"+"fetching url %s: %v", r.Url, err)
				return err
			}
		}
		parserResult := r.ParserFunc(body)
		if parserResult.isParserResultEmpty() {
			return errors.New("url:" + r.Url + string(r.PostData) + " ParserResult is empty!")
		}
		requests = append(requests, parserResult.Requests...)
		for _, item := range parserResult.Items {
			log.Printf("got item %v", item)
			if parserResult.StorageFunc != nil {
				storageResult := parserResult.StorageFunc(item)
				if storageResult.Err != nil {
					log.Printf("Storage Result: error %v", storageResult.Err)
				}
				log.Printf("Storage Insert success: %v", storageResult.Items)
			}
		}

	}
	return nil
}

func (x ParserResult) isParserResultEmpty() bool {
	return reflect.DeepEqual(x, ParserResult{})
}
