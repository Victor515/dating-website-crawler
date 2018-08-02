package engine

import "log"

type ConcurrentEngine struct{
	Scheduler Scheduler
	WorkerCount int
}

type Scheduler interface {
	Submit(request Request)
	ConfigureMasterWorkerChannel(chan Request)
	WorkerReady(chan Request)
	Run()
}

func (e *ConcurrentEngine) Run(seeds ...Request){

	out := make(chan ParserResult)
	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++{
		// create workers
		createWorker(out, e.Scheduler)
	}

	// submit request
	for _, seed := range seeds{
		e.Scheduler.Submit(seed)
	}

	itemCount := 0
	for{
		// receive the result
		result := <- out
		for _, item := range result.Items{
			log.Printf("Got Item#%d: %v", itemCount, item)
			itemCount++
		}

		// add new request to the scheduler
		for _, req := range result.Requests{
			e.Scheduler.Submit(req)
		}
	}
}

func createWorker(out chan ParserResult, s Scheduler){
	in := make(chan Request)
	go func() {
		for{
			// tell scheduler I'm ready!
			s.WorkerReady(in)
			request := <- in
			result, err := worker(request)
			if err != nil{
				continue
			}

			// send result to out channel
			out <- result
		}
	}()

}