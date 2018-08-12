package parser

import (
	"crawler/engine"
	"regexp"
)

const reString = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]+>([^<]+)</a>`

func ParseCityList(contents []byte, _ string) engine.ParserResult{
	re := regexp.MustCompile(reString)
	matches := re.FindAllSubmatch(contents, -1) // match contents

	result := engine.ParserResult{}

	// generate ParserResult
	for _, match := range matches{
		//log.Printf("Got city: %s", match[2])
		// convert item to string before appending
		result.Requests = append(result.Requests, engine.Request{
			Url: string(match[1]),
			Parser: engine.NewFuncParser(ParseCity, "ParseCity"),
		})
	}

	return result
}
