package parser

import (
	"distributedCrawfer/config"
	"distributedCrawfer/engine"
	"distributedCrawfer/helper"
	"distributedCrawfer/model"
	"encoding/json"
)

func ParseBookDetail(contents []byte, bookIndex model.BookIndex, url string) engine.ParserResult {
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
		Items: []engine.Item{
			{
				Url:     url,
				Type:    config.TypeName,
				Id:      bookIndex.Name,
				Payload: bookIndex,
			},
		},
	}

	return result
}

type ParseBookDetailFormat struct {
	bookIndex model.BookIndex
}

func (p *ParseBookDetailFormat) Parse(contents []byte, url string) engine.ParserResult {
	return ParseBookDetail(contents, p.bookIndex, url)
}

func (p *ParseBookDetailFormat) Serialized() (name string, args interface{}) {
	return config.ParseBookDetail, p.bookIndex
}

func NewParseBookDetailFormat(bookIndex model.BookIndex) *ParseBookDetailFormat {
	return &ParseBookDetailFormat{
		bookIndex: bookIndex,
	}
}
