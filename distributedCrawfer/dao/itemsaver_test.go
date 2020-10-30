package dao

import (
	"context"
	"distributedCrawfer/config"
	"distributedCrawfer/engine"
	"distributedCrawfer/model"
	"encoding/json"
	"reflect"
	"testing"

	"gopkg.in/olivere/elastic.v5"
)

func TestItemSaver(t *testing.T) {
	expected := engine.Item{
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

	// TODO : Try to start up elastic search.  here using docker go client
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	const index = "laoma_test"
	// Save expected item
	err = Save(index, client, expected)
	if err != nil {
		panic(err)
	}

	// fetch saved item
	resp, err := client.Get().Index(index).Type(expected.Type).Id(expected.Id).Do(context.Background())
	if err != nil {
		panic(err)
	}
	t.Logf("%s", resp.Source)

	var actual engine.Item

	json.Unmarshal(*resp.Source, &actual)

	actualBookInfo, _ := model.FromJsonObj(actual.Payload)
	actual.Payload = actualBookInfo

	// verify saved item
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("got item : %v,expected : %v", actual, expected)
	}
}
