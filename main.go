package main

import (
	"crawler/engine"
	"crawler/julive/parser"
)

func main() {
	var simpleEngine = &engine.SimpleEngine{}
	httpText := "https://sh.julive.com/project/s"
	simpleEngine.Run(engine.Request{
		Url:        httpText,
		ParserFunc: parser.ParseAreaList,
	})

	// httpText := "https://sh.julive.com/project/s"
	// e := engine.ConcurrentEngine{
	// 	Scheduler: &scheduler.SimpleScheduler{},
	// 	WorkCount: 10,
	// }
	// e.Run(engine.Request{
	// 	Url:        httpText,
	// 	ParserFunc: parser.ParseAreaList,
	// })

}
