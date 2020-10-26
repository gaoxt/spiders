package engine

type Request struct {
	Url        string
	ParserFunc func([]byte) ParserResult
	PostData   []byte
}

type ParserResult struct {
	Requests    []Request
	Items       []interface{}
	StorageFunc func(interface{}) StorageResult
}

type StorageResult struct {
	Items []interface{}
	Err   error
}

// NilParser 返回一个空的ParserResult
func NilParser() ParserResult {
	return ParserResult{}
}
