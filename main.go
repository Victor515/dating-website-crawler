package main

import (
	"crawler/engine"
	"crawler/scheduler"
	"crawler/persist"
	"crawler/zhenai/parser"
)

func main() {
	itemChan, err := persist.ItemSaver("dating_profile")
	if err != nil{
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueueScheduler{},
		WorkerCount: 100,
		ItemChan:    itemChan,
	}
	//e := engine.SimpleEngine{}

	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})

	//e.Run(engine.Request{
	//	Url: "http://www.zhenai.com/zhenghun/aba",
	//	ParserFunc:parser.ParseCity,
	//})

}
