package persist

import (
	"testing"
	"crawler/model"
	"gopkg.in/olivere/elastic.v5"
	"context"
	"encoding/json"
	"crawler/engine"
)

func TestSave(t *testing.T) {
	expected := engine.Item{
		Url:     "http://album.zhenai.com/u/108906739",
		Type:    "zhenai",
		Id:      "108906739",
		Payload: model.Profile{
			Age: 34,
			Height: 162,
			Weight: 57,
			Income: "3001-5000元",
			Gender: "女",
			Name: "安静的雪",
			Xinzuo: "牡羊座",
			Occupation: "人事/行政",
			Marriage: "离异",
			House: "已购房",
			Hukou: "山东菏泽",
			Education: "大学本科",
			Car: "未购车",
		},
	}

	// save item
	err := save(expected)

	if err != nil{
		panic(err)
	}


	// TODO: Try to start up elasticsearch in go client
	client, err := elastic.NewClient(
		elastic.SetSniff(false),
	)

	if err != nil{
		panic(err)
	}

	// get item
	result, err := client.Get().
		Index("dating_profile").
		Type(expected.Type).
		Id(expected.Id).
		Do(context.Background())


	if err != nil{
		panic(err)
	}

	// json decoder
	var actual engine.Item
	json.Unmarshal([]byte(*result.Source), &actual)

	actualProfile, _ := model.FromJsonObject(actual.Payload)
	actual.Payload = actualProfile

	// verify result
	if actual != expected{
		t.Errorf("Expected: +%v, but actual is: %+v", expected, actual)
	}
}
