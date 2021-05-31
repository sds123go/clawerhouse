package engine

import "fmt"

type Scheduler interface {
	Submit(r Request)
	ConfigureWorkerChan(chan Request)
}
type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkCount int
}

func (c *ConcurrentEngine) Run(seeds ...Request) {
	for _, seed := range seeds {
		c.Scheduler.Submit(seed)
	}
	in := make(chan Request)
	out := make(chan ParseResult)
	c.Scheduler.ConfigureWorkerChan(in)
	for i := 0; i < c.WorkCount; i++ {
		createWork(in, out)
	}

	for {
		result := <-out
		for _, item := range result.Items {
			fmt.Printf("get item :", item)
		}
		for _, request := range result.Request {
			c.Scheduler.Submit(request)
		}
	}
}
func createWork(in chan Request, out chan ParseResult) {
	go func() {
		for {
			request := <-in
			result, err := SimpleEngine{}.worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
