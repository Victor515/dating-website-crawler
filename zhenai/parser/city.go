package parser

import (
	"crawler/engine"
	"regexp"
)

var (
	profileRE = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
	cityURL = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/[^"]+)"`)
	)

func ParseCity(contents []byte) engine.ParserResult{
	matches := profileRE.FindAllSubmatch(contents, -1) // match contents

	result := engine.ParserResult{}
	// generate ParserResult
	for _, match := range matches{
		name := string(match[2])

		// convert item to string before appending
		result.Requests = append(result.Requests, engine.Request{
			Url: string(match[1]),
			ParserFunc: func(c []byte) engine.ParserResult {
				return ParseProfile(c, name) // function is first-class citizen...
			},
		})

	}

	matches = cityURL.FindAllSubmatch(contents, -1)

	for _, match := range matches{
		result.Requests = append(result.Requests, engine.Request{
			Url:string(match[1]),
			ParserFunc: ParseCity,
		})
	}

	return result
}
