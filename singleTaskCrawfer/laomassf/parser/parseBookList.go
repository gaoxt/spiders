package parser

import (
	"encoding/json"
	"singleTaskCrawfer/engine"
	"singleTaskCrawfer/helper"
	"singleTaskCrawfer/model"
	"strconv"
)

func ParseBookList(contents []byte) engine.ParserResult {

	firstData := helper.ParserData(string(contents))
	secondData := helper.ParserData(firstData["d"])
	var wxBookListObj []wxBookList
	if secondData["Data"] == nil {
		return engine.NilParser()
	}
	_ = json.Unmarshal([]byte(secondData["Data"].(string)), &wxBookListObj)
	lenWxBookListObj := len(wxBookListObj)
	if lenWxBookListObj == 0 {
		return engine.NilParser()
	}
	result := engine.ParserResult{}
	for i := 0; i < lenWxBookListObj; i++ {
		bookIndex := model.BookIndex{}
		bookID := wxBookListObj[i].ID
		bookIndex.ID = bookID
		bookIndex.Name = wxBookListObj[i].Name
		bookIndex.Author = wxBookListObj[i].Author
		bookIndex.HomeImg = helper.UrlPathFormat(wxBookListObj[i].HomeImg)
		bookIndex.Abstract = wxBookListObj[i].Abstract
		bookIndex.PayPrice = wxBookListObj[i].PayPrice
		bookIndex.CreateDate = helper.CreateDateFormat(wxBookListObj[i].CreateDate)
		bookIndex.Detail = nil

		jsonStr, _ := json.Marshal(map[string]int{"bookId": bookID})
		result.Items = append(result.Items, "List "+strconv.Itoa(bookID)+" "+wxBookListObj[i].Name)
		result.Requests = append(result.Requests, engine.Request{
			Url:      "https://wx.laomassf.com/prointerface/MiniApp/Index.asmx/GetAudioList",
			PostData: jsonStr,
			ParserFunc: func(c []byte) engine.ParserResult {
				return ParseBookDetail(c, bookIndex)
			},
		})
	}

	return result
}
