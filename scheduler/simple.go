package scheduler

import "crawler/engine"

type SimpleScheduler struct {
	WorkChan chan engine.Request
}

func (s *SimpleScheduler) Submit(r engine.Request) {
	go func() {
		s.WorkChan <- r
	}()
}
func (s *SimpleScheduler) ConfigureWorkerChan(c chan engine.Request) {
	s.WorkChan = c
}
