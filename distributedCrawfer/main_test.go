package main

import (
	"distributedCrawfer/config"
	itemSaver "distributedCrawfer/distributed/dao/client"
	worker "distributedCrawfer/distributed/worker/client"
	"distributedCrawfer/engine"
	"distributedCrawfer/laomassf/parser"
	"distributedCrawfer/scheduler"
	"strings"
	"testing"
)

func TestMain(t *testing.T) {
	// var Args interface{}
	// Args = model.BookIndex{
	// 	ID:         245,
	// 	Name:       "逃不开的经济周期",
	// 	Author:     "[挪威] 拉斯.特维德",
	// 	HomeImg:    "https://wx.laomassf.com/upload/img/bookImg/20200426172229_309817792359.jpg",
	// 	Abstract:   "一部跨越300年的经济学演进史",
	// 	PayPrice:   9.99,
	// 	CreateDate: "2020-04-26 17:23:00",
	// 	Detail:     nil,
	// }
	// if bookIndex, ok := Args.(model.BookIndex); ok {
	// 	fmt.Printf("1 %v", bookIndex)
	// } else {
	// 	fmt.Printf("2 %v", bookIndex)
	// }
	// return

	itemChan, err := itemSaver.ItemSaver(":1234")
	if err != nil {
		panic(err)
	}

	pool := createClientPool(strings.Split(":9001", ","))

	processor := worker.CreateProcessor(pool)

	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      2,
		ItemChan:         itemChan,
		RequestProcessor: processor,
	}

	e.Run(engine.Request{
		Parser: engine.NewFuncParser(parser.GetBookListPage, config.GetbookListPage),
	})
}
