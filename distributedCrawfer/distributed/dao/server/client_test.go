package main

import (
	"distributedCrawfer/config"
	"distributedCrawfer/distributed/rpcsupport"
	"distributedCrawfer/engine"
	"distributedCrawfer/model"
	"testing"
	"time"
)

func TestIteamSaver(t *testing.T) {
	const host = ":1234"

	go serveRpc(host, "test1")

	time.Sleep(time.Second)

	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}
	item := engine.Item{
		Url:  "https://wx.laomassf.com/245",
		Type: config.TypeName,
		Id:   "245",
		Payload: model.BookIndex{
			ID:         245,
			Name:       "逃不开的经济周期",
			Author:     "[挪威] 拉斯.特维德",
			HomeImg:    "https://wx.laomassf.com/upload/img/bookImg/20200426172229_309817792359.jpg",
			Abstract:   "一部跨越300年的经济学演进史",
			PayPrice:   9.99,
			CreateDate: "2020-04-26 17:23:00",
			Detail:     nil,
		},
	}
	result := ""
	err = client.Call("ItemSaverService.Save", item, &result)
	if err != nil || result != "OK" {
		t.Errorf("result: %s; err: %s", result, err)
	}
}
