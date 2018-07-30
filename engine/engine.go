package engine

import (
	"crawler/fetcher"
	"log"
)

func Run (seeds ...Request){
	var requests []Request
	for _, r := range seeds{
		requests = append(requests, r)
	}

	for len(requests) > 0{
		r := requests[0]
		requests = requests[1:]

		// fetch text
		log.Printf("Fetching Url: %s", r.Url)
		contents, err := fetcher.Fetch(r.Url)
		if err != nil{
			log.Printf("fetching error " +
				"fetching url: %s, %v", r.Url, err)
			continue
		}

		// feed to parser
		result := r.ParserFunc(contents)

		// process result 1: add new requests to request list
		requests = append(requests, result.Requests...)
		// process result 2: print items
		for _, item := range result.Items{
			log.Printf("Got item: %v", item)
		}

	}


}