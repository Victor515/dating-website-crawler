package parser

import (
	"testing"
	"io/ioutil"
	"crawler/model"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_test_data.html")

	if err != nil{
		panic(err)
	}

	result := ParseProfile(contents, "一切随缘")

	if len(result.Items) != 1{
		t.Errorf("Items should have 1 item; but was %v\n", result.Items)
	}

	profile := result.Items[0].(model.Profile)

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

	if profile != expected{
		t.Errorf("expected %v; but was %v", expected, profile)
	}
}
