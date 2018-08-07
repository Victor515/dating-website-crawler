package parser

import (
	"testing"
	"io/ioutil"
	"crawler/model"
	"crawler/engine"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_test_data.html")

	if err != nil{
		panic(err)
	}

	result := ParseProfile(contents, "http://album.zhenai.com/u/108906739", "安静的雪")

	if len(result.Items) != 1{
		t.Errorf("Items should have 1 item; but was %v\n", result.Items)
	}

	actual := result.Items[0]

	expected := engine.Item{
		Url:     "http://album.zhenai.com/u/108906739",
		Type:    "zhenai",
		Id:      "108906739",
		Payload: model.Profile{
			Age: 35,
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

	if actual != expected{
		t.Errorf("expected %v; but was %v", expected, actual)
	}
}
