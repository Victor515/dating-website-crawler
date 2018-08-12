package engine

import (
	"log"
)

type SimpleEngine struct {}

func (e SimpleEngine) Run (seeds ...Request){
	var requests []Request
	for _, r := range seeds{
		requests = append(requests, r)
	}

	for len(requests) > 0{
		r := requests[0]
		requests = requests[1:]

		// workers complete the request
		result, err := Worker(r)
		if err != nil{
			continue
		}

		// process result 1: add new requests to request list
		requests = append(requests, result.Requests...)
		// process result 2: print items
		for _, item := range result.Items{
			log.Printf("Got item: %v", item)
		}

	}


}

