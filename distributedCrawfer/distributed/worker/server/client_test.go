package main

import (
	"distributedCrawfer/config"
	"distributedCrawfer/distributed/rpcsupport"
	"distributedCrawfer/distributed/worker"
	"distributedCrawfer/model"
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func TestCrawlService(t *testing.T) {
	const host = ":9000"
	go rpcsupport.ServeRpc(host, worker.CrawlService{})
	time.Sleep(time.Second)

	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	bookIndex := model.BookIndex{
		ID:         245,
		Name:       "逃不开的经济周期",
		Author:     "[挪威] 拉斯.特维德",
		HomeImg:    "https://wx.laomassf.com/upload/img/bookImg/20200426172229_309817792359.jpg",
		Abstract:   "一部跨越300年的经济学演进史",
		PayPrice:   9.99,
		CreateDate: "2020-04-26 17:23:00",
		Detail:     nil,
	}
	jsonStr, _ := json.Marshal(map[string]int{"bookId": 245})
	req := worker.Request{
		Url:      "https://wx.laomassf.com/prointerface/MiniApp/Index.asmx/GetAudioList",
		PostData: jsonStr,
		Parser: worker.SerializedParser{
			Name: config.ParseBookDetail,
			Args: bookIndex,
		},
	}
	if bookIndexT, ok := req.Parser.Args.(model.BookIndex); ok {
		fmt.Printf("1 %v", bookIndexT)
	} else {
		fmt.Printf("2 %v", bookIndexT)
	}
	var result worker.ParseResult
	err = client.Call(config.CrawlServiceFunName, req, &result)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(result)
	}
}
