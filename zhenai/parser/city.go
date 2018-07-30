package parser

import (
	"crawler/engine"
	"regexp"
)

const cityRE = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

func ParseCity(contents []byte) engine.ParserResult{
	re := regexp.MustCompile(cityRE)
	matches := re.FindAllSubmatch(contents, -1) // match contents

	result := engine.ParserResult{}
	// generate ParserResult
	for _, match := range matches{

		// convert item to string before appending
		result.Items = append(result.Items, "User " + string(match[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url: string(match[1]),
			ParserFunc: engine.NilParser,
		})

	}

	return result
}
