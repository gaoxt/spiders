package engine

import (
	"distributedCrawfer/fetcher"
	"log"
)

func Worker(r Request) (ParserResult, error) {
	var body []byte
	if r.Url != "" {
		var err error
		log.Printf("Fetching %s %s", r.Url, r.PostData)
		body, err = fetcher.FetchPost(r.Url, r.PostData)
		if err != nil {
			log.Printf("Fetcher: error"+"fetching url %s: %v", r.Url, err)
			return ParserResult{}, err
		}
	}
	return r.Parser.Parse(body, r.Url), nil
}
