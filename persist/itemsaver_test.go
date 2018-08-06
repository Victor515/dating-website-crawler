package persist

import (
	"testing"
	"crawler/model"
	"gopkg.in/olivere/elastic.v5"
	"context"
	"encoding/json"
)

func TestSave(t *testing.T) {
	expected := model.Profile{
		Age: 23,
		Height: 160,
		Weight: 48,
		Income: "8001-12000元",
		Gender: "女",
		Name: "一切随缘",
		Xinzuo: "狮子座",
		Occupation: "其他职业",
		Marriage: "未婚",
		House: "打算婚后购房",
		Hukou: "广东广州",
		Education: "中专",
		Car: "未购车",
	}

	id, err := save(expected)

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

	result, err := client.Get().
		Index("dating_profile").
		Type("zhenai").
		Id(id).
		Do(context.Background())


	if err != nil{
		panic(err)
	}

	// json decoder
	var actual model.Profile
	err = json.Unmarshal([]byte(*result.Source), &actual)

	if err != nil{
		panic(err)
	}

	if actual != expected{
		t.Errorf("Expected: +%v, but actual is: %+v", expected, actual)
	}
}
