package parser

import (
	"crawler/engine"
	"regexp"
)

const reString = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]+>([^<]+)</a>`

func ParseCityList(contents []byte) engine.ParserResult{
	re := regexp.MustCompile(reString)
	matches := re.FindAllSubmatch(contents, -1) // match contents

	result := engine.ParserResult{}

	limit := 10
	// generate ParserResult
	for _, match := range matches{

		// convert item to string before appending
		result.Items = append(result.Items, "City " + string(match[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url: string(match[1]),
			ParserFunc: ParseCity,
		})
		limit--
		if limit == 0{
			break
		}
	}

	return result
}
