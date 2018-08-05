package scheduler

import "crawler/engine"

type SimpleScheduler struct {
	// simple scheduler only has one channel for workers
	workerChan chan engine.Request
}

func (s *SimpleScheduler) WorkerChan() chan engine.Request {
	return s.workerChan
}

func (s *SimpleScheduler) WorkerReady(chan engine.Request) {
}

func (s *SimpleScheduler) Run() {
	s.workerChan = make(chan engine.Request)
}

func (s *SimpleScheduler) Submit(request engine.Request) {
	go func() {
		// open a goroutine for every request
		s.workerChan <- request
	}()
}



