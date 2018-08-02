package scheduler

import "crawler/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) ConfigureMasterWorkerChannel(c chan engine.Request) {
	s.workerChan = c
}

func (s *SimpleScheduler) Submit(request engine.Request) {
	go func() {
		// send request down to worker chan
		s.workerChan <- request
	}()
}



