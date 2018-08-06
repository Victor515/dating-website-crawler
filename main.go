package main

import (
	"crawler/engine"
	"crawler/scheduler"
	"crawler/persist"
	"crawler/zhenai/parser"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.QueueScheduler{},
		WorkerCount: 100,
		ItemChan: persist.ItemSaver(),
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
