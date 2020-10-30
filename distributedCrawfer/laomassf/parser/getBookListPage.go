package parser

import (
	"distributedCrawfer/config"
	"distributedCrawfer/engine"
	"encoding/json"
)

func GetBookListPage(contents []byte, _ string) engine.ParserResult {
	pageSize := 15
	result := engine.ParserResult{}
	for pageIndex := 2; pageIndex <= 3; pageIndex++ {
		jsonStr, _ := json.Marshal(map[string]interface{}{"types": "0", "pageIndex": pageIndex, "pageSize": pageSize})
		result.Requests = append(result.Requests, engine.Request{
			Url:      "https://wx.laomassf.com/prointerface/MiniApp/Index.asmx/GetBookList",
			PostData: jsonStr,
			Parser:   engine.NewFuncParser(ParseBookList, config.ParseBookList),
		})
	}
	return result
}
