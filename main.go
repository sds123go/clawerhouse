package main

import (
	"crawler/engine"
	"crawler/julive/parser"
	"crawler/persist"
	"crawler/scheduler"
)

func main() {
	// var simpleEngine = &engine.SimpleEngine{}
	// httpText := "https://sh.julive.com/project/s"
	// simpleEngine.Run(engine.Request{
	// 	Url:        httpText,
	// 	ParserFunc: parser.ParseAreaList,
	// })

	httpText := "https://sh.julive.com/project/s"
	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.QueuedScheduler{},
		WorkCount: 10,
		ItemChan: persist.ItemSaver(),
	}
	e.Run(engine.Request{
		Url:        httpText,
		ParserFunc: parser.ParseAreaList,
	})

}
