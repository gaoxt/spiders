package parser

import (
	"concurrencyCrawfer/engine"
	"encoding/json"
)

func GetBookListPage(contents []byte, _ string) engine.ParserResult {
	pageSize := 15
	result := engine.ParserResult{}
	for pageIndex := 1; pageIndex <= 10; pageIndex++ {
		jsonStr, _ := json.Marshal(map[string]interface{}{"types": "0", "pageIndex": pageIndex, "pageSize": pageSize})
		result.Requests = append(result.Requests, engine.Request{
			Url:        "https://wx.laomassf.com/prointerface/MiniApp/Index.asmx/GetBookList",
			PostData:   jsonStr,
			ParserFunc: ParseBookList,
		})
	}
	return result
}
