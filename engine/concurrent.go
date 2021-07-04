package engine

import "log"

type Scheduler interface {
	Submit(r Request)
	ConfigureWorkerChan(chan Request)
	Run()
	WorkReady(chan Request)
}
type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkCount int
}

func (c *ConcurrentEngine) Run(seeds ...Request) {

	out := make(chan ParseResult)
	c.Scheduler.Run()
	for i := 0; i < c.WorkCount; i++ {
		createWork(out, c.Scheduler)
	}
	for _, seed := range seeds {
		c.Scheduler.Submit(seed)
	}
	for {
		result := <-out
		for _, item := range result.Items {
			log.Printf("get item %v:", item)
		}
		for _, request := range result.Requests {
			c.Scheduler.Submit(request)
		}
	}
}
func createWork(out chan ParseResult, s Scheduler) {
	in := make(chan Request)
	go func() {
		for {
			s.WorkReady(in)
			request := <-in
			result, err := Worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
