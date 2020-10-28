package engine

type ParserFunc func(contents []byte, url string) ParserResult

type Request struct {
	Url        string
	ParserFunc ParserFunc
	PostData   []byte
}

type ParserResult struct {
	Requests []Request
	Items    []Item
}

type Item struct {
	Url     string
	Type    string
	Id      string
	Payload interface{}
}

// NilParser 返回一个空的ParserResult
func NilParser() ParserResult {
	return ParserResult{}
}
