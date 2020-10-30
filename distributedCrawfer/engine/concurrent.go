package engine

type ConcurrentEngine struct {
	Scheduler        Scheduler
	WorkerCount      int
	ItemChan         chan Item
	RequestProcessor Processor
}
type Processor func(Request) (ParserResult, error)

type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkerChan() chan Request
	Run()
}

type ReadyNotifier interface {
	WorkReady(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	out := make(chan ParserResult)
	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		e.createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	for _, r := range seeds {
		if isDuplicate(r) {
			continue
		}
		e.Scheduler.Submit(r)
	}
	for {
		result := <-out
		for _, item := range result.Items {

			go func() {
				e.ItemChan <- item
			}()
		}

		for _, request := range result.Requests {
			if isDuplicate(request) {
				continue
			}
			e.Scheduler.Submit(request)
		}
	}
}

var visitedUrls = make(map[string]bool)

func isDuplicate(r Request) bool {
	flag := r.Url + string(r.PostData)
	if visitedUrls[flag] {
		return true
	}

	visitedUrls[flag] = true

	return false
}

func (e *ConcurrentEngine) createWorker(in chan Request, out chan ParserResult, ready ReadyNotifier) {
	go func() {
		for {
			ready.WorkReady(in)
			request := <-in
			result, err := e.RequestProcessor(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
