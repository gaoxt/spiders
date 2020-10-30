package main

import (
	"distributedCrawfer/config"
	itemSaver "distributedCrawfer/distributed/dao/client"
	"distributedCrawfer/distributed/rpcsupport"
	worker "distributedCrawfer/distributed/worker/client"
	"distributedCrawfer/engine"
	"distributedCrawfer/laomassf/parser"
	"distributedCrawfer/scheduler"
	"flag"
	"log"
	"net/rpc"
	"strings"
)

var (
	itemSaverHost = flag.String("itemsaver_host", "", "itemsaver host")
	wokerhosts    = flag.String("worker_hosts", "", "worker hosts (comma seprated)")
)

func main() {

	flag.Parse()

	itemChan, err := itemSaver.ItemSaver(*itemSaverHost)
	if err != nil {
		panic(err)
	}

	pool := createClientPool(strings.Split(*wokerhosts, ","))

	processor := worker.CreateProcessor(pool)

	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      10,
		ItemChan:         itemChan,
		RequestProcessor: processor,
	}

	e.Run(engine.Request{
		Parser: engine.NewFuncParser(parser.GetBookListPage, config.GetbookListPage),
	})
}

func createClientPool(hosts []string) chan *rpc.Client {
	var clients []*rpc.Client
	for _, h := range hosts {
		client, err := rpcsupport.NewClient(h)
		if err == nil {
			clients = append(clients, client)
			log.Printf("success connecting to %s", h)
		} else {
			log.Printf("error connecting to %s : %v", h, err)
		}
	}
	out := make(chan *rpc.Client)
	// 轮流分发
	go func() {
		// 始终轮流分发,多套一层for
		for {
			for _, client := range clients {
				out <- client
			}
		}

	}()
	return out
}
