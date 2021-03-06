package parser

import (
	"testing"
	"io/ioutil"
)

func TestParseCityList(t *testing.T) {
	contents, err := ioutil.ReadFile("citylist_test_data.html")

	if err != nil{
		panic(err)
	}

	result := ParseCityList(contents)

	const resultSize = 470
	var expectedUrls = []string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}


	if len(result.Requests) != resultSize{
		t.Errorf("Result should have %d results; but have %d", resultSize, len(result.Requests))
	}

	for i, url := range expectedUrls{
		if result.Requests[i].Url != url{
			t.Errorf("expected url #%d: %s; but get %s", i, url, result.Requests[i].Url)
		}
	}
}
