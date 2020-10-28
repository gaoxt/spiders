package main

import (
	"concurrencyCrawfer/config"
	"concurrencyCrawfer/dao"
	"concurrencyCrawfer/engine"
	"concurrencyCrawfer/laomassf/parser"
	"concurrencyCrawfer/scheduler"
)

func main() {
	itemChan, err := dao.ItemSaver(config.IndexName)
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 10,
		ItemChan:    itemChan,
	}

	e.Run(engine.Request{
		ParserFunc: parser.GetBookListPage,
	})
}
