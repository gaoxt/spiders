package engine

import (
	"concurrencyCrawfer/fetcher"
	"log"
)

func worker(r Request) (ParserResult, error) {
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
	return r.ParserFunc(body, r.Url), nil
}
