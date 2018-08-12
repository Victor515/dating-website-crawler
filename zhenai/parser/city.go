package parser

import (
	"crawler/engine"
	"regexp"
	"crawler/config"
)

var (
	profileRE = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
	cityURL = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/[^"]+)"`)
	)

func ParseCity(contents []byte, _ string) engine.ParserResult{
	matches := profileRE.FindAllSubmatch(contents, -1) // match contents

	result := engine.ParserResult{}
	// generate ParserResult
	for _, match := range matches{
		name := string(match[2])
		url := string(match[1])

		// convert item to string before appending

		result.Requests = append(result.Requests, engine.Request{
			Url: url,
			Parser: NewProfileParser(name),
		})
	}

	matches = cityURL.FindAllSubmatch(contents, -1)

	for _, match := range matches{
		result.Requests = append(result.Requests, engine.Request{
			Url:string(match[1]),
			Parser: engine.NewFuncParser(ParseCity, config.ParseCity),
		})
	}

	return result
}
