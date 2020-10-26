package main

import (
	"encoding/json"
	"fmt"
	"spiders/engine"
	"spiders/laomassf/parser"
)

func main() {
	seed()
}

func seed() {
	pageIndex, pageSize := 1, 15
	for {
		jsonStr, _ := json.Marshal(map[string]interface{}{"types": "0", "pageIndex": pageIndex, "pageSize": pageSize})
		err := engine.Run(engine.Request{
			Url:        "https://wx.laomassf.com/prointerface/MiniApp/Index.asmx/GetBookList",
			PostData:   jsonStr,
			ParserFunc: parser.ParseBookList,
		})
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		pageIndex++
	}
}
