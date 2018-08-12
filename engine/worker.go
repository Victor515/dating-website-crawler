package engine

import (
	"crawler/fetcher"
	"log"
)

func Worker(r Request) (ParserResult, error){
	// fetch text
	//log.Printf("Fetching Url: %s", r.Url)
	contents, err := fetcher.Fetch(r.Url)
	if err != nil{
		log.Printf("fetching error " +
			"fetching url: %s, %v", r.Url, err)
		return ParserResult{}, err
	}
	// feed to parser
	return  r.Parser.Parse(contents, r.Url), nil

}
