package engine

type ConcurrentEngine struct{
	Scheduler Scheduler
	WorkerCount int
	ItemChan    chan interface{}
}

type Scheduler interface {
	ReadyNotifier
	Submit(request Request)
	WorkerChan() chan Request
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request){

	out := make(chan ParserResult)
	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++{
		// create workers
		in := e.Scheduler.WorkerChan()
		createWorker(in, out, e.Scheduler)
	}

	// submit request
	for _, seed := range seeds{
		if(isDup(seed.Url)){
			//log.Printf("Duplicate request: %s", seed.Url)
			continue
		}
		e.Scheduler.Submit(seed)
	}

	for{
		// receive the result
		result := <- out
		for _, item := range result.Items{
			go func(item interface{}) {
				e.ItemChan <- item
			}(item) // need to pass item to the goroutine
		}


		// add new request to the scheduler
		for _, req := range result.Requests{
			// remove URL duplicates
			if(isDup(req.Url)){
				//log.Printf("Duplicate request: %s", req.Url)
				continue
			}
			e.Scheduler.Submit(req)
		}
	}
}

func createWorker(in chan Request, out chan ParserResult, ready ReadyNotifier){
	go func() {
		for{
			// tell scheduler I'm ready!
			ready.WorkerReady(in)
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

var visitedURL = make(map[string]bool)

func isDup(url string) bool{
	if(visitedURL[url]){
		return  true
	}else{
		visitedURL[url] = true
		return  false
	}
}