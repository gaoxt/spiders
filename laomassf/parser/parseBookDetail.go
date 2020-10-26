package parser

import (
	"encoding/json"
	"spiders/engine"
	"spiders/helper"
	"spiders/laomassf/storage"
	"spiders/model"
)

func ParseBookDetail(contents []byte, bookIndex model.BookIndex) engine.ParserResult {
	firstData := helper.ParserData(string(contents))
	secondData := helper.ParserData(firstData["d"])
	var wxBooksObj []wxBookDetail
	_ = json.Unmarshal([]byte(secondData["Data"].(string)), &wxBooksObj)
	var bookDetailObj = make([]model.BookDetail, len(wxBooksObj))
	for i := 0; i < len(wxBooksObj); i++ {
		bookDetailObj[i].Name = wxBooksObj[i].Name
		bookDetailObj[i].Title = wxBooksObj[i].Title
		bookDetailObj[i].HomeImg = helper.UrlPathFormat(wxBooksObj[i].HomeImg)
		bookDetailObj[i].AudioAbstract = wxBooksObj[i].AudioAbstract
		bookDetailObj[i].FileSize = wxBooksObj[i].FileSize
		bookDetailObj[i].FileDuration = wxBooksObj[i].FileDuration
		bookDetailObj[i].CreateDate = helper.CreateDateFormat(wxBooksObj[i].CreateDate)
		bookDetailObj[i].FilePath = helper.UrlPathFormat(wxBooksObj[i].FilePath)
	}
	bookIndex.Detail = bookDetailObj
	result := engine.ParserResult{
		Items: []interface{}{"Detail " + bookIndex.Name},
		StorageFunc: func(c interface{}) engine.StorageResult {
			return storage.BookInsert(bookIndex)
		},
	}
	return result
}
