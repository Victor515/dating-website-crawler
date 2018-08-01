package parser

import (
	"testing"
	"crawler/model"
	"crawler/fetcher"
)

func TestParseProfile2(t *testing.T) {
	contents, err := fetcher.Fetch("http://album.zhenai.com/u/1953772360")

	if err != nil{
		panic(err)
	}

	result := ParseProfile(contents, "找对象")

	if len(result.Items) != 1{
		t.Errorf("Items should have 1 item; but was %v\n", result.Items)
	}

	profile := result.Items[0].(model.Profile)

	expected := model.Profile{
		Age: 31,
		Height: 168,
		Weight: 0,
		Income: "12001-20000元",
		Gender: "男",
		Name: "找对象",
		Xinzuo: "天蝎座",
		Occupation: "--",
		Marriage: "未婚",
		House: "--",
		Hukou: "--",
		Education: "中专",
		Car: "未购车",
	}

	if profile != expected{
		t.Errorf("expected %v; but was %v", expected, profile)
	}
}
